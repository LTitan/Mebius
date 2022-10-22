/*
this is the Mebius Project
*/ // Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/LTitan/Mebius/pkg/apis/generated.proto

package apis

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *Cluster) Reset()      { *m = Cluster{} }
func (*Cluster) ProtoMessage() {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b387b061c0a69a, []int{0}
}
func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return m.Size()
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *ClusterCondition) Reset()      { *m = ClusterCondition{} }
func (*ClusterCondition) ProtoMessage() {}
func (*ClusterCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b387b061c0a69a, []int{1}
}
func (m *ClusterCondition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterCondition.Merge(m, src)
}
func (m *ClusterCondition) XXX_Size() int {
	return m.Size()
}
func (m *ClusterCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterCondition.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterCondition proto.InternalMessageInfo

func (m *ClusterList) Reset()      { *m = ClusterList{} }
func (*ClusterList) ProtoMessage() {}
func (*ClusterList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b387b061c0a69a, []int{2}
}
func (m *ClusterList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterList.Merge(m, src)
}
func (m *ClusterList) XXX_Size() int {
	return m.Size()
}
func (m *ClusterList) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterList.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterList proto.InternalMessageInfo

func (m *ClusterSpec) Reset()      { *m = ClusterSpec{} }
func (*ClusterSpec) ProtoMessage() {}
func (*ClusterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b387b061c0a69a, []int{3}
}
func (m *ClusterSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterSpec.Merge(m, src)
}
func (m *ClusterSpec) XXX_Size() int {
	return m.Size()
}
func (m *ClusterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterSpec proto.InternalMessageInfo

func (m *ClusterStatus) Reset()      { *m = ClusterStatus{} }
func (*ClusterStatus) ProtoMessage() {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b387b061c0a69a, []int{4}
}
func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return m.Size()
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Cluster)(nil), "github.com.LTitan.Mebius.pkg.apis.Cluster")
	proto.RegisterType((*ClusterCondition)(nil), "github.com.LTitan.Mebius.pkg.apis.ClusterCondition")
	proto.RegisterType((*ClusterList)(nil), "github.com.LTitan.Mebius.pkg.apis.ClusterList")
	proto.RegisterType((*ClusterSpec)(nil), "github.com.LTitan.Mebius.pkg.apis.ClusterSpec")
	proto.RegisterMapType((map[string]string)(nil), "github.com.LTitan.Mebius.pkg.apis.ClusterSpec.AttributesEntry")
	proto.RegisterType((*ClusterStatus)(nil), "github.com.LTitan.Mebius.pkg.apis.ClusterStatus")
}

func init() {
	proto.RegisterFile("github.com/LTitan/Mebius/pkg/apis/generated.proto", fileDescriptor_f4b387b061c0a69a)
}

var fileDescriptor_f4b387b061c0a69a = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0xc7, 0xb3, 0x79, 0xa9, 0x66, 0x62, 0x6a, 0x1d, 0x44, 0x42, 0x0e, 0x9b, 0x36, 0xa7, 0x22,
	0x38, 0x6b, 0x54, 0xa4, 0x08, 0x8a, 0x6e, 0x10, 0x11, 0x5a, 0x2a, 0x6b, 0x84, 0xa2, 0x07, 0x9d,
	0x6c, 0xc6, 0xcd, 0x98, 0xee, 0x0b, 0xbb, 0xcf, 0x06, 0x72, 0xf3, 0x23, 0x78, 0xf3, 0x53, 0xf8,
	0x1d, 0x3c, 0xe6, 0xd8, 0x63, 0x4f, 0xc1, 0xac, 0xdf, 0xc2, 0x93, 0xcc, 0xec, 0x74, 0x37, 0x89,
	0x48, 0xd2, 0x4b, 0x6f, 0xf3, 0xec, 0xcc, 0xff, 0xf7, 0xbc, 0xfd, 0x59, 0xd4, 0x71, 0x38, 0x0c,
	0xe3, 0x3e, 0xb1, 0x7d, 0xd7, 0x38, 0xec, 0x71, 0xa0, 0x9e, 0x71, 0xc4, 0xfa, 0x3c, 0x8e, 0x8c,
	0x60, 0xe4, 0x18, 0x34, 0xe0, 0x91, 0xe1, 0x30, 0x8f, 0x85, 0x14, 0xd8, 0x80, 0x04, 0xa1, 0x0f,
	0x3e, 0xde, 0xcb, 0x25, 0x24, 0x95, 0x90, 0x54, 0x42, 0x82, 0x91, 0x43, 0x84, 0xa4, 0x79, 0x6f,
	0x81, 0xea, 0xf8, 0x8e, 0x6f, 0x48, 0x65, 0x3f, 0xfe, 0x2c, 0x23, 0x19, 0xc8, 0x53, 0x4a, 0x6c,
	0x3e, 0x1a, 0x1d, 0x44, 0x84, 0xfb, 0x22, 0x9d, 0x4b, 0xed, 0x21, 0xf7, 0x58, 0x38, 0xc9, 0xf3,
	0xbb, 0x0c, 0xa8, 0x31, 0xee, 0xac, 0xd6, 0xd1, 0x34, 0xfe, 0xa7, 0x0a, 0x63, 0x0f, 0xb8, 0xcb,
	0xfe, 0x11, 0x3c, 0x5e, 0x27, 0x88, 0xec, 0x21, 0x73, 0xe9, 0xaa, 0xae, 0xfd, 0xbd, 0x88, 0xae,
	0x75, 0x4f, 0xe3, 0x08, 0x58, 0x88, 0x3f, 0xa1, 0xeb, 0xa2, 0x9e, 0x01, 0x05, 0xda, 0xd0, 0x76,
	0xb5, 0xfd, 0xda, 0x83, 0xfb, 0x24, 0xc5, 0x92, 0x45, 0x6c, 0x36, 0x0a, 0x22, 0x5e, 0x93, 0x71,
	0x87, 0x1c, 0xf7, 0xbf, 0x30, 0x1b, 0x8e, 0x18, 0x50, 0x13, 0x4f, 0x67, 0xad, 0x42, 0x32, 0x6b,
	0xa1, 0xfc, 0x9b, 0x95, 0x51, 0xf1, 0x1b, 0x54, 0x8e, 0x02, 0x66, 0x37, 0x8a, 0x92, 0x4e, 0xc8,
	0xda, 0x69, 0x13, 0x55, 0xdb, 0xdb, 0x80, 0xd9, 0xe6, 0x0d, 0xc5, 0x2e, 0x8b, 0xc8, 0x92, 0x24,
	0x7c, 0x82, 0xb6, 0x22, 0xa0, 0x10, 0x47, 0x8d, 0x92, 0xaa, 0x78, 0x73, 0xa6, 0xd4, 0x99, 0xdb,
	0x8a, 0xba, 0x95, 0xc6, 0x96, 0xe2, 0xb5, 0x7f, 0x14, 0xd1, 0x8e, 0x7a, 0xd9, 0xf5, 0xbd, 0x01,
	0x07, 0xee, 0x7b, 0x78, 0x17, 0x95, 0x61, 0x12, 0x30, 0x39, 0x9e, 0x6a, 0x5e, 0x50, 0x6f, 0x12,
	0x30, 0x4b, 0xde, 0xe0, 0xe7, 0x59, 0x41, 0x45, 0xf9, 0x66, 0x7f, 0x19, 0xff, 0x67, 0xd6, 0xba,
	0xb3, 0x4a, 0x5d, 0x4e, 0x8c, 0x4f, 0x10, 0xb2, 0x43, 0x46, 0xe1, 0xa3, 0xd8, 0x9b, 0x6a, 0xeb,
	0xee, 0x66, 0x8b, 0xe8, 0x71, 0x97, 0x99, 0xf5, 0x64, 0xd6, 0xaa, 0x76, 0x05, 0x41, 0x84, 0x56,
	0xd5, 0xbe, 0x38, 0xe2, 0x0f, 0xa8, 0x16, 0x07, 0x03, 0x0a, 0x2c, 0x45, 0x97, 0x2f, 0x8d, 0xde,
	0x16, 0x9b, 0x7d, 0x27, 0x11, 0x92, 0x8d, 0xe2, 0xec, 0xdc, 0xfe, 0xa9, 0xa1, 0x9a, 0xea, 0xec,
	0x90, 0x47, 0x70, 0x05, 0x6e, 0x3a, 0x46, 0x15, 0x0e, 0xcc, 0x15, 0x93, 0x2e, 0xc9, 0x46, 0x36,
	0x5e, 0xbd, 0x59, 0x57, 0xe0, 0xca, 0x6b, 0x01, 0xb0, 0x52, 0xce, 0x62, 0x0b, 0xc2, 0x62, 0x38,
	0x44, 0x88, 0x02, 0x84, 0xbc, 0x1f, 0x03, 0x8b, 0x1a, 0x9a, 0xcc, 0xf2, 0xec, 0x72, 0xa6, 0x25,
	0x2f, 0x32, 0xc0, 0x4b, 0x0f, 0xc2, 0x49, 0xde, 0x52, 0x7e, 0x61, 0x2d, 0x64, 0x69, 0x3e, 0x45,
	0x37, 0x57, 0x24, 0x78, 0x07, 0x95, 0x46, 0x6c, 0x92, 0x7a, 0xce, 0x12, 0x47, 0x7c, 0x1b, 0x55,
	0xc6, 0xf4, 0x34, 0x66, 0xa9, 0xc7, 0xac, 0x34, 0x78, 0x52, 0x3c, 0xd0, 0xda, 0x31, 0xaa, 0x2f,
	0xd9, 0x1b, 0x0f, 0x50, 0xd5, 0xbe, 0x30, 0x9a, 0x6a, 0xe1, 0xe1, 0xe6, 0x2d, 0x64, 0x1e, 0x35,
	0x6f, 0xa9, 0xba, 0xab, 0xd9, 0x27, 0x2b, 0x07, 0x9b, 0xaf, 0xa6, 0x73, 0xbd, 0x70, 0x36, 0xd7,
	0x0b, 0xe7, 0x73, 0xbd, 0xf0, 0x35, 0xd1, 0xb5, 0x69, 0xa2, 0x6b, 0x67, 0x89, 0xae, 0x9d, 0x27,
	0xba, 0xf6, 0x2b, 0xd1, 0xb5, 0x6f, 0xbf, 0xf5, 0xc2, 0xfb, 0xbd, 0xb5, 0x3f, 0xe4, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0x6f, 0x63, 0xfb, 0xb4, 0x05, 0x00, 0x00,
}

func (m *Cluster) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cluster) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cluster) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClusterCondition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterCondition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterCondition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdateTime != nil {
		{
			size, err := m.UpdateTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.CreatTime != nil {
		{
			size, err := m.CreatTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	i -= len(m.Status)
	copy(dAtA[i:], m.Status)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Status)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Type)
	copy(dAtA[i:], m.Type)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Type)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClusterList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClusterSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		keysForAttributes := make([]string, 0, len(m.Attributes))
		for k := range m.Attributes {
			keysForAttributes = append(keysForAttributes, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForAttributes)
		for iNdEx := len(keysForAttributes) - 1; iNdEx >= 0; iNdEx-- {
			v := m.Attributes[string(keysForAttributes[iNdEx])]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenerated(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(keysForAttributes[iNdEx])
			copy(dAtA[i:], keysForAttributes[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(keysForAttributes[iNdEx])))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenerated(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClusterStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Condition) > 0 {
		for iNdEx := len(m.Condition) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Condition[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ClusterCondition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Status)
	n += 1 + l + sovGenerated(uint64(l))
	if m.CreatTime != nil {
		l = m.CreatTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.UpdateTime != nil {
		l = m.UpdateTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *ClusterList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ClusterSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for k, v := range m.Attributes {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ClusterStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Condition) > 0 {
		for _, e := range m.Condition {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Cluster) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Cluster{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ClusterSpec", "ClusterSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ClusterStatus", "ClusterStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterCondition) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterCondition{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`CreatTime:` + strings.Replace(fmt.Sprintf("%v", this.CreatTime), "Time", "v1.Time", 1) + `,`,
		`UpdateTime:` + strings.Replace(fmt.Sprintf("%v", this.UpdateTime), "Time", "v1.Time", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]Cluster{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "Cluster", "Cluster", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&ClusterList{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterSpec) String() string {
	if this == nil {
		return "nil"
	}
	keysForAttributes := make([]string, 0, len(this.Attributes))
	for k := range this.Attributes {
		keysForAttributes = append(keysForAttributes, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAttributes)
	mapStringForAttributes := "map[string]string{"
	for _, k := range keysForAttributes {
		mapStringForAttributes += fmt.Sprintf("%v: %v,", k, this.Attributes[k])
	}
	mapStringForAttributes += "}"
	s := strings.Join([]string{`&ClusterSpec{`,
		`Attributes:` + mapStringForAttributes + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterStatus) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForCondition := "[]ClusterCondition{"
	for _, f := range this.Condition {
		repeatedStringForCondition += strings.Replace(strings.Replace(f.String(), "ClusterCondition", "ClusterCondition", 1), `&`, ``, 1) + ","
	}
	repeatedStringForCondition += "}"
	s := strings.Join([]string{`&ClusterStatus{`,
		`Condition:` + repeatedStringForCondition + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Cluster) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cluster: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cluster: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterCondition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = ClusterConditionStatus(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreatTime == nil {
				m.CreatTime = &v1.Time{}
			}
			if err := m.CreatTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdateTime == nil {
				m.UpdateTime = &v1.Time{}
			}
			if err := m.UpdateTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, Cluster{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attributes == nil {
				m.Attributes = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenerated
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenerated
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenerated(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenerated
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Attributes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClusterStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Condition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Condition = append(m.Condition, ClusterCondition{})
			if err := m.Condition[len(m.Condition)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
