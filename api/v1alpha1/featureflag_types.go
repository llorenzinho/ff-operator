/*
Copyright 2023.

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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FeatureFlagSpec defines the desired state of FeatureFlag
type FeatureFlagSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Enabled bool `json:"enabled"`
	Status  bool `json:"status"`
}

// FeatureFlagStatus defines the observed state of FeatureFlag
type FeatureFlagStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Active bool `json:"active"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// FeatureFlag is the Schema for the featureflags API
type FeatureFlag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FeatureFlagSpec   `json:"spec,omitempty"`
	Status FeatureFlagStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FeatureFlagList contains a list of FeatureFlag
type FeatureFlagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FeatureFlag `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FeatureFlag{}, &FeatureFlagList{})
}
