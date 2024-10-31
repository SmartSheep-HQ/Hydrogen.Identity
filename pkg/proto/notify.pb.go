// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: notify.proto

package proto

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

type NotifyInfoPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This bytes is directly encoded from pushkit.Notification
	// Which is passed to the pusher service, we don't need to care about the content
	Data    []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Unsaved bool   `protobuf:"varint,2,opt,name=unsaved,proto3" json:"unsaved,omitempty"`
}

func (x *NotifyInfoPayload) Reset() {
	*x = NotifyInfoPayload{}
	mi := &file_notify_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyInfoPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyInfoPayload) ProtoMessage() {}

func (x *NotifyInfoPayload) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyInfoPayload.ProtoReflect.Descriptor instead.
func (*NotifyInfoPayload) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyInfoPayload) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *NotifyInfoPayload) GetUnsaved() bool {
	if x != nil {
		return x.Unsaved
	}
	return false
}

type NotifyUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64             `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Notify *NotifyInfoPayload `protobuf:"bytes,2,opt,name=notify,proto3" json:"notify,omitempty"`
}

func (x *NotifyUserRequest) Reset() {
	*x = NotifyUserRequest{}
	mi := &file_notify_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyUserRequest) ProtoMessage() {}

func (x *NotifyUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyUserRequest.ProtoReflect.Descriptor instead.
func (*NotifyUserRequest) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyUserRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NotifyUserRequest) GetNotify() *NotifyInfoPayload {
	if x != nil {
		return x.Notify
	}
	return nil
}

type NotifyUserBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId []uint64           `protobuf:"varint,1,rep,packed,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Notify *NotifyInfoPayload `protobuf:"bytes,2,opt,name=notify,proto3" json:"notify,omitempty"`
}

func (x *NotifyUserBatchRequest) Reset() {
	*x = NotifyUserBatchRequest{}
	mi := &file_notify_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyUserBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyUserBatchRequest) ProtoMessage() {}

func (x *NotifyUserBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyUserBatchRequest.ProtoReflect.Descriptor instead.
func (*NotifyUserBatchRequest) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyUserBatchRequest) GetUserId() []uint64 {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *NotifyUserBatchRequest) GetNotify() *NotifyInfoPayload {
	if x != nil {
		return x.Notify
	}
	return nil
}

type NotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *NotifyResponse) Reset() {
	*x = NotifyResponse{}
	mi := &file_notify_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyResponse) ProtoMessage() {}

func (x *NotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_notify_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyResponse.ProtoReflect.Descriptor instead.
func (*NotifyResponse) Descriptor() ([]byte, []int) {
	return file_notify_proto_rawDescGZIP(), []int{3}
}

func (x *NotifyResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_notify_proto protoreflect.FileDescriptor

var file_notify_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x6e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x75, 0x6e, 0x73, 0x61, 0x76, 0x65, 0x64, 0x22, 0x5e, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x63, 0x0a, 0x16, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x06, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x2f, 0x0a,
	0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xdf,
	0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0d,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_notify_proto_rawDescOnce sync.Once
	file_notify_proto_rawDescData = file_notify_proto_rawDesc
)

func file_notify_proto_rawDescGZIP() []byte {
	file_notify_proto_rawDescOnce.Do(func() {
		file_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_notify_proto_rawDescData)
	})
	return file_notify_proto_rawDescData
}

var file_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_notify_proto_goTypes = []any{
	(*NotifyInfoPayload)(nil),      // 0: proto.NotifyInfoPayload
	(*NotifyUserRequest)(nil),      // 1: proto.NotifyUserRequest
	(*NotifyUserBatchRequest)(nil), // 2: proto.NotifyUserBatchRequest
	(*NotifyResponse)(nil),         // 3: proto.NotifyResponse
}
var file_notify_proto_depIdxs = []int32{
	0, // 0: proto.NotifyUserRequest.notify:type_name -> proto.NotifyInfoPayload
	0, // 1: proto.NotifyUserBatchRequest.notify:type_name -> proto.NotifyInfoPayload
	1, // 2: proto.NotifyService.NotifyUser:input_type -> proto.NotifyUserRequest
	2, // 3: proto.NotifyService.NotifyUserBatch:input_type -> proto.NotifyUserBatchRequest
	0, // 4: proto.NotifyService.NotifyAllUser:input_type -> proto.NotifyInfoPayload
	3, // 5: proto.NotifyService.NotifyUser:output_type -> proto.NotifyResponse
	3, // 6: proto.NotifyService.NotifyUserBatch:output_type -> proto.NotifyResponse
	3, // 7: proto.NotifyService.NotifyAllUser:output_type -> proto.NotifyResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notify_proto_init() }
func file_notify_proto_init() {
	if File_notify_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notify_proto_goTypes,
		DependencyIndexes: file_notify_proto_depIdxs,
		MessageInfos:      file_notify_proto_msgTypes,
	}.Build()
	File_notify_proto = out.File
	file_notify_proto_rawDesc = nil
	file_notify_proto_goTypes = nil
	file_notify_proto_depIdxs = nil
}
