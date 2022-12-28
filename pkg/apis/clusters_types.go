package apis

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusters
// +kubebuilder:resource:shortName=cluster
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ClusterSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ClusterStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ClusterConditionStatus string

const (
	ClusterStatusOffline   ClusterConditionStatus = "offline"
	ClusterStatusDeploying ClusterConditionStatus = "deploying"
	ClusterStatusOnline    ClusterConditionStatus = "online"
)

type ClusterCondition struct {
	Type       string                 `json:"type,omitempty" protobuf:"bytes,1,opt,name=type"`
	Status     ClusterConditionStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status,casttype=ClusterConditionStatus"`
	CreatTime  *metav1.Time           `json:"creat_time,omitempty" protobuf:"bytes,3,opt,name=creat_time,json=creatTime"`
	UpdateTime *metav1.Time           `json:"update_time,omitempty" protobuf:"bytes,4,opt,name=update_time,json=updateTime"`
}

type ClusterStatus struct {
	Condition []ClusterCondition `json:"condition,omitempty" protobuf:"bytes,1,rep,name=condition"`
}

type ClusterSpec struct {
	// +optional
	Attributes map[string]string `json:"attributes,omitempty" protobuf:"bytes,1,rep,name=attributes"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Cluster `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
