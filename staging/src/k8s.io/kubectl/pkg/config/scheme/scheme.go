/*
Copyright 2025 The Kubernetes Authors.

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

package scheme

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubectl/pkg/config"
	"k8s.io/kubectl/pkg/config/v1alpha1"
	"k8s.io/kubectl/pkg/config/v1beta1"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()
	// StrictCodecs provides methods for retrieving codecs and serializers
	// for specific versions and content types.
	StrictCodecs = serializer.NewCodecFactory(Scheme, serializer.EnableStrict)
	// LenientCodecs provides methods for retrieving codecs and serializers
	// for specific versions and content types.
	LenientCodecs = serializer.NewCodecFactory(Scheme, serializer.DisableStrict)
)

func init() {
	AddToScheme(Scheme)
}

// AddToScheme registers the API group and adds types to a scheme
func AddToScheme(scheme *runtime.Scheme) {
	utilruntime.Must(config.AddToScheme(scheme))
	utilruntime.Must(v1beta1.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
}
