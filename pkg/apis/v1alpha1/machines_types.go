package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=machines
// +kubebuilder:resource:shortName=machine
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Machine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              MachineSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            MachineStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type MachineStatus struct {
	Condition []ClusterCondition `json:"condition,omitempty" protobuf:"bytes,1,rep,name=condition"`
}

type MachineSpec struct {
	// +optional
	Attributes map[string]string `json:"attributes,omitempty" protobuf:"bytes,1,rep,name=attributes"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Machine `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
