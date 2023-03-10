// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: apps/policy/pb/policy.proto

package policy

import (
	request "github.com/infraboard/mcube/http/request"
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

// 鉴权请求：用户能不发操作该资源
type ValidatePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	// @gotags: json:"username" bson:"username"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" bson:"username"`
	// 空间
	// @gotags: json:"namespace" bson:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 服务
	// @gotags: json:"service" bson:"service"
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service" bson:"service"`
	// 资源
	// @gotags: json:"resource" bson:"resource"
	Resource string `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource" bson:"resource"`
	// 操作行为
	// @gotags: json:"action" bson:"action"
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action" bson:"action"`
}

func (x *ValidatePermissionRequest) Reset() {
	*x = ValidatePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePermissionRequest) ProtoMessage() {}

func (x *ValidatePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePermissionRequest.ProtoReflect.Descriptor instead.
func (*ValidatePermissionRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatePermissionRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ValidatePermissionRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ValidatePermissionRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ValidatePermissionRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *ValidatePermissionRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

// 策略定义
type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色id
	// @gotags: json:"id" bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 角色创建时间
	// @gotags: json:"create_at" bson:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 策略定义
	// @gotags: json:"spec" bson:"spec"
	Spec *CreatePolicyRequest `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Policy) GetSpec() *CreatePolicyRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Policy RBAC
// 一个人可以有多个角色，从而获得对应角色的操作权限
type CreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名
	// @gotags: json:"username" bson:"username" validate:"required"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" bson:"username" validate:"required"`
	// 角色名称
	// @gotags: json:"role" bson:"role" validate:"required"
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role" bson:"role" validate:"required"`
	// 空间
	// @gotags: json:"namespace" bson:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
}

func (x *CreatePolicyRequest) Reset() {
	*x = CreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRequest) ProtoMessage() {}

func (x *CreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePolicyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreatePolicyRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CreatePolicyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type QueryPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 用户名
	// @gotags: json:"username" bson:"username"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" bson:"username"`
	// 操作空间
	// @gotags: json:"namespace" bson:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace" bson:"namespace"`
	// 角色名称
	// @gotags: json:"role" bson:"role"
	Role string `protobuf:"bytes,4,opt,name=role,proto3" json:"role" bson:"role"`
}

func (x *QueryPolicyRequest) Reset() {
	*x = QueryPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPolicyRequest) ProtoMessage() {}

func (x *QueryPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPolicyRequest.ProtoReflect.Descriptor instead.
func (*QueryPolicyRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{3}
}

func (x *QueryPolicyRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPolicyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *QueryPolicyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryPolicyRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type PolicySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 角色名称
	// @gotags: json:"id" bson:"_id"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"id" bson:"_id"`
	// 角色名称
	// @gotags: json:"items" bson:"items"
	Items []*Policy `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *PolicySet) Reset() {
	*x = PolicySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicySet) ProtoMessage() {}

func (x *PolicySet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicySet.ProtoReflect.Descriptor instead.
func (*PolicySet) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_policy_proto_rawDescGZIP(), []int{4}
}

func (x *PolicySet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PolicySet) GetItems() []*Policy {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_policy_pb_policy_proto protoreflect.FileDescriptor

var file_apps_policy_pb_policy_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x01, 0x0a, 0x19, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x71, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x3a, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x63, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x9a,
	0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x52, 0x0a, 0x09, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32,
	0xb8, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x5d, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e,
	0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6b, 0x65,
	0x79, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x52, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x25, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x67, 0x37, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6b,
	0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x74, 0x6f, 0x6e, 0x67, 0x78,
	0x75, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x67, 0x37, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_policy_pb_policy_proto_rawDescOnce sync.Once
	file_apps_policy_pb_policy_proto_rawDescData = file_apps_policy_pb_policy_proto_rawDesc
)

func file_apps_policy_pb_policy_proto_rawDescGZIP() []byte {
	file_apps_policy_pb_policy_proto_rawDescOnce.Do(func() {
		file_apps_policy_pb_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_policy_pb_policy_proto_rawDescData)
	})
	return file_apps_policy_pb_policy_proto_rawDescData
}

var file_apps_policy_pb_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_policy_pb_policy_proto_goTypes = []interface{}{
	(*ValidatePermissionRequest)(nil), // 0: keyauth_g7.policy.ValidatePermissionRequest
	(*Policy)(nil),                    // 1: keyauth_g7.policy.Policy
	(*CreatePolicyRequest)(nil),       // 2: keyauth_g7.policy.CreatePolicyRequest
	(*QueryPolicyRequest)(nil),        // 3: keyauth_g7.policy.QueryPolicyRequest
	(*PolicySet)(nil),                 // 4: keyauth_g7.policy.PolicySet
	(*request.PageRequest)(nil),       // 5: infraboard.mcube.page.PageRequest
}
var file_apps_policy_pb_policy_proto_depIdxs = []int32{
	2, // 0: keyauth_g7.policy.Policy.spec:type_name -> keyauth_g7.policy.CreatePolicyRequest
	5, // 1: keyauth_g7.policy.QueryPolicyRequest.page:type_name -> infraboard.mcube.page.PageRequest
	1, // 2: keyauth_g7.policy.PolicySet.items:type_name -> keyauth_g7.policy.Policy
	0, // 3: keyauth_g7.policy.RPC.ValidatePermission:input_type -> keyauth_g7.policy.ValidatePermissionRequest
	3, // 4: keyauth_g7.policy.RPC.QueryPolicy:input_type -> keyauth_g7.policy.QueryPolicyRequest
	1, // 5: keyauth_g7.policy.RPC.ValidatePermission:output_type -> keyauth_g7.policy.Policy
	4, // 6: keyauth_g7.policy.RPC.QueryPolicy:output_type -> keyauth_g7.policy.PolicySet
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_policy_pb_policy_proto_init() }
func file_apps_policy_pb_policy_proto_init() {
	if File_apps_policy_pb_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_policy_pb_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePermissionRequest); i {
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
		file_apps_policy_pb_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_apps_policy_pb_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyRequest); i {
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
		file_apps_policy_pb_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPolicyRequest); i {
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
		file_apps_policy_pb_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicySet); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_policy_pb_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_policy_pb_policy_proto_goTypes,
		DependencyIndexes: file_apps_policy_pb_policy_proto_depIdxs,
		MessageInfos:      file_apps_policy_pb_policy_proto_msgTypes,
	}.Build()
	File_apps_policy_pb_policy_proto = out.File
	file_apps_policy_pb_policy_proto_rawDesc = nil
	file_apps_policy_pb_policy_proto_goTypes = nil
	file_apps_policy_pb_policy_proto_depIdxs = nil
}
