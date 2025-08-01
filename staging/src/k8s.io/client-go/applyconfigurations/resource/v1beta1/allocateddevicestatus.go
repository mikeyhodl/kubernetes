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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// AllocatedDeviceStatusApplyConfiguration represents a declarative configuration of the AllocatedDeviceStatus type for use
// with apply.
type AllocatedDeviceStatusApplyConfiguration struct {
	Driver      *string                              `json:"driver,omitempty"`
	Pool        *string                              `json:"pool,omitempty"`
	Device      *string                              `json:"device,omitempty"`
	ShareID     *string                              `json:"shareID,omitempty"`
	Conditions  []v1.ConditionApplyConfiguration     `json:"conditions,omitempty"`
	Data        *runtime.RawExtension                `json:"data,omitempty"`
	NetworkData *NetworkDeviceDataApplyConfiguration `json:"networkData,omitempty"`
}

// AllocatedDeviceStatusApplyConfiguration constructs a declarative configuration of the AllocatedDeviceStatus type for use with
// apply.
func AllocatedDeviceStatus() *AllocatedDeviceStatusApplyConfiguration {
	return &AllocatedDeviceStatusApplyConfiguration{}
}

// WithDriver sets the Driver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Driver field is set to the value of the last call.
func (b *AllocatedDeviceStatusApplyConfiguration) WithDriver(value string) *AllocatedDeviceStatusApplyConfiguration {
	b.Driver = &value
	return b
}

// WithPool sets the Pool field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pool field is set to the value of the last call.
func (b *AllocatedDeviceStatusApplyConfiguration) WithPool(value string) *AllocatedDeviceStatusApplyConfiguration {
	b.Pool = &value
	return b
}

// WithDevice sets the Device field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Device field is set to the value of the last call.
func (b *AllocatedDeviceStatusApplyConfiguration) WithDevice(value string) *AllocatedDeviceStatusApplyConfiguration {
	b.Device = &value
	return b
}

// WithShareID sets the ShareID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ShareID field is set to the value of the last call.
func (b *AllocatedDeviceStatusApplyConfiguration) WithShareID(value string) *AllocatedDeviceStatusApplyConfiguration {
	b.ShareID = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *AllocatedDeviceStatusApplyConfiguration) WithConditions(values ...*v1.ConditionApplyConfiguration) *AllocatedDeviceStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithData sets the Data field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Data field is set to the value of the last call.
func (b *AllocatedDeviceStatusApplyConfiguration) WithData(value runtime.RawExtension) *AllocatedDeviceStatusApplyConfiguration {
	b.Data = &value
	return b
}

// WithNetworkData sets the NetworkData field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkData field is set to the value of the last call.
func (b *AllocatedDeviceStatusApplyConfiguration) WithNetworkData(value *NetworkDeviceDataApplyConfiguration) *AllocatedDeviceStatusApplyConfiguration {
	b.NetworkData = value
	return b
}
