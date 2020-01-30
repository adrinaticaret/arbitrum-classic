// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package rollupvalidator

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

type LogInfo struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BlockHash            string   `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber          string   `protobuf:"bytes,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	LogIndex             string   `protobuf:"bytes,5,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Topics               []string `protobuf:"bytes,6,rep,name=topics,proto3" json:"topics,omitempty"`
	TransactionIndex     string   `protobuf:"bytes,7,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	TransactionHash      string   `protobuf:"bytes,8,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInfo) Reset()         { *m = LogInfo{} }
func (m *LogInfo) String() string { return proto.CompactTextString(m) }
func (*LogInfo) ProtoMessage()    {}
func (*LogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *LogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInfo.Unmarshal(m, b)
}
func (m *LogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInfo.Marshal(b, m, deterministic)
}
func (m *LogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInfo.Merge(m, src)
}
func (m *LogInfo) XXX_Size() int {
	return xxx_messageInfo_LogInfo.Size(m)
}
func (m *LogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogInfo proto.InternalMessageInfo

func (m *LogInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *LogInfo) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *LogInfo) GetBlockNumber() string {
	if m != nil {
		return m.BlockNumber
	}
	return ""
}

func (m *LogInfo) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *LogInfo) GetLogIndex() string {
	if m != nil {
		return m.LogIndex
	}
	return ""
}

func (m *LogInfo) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *LogInfo) GetTransactionIndex() string {
	if m != nil {
		return m.TransactionIndex
	}
	return ""
}

func (m *LogInfo) GetTransactionHash() string {
	if m != nil {
		return m.TransactionHash
	}
	return ""
}

type FindLogsArgs struct {
	FromHeight           string   `protobuf:"bytes,1,opt,name=fromHeight,proto3" json:"fromHeight,omitempty"`
	ToHeight             string   `protobuf:"bytes,2,opt,name=toHeight,proto3" json:"toHeight,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Topics               []string `protobuf:"bytes,4,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindLogsArgs) Reset()         { *m = FindLogsArgs{} }
func (m *FindLogsArgs) String() string { return proto.CompactTextString(m) }
func (*FindLogsArgs) ProtoMessage()    {}
func (*FindLogsArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *FindLogsArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLogsArgs.Unmarshal(m, b)
}
func (m *FindLogsArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLogsArgs.Marshal(b, m, deterministic)
}
func (m *FindLogsArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLogsArgs.Merge(m, src)
}
func (m *FindLogsArgs) XXX_Size() int {
	return xxx_messageInfo_FindLogsArgs.Size(m)
}
func (m *FindLogsArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLogsArgs.DiscardUnknown(m)
}

var xxx_messageInfo_FindLogsArgs proto.InternalMessageInfo

func (m *FindLogsArgs) GetFromHeight() string {
	if m != nil {
		return m.FromHeight
	}
	return ""
}

func (m *FindLogsArgs) GetToHeight() string {
	if m != nil {
		return m.ToHeight
	}
	return ""
}

func (m *FindLogsArgs) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *FindLogsArgs) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type FindLogsReply struct {
	Logs                 []*LogInfo `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindLogsReply) Reset()         { *m = FindLogsReply{} }
func (m *FindLogsReply) String() string { return proto.CompactTextString(m) }
func (*FindLogsReply) ProtoMessage()    {}
func (*FindLogsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *FindLogsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindLogsReply.Unmarshal(m, b)
}
func (m *FindLogsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindLogsReply.Marshal(b, m, deterministic)
}
func (m *FindLogsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindLogsReply.Merge(m, src)
}
func (m *FindLogsReply) XXX_Size() int {
	return xxx_messageInfo_FindLogsReply.Size(m)
}
func (m *FindLogsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FindLogsReply.DiscardUnknown(m)
}

var xxx_messageInfo_FindLogsReply proto.InternalMessageInfo

func (m *FindLogsReply) GetLogs() []*LogInfo {
	if m != nil {
		return m.Logs
	}
	return nil
}

type GetMessageResultArgs struct {
	TxHash               string   `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageResultArgs) Reset()         { *m = GetMessageResultArgs{} }
func (m *GetMessageResultArgs) String() string { return proto.CompactTextString(m) }
func (*GetMessageResultArgs) ProtoMessage()    {}
func (*GetMessageResultArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *GetMessageResultArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageResultArgs.Unmarshal(m, b)
}
func (m *GetMessageResultArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageResultArgs.Marshal(b, m, deterministic)
}
func (m *GetMessageResultArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageResultArgs.Merge(m, src)
}
func (m *GetMessageResultArgs) XXX_Size() int {
	return xxx_messageInfo_GetMessageResultArgs.Size(m)
}
func (m *GetMessageResultArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageResultArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageResultArgs proto.InternalMessageInfo

func (m *GetMessageResultArgs) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

type GetMessageResultReply struct {
	Found                bool     `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	RawVal               string   `protobuf:"bytes,2,opt,name=rawVal,proto3" json:"rawVal,omitempty"`
	LogPreHash           string   `protobuf:"bytes,3,opt,name=logPreHash,proto3" json:"logPreHash,omitempty"`
	LogPostHash          string   `protobuf:"bytes,4,opt,name=logPostHash,proto3" json:"logPostHash,omitempty"`
	LogValHashes         []string `protobuf:"bytes,5,rep,name=logValHashes,proto3" json:"logValHashes,omitempty"`
	OnChainTxHash        string   `protobuf:"bytes,6,opt,name=onChainTxHash,proto3" json:"onChainTxHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageResultReply) Reset()         { *m = GetMessageResultReply{} }
func (m *GetMessageResultReply) String() string { return proto.CompactTextString(m) }
func (*GetMessageResultReply) ProtoMessage()    {}
func (*GetMessageResultReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *GetMessageResultReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageResultReply.Unmarshal(m, b)
}
func (m *GetMessageResultReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageResultReply.Marshal(b, m, deterministic)
}
func (m *GetMessageResultReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageResultReply.Merge(m, src)
}
func (m *GetMessageResultReply) XXX_Size() int {
	return xxx_messageInfo_GetMessageResultReply.Size(m)
}
func (m *GetMessageResultReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageResultReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageResultReply proto.InternalMessageInfo

func (m *GetMessageResultReply) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *GetMessageResultReply) GetRawVal() string {
	if m != nil {
		return m.RawVal
	}
	return ""
}

func (m *GetMessageResultReply) GetLogPreHash() string {
	if m != nil {
		return m.LogPreHash
	}
	return ""
}

func (m *GetMessageResultReply) GetLogPostHash() string {
	if m != nil {
		return m.LogPostHash
	}
	return ""
}

func (m *GetMessageResultReply) GetLogValHashes() []string {
	if m != nil {
		return m.LogValHashes
	}
	return nil
}

func (m *GetMessageResultReply) GetOnChainTxHash() string {
	if m != nil {
		return m.OnChainTxHash
	}
	return ""
}

type GetAssertionCountArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssertionCountArgs) Reset()         { *m = GetAssertionCountArgs{} }
func (m *GetAssertionCountArgs) String() string { return proto.CompactTextString(m) }
func (*GetAssertionCountArgs) ProtoMessage()    {}
func (*GetAssertionCountArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *GetAssertionCountArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssertionCountArgs.Unmarshal(m, b)
}
func (m *GetAssertionCountArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssertionCountArgs.Marshal(b, m, deterministic)
}
func (m *GetAssertionCountArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssertionCountArgs.Merge(m, src)
}
func (m *GetAssertionCountArgs) XXX_Size() int {
	return xxx_messageInfo_GetAssertionCountArgs.Size(m)
}
func (m *GetAssertionCountArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssertionCountArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssertionCountArgs proto.InternalMessageInfo

type GetAssertionCountReply struct {
	AssertionCount       int32    `protobuf:"varint,1,opt,name=assertionCount,proto3" json:"assertionCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssertionCountReply) Reset()         { *m = GetAssertionCountReply{} }
func (m *GetAssertionCountReply) String() string { return proto.CompactTextString(m) }
func (*GetAssertionCountReply) ProtoMessage()    {}
func (*GetAssertionCountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{6}
}

func (m *GetAssertionCountReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssertionCountReply.Unmarshal(m, b)
}
func (m *GetAssertionCountReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssertionCountReply.Marshal(b, m, deterministic)
}
func (m *GetAssertionCountReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssertionCountReply.Merge(m, src)
}
func (m *GetAssertionCountReply) XXX_Size() int {
	return xxx_messageInfo_GetAssertionCountReply.Size(m)
}
func (m *GetAssertionCountReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssertionCountReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssertionCountReply proto.InternalMessageInfo

func (m *GetAssertionCountReply) GetAssertionCount() int32 {
	if m != nil {
		return m.AssertionCount
	}
	return 0
}

type GetVMInfoArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoArgs) Reset()         { *m = GetVMInfoArgs{} }
func (m *GetVMInfoArgs) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoArgs) ProtoMessage()    {}
func (*GetVMInfoArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{7}
}

func (m *GetVMInfoArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoArgs.Unmarshal(m, b)
}
func (m *GetVMInfoArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoArgs.Marshal(b, m, deterministic)
}
func (m *GetVMInfoArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoArgs.Merge(m, src)
}
func (m *GetVMInfoArgs) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoArgs.Size(m)
}
func (m *GetVMInfoArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoArgs proto.InternalMessageInfo

type GetVMInfoReply struct {
	VmID                 string   `protobuf:"bytes,1,opt,name=vmID,proto3" json:"vmID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVMInfoReply) Reset()         { *m = GetVMInfoReply{} }
func (m *GetVMInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetVMInfoReply) ProtoMessage()    {}
func (*GetVMInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{8}
}

func (m *GetVMInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVMInfoReply.Unmarshal(m, b)
}
func (m *GetVMInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVMInfoReply.Marshal(b, m, deterministic)
}
func (m *GetVMInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVMInfoReply.Merge(m, src)
}
func (m *GetVMInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetVMInfoReply.Size(m)
}
func (m *GetVMInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVMInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetVMInfoReply proto.InternalMessageInfo

func (m *GetVMInfoReply) GetVmID() string {
	if m != nil {
		return m.VmID
	}
	return ""
}

type CallMessageArgs struct {
	ContractAddress      string   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	Sender               string   `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallMessageArgs) Reset()         { *m = CallMessageArgs{} }
func (m *CallMessageArgs) String() string { return proto.CompactTextString(m) }
func (*CallMessageArgs) ProtoMessage()    {}
func (*CallMessageArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{9}
}

func (m *CallMessageArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallMessageArgs.Unmarshal(m, b)
}
func (m *CallMessageArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallMessageArgs.Marshal(b, m, deterministic)
}
func (m *CallMessageArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallMessageArgs.Merge(m, src)
}
func (m *CallMessageArgs) XXX_Size() int {
	return xxx_messageInfo_CallMessageArgs.Size(m)
}
func (m *CallMessageArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_CallMessageArgs.DiscardUnknown(m)
}

var xxx_messageInfo_CallMessageArgs proto.InternalMessageInfo

func (m *CallMessageArgs) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *CallMessageArgs) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *CallMessageArgs) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type CallMessageReply struct {
	RawVal               string   `protobuf:"bytes,1,opt,name=rawVal,proto3" json:"rawVal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallMessageReply) Reset()         { *m = CallMessageReply{} }
func (m *CallMessageReply) String() string { return proto.CompactTextString(m) }
func (*CallMessageReply) ProtoMessage()    {}
func (*CallMessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{10}
}

func (m *CallMessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallMessageReply.Unmarshal(m, b)
}
func (m *CallMessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallMessageReply.Marshal(b, m, deterministic)
}
func (m *CallMessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallMessageReply.Merge(m, src)
}
func (m *CallMessageReply) XXX_Size() int {
	return xxx_messageInfo_CallMessageReply.Size(m)
}
func (m *CallMessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CallMessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_CallMessageReply proto.InternalMessageInfo

func (m *CallMessageReply) GetRawVal() string {
	if m != nil {
		return m.RawVal
	}
	return ""
}

func init() {
	proto.RegisterType((*LogInfo)(nil), "rollupvalidator.LogInfo")
	proto.RegisterType((*FindLogsArgs)(nil), "rollupvalidator.FindLogsArgs")
	proto.RegisterType((*FindLogsReply)(nil), "rollupvalidator.FindLogsReply")
	proto.RegisterType((*GetMessageResultArgs)(nil), "rollupvalidator.GetMessageResultArgs")
	proto.RegisterType((*GetMessageResultReply)(nil), "rollupvalidator.GetMessageResultReply")
	proto.RegisterType((*GetAssertionCountArgs)(nil), "rollupvalidator.GetAssertionCountArgs")
	proto.RegisterType((*GetAssertionCountReply)(nil), "rollupvalidator.GetAssertionCountReply")
	proto.RegisterType((*GetVMInfoArgs)(nil), "rollupvalidator.GetVMInfoArgs")
	proto.RegisterType((*GetVMInfoReply)(nil), "rollupvalidator.GetVMInfoReply")
	proto.RegisterType((*CallMessageArgs)(nil), "rollupvalidator.CallMessageArgs")
	proto.RegisterType((*CallMessageReply)(nil), "rollupvalidator.CallMessageReply")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x56, 0x9a, 0x8f, 0x26, 0xd3, 0x8f, 0xf4, 0x5d, 0xf5, 0x2d, 0x56, 0x04, 0x25, 0x58, 0xa5,
	0x44, 0x15, 0x24, 0x52, 0x39, 0x23, 0x51, 0x8a, 0x68, 0x23, 0xb5, 0x08, 0x59, 0x28, 0x07, 0x6e,
	0x6b, 0x7b, 0xe3, 0x58, 0xdd, 0x78, 0xa3, 0xdd, 0x75, 0x29, 0x12, 0x17, 0x7e, 0x1d, 0x67, 0xfe,
	0x11, 0xda, 0x59, 0xc7, 0x71, 0xec, 0x96, 0xde, 0x3c, 0xcf, 0x7c, 0xed, 0xf3, 0xcc, 0x8c, 0x0c,
	0xdb, 0x8a, 0xc9, 0x5b, 0x26, 0x87, 0x0b, 0x29, 0xb4, 0x20, 0x5d, 0x29, 0x38, 0x4f, 0x17, 0xb7,
	0x94, 0xc7, 0x21, 0xd5, 0x42, 0xba, 0xbf, 0x36, 0x60, 0xf3, 0x4a, 0x44, 0xe3, 0x64, 0x2a, 0x88,
	0x03, 0x9b, 0x34, 0x0c, 0x25, 0x53, 0xca, 0xa9, 0xf5, 0x6b, 0x83, 0x8e, 0xb7, 0x34, 0xc9, 0x53,
	0xe8, 0xf8, 0x5c, 0x04, 0x37, 0x97, 0x54, 0xcd, 0x9c, 0x0d, 0xf4, 0xad, 0x00, 0xd2, 0x87, 0x2d,
	0x34, 0x3e, 0xa7, 0x73, 0x9f, 0x49, 0xa7, 0x8e, 0xfe, 0x22, 0x44, 0x08, 0x34, 0x42, 0xaa, 0xa9,
	0xd3, 0x40, 0x17, 0x7e, 0x93, 0x1e, 0xb4, 0xb9, 0x69, 0x1c, 0xb2, 0x3b, 0xa7, 0x89, 0x78, 0x6e,
	0x93, 0x03, 0x68, 0x69, 0xb1, 0x88, 0x03, 0xe5, 0xb4, 0xfa, 0xf5, 0x41, 0xc7, 0xcb, 0x2c, 0x72,
	0x02, 0x7b, 0x5a, 0xd2, 0x44, 0xd1, 0x40, 0xc7, 0x22, 0xb1, 0xb9, 0x9b, 0x98, 0x5b, 0xc1, 0xc9,
	0x00, 0xba, 0x05, 0x0c, 0x5f, 0xde, 0xc6, 0xd0, 0x32, 0xec, 0xfe, 0x84, 0xed, 0x4f, 0x71, 0x12,
	0x5e, 0x89, 0x48, 0x9d, 0xc9, 0x48, 0x91, 0x43, 0x80, 0xa9, 0x14, 0xf3, 0x4b, 0x16, 0x47, 0x33,
	0x9d, 0x49, 0x51, 0x40, 0xcc, 0xcb, 0xb5, 0xc8, 0xbc, 0x56, 0x8c, 0xdc, 0x2e, 0x6a, 0x58, 0x5f,
	0xd7, 0x70, 0xc5, 0xa9, 0x51, 0xe4, 0xe4, 0xbe, 0x83, 0x9d, 0x65, 0x77, 0x8f, 0x2d, 0xf8, 0x0f,
	0xf2, 0x1a, 0x1a, 0x5c, 0x44, 0x36, 0x6c, 0xeb, 0xd4, 0x19, 0x96, 0x46, 0x36, 0xcc, 0xc6, 0xe5,
	0x61, 0x94, 0x3b, 0x84, 0xfd, 0x0b, 0xa6, 0xaf, 0x99, 0x52, 0x34, 0x62, 0x1e, 0x53, 0x29, 0xd7,
	0x48, 0xc2, 0xb4, 0xbb, 0x43, 0xd6, 0x96, 0x40, 0x66, 0xb9, 0x7f, 0x6a, 0xf0, 0x7f, 0x39, 0xc1,
	0xf6, 0xdd, 0x87, 0xe6, 0x54, 0xa4, 0x49, 0x88, 0x09, 0x6d, 0xcf, 0x1a, 0xa6, 0x8e, 0xa4, 0xdf,
	0x27, 0x94, 0x67, 0x54, 0x33, 0xcb, 0x88, 0xc4, 0x45, 0xf4, 0x45, 0x32, 0xec, 0x61, 0xb9, 0x16,
	0x10, 0xb3, 0x14, 0xc6, 0x12, 0x4a, 0x63, 0x80, 0x9d, 0x7c, 0x11, 0x22, 0x2e, 0x6c, 0x73, 0x11,
	0x4d, 0x28, 0x37, 0x16, 0x53, 0x4e, 0x13, 0x65, 0x59, 0xc3, 0xc8, 0x11, 0xec, 0x88, 0xe4, 0x7c,
	0x46, 0xe3, 0xe4, 0xab, 0x25, 0xd3, 0xc2, 0x3a, 0xeb, 0xa0, 0xfb, 0x04, 0x29, 0x9d, 0x29, 0xc5,
	0xa4, 0x19, 0xea, 0xb9, 0x48, 0x13, 0x14, 0xc1, 0x7d, 0x0f, 0x07, 0x15, 0x87, 0x25, 0x7b, 0x0c,
	0xbb, 0x74, 0x0d, 0x46, 0xd6, 0x4d, 0xaf, 0x84, 0xba, 0x5d, 0xd8, 0xb9, 0x60, 0x7a, 0x72, 0x6d,
	0x14, 0xc7, 0x92, 0x47, 0xb0, 0x9b, 0x03, 0xb6, 0x14, 0x81, 0xc6, 0xed, 0x7c, 0xfc, 0x31, 0xd3,
	0x19, 0xbf, 0xdd, 0x08, 0xba, 0xe7, 0x94, 0xf3, 0x4c, 0x65, 0x1c, 0xc8, 0x00, 0xba, 0x81, 0x48,
	0xb4, 0xa4, 0x81, 0x3e, 0x5b, 0xbb, 0xb2, 0x32, 0x6c, 0x24, 0x57, 0x2c, 0x09, 0x99, 0x5c, 0x4a,
	0x6e, 0xad, 0xfc, 0x8a, 0xea, 0xab, 0x2b, 0x72, 0x4f, 0x60, 0xaf, 0xd0, 0xc8, 0x3e, 0x68, 0x35,
	0xb2, 0x5a, 0x71, 0x64, 0xa7, 0xbf, 0xeb, 0xd0, 0xf5, 0x70, 0x99, 0x26, 0xcb, 0x65, 0x22, 0x14,
	0xf6, 0xca, 0xdb, 0x40, 0x5e, 0x56, 0x56, 0xee, 0xbe, 0x0d, 0xeb, 0x1d, 0x3f, 0x1a, 0x66, 0x9f,
	0xe3, 0xc1, 0x56, 0xe1, 0x89, 0xa4, 0x5f, 0x49, 0x2b, 0x29, 0xd5, 0x7b, 0xf1, 0xaf, 0x08, 0x5b,
	0x73, 0x0c, 0xed, 0xe5, 0xd1, 0x90, 0x67, 0x95, 0xf0, 0xe2, 0x35, 0xf7, 0x0e, 0x1f, 0x74, 0xdb,
	0x52, 0x21, 0xfc, 0x57, 0xd9, 0x11, 0x72, 0x2f, 0xb7, 0xea, 0x82, 0xf5, 0x5e, 0x3d, 0x1e, 0x67,
	0xbb, 0x5c, 0x41, 0x27, 0x5f, 0x1b, 0x72, 0x78, 0x5f, 0xd6, 0x6a, 0xc7, 0x7a, 0xcf, 0x1f, 0xf6,
	0x63, 0xb5, 0x0f, 0xe3, 0x6f, 0x17, 0x51, 0xac, 0x67, 0xa9, 0x3f, 0x0c, 0xc4, 0x7c, 0x24, 0xa6,
	0xd3, 0xc0, 0x5c, 0x03, 0xa7, 0xbe, 0x1a, 0x51, 0xe9, 0xc7, 0x5a, 0xa6, 0xf3, 0xd1, 0x82, 0x06,
	0x37, 0x34, 0x62, 0x88, 0xbc, 0xc9, 0x2b, 0x8d, 0x4a, 0x95, 0xfd, 0x16, 0xfe, 0x18, 0xde, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x7d, 0xc8, 0xf9, 0x28, 0x06, 0x00, 0x00,
}