// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/file.proto

package file

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

type ListRequest struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix               []string `protobuf:"bytes,2,rep,name=suffix,proto3" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{0}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ListRequest) GetSuffix() []string {
	if m != nil {
		return m.Suffix
	}
	return nil
}

type ListResponse struct {
	Files                []*FileObj `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{1}
}
func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (dst *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(dst, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetFiles() []*FileObj {
	if m != nil {
		return m.Files
	}
	return nil
}

type OpenRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	New                  bool     `protobuf:"varint,2,opt,name=new,proto3" json:"new,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenRequest) Reset()         { *m = OpenRequest{} }
func (m *OpenRequest) String() string { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()    {}
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{2}
}
func (m *OpenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenRequest.Unmarshal(m, b)
}
func (m *OpenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenRequest.Marshal(b, m, deterministic)
}
func (dst *OpenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenRequest.Merge(dst, src)
}
func (m *OpenRequest) XXX_Size() int {
	return xxx_messageInfo_OpenRequest.Size(m)
}
func (m *OpenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenRequest proto.InternalMessageInfo

func (m *OpenRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *OpenRequest) GetNew() bool {
	if m != nil {
		return m.New
	}
	return false
}

type OpenResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenResponse) Reset()         { *m = OpenResponse{} }
func (m *OpenResponse) String() string { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()    {}
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{3}
}
func (m *OpenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenResponse.Unmarshal(m, b)
}
func (m *OpenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenResponse.Marshal(b, m, deterministic)
}
func (dst *OpenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenResponse.Merge(dst, src)
}
func (m *OpenResponse) XXX_Size() int {
	return xxx_messageInfo_OpenResponse.Size(m)
}
func (m *OpenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenResponse proto.InternalMessageInfo

func (m *OpenResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OpenResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type CloseRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{4}
}
func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (dst *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(dst, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CloseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{5}
}
func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (dst *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(dst, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

type StatRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatRequest) Reset()         { *m = StatRequest{} }
func (m *StatRequest) String() string { return proto.CompactTextString(m) }
func (*StatRequest) ProtoMessage()    {}
func (*StatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{6}
}
func (m *StatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatRequest.Unmarshal(m, b)
}
func (m *StatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatRequest.Marshal(b, m, deterministic)
}
func (dst *StatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatRequest.Merge(dst, src)
}
func (m *StatRequest) XXX_Size() int {
	return xxx_messageInfo_StatRequest.Size(m)
}
func (m *StatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatRequest proto.InternalMessageInfo

func (m *StatRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type StatResponse struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	LastModified         int64    `protobuf:"varint,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatResponse) Reset()         { *m = StatResponse{} }
func (m *StatResponse) String() string { return proto.CompactTextString(m) }
func (*StatResponse) ProtoMessage()    {}
func (*StatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{7}
}
func (m *StatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatResponse.Unmarshal(m, b)
}
func (m *StatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatResponse.Marshal(b, m, deterministic)
}
func (dst *StatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatResponse.Merge(dst, src)
}
func (m *StatResponse) XXX_Size() int {
	return xxx_messageInfo_StatResponse.Size(m)
}
func (m *StatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatResponse proto.InternalMessageInfo

func (m *StatResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StatResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *StatResponse) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

type ReadRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size                 int64    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{8}
}
func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (dst *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(dst, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ReadRequest) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ReadResponse struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Eof                  bool     `protobuf:"varint,3,opt,name=eof,proto3" json:"eof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{9}
}
func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (dst *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(dst, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ReadResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReadResponse) GetEof() bool {
	if m != nil {
		return m.Eof
	}
	return false
}

type WriteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{10}
}
func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (dst *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(dst, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WriteRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *WriteRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteResponse struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{11}
}
func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (dst *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(dst, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type RemoveRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{12}
}
func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(dst, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type RemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{13}
}
func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(dst, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

type FileObj struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	LastModified         int64    `protobuf:"varint,3,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileObj) Reset()         { *m = FileObj{} }
func (m *FileObj) String() string { return proto.CompactTextString(m) }
func (*FileObj) ProtoMessage()    {}
func (*FileObj) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_025f2c97b9e19b0b, []int{14}
}
func (m *FileObj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileObj.Unmarshal(m, b)
}
func (m *FileObj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileObj.Marshal(b, m, deterministic)
}
func (dst *FileObj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileObj.Merge(dst, src)
}
func (m *FileObj) XXX_Size() int {
	return xxx_messageInfo_FileObj.Size(m)
}
func (m *FileObj) XXX_DiscardUnknown() {
	xxx_messageInfo_FileObj.DiscardUnknown(m)
}

var xxx_messageInfo_FileObj proto.InternalMessageInfo

func (m *FileObj) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FileObj) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FileObj) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "ListRequest")
	proto.RegisterType((*ListResponse)(nil), "ListResponse")
	proto.RegisterType((*OpenRequest)(nil), "OpenRequest")
	proto.RegisterType((*OpenResponse)(nil), "OpenResponse")
	proto.RegisterType((*CloseRequest)(nil), "CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "CloseResponse")
	proto.RegisterType((*StatRequest)(nil), "StatRequest")
	proto.RegisterType((*StatResponse)(nil), "StatResponse")
	proto.RegisterType((*ReadRequest)(nil), "ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "WriteResponse")
	proto.RegisterType((*RemoveRequest)(nil), "RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "RemoveResponse")
	proto.RegisterType((*FileObj)(nil), "FileObj")
}

func init() { proto.RegisterFile("proto/file.proto", fileDescriptor_file_025f2c97b9e19b0b) }

var fileDescriptor_file_025f2c97b9e19b0b = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x6b, 0x13, 0x41,
	0x10, 0x4e, 0xee, 0xd2, 0x1a, 0xe7, 0x76, 0xaf, 0x61, 0x1f, 0x24, 0xdc, 0x43, 0x09, 0x5b, 0x84,
	0x48, 0x61, 0x85, 0x0a, 0xbe, 0x88, 0x4f, 0x82, 0xa8, 0x28, 0x85, 0xf5, 0xc1, 0x07, 0x41, 0xb9,
	0x72, 0x73, 0xb0, 0x72, 0xc9, 0x9d, 0xd9, 0x8d, 0xbf, 0xfe, 0x1d, 0xff, 0x51, 0x99, 0xdd, 0x4d,
	0xb3, 0x29, 0x54, 0x4a, 0xdf, 0x66, 0xe6, 0x66, 0xbe, 0xf9, 0xb2, 0xdf, 0x37, 0x81, 0xd9, 0xb0,
	0xe9, 0x5d, 0xff, 0xb4, 0x35, 0x1d, 0x2a, 0x1f, 0xca, 0x97, 0x50, 0xbc, 0x37, 0xd6, 0x69, 0xfc,
	0xbe, 0x45, 0xeb, 0xc4, 0x23, 0x38, 0x1e, 0x36, 0xd8, 0x9a, 0x5f, 0xf3, 0xf1, 0x62, 0xbc, 0x7c,
	0xa8, 0x63, 0x46, 0x75, 0xbb, 0x6d, 0xa9, 0x9e, 0x2d, 0x72, 0xaa, 0x87, 0x4c, 0x2a, 0x60, 0x61,
	0xdc, 0x0e, 0xfd, 0xda, 0xa2, 0x38, 0x85, 0x23, 0x02, 0xb7, 0xf3, 0xf1, 0x22, 0x5f, 0x16, 0x17,
	0x53, 0xf5, 0xda, 0x74, 0x78, 0x79, 0xf5, 0x4d, 0x87, 0xb2, 0x7c, 0x01, 0xc5, 0xe5, 0x80, 0xeb,
	0xdd, 0xba, 0x0a, 0xa6, 0x54, 0x5f, 0xd7, 0x2b, 0x8c, 0x0b, 0xaf, 0x73, 0x31, 0x83, 0x7c, 0x8d,
	0x3f, 0xe7, 0xd9, 0x62, 0xbc, 0x9c, 0x6a, 0x0a, 0xe5, 0x73, 0x60, 0x61, 0x38, 0x2e, 0x2b, 0x21,
	0x33, 0x8d, 0x9f, 0xcb, 0x75, 0x66, 0x1a, 0x22, 0xb9, 0x41, 0xbb, 0xed, 0x5c, 0x1c, 0x8a, 0x99,
	0x3c, 0x05, 0xf6, 0xaa, 0xeb, 0x2d, 0xee, 0xb6, 0xde, 0x98, 0x93, 0x27, 0xc0, 0xe3, 0xf7, 0x00,
	0x2c, 0x9f, 0x40, 0xf1, 0xd1, 0xd5, 0xee, 0x0e, 0x2c, 0xe5, 0x67, 0x60, 0xa1, 0x35, 0x72, 0x12,
	0x30, 0x71, 0xbf, 0x87, 0x5d, 0x9f, 0x8f, 0xa9, 0x66, 0xcd, 0x1f, 0xf4, 0xac, 0x72, 0xed, 0x63,
	0x71, 0x06, 0xbc, 0xab, 0xad, 0xfb, 0xba, 0xea, 0x1b, 0xd3, 0x1a, 0x6c, 0xe6, 0xb9, 0xff, 0xc8,
	0xa8, 0xf8, 0x21, 0xd6, 0xe4, 0x5b, 0x28, 0x34, 0xd6, 0xcd, 0x2d, 0xbc, 0xe9, 0xf7, 0xf6, 0x6d,
	0x6b, 0xd1, 0x45, 0xe4, 0x98, 0x5d, 0xef, 0xcb, 0xf7, 0xfb, 0xe4, 0x1b, 0x60, 0x01, 0x6a, 0xcf,
	0xd3, 0xf7, 0x8c, 0x13, 0x4e, 0x02, 0x26, 0x4d, 0xed, 0x6a, 0x8f, 0xc6, 0xb4, 0x8f, 0x49, 0x05,
	0xec, 0x5b, 0x0f, 0x35, 0xd5, 0x14, 0xca, 0x77, 0xc0, 0x3e, 0x6d, 0x8c, 0xc3, 0x7b, 0xb0, 0xf2,
	0xe8, 0xf9, 0x1e, 0x5d, 0x9e, 0x01, 0x8f, 0x58, 0xb7, 0xd3, 0x92, 0xe7, 0xc0, 0x35, 0xae, 0xfa,
	0x1f, 0x78, 0x17, 0x3d, 0x66, 0x50, 0xee, 0x9a, 0xa3, 0x98, 0x5f, 0xe0, 0x41, 0x34, 0xe1, 0x7f,
	0xed, 0x76, 0x5f, 0x91, 0x2e, 0xfe, 0x66, 0x30, 0xa1, 0x05, 0xe2, 0x31, 0x4c, 0xe8, 0x16, 0x04,
	0x53, 0xc9, 0x45, 0x55, 0x5c, 0xa5, 0x07, 0x22, 0x47, 0xd4, 0x46, 0x2e, 0x16, 0x4c, 0x25, 0x97,
	0x50, 0x71, 0x95, 0x5a, 0x3b, 0xb4, 0x91, 0xb1, 0x04, 0x53, 0x89, 0x15, 0x2b, 0xae, 0x52, 0xb7,
	0x85, 0x36, 0xd2, 0x55, 0x30, 0x95, 0x38, 0xa5, 0xe2, 0x2a, 0x15, 0x5b, 0x8e, 0xc4, 0x12, 0x8e,
	0xfc, 0x43, 0x0b, 0xae, 0x52, 0xf1, 0xaa, 0x52, 0x1d, 0xbc, 0x7f, 0xe8, 0xf4, 0xc7, 0x20, 0xb8,
	0x4a, 0x8f, 0xa6, 0x2a, 0xd5, 0xe1, 0x8d, 0x8c, 0xc4, 0x39, 0x1c, 0x87, 0xa7, 0x16, 0xa5, 0x3a,
	0x10, 0xa8, 0x3a, 0x51, 0x37, 0x34, 0x18, 0x5d, 0x1d, 0xfb, 0xbf, 0x9b, 0x67, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x89, 0x4f, 0x34, 0xa3, 0x82, 0x04, 0x00, 0x00,
}
