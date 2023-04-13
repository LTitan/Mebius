package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

type MachineSpec struct {
	// Whether the Machine can use Mebius to perform operations.
	// The default false indicates that the machine can accept operations on the Mebius server.
	// +optional
	Unschedulable bool `json:"unschedulable,omitempty" protobuf:"bytes,1,opt,name=unschedulable"`
	// Whether Machine can be clustered.
	// The default value is false.
	// +optional
	Unclusterable bool `json:"unclusterable,omitempty" protobuf:"bytes,2,opt,name=unclusterable"`
	// Refer to the taints mechanism of k8s.
	// +optional
	Taints []corev1.Taint `json:"taints,omitempty" protobuf:"bytes,3,rep,name=taints"`
	// Flags which cluster the machine expects to join.
	// The default value is null, meaning it does not want to join any cluster.
	// +optional
	DesiredCluster ClusterID `json:"desired_cluster,omitempty" protobuf:"bytes,4,opt,name=desired_cluster,casttype=ClusterID"`
	// Flag the role Machine expects to become in the Cluster.
	// The default is Vagrant, meaning it does not join the Cluster.
	// +optional
	DesiredRole MachineRole `json:"desired_role,omitempty" protobuf:"bytes,5,opt,name=desired_role,casttype=MachineRole"`

	// TODO: Related to the Resource to be added.
}

type MachineStatus struct {
	// Records the port of the agent.
	// +optional
	MebiusEndpoint MachineMebiusEndpoint `json:"mebius_endpoint,omitempty" protobuf:"bytes,1,opt,name=mebius_endpoint,casttype=MachineMebiusEndpoint"`
	// Records the ports of daemons enabled by the agent.
	// +optional
	DaemonEndpoints []MachineDaemonEndpoint `json:"daemon_endpoints,omitempty" protobuf:"bytes,2,rep,name=daemon_endpoints"`
	// Various stages of machine.
	// +optional
	Phase MachinePhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=MachinePhase"`
	// Conditions of the Machine (agent running properly, memory pressure free, disk pressure free, network reachable).
	// +optional
	Conditions []MachineCondition `json:"conditions,omitempty" protobuf:"bytes,4,rep,name=conditions"`
	// Information about the Machine system to be reported by the agent.
	// +optional
	MachineInfo MachineSystemInfo `json:"machine_info,omitempty" protobuf:"bytes,5,opt,name=machine_info,casttype=MachineSystemInfo"`
	// Machine address, a variety of types, temporarily on a LAN IP is almost.
	// TODO: If the Agent supports the IPIP tunnel, it can also use non-LAN IP addresses to remove Layer 2 reachability
	// +optional
	Addresses []MachineAddress `json:"addresses,omitempty" protobuf:"bytes,6,rep,name=addresses"`
	// The cluster that machine actually joins.
	// +optional
	OwnerCluster ClusterID `json:"owner_cluster,omitempty" protobuf:"bytes,7,opt,name=owner_cluster,casttype=ClusterID"`
	// The role of Machine in the Cluster.
	// If Machine does not join the Cluster, it is positioned as Vagrant.
	// +optional
	Role MachineRole `json:"machine_role,omitempty" protobuf:"bytes,8,opt,name=machine_role,casttype=MachineRole"`

	// TODO: Related to the Resource to be added.
}

type MachineCondition struct {
	Type   MachineConditionType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=MachineConditionType"`
	Status ConditionStatus      `json:"status,omitempty" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	// +optional
	LastHeartbeatTime metav1.Time `json:"last_heartbeat_time,omitempty" protobuf:"bytes,3,opt,name=last_heartbeat_time"`
	// +optional
	LastTransitionTime metav1.Time `json:"last_transition_time,omitempty" protobuf:"bytes,4,opt,name=last_transition_time"`
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

type MachineMebiusEndpoint struct {
	Port int32 `json:"port,omitempty" protobuf:"bytes,1,opt,name=port"`
}

type MachineDaemonEndpoint struct {
	Port int32 `json:"port,omitempty" protobuf:"bytes,1,opt,name=port"`
}

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

type MachinePhase string

const (
	MachinePending    MachinePhase = "Pending"
	MachineRunning    MachinePhase = "Running"
	MachineTerminated MachinePhase = "Terminated"
)

type MachineConditionType string

const (
	MachineClientReady      MachineConditionType = "ClientReady"
	MachineMemoryUnpressure MachineConditionType = "MemoryUnpressure"
	MachineDiskUnpressure   MachineConditionType = "DiskUnpressure"
	MachineNetworkAvailable MachineConditionType = "NetworkAvailable"
)

type MachineSystemInfo struct {
	// MachineID reported by the machine. For unique machine identification
	// in the cluster this field is preferred. Learn more from man(5)
	// machine-id: <http://man7.org/linux/man-pages/man5/machine-id.5.html>
	MachineID       string `json:"machine_id,omitempty" protobuf:"bytes,1,opt,name=machine_id"`
	KernelVersion   string `json:"kernel_version,omitempty" protobuf:"bytes,2,opt,name=kernel_version"`
	OperationSystem string `json:"operation_system,omitempty" protobuf:"bytes,3,opt,name=operation_system"`
	Architecture    string `json:"architecture,omitempty" protobuf:"bytes,4,opt,name=architecture"`
}

type MachineAddressType string

const (
	MachineHostName   MachineAddressType = "Hostname"
	MachineLanIP      MachineAddressType = "LanIP"
	MachineInternetIP MachineAddressType = "InternetIP"
	// TODO: DNS type
)

type MachineAddress struct {
	Type    MachineAddressType `json:"type,omitempty" protobuf:"bytes,1,opt,name=type,casttype=MachineAddressType"`
	Address string             `json:"address,omitempty" protobuf:"bytes,2,opt,name=address"`
}

type MachineRole string

const (
	//
	MachineMaster  MachineRole = "Master"
	MachineWorker  MachineRole = "Worker"
	MachineVagrant MachineRole = "Vagrant"
)

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Machine `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
