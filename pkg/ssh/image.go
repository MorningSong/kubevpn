package ssh

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/distribution/reference"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/image"
	"github.com/docker/cli/cli/flags"
	"github.com/docker/cli/cli/streams"
	"github.com/docker/cli/cli/trust"
	typesimage "github.com/docker/docker/api/types/image"
	registrytypes "github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/moby/term"
	"github.com/opencontainers/image-spec/specs-go/v1"
	"golang.org/x/crypto/ssh"

	plog "github.com/wencaiwulue/kubevpn/v2/pkg/log"
)

func GetClient() (*client.Client, *command.DockerCli, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, nil, err
	}
	var dockerCli *command.DockerCli
	dockerCli, err = command.NewDockerCli(command.WithAPIClient(cli))
	if err != nil {
		return nil, nil, err
	}
	err = dockerCli.Initialize(flags.NewClientOptions())
	if err != nil {
		return nil, nil, err
	}
	return cli, dockerCli, nil
}

// TransferImage
// 1) if not special ssh config, just pull image and tag and push
// 2) if special ssh config, pull image, tag image, save image and scp image to remote, load image and push
func TransferImage(ctx context.Context, conf *SshConfig, imageSource, imageTarget string, out io.Writer) error {
	client, cli, err := GetClient()
	if err != nil {
		plog.G(ctx).Errorf("Failed to get docker client: %v", err)
		return err
	}
	// todo add flags? or detect k8s node runtime ?
	platform := &v1.Platform{Architecture: "amd64", OS: "linux"}
	err = PullImage(ctx, platform, client, cli, imageSource, out)
	if err != nil {
		plog.G(ctx).Errorf("Failed to pull image: %v", err)
		return err
	}

	err = client.ImageTag(ctx, imageSource, imageTarget)
	if err != nil {
		plog.G(ctx).Errorf("Failed to tag image %s to %s: %v", imageSource, imageTarget, err)
		return err
	}

	// use it if sshConfig is not empty
	if conf.ConfigAlias == "" && conf.Addr == "" {
		var distributionRef reference.Named
		distributionRef, err = reference.ParseNormalizedNamed(imageTarget)
		if err != nil {
			plog.G(ctx).Errorf("Failed to parse image name %s: %v", imageTarget, err)
			return err
		}
		var imgRefAndAuth trust.ImageRefAndAuth
		imgRefAndAuth, err = trust.GetImageReferencesAndAuth(ctx, image.AuthResolver(cli), distributionRef.String())
		if err != nil {
			plog.G(ctx).Errorf("Failed to get image auth: %v", err)
			return err
		}
		var encodedAuth string
		encodedAuth, err = registrytypes.EncodeAuthConfig(*imgRefAndAuth.AuthConfig())
		if err != nil {
			plog.G(ctx).Errorf("Failed to encode auth config to base64: %v", err)
			return err
		}
		requestPrivilege := command.RegistryAuthenticationPrivilegedFunc(cli, imgRefAndAuth.RepoInfo().Index, "push")
		var readCloser io.ReadCloser
		readCloser, err = client.ImagePush(ctx, imageTarget, typesimage.PushOptions{
			RegistryAuth:  encodedAuth,
			PrivilegeFunc: requestPrivilege,
		})
		if err != nil {
			plog.G(ctx).Errorf("Failed to push image %s, err: %v", imageTarget, err)
			return err
		}
		defer readCloser.Close()
		if out == nil {
			_, out, _ = term.StdStreams()
		}
		outWarp := streams.NewOut(out)
		err = jsonmessage.DisplayJSONMessagesToStream(readCloser, outWarp, nil)
		if err != nil {
			plog.G(ctx).Errorf("Failed to display message, err: %v", err)
			return err
		}
		return nil
	}

	// transfer image to remote
	var sshClient *ssh.Client
	sshClient, err = DialSshRemote(ctx, conf, ctx.Done())
	if err != nil {
		return err
	}
	defer sshClient.Close()
	var responseReader io.ReadCloser
	responseReader, err = client.ImageSave(ctx, []string{imageTarget})
	if err != nil {
		plog.G(ctx).Errorf("Failed to save image %s: %v", imageTarget, err)
		return err
	}
	defer responseReader.Close()
	file, err := os.CreateTemp("", "*.tar")
	if err != nil {
		return err
	}
	plog.G(ctx).Infof("Saving image %s to temp file %s", imageTarget, file.Name())
	if _, err = io.Copy(file, responseReader); err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	defer os.Remove(file.Name())

	plog.G(ctx).Infof("Transferring image %s", imageTarget)
	filename := filepath.Base(file.Name())
	cmd := fmt.Sprintf(
		"(docker load -i ~/.kubevpn/%s && docker push %s) || (nerdctl image load -i ~/.kubevpn/%s && nerdctl image push %s)",
		filename, imageTarget,
		filename, imageTarget,
	)
	stdout := plog.G(ctx).Out
	err = SCPAndExec(ctx, stdout, stdout, sshClient, file.Name(), filename, []string{cmd}...)
	if err != nil {
		return err
	}
	plog.G(ctx).Infof("Loaded image: %s", imageTarget)
	return nil
}

// PullImage image.RunPull(ctx, c, image.PullOptions{})
func PullImage(ctx context.Context, platform *v1.Platform, cli *client.Client, dockerCli *command.DockerCli, img string, out io.Writer) error {
	var readCloser io.ReadCloser
	var plat string
	if platform != nil && platform.Architecture != "" && platform.OS != "" {
		plat = fmt.Sprintf("%s/%s", platform.OS, platform.Architecture)
	}
	distributionRef, err := reference.ParseNormalizedNamed(img)
	if err != nil {
		plog.G(ctx).Errorf("Failed to parse image name %s: %v", img, err)
		return err
	}
	var imgRefAndAuth trust.ImageRefAndAuth
	imgRefAndAuth, err = trust.GetImageReferencesAndAuth(ctx, image.AuthResolver(dockerCli), distributionRef.String())
	if err != nil {
		plog.G(ctx).Errorf("Failed to get image auth: %v", err)
		return err
	}
	var encodedAuth string
	encodedAuth, err = registrytypes.EncodeAuthConfig(*imgRefAndAuth.AuthConfig())
	if err != nil {
		plog.G(ctx).Errorf("Failed to encode auth config to base64: %v", err)
		return err
	}
	requestPrivilege := command.RegistryAuthenticationPrivilegedFunc(dockerCli, imgRefAndAuth.RepoInfo().Index, "pull")
	readCloser, err = cli.ImagePull(ctx, img, typesimage.PullOptions{
		All:           false,
		RegistryAuth:  encodedAuth,
		PrivilegeFunc: requestPrivilege,
		Platform:      plat,
	})
	if err != nil {
		plog.G(ctx).Errorf("Failed to pull image %s: %v", img, err)
		return err
	}
	defer readCloser.Close()
	if out == nil {
		_, out, _ = term.StdStreams()
	}
	outWarp := streams.NewOut(out)
	err = jsonmessage.DisplayJSONMessagesToStream(readCloser, outWarp, nil)
	if err != nil {
		plog.G(ctx).Errorf("Failed to display message, err: %v", err)
		return err
	}
	return nil
}
