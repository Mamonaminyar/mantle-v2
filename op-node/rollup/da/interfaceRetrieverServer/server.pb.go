// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: server.proto

package interfaceRetrieverServer

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

type FramesAndDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*FramesAndDataRequest_DataStoreId
	//	*FramesAndDataRequest_DataConfirmHash
	Value isFramesAndDataRequest_Value `protobuf_oneof:"value"`
}

func (x *FramesAndDataRequest) Reset() {
	*x = FramesAndDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FramesAndDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FramesAndDataRequest) ProtoMessage() {}

func (x *FramesAndDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FramesAndDataRequest.ProtoReflect.Descriptor instead.
func (*FramesAndDataRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (m *FramesAndDataRequest) GetValue() isFramesAndDataRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *FramesAndDataRequest) GetDataStoreId() uint32 {
	if x, ok := x.GetValue().(*FramesAndDataRequest_DataStoreId); ok {
		return x.DataStoreId
	}
	return 0
}

func (x *FramesAndDataRequest) GetDataConfirmHash() string {
	if x, ok := x.GetValue().(*FramesAndDataRequest_DataConfirmHash); ok {
		return x.DataConfirmHash
	}
	return ""
}

type isFramesAndDataRequest_Value interface {
	isFramesAndDataRequest_Value()
}

type FramesAndDataRequest_DataStoreId struct {
	DataStoreId uint32 `protobuf:"varint,1,opt,name=DataStoreId,proto3,oneof"`
}

type FramesAndDataRequest_DataConfirmHash struct {
	DataConfirmHash string `protobuf:"bytes,2,opt,name=DataConfirmHash,proto3,oneof"`
}

func (*FramesAndDataRequest_DataStoreId) isFramesAndDataRequest_Value() {}

func (*FramesAndDataRequest_DataConfirmHash) isFramesAndDataRequest_Value() {}

type FramesAndDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Frames [][]byte `protobuf:"bytes,2,rep,name=Frames,proto3" json:"Frames,omitempty"`
}

func (x *FramesAndDataReply) Reset() {
	*x = FramesAndDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FramesAndDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FramesAndDataReply) ProtoMessage() {}

func (x *FramesAndDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FramesAndDataReply.ProtoReflect.Descriptor instead.
func (*FramesAndDataReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *FramesAndDataReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FramesAndDataReply) GetFrames() [][]byte {
	if x != nil {
		return x.Frames
	}
	return nil
}

type LastDataStoreIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LastDataStoreIdRequest) Reset() {
	*x = LastDataStoreIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastDataStoreIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastDataStoreIdRequest) ProtoMessage() {}

func (x *LastDataStoreIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastDataStoreIdRequest.ProtoReflect.Descriptor instead.
func (*LastDataStoreIdRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

type LastDataStoreIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataStoreId uint32 `protobuf:"varint,1,opt,name=DataStoreId,proto3" json:"DataStoreId,omitempty"`
}

func (x *LastDataStoreIdReply) Reset() {
	*x = LastDataStoreIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LastDataStoreIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LastDataStoreIdReply) ProtoMessage() {}

func (x *LastDataStoreIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LastDataStoreIdReply.ProtoReflect.Descriptor instead.
func (*LastDataStoreIdReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *LastDataStoreIdReply) GetDataStoreId() uint32 {
	if x != nil {
		return x.DataStoreId
	}
	return 0
}

type RetrieveNextDataStoreAndFrameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataStoreId uint32 `protobuf:"varint,1,opt,name=DataStoreId,proto3" json:"DataStoreId,omitempty"`
}

func (x *RetrieveNextDataStoreAndFrameRequest) Reset() {
	*x = RetrieveNextDataStoreAndFrameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveNextDataStoreAndFrameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveNextDataStoreAndFrameRequest) ProtoMessage() {}

func (x *RetrieveNextDataStoreAndFrameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveNextDataStoreAndFrameRequest.ProtoReflect.Descriptor instead.
func (*RetrieveNextDataStoreAndFrameRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *RetrieveNextDataStoreAndFrameRequest) GetDataStoreId() uint32 {
	if x != nil {
		return x.DataStoreId
	}
	return 0
}

type DataStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GUID               string `protobuf:"bytes,1,opt,name=GUID,proto3" json:"GUID,omitempty"`
	DataStoreId        uint64 `protobuf:"varint,2,opt,name=DataStoreId,proto3" json:"DataStoreId,omitempty"`
	DataConfirmHash    string `protobuf:"bytes,3,opt,name=DataConfirmHash,proto3" json:"DataConfirmHash,omitempty"`
	FromStoreNumber    uint64 `protobuf:"varint,4,opt,name=FromStoreNumber,proto3" json:"FromStoreNumber,omitempty"`
	ConfirmGasUsed     string `protobuf:"bytes,5,opt,name=ConfirmGasUsed,proto3" json:"ConfirmGasUsed,omitempty"`
	ConfirmBlockNumber uint64 `protobuf:"varint,6,opt,name=ConfirmBlockNumber,proto3" json:"ConfirmBlockNumber,omitempty"`
	DataHash           string `protobuf:"bytes,7,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
	SignatoryRecord    string `protobuf:"bytes,8,opt,name=SignatoryRecord,proto3" json:"SignatoryRecord,omitempty"`
	Status             bool   `protobuf:"varint,9,opt,name=Status,proto3" json:"Status,omitempty"`
	Confirmer          string `protobuf:"bytes,10,opt,name=Confirmer,proto3" json:"Confirmer,omitempty"`
	Header             string `protobuf:"bytes,11,opt,name=Header,proto3" json:"Header,omitempty"`
	DataCommitment     string `protobuf:"bytes,12,opt,name=DataCommitment,proto3" json:"DataCommitment,omitempty"`
	Timestamp          uint64 `protobuf:"varint,13,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	DataSize           uint64 `protobuf:"varint,14,opt,name=DataSize,proto3" json:"DataSize,omitempty"`
	Source             int32  `protobuf:"varint,15,opt,name=Source,proto3" json:"Source,omitempty"`
	SourceId           string `protobuf:"bytes,16,opt,name=SourceId,proto3" json:"SourceId,omitempty"`
}

func (x *DataStore) Reset() {
	*x = DataStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStore) ProtoMessage() {}

func (x *DataStore) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStore.ProtoReflect.Descriptor instead.
func (*DataStore) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *DataStore) GetGUID() string {
	if x != nil {
		return x.GUID
	}
	return ""
}

func (x *DataStore) GetDataStoreId() uint64 {
	if x != nil {
		return x.DataStoreId
	}
	return 0
}

func (x *DataStore) GetDataConfirmHash() string {
	if x != nil {
		return x.DataConfirmHash
	}
	return ""
}

func (x *DataStore) GetFromStoreNumber() uint64 {
	if x != nil {
		return x.FromStoreNumber
	}
	return 0
}

func (x *DataStore) GetConfirmGasUsed() string {
	if x != nil {
		return x.ConfirmGasUsed
	}
	return ""
}

func (x *DataStore) GetConfirmBlockNumber() uint64 {
	if x != nil {
		return x.ConfirmBlockNumber
	}
	return 0
}

func (x *DataStore) GetDataHash() string {
	if x != nil {
		return x.DataHash
	}
	return ""
}

func (x *DataStore) GetSignatoryRecord() string {
	if x != nil {
		return x.SignatoryRecord
	}
	return ""
}

func (x *DataStore) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *DataStore) GetConfirmer() string {
	if x != nil {
		return x.Confirmer
	}
	return ""
}

func (x *DataStore) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *DataStore) GetDataCommitment() string {
	if x != nil {
		return x.DataCommitment
	}
	return ""
}

func (x *DataStore) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *DataStore) GetDataSize() uint64 {
	if x != nil {
		return x.DataSize
	}
	return 0
}

func (x *DataStore) GetSource() int32 {
	if x != nil {
		return x.Source
	}
	return 0
}

func (x *DataStore) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

type DataStoreBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataStoreId uint32 `protobuf:"varint,1,opt,name=DataStoreId,proto3" json:"DataStoreId,omitempty"`
	BlockData   []byte `protobuf:"bytes,2,opt,name=BlockData,proto3" json:"BlockData,omitempty"`
	Canonical   bool   `protobuf:"varint,3,opt,name=Canonical,proto3" json:"Canonical,omitempty"`
	Timestamp   int64  `protobuf:"varint,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *DataStoreBlock) Reset() {
	*x = DataStoreBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStoreBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStoreBlock) ProtoMessage() {}

func (x *DataStoreBlock) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStoreBlock.ProtoReflect.Descriptor instead.
func (*DataStoreBlock) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{6}
}

func (x *DataStoreBlock) GetDataStoreId() uint32 {
	if x != nil {
		return x.DataStoreId
	}
	return 0
}

func (x *DataStoreBlock) GetBlockData() []byte {
	if x != nil {
		return x.BlockData
	}
	return nil
}

func (x *DataStoreBlock) GetCanonical() bool {
	if x != nil {
		return x.Canonical
	}
	return false
}

func (x *DataStoreBlock) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type RetrieveNextDataStoreAndFrameReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataStoreId    uint32          `protobuf:"varint,1,opt,name=DataStoreId,proto3" json:"DataStoreId,omitempty"`
	DataStore      *DataStore      `protobuf:"bytes,2,opt,name=DataStore,proto3" json:"DataStore,omitempty"`
	DataStoreBlock *DataStoreBlock `protobuf:"bytes,3,opt,name=DataStoreBlock,proto3" json:"DataStoreBlock,omitempty"`
}

func (x *RetrieveNextDataStoreAndFrameReply) Reset() {
	*x = RetrieveNextDataStoreAndFrameReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveNextDataStoreAndFrameReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveNextDataStoreAndFrameReply) ProtoMessage() {}

func (x *RetrieveNextDataStoreAndFrameReply) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveNextDataStoreAndFrameReply.ProtoReflect.Descriptor instead.
func (*RetrieveNextDataStoreAndFrameReply) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{7}
}

func (x *RetrieveNextDataStoreAndFrameReply) GetDataStoreId() uint32 {
	if x != nil {
		return x.DataStoreId
	}
	return 0
}

func (x *RetrieveNextDataStoreAndFrameReply) GetDataStore() *DataStore {
	if x != nil {
		return x.DataStore
	}
	return nil
}

func (x *RetrieveNextDataStoreAndFrameReply) GetDataStoreBlock() *DataStoreBlock {
	if x != nil {
		return x.DataStoreBlock
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x6f, 0x0a, 0x14, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0f, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x48, 0x61, 0x73, 0x68,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x06, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x4c,
	0x61, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x14, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22,
	0x48, 0x0a, 0x24, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x22, 0x97, 0x04, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x55, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x47, 0x55, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x46, 0x72, 0x6f, 0x6d, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0f, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x47, 0x61, 0x73, 0x55,
	0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x61, 0x73, 0x68, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x43, 0x61, 0x6e, 0x6f, 0x6e,
	0x69, 0x63, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0xdb, 0x01, 0x0a, 0x22, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4e,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x50,
	0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x32, 0xb1, 0x03, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x61, 0x6c, 0x12, 0x77, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x73, 0x41, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x41, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x41, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x1e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x30,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x73, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x9f, 0x01, 0x0a, 0x1d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4e,
	0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x6e, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x6a, 0x5a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x6c, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6f, 0x2f, 0x6c, 0x69, 0x74, 0x68, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x73,
	0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x74,
	0x6c, 0x65, 0x2d, 0x64, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_server_proto_goTypes = []interface{}{
	(*FramesAndDataRequest)(nil),                 // 0: interfaceRetrieverServer.FramesAndDataRequest
	(*FramesAndDataReply)(nil),                   // 1: interfaceRetrieverServer.FramesAndDataReply
	(*LastDataStoreIdRequest)(nil),               // 2: interfaceRetrieverServer.LastDataStoreIdRequest
	(*LastDataStoreIdReply)(nil),                 // 3: interfaceRetrieverServer.LastDataStoreIdReply
	(*RetrieveNextDataStoreAndFrameRequest)(nil), // 4: interfaceRetrieverServer.RetrieveNextDataStoreAndFrameRequest
	(*DataStore)(nil),                            // 5: interfaceRetrieverServer.DataStore
	(*DataStoreBlock)(nil),                       // 6: interfaceRetrieverServer.DataStoreBlock
	(*RetrieveNextDataStoreAndFrameReply)(nil),   // 7: interfaceRetrieverServer.RetrieveNextDataStoreAndFrameReply
}
var file_server_proto_depIdxs = []int32{
	5, // 0: interfaceRetrieverServer.RetrieveNextDataStoreAndFrameReply.DataStore:type_name -> interfaceRetrieverServer.DataStore
	6, // 1: interfaceRetrieverServer.RetrieveNextDataStoreAndFrameReply.DataStoreBlock:type_name -> interfaceRetrieverServer.DataStoreBlock
	0, // 2: interfaceRetrieverServer.DataRetrieval.RetrieveFramesAndData:input_type -> interfaceRetrieverServer.FramesAndDataRequest
	2, // 3: interfaceRetrieverServer.DataRetrieval.RetrieveLastConfirmDataStoreId:input_type -> interfaceRetrieverServer.LastDataStoreIdRequest
	4, // 4: interfaceRetrieverServer.DataRetrieval.RetrieveNextDataStoreAndFrame:input_type -> interfaceRetrieverServer.RetrieveNextDataStoreAndFrameRequest
	1, // 5: interfaceRetrieverServer.DataRetrieval.RetrieveFramesAndData:output_type -> interfaceRetrieverServer.FramesAndDataReply
	3, // 6: interfaceRetrieverServer.DataRetrieval.RetrieveLastConfirmDataStoreId:output_type -> interfaceRetrieverServer.LastDataStoreIdReply
	7, // 7: interfaceRetrieverServer.DataRetrieval.RetrieveNextDataStoreAndFrame:output_type -> interfaceRetrieverServer.RetrieveNextDataStoreAndFrameReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FramesAndDataRequest); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FramesAndDataReply); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastDataStoreIdRequest); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LastDataStoreIdReply); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveNextDataStoreAndFrameRequest); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStore); i {
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
		file_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStoreBlock); i {
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
		file_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveNextDataStoreAndFrameReply); i {
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
	file_server_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FramesAndDataRequest_DataStoreId)(nil),
		(*FramesAndDataRequest_DataConfirmHash)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
