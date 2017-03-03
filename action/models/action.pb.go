// Code generated by protoc-gen-go.
// source: action.proto
// DO NOT EDIT!

/*
Package models is a generated protocol buffer package.

It is generated from these files:
	action.proto
	request.proto

It has these top-level messages:
	FollowAction
	LikeAction
	Request
*/
package models

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

type FollowAction struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Target      uint64 `protobuf:"varint,3,opt,name=target" json:"target,omitempty"`
	CreatedUtc  int64  `protobuf:"varint,14,opt,name=created_utc,json=createdUtc" json:"created_utc,omitempty"`
	ModifiedUtc int64  `protobuf:"varint,15,opt,name=modified_utc,json=modifiedUtc" json:"modified_utc,omitempty"`
	Deleted     bool   `protobuf:"varint,16,opt,name=deleted" json:"deleted,omitempty"`
	DeletedUtc  int64  `protobuf:"varint,17,opt,name=deleted_utc,json=deletedUtc" json:"deleted_utc,omitempty"`
}

func (m *FollowAction) Reset()                    { *m = FollowAction{} }
func (m *FollowAction) String() string            { return proto.CompactTextString(m) }
func (*FollowAction) ProtoMessage()               {}
func (*FollowAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FollowAction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FollowAction) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *FollowAction) GetTarget() uint64 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *FollowAction) GetCreatedUtc() int64 {
	if m != nil {
		return m.CreatedUtc
	}
	return 0
}

func (m *FollowAction) GetModifiedUtc() int64 {
	if m != nil {
		return m.ModifiedUtc
	}
	return 0
}

func (m *FollowAction) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *FollowAction) GetDeletedUtc() int64 {
	if m != nil {
		return m.DeletedUtc
	}
	return 0
}

type LikeAction struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Target     uint64 `protobuf:"varint,2,opt,name=target" json:"target,omitempty"`
	UserId     uint64 `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	IsFriend   bool   `protobuf:"varint,4,opt,name=is_friend,json=isFriend" json:"is_friend,omitempty"`
	Index      uint64 `protobuf:"varint,13,opt,name=index" json:"index,omitempty"`
	CreatedUtc int64  `protobuf:"varint,14,opt,name=created_utc,json=createdUtc" json:"created_utc,omitempty"`
	Deleted    bool   `protobuf:"varint,16,opt,name=deleted" json:"deleted,omitempty"`
	DeletedUtc int64  `protobuf:"varint,17,opt,name=deleted_utc,json=deletedUtc" json:"deleted_utc,omitempty"`
}

func (m *LikeAction) Reset()                    { *m = LikeAction{} }
func (m *LikeAction) String() string            { return proto.CompactTextString(m) }
func (*LikeAction) ProtoMessage()               {}
func (*LikeAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LikeAction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LikeAction) GetTarget() uint64 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *LikeAction) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LikeAction) GetIsFriend() bool {
	if m != nil {
		return m.IsFriend
	}
	return false
}

func (m *LikeAction) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *LikeAction) GetCreatedUtc() int64 {
	if m != nil {
		return m.CreatedUtc
	}
	return 0
}

func (m *LikeAction) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *LikeAction) GetDeletedUtc() int64 {
	if m != nil {
		return m.DeletedUtc
	}
	return 0
}

func init() {
	proto.RegisterType((*FollowAction)(nil), "models.FollowAction")
	proto.RegisterType((*LikeAction)(nil), "models.LikeAction")
}

func init() { proto.RegisterFile("action.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0xa4, 0xa6, 0x75, 0x1a, 0xab, 0x2e, 0xa2, 0x0b, 0x1e, 0x8c, 0x3d, 0xe5, 0xe4,
	0xc5, 0x27, 0xf0, 0x52, 0x10, 0x3c, 0x05, 0x7a, 0x0e, 0x71, 0x67, 0x2a, 0x83, 0x69, 0x56, 0x36,
	0x53, 0xf4, 0x29, 0x7d, 0x02, 0x1f, 0x46, 0xb2, 0x49, 0xd1, 0x1c, 0x44, 0xe8, 0x6d, 0xff, 0x6f,
	0xff, 0x19, 0xfe, 0x7f, 0x17, 0xd2, 0xca, 0x0a, 0xbb, 0xe6, 0xee, 0xcd, 0x3b, 0x71, 0x3a, 0xd9,
	0x3a, 0xa4, 0xba, 0x5d, 0x7e, 0x2a, 0x48, 0x57, 0xae, 0xae, 0xdd, 0xfb, 0x43, 0xb8, 0xd6, 0x0b,
	0x88, 0x18, 0x8d, 0xca, 0x54, 0x3e, 0x29, 0x22, 0x46, 0x7d, 0x05, 0xd3, 0x5d, 0x4b, 0xbe, 0x64,
	0x34, 0x51, 0x80, 0x49, 0x27, 0x1f, 0x51, 0x5f, 0x42, 0x22, 0x95, 0x7f, 0x21, 0x31, 0x71, 0xcf,
	0x7b, 0xa5, 0x6f, 0x60, 0x6e, 0x3d, 0x55, 0x42, 0x58, 0xee, 0xc4, 0x9a, 0x45, 0xa6, 0xf2, 0xb8,
	0x80, 0x01, 0xad, 0xc5, 0xea, 0x5b, 0x48, 0xb7, 0x0e, 0x79, 0xc3, 0x83, 0xe3, 0x34, 0x38, 0xe6,
	0x7b, 0xd6, 0x59, 0x0c, 0x4c, 0x91, 0x6a, 0x12, 0x42, 0x73, 0x96, 0xa9, 0x7c, 0x56, 0xec, 0x65,
	0xb7, 0x7d, 0x38, 0x86, 0xd9, 0xf3, 0x7e, 0xfb, 0x80, 0xd6, 0x62, 0x97, 0x5f, 0x0a, 0xe0, 0x89,
	0x5f, 0xe9, 0x8f, 0x3a, 0x3f, 0xa9, 0xa3, 0x51, 0xea, 0x5f, 0x35, 0xe3, 0x51, 0xcd, 0x6b, 0x38,
	0xe6, 0xb6, 0xdc, 0x78, 0xa6, 0x06, 0xcd, 0x24, 0x84, 0x99, 0x71, 0xbb, 0x0a, 0x5a, 0x5f, 0xc0,
	0x11, 0x37, 0x48, 0x1f, 0xe6, 0x24, 0xcc, 0xf4, 0xe2, 0xff, 0x17, 0x38, 0xbc, 0xde, 0x73, 0x12,
	0xbe, 0xef, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xee, 0xf7, 0x04, 0x18, 0xce, 0x01, 0x00, 0x00,
}