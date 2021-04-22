/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// FileParameters are the configurable fields of a File.
type FileParameters struct {
	File string `json:"file"`
}

// FileObservation are the observable fields of a File.
type FileObservation struct {
	Status string `json:"status,omitempty"`
}

// A FileSpec defines the desired state of a File.
type FileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FileParameters `json:"forProvider"`
}

// A FileStatus represents the observed state of a File.
type FileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// A File is an example API type.
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FILE",type="date",JSONPath=".spec.providerConfigRef.name"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// Please replace `PROVIDER-NAME` with your actual provider name, like `aws`, `azure`, `gcp`, `alibaba`
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ssh}
type File struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FileSpec   `json:"spec"`
	Status FileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileList contains a list of File
type FileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []File `json:"items"`
}
