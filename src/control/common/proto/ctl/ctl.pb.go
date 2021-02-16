// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ctl/ctl.proto

package ctl

import (
	context "context"
	fmt "fmt"
	_ "github.com/daos-stack/daos/src/control/common/proto/shared"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("ctl/ctl.proto", fileDescriptor_1d50d0bc048a9dfa)
}

var fileDescriptor_1d50d0bc048a9dfa = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0x4d, 0x8c, 0x44, 0x6b, 0x40, 0xa8, 0x8a, 0x09, 0x47, 0xcf, 0xc2, 0x46, 0x31, 0x31,
	0x7a, 0x32, 0x92, 0x70, 0x34, 0x48, 0xe3, 0xc5, 0x5b, 0xe9, 0x56, 0x20, 0x6c, 0xdb, 0x75, 0x3a,
	0x48, 0xfc, 0x64, 0x7e, 0x3d, 0xd3, 0x3f, 0x1b, 0x76, 0x91, 0xc3, 0xde, 0x66, 0x7e, 0x79, 0x6f,
	0x5e, 0x99, 0x61, 0x49, 0x53, 0x60, 0x96, 0x08, 0xcc, 0x06, 0x39, 0x18, 0x34, 0xf4, 0x50, 0x60,
	0xd6, 0xeb, 0x38, 0x66, 0xd1, 0x00, 0x9f, 0xcb, 0xc0, 0x03, 0xd2, 0x12, 0x37, 0x06, 0x56, 0x11,
	0x51, 0x87, 0x3e, 0x97, 0xa0, 0x36, 0x1c, 0x0a, 0x99, 0x9f, 0x66, 0x55, 0x1a, 0xdb, 0x33, 0xd7,
	0x02, 0xd7, 0x2b, 0x5b, 0x78, 0xec, 0x82, 0x83, 0x4c, 0xcb, 0xec, 0xee, 0xf7, 0x88, 0x34, 0x46,
	0x98, 0xb1, 0x6f, 0x41, 0x47, 0xa4, 0xc5, 0x42, 0xec, 0x04, 0x64, 0xce, 0x41, 0xd2, 0xee, 0xc0,
	0xbd, 0xad, 0x0a, 0xa7, 0xf2, 0xab, 0x77, 0xb5, 0x97, 0xdb, 0xfc, 0xfa, 0x80, 0x3e, 0x91, 0xd3,
	0xc8, 0x99, 0xe0, 0x9a, 0x9e, 0x97, 0x95, 0x8e, 0x38, 0xfb, 0xc5, 0x7f, 0xe8, 0xbd, 0xcf, 0xa4,
	0x19, 0xe1, 0xd8, 0x80, 0xe2, 0x48, 0x2f, 0xcb, 0xc2, 0xc0, 0x9c, 0xbf, 0xbb, 0x0f, 0x17, 0xe9,
	0xaf, 0x61, 0x4d, 0xa5, 0xf4, 0x12, 0xd9, 0xa6, 0x57, 0x60, 0x91, 0x3e, 0x8e, 0xfb, 0x7c, 0x5b,
	0x4b, 0xf8, 0x89, 0xe9, 0x15, 0xb6, 0x4d, 0xdf, 0xc1, 0x7e, 0xc2, 0x88, 0xb4, 0x0a, 0xfc, 0x9e,
	0xa7, 0x1c, 0x8b, 0x05, 0x56, 0xe1, 0x76, 0x81, 0xbb, 0xdc, 0x0f, 0xb9, 0x25, 0xc7, 0x4c, 0xa5,
	0xe1, 0x05, 0xed, 0xf0, 0x43, 0x63, 0xeb, 0x8c, 0x9d, 0x1d, 0xe2, 0x2d, 0xf7, 0xa4, 0xe3, 0x8e,
	0xc0, 0x16, 0x6b, 0x4c, 0xcd, 0x46, 0x4f, 0xdd, 0x79, 0x69, 0xd3, 0x2b, 0x7d, 0xed, 0x8c, 0xad,
	0x72, 0xeb, 0x5d, 0x37, 0xe4, 0x84, 0xa1, 0xc9, 0xeb, 0xab, 0x27, 0x4b, 0x3d, 0xaf, 0xa9, 0x1e,
	0x92, 0xf6, 0x54, 0x5a, 0x89, 0xf1, 0x38, 0xf5, 0x4c, 0x7d, 0x42, 0x18, 0x72, 0xa8, 0x29, 0x7f,
	0x79, 0xfc, 0x78, 0x98, 0x2f, 0x71, 0xb1, 0x9e, 0x0d, 0x84, 0x51, 0x49, 0xca, 0x8d, 0xed, 0x5b,
	0xe4, 0x62, 0xe5, 0xcb, 0xc4, 0x82, 0x48, 0x84, 0xd1, 0x08, 0x26, 0x4b, 0x84, 0x51, 0xca, 0xe8,
	0xc4, 0xff, 0xe1, 0xdd, 0xd7, 0x36, 0x6b, 0xf8, 0x72, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x0a,
	0x90, 0xcd, 0x70, 0x7f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CtlSvcClient is the client API for CtlSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CtlSvcClient interface {
	// Prepare nonvolatile storage devices for use with DAOS
	StoragePrepare(ctx context.Context, in *StoragePrepareReq, opts ...grpc.CallOption) (*StoragePrepareResp, error)
	// Retrieve details of nonvolatile storage on server, including health info
	StorageScan(ctx context.Context, in *StorageScanReq, opts ...grpc.CallOption) (*StorageScanResp, error)
	// Format nonvolatile storage devices for use with DAOS
	StorageFormat(ctx context.Context, in *StorageFormatReq, opts ...grpc.CallOption) (*StorageFormatResp, error)
	// Perform a fabric scan to determine the available provider, device, NUMA node combinations
	NetworkScan(ctx context.Context, in *NetworkScanReq, opts ...grpc.CallOption) (*NetworkScanResp, error)
	// Retrieve firmware details from storage devices on server
	FirmwareQuery(ctx context.Context, in *FirmwareQueryReq, opts ...grpc.CallOption) (*FirmwareQueryResp, error)
	// Update firmware on storage devices on server
	FirmwareUpdate(ctx context.Context, in *FirmwareUpdateReq, opts ...grpc.CallOption) (*FirmwareUpdateResp, error)
	// Query the per-server metadata
	SmdQuery(ctx context.Context, in *SmdQueryReq, opts ...grpc.CallOption) (*SmdQueryResp, error)
	// Prepare DAOS I/O Engines on a host for controlled shutdown. (gRPC fanout)
	PrepShutdownRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error)
	// Stop DAOS I/O Engines on a host. (gRPC fanout)
	StopRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error)
	// Ping DAOS I/O Engines on a host. (gRPC fanout)
	PingRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error)
	// ResetFormat DAOS I/O Engines on a host. (gRPC fanout)
	ResetFormatRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error)
	// Start DAOS I/O Engines on a host. (gRPC fanout)
	StartRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error)
}

type ctlSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewCtlSvcClient(cc grpc.ClientConnInterface) CtlSvcClient {
	return &ctlSvcClient{cc}
}

func (c *ctlSvcClient) StoragePrepare(ctx context.Context, in *StoragePrepareReq, opts ...grpc.CallOption) (*StoragePrepareResp, error) {
	out := new(StoragePrepareResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/StoragePrepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) StorageScan(ctx context.Context, in *StorageScanReq, opts ...grpc.CallOption) (*StorageScanResp, error) {
	out := new(StorageScanResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/StorageScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) StorageFormat(ctx context.Context, in *StorageFormatReq, opts ...grpc.CallOption) (*StorageFormatResp, error) {
	out := new(StorageFormatResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/StorageFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) NetworkScan(ctx context.Context, in *NetworkScanReq, opts ...grpc.CallOption) (*NetworkScanResp, error) {
	out := new(NetworkScanResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/NetworkScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) FirmwareQuery(ctx context.Context, in *FirmwareQueryReq, opts ...grpc.CallOption) (*FirmwareQueryResp, error) {
	out := new(FirmwareQueryResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/FirmwareQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) FirmwareUpdate(ctx context.Context, in *FirmwareUpdateReq, opts ...grpc.CallOption) (*FirmwareUpdateResp, error) {
	out := new(FirmwareUpdateResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/FirmwareUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) SmdQuery(ctx context.Context, in *SmdQueryReq, opts ...grpc.CallOption) (*SmdQueryResp, error) {
	out := new(SmdQueryResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/SmdQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) PrepShutdownRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error) {
	out := new(RanksResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/PrepShutdownRanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) StopRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error) {
	out := new(RanksResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/StopRanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) PingRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error) {
	out := new(RanksResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/PingRanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) ResetFormatRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error) {
	out := new(RanksResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/ResetFormatRanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctlSvcClient) StartRanks(ctx context.Context, in *RanksReq, opts ...grpc.CallOption) (*RanksResp, error) {
	out := new(RanksResp)
	err := c.cc.Invoke(ctx, "/ctl.CtlSvc/StartRanks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CtlSvcServer is the server API for CtlSvc service.
type CtlSvcServer interface {
	// Prepare nonvolatile storage devices for use with DAOS
	StoragePrepare(context.Context, *StoragePrepareReq) (*StoragePrepareResp, error)
	// Retrieve details of nonvolatile storage on server, including health info
	StorageScan(context.Context, *StorageScanReq) (*StorageScanResp, error)
	// Format nonvolatile storage devices for use with DAOS
	StorageFormat(context.Context, *StorageFormatReq) (*StorageFormatResp, error)
	// Perform a fabric scan to determine the available provider, device, NUMA node combinations
	NetworkScan(context.Context, *NetworkScanReq) (*NetworkScanResp, error)
	// Retrieve firmware details from storage devices on server
	FirmwareQuery(context.Context, *FirmwareQueryReq) (*FirmwareQueryResp, error)
	// Update firmware on storage devices on server
	FirmwareUpdate(context.Context, *FirmwareUpdateReq) (*FirmwareUpdateResp, error)
	// Query the per-server metadata
	SmdQuery(context.Context, *SmdQueryReq) (*SmdQueryResp, error)
	// Prepare DAOS I/O Engines on a host for controlled shutdown. (gRPC fanout)
	PrepShutdownRanks(context.Context, *RanksReq) (*RanksResp, error)
	// Stop DAOS I/O Engines on a host. (gRPC fanout)
	StopRanks(context.Context, *RanksReq) (*RanksResp, error)
	// Ping DAOS I/O Engines on a host. (gRPC fanout)
	PingRanks(context.Context, *RanksReq) (*RanksResp, error)
	// ResetFormat DAOS I/O Engines on a host. (gRPC fanout)
	ResetFormatRanks(context.Context, *RanksReq) (*RanksResp, error)
	// Start DAOS I/O Engines on a host. (gRPC fanout)
	StartRanks(context.Context, *RanksReq) (*RanksResp, error)
}

// UnimplementedCtlSvcServer can be embedded to have forward compatible implementations.
type UnimplementedCtlSvcServer struct {
}

func (*UnimplementedCtlSvcServer) StoragePrepare(ctx context.Context, req *StoragePrepareReq) (*StoragePrepareResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoragePrepare not implemented")
}
func (*UnimplementedCtlSvcServer) StorageScan(ctx context.Context, req *StorageScanReq) (*StorageScanResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageScan not implemented")
}
func (*UnimplementedCtlSvcServer) StorageFormat(ctx context.Context, req *StorageFormatReq) (*StorageFormatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorageFormat not implemented")
}
func (*UnimplementedCtlSvcServer) NetworkScan(ctx context.Context, req *NetworkScanReq) (*NetworkScanResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetworkScan not implemented")
}
func (*UnimplementedCtlSvcServer) FirmwareQuery(ctx context.Context, req *FirmwareQueryReq) (*FirmwareQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FirmwareQuery not implemented")
}
func (*UnimplementedCtlSvcServer) FirmwareUpdate(ctx context.Context, req *FirmwareUpdateReq) (*FirmwareUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FirmwareUpdate not implemented")
}
func (*UnimplementedCtlSvcServer) SmdQuery(ctx context.Context, req *SmdQueryReq) (*SmdQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmdQuery not implemented")
}
func (*UnimplementedCtlSvcServer) PrepShutdownRanks(ctx context.Context, req *RanksReq) (*RanksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepShutdownRanks not implemented")
}
func (*UnimplementedCtlSvcServer) StopRanks(ctx context.Context, req *RanksReq) (*RanksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRanks not implemented")
}
func (*UnimplementedCtlSvcServer) PingRanks(ctx context.Context, req *RanksReq) (*RanksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingRanks not implemented")
}
func (*UnimplementedCtlSvcServer) ResetFormatRanks(ctx context.Context, req *RanksReq) (*RanksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetFormatRanks not implemented")
}
func (*UnimplementedCtlSvcServer) StartRanks(ctx context.Context, req *RanksReq) (*RanksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRanks not implemented")
}

func RegisterCtlSvcServer(s *grpc.Server, srv CtlSvcServer) {
	s.RegisterService(&_CtlSvc_serviceDesc, srv)
}

func _CtlSvc_StoragePrepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoragePrepareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).StoragePrepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/StoragePrepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).StoragePrepare(ctx, req.(*StoragePrepareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_StorageScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageScanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).StorageScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/StorageScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).StorageScan(ctx, req.(*StorageScanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_StorageFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageFormatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).StorageFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/StorageFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).StorageFormat(ctx, req.(*StorageFormatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_NetworkScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkScanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).NetworkScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/NetworkScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).NetworkScan(ctx, req.(*NetworkScanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_FirmwareQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirmwareQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).FirmwareQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/FirmwareQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).FirmwareQuery(ctx, req.(*FirmwareQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_FirmwareUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirmwareUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).FirmwareUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/FirmwareUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).FirmwareUpdate(ctx, req.(*FirmwareUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_SmdQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmdQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).SmdQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/SmdQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).SmdQuery(ctx, req.(*SmdQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_PrepShutdownRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RanksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).PrepShutdownRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/PrepShutdownRanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).PrepShutdownRanks(ctx, req.(*RanksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_StopRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RanksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).StopRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/StopRanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).StopRanks(ctx, req.(*RanksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_PingRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RanksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).PingRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/PingRanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).PingRanks(ctx, req.(*RanksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_ResetFormatRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RanksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).ResetFormatRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/ResetFormatRanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).ResetFormatRanks(ctx, req.(*RanksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CtlSvc_StartRanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RanksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtlSvcServer).StartRanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ctl.CtlSvc/StartRanks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtlSvcServer).StartRanks(ctx, req.(*RanksReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CtlSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ctl.CtlSvc",
	HandlerType: (*CtlSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoragePrepare",
			Handler:    _CtlSvc_StoragePrepare_Handler,
		},
		{
			MethodName: "StorageScan",
			Handler:    _CtlSvc_StorageScan_Handler,
		},
		{
			MethodName: "StorageFormat",
			Handler:    _CtlSvc_StorageFormat_Handler,
		},
		{
			MethodName: "NetworkScan",
			Handler:    _CtlSvc_NetworkScan_Handler,
		},
		{
			MethodName: "FirmwareQuery",
			Handler:    _CtlSvc_FirmwareQuery_Handler,
		},
		{
			MethodName: "FirmwareUpdate",
			Handler:    _CtlSvc_FirmwareUpdate_Handler,
		},
		{
			MethodName: "SmdQuery",
			Handler:    _CtlSvc_SmdQuery_Handler,
		},
		{
			MethodName: "PrepShutdownRanks",
			Handler:    _CtlSvc_PrepShutdownRanks_Handler,
		},
		{
			MethodName: "StopRanks",
			Handler:    _CtlSvc_StopRanks_Handler,
		},
		{
			MethodName: "PingRanks",
			Handler:    _CtlSvc_PingRanks_Handler,
		},
		{
			MethodName: "ResetFormatRanks",
			Handler:    _CtlSvc_ResetFormatRanks_Handler,
		},
		{
			MethodName: "StartRanks",
			Handler:    _CtlSvc_StartRanks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ctl/ctl.proto",
}
