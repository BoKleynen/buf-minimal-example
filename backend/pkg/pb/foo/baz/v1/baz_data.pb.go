// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: foo/baz/v1/baz_data.proto

package bazv1

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

type AData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QData []*AData_QData `protobuf:"bytes,1,rep,name=q_data,json=qData,proto3" json:"q_data,omitempty"`
}

func (x *AData) Reset() {
	*x = AData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_baz_v1_baz_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AData) ProtoMessage() {}

func (x *AData) ProtoReflect() protoreflect.Message {
	mi := &file_foo_baz_v1_baz_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AData.ProtoReflect.Descriptor instead.
func (*AData) Descriptor() ([]byte, []int) {
	return file_foo_baz_v1_baz_data_proto_rawDescGZIP(), []int{0}
}

func (x *AData) GetQData() []*AData_QData {
	if x != nil {
		return x.QData
	}
	return nil
}

type CCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CCount) Reset() {
	*x = CCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_baz_v1_baz_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCount) ProtoMessage() {}

func (x *CCount) ProtoReflect() protoreflect.Message {
	mi := &file_foo_baz_v1_baz_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCount.ProtoReflect.Descriptor instead.
func (*CCount) Descriptor() ([]byte, []int) {
	return file_foo_baz_v1_baz_data_proto_rawDescGZIP(), []int{1}
}

func (x *CCount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CCount) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CCount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CCounts []*CCount `protobuf:"bytes,1,rep,name=c_counts,json=cCounts,proto3" json:"c_counts,omitempty"`
}

func (x *BData) Reset() {
	*x = BData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_baz_v1_baz_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BData) ProtoMessage() {}

func (x *BData) ProtoReflect() protoreflect.Message {
	mi := &file_foo_baz_v1_baz_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BData.ProtoReflect.Descriptor instead.
func (*BData) Descriptor() ([]byte, []int) {
	return file_foo_baz_v1_baz_data_proto_rawDescGZIP(), []int{2}
}

func (x *BData) GetCCounts() []*CCount {
	if x != nil {
		return x.CCounts
	}
	return nil
}

type StabilityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium
	// doloremque laudantium, totam rem aperiam, eaque ipsa :
	//   - quae ab illo inventore veritatis et quasi architecto beatae vitae.
	//   - dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit
	//     aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui.
	Filters []string `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *StabilityData) Reset() {
	*x = StabilityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_baz_v1_baz_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StabilityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StabilityData) ProtoMessage() {}

func (x *StabilityData) ProtoReflect() protoreflect.Message {
	mi := &file_foo_baz_v1_baz_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StabilityData.ProtoReflect.Descriptor instead.
func (*StabilityData) Descriptor() ([]byte, []int) {
	return file_foo_baz_v1_baz_data_proto_rawDescGZIP(), []int{3}
}

func (x *StabilityData) GetFilters() []string {
	if x != nil {
		return x.Filters
	}
	return nil
}

type AData_QData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the something.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// count of something
	BCount uint64 `protobuf:"varint,2,opt,name=b_count,json=bCount,proto3" json:"b_count,omitempty"`
}

func (x *AData_QData) Reset() {
	*x = AData_QData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_baz_v1_baz_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AData_QData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AData_QData) ProtoMessage() {}

func (x *AData_QData) ProtoReflect() protoreflect.Message {
	mi := &file_foo_baz_v1_baz_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AData_QData.ProtoReflect.Descriptor instead.
func (*AData_QData) Descriptor() ([]byte, []int) {
	return file_foo_baz_v1_baz_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AData_QData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AData_QData) GetBCount() uint64 {
	if x != nil {
		return x.BCount
	}
	return 0
}

var File_foo_baz_v1_baz_data_proto protoreflect.FileDescriptor

var file_foo_baz_v1_baz_data_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x6f, 0x6f, 0x2f, 0x62, 0x61, 0x7a, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x7a,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x6f, 0x6f,
	0x2e, 0x62, 0x61, 0x7a, 0x2e, 0x76, 0x31, 0x22, 0x6d, 0x0a, 0x05, 0x41, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x06, 0x71, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x62, 0x61, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x51, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x71, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x34, 0x0a, 0x05, 0x51, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x62, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x62, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x06, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x05, 0x42, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x62, 0x61, 0x7a, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0x91, 0x01,
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x62, 0x61, 0x7a, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x42, 0x61, 0x7a, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x27, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6f,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x6f, 0x6f, 0x2f, 0x62, 0x61, 0x7a,
	0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x7a, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x46, 0x42, 0x58, 0xaa,
	0x02, 0x0a, 0x46, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x7a, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x46,
	0x6f, 0x6f, 0x5c, 0x42, 0x61, 0x7a, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x46, 0x6f, 0x6f, 0x5c,
	0x42, 0x61, 0x7a, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0c, 0x46, 0x6f, 0x6f, 0x3a, 0x3a, 0x42, 0x61, 0x7a, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foo_baz_v1_baz_data_proto_rawDescOnce sync.Once
	file_foo_baz_v1_baz_data_proto_rawDescData = file_foo_baz_v1_baz_data_proto_rawDesc
)

func file_foo_baz_v1_baz_data_proto_rawDescGZIP() []byte {
	file_foo_baz_v1_baz_data_proto_rawDescOnce.Do(func() {
		file_foo_baz_v1_baz_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_baz_v1_baz_data_proto_rawDescData)
	})
	return file_foo_baz_v1_baz_data_proto_rawDescData
}

var file_foo_baz_v1_baz_data_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_foo_baz_v1_baz_data_proto_goTypes = []interface{}{
	(*AData)(nil),         // 0: foo.baz.v1.AData
	(*CCount)(nil),        // 1: foo.baz.v1.CCount
	(*BData)(nil),         // 2: foo.baz.v1.BData
	(*StabilityData)(nil), // 3: foo.baz.v1.StabilityData
	(*AData_QData)(nil),   // 4: foo.baz.v1.AData.QData
}
var file_foo_baz_v1_baz_data_proto_depIdxs = []int32{
	4, // 0: foo.baz.v1.AData.q_data:type_name -> foo.baz.v1.AData.QData
	1, // 1: foo.baz.v1.BData.c_counts:type_name -> foo.baz.v1.CCount
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foo_baz_v1_baz_data_proto_init() }
func file_foo_baz_v1_baz_data_proto_init() {
	if File_foo_baz_v1_baz_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foo_baz_v1_baz_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AData); i {
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
		file_foo_baz_v1_baz_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCount); i {
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
		file_foo_baz_v1_baz_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BData); i {
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
		file_foo_baz_v1_baz_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StabilityData); i {
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
		file_foo_baz_v1_baz_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AData_QData); i {
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
			RawDescriptor: file_foo_baz_v1_baz_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foo_baz_v1_baz_data_proto_goTypes,
		DependencyIndexes: file_foo_baz_v1_baz_data_proto_depIdxs,
		MessageInfos:      file_foo_baz_v1_baz_data_proto_msgTypes,
	}.Build()
	File_foo_baz_v1_baz_data_proto = out.File
	file_foo_baz_v1_baz_data_proto_rawDesc = nil
	file_foo_baz_v1_baz_data_proto_goTypes = nil
	file_foo_baz_v1_baz_data_proto_depIdxs = nil
}
