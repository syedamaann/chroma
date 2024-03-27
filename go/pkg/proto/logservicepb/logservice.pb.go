// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.4
// source: chromadb/proto/logservice.proto

package logservicepb

import (
	coordinatorpb "github.com/chroma-core/chroma/go/pkg/proto/coordinatorpb"
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

type PushLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId string                           `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Records      []*coordinatorpb.OperationRecord `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *PushLogsRequest) Reset() {
	*x = PushLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushLogsRequest) ProtoMessage() {}

func (x *PushLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushLogsRequest.ProtoReflect.Descriptor instead.
func (*PushLogsRequest) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{0}
}

func (x *PushLogsRequest) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *PushLogsRequest) GetRecords() []*coordinatorpb.OperationRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type PushLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordCount int32 `protobuf:"varint,1,opt,name=record_count,json=recordCount,proto3" json:"record_count,omitempty"`
}

func (x *PushLogsResponse) Reset() {
	*x = PushLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushLogsResponse) ProtoMessage() {}

func (x *PushLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushLogsResponse.ProtoReflect.Descriptor instead.
func (*PushLogsResponse) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{1}
}

func (x *PushLogsResponse) GetRecordCount() int32 {
	if x != nil {
		return x.RecordCount
	}
	return 0
}

type PullLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	StartFromId  int64  `protobuf:"varint,2,opt,name=start_from_id,json=startFromId,proto3" json:"start_from_id,omitempty"`
	BatchSize    int32  `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	EndTimestamp int64  `protobuf:"varint,4,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`
}

func (x *PullLogsRequest) Reset() {
	*x = PullLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullLogsRequest) ProtoMessage() {}

func (x *PullLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullLogsRequest.ProtoReflect.Descriptor instead.
func (*PullLogsRequest) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{2}
}

func (x *PullLogsRequest) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *PullLogsRequest) GetStartFromId() int64 {
	if x != nil {
		return x.StartFromId
	}
	return 0
}

func (x *PullLogsRequest) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *PullLogsRequest) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

// Represents an operation from the log
type LogRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogOffset int64                          `protobuf:"varint,1,opt,name=log_offset,json=logOffset,proto3" json:"log_offset,omitempty"`
	Record    *coordinatorpb.OperationRecord `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *LogRecord) Reset() {
	*x = LogRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRecord) ProtoMessage() {}

func (x *LogRecord) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRecord.ProtoReflect.Descriptor instead.
func (*LogRecord) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{3}
}

func (x *LogRecord) GetLogOffset() int64 {
	if x != nil {
		return x.LogOffset
	}
	return 0
}

func (x *LogRecord) GetRecord() *coordinatorpb.OperationRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type PullLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*LogRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *PullLogsResponse) Reset() {
	*x = PullLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullLogsResponse) ProtoMessage() {}

func (x *PullLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullLogsResponse.ProtoReflect.Descriptor instead.
func (*PullLogsResponse) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{4}
}

func (x *PullLogsResponse) GetRecords() []*LogRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type CollectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// The first log id of the collection that needs to be compacted
	FirstLogId   int64 `protobuf:"varint,2,opt,name=first_log_id,json=firstLogId,proto3" json:"first_log_id,omitempty"`
	FirstLogIdTs int64 `protobuf:"varint,3,opt,name=first_log_id_ts,json=firstLogIdTs,proto3" json:"first_log_id_ts,omitempty"`
}

func (x *CollectionInfo) Reset() {
	*x = CollectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionInfo) ProtoMessage() {}

func (x *CollectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionInfo.ProtoReflect.Descriptor instead.
func (*CollectionInfo) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{5}
}

func (x *CollectionInfo) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *CollectionInfo) GetFirstLogId() int64 {
	if x != nil {
		return x.FirstLogId
	}
	return 0
}

func (x *CollectionInfo) GetFirstLogIdTs() int64 {
	if x != nil {
		return x.FirstLogIdTs
	}
	return 0
}

type GetAllCollectionInfoToCompactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllCollectionInfoToCompactRequest) Reset() {
	*x = GetAllCollectionInfoToCompactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCollectionInfoToCompactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCollectionInfoToCompactRequest) ProtoMessage() {}

func (x *GetAllCollectionInfoToCompactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCollectionInfoToCompactRequest.ProtoReflect.Descriptor instead.
func (*GetAllCollectionInfoToCompactRequest) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{6}
}

type GetAllCollectionInfoToCompactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllCollectionInfo []*CollectionInfo `protobuf:"bytes,1,rep,name=all_collection_info,json=allCollectionInfo,proto3" json:"all_collection_info,omitempty"`
}

func (x *GetAllCollectionInfoToCompactResponse) Reset() {
	*x = GetAllCollectionInfoToCompactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chromadb_proto_logservice_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCollectionInfoToCompactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCollectionInfoToCompactResponse) ProtoMessage() {}

func (x *GetAllCollectionInfoToCompactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chromadb_proto_logservice_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCollectionInfoToCompactResponse.ProtoReflect.Descriptor instead.
func (*GetAllCollectionInfoToCompactResponse) Descriptor() ([]byte, []int) {
	return file_chromadb_proto_logservice_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllCollectionInfoToCompactResponse) GetAllCollectionInfo() []*CollectionInfo {
	if x != nil {
		return x.AllCollectionInfo
	}
	return nil
}

var File_chromadb_proto_logservice_proto protoreflect.FileDescriptor

var file_chromadb_proto_logservice_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x1a, 0x1b, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x61, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x31,
	0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x35, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x6c,
	0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x5b, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x3f, 0x0a, 0x10, 0x50, 0x75, 0x6c, 0x6c, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x7e, 0x0a, 0x0e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64,
	0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x49, 0x64, 0x54, 0x73, 0x22, 0x26, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x54,
	0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6f, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x5f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x61,
	0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0x8e, 0x02, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3f, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x08, 0x50, 0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x17, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x50,
	0x75, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x7e, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x12, 0x2c, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x54, 0x6f,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x72, 0x6f, 0x6d, 0x61, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x61, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chromadb_proto_logservice_proto_rawDescOnce sync.Once
	file_chromadb_proto_logservice_proto_rawDescData = file_chromadb_proto_logservice_proto_rawDesc
)

func file_chromadb_proto_logservice_proto_rawDescGZIP() []byte {
	file_chromadb_proto_logservice_proto_rawDescOnce.Do(func() {
		file_chromadb_proto_logservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_chromadb_proto_logservice_proto_rawDescData)
	})
	return file_chromadb_proto_logservice_proto_rawDescData
}

var file_chromadb_proto_logservice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chromadb_proto_logservice_proto_goTypes = []interface{}{
	(*PushLogsRequest)(nil),                       // 0: chroma.PushLogsRequest
	(*PushLogsResponse)(nil),                      // 1: chroma.PushLogsResponse
	(*PullLogsRequest)(nil),                       // 2: chroma.PullLogsRequest
	(*LogRecord)(nil),                             // 3: chroma.LogRecord
	(*PullLogsResponse)(nil),                      // 4: chroma.PullLogsResponse
	(*CollectionInfo)(nil),                        // 5: chroma.CollectionInfo
	(*GetAllCollectionInfoToCompactRequest)(nil),  // 6: chroma.GetAllCollectionInfoToCompactRequest
	(*GetAllCollectionInfoToCompactResponse)(nil), // 7: chroma.GetAllCollectionInfoToCompactResponse
	(*coordinatorpb.OperationRecord)(nil),         // 8: chroma.OperationRecord
}
var file_chromadb_proto_logservice_proto_depIdxs = []int32{
	8, // 0: chroma.PushLogsRequest.records:type_name -> chroma.OperationRecord
	8, // 1: chroma.LogRecord.record:type_name -> chroma.OperationRecord
	3, // 2: chroma.PullLogsResponse.records:type_name -> chroma.LogRecord
	5, // 3: chroma.GetAllCollectionInfoToCompactResponse.all_collection_info:type_name -> chroma.CollectionInfo
	0, // 4: chroma.LogService.PushLogs:input_type -> chroma.PushLogsRequest
	2, // 5: chroma.LogService.PullLogs:input_type -> chroma.PullLogsRequest
	6, // 6: chroma.LogService.GetAllCollectionInfoToCompact:input_type -> chroma.GetAllCollectionInfoToCompactRequest
	1, // 7: chroma.LogService.PushLogs:output_type -> chroma.PushLogsResponse
	4, // 8: chroma.LogService.PullLogs:output_type -> chroma.PullLogsResponse
	7, // 9: chroma.LogService.GetAllCollectionInfoToCompact:output_type -> chroma.GetAllCollectionInfoToCompactResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chromadb_proto_logservice_proto_init() }
func file_chromadb_proto_logservice_proto_init() {
	if File_chromadb_proto_logservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chromadb_proto_logservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushLogsRequest); i {
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
		file_chromadb_proto_logservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushLogsResponse); i {
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
		file_chromadb_proto_logservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullLogsRequest); i {
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
		file_chromadb_proto_logservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRecord); i {
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
		file_chromadb_proto_logservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PullLogsResponse); i {
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
		file_chromadb_proto_logservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionInfo); i {
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
		file_chromadb_proto_logservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCollectionInfoToCompactRequest); i {
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
		file_chromadb_proto_logservice_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCollectionInfoToCompactResponse); i {
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
			RawDescriptor: file_chromadb_proto_logservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chromadb_proto_logservice_proto_goTypes,
		DependencyIndexes: file_chromadb_proto_logservice_proto_depIdxs,
		MessageInfos:      file_chromadb_proto_logservice_proto_msgTypes,
	}.Build()
	File_chromadb_proto_logservice_proto = out.File
	file_chromadb_proto_logservice_proto_rawDesc = nil
	file_chromadb_proto_logservice_proto_goTypes = nil
	file_chromadb_proto_logservice_proto_depIdxs = nil
}
