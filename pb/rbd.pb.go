// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.19.4
// source: rbd.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Command_OpType int32

const (
	Command_CmdUnknown Command_OpType = 0
	Command_Get        Command_OpType = 1
	Command_Put        Command_OpType = 2
	Command_Delete     Command_OpType = 3
)

// Enum value maps for Command_OpType.
var (
	Command_OpType_name = map[int32]string{
		0: "CmdUnknown",
		1: "Get",
		2: "Put",
		3: "Delete",
	}
	Command_OpType_value = map[string]int32{
		"CmdUnknown": 0,
		"Get":        1,
		"Put":        2,
		"Delete":     3,
	}
)

func (x Command_OpType) Enum() *Command_OpType {
	p := new(Command_OpType)
	*p = x
	return p
}

func (x Command_OpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command_OpType) Descriptor() protoreflect.EnumDescriptor {
	return file_rbd_proto_enumTypes[0].Descriptor()
}

func (Command_OpType) Type() protoreflect.EnumType {
	return &file_rbd_proto_enumTypes[0]
}

func (x Command_OpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command_OpType.Descriptor instead.
func (Command_OpType) EnumDescriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{0, 0}
}

type Command_ReadConsistency int32

const (
	Command_RCUnknown    Command_ReadConsistency = 0
	Command_Serializable Command_ReadConsistency = 1
	Command_Linearizable Command_ReadConsistency = 2
)

// Enum value maps for Command_ReadConsistency.
var (
	Command_ReadConsistency_name = map[int32]string{
		0: "RCUnknown",
		1: "Serializable",
		2: "Linearizable",
	}
	Command_ReadConsistency_value = map[string]int32{
		"RCUnknown":    0,
		"Serializable": 1,
		"Linearizable": 2,
	}
)

func (x Command_ReadConsistency) Enum() *Command_ReadConsistency {
	p := new(Command_ReadConsistency)
	*p = x
	return p
}

func (x Command_ReadConsistency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command_ReadConsistency) Descriptor() protoreflect.EnumDescriptor {
	return file_rbd_proto_enumTypes[1].Descriptor()
}

func (Command_ReadConsistency) Type() protoreflect.EnumType {
	return &file_rbd_proto_enumTypes[1]
}

func (x Command_ReadConsistency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command_ReadConsistency.Descriptor instead.
func (Command_ReadConsistency) EnumDescriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{0, 1}
}

type Event_EventType int32

const (
	Event_EventUnknown Event_EventType = 0
	Event_Put          Event_EventType = 1
	Event_Delete       Event_EventType = 2
)

// Enum value maps for Event_EventType.
var (
	Event_EventType_name = map[int32]string{
		0: "EventUnknown",
		1: "Put",
		2: "Delete",
	}
	Event_EventType_value = map[string]int32{
		"EventUnknown": 0,
		"Put":          1,
		"Delete":       2,
	}
)

func (x Event_EventType) Enum() *Event_EventType {
	p := new(Event_EventType)
	*p = x
	return p
}

func (x Event_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_rbd_proto_enumTypes[2].Descriptor()
}

func (Event_EventType) Type() protoreflect.EnumType {
	return &file_rbd_proto_enumTypes[2]
}

func (x Event_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventType.Descriptor instead.
func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{3, 0}
}

type Command struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Op            Command_OpType          `protobuf:"varint,1,opt,name=op,proto3,enum=rbdkv.Command_OpType" json:"op,omitempty"`
	Rc            Command_ReadConsistency `protobuf:"varint,2,opt,name=rc,proto3,enum=rbdkv.Command_ReadConsistency" json:"rc,omitempty"`
	Key           []byte                  `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                  `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Ttl           uint64                  `protobuf:"varint,5,opt,name=ttl,proto3" json:"ttl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Command) Reset() {
	*x = Command{}
	mi := &file_rbd_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{0}
}

func (x *Command) GetOp() Command_OpType {
	if x != nil {
		return x.Op
	}
	return Command_CmdUnknown
}

func (x *Command) GetRc() Command_ReadConsistency {
	if x != nil {
		return x.Rc
	}
	return Command_RCUnknown
}

func (x *Command) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Command) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Command) GetTtl() uint64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type CommandResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           []byte                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	mi := &file_rbd_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{1}
}

func (x *CommandResponse) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *CommandResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type WatchRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	WatcherId      string                 `protobuf:"bytes,1,opt,name=watcher_id,json=watcherId,proto3" json:"watcher_id,omitempty"`
	Prefixes       [][]byte               `protobuf:"bytes,2,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	EventCapacity  uint64                 `protobuf:"varint,3,opt,name=event_capacity,json=eventCapacity,proto3" json:"event_capacity,omitempty"`
	LeaderRequired bool                   `protobuf:"varint,4,opt,name=leader_required,json=leaderRequired,proto3" json:"leader_required,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	mi := &file_rbd_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{2}
}

func (x *WatchRequest) GetWatcherId() string {
	if x != nil {
		return x.WatcherId
	}
	return ""
}

func (x *WatchRequest) GetPrefixes() [][]byte {
	if x != nil {
		return x.Prefixes
	}
	return nil
}

func (x *WatchRequest) GetEventCapacity() uint64 {
	if x != nil {
		return x.EventCapacity
	}
	return 0
}

func (x *WatchRequest) GetLeaderRequired() bool {
	if x != nil {
		return x.LeaderRequired
	}
	return false
}

type Event struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          Event_EventType        `protobuf:"varint,1,opt,name=type,proto3,enum=rbdkv.Event_EventType" json:"type,omitempty"`
	Key           []byte                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value         []byte                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Version       uint64                 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	ExpireAt      uint64                 `protobuf:"varint,5,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	WatcherId     string                 `protobuf:"bytes,6,opt,name=watcher_id,json=watcherId,proto3" json:"watcher_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_rbd_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetType() Event_EventType {
	if x != nil {
		return x.Type
	}
	return Event_EventUnknown
}

func (x *Event) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Event) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Event) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Event) GetExpireAt() uint64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *Event) GetWatcherId() string {
	if x != nil {
		return x.WatcherId
	}
	return ""
}

type WatchResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WatcherId     string                 `protobuf:"bytes,1,opt,name=watcher_id,json=watcherId,proto3" json:"watcher_id,omitempty"`
	Event         *Event                 `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	mi := &file_rbd_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{4}
}

func (x *WatchResponse) GetWatcherId() string {
	if x != nil {
		return x.WatcherId
	}
	return ""
}

func (x *WatchResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type LeaderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LeaderRequest) Reset() {
	*x = LeaderRequest{}
	mi := &file_rbd_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderRequest) ProtoMessage() {}

func (x *LeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderRequest.ProtoReflect.Descriptor instead.
func (*LeaderRequest) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{5}
}

type LeaderInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LeaderAddr    string                 `protobuf:"bytes,1,opt,name=leader_addr,json=leaderAddr,proto3" json:"leader_addr,omitempty"`
	LeaderId      string                 `protobuf:"bytes,2,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
	Term          uint64                 `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LeaderInfoResponse) Reset() {
	*x = LeaderInfoResponse{}
	mi := &file_rbd_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LeaderInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderInfoResponse) ProtoMessage() {}

func (x *LeaderInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderInfoResponse.ProtoReflect.Descriptor instead.
func (*LeaderInfoResponse) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{6}
}

func (x *LeaderInfoResponse) GetLeaderAddr() string {
	if x != nil {
		return x.LeaderAddr
	}
	return ""
}

func (x *LeaderInfoResponse) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

func (x *LeaderInfoResponse) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

type JoinRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr          string                 `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	mi := &file_rbd_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{7}
}

func (x *JoinRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JoinRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type JoinResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LeaderInfo    *LeaderInfoResponse    `protobuf:"bytes,1,opt,name=leader_info,json=leaderInfo,proto3" json:"leader_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	mi := &file_rbd_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rbd_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_rbd_proto_rawDescGZIP(), []int{8}
}

func (x *JoinResponse) GetLeaderInfo() *LeaderInfoResponse {
	if x != nil {
		return x.LeaderInfo
	}
	return nil
}

var File_rbd_proto protoreflect.FileDescriptor

var file_rbd_proto_rawDesc = string([]byte{
	0x0a, 0x09, 0x72, 0x62, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x62, 0x64,
	0x6b, 0x76, 0x22, 0x98, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x25,
	0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x62, 0x64,
	0x6b, 0x76, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x2e, 0x0a, 0x02, 0x72, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x02, 0x72, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22,
	0x36, 0x0a, 0x06, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6d, 0x64,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x03, 0x22, 0x44, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x43,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4c,
	0x69, 0x6e, 0x65, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x02, 0x22, 0x39, 0x0a,
	0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x0c, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x22, 0xe5, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72,
	0x62, 0x64, 0x6b, 0x76, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x02, 0x22, 0x52, 0x0a, 0x0d,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x62,
	0x64, 0x6b, 0x76, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x66, 0x0a, 0x12, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x22, 0x31, 0x0a, 0x0b, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x4a, 0x0a, 0x0c,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb5, 0x01, 0x0a, 0x05, 0x52, 0x62, 0x64,
	0x6b, 0x76, 0x12, 0x33, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x2e,
	0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x16, 0x2e,
	0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x62,
	0x64, 0x6b, 0x76, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x13, 0x2e, 0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x32, 0x3f, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x72, 0x62, 0x64, 0x6b, 0x76, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x62, 0x64,
	0x6b, 0x76, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x4c, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x62, 0x64,
	0x2d, 0x6b, 0x76, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_rbd_proto_rawDescOnce sync.Once
	file_rbd_proto_rawDescData []byte
)

func file_rbd_proto_rawDescGZIP() []byte {
	file_rbd_proto_rawDescOnce.Do(func() {
		file_rbd_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rbd_proto_rawDesc), len(file_rbd_proto_rawDesc)))
	})
	return file_rbd_proto_rawDescData
}

var file_rbd_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_rbd_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rbd_proto_goTypes = []any{
	(Command_OpType)(0),          // 0: rbdkv.Command.OpType
	(Command_ReadConsistency)(0), // 1: rbdkv.Command.ReadConsistency
	(Event_EventType)(0),         // 2: rbdkv.Event.EventType
	(*Command)(nil),              // 3: rbdkv.Command
	(*CommandResponse)(nil),      // 4: rbdkv.CommandResponse
	(*WatchRequest)(nil),         // 5: rbdkv.WatchRequest
	(*Event)(nil),                // 6: rbdkv.Event
	(*WatchResponse)(nil),        // 7: rbdkv.WatchResponse
	(*LeaderRequest)(nil),        // 8: rbdkv.LeaderRequest
	(*LeaderInfoResponse)(nil),   // 9: rbdkv.LeaderInfoResponse
	(*JoinRequest)(nil),          // 10: rbdkv.JoinRequest
	(*JoinResponse)(nil),         // 11: rbdkv.JoinResponse
}
var file_rbd_proto_depIdxs = []int32{
	0,  // 0: rbdkv.Command.op:type_name -> rbdkv.Command.OpType
	1,  // 1: rbdkv.Command.rc:type_name -> rbdkv.Command.ReadConsistency
	2,  // 2: rbdkv.Event.type:type_name -> rbdkv.Event.EventType
	6,  // 3: rbdkv.WatchResponse.event:type_name -> rbdkv.Event
	9,  // 4: rbdkv.JoinResponse.leader_info:type_name -> rbdkv.LeaderInfoResponse
	3,  // 5: rbdkv.Rbdkv.Execute:input_type -> rbdkv.Command
	8,  // 6: rbdkv.Rbdkv.LeaderInfo:input_type -> rbdkv.LeaderRequest
	5,  // 7: rbdkv.Rbdkv.Watch:input_type -> rbdkv.WatchRequest
	10, // 8: rbdkv.Controller.Join:input_type -> rbdkv.JoinRequest
	4,  // 9: rbdkv.Rbdkv.Execute:output_type -> rbdkv.CommandResponse
	9,  // 10: rbdkv.Rbdkv.LeaderInfo:output_type -> rbdkv.LeaderInfoResponse
	7,  // 11: rbdkv.Rbdkv.Watch:output_type -> rbdkv.WatchResponse
	11, // 12: rbdkv.Controller.Join:output_type -> rbdkv.JoinResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rbd_proto_init() }
func file_rbd_proto_init() {
	if File_rbd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rbd_proto_rawDesc), len(file_rbd_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_rbd_proto_goTypes,
		DependencyIndexes: file_rbd_proto_depIdxs,
		EnumInfos:         file_rbd_proto_enumTypes,
		MessageInfos:      file_rbd_proto_msgTypes,
	}.Build()
	File_rbd_proto = out.File
	file_rbd_proto_goTypes = nil
	file_rbd_proto_depIdxs = nil
}
