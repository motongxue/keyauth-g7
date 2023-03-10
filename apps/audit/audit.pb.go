// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: apps/audit/pb/audit.proto

package audit

import (
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

// 操作日志
type OperateLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 谁
	// @gotags: json:"username" bson:"username"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" bson:"username"`
	// 什么时候
	// @gotags: json:"when" bson:"when"
	When int64 `protobuf:"varint,2,opt,name=when,proto3" json:"when" bson:"when"`
	// 对哪个服务
	// @gotags: json:"service" bson:"service"
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service" bson:"service"`
	// 对哪种资源
	// @gotags: json:"resource" bson:"resource"
	Resource string `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource" bson:"resource"`
	// 对哪种资源
	// @gotags: json:"action" bson:"action"
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action" bson:"action"`
	// 请求参数, URL部分
	// @gotags: json:"url" bson:"url"
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url" bson:"url"`
	// 请求参数, body部分
	// @gotags: json:"request" bson:"request"
	Request string `protobuf:"bytes,7,opt,name=request,proto3" json:"request" bson:"request"`
	// 请求响应
	// @gotags: json:"response" bson:"response"
	Response string `protobuf:"bytes,8,opt,name=response,proto3" json:"response" bson:"response"`
	// 调用时间
	// @gotags: json:"cost" bson:"cost"
	Cost int64 `protobuf:"varint,9,opt,name=cost,proto3" json:"cost" bson:"cost"`
	// 状态码
	// @gotags: json:"status_code" bson:"status_code"
	StatusCode int64 `protobuf:"varint,10,opt,name=status_code,json=statusCode,proto3" json:"status_code" bson:"status_code"`
	// 用户代理
	// @gotags: json:"user_agent" bson:"user_agent"
	UserAgent string `protobuf:"bytes,11,opt,name=user_agent,json=userAgent,proto3" json:"user_agent" bson:"user_agent"`
	// 用户ip
	// @gotags: json:"remote_ip" bson:"remote_ip"
	RemoteIp string `protobuf:"bytes,12,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip" bson:"remote_ip"`
	// 其他信息
	// @gotags: json:"meta" bson:"meta"
	Meta map[string]string `protobuf:"bytes,13,rep,name=meta,proto3" json:"meta" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"meta"`
}

func (x *OperateLog) Reset() {
	*x = OperateLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_audit_pb_audit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateLog) ProtoMessage() {}

func (x *OperateLog) ProtoReflect() protoreflect.Message {
	mi := &file_apps_audit_pb_audit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateLog.ProtoReflect.Descriptor instead.
func (*OperateLog) Descriptor() ([]byte, []int) {
	return file_apps_audit_pb_audit_proto_rawDescGZIP(), []int{0}
}

func (x *OperateLog) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *OperateLog) GetWhen() int64 {
	if x != nil {
		return x.When
	}
	return 0
}

func (x *OperateLog) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *OperateLog) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *OperateLog) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *OperateLog) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *OperateLog) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *OperateLog) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *OperateLog) GetCost() int64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *OperateLog) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *OperateLog) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *OperateLog) GetRemoteIp() string {
	if x != nil {
		return x.RemoteIp
	}
	return ""
}

func (x *OperateLog) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type AuditOperateLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuditOperateLogResponse) Reset() {
	*x = AuditOperateLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_audit_pb_audit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditOperateLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditOperateLogResponse) ProtoMessage() {}

func (x *AuditOperateLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_audit_pb_audit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditOperateLogResponse.ProtoReflect.Descriptor instead.
func (*AuditOperateLogResponse) Descriptor() ([]byte, []int) {
	return file_apps_audit_pb_audit_proto_rawDescGZIP(), []int{1}
}

var File_apps_audit_pb_audit_proto protoreflect.FileDescriptor

var file_apps_audit_pb_audit_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x22, 0xb8, 0x03,
	0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x5f, 0x69, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x49, 0x70, 0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a,
	0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x19, 0x0a, 0x17, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x5e, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x57, 0x0a, 0x0c, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6b, 0x65, 0x79,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x1a, 0x29, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x67, 0x37, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x6f, 0x74, 0x6f, 0x6e, 0x67, 0x78, 0x75, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2d, 0x67, 0x37, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_audit_pb_audit_proto_rawDescOnce sync.Once
	file_apps_audit_pb_audit_proto_rawDescData = file_apps_audit_pb_audit_proto_rawDesc
)

func file_apps_audit_pb_audit_proto_rawDescGZIP() []byte {
	file_apps_audit_pb_audit_proto_rawDescOnce.Do(func() {
		file_apps_audit_pb_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_audit_pb_audit_proto_rawDescData)
	})
	return file_apps_audit_pb_audit_proto_rawDescData
}

var file_apps_audit_pb_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_audit_pb_audit_proto_goTypes = []interface{}{
	(*OperateLog)(nil),              // 0: keyauth_g7.audit.OperateLog
	(*AuditOperateLogResponse)(nil), // 1: keyauth_g7.audit.AuditOperateLogResponse
	nil,                             // 2: keyauth_g7.audit.OperateLog.MetaEntry
}
var file_apps_audit_pb_audit_proto_depIdxs = []int32{
	2, // 0: keyauth_g7.audit.OperateLog.meta:type_name -> keyauth_g7.audit.OperateLog.MetaEntry
	0, // 1: keyauth_g7.audit.RPC.AuditOperate:input_type -> keyauth_g7.audit.OperateLog
	1, // 2: keyauth_g7.audit.RPC.AuditOperate:output_type -> keyauth_g7.audit.AuditOperateLogResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_audit_pb_audit_proto_init() }
func file_apps_audit_pb_audit_proto_init() {
	if File_apps_audit_pb_audit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_audit_pb_audit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateLog); i {
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
		file_apps_audit_pb_audit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditOperateLogResponse); i {
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
			RawDescriptor: file_apps_audit_pb_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_audit_pb_audit_proto_goTypes,
		DependencyIndexes: file_apps_audit_pb_audit_proto_depIdxs,
		MessageInfos:      file_apps_audit_pb_audit_proto_msgTypes,
	}.Build()
	File_apps_audit_pb_audit_proto = out.File
	file_apps_audit_pb_audit_proto_rawDesc = nil
	file_apps_audit_pb_audit_proto_goTypes = nil
	file_apps_audit_pb_audit_proto_depIdxs = nil
}
