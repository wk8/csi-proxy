// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
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

// Context of the paths used for path prefix validation
type PathContext int32

const (
	// Indicates the kubelet-csi-plugins-path parameter of csi-proxy be used as the path context
	PathContext_PLUGIN PathContext = 0
	// Indicates the kubelet-pod-path parameter of csi-proxy be used as the path context
	PathContext_CONTAINER PathContext = 1
)

var PathContext_name = map[int32]string{
	0: "PLUGIN",
	1: "CONTAINER",
}

var PathContext_value = map[string]int32{
	"PLUGIN":    0,
	"CONTAINER": 1,
}

func (x PathContext) String() string {
	return proto.EnumName(PathContext_name, int32(x))
}

func (PathContext) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type PathExistsRequest struct {
	// The path whose existence we want to check in the host's filesystem
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Context of the path parameter.
	// This is used to determine the root for relative path parameters
	Context              PathContext `protobuf:"varint,2,opt,name=context,proto3,enum=v1alpha1.PathContext" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PathExistsRequest) Reset()         { *m = PathExistsRequest{} }
func (m *PathExistsRequest) String() string { return proto.CompactTextString(m) }
func (*PathExistsRequest) ProtoMessage()    {}
func (*PathExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *PathExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathExistsRequest.Unmarshal(m, b)
}
func (m *PathExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathExistsRequest.Marshal(b, m, deterministic)
}
func (m *PathExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathExistsRequest.Merge(m, src)
}
func (m *PathExistsRequest) XXX_Size() int {
	return xxx_messageInfo_PathExistsRequest.Size(m)
}
func (m *PathExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PathExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PathExistsRequest proto.InternalMessageInfo

func (m *PathExistsRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathExistsRequest) GetContext() PathContext {
	if m != nil {
		return m.Context
	}
	return PathContext_PLUGIN
}

type PathExistsResponse struct {
	// Error message if any. Empty string indicates success
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// Indicates whether the path in PathExistsRequest exists in the host's filesystem
	Exists               bool     `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathExistsResponse) Reset()         { *m = PathExistsResponse{} }
func (m *PathExistsResponse) String() string { return proto.CompactTextString(m) }
func (*PathExistsResponse) ProtoMessage()    {}
func (*PathExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *PathExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathExistsResponse.Unmarshal(m, b)
}
func (m *PathExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathExistsResponse.Marshal(b, m, deterministic)
}
func (m *PathExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathExistsResponse.Merge(m, src)
}
func (m *PathExistsResponse) XXX_Size() int {
	return xxx_messageInfo_PathExistsResponse.Size(m)
}
func (m *PathExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PathExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PathExistsResponse proto.InternalMessageInfo

func (m *PathExistsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PathExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type MkdirRequest struct {
	// The path to create in the host's filesystem.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	// Non-existent parent directories in the path will be automatically created.
	// Directories will be created with Read and Write privileges of the Windows
	// User account under which csi-proxy is started (typically LocalSystem).
	//
	// Restrictions:
	// If an absolute path (indicated by a drive letter prefix: e.g. "C:\") is passed,
	// depending on the context parameter of this function, the path prefix needs
	// to match the paths specified either as kubelet-csi-plugins-path
	// or as kubelet-pod-path parameters of csi-proxy.
	// If a relative path is passed, depending on the context parameter of this
	// function, the path will be considered relative to the path specified either as
	// kubelet-csi-plugins-path or as kubelet-pod-path parameters of csi-proxy.
	// The path parameter cannot already exist on host filesystem.
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// Maximum path length will be capped to 260 characters.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Context of the path parameter.
	// This is used to [1] determine the root for relative path parameters
	// or [2] validate prefix for absolute paths (indicated by a drive letter
	// prefix: e.g. "C:\")
	Context              PathContext `protobuf:"varint,2,opt,name=context,proto3,enum=v1alpha1.PathContext" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MkdirRequest) Reset()         { *m = MkdirRequest{} }
func (m *MkdirRequest) String() string { return proto.CompactTextString(m) }
func (*MkdirRequest) ProtoMessage()    {}
func (*MkdirRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *MkdirRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MkdirRequest.Unmarshal(m, b)
}
func (m *MkdirRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MkdirRequest.Marshal(b, m, deterministic)
}
func (m *MkdirRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MkdirRequest.Merge(m, src)
}
func (m *MkdirRequest) XXX_Size() int {
	return xxx_messageInfo_MkdirRequest.Size(m)
}
func (m *MkdirRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MkdirRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MkdirRequest proto.InternalMessageInfo

func (m *MkdirRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MkdirRequest) GetContext() PathContext {
	if m != nil {
		return m.Context
	}
	return PathContext_PLUGIN
}

type MkdirResponse struct {
	// Error message if any. Empty string indicates success
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MkdirResponse) Reset()         { *m = MkdirResponse{} }
func (m *MkdirResponse) String() string { return proto.CompactTextString(m) }
func (*MkdirResponse) ProtoMessage()    {}
func (*MkdirResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *MkdirResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MkdirResponse.Unmarshal(m, b)
}
func (m *MkdirResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MkdirResponse.Marshal(b, m, deterministic)
}
func (m *MkdirResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MkdirResponse.Merge(m, src)
}
func (m *MkdirResponse) XXX_Size() int {
	return xxx_messageInfo_MkdirResponse.Size(m)
}
func (m *MkdirResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MkdirResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MkdirResponse proto.InternalMessageInfo

func (m *MkdirResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RmdirRequest struct {
	// The path to remove in the host's filesystem.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// If an absolute path (indicated by a drive letter prefix: e.g. "C:\") is passed,
	// depending on the context parameter of this function, the path prefix needs
	// to match the paths specified either as kubelet-csi-plugins-path
	// or as kubelet-pod-path parameters of csi-proxy.
	// If a relative path is passed, depending on the context parameter of this
	// function, the path will be considered relative to the path specified either as
	// kubelet-csi-plugins-path or as kubelet-pod-path parameters of csi-proxy.
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// Path cannot be a file of type symlink.
	// Maximum path length will be capped to 260 characters.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Context of the path creation used for path prefix validation
	// This is used to [1] determine the root for relative path parameters
	// or [2] validate prefix for absolute paths (indicated by a drive letter
	// prefix: e.g. "C:\")
	Context              PathContext `protobuf:"varint,2,opt,name=context,proto3,enum=v1alpha1.PathContext" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RmdirRequest) Reset()         { *m = RmdirRequest{} }
func (m *RmdirRequest) String() string { return proto.CompactTextString(m) }
func (*RmdirRequest) ProtoMessage()    {}
func (*RmdirRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *RmdirRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmdirRequest.Unmarshal(m, b)
}
func (m *RmdirRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmdirRequest.Marshal(b, m, deterministic)
}
func (m *RmdirRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmdirRequest.Merge(m, src)
}
func (m *RmdirRequest) XXX_Size() int {
	return xxx_messageInfo_RmdirRequest.Size(m)
}
func (m *RmdirRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RmdirRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RmdirRequest proto.InternalMessageInfo

func (m *RmdirRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RmdirRequest) GetContext() PathContext {
	if m != nil {
		return m.Context
	}
	return PathContext_PLUGIN
}

type RmdirResponse struct {
	// Error message if any. Empty string indicates success
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RmdirResponse) Reset()         { *m = RmdirResponse{} }
func (m *RmdirResponse) String() string { return proto.CompactTextString(m) }
func (*RmdirResponse) ProtoMessage()    {}
func (*RmdirResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *RmdirResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RmdirResponse.Unmarshal(m, b)
}
func (m *RmdirResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RmdirResponse.Marshal(b, m, deterministic)
}
func (m *RmdirResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RmdirResponse.Merge(m, src)
}
func (m *RmdirResponse) XXX_Size() int {
	return xxx_messageInfo_RmdirResponse.Size(m)
}
func (m *RmdirResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RmdirResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RmdirResponse proto.InternalMessageInfo

func (m *RmdirResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LinkPathRequest struct {
	// The path where the symlink is created in the host's filesystem.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// If an absolute path (indicated by a drive letter prefix: e.g. "C:\") is passed,
	// the path prefix needs to match the path specified as kubelet-csi-plugins-path
	// parameter of csi-proxy.
	// If a relative path is passed, the path will be considered relative to the
	// path specified as kubelet-csi-plugins-path parameter of csi-proxy.
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// source_path cannot already exist in the host filesystem.
	// Maximum path length will be capped to 260 characters.
	SourcePath string `protobuf:"bytes,1,opt,name=source_path,json=sourcePath,proto3" json:"source_path,omitempty"`
	// Target path in the host's filesystem used for the symlink creation.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// If an absolute path (indicated by a drive letter prefix: e.g. "C:\") is passed,
	// the path prefix needs to match the path specified as kubelet-pod-path
	// parameter of csi-proxy.
	// If a relative path is passed, the path will be considered relative to the
	// path specified as kubelet-pod-path parameter of csi-proxy.
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// target_path needs to exist as a directory in the host that is empty.
	// target_path cannot be a symbolic link.
	// Maximum path length will be capped to 260 characters.
	TargetPath           string   `protobuf:"bytes,2,opt,name=target_path,json=targetPath,proto3" json:"target_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkPathRequest) Reset()         { *m = LinkPathRequest{} }
func (m *LinkPathRequest) String() string { return proto.CompactTextString(m) }
func (*LinkPathRequest) ProtoMessage()    {}
func (*LinkPathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *LinkPathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkPathRequest.Unmarshal(m, b)
}
func (m *LinkPathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkPathRequest.Marshal(b, m, deterministic)
}
func (m *LinkPathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkPathRequest.Merge(m, src)
}
func (m *LinkPathRequest) XXX_Size() int {
	return xxx_messageInfo_LinkPathRequest.Size(m)
}
func (m *LinkPathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkPathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LinkPathRequest proto.InternalMessageInfo

func (m *LinkPathRequest) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *LinkPathRequest) GetTargetPath() string {
	if m != nil {
		return m.TargetPath
	}
	return ""
}

type LinkPathResponse struct {
	// Error message if any. Empty string indicates success
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkPathResponse) Reset()         { *m = LinkPathResponse{} }
func (m *LinkPathResponse) String() string { return proto.CompactTextString(m) }
func (*LinkPathResponse) ProtoMessage()    {}
func (*LinkPathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *LinkPathResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkPathResponse.Unmarshal(m, b)
}
func (m *LinkPathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkPathResponse.Marshal(b, m, deterministic)
}
func (m *LinkPathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkPathResponse.Merge(m, src)
}
func (m *LinkPathResponse) XXX_Size() int {
	return xxx_messageInfo_LinkPathResponse.Size(m)
}
func (m *LinkPathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkPathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LinkPathResponse proto.InternalMessageInfo

func (m *LinkPathResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("v1alpha1.PathContext", PathContext_name, PathContext_value)
	proto.RegisterType((*PathExistsRequest)(nil), "v1alpha1.PathExistsRequest")
	proto.RegisterType((*PathExistsResponse)(nil), "v1alpha1.PathExistsResponse")
	proto.RegisterType((*MkdirRequest)(nil), "v1alpha1.MkdirRequest")
	proto.RegisterType((*MkdirResponse)(nil), "v1alpha1.MkdirResponse")
	proto.RegisterType((*RmdirRequest)(nil), "v1alpha1.RmdirRequest")
	proto.RegisterType((*RmdirResponse)(nil), "v1alpha1.RmdirResponse")
	proto.RegisterType((*LinkPathRequest)(nil), "v1alpha1.LinkPathRequest")
	proto.RegisterType((*LinkPathResponse)(nil), "v1alpha1.LinkPathResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x5f, 0x4b, 0x02, 0x41,
	0x14, 0xc5, 0x5d, 0x49, 0xd3, 0x6b, 0x96, 0x5d, 0xca, 0x6c, 0x0b, 0x92, 0x85, 0x60, 0xe9, 0xc1,
	0xd0, 0xde, 0x7a, 0x2b, 0xb1, 0x10, 0xcc, 0x64, 0x2b, 0xe8, 0x2d, 0x36, 0x1b, 0x72, 0xf1, 0xcf,
	0x6c, 0x33, 0x63, 0xd8, 0xf7, 0xe8, 0x03, 0xc7, 0xcc, 0xec, 0xb2, 0xb3, 0x25, 0xf6, 0x52, 0x6f,
	0x3b, 0x73, 0xcf, 0xfd, 0xdd, 0xc3, 0x3d, 0xb3, 0x50, 0xf4, 0xc3, 0xa0, 0x11, 0x32, 0x2a, 0x28,
	0x16, 0xde, 0x9b, 0xfe, 0x24, 0x1c, 0xf9, 0x4d, 0xe7, 0x11, 0xb6, 0x07, 0xbe, 0x18, 0x75, 0x16,
	0x01, 0x17, 0xdc, 0x23, 0x6f, 0x73, 0xc2, 0x05, 0x22, 0xac, 0x85, 0xbe, 0x18, 0xd5, 0xac, 0xba,
	0xe5, 0x16, 0x3d, 0xf5, 0x8d, 0xa7, 0xb0, 0x3e, 0xa4, 0x33, 0x41, 0x16, 0xa2, 0x96, 0xad, 0x5b,
	0xee, 0x66, 0x6b, 0xb7, 0x11, 0x43, 0x1a, 0x92, 0xd0, 0xd6, 0x45, 0x2f, 0x56, 0x39, 0x97, 0x80,
	0x26, 0x99, 0x87, 0x74, 0xc6, 0x09, 0xee, 0x40, 0x8e, 0x30, 0x46, 0x59, 0xc4, 0xd6, 0x07, 0xac,
	0x42, 0x9e, 0x28, 0x9d, 0x62, 0x17, 0xbc, 0xe8, 0xe4, 0xdc, 0xc1, 0xc6, 0xcd, 0xf8, 0x25, 0x60,
	0x7f, 0x6a, 0xec, 0x18, 0xca, 0x11, 0x74, 0x95, 0x27, 0x39, 0xdb, 0x9b, 0xfe, 0xc3, 0xec, 0x08,
	0xfa, 0xcb, 0xec, 0xad, 0x5e, 0x30, 0x1b, 0x4b, 0x44, 0x3c, 0xfe, 0x08, 0x4a, 0x9c, 0xce, 0xd9,
	0x90, 0x3c, 0x19, 0x2e, 0x40, 0x5f, 0x49, 0x9d, 0x14, 0x08, 0x9f, 0xbd, 0x12, 0xa1, 0x05, 0x59,
	0x2d, 0xd0, 0x57, 0x52, 0xe0, 0xb8, 0x50, 0x49, 0xa0, 0xab, 0xc6, 0x9f, 0xb8, 0x50, 0x32, 0xdc,
	0x23, 0x40, 0x7e, 0xd0, 0x7b, 0xb8, 0xee, 0xf6, 0x2b, 0x19, 0x2c, 0x43, 0xb1, 0x7d, 0xdb, 0xbf,
	0xbf, 0xe8, 0xf6, 0x3b, 0x5e, 0xc5, 0x6a, 0x7d, 0x66, 0x01, 0xae, 0x82, 0x09, 0xe1, 0x1f, 0x5c,
	0x90, 0x29, 0x76, 0x01, 0x92, 0xcc, 0xf1, 0x20, 0xbd, 0x8c, 0xd4, 0x1b, 0xb3, 0x0f, 0x97, 0x17,
	0xb5, 0x2f, 0x27, 0x83, 0xe7, 0x90, 0x53, 0x29, 0x61, 0x35, 0x11, 0x9a, 0x6f, 0xc1, 0xde, 0xfb,
	0x71, 0x6f, 0xf6, 0xaa, 0x2d, 0x9b, 0xbd, 0x66, 0x96, 0x66, 0x6f, 0x2a, 0x0e, 0x27, 0x83, 0x6d,
	0x28, 0xc4, 0x5b, 0xc2, 0xfd, 0x44, 0xf6, 0x2d, 0x0e, 0xdb, 0x5e, 0x56, 0x8a, 0x21, 0xcf, 0x79,
	0xf5, 0x9b, 0x9d, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x62, 0xea, 0x63, 0x73, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FilesystemClient is the client API for Filesystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FilesystemClient interface {
	// PathExists checks if the requested path exists in the host's filesystem
	PathExists(ctx context.Context, in *PathExistsRequest, opts ...grpc.CallOption) (*PathExistsResponse, error)
	// Mkdir creates a directory at the requested path in the host's filesystem
	Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*MkdirResponse, error)
	// Rmdir removes the directory at the requested path in the host's filesystem.
	// This may be used for unlinking a symlink created through LinkPath
	Rmdir(ctx context.Context, in *RmdirRequest, opts ...grpc.CallOption) (*RmdirResponse, error)
	// LinkPath creates a local directory symbolic link between a source path
	// and target path in the host's filesystem
	LinkPath(ctx context.Context, in *LinkPathRequest, opts ...grpc.CallOption) (*LinkPathResponse, error)
}

type filesystemClient struct {
	cc *grpc.ClientConn
}

func NewFilesystemClient(cc *grpc.ClientConn) FilesystemClient {
	return &filesystemClient{cc}
}

func (c *filesystemClient) PathExists(ctx context.Context, in *PathExistsRequest, opts ...grpc.CallOption) (*PathExistsResponse, error) {
	out := new(PathExistsResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Filesystem/PathExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemClient) Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*MkdirResponse, error) {
	out := new(MkdirResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Filesystem/Mkdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemClient) Rmdir(ctx context.Context, in *RmdirRequest, opts ...grpc.CallOption) (*RmdirResponse, error) {
	out := new(RmdirResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Filesystem/Rmdir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesystemClient) LinkPath(ctx context.Context, in *LinkPathRequest, opts ...grpc.CallOption) (*LinkPathResponse, error) {
	out := new(LinkPathResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Filesystem/LinkPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesystemServer is the server API for Filesystem service.
type FilesystemServer interface {
	// PathExists checks if the requested path exists in the host's filesystem
	PathExists(context.Context, *PathExistsRequest) (*PathExistsResponse, error)
	// Mkdir creates a directory at the requested path in the host's filesystem
	Mkdir(context.Context, *MkdirRequest) (*MkdirResponse, error)
	// Rmdir removes the directory at the requested path in the host's filesystem.
	// This may be used for unlinking a symlink created through LinkPath
	Rmdir(context.Context, *RmdirRequest) (*RmdirResponse, error)
	// LinkPath creates a local directory symbolic link between a source path
	// and target path in the host's filesystem
	LinkPath(context.Context, *LinkPathRequest) (*LinkPathResponse, error)
}

// UnimplementedFilesystemServer can be embedded to have forward compatible implementations.
type UnimplementedFilesystemServer struct {
}

func (*UnimplementedFilesystemServer) PathExists(ctx context.Context, req *PathExistsRequest) (*PathExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PathExists not implemented")
}
func (*UnimplementedFilesystemServer) Mkdir(ctx context.Context, req *MkdirRequest) (*MkdirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mkdir not implemented")
}
func (*UnimplementedFilesystemServer) Rmdir(ctx context.Context, req *RmdirRequest) (*RmdirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rmdir not implemented")
}
func (*UnimplementedFilesystemServer) LinkPath(ctx context.Context, req *LinkPathRequest) (*LinkPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkPath not implemented")
}

func RegisterFilesystemServer(s *grpc.Server, srv FilesystemServer) {
	s.RegisterService(&_Filesystem_serviceDesc, srv)
}

func _Filesystem_PathExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PathExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServer).PathExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Filesystem/PathExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServer).PathExists(ctx, req.(*PathExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filesystem_Mkdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServer).Mkdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Filesystem/Mkdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServer).Mkdir(ctx, req.(*MkdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filesystem_Rmdir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RmdirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServer).Rmdir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Filesystem/Rmdir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServer).Rmdir(ctx, req.(*RmdirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filesystem_LinkPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesystemServer).LinkPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Filesystem/LinkPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesystemServer).LinkPath(ctx, req.(*LinkPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Filesystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Filesystem",
	HandlerType: (*FilesystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PathExists",
			Handler:    _Filesystem_PathExists_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _Filesystem_Mkdir_Handler,
		},
		{
			MethodName: "Rmdir",
			Handler:    _Filesystem_Rmdir_Handler,
		},
		{
			MethodName: "LinkPath",
			Handler:    _Filesystem_LinkPath_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
