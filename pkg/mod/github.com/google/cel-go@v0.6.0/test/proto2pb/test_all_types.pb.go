// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/proto2pb/test_all_types.proto

package google_expr_proto2_test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type GlobalEnum int32

const (
	GlobalEnum_GOO GlobalEnum = 0
	GlobalEnum_GAR GlobalEnum = 1
	GlobalEnum_GAZ GlobalEnum = 2
)

var GlobalEnum_name = map[int32]string{
	0: "GOO",
	1: "GAR",
	2: "GAZ",
}

var GlobalEnum_value = map[string]int32{
	"GOO": 0,
	"GAR": 1,
	"GAZ": 2,
}

func (x GlobalEnum) Enum() *GlobalEnum {
	p := new(GlobalEnum)
	*p = x
	return p
}

func (x GlobalEnum) String() string {
	return proto.EnumName(GlobalEnum_name, int32(x))
}

func (x *GlobalEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GlobalEnum_value, data, "GlobalEnum")
	if err != nil {
		return err
	}
	*x = GlobalEnum(value)
	return nil
}

func (GlobalEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a87355aedb05f01f, []int{0}
}

type TestAllTypes_NestedEnum int32

const (
	TestAllTypes_FOO TestAllTypes_NestedEnum = 0
	TestAllTypes_BAR TestAllTypes_NestedEnum = 1
	TestAllTypes_BAZ TestAllTypes_NestedEnum = 2
)

var TestAllTypes_NestedEnum_name = map[int32]string{
	0: "FOO",
	1: "BAR",
	2: "BAZ",
}

var TestAllTypes_NestedEnum_value = map[string]int32{
	"FOO": 0,
	"BAR": 1,
	"BAZ": 2,
}

func (x TestAllTypes_NestedEnum) Enum() *TestAllTypes_NestedEnum {
	p := new(TestAllTypes_NestedEnum)
	*p = x
	return p
}

func (x TestAllTypes_NestedEnum) String() string {
	return proto.EnumName(TestAllTypes_NestedEnum_name, int32(x))
}

func (x *TestAllTypes_NestedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TestAllTypes_NestedEnum_value, data, "TestAllTypes_NestedEnum")
	if err != nil {
		return err
	}
	*x = TestAllTypes_NestedEnum(value)
	return nil
}

func (TestAllTypes_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a87355aedb05f01f, []int{0, 0}
}

type TestAllTypes struct {
	SingleInt32         *int32                   `protobuf:"varint,1,opt,name=single_int32,json=singleInt32,def=-32" json:"single_int32,omitempty"`
	SingleInt64         *int64                   `protobuf:"varint,2,opt,name=single_int64,json=singleInt64,def=-64" json:"single_int64,omitempty"`
	SingleUint32        *uint32                  `protobuf:"varint,3,opt,name=single_uint32,json=singleUint32,def=32" json:"single_uint32,omitempty"`
	SingleUint64        *uint64                  `protobuf:"varint,4,opt,name=single_uint64,json=singleUint64,def=64" json:"single_uint64,omitempty"`
	SingleSint32        *int32                   `protobuf:"zigzag32,5,opt,name=single_sint32,json=singleSint32" json:"single_sint32,omitempty"`
	SingleSint64        *int64                   `protobuf:"zigzag64,6,opt,name=single_sint64,json=singleSint64" json:"single_sint64,omitempty"`
	SingleFixed32       *uint32                  `protobuf:"fixed32,7,opt,name=single_fixed32,json=singleFixed32" json:"single_fixed32,omitempty"`
	SingleFixed64       *uint64                  `protobuf:"fixed64,8,opt,name=single_fixed64,json=singleFixed64" json:"single_fixed64,omitempty"`
	SingleSfixed32      *int32                   `protobuf:"fixed32,9,opt,name=single_sfixed32,json=singleSfixed32" json:"single_sfixed32,omitempty"`
	SingleSfixed64      *int64                   `protobuf:"fixed64,10,opt,name=single_sfixed64,json=singleSfixed64" json:"single_sfixed64,omitempty"`
	SingleFloat         *float32                 `protobuf:"fixed32,11,opt,name=single_float,json=singleFloat,def=3" json:"single_float,omitempty"`
	SingleDouble        *float64                 `protobuf:"fixed64,12,opt,name=single_double,json=singleDouble,def=6.4" json:"single_double,omitempty"`
	SingleBool          *bool                    `protobuf:"varint,13,opt,name=single_bool,json=singleBool,def=1" json:"single_bool,omitempty"`
	SingleString        *string                  `protobuf:"bytes,14,opt,name=single_string,json=singleString,def=empty" json:"single_string,omitempty"`
	SingleBytes         []byte                   `protobuf:"bytes,15,opt,name=single_bytes,json=singleBytes,def=none" json:"single_bytes,omitempty"`
	StandaloneEnum      *TestAllTypes_NestedEnum `protobuf:"varint,22,opt,name=standalone_enum,json=standaloneEnum,enum=google.expr.proto2.test.TestAllTypes_NestedEnum" json:"standalone_enum,omitempty"`
	SingleAny           *any.Any                 `protobuf:"bytes,100,opt,name=single_any,json=singleAny" json:"single_any,omitempty"`
	SingleDuration      *duration.Duration       `protobuf:"bytes,101,opt,name=single_duration,json=singleDuration" json:"single_duration,omitempty"`
	SingleTimestamp     *timestamp.Timestamp     `protobuf:"bytes,102,opt,name=single_timestamp,json=singleTimestamp" json:"single_timestamp,omitempty"`
	SingleStruct        *_struct.Struct          `protobuf:"bytes,103,opt,name=single_struct,json=singleStruct" json:"single_struct,omitempty"`
	SingleValue         *_struct.Value           `protobuf:"bytes,104,opt,name=single_value,json=singleValue" json:"single_value,omitempty"`
	SingleInt64Wrapper  *wrappers.Int64Value     `protobuf:"bytes,105,opt,name=single_int64_wrapper,json=singleInt64Wrapper" json:"single_int64_wrapper,omitempty"`
	SingleInt32Wrapper  *wrappers.Int32Value     `protobuf:"bytes,106,opt,name=single_int32_wrapper,json=singleInt32Wrapper" json:"single_int32_wrapper,omitempty"`
	SingleDoubleWrapper *wrappers.DoubleValue    `protobuf:"bytes,107,opt,name=single_double_wrapper,json=singleDoubleWrapper" json:"single_double_wrapper,omitempty"`
	SingleFloatWrapper  *wrappers.FloatValue     `protobuf:"bytes,108,opt,name=single_float_wrapper,json=singleFloatWrapper" json:"single_float_wrapper,omitempty"`
	SingleUint64Wrapper *wrappers.UInt64Value    `protobuf:"bytes,109,opt,name=single_uint64_wrapper,json=singleUint64Wrapper" json:"single_uint64_wrapper,omitempty"`
	SingleUint32Wrapper *wrappers.UInt32Value    `protobuf:"bytes,110,opt,name=single_uint32_wrapper,json=singleUint32Wrapper" json:"single_uint32_wrapper,omitempty"`
	SingleStringWrapper *wrappers.StringValue    `protobuf:"bytes,111,opt,name=single_string_wrapper,json=singleStringWrapper" json:"single_string_wrapper,omitempty"`
	SingleBoolWrapper   *wrappers.BoolValue      `protobuf:"bytes,112,opt,name=single_bool_wrapper,json=singleBoolWrapper" json:"single_bool_wrapper,omitempty"`
	SingleBytesWrapper  *wrappers.BytesValue     `protobuf:"bytes,113,opt,name=single_bytes_wrapper,json=singleBytesWrapper" json:"single_bytes_wrapper,omitempty"`
	// Types that are valid to be assigned to NestedType:
	//	*TestAllTypes_SingleNestedMessage
	//	*TestAllTypes_SingleNestedEnum
	NestedType            isTestAllTypes_NestedType     `protobuf_oneof:"nested_type"`
	RepeatedInt32         []int32                       `protobuf:"varint,31,rep,name=repeated_int32,json=repeatedInt32" json:"repeated_int32,omitempty"`
	RepeatedInt64         []int64                       `protobuf:"varint,32,rep,name=repeated_int64,json=repeatedInt64" json:"repeated_int64,omitempty"`
	RepeatedUint32        []uint32                      `protobuf:"varint,33,rep,name=repeated_uint32,json=repeatedUint32" json:"repeated_uint32,omitempty"`
	RepeatedUint64        []uint64                      `protobuf:"varint,34,rep,name=repeated_uint64,json=repeatedUint64" json:"repeated_uint64,omitempty"`
	RepeatedSint32        []int32                       `protobuf:"zigzag32,35,rep,name=repeated_sint32,json=repeatedSint32" json:"repeated_sint32,omitempty"`
	RepeatedSint64        []int64                       `protobuf:"zigzag64,36,rep,name=repeated_sint64,json=repeatedSint64" json:"repeated_sint64,omitempty"`
	RepeatedFixed32       []uint32                      `protobuf:"fixed32,37,rep,name=repeated_fixed32,json=repeatedFixed32" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64       []uint64                      `protobuf:"fixed64,38,rep,name=repeated_fixed64,json=repeatedFixed64" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32      []int32                       `protobuf:"fixed32,39,rep,name=repeated_sfixed32,json=repeatedSfixed32" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64      []int64                       `protobuf:"fixed64,40,rep,name=repeated_sfixed64,json=repeatedSfixed64" json:"repeated_sfixed64,omitempty"`
	RepeatedFloat         []float32                     `protobuf:"fixed32,41,rep,name=repeated_float,json=repeatedFloat" json:"repeated_float,omitempty"`
	RepeatedDouble        []float64                     `protobuf:"fixed64,42,rep,name=repeated_double,json=repeatedDouble" json:"repeated_double,omitempty"`
	RepeatedBool          []bool                        `protobuf:"varint,43,rep,name=repeated_bool,json=repeatedBool" json:"repeated_bool,omitempty"`
	RepeatedString        []string                      `protobuf:"bytes,44,rep,name=repeated_string,json=repeatedString" json:"repeated_string,omitempty"`
	RepeatedBytes         [][]byte                      `protobuf:"bytes,45,rep,name=repeated_bytes,json=repeatedBytes" json:"repeated_bytes,omitempty"`
	RepeatedNestedMessage []*TestAllTypes_NestedMessage `protobuf:"bytes,48,rep,name=repeated_nested_message,json=repeatedNestedMessage" json:"repeated_nested_message,omitempty"`
	RepeatedNestedEnum    []TestAllTypes_NestedEnum     `protobuf:"varint,51,rep,name=repeated_nested_enum,json=repeatedNestedEnum,enum=google.expr.proto2.test.TestAllTypes_NestedEnum" json:"repeated_nested_enum,omitempty"`
	RepeatedStringPiece   []string                      `protobuf:"bytes,54,rep,name=repeated_string_piece,json=repeatedStringPiece" json:"repeated_string_piece,omitempty"`
	RepeatedCord          []string                      `protobuf:"bytes,55,rep,name=repeated_cord,json=repeatedCord" json:"repeated_cord,omitempty"`
	RepeatedLazyMessage   []*TestAllTypes_NestedMessage `protobuf:"bytes,57,rep,name=repeated_lazy_message,json=repeatedLazyMessage" json:"repeated_lazy_message,omitempty"`
	MapStringString       map[string]string             `protobuf:"bytes,58,rep,name=map_string_string,json=mapStringString" json:"map_string_string,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MapInt64NestedType    map[int64]*NestedTestAllTypes `protobuf:"bytes,59,rep,name=map_int64_nested_type,json=mapInt64NestedType" json:"map_int64_nested_type,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral  struct{}                      `json:"-"`
	XXX_unrecognized      []byte                        `json:"-"`
	XXX_sizecache         int32                         `json:"-"`
}

func (m *TestAllTypes) Reset()         { *m = TestAllTypes{} }
func (m *TestAllTypes) String() string { return proto.CompactTextString(m) }
func (*TestAllTypes) ProtoMessage()    {}
func (*TestAllTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87355aedb05f01f, []int{0}
}

func (m *TestAllTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAllTypes.Unmarshal(m, b)
}
func (m *TestAllTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAllTypes.Marshal(b, m, deterministic)
}
func (m *TestAllTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAllTypes.Merge(m, src)
}
func (m *TestAllTypes) XXX_Size() int {
	return xxx_messageInfo_TestAllTypes.Size(m)
}
func (m *TestAllTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAllTypes.DiscardUnknown(m)
}

var xxx_messageInfo_TestAllTypes proto.InternalMessageInfo

const Default_TestAllTypes_SingleInt32 int32 = -32
const Default_TestAllTypes_SingleInt64 int64 = -64
const Default_TestAllTypes_SingleUint32 uint32 = 32
const Default_TestAllTypes_SingleUint64 uint64 = 64
const Default_TestAllTypes_SingleFloat float32 = 3
const Default_TestAllTypes_SingleDouble float64 = 6.4
const Default_TestAllTypes_SingleBool bool = true
const Default_TestAllTypes_SingleString string = "empty"

var Default_TestAllTypes_SingleBytes []byte = []byte("none")

const Default_TestAllTypes_SingleNestedEnum TestAllTypes_NestedEnum = TestAllTypes_BAR

func (m *TestAllTypes) GetSingleInt32() int32 {
	if m != nil && m.SingleInt32 != nil {
		return *m.SingleInt32
	}
	return Default_TestAllTypes_SingleInt32
}

func (m *TestAllTypes) GetSingleInt64() int64 {
	if m != nil && m.SingleInt64 != nil {
		return *m.SingleInt64
	}
	return Default_TestAllTypes_SingleInt64
}

func (m *TestAllTypes) GetSingleUint32() uint32 {
	if m != nil && m.SingleUint32 != nil {
		return *m.SingleUint32
	}
	return Default_TestAllTypes_SingleUint32
}

func (m *TestAllTypes) GetSingleUint64() uint64 {
	if m != nil && m.SingleUint64 != nil {
		return *m.SingleUint64
	}
	return Default_TestAllTypes_SingleUint64
}

func (m *TestAllTypes) GetSingleSint32() int32 {
	if m != nil && m.SingleSint32 != nil {
		return *m.SingleSint32
	}
	return 0
}

func (m *TestAllTypes) GetSingleSint64() int64 {
	if m != nil && m.SingleSint64 != nil {
		return *m.SingleSint64
	}
	return 0
}

func (m *TestAllTypes) GetSingleFixed32() uint32 {
	if m != nil && m.SingleFixed32 != nil {
		return *m.SingleFixed32
	}
	return 0
}

func (m *TestAllTypes) GetSingleFixed64() uint64 {
	if m != nil && m.SingleFixed64 != nil {
		return *m.SingleFixed64
	}
	return 0
}

func (m *TestAllTypes) GetSingleSfixed32() int32 {
	if m != nil && m.SingleSfixed32 != nil {
		return *m.SingleSfixed32
	}
	return 0
}

func (m *TestAllTypes) GetSingleSfixed64() int64 {
	if m != nil && m.SingleSfixed64 != nil {
		return *m.SingleSfixed64
	}
	return 0
}

func (m *TestAllTypes) GetSingleFloat() float32 {
	if m != nil && m.SingleFloat != nil {
		return *m.SingleFloat
	}
	return Default_TestAllTypes_SingleFloat
}

func (m *TestAllTypes) GetSingleDouble() float64 {
	if m != nil && m.SingleDouble != nil {
		return *m.SingleDouble
	}
	return Default_TestAllTypes_SingleDouble
}

func (m *TestAllTypes) GetSingleBool() bool {
	if m != nil && m.SingleBool != nil {
		return *m.SingleBool
	}
	return Default_TestAllTypes_SingleBool
}

func (m *TestAllTypes) GetSingleString() string {
	if m != nil && m.SingleString != nil {
		return *m.SingleString
	}
	return Default_TestAllTypes_SingleString
}

func (m *TestAllTypes) GetSingleBytes() []byte {
	if m != nil && m.SingleBytes != nil {
		return m.SingleBytes
	}
	return append([]byte(nil), Default_TestAllTypes_SingleBytes...)
}

func (m *TestAllTypes) GetStandaloneEnum() TestAllTypes_NestedEnum {
	if m != nil && m.StandaloneEnum != nil {
		return *m.StandaloneEnum
	}
	return TestAllTypes_FOO
}

func (m *TestAllTypes) GetSingleAny() *any.Any {
	if m != nil {
		return m.SingleAny
	}
	return nil
}

func (m *TestAllTypes) GetSingleDuration() *duration.Duration {
	if m != nil {
		return m.SingleDuration
	}
	return nil
}

func (m *TestAllTypes) GetSingleTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.SingleTimestamp
	}
	return nil
}

func (m *TestAllTypes) GetSingleStruct() *_struct.Struct {
	if m != nil {
		return m.SingleStruct
	}
	return nil
}

func (m *TestAllTypes) GetSingleValue() *_struct.Value {
	if m != nil {
		return m.SingleValue
	}
	return nil
}

func (m *TestAllTypes) GetSingleInt64Wrapper() *wrappers.Int64Value {
	if m != nil {
		return m.SingleInt64Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleInt32Wrapper() *wrappers.Int32Value {
	if m != nil {
		return m.SingleInt32Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleDoubleWrapper() *wrappers.DoubleValue {
	if m != nil {
		return m.SingleDoubleWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleFloatWrapper() *wrappers.FloatValue {
	if m != nil {
		return m.SingleFloatWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleUint64Wrapper() *wrappers.UInt64Value {
	if m != nil {
		return m.SingleUint64Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleUint32Wrapper() *wrappers.UInt32Value {
	if m != nil {
		return m.SingleUint32Wrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleStringWrapper() *wrappers.StringValue {
	if m != nil {
		return m.SingleStringWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleBoolWrapper() *wrappers.BoolValue {
	if m != nil {
		return m.SingleBoolWrapper
	}
	return nil
}

func (m *TestAllTypes) GetSingleBytesWrapper() *wrappers.BytesValue {
	if m != nil {
		return m.SingleBytesWrapper
	}
	return nil
}

type isTestAllTypes_NestedType interface {
	isTestAllTypes_NestedType()
}

type TestAllTypes_SingleNestedMessage struct {
	SingleNestedMessage *TestAllTypes_NestedMessage `protobuf:"bytes,18,opt,name=single_nested_message,json=singleNestedMessage,oneof"`
}

type TestAllTypes_SingleNestedEnum struct {
	SingleNestedEnum TestAllTypes_NestedEnum `protobuf:"varint,21,opt,name=single_nested_enum,json=singleNestedEnum,enum=google.expr.proto2.test.TestAllTypes_NestedEnum,oneof,def=1"`
}

func (*TestAllTypes_SingleNestedMessage) isTestAllTypes_NestedType() {}

func (*TestAllTypes_SingleNestedEnum) isTestAllTypes_NestedType() {}

func (m *TestAllTypes) GetNestedType() isTestAllTypes_NestedType {
	if m != nil {
		return m.NestedType
	}
	return nil
}

func (m *TestAllTypes) GetSingleNestedMessage() *TestAllTypes_NestedMessage {
	if x, ok := m.GetNestedType().(*TestAllTypes_SingleNestedMessage); ok {
		return x.SingleNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetSingleNestedEnum() TestAllTypes_NestedEnum {
	if x, ok := m.GetNestedType().(*TestAllTypes_SingleNestedEnum); ok {
		return x.SingleNestedEnum
	}
	return Default_TestAllTypes_SingleNestedEnum
}

func (m *TestAllTypes) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedInt64() []int64 {
	if m != nil {
		return m.RepeatedInt64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint32() []uint32 {
	if m != nil {
		return m.RepeatedUint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedUint64() []uint64 {
	if m != nil {
		return m.RepeatedUint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint32() []int32 {
	if m != nil {
		return m.RepeatedSint32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSint64() []int64 {
	if m != nil {
		return m.RepeatedSint64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed32() []uint32 {
	if m != nil {
		return m.RepeatedFixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFixed64() []uint64 {
	if m != nil {
		return m.RepeatedFixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed32() []int32 {
	if m != nil {
		return m.RepeatedSfixed32
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedSfixed64() []int64 {
	if m != nil {
		return m.RepeatedSfixed64
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedFloat() []float32 {
	if m != nil {
		return m.RepeatedFloat
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedDouble() []float64 {
	if m != nil {
		return m.RepeatedDouble
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBool() []bool {
	if m != nil {
		return m.RepeatedBool
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedString() []string {
	if m != nil {
		return m.RepeatedString
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedBytes() [][]byte {
	if m != nil {
		return m.RepeatedBytes
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedMessage() []*TestAllTypes_NestedMessage {
	if m != nil {
		return m.RepeatedNestedMessage
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedNestedEnum() []TestAllTypes_NestedEnum {
	if m != nil {
		return m.RepeatedNestedEnum
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedStringPiece() []string {
	if m != nil {
		return m.RepeatedStringPiece
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedCord() []string {
	if m != nil {
		return m.RepeatedCord
	}
	return nil
}

func (m *TestAllTypes) GetRepeatedLazyMessage() []*TestAllTypes_NestedMessage {
	if m != nil {
		return m.RepeatedLazyMessage
	}
	return nil
}

func (m *TestAllTypes) GetMapStringString() map[string]string {
	if m != nil {
		return m.MapStringString
	}
	return nil
}

func (m *TestAllTypes) GetMapInt64NestedType() map[int64]*NestedTestAllTypes {
	if m != nil {
		return m.MapInt64NestedType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestAllTypes) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestAllTypes_SingleNestedMessage)(nil),
		(*TestAllTypes_SingleNestedEnum)(nil),
	}
}

type TestAllTypes_NestedMessage struct {
	Bb                   *int32   `protobuf:"varint,1,opt,name=bb" json:"bb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestAllTypes_NestedMessage) Reset()         { *m = TestAllTypes_NestedMessage{} }
func (m *TestAllTypes_NestedMessage) String() string { return proto.CompactTextString(m) }
func (*TestAllTypes_NestedMessage) ProtoMessage()    {}
func (*TestAllTypes_NestedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87355aedb05f01f, []int{0, 0}
}

func (m *TestAllTypes_NestedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Unmarshal(m, b)
}
func (m *TestAllTypes_NestedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Marshal(b, m, deterministic)
}
func (m *TestAllTypes_NestedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAllTypes_NestedMessage.Merge(m, src)
}
func (m *TestAllTypes_NestedMessage) XXX_Size() int {
	return xxx_messageInfo_TestAllTypes_NestedMessage.Size(m)
}
func (m *TestAllTypes_NestedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAllTypes_NestedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestAllTypes_NestedMessage proto.InternalMessageInfo

func (m *TestAllTypes_NestedMessage) GetBb() int32 {
	if m != nil && m.Bb != nil {
		return *m.Bb
	}
	return 0
}

type NestedTestAllTypes struct {
	Child                *NestedTestAllTypes `protobuf:"bytes,1,opt,name=child" json:"child,omitempty"`
	Payload              *TestAllTypes       `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NestedTestAllTypes) Reset()         { *m = NestedTestAllTypes{} }
func (m *NestedTestAllTypes) String() string { return proto.CompactTextString(m) }
func (*NestedTestAllTypes) ProtoMessage()    {}
func (*NestedTestAllTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a87355aedb05f01f, []int{1}
}

func (m *NestedTestAllTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NestedTestAllTypes.Unmarshal(m, b)
}
func (m *NestedTestAllTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NestedTestAllTypes.Marshal(b, m, deterministic)
}
func (m *NestedTestAllTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NestedTestAllTypes.Merge(m, src)
}
func (m *NestedTestAllTypes) XXX_Size() int {
	return xxx_messageInfo_NestedTestAllTypes.Size(m)
}
func (m *NestedTestAllTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_NestedTestAllTypes.DiscardUnknown(m)
}

var xxx_messageInfo_NestedTestAllTypes proto.InternalMessageInfo

func (m *NestedTestAllTypes) GetChild() *NestedTestAllTypes {
	if m != nil {
		return m.Child
	}
	return nil
}

func (m *NestedTestAllTypes) GetPayload() *TestAllTypes {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.expr.proto2.test.GlobalEnum", GlobalEnum_name, GlobalEnum_value)
	proto.RegisterEnum("google.expr.proto2.test.TestAllTypes_NestedEnum", TestAllTypes_NestedEnum_name, TestAllTypes_NestedEnum_value)
	proto.RegisterType((*TestAllTypes)(nil), "google.expr.proto2.test.TestAllTypes")
	proto.RegisterMapType((map[int64]*NestedTestAllTypes)(nil), "google.expr.proto2.test.TestAllTypes.MapInt64NestedTypeEntry")
	proto.RegisterMapType((map[string]string)(nil), "google.expr.proto2.test.TestAllTypes.MapStringStringEntry")
	proto.RegisterType((*TestAllTypes_NestedMessage)(nil), "google.expr.proto2.test.TestAllTypes.NestedMessage")
	proto.RegisterType((*NestedTestAllTypes)(nil), "google.expr.proto2.test.NestedTestAllTypes")
}

func init() { proto.RegisterFile("test/proto2pb/test_all_types.proto", fileDescriptor_a87355aedb05f01f) }

var fileDescriptor_a87355aedb05f01f = []byte{
	// 1327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x6d, 0x73, 0xd3, 0xc6,
	0x16, 0x66, 0xb5, 0x09, 0x49, 0x36, 0x71, 0xe2, 0x6c, 0x12, 0xb2, 0x04, 0xe6, 0xb2, 0x37, 0x10,
	0xb2, 0xc0, 0xc5, 0x30, 0xb6, 0x46, 0xb7, 0xa8, 0xed, 0x74, 0xe2, 0x02, 0xa1, 0x9d, 0x52, 0x18,
	0x05, 0xda, 0x69, 0xbf, 0x78, 0xe4, 0x78, 0x13, 0xdc, 0xc8, 0x92, 0x2a, 0xcb, 0x2d, 0xe6, 0x3f,
	0xf4, 0x7b, 0x7f, 0x6e, 0x47, 0xfb, 0xe6, 0x95, 0x84, 0x3b, 0x90, 0x2f, 0x19, 0xeb, 0x9c, 0xe7,
	0x3c, 0xe7, 0x4d, 0xe7, 0x1c, 0x05, 0xed, 0xe7, 0x7c, 0x9c, 0x3f, 0x4a, 0xb3, 0x24, 0x4f, 0xda,
	0x69, 0xff, 0x51, 0xf1, 0xd4, 0x0b, 0xa3, 0xa8, 0x97, 0x4f, 0x53, 0x3e, 0x6e, 0x09, 0x39, 0xde,
	0x3d, 0x4f, 0x92, 0xf3, 0x88, 0xb7, 0xf8, 0xfb, 0x34, 0x93, 0xa2, 0x76, 0xab, 0x00, 0xee, 0x5d,
	0x97, 0x0a, 0x69, 0xde, 0x9f, 0x9c, 0x3d, 0x0a, 0xe3, 0xa9, 0x04, 0xec, 0xfd, 0xa7, 0xaa, 0x1a,
	0x4c, 0xb2, 0x30, 0x1f, 0x26, 0xb1, 0xd2, 0xdf, 0xac, 0xea, 0xc7, 0x79, 0x36, 0x39, 0xcd, 0x95,
	0xf6, 0x56, 0x55, 0x9b, 0x0f, 0x47, 0x7c, 0x9c, 0x87, 0xa3, 0x74, 0x1e, 0xfd, 0x9f, 0x59, 0x98,
	0xa6, 0x3c, 0x53, 0x21, 0xef, 0xff, 0x75, 0x03, 0xad, 0xbd, 0xe1, 0xe3, 0xfc, 0x28, 0x8a, 0xde,
	0x14, 0x99, 0xe0, 0xbb, 0x68, 0x6d, 0x3c, 0x8c, 0xcf, 0x23, 0xde, 0x1b, 0xc6, 0x79, 0xa7, 0x4d,
	0x00, 0x05, 0x6c, 0xd1, 0x87, 0x0f, 0x3b, 0xed, 0x60, 0x55, 0x2a, 0xbe, 0x2b, 0xe4, 0x65, 0x9c,
	0xe7, 0x12, 0x87, 0x02, 0x06, 0x7d, 0xf8, 0xd0, 0x73, 0x2d, 0x9c, 0xe7, 0xe2, 0x43, 0xd4, 0x50,
	0xb8, 0x89, 0x24, 0x84, 0x14, 0xb0, 0x86, 0xef, 0x74, 0xda, 0x81, 0x22, 0x78, 0x2b, 0xe4, 0x15,
	0xa0, 0xe7, 0x92, 0x05, 0x0a, 0xd8, 0x82, 0xef, 0x78, 0xae, 0x0d, 0xf4, 0x5c, 0x7c, 0xdb, 0x00,
	0xc7, 0x92, 0x71, 0x91, 0x02, 0xb6, 0xa9, 0x41, 0x27, 0x92, 0xad, 0x0c, 0xf2, 0x5c, 0x72, 0x95,
	0x02, 0x86, 0x6d, 0x90, 0xe7, 0xe2, 0x03, 0xb4, 0xae, 0x40, 0x67, 0xc3, 0xf7, 0x7c, 0xd0, 0x69,
	0x93, 0x25, 0x0a, 0xd8, 0x52, 0xa0, 0x4c, 0x9f, 0x4b, 0x61, 0x15, 0xe6, 0xb9, 0x64, 0x99, 0x02,
	0x76, 0xb5, 0x04, 0x13, 0x99, 0x6e, 0x68, 0x97, 0x9a, 0x6e, 0x85, 0x02, 0xb6, 0x11, 0x28, 0xeb,
	0x13, 0x25, 0xad, 0x01, 0x3d, 0x97, 0x20, 0x0a, 0x58, 0xb3, 0x0c, 0xf4, 0x5c, 0x7c, 0xc7, 0xd4,
	0xf8, 0x2c, 0x4a, 0xc2, 0x9c, 0xac, 0x52, 0xc0, 0x1c, 0x1f, 0x74, 0x74, 0x85, 0x9f, 0x17, 0x52,
	0xcc, 0x4c, 0xaa, 0x83, 0x64, 0xd2, 0x8f, 0x38, 0x59, 0xa3, 0x80, 0x01, 0x1f, 0x7a, 0x2d, 0x53,
	0xb9, 0xa7, 0x42, 0x81, 0x0f, 0x90, 0x32, 0xec, 0xf5, 0x93, 0x24, 0x22, 0x0d, 0x0a, 0xd8, 0xb2,
	0xbf, 0x90, 0x67, 0x13, 0x1e, 0x20, 0xa9, 0xe8, 0x26, 0x49, 0x84, 0xef, 0xcf, 0x6a, 0x97, 0x67,
	0xc3, 0xf8, 0x9c, 0xac, 0x53, 0xc0, 0x56, 0xfc, 0x45, 0x3e, 0x4a, 0xf3, 0xa9, 0x29, 0xa1, 0x50,
	0xe1, 0x43, 0x13, 0x62, 0x7f, 0x9a, 0xf3, 0x31, 0xd9, 0xa0, 0x80, 0xad, 0xf9, 0x0b, 0x71, 0x12,
	0x73, 0x1d, 0x65, 0xb7, 0x50, 0xe0, 0x5f, 0xd0, 0xc6, 0x38, 0x0f, 0xe3, 0x41, 0x18, 0x25, 0x31,
	0xef, 0xf1, 0x78, 0x32, 0x22, 0xd7, 0x28, 0x60, 0xeb, 0xed, 0xc7, 0xad, 0x39, 0x53, 0xd3, 0xb2,
	0xdf, 0xcb, 0xd6, 0x8f, 0x7c, 0x9c, 0xf3, 0xc1, 0xb3, 0x78, 0x32, 0x0a, 0xd6, 0x67, 0x44, 0xc5,
	0x33, 0xee, 0x20, 0x15, 0x7d, 0x2f, 0x8c, 0xa7, 0x64, 0x40, 0x01, 0x5b, 0x6d, 0x6f, 0x6b, 0x56,
	0xfd, 0xe2, 0xb7, 0x8e, 0xe2, 0x69, 0xb0, 0x22, 0x71, 0x47, 0xf1, 0x14, 0x77, 0x4d, 0x13, 0xf4,
	0xc0, 0x11, 0x2e, 0x2c, 0xaf, 0xd7, 0x2c, 0x9f, 0x2a, 0x80, 0xee, 0x8f, 0x7e, 0xc6, 0xcf, 0x50,
	0x53, 0x71, 0x98, 0xb1, 0x23, 0x67, 0x82, 0x64, 0xaf, 0x46, 0xf2, 0x46, 0x23, 0x02, 0xe5, 0xd7,
	0x08, 0xf0, 0x57, 0x76, 0xbd, 0x27, 0xa7, 0x39, 0x39, 0x17, 0x1c, 0xbb, 0x35, 0x8e, 0x13, 0xa1,
	0xb6, 0x3a, 0x30, 0x39, 0xcd, 0xf1, 0x13, 0xd3, 0x81, 0x3f, 0xc2, 0x68, 0xc2, 0xc9, 0x3b, 0x61,
	0x7c, 0xad, 0x66, 0xfc, 0x53, 0xa1, 0xd5, 0x3d, 0x11, 0x0f, 0xf8, 0x25, 0xda, 0xb6, 0x67, 0xb8,
	0xa7, 0x76, 0x03, 0x19, 0x0a, 0x8a, 0x1b, 0x35, 0x0a, 0x31, 0xd1, 0x92, 0x07, 0x5b, 0x33, 0xfe,
	0xb3, 0x34, 0x2b, 0xd3, 0x75, 0xda, 0x86, 0xee, 0xb7, 0xf9, 0x74, 0x9d, 0x76, 0x95, 0xae, 0xd3,
	0xd6, 0x74, 0xaf, 0xd1, 0x4e, 0xe9, 0xbd, 0x36, 0x7c, 0x17, 0x82, 0xef, 0x66, 0xbd, 0x4f, 0x02,
	0x26, 0x09, 0xb7, 0xec, 0x17, 0xbf, 0x1e, 0xa0, 0x98, 0x27, 0x43, 0x18, 0xcd, 0x09, 0x50, 0xcc,
	0x57, 0x29, 0x40, 0x21, 0xa9, 0x07, 0x38, 0x29, 0xd7, 0x6f, 0x34, 0x27, 0xc0, 0xb7, 0x56, 0x01,
	0xb7, 0xec, 0x9d, 0xf6, 0x71, 0x46, 0xab, 0x84, 0xf1, 0xbf, 0x30, 0xea, 0x1a, 0x6e, 0xd9, 0xeb,
	0xb4, 0xce, 0x28, 0x67, 0xd9, 0x30, 0x26, 0x73, 0x18, 0xe5, 0x5c, 0x97, 0x18, 0xa5, 0x48, 0x33,
	0x7e, 0x8f, 0xb6, 0xac, 0x25, 0x62, 0xf8, 0xd2, 0x39, 0xef, 0x7d, 0xb1, 0x51, 0x24, 0xdb, 0xe6,
	0x6c, 0xc5, 0xd4, 0x1b, 0x22, 0xb6, 0x87, 0x21, 0xfb, 0x7d, 0x4e, 0x43, 0xc4, 0x2a, 0x29, 0x35,
	0x44, 0x48, 0x34, 0xdd, 0xd0, 0x24, 0x1b, 0x8b, 0x6d, 0xd1, 0x1b, 0xf1, 0xf1, 0x38, 0x3c, 0xe7,
	0x04, 0x0b, 0xbe, 0xce, 0xe7, 0x6c, 0x9a, 0x97, 0xd2, 0xf4, 0xc5, 0x15, 0x5d, 0x85, 0x92, 0x18,
	0x73, 0x84, 0xcb, 0xae, 0xc4, 0x46, 0xdb, 0xb9, 0xdc, 0x46, 0xf3, 0x61, 0xf7, 0x28, 0x78, 0x71,
	0x25, 0x68, 0xda, 0x9e, 0xc4, 0x6a, 0x3b, 0x40, 0xeb, 0x19, 0x4f, 0x79, 0x58, 0x78, 0x90, 0xc7,
	0xee, 0x16, 0x85, 0x6c, 0x31, 0x68, 0x68, 0xa9, 0x3c, 0xc6, 0x15, 0x98, 0xe7, 0x12, 0x4a, 0x21,
	0x83, 0x25, 0x98, 0xbc, 0x50, 0x06, 0xa6, 0xae, 0xf1, 0x7f, 0x29, 0x64, 0x8d, 0xc0, 0x58, 0x9b,
	0x5b, 0x5c, 0x06, 0x7a, 0x2e, 0xd9, 0xa7, 0x90, 0x2d, 0x94, 0x81, 0x15, 0x46, 0x75, 0x8d, 0x6f,
	0x53, 0xc8, 0x36, 0x67, 0xc0, 0x93, 0x3a, 0xa3, 0xba, 0xc8, 0x77, 0x28, 0x64, 0xb8, 0x0c, 0xf4,
	0x5c, 0x7c, 0x0f, 0x35, 0x0d, 0x50, 0x9f, 0xd1, 0x03, 0x0a, 0xd9, 0x52, 0x60, 0x08, 0xf4, 0x5d,
	0xae, 0x41, 0x3d, 0x97, 0xdc, 0xa5, 0x90, 0x5d, 0xad, 0x40, 0x3d, 0x17, 0x3f, 0x40, 0x9b, 0x33,
	0xf7, 0x9a, 0xf6, 0x90, 0x42, 0xb6, 0x11, 0x18, 0x0e, 0x73, 0x9f, 0xeb, 0x60, 0xcf, 0x25, 0x8c,
	0x42, 0xd6, 0xac, 0x82, 0xe5, 0x37, 0xc4, 0x2c, 0x08, 0x71, 0xa5, 0xef, 0x51, 0xc8, 0x9c, 0x59,
	0xe9, 0xe5, 0x91, 0xb6, 0xf3, 0x57, 0x67, 0xfa, 0x3e, 0x85, 0x0c, 0xcc, 0xf2, 0x57, 0x37, 0xfa,
	0x36, 0x32, 0x96, 0xf2, 0x4a, 0x3f, 0xa0, 0x90, 0x2d, 0x07, 0x6b, 0x5a, 0x28, 0x2e, 0x74, 0xa9,
	0x9a, 0xf2, 0x46, 0xff, 0x8f, 0x42, 0xb6, 0x62, 0x55, 0x53, 0x9e, 0x67, 0x3b, 0x3a, 0x79, 0xa0,
	0x1f, 0x52, 0xc8, 0xd6, 0x66, 0xd1, 0xc9, 0xe3, 0x7c, 0x81, 0x76, 0x0d, 0xac, 0x32, 0x3a, 0x8f,
	0x29, 0xbc, 0xe4, 0xe8, 0x04, 0x3b, 0x9a, 0xb3, 0x3c, 0x3a, 0x7d, 0xb4, 0x5d, 0x75, 0x26, 0x86,
	0xa7, 0x43, 0xe1, 0xa5, 0x3e, 0x07, 0x70, 0xd9, 0x8d, 0x98, 0x1b, 0x0f, 0xed, 0x54, 0x0a, 0xd4,
	0x4b, 0x87, 0xfc, 0x94, 0x13, 0xaf, 0x28, 0x53, 0xd7, 0x59, 0x76, 0x82, 0xad, 0x72, 0xa9, 0x5e,
	0x17, 0xea, 0xe2, 0x23, 0xd4, 0xd8, 0x9d, 0x26, 0xd9, 0x80, 0xfc, 0x5f, 0xe1, 0xc1, 0xac, 0x03,
	0xdf, 0x26, 0xd9, 0x00, 0x5f, 0x58, 0x0e, 0xa2, 0xf0, 0xc3, 0xd4, 0xd4, 0xeb, 0xc9, 0xa5, 0xeb,
	0xd5, 0x75, 0x18, 0x98, 0x45, 0xf5, 0x43, 0xf8, 0x61, 0xaa, 0x2b, 0x76, 0x86, 0x36, 0x47, 0x61,
	0xaa, 0x13, 0x51, 0x0d, 0xf7, 0x85, 0x23, 0xff, 0xd3, 0x1c, 0xbd, 0x0c, 0x53, 0x99, 0xa6, 0xfc,
	0xfb, 0x2c, 0xce, 0xb3, 0x69, 0xb0, 0x31, 0x2a, 0x4b, 0x71, 0x8a, 0x76, 0x0a, 0x3f, 0xf2, 0x98,
	0xa9, 0xd6, 0x14, 0xff, 0xdf, 0x90, 0x2f, 0x85, 0xaf, 0xaf, 0x3f, 0xd9, 0x97, 0x58, 0x37, 0x32,
	0xb9, 0x42, 0x26, 0xdd, 0xe1, 0x51, 0x4d, 0xb1, 0x77, 0x0b, 0x35, 0xca, 0x2f, 0xc7, 0x3a, 0x72,
	0xfa, 0x7d, 0xf9, 0x4f, 0x47, 0xe0, 0xf4, 0xfb, 0x7b, 0x5d, 0xb4, 0xfd, 0xb1, 0xd8, 0x71, 0x13,
	0xc1, 0x0b, 0x3e, 0x15, 0xc0, 0x95, 0xa0, 0xf8, 0x89, 0xb7, 0xd1, 0xa2, 0xfc, 0x00, 0x72, 0x84,
	0x4c, 0x3e, 0xf8, 0xce, 0x17, 0x60, 0x2f, 0x43, 0xbb, 0x73, 0x62, 0xb2, 0x69, 0xa0, 0xa4, 0x39,
	0xb2, 0x69, 0x56, 0xdb, 0x0f, 0xe6, 0xe6, 0xac, 0xa8, 0xac, 0xcc, 0x2d, 0x9f, 0xfb, 0x87, 0x08,
	0x59, 0xaf, 0xe3, 0x12, 0x82, 0xcf, 0x5f, 0xbd, 0x6a, 0x5e, 0x29, 0x7e, 0x74, 0x8f, 0x82, 0x26,
	0x90, 0x3f, 0x7e, 0x6d, 0x3a, 0xdd, 0x06, 0x5a, 0xb5, 0x2a, 0xbd, 0xff, 0x37, 0x40, 0xb8, 0xce,
	0x5c, 0x44, 0x75, 0xfa, 0x6e, 0x18, 0x0d, 0x44, 0xa4, 0x9f, 0x1b, 0x95, 0xb0, 0xc4, 0xdf, 0xa0,
	0xa5, 0x34, 0x9c, 0x46, 0x49, 0x38, 0x50, 0xa9, 0x1d, 0x7c, 0x52, 0x3b, 0x03, 0x6d, 0x75, 0xff,
	0x10, 0xa1, 0xe3, 0x28, 0xe9, 0x87, 0x91, 0x4e, 0xe9, 0x58, 0xa7, 0x74, 0xac, 0x53, 0x3a, 0x2e,
	0x52, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x40, 0x40, 0x1a, 0xe9, 0x2b, 0x0f, 0x00, 0x00,
}