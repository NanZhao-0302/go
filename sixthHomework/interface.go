package apis


import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
	"sync"
)


//type Circle struct {
//	ID           uint
//	Timestamp    int64
//	PersonID     uint
//	PersonName   string
//	Content      string
//	AtTimeHeight float32
//	AtTimeWeight float32
//	Visible      bool
//}

type TopPost struct {
	ID            uint32
	Timestamp     int64
	PersonID      uint32
	PersonName    string
	Content       string
	AtTimeHeight  float32
	AtTimeWeight  float32
	AtTimeFatRate float32
}

func (*Circle) TableName() string {
	return "testdb.circle"
}



const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CirCleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Circle `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CirCleList) Reset() {
	*x = CirCleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CirCleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CirCleList) ProtoMessage() {}

func (x *CirCleList) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CirCleList.ProtoReflect.Descriptor instead.
func (*CirCleList) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *CirCleList) GetItems() []*Circle {
	if x != nil {
		return x.Items
	}
	return nil
}

type Circle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Circle) Reset() {
	*x = Circle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Circle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Circle) ProtoMessage() {}

func (x *Circle) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Circle.ProtoReflect.Descriptor instead.
func (*Circle) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Circle) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Circle) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Circle) GetPersonId() uint32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *Circle) GetPersonName() string {
	if x != nil {
		return x.PersonName
	}
	return ""
}

func (x *Circle) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *Circle) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Circle) GetAtTimeHeight() float32 {
	if x != nil {
		return x.AtTimeHeight
	}
	return 0
}

func (x *Circle) GetAtTimeWeight() float32 {
	if x != nil {
		return x.AtTimeWeight
	}
	return 0
}

func (x *Circle) GetAtTimeAge() uint32 {
	if x != nil {
		return x.AtTimeAge
	}
	return 0
}

func (x *Circle) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

type ServerInterface interface {
	PostStatus(c *apis.Circle) error
	DeletePost(id uint32) error
	ListPost() ([]*apis.TopPost, error)
}

type CircleInitInterface interface {
	Init() error
}

var File_types_proto protoreflect.FileDescriptor


var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_proto_goTypes = []interface{}{
	(*CirCleList)(nil),
	(*Circle)(nil),
}
var file_types_proto_depIdxs = []int32{
	1,
	1,
	1,
	1,
	1,
	0,
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CirCleList); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Circle); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
