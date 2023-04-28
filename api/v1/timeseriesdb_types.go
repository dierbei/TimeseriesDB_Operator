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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TimeseriesDBSpec defines the desired state of TimeseriesDB
type TimeseriesDBSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	DBType   string `json:"dbType,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}

// TimeseriesDBStatus defines the observed state of TimeseriesDB
type TimeseriesDBStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// Root = true 注释表明 TimeseriesDB 对象是根对象，应该是 API 中的顶级 Kubernetes 资源。
// 状态注释表明状态字段应该被视为子资源，这意味着状态更新可以与主资源更新分开进行。

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TimeseriesDB is the Schema for the timeseriesdbs API
type TimeseriesDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TimeseriesDBSpec   `json:"spec,omitempty"`
	Status TimeseriesDBStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TimeseriesDBList contains a list of TimeseriesDB
type TimeseriesDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TimeseriesDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TimeseriesDB{}, &TimeseriesDBList{})
}
