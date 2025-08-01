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

package v1beta2

import (
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "k8s.io/client-go/applyconfigurations/internal"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// DaemonSetApplyConfiguration represents a declarative configuration of the DaemonSet type for use
// with apply.
type DaemonSetApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *DaemonSetSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *DaemonSetStatusApplyConfiguration `json:"status,omitempty"`
}

// DaemonSet constructs a declarative configuration of the DaemonSet type for use with
// apply.
func DaemonSet(name, namespace string) *DaemonSetApplyConfiguration {
	b := &DaemonSetApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("DaemonSet")
	b.WithAPIVersion("apps/v1beta2")
	return b
}

// ExtractDaemonSet extracts the applied configuration owned by fieldManager from
// daemonSet. If no managedFields are found in daemonSet for fieldManager, a
// DaemonSetApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// daemonSet must be a unmodified DaemonSet API object that was retrieved from the Kubernetes API.
// ExtractDaemonSet provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractDaemonSet(daemonSet *appsv1beta2.DaemonSet, fieldManager string) (*DaemonSetApplyConfiguration, error) {
	return extractDaemonSet(daemonSet, fieldManager, "")
}

// ExtractDaemonSetStatus is the same as ExtractDaemonSet except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractDaemonSetStatus(daemonSet *appsv1beta2.DaemonSet, fieldManager string) (*DaemonSetApplyConfiguration, error) {
	return extractDaemonSet(daemonSet, fieldManager, "status")
}

func extractDaemonSet(daemonSet *appsv1beta2.DaemonSet, fieldManager string, subresource string) (*DaemonSetApplyConfiguration, error) {
	b := &DaemonSetApplyConfiguration{}
	err := managedfields.ExtractInto(daemonSet, internal.Parser().Type("io.k8s.api.apps.v1beta2.DaemonSet"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(daemonSet.Name)
	b.WithNamespace(daemonSet.Namespace)

	b.WithKind("DaemonSet")
	b.WithAPIVersion("apps/v1beta2")
	return b, nil
}
func (b DaemonSetApplyConfiguration) IsApplyConfiguration() {}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithKind(value string) *DaemonSetApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithAPIVersion(value string) *DaemonSetApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithName(value string) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithGenerateName(value string) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithNamespace(value string) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithUID(value types.UID) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithResourceVersion(value string) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithGeneration(value int64) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithCreationTimestamp(value metav1.Time) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *DaemonSetApplyConfiguration) WithLabels(entries map[string]string) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *DaemonSetApplyConfiguration) WithAnnotations(entries map[string]string) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *DaemonSetApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *DaemonSetApplyConfiguration) WithFinalizers(values ...string) *DaemonSetApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *DaemonSetApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithSpec(value *DaemonSetSpecApplyConfiguration) *DaemonSetApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *DaemonSetApplyConfiguration) WithStatus(value *DaemonSetStatusApplyConfiguration) *DaemonSetApplyConfiguration {
	b.Status = value
	return b
}

// GetKind retrieves the value of the Kind field in the declarative configuration.
func (b *DaemonSetApplyConfiguration) GetKind() *string {
	return b.TypeMetaApplyConfiguration.Kind
}

// GetAPIVersion retrieves the value of the APIVersion field in the declarative configuration.
func (b *DaemonSetApplyConfiguration) GetAPIVersion() *string {
	return b.TypeMetaApplyConfiguration.APIVersion
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *DaemonSetApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}

// GetNamespace retrieves the value of the Namespace field in the declarative configuration.
func (b *DaemonSetApplyConfiguration) GetNamespace() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Namespace
}
