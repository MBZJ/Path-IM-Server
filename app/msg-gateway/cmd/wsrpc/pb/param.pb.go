// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: param.proto

package pb

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqIdentifier uint32 `protobuf:"varint,1,opt,name=ReqIdentifier,proto3" json:"ReqIdentifier,omitempty"`
	Token         string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	SendID        string `protobuf:"bytes,3,opt,name=SendID,proto3" json:"SendID,omitempty"`
	MsgIncr       string `protobuf:"bytes,4,opt,name=MsgIncr,proto3" json:"MsgIncr,omitempty"`
	Data          []byte `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_param_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_param_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_param_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetReqIdentifier() uint32 {
	if x != nil {
		return x.ReqIdentifier
	}
	return 0
}

func (x *Req) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Req) GetSendID() string {
	if x != nil {
		return x.SendID
	}
	return ""
}

func (x *Req) GetMsgIncr() string {
	if x != nil {
		return x.MsgIncr
	}
	return ""
}

func (x *Req) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqIdentifier uint32 `protobuf:"varint,1,opt,name=ReqIdentifier,proto3" json:"ReqIdentifier,omitempty"`
	MsgIncr       string `protobuf:"bytes,2,opt,name=MsgIncr,proto3" json:"MsgIncr,omitempty"`
	ErrCode       uint32 `protobuf:"varint,3,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg        string `protobuf:"bytes,4,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Data          []byte `protobuf:"bytes,5,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_param_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_param_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_param_proto_rawDescGZIP(), []int{1}
}

func (x *Resp) GetReqIdentifier() uint32 {
	if x != nil {
		return x.ReqIdentifier
	}
	return 0
}

func (x *Resp) GetMsgIncr() string {
	if x != nil {
		return x.MsgIncr
	}
	return ""
}

func (x *Resp) GetErrCode() uint32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *Resp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *Resp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_param_proto protoreflect.FileDescriptor

var file_param_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d,
	0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x03, 0x52,
	0x65, 0x71, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x52, 0x65, 0x71, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x63,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x63, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x8c, 0x01, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x52, 0x65, 0x71, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x63, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x63, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x4d, 0x73,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_param_proto_rawDescOnce sync.Once
	file_param_proto_rawDescData = file_param_proto_rawDesc
)

func file_param_proto_rawDescGZIP() []byte {
	file_param_proto_rawDescOnce.Do(func() {
		file_param_proto_rawDescData = protoimpl.X.CompressGZIP(file_param_proto_rawDescData)
	})
	return file_param_proto_rawDescData
}

var file_param_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_param_proto_goTypes = []interface{}{
	(*Req)(nil),  // 0: msg_gateway.Req
	(*Resp)(nil), // 1: msg_gateway.Resp
}
var file_param_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_param_proto_init() }
func file_param_proto_init() {
	if File_param_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_param_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_param_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
			RawDescriptor: file_param_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_param_proto_goTypes,
		DependencyIndexes: file_param_proto_depIdxs,
		MessageInfos:      file_param_proto_msgTypes,
	}.Build()
	File_param_proto = out.File
	file_param_proto_rawDesc = nil
	file_param_proto_goTypes = nil
	file_param_proto_depIdxs = nil
}
