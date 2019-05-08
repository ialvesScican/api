// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package fileshareproto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// CreateFileShareOpts is a structure which indicates all required properties for creating a file share.
type CreateFileShareOpts struct {
	// The uuid of the file share, optional when creating.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the file share, required.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The requested capacity of the file share, required.
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// The description of the file share, optional.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// The locality that file share belongs to, required.
	AvailabilityZone string `protobuf:"bytes,6,opt,name=availabilityZone,proto3" json:"availabilityZone,omitempty"`
	// The service level that file share belongs to, required.
	ProfileId string `protobuf:"bytes,7,opt,name=profileId,proto3" json:"profileId,omitempty"`
	// The uuid of the pool on which file share will be created, required.
	PoolId string `protobuf:"bytes,8,opt,name=poolId,proto3" json:"poolId,omitempty"`
	// The name of the pool on which file share will be created, required.
	PoolName string `protobuf:"bytes,9,opt,name=poolName,proto3" json:"poolName,omitempty"`
	// The metadata of the file share, optional.
	Metadata map[string]string `protobuf:"bytes,10,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The storage driver type.
	DriverName string `protobuf:"bytes,11,opt,name=driverName,proto3" json:"driverName,omitempty"`
	// The Context
	Context string `protobuf:"bytes,12,opt,name=context,proto3" json:"context,omitempty"`
	// The Serialized profile
	Profile              string   `protobuf:"bytes,13,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFileShareOpts) Reset()         { *m = CreateFileShareOpts{} }
func (m *CreateFileShareOpts) String() string { return proto.CompactTextString(m) }
func (*CreateFileShareOpts) ProtoMessage()    {}
func (*CreateFileShareOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

func (m *CreateFileShareOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFileShareOpts.Unmarshal(m, b)
}
func (m *CreateFileShareOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFileShareOpts.Marshal(b, m, deterministic)
}
func (m *CreateFileShareOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFileShareOpts.Merge(m, src)
}
func (m *CreateFileShareOpts) XXX_Size() int {
	return xxx_messageInfo_CreateFileShareOpts.Size(m)
}
func (m *CreateFileShareOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFileShareOpts.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFileShareOpts proto.InternalMessageInfo

func (m *CreateFileShareOpts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateFileShareOpts) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateFileShareOpts) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *CreateFileShareOpts) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateFileShareOpts) GetAvailabilityZone() string {
	if m != nil {
		return m.AvailabilityZone
	}
	return ""
}

func (m *CreateFileShareOpts) GetProfileId() string {
	if m != nil {
		return m.ProfileId
	}
	return ""
}

func (m *CreateFileShareOpts) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *CreateFileShareOpts) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *CreateFileShareOpts) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CreateFileShareOpts) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

func (m *CreateFileShareOpts) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *CreateFileShareOpts) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

// DeleteFileShareOpts is a structure which indicates all required properties
// for deleting a file share.
type DeleteFileShareOpts struct {
	// The uuid of the fileshare, required.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The service level that fileshare belongs to, required.
	// This item will be replace by profile
	ProfileId string `protobuf:"bytes,2,opt,name=profileId,proto3" json:"profileId,omitempty"`
	// The uuid of the pool on which fileshare will be created, required.
	PoolId string `protobuf:"bytes,3,opt,name=poolId,proto3" json:"poolId,omitempty"`
	// The metadata of the fileshare, optional.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The storage driver type.
	DriverName string `protobuf:"bytes,5,opt,name=driverName,proto3" json:"driverName,omitempty"`
	// The Context
	Context string `protobuf:"bytes,6,opt,name=context,proto3" json:"context,omitempty"`
	// The Serialized profile
	Profile              string   `protobuf:"bytes,7,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFileShareOpts) Reset()         { *m = DeleteFileShareOpts{} }
func (m *DeleteFileShareOpts) String() string { return proto.CompactTextString(m) }
func (*DeleteFileShareOpts) ProtoMessage()    {}
func (*DeleteFileShareOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

func (m *DeleteFileShareOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFileShareOpts.Unmarshal(m, b)
}
func (m *DeleteFileShareOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFileShareOpts.Marshal(b, m, deterministic)
}
func (m *DeleteFileShareOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFileShareOpts.Merge(m, src)
}
func (m *DeleteFileShareOpts) XXX_Size() int {
	return xxx_messageInfo_DeleteFileShareOpts.Size(m)
}
func (m *DeleteFileShareOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFileShareOpts.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFileShareOpts proto.InternalMessageInfo

func (m *DeleteFileShareOpts) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteFileShareOpts) GetProfileId() string {
	if m != nil {
		return m.ProfileId
	}
	return ""
}

func (m *DeleteFileShareOpts) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *DeleteFileShareOpts) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DeleteFileShareOpts) GetDriverName() string {
	if m != nil {
		return m.DriverName
	}
	return ""
}

func (m *DeleteFileShareOpts) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *DeleteFileShareOpts) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

// Generic response, it return:
// 1. Return result with message when create/update resource successfully.
// 2. Return result without message when delete resource successfully.
// 3. Return Error with error code and message when operate unsuccessfully.
type GenericResponse struct {
	// Types that are valid to be assigned to Reply:
	//	*GenericResponse_Result_
	//	*GenericResponse_Error_
	Reply                isGenericResponse_Reply `protobuf_oneof:"reply"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GenericResponse) Reset()         { *m = GenericResponse{} }
func (m *GenericResponse) String() string { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()    {}
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

func (m *GenericResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse.Unmarshal(m, b)
}
func (m *GenericResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse.Marshal(b, m, deterministic)
}
func (m *GenericResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse.Merge(m, src)
}
func (m *GenericResponse) XXX_Size() int {
	return xxx_messageInfo_GenericResponse.Size(m)
}
func (m *GenericResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse proto.InternalMessageInfo

type isGenericResponse_Reply interface {
	isGenericResponse_Reply()
}

type GenericResponse_Result_ struct {
	Result *GenericResponse_Result `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

type GenericResponse_Error_ struct {
	Error *GenericResponse_Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GenericResponse_Result_) isGenericResponse_Reply() {}

func (*GenericResponse_Error_) isGenericResponse_Reply() {}

func (m *GenericResponse) GetReply() isGenericResponse_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *GenericResponse) GetResult() *GenericResponse_Result {
	if x, ok := m.GetReply().(*GenericResponse_Result_); ok {
		return x.Result
	}
	return nil
}

func (m *GenericResponse) GetError() *GenericResponse_Error {
	if x, ok := m.GetReply().(*GenericResponse_Error_); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GenericResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GenericResponse_Result_)(nil),
		(*GenericResponse_Error_)(nil),
	}
}

type GenericResponse_Result struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse_Result) Reset()         { *m = GenericResponse_Result{} }
func (m *GenericResponse_Result) String() string { return proto.CompactTextString(m) }
func (*GenericResponse_Result) ProtoMessage()    {}
func (*GenericResponse_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2, 0}
}

func (m *GenericResponse_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse_Result.Unmarshal(m, b)
}
func (m *GenericResponse_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse_Result.Marshal(b, m, deterministic)
}
func (m *GenericResponse_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse_Result.Merge(m, src)
}
func (m *GenericResponse_Result) XXX_Size() int {
	return xxx_messageInfo_GenericResponse_Result.Size(m)
}
func (m *GenericResponse_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse_Result.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse_Result proto.InternalMessageInfo

func (m *GenericResponse_Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GenericResponse_Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse_Error) Reset()         { *m = GenericResponse_Error{} }
func (m *GenericResponse_Error) String() string { return proto.CompactTextString(m) }
func (*GenericResponse_Error) ProtoMessage()    {}
func (*GenericResponse_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2, 1}
}

func (m *GenericResponse_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse_Error.Unmarshal(m, b)
}
func (m *GenericResponse_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse_Error.Marshal(b, m, deterministic)
}
func (m *GenericResponse_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse_Error.Merge(m, src)
}
func (m *GenericResponse_Error) XXX_Size() int {
	return xxx_messageInfo_GenericResponse_Error.Size(m)
}
func (m *GenericResponse_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse_Error.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse_Error proto.InternalMessageInfo

func (m *GenericResponse_Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *GenericResponse_Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateFileShareOpts)(nil), "fileshareproto.CreateFileShareOpts")
	proto.RegisterMapType((map[string]string)(nil), "fileshareproto.CreateFileShareOpts.MetadataEntry")
	proto.RegisterType((*DeleteFileShareOpts)(nil), "fileshareproto.DeleteFileShareOpts")
	proto.RegisterMapType((map[string]string)(nil), "fileshareproto.DeleteFileShareOpts.MetadataEntry")
	proto.RegisterType((*GenericResponse)(nil), "fileshareproto.GenericResponse")
	proto.RegisterType((*GenericResponse_Result)(nil), "fileshareproto.GenericResponse.Result")
	proto.RegisterType((*GenericResponse_Error)(nil), "fileshareproto.GenericResponse.Error")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0xed, 0xc4, 0x49, 0xc6, 0xbf, 0xa6, 0xd5, 0xf6, 0x27, 0xb4, 0x8a, 0x10, 0x44, 0x41,
	0xa0, 0x88, 0x83, 0x25, 0xcc, 0x05, 0x81, 0x2a, 0x21, 0xfa, 0x87, 0xf6, 0x50, 0x40, 0xe6, 0x04,
	0xb7, 0xad, 0x3d, 0xc0, 0xaa, 0x1b, 0xaf, 0xb5, 0xbb, 0x8d, 0x08, 0x5f, 0x8d, 0x1b, 0x07, 0x6e,
	0x7c, 0x25, 0x84, 0x76, 0xe3, 0x94, 0xc4, 0x49, 0x49, 0x0e, 0x5c, 0xb8, 0xcd, 0x9b, 0x37, 0xf3,
	0xc6, 0xf3, 0xd6, 0x03, 0xd1, 0x58, 0xe6, 0x28, 0xe2, 0x52, 0x49, 0x23, 0x49, 0xef, 0x03, 0x17,
	0xa8, 0x3f, 0x31, 0x85, 0x0e, 0x0f, 0x7f, 0x04, 0xb0, 0x7f, 0xa8, 0x90, 0x19, 0x3c, 0xe1, 0x02,
	0xdf, 0x5a, 0xe2, 0x75, 0x69, 0x34, 0xe9, 0x81, 0xcf, 0x73, 0xea, 0x0d, 0xbc, 0x51, 0x37, 0xf5,
	0x79, 0x4e, 0x08, 0x34, 0x0b, 0x36, 0x46, 0xea, 0xbb, 0x8c, 0x8b, 0x6d, 0x4e, 0xf3, 0x2f, 0x48,
	0x83, 0x81, 0x37, 0x0a, 0x52, 0x17, 0x93, 0x01, 0x44, 0x39, 0xea, 0x4c, 0xf1, 0xd2, 0x70, 0x59,
	0xd0, 0xa6, 0x2b, 0x5f, 0x4c, 0x91, 0x87, 0xb0, 0xc7, 0x26, 0x8c, 0x0b, 0x76, 0xc1, 0x05, 0x37,
	0xd3, 0xf7, 0xb2, 0x40, 0x1a, 0xba, 0xb2, 0x95, 0x3c, 0xb9, 0x0d, 0xdd, 0x52, 0x49, 0xfb, 0xc9,
	0x67, 0x39, 0x6d, 0xbb, 0xa2, 0xdf, 0x09, 0x72, 0x0b, 0xc2, 0x52, 0x4a, 0x71, 0x96, 0xd3, 0x8e,
	0xa3, 0x2a, 0x44, 0xfa, 0xd0, 0xb1, 0xd1, 0x2b, 0xfb, 0xbd, 0x5d, 0xc7, 0x5c, 0x63, 0x72, 0x0e,
	0x9d, 0x31, 0x1a, 0x96, 0x33, 0xc3, 0x28, 0x0c, 0x82, 0x51, 0x94, 0x3c, 0x8a, 0x97, 0x2d, 0x89,
	0xd7, 0xd8, 0x11, 0x9f, 0x57, 0x3d, 0xc7, 0x85, 0x51, 0xd3, 0xf4, 0x5a, 0x82, 0xdc, 0x01, 0xc8,
	0x15, 0x9f, 0xa0, 0x72, 0xc3, 0x22, 0x37, 0x6c, 0x21, 0x43, 0x28, 0xb4, 0x33, 0x59, 0x18, 0xfc,
	0x6c, 0xe8, 0x7f, 0x8e, 0x9c, 0x43, 0xcb, 0x54, 0x9b, 0xd0, 0x9d, 0x19, 0x53, 0xc1, 0xfe, 0x33,
	0xd8, 0x59, 0x1a, 0x47, 0xf6, 0x20, 0xb8, 0xc4, 0x69, 0xf5, 0x18, 0x36, 0x24, 0xff, 0x43, 0x6b,
	0xc2, 0xc4, 0xd5, 0xfc, 0x39, 0x66, 0xe0, 0xa9, 0xff, 0xc4, 0x1b, 0x7e, 0xf5, 0x61, 0xff, 0x08,
	0x05, 0x6e, 0x7a, 0xcf, 0x25, 0x67, 0xfd, 0x9b, 0x9d, 0x0d, 0x96, 0x9c, 0x5d, 0x74, 0xaf, 0xb9,
	0xde, 0xbd, 0x35, 0xc3, 0xb7, 0x74, 0xaf, 0xf5, 0x27, 0xf7, 0xc2, 0x1b, 0xdd, 0x6b, 0xff, 0x45,
	0xf7, 0x7e, 0x7a, 0xb0, 0xfb, 0x12, 0x0b, 0x54, 0x3c, 0x4b, 0x51, 0x97, 0xb2, 0xd0, 0x48, 0x9e,
	0x43, 0xa8, 0x50, 0x5f, 0x09, 0xe3, 0x24, 0xa2, 0xe4, 0x41, 0x7d, 0xe3, 0x5a, 0x43, 0x9c, 0xba,
	0xea, 0xd3, 0x46, 0x5a, 0xf5, 0x91, 0x03, 0x68, 0xa1, 0x52, 0x52, 0xb9, 0x79, 0x51, 0x72, 0x7f,
	0x93, 0xc0, 0xb1, 0x2d, 0x3e, 0x6d, 0xa4, 0xb3, 0xae, 0xfe, 0x10, 0xc2, 0x99, 0xa4, 0xdd, 0x7a,
	0x8c, 0x5a, 0xb3, 0x8f, 0x58, 0xad, 0x33, 0x87, 0xfd, 0x03, 0x68, 0xb9, 0x2e, 0x7b, 0x93, 0x99,
	0xcc, 0xe7, 0xbc, 0x8b, 0xeb, 0x37, 0xe9, 0xaf, 0xdc, 0xe4, 0x8b, 0x36, 0xb4, 0x14, 0x96, 0x62,
	0x9a, 0x7c, 0xf3, 0x00, 0x0e, 0x65, 0x61, 0x94, 0x14, 0x02, 0x15, 0x79, 0x07, 0xbb, 0xb5, 0x6b,
	0x20, 0xf7, 0xb6, 0x38, 0x97, 0xfe, 0xdd, 0x0d, 0x2b, 0x0e, 0x1b, 0x56, 0xba, 0xf6, 0xab, 0xac,
	0x4a, 0xaf, 0xf9, 0x97, 0xb6, 0x90, 0x4e, 0xbe, 0x7b, 0xd0, 0x3b, 0x79, 0xa3, 0xe4, 0x84, 0x6b,
	0x2e, 0x8b, 0x23, 0x99, 0x5d, 0xfe, 0x9b, 0x8b, 0x5c, 0x84, 0x8e, 0x78, 0xfc, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0xab, 0x4a, 0x9e, 0x6a, 0xc2, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerClient interface {
	// Create a file share
	CreateFileShare(ctx context.Context, in *CreateFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error)
	// Delete a file share
	DeleteFileShare(ctx context.Context, in *DeleteFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error)
}

type controllerClient struct {
	cc *grpc.ClientConn
}

func NewControllerClient(cc *grpc.ClientConn) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) CreateFileShare(ctx context.Context, in *CreateFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/fileshareproto.Controller/CreateFileShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) DeleteFileShare(ctx context.Context, in *DeleteFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/fileshareproto.Controller/DeleteFileShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerServer is the server API for Controller service.
type ControllerServer interface {
	// Create a file share
	CreateFileShare(context.Context, *CreateFileShareOpts) (*GenericResponse, error)
	// Delete a file share
	DeleteFileShare(context.Context, *DeleteFileShareOpts) (*GenericResponse, error)
}

// UnimplementedControllerServer can be embedded to have forward compatible implementations.
type UnimplementedControllerServer struct {
}

func (*UnimplementedControllerServer) CreateFileShare(ctx context.Context, req *CreateFileShareOpts) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileShare not implemented")
}
func (*UnimplementedControllerServer) DeleteFileShare(ctx context.Context, req *DeleteFileShareOpts) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileShare not implemented")
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_CreateFileShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileShareOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).CreateFileShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileshareproto.Controller/CreateFileShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).CreateFileShare(ctx, req.(*CreateFileShareOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_DeleteFileShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileShareOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).DeleteFileShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileshareproto.Controller/DeleteFileShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).DeleteFileShare(ctx, req.(*DeleteFileShareOpts))
	}
	return interceptor(ctx, in, info, handler)
}

var _Controller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fileshareproto.Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFileShare",
			Handler:    _Controller_CreateFileShare_Handler,
		},
		{
			MethodName: "DeleteFileShare",
			Handler:    _Controller_DeleteFileShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}

// FProvisionDockClient is the client API for FProvisionDock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FProvisionDockClient interface {
	// Create a file share
	CreateFileShare(ctx context.Context, in *CreateFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error)
	// Delete a file share
	DeleteFileShare(ctx context.Context, in *DeleteFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error)
}

type fProvisionDockClient struct {
	cc *grpc.ClientConn
}

func NewFProvisionDockClient(cc *grpc.ClientConn) FProvisionDockClient {
	return &fProvisionDockClient{cc}
}

func (c *fProvisionDockClient) CreateFileShare(ctx context.Context, in *CreateFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/fileshareproto.FProvisionDock/CreateFileShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fProvisionDockClient) DeleteFileShare(ctx context.Context, in *DeleteFileShareOpts, opts ...grpc.CallOption) (*GenericResponse, error) {
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, "/fileshareproto.FProvisionDock/DeleteFileShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FProvisionDockServer is the server API for FProvisionDock service.
type FProvisionDockServer interface {
	// Create a file share
	CreateFileShare(context.Context, *CreateFileShareOpts) (*GenericResponse, error)
	// Delete a file share
	DeleteFileShare(context.Context, *DeleteFileShareOpts) (*GenericResponse, error)
}

// UnimplementedFProvisionDockServer can be embedded to have forward compatible implementations.
type UnimplementedFProvisionDockServer struct {
}

func (*UnimplementedFProvisionDockServer) CreateFileShare(ctx context.Context, req *CreateFileShareOpts) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileShare not implemented")
}
func (*UnimplementedFProvisionDockServer) DeleteFileShare(ctx context.Context, req *DeleteFileShareOpts) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileShare not implemented")
}

func RegisterFProvisionDockServer(s *grpc.Server, srv FProvisionDockServer) {
	s.RegisterService(&_FProvisionDock_serviceDesc, srv)
}

func _FProvisionDock_CreateFileShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileShareOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FProvisionDockServer).CreateFileShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileshareproto.FProvisionDock/CreateFileShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FProvisionDockServer).CreateFileShare(ctx, req.(*CreateFileShareOpts))
	}
	return interceptor(ctx, in, info, handler)
}

func _FProvisionDock_DeleteFileShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileShareOpts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FProvisionDockServer).DeleteFileShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fileshareproto.FProvisionDock/DeleteFileShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FProvisionDockServer).DeleteFileShare(ctx, req.(*DeleteFileShareOpts))
	}
	return interceptor(ctx, in, info, handler)
}

var _FProvisionDock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fileshareproto.FProvisionDock",
	HandlerType: (*FProvisionDockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFileShare",
			Handler:    _FProvisionDock_CreateFileShare_Handler,
		},
		{
			MethodName: "DeleteFileShare",
			Handler:    _FProvisionDock_DeleteFileShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}
