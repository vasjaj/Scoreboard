// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: proto/scoreboard.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PlayerScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Points int64  `protobuf:"varint,2,opt,name=points,proto3" json:"points,omitempty"`
}

func (x *PlayerScoreRequest) Reset() {
	*x = PlayerScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scoreboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerScoreRequest) ProtoMessage() {}

func (x *PlayerScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scoreboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerScoreRequest.ProtoReflect.Descriptor instead.
func (*PlayerScoreRequest) Descriptor() ([]byte, []int) {
	return file_proto_scoreboard_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerScoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerScoreRequest) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

type PlayerScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank int64 `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *PlayerScoreResponse) Reset() {
	*x = PlayerScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scoreboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerScoreResponse) ProtoMessage() {}

func (x *PlayerScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scoreboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerScoreResponse.ProtoReflect.Descriptor instead.
func (*PlayerScoreResponse) Descriptor() ([]byte, []int) {
	return file_proto_scoreboard_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerScoreResponse) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type LeaderboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Page     int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Monthly  bool   `protobuf:"varint,4,opt,name=monthly,proto3" json:"monthly,omitempty"`
}

func (x *LeaderboardRequest) Reset() {
	*x = LeaderboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scoreboard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardRequest) ProtoMessage() {}

func (x *LeaderboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scoreboard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardRequest.ProtoReflect.Descriptor instead.
func (*LeaderboardRequest) Descriptor() ([]byte, []int) {
	return file_proto_scoreboard_proto_rawDescGZIP(), []int{2}
}

func (x *LeaderboardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LeaderboardRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *LeaderboardRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *LeaderboardRequest) GetMonthly() bool {
	if x != nil {
		return x.Monthly
	}
	return false
}

type LeaderboardScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Points   int64  `protobuf:"varint,2,opt,name=points,proto3" json:"points,omitempty"`
	Position int64  `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *LeaderboardScore) Reset() {
	*x = LeaderboardScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scoreboard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardScore) ProtoMessage() {}

func (x *LeaderboardScore) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scoreboard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardScore.ProtoReflect.Descriptor instead.
func (*LeaderboardScore) Descriptor() ([]byte, []int) {
	return file_proto_scoreboard_proto_rawDescGZIP(), []int{3}
}

func (x *LeaderboardScore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LeaderboardScore) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *LeaderboardScore) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

type LeaderboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextPage int64               `protobuf:"varint,1,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	Around   bool                `protobuf:"varint,2,opt,name=around,proto3" json:"around,omitempty"`
	Score    []*LeaderboardScore `protobuf:"bytes,3,rep,name=score,proto3" json:"score,omitempty"`
	AroundMe []*LeaderboardScore `protobuf:"bytes,4,rep,name=around_me,json=aroundMe,proto3" json:"around_me,omitempty"`
}

func (x *LeaderboardResponse) Reset() {
	*x = LeaderboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scoreboard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaderboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaderboardResponse) ProtoMessage() {}

func (x *LeaderboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scoreboard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaderboardResponse.ProtoReflect.Descriptor instead.
func (*LeaderboardResponse) Descriptor() ([]byte, []int) {
	return file_proto_scoreboard_proto_rawDescGZIP(), []int{4}
}

func (x *LeaderboardResponse) GetNextPage() int64 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

func (x *LeaderboardResponse) GetAround() bool {
	if x != nil {
		return x.Around
	}
	return false
}

func (x *LeaderboardResponse) GetScore() []*LeaderboardScore {
	if x != nil {
		return x.Score
	}
	return nil
}

func (x *LeaderboardResponse) GetAroundMe() []*LeaderboardScore {
	if x != nil {
		return x.AroundMe
	}
	return nil
}

var File_proto_scoreboard_proto protoreflect.FileDescriptor

var file_proto_scoreboard_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x12,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x29,
	0x0a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x73, 0x0a, 0x12, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x22, 0x5a,
	0x0a, 0x10, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xaf, 0x01, 0x0a, 0x13, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x08, 0x61, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x32, 0xdc, 0x01, 0x0a,
	0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x49, 0x0a, 0x0a, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x04, 0x53, 0x65, 0x65, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_scoreboard_proto_rawDescOnce sync.Once
	file_proto_scoreboard_proto_rawDescData = file_proto_scoreboard_proto_rawDesc
)

func file_proto_scoreboard_proto_rawDescGZIP() []byte {
	file_proto_scoreboard_proto_rawDescOnce.Do(func() {
		file_proto_scoreboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_scoreboard_proto_rawDescData)
	})
	return file_proto_scoreboard_proto_rawDescData
}

var file_proto_scoreboard_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_scoreboard_proto_goTypes = []interface{}{
	(*PlayerScoreRequest)(nil),  // 0: proto.PlayerScoreRequest
	(*PlayerScoreResponse)(nil), // 1: proto.PlayerScoreResponse
	(*LeaderboardRequest)(nil),  // 2: proto.LeaderboardRequest
	(*LeaderboardScore)(nil),    // 3: proto.LeaderboardScore
	(*LeaderboardResponse)(nil), // 4: proto.LeaderboardResponse
	(*empty.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_proto_scoreboard_proto_depIdxs = []int32{
	3, // 0: proto.LeaderboardResponse.score:type_name -> proto.LeaderboardScore
	3, // 1: proto.LeaderboardResponse.around_me:type_name -> proto.LeaderboardScore
	0, // 2: proto.Scoreboard.StoreScore:input_type -> proto.PlayerScoreRequest
	2, // 3: proto.Scoreboard.GetLeaderboard:input_type -> proto.LeaderboardRequest
	5, // 4: proto.Scoreboard.Seed:input_type -> google.protobuf.Empty
	1, // 5: proto.Scoreboard.StoreScore:output_type -> proto.PlayerScoreResponse
	4, // 6: proto.Scoreboard.GetLeaderboard:output_type -> proto.LeaderboardResponse
	5, // 7: proto.Scoreboard.Seed:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_scoreboard_proto_init() }
func file_proto_scoreboard_proto_init() {
	if File_proto_scoreboard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_scoreboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerScoreRequest); i {
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
		file_proto_scoreboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerScoreResponse); i {
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
		file_proto_scoreboard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardRequest); i {
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
		file_proto_scoreboard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardScore); i {
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
		file_proto_scoreboard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaderboardResponse); i {
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
			RawDescriptor: file_proto_scoreboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_scoreboard_proto_goTypes,
		DependencyIndexes: file_proto_scoreboard_proto_depIdxs,
		MessageInfos:      file_proto_scoreboard_proto_msgTypes,
	}.Build()
	File_proto_scoreboard_proto = out.File
	file_proto_scoreboard_proto_rawDesc = nil
	file_proto_scoreboard_proto_goTypes = nil
	file_proto_scoreboard_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ScoreboardClient is the client API for Scoreboard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScoreboardClient interface {
	StoreScore(ctx context.Context, opts ...grpc.CallOption) (Scoreboard_StoreScoreClient, error)
	GetLeaderboard(ctx context.Context, in *LeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error)
	Seed(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type scoreboardClient struct {
	cc grpc.ClientConnInterface
}

func NewScoreboardClient(cc grpc.ClientConnInterface) ScoreboardClient {
	return &scoreboardClient{cc}
}

func (c *scoreboardClient) StoreScore(ctx context.Context, opts ...grpc.CallOption) (Scoreboard_StoreScoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Scoreboard_serviceDesc.Streams[0], "/proto.Scoreboard/StoreScore", opts...)
	if err != nil {
		return nil, err
	}
	x := &scoreboardStoreScoreClient{stream}
	return x, nil
}

type Scoreboard_StoreScoreClient interface {
	Send(*PlayerScoreRequest) error
	Recv() (*PlayerScoreResponse, error)
	grpc.ClientStream
}

type scoreboardStoreScoreClient struct {
	grpc.ClientStream
}

func (x *scoreboardStoreScoreClient) Send(m *PlayerScoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scoreboardStoreScoreClient) Recv() (*PlayerScoreResponse, error) {
	m := new(PlayerScoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scoreboardClient) GetLeaderboard(ctx context.Context, in *LeaderboardRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error) {
	out := new(LeaderboardResponse)
	err := c.cc.Invoke(ctx, "/proto.Scoreboard/GetLeaderboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoreboardClient) Seed(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Scoreboard/Seed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreboardServer is the server API for Scoreboard service.
type ScoreboardServer interface {
	StoreScore(Scoreboard_StoreScoreServer) error
	GetLeaderboard(context.Context, *LeaderboardRequest) (*LeaderboardResponse, error)
	Seed(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedScoreboardServer can be embedded to have forward compatible implementations.
type UnimplementedScoreboardServer struct {
}

func (*UnimplementedScoreboardServer) StoreScore(Scoreboard_StoreScoreServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreScore not implemented")
}
func (*UnimplementedScoreboardServer) GetLeaderboard(context.Context, *LeaderboardRequest) (*LeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderboard not implemented")
}
func (*UnimplementedScoreboardServer) Seed(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seed not implemented")
}

func RegisterScoreboardServer(s *grpc.Server, srv ScoreboardServer) {
	s.RegisterService(&_Scoreboard_serviceDesc, srv)
}

func _Scoreboard_StoreScore_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScoreboardServer).StoreScore(&scoreboardStoreScoreServer{stream})
}

type Scoreboard_StoreScoreServer interface {
	Send(*PlayerScoreResponse) error
	Recv() (*PlayerScoreRequest, error)
	grpc.ServerStream
}

type scoreboardStoreScoreServer struct {
	grpc.ServerStream
}

func (x *scoreboardStoreScoreServer) Send(m *PlayerScoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scoreboardStoreScoreServer) Recv() (*PlayerScoreRequest, error) {
	m := new(PlayerScoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Scoreboard_GetLeaderboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreboardServer).GetLeaderboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Scoreboard/GetLeaderboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreboardServer).GetLeaderboard(ctx, req.(*LeaderboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scoreboard_Seed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreboardServer).Seed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Scoreboard/Seed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreboardServer).Seed(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scoreboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Scoreboard",
	HandlerType: (*ScoreboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeaderboard",
			Handler:    _Scoreboard_GetLeaderboard_Handler,
		},
		{
			MethodName: "Seed",
			Handler:    _Scoreboard_Seed_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StoreScore",
			Handler:       _Scoreboard_StoreScore_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/scoreboard.proto",
}