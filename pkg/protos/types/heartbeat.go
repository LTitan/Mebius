package types

import "github.com/LTitan/Mebius/apis/v1alpha1"

type HeartbeatRequest struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	UUID string `json:"uuid,omitempty" protobuf:"bytes,2,opt,name=uuid"`
}

type HeartbeatResponse struct {
	Code    int32               `json:"code,omitempty" protobuf:"varint,1,opt,name=code"`
	Message string              `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
	Data    []*v1alpha1.Machine `json:"data,omitempty" protobuf:"bytes,3,rep,name=data"`
}
