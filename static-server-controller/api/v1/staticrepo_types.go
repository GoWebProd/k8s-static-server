/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StaticRepoSpec defines the desired state of StaticRepo.
type StaticRepoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Hostname string                  `json:"hostname"`
	Git      *StaticRepoSpecGit      `json:"git,omitempty"`
	Cmd      *StaticRepoSpecCmd      `json:"cmd,omitempty"`
	ImageDir *StaticRepoSpecImageDir `json:"imageDir,omitempty"`
}

type StaticRepoSpecGit struct {
	Repo   string `json:"repo"`
	Commit string `json:"commit"`
}

type StaticRepoSpecCmd struct {
	Image   string   `json:"image"`
	Command []string `json:"command"`
}

type StaticRepoSpecImageDir struct {
	Image string `json:"image"`
	Path  string `json:"path"`
}

// StaticRepoStatus defines the observed state of StaticRepo.
type StaticRepoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StaticRepo is the Schema for the staticrepoes API.
type StaticRepo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StaticRepoSpec   `json:"spec,omitempty"`
	Status StaticRepoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StaticRepoList contains a list of StaticRepo.
type StaticRepoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticRepo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StaticRepo{}, &StaticRepoList{})
}
