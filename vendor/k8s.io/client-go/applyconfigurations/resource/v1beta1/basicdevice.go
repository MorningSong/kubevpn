/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	resourcev1beta1 "k8s.io/api/resource/v1beta1"
)

// BasicDeviceApplyConfiguration represents a declarative configuration of the BasicDevice type for use
// with apply.
type BasicDeviceApplyConfiguration struct {
	Attributes map[resourcev1beta1.QualifiedName]DeviceAttributeApplyConfiguration `json:"attributes,omitempty"`
	Capacity   map[resourcev1beta1.QualifiedName]DeviceCapacityApplyConfiguration  `json:"capacity,omitempty"`
}

// BasicDeviceApplyConfiguration constructs a declarative configuration of the BasicDevice type for use with
// apply.
func BasicDevice() *BasicDeviceApplyConfiguration {
	return &BasicDeviceApplyConfiguration{}
}

// WithAttributes puts the entries into the Attributes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Attributes field,
// overwriting an existing map entries in Attributes field with the same key.
func (b *BasicDeviceApplyConfiguration) WithAttributes(entries map[resourcev1beta1.QualifiedName]DeviceAttributeApplyConfiguration) *BasicDeviceApplyConfiguration {
	if b.Attributes == nil && len(entries) > 0 {
		b.Attributes = make(map[resourcev1beta1.QualifiedName]DeviceAttributeApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Attributes[k] = v
	}
	return b
}

// WithCapacity puts the entries into the Capacity field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Capacity field,
// overwriting an existing map entries in Capacity field with the same key.
func (b *BasicDeviceApplyConfiguration) WithCapacity(entries map[resourcev1beta1.QualifiedName]DeviceCapacityApplyConfiguration) *BasicDeviceApplyConfiguration {
	if b.Capacity == nil && len(entries) > 0 {
		b.Capacity = make(map[resourcev1beta1.QualifiedName]DeviceCapacityApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Capacity[k] = v
	}
	return b
}
