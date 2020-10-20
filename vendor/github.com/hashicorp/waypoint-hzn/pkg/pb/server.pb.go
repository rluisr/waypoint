// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: server.proto

package pb

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type LabelSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels []*Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *LabelSet) Reset() {
	*x = LabelSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSet) ProtoMessage() {}

func (x *LabelSet) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSet.ProtoReflect.Descriptor instead.
func (*LabelSet) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *LabelSet) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

type RegisterGuestAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// server ID is the unique ULID of the Waypoint server requesting a
	// guest account. If this server already has a guest account registered,
	// the same token will be returned.
	ServerId string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Indicates that the user accepted the EULA to access Horizon
	AcceptTos bool `protobuf:"varint,2,opt,name=accept_tos,json=acceptTos,proto3" json:"accept_tos,omitempty"`
}

func (x *RegisterGuestAccountRequest) Reset() {
	*x = RegisterGuestAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterGuestAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterGuestAccountRequest) ProtoMessage() {}

func (x *RegisterGuestAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterGuestAccountRequest.ProtoReflect.Descriptor instead.
func (*RegisterGuestAccountRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterGuestAccountRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *RegisterGuestAccountRequest) GetAcceptTos() bool {
	if x != nil {
		return x.AcceptTos
	}
	return false
}

type RegisterGuestAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// API token to use for protected endpoints.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RegisterGuestAccountResponse) Reset() {
	*x = RegisterGuestAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterGuestAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterGuestAccountResponse) ProtoMessage() {}

func (x *RegisterGuestAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterGuestAccountResponse.ProtoReflect.Descriptor instead.
func (*RegisterGuestAccountResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterGuestAccountResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RegisterHostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hostname to register
	//
	// Types that are assignable to Hostname:
	//	*RegisterHostnameRequest_Generate
	//	*RegisterHostnameRequest_Exact
	Hostname isRegisterHostnameRequest_Hostname `protobuf_oneof:"hostname"`
	// labels to link this hostname to.
	Labels *LabelSet `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
}

func (x *RegisterHostnameRequest) Reset() {
	*x = RegisterHostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterHostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterHostnameRequest) ProtoMessage() {}

func (x *RegisterHostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterHostnameRequest.ProtoReflect.Descriptor instead.
func (*RegisterHostnameRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (m *RegisterHostnameRequest) GetHostname() isRegisterHostnameRequest_Hostname {
	if m != nil {
		return m.Hostname
	}
	return nil
}

func (x *RegisterHostnameRequest) GetGenerate() *empty.Empty {
	if x, ok := x.GetHostname().(*RegisterHostnameRequest_Generate); ok {
		return x.Generate
	}
	return nil
}

func (x *RegisterHostnameRequest) GetExact() string {
	if x, ok := x.GetHostname().(*RegisterHostnameRequest_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *RegisterHostnameRequest) GetLabels() *LabelSet {
	if x != nil {
		return x.Labels
	}
	return nil
}

type isRegisterHostnameRequest_Hostname interface {
	isRegisterHostnameRequest_Hostname()
}

type RegisterHostnameRequest_Generate struct {
	// auto-generate a hostname
	Generate *empty.Empty `protobuf:"bytes,1,opt,name=generate,proto3,oneof"`
}

type RegisterHostnameRequest_Exact struct {
	// specific hostname request
	Exact string `protobuf:"bytes,2,opt,name=exact,proto3,oneof"`
}

func (*RegisterHostnameRequest_Generate) isRegisterHostnameRequest_Hostname() {}

func (*RegisterHostnameRequest_Exact) isRegisterHostnameRequest_Hostname() {}

type RegisterHostnameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Fqdn     string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
}

func (x *RegisterHostnameResponse) Reset() {
	*x = RegisterHostnameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterHostnameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterHostnameResponse) ProtoMessage() {}

func (x *RegisterHostnameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterHostnameResponse.ProtoReflect.Descriptor instead.
func (*RegisterHostnameResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterHostnameResponse) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *RegisterHostnameResponse) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

type ListHostnamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHostnamesRequest) Reset() {
	*x = ListHostnamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostnamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostnamesRequest) ProtoMessage() {}

func (x *ListHostnamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostnamesRequest.ProtoReflect.Descriptor instead.
func (*ListHostnamesRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{6}
}

type ListHostnamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostnames []*ListHostnamesResponse_Hostname `protobuf:"bytes,1,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
}

func (x *ListHostnamesResponse) Reset() {
	*x = ListHostnamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostnamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostnamesResponse) ProtoMessage() {}

func (x *ListHostnamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostnamesResponse.ProtoReflect.Descriptor instead.
func (*ListHostnamesResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{7}
}

func (x *ListHostnamesResponse) GetHostnames() []*ListHostnamesResponse_Hostname {
	if x != nil {
		return x.Hostnames
	}
	return nil
}

type DeleteHostnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *DeleteHostnameRequest) Reset() {
	*x = DeleteHostnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHostnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHostnameRequest) ProtoMessage() {}

func (x *DeleteHostnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHostnameRequest.ProtoReflect.Descriptor instead.
func (*DeleteHostnameRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteHostnameRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type ListHostnamesResponse_Hostname struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname string    `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Fqdn     string    `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Labels   *LabelSet `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
}

func (x *ListHostnamesResponse_Hostname) Reset() {
	*x = ListHostnamesResponse_Hostname{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostnamesResponse_Hostname) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostnamesResponse_Hostname) ProtoMessage() {}

func (x *ListHostnamesResponse_Hostname) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostnamesResponse_Hostname.ProtoReflect.Descriptor instead.
func (*ListHostnamesResponse_Hostname) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListHostnamesResponse_Hostname) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ListHostnamesResponse_Hostname) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *ListHostnamesResponse_Hostname) GetLabels() *LabelSet {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x64, 0x65, 0x66, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x4b, 0x0a, 0x08, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x59, 0x0a, 0x1b, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x74,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x54, 0x6f, 0x73, 0x22, 0x34, 0x0a, 0x1c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xfc, 0x01, 0x0a, 0x17, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48,
	0x00, 0x52, 0x08, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x56, 0x0a, 0x05, 0x65,
	0x78, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3e, 0xfa, 0x42, 0x3b, 0x72,
	0x39, 0x10, 0x03, 0x32, 0x0b, 0x5c, 0x77, 0x2b, 0x5b, 0x5c, 0x77, 0x5c, 0x64, 0x2d, 0x5d, 0x2a,
	0x5a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5a, 0x03, 0x61, 0x70, 0x69, 0x5a, 0x04, 0x62, 0x6c,
	0x6f, 0x67, 0x5a, 0x03, 0x68, 0x7a, 0x6e, 0x5a, 0x07, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e,
	0x5a, 0x08, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78,
	0x61, 0x63, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x53, 0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x42, 0x0f, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x4a, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x71, 0x64, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xe3, 0x01, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68,
	0x7a, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x74, 0x0a, 0x08,
	0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a,
	0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x22, 0x33, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xcf, 0x03, 0x0a, 0x0b, 0x57, 0x61, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x48, 0x7a, 0x6e, 0x12, 0x81, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x33, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x2c, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77, 0x61,
	0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x57, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x77,
	0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x68, 0x7a, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x04, 0x5a, 0x02, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_server_proto_goTypes = []interface{}{
	(*Label)(nil),                          // 0: hashicorp.waypoint_hzn.Label
	(*LabelSet)(nil),                       // 1: hashicorp.waypoint_hzn.LabelSet
	(*RegisterGuestAccountRequest)(nil),    // 2: hashicorp.waypoint_hzn.RegisterGuestAccountRequest
	(*RegisterGuestAccountResponse)(nil),   // 3: hashicorp.waypoint_hzn.RegisterGuestAccountResponse
	(*RegisterHostnameRequest)(nil),        // 4: hashicorp.waypoint_hzn.RegisterHostnameRequest
	(*RegisterHostnameResponse)(nil),       // 5: hashicorp.waypoint_hzn.RegisterHostnameResponse
	(*ListHostnamesRequest)(nil),           // 6: hashicorp.waypoint_hzn.ListHostnamesRequest
	(*ListHostnamesResponse)(nil),          // 7: hashicorp.waypoint_hzn.ListHostnamesResponse
	(*DeleteHostnameRequest)(nil),          // 8: hashicorp.waypoint_hzn.DeleteHostnameRequest
	(*ListHostnamesResponse_Hostname)(nil), // 9: hashicorp.waypoint_hzn.ListHostnamesResponse.Hostname
	(*empty.Empty)(nil),                    // 10: google.protobuf.Empty
}
var file_server_proto_depIdxs = []int32{
	0,  // 0: hashicorp.waypoint_hzn.LabelSet.labels:type_name -> hashicorp.waypoint_hzn.Label
	10, // 1: hashicorp.waypoint_hzn.RegisterHostnameRequest.generate:type_name -> google.protobuf.Empty
	1,  // 2: hashicorp.waypoint_hzn.RegisterHostnameRequest.labels:type_name -> hashicorp.waypoint_hzn.LabelSet
	9,  // 3: hashicorp.waypoint_hzn.ListHostnamesResponse.hostnames:type_name -> hashicorp.waypoint_hzn.ListHostnamesResponse.Hostname
	1,  // 4: hashicorp.waypoint_hzn.ListHostnamesResponse.Hostname.labels:type_name -> hashicorp.waypoint_hzn.LabelSet
	2,  // 5: hashicorp.waypoint_hzn.WaypointHzn.RegisterGuestAccount:input_type -> hashicorp.waypoint_hzn.RegisterGuestAccountRequest
	4,  // 6: hashicorp.waypoint_hzn.WaypointHzn.RegisterHostname:input_type -> hashicorp.waypoint_hzn.RegisterHostnameRequest
	6,  // 7: hashicorp.waypoint_hzn.WaypointHzn.ListHostnames:input_type -> hashicorp.waypoint_hzn.ListHostnamesRequest
	8,  // 8: hashicorp.waypoint_hzn.WaypointHzn.DeleteHostname:input_type -> hashicorp.waypoint_hzn.DeleteHostnameRequest
	3,  // 9: hashicorp.waypoint_hzn.WaypointHzn.RegisterGuestAccount:output_type -> hashicorp.waypoint_hzn.RegisterGuestAccountResponse
	5,  // 10: hashicorp.waypoint_hzn.WaypointHzn.RegisterHostname:output_type -> hashicorp.waypoint_hzn.RegisterHostnameResponse
	7,  // 11: hashicorp.waypoint_hzn.WaypointHzn.ListHostnames:output_type -> hashicorp.waypoint_hzn.ListHostnamesResponse
	10, // 12: hashicorp.waypoint_hzn.WaypointHzn.DeleteHostname:output_type -> google.protobuf.Empty
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterGuestAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterGuestAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterHostnameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterHostnameResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostnamesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostnamesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHostnameRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostnamesResponse_Hostname); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_server_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*RegisterHostnameRequest_Generate)(nil),
		(*RegisterHostnameRequest_Exact)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WaypointHznClient is the client API for WaypointHzn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WaypointHznClient interface {
	RegisterGuestAccount(ctx context.Context, in *RegisterGuestAccountRequest, opts ...grpc.CallOption) (*RegisterGuestAccountResponse, error)
	RegisterHostname(ctx context.Context, in *RegisterHostnameRequest, opts ...grpc.CallOption) (*RegisterHostnameResponse, error)
	ListHostnames(ctx context.Context, in *ListHostnamesRequest, opts ...grpc.CallOption) (*ListHostnamesResponse, error)
	DeleteHostname(ctx context.Context, in *DeleteHostnameRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type waypointHznClient struct {
	cc grpc.ClientConnInterface
}

func NewWaypointHznClient(cc grpc.ClientConnInterface) WaypointHznClient {
	return &waypointHznClient{cc}
}

func (c *waypointHznClient) RegisterGuestAccount(ctx context.Context, in *RegisterGuestAccountRequest, opts ...grpc.CallOption) (*RegisterGuestAccountResponse, error) {
	out := new(RegisterGuestAccountResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/RegisterGuestAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waypointHznClient) RegisterHostname(ctx context.Context, in *RegisterHostnameRequest, opts ...grpc.CallOption) (*RegisterHostnameResponse, error) {
	out := new(RegisterHostnameResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/RegisterHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waypointHznClient) ListHostnames(ctx context.Context, in *ListHostnamesRequest, opts ...grpc.CallOption) (*ListHostnamesResponse, error) {
	out := new(ListHostnamesResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/ListHostnames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *waypointHznClient) DeleteHostname(ctx context.Context, in *DeleteHostnameRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/hashicorp.waypoint_hzn.WaypointHzn/DeleteHostname", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaypointHznServer is the server API for WaypointHzn service.
type WaypointHznServer interface {
	RegisterGuestAccount(context.Context, *RegisterGuestAccountRequest) (*RegisterGuestAccountResponse, error)
	RegisterHostname(context.Context, *RegisterHostnameRequest) (*RegisterHostnameResponse, error)
	ListHostnames(context.Context, *ListHostnamesRequest) (*ListHostnamesResponse, error)
	DeleteHostname(context.Context, *DeleteHostnameRequest) (*empty.Empty, error)
}

// UnimplementedWaypointHznServer can be embedded to have forward compatible implementations.
type UnimplementedWaypointHznServer struct {
}

func (*UnimplementedWaypointHznServer) RegisterGuestAccount(context.Context, *RegisterGuestAccountRequest) (*RegisterGuestAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterGuestAccount not implemented")
}
func (*UnimplementedWaypointHznServer) RegisterHostname(context.Context, *RegisterHostnameRequest) (*RegisterHostnameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterHostname not implemented")
}
func (*UnimplementedWaypointHznServer) ListHostnames(context.Context, *ListHostnamesRequest) (*ListHostnamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHostnames not implemented")
}
func (*UnimplementedWaypointHznServer) DeleteHostname(context.Context, *DeleteHostnameRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHostname not implemented")
}

func RegisterWaypointHznServer(s *grpc.Server, srv WaypointHznServer) {
	s.RegisterService(&_WaypointHzn_serviceDesc, srv)
}

func _WaypointHzn_RegisterGuestAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterGuestAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).RegisterGuestAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/RegisterGuestAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).RegisterGuestAccount(ctx, req.(*RegisterGuestAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaypointHzn_RegisterHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterHostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).RegisterHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/RegisterHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).RegisterHostname(ctx, req.(*RegisterHostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaypointHzn_ListHostnames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHostnamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).ListHostnames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/ListHostnames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).ListHostnames(ctx, req.(*ListHostnamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WaypointHzn_DeleteHostname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaypointHznServer).DeleteHostname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.waypoint_hzn.WaypointHzn/DeleteHostname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaypointHznServer).DeleteHostname(ctx, req.(*DeleteHostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WaypointHzn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.waypoint_hzn.WaypointHzn",
	HandlerType: (*WaypointHznServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterGuestAccount",
			Handler:    _WaypointHzn_RegisterGuestAccount_Handler,
		},
		{
			MethodName: "RegisterHostname",
			Handler:    _WaypointHzn_RegisterHostname_Handler,
		},
		{
			MethodName: "ListHostnames",
			Handler:    _WaypointHzn_ListHostnames_Handler,
		},
		{
			MethodName: "DeleteHostname",
			Handler:    _WaypointHzn_DeleteHostname_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}