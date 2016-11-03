// Code generated by protoc-gen-go.
// source: transicator.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	transicator.proto

It has these top-level messages:
	ValuePb
	ColumnPb
	ChangePb
	ChangeListPb
	SnapshotHeaderPb
	TableHeaderPb
	RowPb
	StreamMessagePb
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ValuePb struct {
	// Types that are valid to be assigned to Value:
	//	*ValuePb_String_
	//	*ValuePb_Int
	//	*ValuePb_Uint
	//	*ValuePb_Double
	//	*ValuePb_Bytes
	//	*ValuePb_Bool
	Value            isValuePb_Value `protobuf_oneof:"value"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ValuePb) Reset()                    { *m = ValuePb{} }
func (m *ValuePb) String() string            { return proto.CompactTextString(m) }
func (*ValuePb) ProtoMessage()               {}
func (*ValuePb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isValuePb_Value interface {
	isValuePb_Value()
}

type ValuePb_String_ struct {
	String_ string `protobuf:"bytes,1,opt,name=string,oneof"`
}
type ValuePb_Int struct {
	Int int64 `protobuf:"varint,2,opt,name=int,oneof"`
}
type ValuePb_Uint struct {
	Uint uint64 `protobuf:"varint,3,opt,name=uint,oneof"`
}
type ValuePb_Double struct {
	Double float64 `protobuf:"fixed64,4,opt,name=double,oneof"`
}
type ValuePb_Bytes struct {
	Bytes []byte `protobuf:"bytes,5,opt,name=bytes,oneof"`
}
type ValuePb_Bool struct {
	Bool bool `protobuf:"varint,6,opt,name=bool,oneof"`
}

func (*ValuePb_String_) isValuePb_Value() {}
func (*ValuePb_Int) isValuePb_Value()     {}
func (*ValuePb_Uint) isValuePb_Value()    {}
func (*ValuePb_Double) isValuePb_Value()  {}
func (*ValuePb_Bytes) isValuePb_Value()   {}
func (*ValuePb_Bool) isValuePb_Value()    {}

func (m *ValuePb) GetValue() isValuePb_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValuePb) GetString_() string {
	if x, ok := m.GetValue().(*ValuePb_String_); ok {
		return x.String_
	}
	return ""
}

func (m *ValuePb) GetInt() int64 {
	if x, ok := m.GetValue().(*ValuePb_Int); ok {
		return x.Int
	}
	return 0
}

func (m *ValuePb) GetUint() uint64 {
	if x, ok := m.GetValue().(*ValuePb_Uint); ok {
		return x.Uint
	}
	return 0
}

func (m *ValuePb) GetDouble() float64 {
	if x, ok := m.GetValue().(*ValuePb_Double); ok {
		return x.Double
	}
	return 0
}

func (m *ValuePb) GetBytes() []byte {
	if x, ok := m.GetValue().(*ValuePb_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (m *ValuePb) GetBool() bool {
	if x, ok := m.GetValue().(*ValuePb_Bool); ok {
		return x.Bool
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ValuePb) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ValuePb_OneofMarshaler, _ValuePb_OneofUnmarshaler, _ValuePb_OneofSizer, []interface{}{
		(*ValuePb_String_)(nil),
		(*ValuePb_Int)(nil),
		(*ValuePb_Uint)(nil),
		(*ValuePb_Double)(nil),
		(*ValuePb_Bytes)(nil),
		(*ValuePb_Bool)(nil),
	}
}

func _ValuePb_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ValuePb)
	// value
	switch x := m.Value.(type) {
	case *ValuePb_String_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.String_)
	case *ValuePb_Int:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int))
	case *ValuePb_Uint:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint))
	case *ValuePb_Double:
		b.EncodeVarint(4<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Double))
	case *ValuePb_Bytes:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Bytes)
	case *ValuePb_Bool:
		t := uint64(0)
		if x.Bool {
			t = 1
		}
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("ValuePb.Value has unexpected type %T", x)
	}
	return nil
}

func _ValuePb_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ValuePb)
	switch tag {
	case 1: // value.string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &ValuePb_String_{x}
		return true, err
	case 2: // value.int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ValuePb_Int{int64(x)}
		return true, err
	case 3: // value.uint
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ValuePb_Uint{x}
		return true, err
	case 4: // value.double
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &ValuePb_Double{math.Float64frombits(x)}
		return true, err
	case 5: // value.bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &ValuePb_Bytes{x}
		return true, err
	case 6: // value.bool
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ValuePb_Bool{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _ValuePb_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ValuePb)
	// value
	switch x := m.Value.(type) {
	case *ValuePb_String_:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.String_)))
		n += len(x.String_)
	case *ValuePb_Int:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int))
	case *ValuePb_Uint:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint))
	case *ValuePb_Double:
		n += proto.SizeVarint(4<<3 | proto.WireFixed64)
		n += 8
	case *ValuePb_Bytes:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Bytes)))
		n += len(x.Bytes)
	case *ValuePb_Bool:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ColumnPb struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *ValuePb `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Type             *int32   `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ColumnPb) Reset()                    { *m = ColumnPb{} }
func (m *ColumnPb) String() string            { return proto.CompactTextString(m) }
func (*ColumnPb) ProtoMessage()               {}
func (*ColumnPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ColumnPb) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ColumnPb) GetValue() *ValuePb {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ColumnPb) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type ChangePb struct {
	Operation        *int32      `protobuf:"varint,1,req,name=operation" json:"operation,omitempty"`
	Table            *string     `protobuf:"bytes,2,req,name=table" json:"table,omitempty"`
	Sequence         *string     `protobuf:"bytes,3,opt,name=sequence" json:"sequence,omitempty"`
	CommitSequence   *uint64     `protobuf:"varint,4,opt,name=commitSequence" json:"commitSequence,omitempty"`
	ChangeSequence   *uint64     `protobuf:"varint,5,opt,name=changeSequence" json:"changeSequence,omitempty"`
	CommitIndex      *uint32     `protobuf:"varint,6,opt,name=commitIndex" json:"commitIndex,omitempty"`
	TransactionID    *uint32     `protobuf:"varint,7,opt,name=transactionID" json:"transactionID,omitempty"`
	NewColumns       []*ColumnPb `protobuf:"bytes,8,rep,name=newColumns" json:"newColumns,omitempty"`
	OldColumns       []*ColumnPb `protobuf:"bytes,9,rep,name=oldColumns" json:"oldColumns,omitempty"`
	Timestamp        *int64      `protobuf:"varint,10,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ChangePb) Reset()                    { *m = ChangePb{} }
func (m *ChangePb) String() string            { return proto.CompactTextString(m) }
func (*ChangePb) ProtoMessage()               {}
func (*ChangePb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChangePb) GetOperation() int32 {
	if m != nil && m.Operation != nil {
		return *m.Operation
	}
	return 0
}

func (m *ChangePb) GetTable() string {
	if m != nil && m.Table != nil {
		return *m.Table
	}
	return ""
}

func (m *ChangePb) GetSequence() string {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return ""
}

func (m *ChangePb) GetCommitSequence() uint64 {
	if m != nil && m.CommitSequence != nil {
		return *m.CommitSequence
	}
	return 0
}

func (m *ChangePb) GetChangeSequence() uint64 {
	if m != nil && m.ChangeSequence != nil {
		return *m.ChangeSequence
	}
	return 0
}

func (m *ChangePb) GetCommitIndex() uint32 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func (m *ChangePb) GetTransactionID() uint32 {
	if m != nil && m.TransactionID != nil {
		return *m.TransactionID
	}
	return 0
}

func (m *ChangePb) GetNewColumns() []*ColumnPb {
	if m != nil {
		return m.NewColumns
	}
	return nil
}

func (m *ChangePb) GetOldColumns() []*ColumnPb {
	if m != nil {
		return m.OldColumns
	}
	return nil
}

func (m *ChangePb) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type ChangeListPb struct {
	LastSequence     *string     `protobuf:"bytes,1,opt,name=lastSequence" json:"lastSequence,omitempty"`
	FirstSequence    *string     `protobuf:"bytes,2,opt,name=firstSequence" json:"firstSequence,omitempty"`
	Changes          []*ChangePb `protobuf:"bytes,3,rep,name=changes" json:"changes,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ChangeListPb) Reset()                    { *m = ChangeListPb{} }
func (m *ChangeListPb) String() string            { return proto.CompactTextString(m) }
func (*ChangeListPb) ProtoMessage()               {}
func (*ChangeListPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ChangeListPb) GetLastSequence() string {
	if m != nil && m.LastSequence != nil {
		return *m.LastSequence
	}
	return ""
}

func (m *ChangeListPb) GetFirstSequence() string {
	if m != nil && m.FirstSequence != nil {
		return *m.FirstSequence
	}
	return ""
}

func (m *ChangeListPb) GetChanges() []*ChangePb {
	if m != nil {
		return m.Changes
	}
	return nil
}

type SnapshotHeaderPb struct {
	Timestamp        *string `protobuf:"bytes,1,req,name=timestamp" json:"timestamp,omitempty"`
	Snapshot         *string `protobuf:"bytes,2,req,name=snapshot" json:"snapshot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SnapshotHeaderPb) Reset()                    { *m = SnapshotHeaderPb{} }
func (m *SnapshotHeaderPb) String() string            { return proto.CompactTextString(m) }
func (*SnapshotHeaderPb) ProtoMessage()               {}
func (*SnapshotHeaderPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SnapshotHeaderPb) GetTimestamp() string {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return ""
}

func (m *SnapshotHeaderPb) GetSnapshot() string {
	if m != nil && m.Snapshot != nil {
		return *m.Snapshot
	}
	return ""
}

type TableHeaderPb struct {
	Name             *string     `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Columns          []*ColumnPb `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TableHeaderPb) Reset()                    { *m = TableHeaderPb{} }
func (m *TableHeaderPb) String() string            { return proto.CompactTextString(m) }
func (*TableHeaderPb) ProtoMessage()               {}
func (*TableHeaderPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TableHeaderPb) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TableHeaderPb) GetColumns() []*ColumnPb {
	if m != nil {
		return m.Columns
	}
	return nil
}

type RowPb struct {
	Values           []*ValuePb `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RowPb) Reset()                    { *m = RowPb{} }
func (m *RowPb) String() string            { return proto.CompactTextString(m) }
func (*RowPb) ProtoMessage()               {}
func (*RowPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RowPb) GetValues() []*ValuePb {
	if m != nil {
		return m.Values
	}
	return nil
}

type StreamMessagePb struct {
	// Types that are valid to be assigned to Message:
	//	*StreamMessagePb_Table
	//	*StreamMessagePb_Row
	Message          isStreamMessagePb_Message `protobuf_oneof:"message"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *StreamMessagePb) Reset()                    { *m = StreamMessagePb{} }
func (m *StreamMessagePb) String() string            { return proto.CompactTextString(m) }
func (*StreamMessagePb) ProtoMessage()               {}
func (*StreamMessagePb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isStreamMessagePb_Message interface {
	isStreamMessagePb_Message()
}

type StreamMessagePb_Table struct {
	Table *TableHeaderPb `protobuf:"bytes,1,opt,name=table,oneof"`
}
type StreamMessagePb_Row struct {
	Row *RowPb `protobuf:"bytes,2,opt,name=row,oneof"`
}

func (*StreamMessagePb_Table) isStreamMessagePb_Message() {}
func (*StreamMessagePb_Row) isStreamMessagePb_Message()   {}

func (m *StreamMessagePb) GetMessage() isStreamMessagePb_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *StreamMessagePb) GetTable() *TableHeaderPb {
	if x, ok := m.GetMessage().(*StreamMessagePb_Table); ok {
		return x.Table
	}
	return nil
}

func (m *StreamMessagePb) GetRow() *RowPb {
	if x, ok := m.GetMessage().(*StreamMessagePb_Row); ok {
		return x.Row
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamMessagePb) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamMessagePb_OneofMarshaler, _StreamMessagePb_OneofUnmarshaler, _StreamMessagePb_OneofSizer, []interface{}{
		(*StreamMessagePb_Table)(nil),
		(*StreamMessagePb_Row)(nil),
	}
}

func _StreamMessagePb_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamMessagePb)
	// message
	switch x := m.Message.(type) {
	case *StreamMessagePb_Table:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Table); err != nil {
			return err
		}
	case *StreamMessagePb_Row:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Row); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamMessagePb.Message has unexpected type %T", x)
	}
	return nil
}

func _StreamMessagePb_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamMessagePb)
	switch tag {
	case 1: // message.table
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TableHeaderPb)
		err := b.DecodeMessage(msg)
		m.Message = &StreamMessagePb_Table{msg}
		return true, err
	case 2: // message.row
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RowPb)
		err := b.DecodeMessage(msg)
		m.Message = &StreamMessagePb_Row{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamMessagePb_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamMessagePb)
	// message
	switch x := m.Message.(type) {
	case *StreamMessagePb_Table:
		s := proto.Size(x.Table)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamMessagePb_Row:
		s := proto.Size(x.Row)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ValuePb)(nil), "common.ValuePb")
	proto.RegisterType((*ColumnPb)(nil), "common.ColumnPb")
	proto.RegisterType((*ChangePb)(nil), "common.ChangePb")
	proto.RegisterType((*ChangeListPb)(nil), "common.ChangeListPb")
	proto.RegisterType((*SnapshotHeaderPb)(nil), "common.SnapshotHeaderPb")
	proto.RegisterType((*TableHeaderPb)(nil), "common.TableHeaderPb")
	proto.RegisterType((*RowPb)(nil), "common.RowPb")
	proto.RegisterType((*StreamMessagePb)(nil), "common.StreamMessagePb")
}

func init() { proto.RegisterFile("transicator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xbf, 0xe2, 0x78, 0x1a, 0x37, 0xa9, 0xa1, 0xc8, 0x07, 0x04, 0xc5, 0x42, 0x28, 0xa7,
	0x1c, 0xb8, 0x70, 0x45, 0x80, 0x50, 0x2a, 0x81, 0x84, 0x08, 0xe2, 0x88, 0xb4, 0x4e, 0x96, 0xd6,
	0x92, 0x77, 0xd7, 0xec, 0xae, 0x1b, 0xfa, 0xd3, 0xb9, 0x31, 0x3b, 0x6b, 0x07, 0x82, 0xca, 0x71,
	0xde, 0xbe, 0x99, 0x37, 0x6f, 0xf6, 0xc1, 0xb9, 0xd5, 0x4c, 0x9a, 0x66, 0xcb, 0xac, 0xd2, 0xab,
	0x4e, 0x2b, 0xab, 0x8a, 0xc9, 0x56, 0x09, 0xa1, 0x64, 0x75, 0x0b, 0xe9, 0x57, 0xd6, 0xf6, 0xfc,
	0x53, 0x5d, 0x2c, 0x60, 0x62, 0xac, 0x6e, 0xe4, 0x75, 0x19, 0x5c, 0x06, 0xcb, 0x6c, 0x7d, 0x52,
	0xe4, 0x10, 0x35, 0xd2, 0x96, 0x21, 0x96, 0x11, 0x96, 0x67, 0x10, 0xf7, 0xae, 0x8e, 0xb0, 0x8e,
	0xb1, 0xc6, 0x86, 0x9d, 0xea, 0xeb, 0x96, 0x97, 0x31, 0x22, 0x01, 0x22, 0x73, 0x48, 0xea, 0x3b,
	0xcb, 0x4d, 0x99, 0x20, 0x30, 0xf3, 0x2d, 0xb5, 0x52, 0x6d, 0x39, 0xc1, 0x7a, 0xba, 0x3e, 0x79,
	0x93, 0x42, 0x72, 0xeb, 0xe4, 0xaa, 0xf7, 0x30, 0x7d, 0xab, 0xda, 0x5e, 0x48, 0x14, 0x9e, 0x41,
	0x2c, 0x99, 0xe0, 0x28, 0x1b, 0x2e, 0xb3, 0xe2, 0xc9, 0x40, 0x21, 0xd9, 0xd3, 0x97, 0xf3, 0x95,
	0xdf, 0x74, 0x35, 0xae, 0x89, 0x6c, 0x7b, 0xd7, 0x71, 0xda, 0x22, 0xa9, 0x7e, 0x05, 0x38, 0xe8,
	0x86, 0xc9, 0x6b, 0xf7, 0x74, 0x0e, 0x99, 0xea, 0xb8, 0x66, 0xb6, 0x51, 0x92, 0xa6, 0x25, 0x68,
	0x21, 0xb1, 0xcc, 0xad, 0x18, 0xd2, 0xf0, 0x05, 0x4c, 0x0d, 0xff, 0xd1, 0x73, 0xb9, 0xf5, 0x03,
	0xb2, 0xe2, 0x11, 0x9c, 0x39, 0x81, 0xc6, 0x6e, 0x46, 0xdc, 0x99, 0x89, 0x09, 0xa7, 0xb9, 0x07,
	0x3c, 0x21, 0xfc, 0x01, 0x9c, 0x7a, 0xfe, 0x95, 0xdc, 0xf1, 0x9f, 0x64, 0x2c, 0x2f, 0x2e, 0x20,
	0xa7, 0x13, 0xb3, 0xad, 0x93, 0xbe, 0x7a, 0x57, 0xa6, 0x04, 0x3f, 0x07, 0x90, 0x7c, 0xef, 0x7d,
	0x9a, 0x72, 0x7a, 0x19, 0xa1, 0x9f, 0xc5, 0xe8, 0xe7, 0x60, 0x1f, 0x59, 0xaa, 0xdd, 0x8d, 0xac,
	0xec, 0x3f, 0x2c, 0xf4, 0x66, 0x1b, 0xc1, 0x8d, 0x65, 0xa2, 0x2b, 0xc1, 0xfd, 0x48, 0xf5, 0x0d,
	0x66, 0xde, 0xfa, 0x87, 0xc6, 0x58, 0xa4, 0x3c, 0x84, 0x59, 0xcb, 0xcc, 0x1f, 0x23, 0xf4, 0x8d,
	0x6e, 0xb7, 0xef, 0x8d, 0xfe, 0x0b, 0x0e, 0x09, 0x7e, 0x06, 0xa9, 0xf7, 0x67, 0xf0, 0x10, 0xc7,
	0x92, 0xc3, 0x39, 0xab, 0x57, 0xb0, 0xd8, 0x48, 0xd6, 0x99, 0x1b, 0x65, 0xd7, 0x9c, 0xed, 0xb8,
	0xfe, 0x77, 0x8d, 0xe0, 0x70, 0xd3, 0x81, 0xe6, 0xaf, 0x5c, 0xbd, 0x86, 0xfc, 0x8b, 0x3b, 0xfa,
	0xa1, 0xeb, 0xf8, 0x87, 0x9d, 0xf4, 0xe0, 0x36, 0xbc, 0xdf, 0x6d, 0xb5, 0x84, 0xe4, 0xb3, 0xda,
	0x63, 0xe7, 0x53, 0x98, 0x50, 0x1a, 0x0c, 0xf6, 0x46, 0xf7, 0xc4, 0xa1, 0xaa, 0x61, 0xbe, 0xb1,
	0x9a, 0x33, 0xf1, 0x91, 0x1b, 0xc3, 0x28, 0x06, 0x2f, 0xc6, 0x3f, 0x0f, 0x28, 0x41, 0x17, 0x63,
	0xcb, 0xd1, 0x4e, 0x18, 0xce, 0xc7, 0x10, 0x69, 0xb5, 0x1f, 0x72, 0x96, 0x8f, 0x2c, 0xd2, 0xc5,
	0xa8, 0x66, 0x90, 0x0a, 0x3f, 0xf2, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xad, 0x5c, 0x98,
	0x40, 0x03, 0x00, 0x00,
}
