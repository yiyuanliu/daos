// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package mgmt

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// RankStateEventInfo defines extended fields for rank state change related events.
type RankStateEventInfo struct {
	Instance             uint32   `protobuf:"varint,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Errored              bool     `protobuf:"varint,2,opt,name=errored,proto3" json:"errored,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankStateEventInfo) Reset()         { *m = RankStateEventInfo{} }
func (m *RankStateEventInfo) String() string { return proto.CompactTextString(m) }
func (*RankStateEventInfo) ProtoMessage()    {}
func (*RankStateEventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *RankStateEventInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankStateEventInfo.Unmarshal(m, b)
}
func (m *RankStateEventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankStateEventInfo.Marshal(b, m, deterministic)
}
func (m *RankStateEventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankStateEventInfo.Merge(m, src)
}
func (m *RankStateEventInfo) XXX_Size() int {
	return xxx_messageInfo_RankStateEventInfo.Size(m)
}
func (m *RankStateEventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RankStateEventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RankStateEventInfo proto.InternalMessageInfo

func (m *RankStateEventInfo) GetInstance() uint32 {
	if m != nil {
		return m.Instance
	}
	return 0
}

func (m *RankStateEventInfo) GetErrored() bool {
	if m != nil {
		return m.Errored
	}
	return false
}

func (m *RankStateEventInfo) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// PoolSvcEventInfo defines extended fields for pool service change events.
type PoolSvcEventInfo struct {
	PoolUuid             string   `protobuf:"bytes,1,opt,name=pool_uuid,json=poolUuid,proto3" json:"pool_uuid,omitempty"`
	SvcReps              []uint32 `protobuf:"varint,2,rep,packed,name=svc_reps,json=svcReps,proto3" json:"svc_reps,omitempty"`
	Version              uint64   `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolSvcEventInfo) Reset()         { *m = PoolSvcEventInfo{} }
func (m *PoolSvcEventInfo) String() string { return proto.CompactTextString(m) }
func (*PoolSvcEventInfo) ProtoMessage()    {}
func (*PoolSvcEventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{1}
}

func (m *PoolSvcEventInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolSvcEventInfo.Unmarshal(m, b)
}
func (m *PoolSvcEventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolSvcEventInfo.Marshal(b, m, deterministic)
}
func (m *PoolSvcEventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolSvcEventInfo.Merge(m, src)
}
func (m *PoolSvcEventInfo) XXX_Size() int {
	return xxx_messageInfo_PoolSvcEventInfo.Size(m)
}
func (m *PoolSvcEventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolSvcEventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PoolSvcEventInfo proto.InternalMessageInfo

func (m *PoolSvcEventInfo) GetPoolUuid() string {
	if m != nil {
		return m.PoolUuid
	}
	return ""
}

func (m *PoolSvcEventInfo) GetSvcReps() []uint32 {
	if m != nil {
		return m.SvcReps
	}
	return nil
}

func (m *PoolSvcEventInfo) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// RASEvent describes a RAS event in the DAOS system.
type RASEvent struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Severity  uint32 `protobuf:"varint,3,opt,name=severity,proto3" json:"severity,omitempty"`
	Msg       string `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	Type      uint32 `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Rank      uint32 `protobuf:"varint,6,opt,name=rank,proto3" json:"rank,omitempty"`
	Hostname  string `protobuf:"bytes,7,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Types that are valid to be assigned to ExtendedInfo:
	//	*RASEvent_RankStateInfo
	//	*RASEvent_PoolSvcInfo
	ExtendedInfo         isRASEvent_ExtendedInfo `protobuf_oneof:"extended_info"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RASEvent) Reset()         { *m = RASEvent{} }
func (m *RASEvent) String() string { return proto.CompactTextString(m) }
func (*RASEvent) ProtoMessage()    {}
func (*RASEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{2}
}

func (m *RASEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RASEvent.Unmarshal(m, b)
}
func (m *RASEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RASEvent.Marshal(b, m, deterministic)
}
func (m *RASEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RASEvent.Merge(m, src)
}
func (m *RASEvent) XXX_Size() int {
	return xxx_messageInfo_RASEvent.Size(m)
}
func (m *RASEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RASEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RASEvent proto.InternalMessageInfo

func (m *RASEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RASEvent) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *RASEvent) GetSeverity() uint32 {
	if m != nil {
		return m.Severity
	}
	return 0
}

func (m *RASEvent) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RASEvent) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RASEvent) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *RASEvent) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type isRASEvent_ExtendedInfo interface {
	isRASEvent_ExtendedInfo()
}

type RASEvent_RankStateInfo struct {
	RankStateInfo *RankStateEventInfo `protobuf:"bytes,8,opt,name=rank_state_info,json=rankStateInfo,proto3,oneof"`
}

type RASEvent_PoolSvcInfo struct {
	PoolSvcInfo *PoolSvcEventInfo `protobuf:"bytes,9,opt,name=pool_svc_info,json=poolSvcInfo,proto3,oneof"`
}

func (*RASEvent_RankStateInfo) isRASEvent_ExtendedInfo() {}

func (*RASEvent_PoolSvcInfo) isRASEvent_ExtendedInfo() {}

func (m *RASEvent) GetExtendedInfo() isRASEvent_ExtendedInfo {
	if m != nil {
		return m.ExtendedInfo
	}
	return nil
}

func (m *RASEvent) GetRankStateInfo() *RankStateEventInfo {
	if x, ok := m.GetExtendedInfo().(*RASEvent_RankStateInfo); ok {
		return x.RankStateInfo
	}
	return nil
}

func (m *RASEvent) GetPoolSvcInfo() *PoolSvcEventInfo {
	if x, ok := m.GetExtendedInfo().(*RASEvent_PoolSvcInfo); ok {
		return x.PoolSvcInfo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RASEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RASEvent_RankStateInfo)(nil),
		(*RASEvent_PoolSvcInfo)(nil),
	}
}

// ClusterEventReq contains event details, request sequence and extended info.
type ClusterEventReq struct {
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*ClusterEventReq_Ras
	Event                isClusterEventReq_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClusterEventReq) Reset()         { *m = ClusterEventReq{} }
func (m *ClusterEventReq) String() string { return proto.CompactTextString(m) }
func (*ClusterEventReq) ProtoMessage()    {}
func (*ClusterEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{3}
}

func (m *ClusterEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterEventReq.Unmarshal(m, b)
}
func (m *ClusterEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterEventReq.Marshal(b, m, deterministic)
}
func (m *ClusterEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterEventReq.Merge(m, src)
}
func (m *ClusterEventReq) XXX_Size() int {
	return xxx_messageInfo_ClusterEventReq.Size(m)
}
func (m *ClusterEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterEventReq proto.InternalMessageInfo

func (m *ClusterEventReq) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

type isClusterEventReq_Event interface {
	isClusterEventReq_Event()
}

type ClusterEventReq_Ras struct {
	Ras *RASEvent `protobuf:"bytes,2,opt,name=ras,proto3,oneof"`
}

func (*ClusterEventReq_Ras) isClusterEventReq_Event() {}

func (m *ClusterEventReq) GetEvent() isClusterEventReq_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ClusterEventReq) GetRas() *RASEvent {
	if x, ok := m.GetEvent().(*ClusterEventReq_Ras); ok {
		return x.Ras
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ClusterEventReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ClusterEventReq_Ras)(nil),
	}
}

// ClusterEventResp acknowledges receipt of an event notification and an
// error status.
type ClusterEventResp struct {
	Sequence             uint64   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterEventResp) Reset()         { *m = ClusterEventResp{} }
func (m *ClusterEventResp) String() string { return proto.CompactTextString(m) }
func (*ClusterEventResp) ProtoMessage()    {}
func (*ClusterEventResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{4}
}

func (m *ClusterEventResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterEventResp.Unmarshal(m, b)
}
func (m *ClusterEventResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterEventResp.Marshal(b, m, deterministic)
}
func (m *ClusterEventResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterEventResp.Merge(m, src)
}
func (m *ClusterEventResp) XXX_Size() int {
	return xxx_messageInfo_ClusterEventResp.Size(m)
}
func (m *ClusterEventResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterEventResp.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterEventResp proto.InternalMessageInfo

func (m *ClusterEventResp) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ClusterEventResp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*RankStateEventInfo)(nil), "mgmt.RankStateEventInfo")
	proto.RegisterType((*PoolSvcEventInfo)(nil), "mgmt.PoolSvcEventInfo")
	proto.RegisterType((*RASEvent)(nil), "mgmt.RASEvent")
	proto.RegisterType((*ClusterEventReq)(nil), "mgmt.ClusterEventReq")
	proto.RegisterType((*ClusterEventResp)(nil), "mgmt.ClusterEventResp")
}

func init() {
	proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e)
}

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x9b, 0x34, 0x6d, 0x93, 0xa9, 0xb2, 0xad, 0x2c, 0xb4, 0x32, 0x7f, 0x0e, 0x51, 0x4e,
	0x39, 0xf5, 0xb0, 0x5c, 0xb9, 0x50, 0x04, 0x2a, 0x37, 0x34, 0x15, 0x17, 0x2e, 0x25, 0x34, 0xb3,
	0x4b, 0xb4, 0x8d, 0x9d, 0xb5, 0x9d, 0x88, 0xfd, 0x70, 0x7c, 0x37, 0xe4, 0x49, 0xd3, 0x15, 0x20,
	0xed, 0x6d, 0xde, 0xb3, 0xfd, 0x1b, 0xfb, 0x79, 0x60, 0x49, 0x3d, 0x29, 0xb7, 0x69, 0x8d, 0x76,
	0x5a, 0x44, 0xcd, 0x5d, 0xe3, 0xf2, 0xef, 0x20, 0xb0, 0x54, 0xf7, 0x7b, 0x57, 0x3a, 0xfa, 0xe8,
	0x57, 0x3f, 0xab, 0x5b, 0x2d, 0x5e, 0x41, 0x5c, 0x2b, 0xeb, 0x4a, 0x75, 0x24, 0x19, 0x64, 0x41,
	0x91, 0xe2, 0x45, 0x0b, 0x09, 0x0b, 0x32, 0x46, 0x1b, 0xaa, 0x64, 0x98, 0x05, 0x45, 0x8c, 0xa3,
	0x14, 0x2f, 0x60, 0xc6, 0xa5, 0x9c, 0x66, 0x41, 0x91, 0xe0, 0x20, 0xf2, 0x0a, 0xd6, 0x5f, 0xb4,
	0x3e, 0xed, 0xfb, 0xe3, 0x13, 0xff, 0x35, 0x24, 0xad, 0xd6, 0xa7, 0x43, 0xd7, 0xd5, 0x15, 0x37,
	0x48, 0x30, 0xf6, 0xc6, 0xd7, 0xae, 0xae, 0xc4, 0x4b, 0x88, 0x6d, 0x7f, 0x3c, 0x18, 0x6a, 0xad,
	0x0c, 0xb3, 0x69, 0x91, 0xe2, 0xc2, 0xf6, 0x47, 0xa4, 0xd6, 0xfa, 0xde, 0x3d, 0x19, 0x5b, 0x6b,
	0xc5, 0x3d, 0x22, 0x1c, 0x65, 0xfe, 0x3b, 0x84, 0x18, 0xdf, 0xef, 0xb9, 0x85, 0xb8, 0x82, 0xf0,
	0xc2, 0x0d, 0xeb, 0x4a, 0xbc, 0x81, 0xc4, 0xd5, 0x0d, 0x59, 0x57, 0x36, 0x2d, 0x5f, 0x3a, 0xc1,
	0x27, 0xc3, 0x3f, 0xd6, 0x52, 0x4f, 0xa6, 0x76, 0x8f, 0x4c, 0x4d, 0xf1, 0xa2, 0xc5, 0x1a, 0xa6,
	0x8d, 0xbd, 0x93, 0x11, 0x9f, 0xf1, 0xa5, 0x10, 0x10, 0xb9, 0xc7, 0x96, 0xe4, 0x8c, 0x77, 0x72,
	0xed, 0x3d, 0x53, 0xaa, 0x7b, 0x39, 0x1f, 0x3c, 0x5f, 0x7b, 0xea, 0x4f, 0x6d, 0x9d, 0x2a, 0x1b,
	0x92, 0x8b, 0xe1, 0x85, 0xa3, 0x16, 0x5b, 0x58, 0xf9, 0x3d, 0x07, 0xeb, 0x53, 0x3f, 0xd4, 0xea,
	0x56, 0xcb, 0x38, 0x0b, 0x8a, 0xe5, 0x8d, 0xdc, 0xf8, 0x4f, 0xd9, 0xfc, 0xff, 0x23, 0xbb, 0x09,
	0xa6, 0x66, 0x74, 0x39, 0xc2, 0x77, 0x90, 0x72, 0x84, 0x3e, 0x2a, 0x26, 0x24, 0x4c, 0xb8, 0x1e,
	0x08, 0xff, 0x26, 0xbe, 0x9b, 0xe0, 0xb2, 0x1d, 0x3c, 0x2f, 0xb7, 0x2b, 0x48, 0xe9, 0x97, 0x23,
	0x55, 0x51, 0xc5, 0xa7, 0xf3, 0x6f, 0xb0, 0xfa, 0x70, 0xea, 0xac, 0x23, 0xc3, 0x67, 0x90, 0x1e,
	0x86, 0x5c, 0x1e, 0x3a, 0x1a, 0x87, 0x20, 0xc2, 0x8b, 0x16, 0x39, 0x4c, 0x4d, 0x69, 0x39, 0xcb,
	0xe5, 0xcd, 0xd5, 0xf9, 0xd6, 0xe7, 0xf8, 0x77, 0x13, 0xf4, 0x8b, 0xdb, 0x05, 0xcc, 0x78, 0xde,
	0xf2, 0x4f, 0xb0, 0xfe, 0x9b, 0x6d, 0xdb, 0x67, 0xe1, 0xd7, 0x30, 0xf7, 0xc9, 0x74, 0x03, 0x7f,
	0x86, 0x67, 0xf5, 0x63, 0xce, 0x83, 0xfb, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x4f,
	0x30, 0xe9, 0xc7, 0x02, 0x00, 0x00,
}