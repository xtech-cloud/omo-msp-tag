// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/tag/dummy.proto

package tag

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 添加标签的请求
type DummyAddTagRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyAddTagRequest) Reset()         { *m = DummyAddTagRequest{} }
func (m *DummyAddTagRequest) String() string { return proto.CompactTextString(m) }
func (*DummyAddTagRequest) ProtoMessage()    {}
func (*DummyAddTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ab4a3ff9059ec40, []int{0}
}

func (m *DummyAddTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyAddTagRequest.Unmarshal(m, b)
}
func (m *DummyAddTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyAddTagRequest.Marshal(b, m, deterministic)
}
func (m *DummyAddTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyAddTagRequest.Merge(m, src)
}
func (m *DummyAddTagRequest) XXX_Size() int {
	return xxx_messageInfo_DummyAddTagRequest.Size(m)
}
func (m *DummyAddTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyAddTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DummyAddTagRequest proto.InternalMessageInfo

func (m *DummyAddTagRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DummyAddTagRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// 删除标签的请求
type DummyRemoveTagRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyRemoveTagRequest) Reset()         { *m = DummyRemoveTagRequest{} }
func (m *DummyRemoveTagRequest) String() string { return proto.CompactTextString(m) }
func (*DummyRemoveTagRequest) ProtoMessage()    {}
func (*DummyRemoveTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ab4a3ff9059ec40, []int{1}
}

func (m *DummyRemoveTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyRemoveTagRequest.Unmarshal(m, b)
}
func (m *DummyRemoveTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyRemoveTagRequest.Marshal(b, m, deterministic)
}
func (m *DummyRemoveTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyRemoveTagRequest.Merge(m, src)
}
func (m *DummyRemoveTagRequest) XXX_Size() int {
	return xxx_messageInfo_DummyRemoveTagRequest.Size(m)
}
func (m *DummyRemoveTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyRemoveTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DummyRemoveTagRequest proto.InternalMessageInfo

func (m *DummyRemoveTagRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DummyRemoveTagRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// 检索标签的请求
type DummyFilterTagRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Code                 []string `protobuf:"bytes,3,rep,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyFilterTagRequest) Reset()         { *m = DummyFilterTagRequest{} }
func (m *DummyFilterTagRequest) String() string { return proto.CompactTextString(m) }
func (*DummyFilterTagRequest) ProtoMessage()    {}
func (*DummyFilterTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ab4a3ff9059ec40, []int{2}
}

func (m *DummyFilterTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyFilterTagRequest.Unmarshal(m, b)
}
func (m *DummyFilterTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyFilterTagRequest.Marshal(b, m, deterministic)
}
func (m *DummyFilterTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyFilterTagRequest.Merge(m, src)
}
func (m *DummyFilterTagRequest) XXX_Size() int {
	return xxx_messageInfo_DummyFilterTagRequest.Size(m)
}
func (m *DummyFilterTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyFilterTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DummyFilterTagRequest proto.InternalMessageInfo

func (m *DummyFilterTagRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DummyFilterTagRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *DummyFilterTagRequest) GetCode() []string {
	if m != nil {
		return m.Code
	}
	return nil
}

// 检索标签的回复
type DummyFilterTagResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Owner                []string `protobuf:"bytes,3,rep,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyFilterTagResponse) Reset()         { *m = DummyFilterTagResponse{} }
func (m *DummyFilterTagResponse) String() string { return proto.CompactTextString(m) }
func (*DummyFilterTagResponse) ProtoMessage()    {}
func (*DummyFilterTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ab4a3ff9059ec40, []int{3}
}

func (m *DummyFilterTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyFilterTagResponse.Unmarshal(m, b)
}
func (m *DummyFilterTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyFilterTagResponse.Marshal(b, m, deterministic)
}
func (m *DummyFilterTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyFilterTagResponse.Merge(m, src)
}
func (m *DummyFilterTagResponse) XXX_Size() int {
	return xxx_messageInfo_DummyFilterTagResponse.Size(m)
}
func (m *DummyFilterTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyFilterTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DummyFilterTagResponse proto.InternalMessageInfo

func (m *DummyFilterTagResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DummyFilterTagResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *DummyFilterTagResponse) GetOwner() []string {
	if m != nil {
		return m.Owner
	}
	return nil
}

// 列举标签的请求
type DummyListTagRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyListTagRequest) Reset()         { *m = DummyListTagRequest{} }
func (m *DummyListTagRequest) String() string { return proto.CompactTextString(m) }
func (*DummyListTagRequest) ProtoMessage()    {}
func (*DummyListTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ab4a3ff9059ec40, []int{4}
}

func (m *DummyListTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyListTagRequest.Unmarshal(m, b)
}
func (m *DummyListTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyListTagRequest.Marshal(b, m, deterministic)
}
func (m *DummyListTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyListTagRequest.Merge(m, src)
}
func (m *DummyListTagRequest) XXX_Size() int {
	return xxx_messageInfo_DummyListTagRequest.Size(m)
}
func (m *DummyListTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyListTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DummyListTagRequest proto.InternalMessageInfo

func (m *DummyListTagRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DummyListTagRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *DummyListTagRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// 列举标签的回复
type DummyListTagResponse struct {
	Status               *Status      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                int64        `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Tag                  []*TagEntity `protobuf:"bytes,3,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DummyListTagResponse) Reset()         { *m = DummyListTagResponse{} }
func (m *DummyListTagResponse) String() string { return proto.CompactTextString(m) }
func (*DummyListTagResponse) ProtoMessage()    {}
func (*DummyListTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ab4a3ff9059ec40, []int{5}
}

func (m *DummyListTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyListTagResponse.Unmarshal(m, b)
}
func (m *DummyListTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyListTagResponse.Marshal(b, m, deterministic)
}
func (m *DummyListTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyListTagResponse.Merge(m, src)
}
func (m *DummyListTagResponse) XXX_Size() int {
	return xxx_messageInfo_DummyListTagResponse.Size(m)
}
func (m *DummyListTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyListTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DummyListTagResponse proto.InternalMessageInfo

func (m *DummyListTagResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DummyListTagResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *DummyListTagResponse) GetTag() []*TagEntity {
	if m != nil {
		return m.Tag
	}
	return nil
}

func init() {
	proto.RegisterType((*DummyAddTagRequest)(nil), "tag.DummyAddTagRequest")
	proto.RegisterType((*DummyRemoveTagRequest)(nil), "tag.DummyRemoveTagRequest")
	proto.RegisterType((*DummyFilterTagRequest)(nil), "tag.DummyFilterTagRequest")
	proto.RegisterType((*DummyFilterTagResponse)(nil), "tag.DummyFilterTagResponse")
	proto.RegisterType((*DummyListTagRequest)(nil), "tag.DummyListTagRequest")
	proto.RegisterType((*DummyListTagResponse)(nil), "tag.DummyListTagResponse")
}

func init() {
	proto.RegisterFile("proto/tag/dummy.proto", fileDescriptor_0ab4a3ff9059ec40)
}

var fileDescriptor_0ab4a3ff9059ec40 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x3f, 0x53, 0xc2, 0x40,
	0x10, 0xc5, 0x85, 0x48, 0x1c, 0x96, 0x19, 0x8b, 0x15, 0x10, 0x63, 0xc3, 0xc4, 0xc6, 0x0a, 0x66,
	0xb0, 0xb0, 0xd2, 0x11, 0x47, 0xa9, 0xac, 0x4e, 0x1a, 0xcb, 0x93, 0x1c, 0x31, 0x23, 0xe4, 0x30,
	0xb7, 0xd1, 0xe1, 0x23, 0xf8, 0xad, 0x9d, 0x6c, 0xfe, 0x82, 0xd8, 0x60, 0xc7, 0xde, 0xdb, 0xfd,
	0x3d, 0xf6, 0xee, 0x05, 0x3a, 0xab, 0x48, 0x93, 0x1e, 0x92, 0xf4, 0x87, 0x5e, 0xbc, 0x5c, 0xae,
	0x07, 0x5c, 0xa3, 0x45, 0xd2, 0x77, 0xba, 0xa5, 0x66, 0xde, 0x64, 0xa4, 0xbc, 0x54, 0x74, 0x6f,
	0x01, 0x1f, 0x92, 0xde, 0xb1, 0xe7, 0x4d, 0xa5, 0x2f, 0xd4, 0x47, 0xac, 0x0c, 0x21, 0xc2, 0xe1,
	0x4c, 0x7b, 0xaa, 0x57, 0xeb, 0xd7, 0x2e, 0x9b, 0x82, 0x7f, 0x63, 0x1b, 0x1a, 0xfa, 0x2b, 0x54,
	0x51, 0xaf, 0xce, 0x87, 0x69, 0xe1, 0x8e, 0xa1, 0xc3, 0xf3, 0x42, 0x2d, 0xf5, 0xa7, 0xda, 0x0b,
	0xf1, 0x92, 0x21, 0x26, 0xc1, 0x82, 0x54, 0x54, 0x41, 0x74, 0xc1, 0xd6, 0xf3, 0xb9, 0x51, 0xc4,
	0x10, 0x4b, 0x64, 0x55, 0x82, 0x99, 0xe9, 0x38, 0x24, 0xc6, 0x58, 0x22, 0x2d, 0x0a, 0x43, 0xab,
	0x6f, 0xe5, 0x86, 0x6e, 0x00, 0xdd, 0x6d, 0xb4, 0x59, 0xe9, 0xd0, 0x28, 0xbc, 0x00, 0xdb, 0x90,
	0xa4, 0xd8, 0x30, 0xbb, 0x35, 0x6a, 0x0d, 0x48, 0xfa, 0x83, 0x67, 0x3e, 0x12, 0x99, 0x94, 0x18,
	0x91, 0x26, 0xb9, 0xc8, 0x8d, 0xb8, 0x28, 0xb7, 0x48, 0x9d, 0x8a, 0x2d, 0x4e, 0xd8, 0xea, 0x29,
	0x30, 0xb4, 0xf7, 0x0e, 0x15, 0x74, 0xe5, 0x82, 0x0c, 0xb4, 0x37, 0xd1, 0xff, 0xdf, 0xa1, 0x0f,
	0x49, 0x2a, 0x78, 0x83, 0xd6, 0xe8, 0x98, 0xe7, 0xa6, 0xd2, 0x7f, 0x0c, 0x29, 0xa0, 0xb5, 0x48,
	0xa4, 0xd1, 0x77, 0x1d, 0x1a, 0xec, 0x8a, 0xd7, 0x60, 0xa7, 0xe9, 0xc0, 0x53, 0x6e, 0xfc, 0x9d,
	0x17, 0x07, 0x59, 0xb8, 0x5f, 0xc8, 0xf0, 0x3d, 0xff, 0x77, 0xee, 0x01, 0xde, 0x40, 0xb3, 0x88,
	0x05, 0x3a, 0xe5, 0xec, 0x76, 0x56, 0xfe, 0x18, 0x9f, 0x40, 0xb3, 0x78, 0xb7, 0xea, 0xf8, 0x76,
	0x4e, 0x9c, 0xf3, 0x9d, 0x5a, 0xc1, 0xb9, 0x83, 0xa3, 0xec, 0xe6, 0xb0, 0x57, 0x76, 0x6e, 0xbe,
	0x93, 0x73, 0xb6, 0x43, 0xc9, 0x09, 0xaf, 0x36, 0x7f, 0x2b, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0xb1, 0xde, 0x76, 0x61, 0x03, 0x00, 0x00,
}
