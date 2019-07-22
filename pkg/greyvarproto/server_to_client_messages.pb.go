// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server_to_client_messages.proto

package greyvarproto

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

type MoveResponse struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveResponse) Reset()         { *m = MoveResponse{} }
func (m *MoveResponse) String() string { return proto.CompactTextString(m) }
func (*MoveResponse) ProtoMessage()    {}
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{0}
}

func (m *MoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveResponse.Unmarshal(m, b)
}
func (m *MoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveResponse.Marshal(b, m, deterministic)
}
func (m *MoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveResponse.Merge(m, src)
}
func (m *MoveResponse) XXX_Size() int {
	return xxx_messageInfo_MoveResponse.Size(m)
}
func (m *MoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MoveResponse proto.InternalMessageInfo

func (m *MoveResponse) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type EntityMoved struct {
	EntityId             int32    `protobuf:"varint,1,opt,name=entityId,proto3" json:"entityId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityMoved) Reset()         { *m = EntityMoved{} }
func (m *EntityMoved) String() string { return proto.CompactTextString(m) }
func (*EntityMoved) ProtoMessage()    {}
func (*EntityMoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{1}
}

func (m *EntityMoved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityMoved.Unmarshal(m, b)
}
func (m *EntityMoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityMoved.Marshal(b, m, deterministic)
}
func (m *EntityMoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityMoved.Merge(m, src)
}
func (m *EntityMoved) XXX_Size() int {
	return xxx_messageInfo_EntityMoved.Size(m)
}
func (m *EntityMoved) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityMoved.DiscardUnknown(m)
}

var xxx_messageInfo_EntityMoved proto.InternalMessageInfo

func (m *EntityMoved) GetEntityId() int32 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

type PlayerQuit struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerQuit) Reset()         { *m = PlayerQuit{} }
func (m *PlayerQuit) String() string { return proto.CompactTextString(m) }
func (*PlayerQuit) ProtoMessage()    {}
func (*PlayerQuit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{2}
}

func (m *PlayerQuit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerQuit.Unmarshal(m, b)
}
func (m *PlayerQuit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerQuit.Marshal(b, m, deterministic)
}
func (m *PlayerQuit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerQuit.Merge(m, src)
}
func (m *PlayerQuit) XXX_Size() int {
	return xxx_messageInfo_PlayerQuit.Size(m)
}
func (m *PlayerQuit) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerQuit.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerQuit proto.InternalMessageInfo

func (m *PlayerQuit) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type PlayerAlreadyHere struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerAlreadyHere) Reset()         { *m = PlayerAlreadyHere{} }
func (m *PlayerAlreadyHere) String() string { return proto.CompactTextString(m) }
func (*PlayerAlreadyHere) ProtoMessage()    {}
func (*PlayerAlreadyHere) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{3}
}

func (m *PlayerAlreadyHere) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerAlreadyHere.Unmarshal(m, b)
}
func (m *PlayerAlreadyHere) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerAlreadyHere.Marshal(b, m, deterministic)
}
func (m *PlayerAlreadyHere) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerAlreadyHere.Merge(m, src)
}
func (m *PlayerAlreadyHere) XXX_Size() int {
	return xxx_messageInfo_PlayerAlreadyHere.Size(m)
}
func (m *PlayerAlreadyHere) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerAlreadyHere.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerAlreadyHere proto.InternalMessageInfo

func (m *PlayerAlreadyHere) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type PlayerJoined struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerJoined) Reset()         { *m = PlayerJoined{} }
func (m *PlayerJoined) String() string { return proto.CompactTextString(m) }
func (*PlayerJoined) ProtoMessage()    {}
func (*PlayerJoined) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{4}
}

func (m *PlayerJoined) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerJoined.Unmarshal(m, b)
}
func (m *PlayerJoined) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerJoined.Marshal(b, m, deterministic)
}
func (m *PlayerJoined) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerJoined.Merge(m, src)
}
func (m *PlayerJoined) XXX_Size() int {
	return xxx_messageInfo_PlayerJoined.Size(m)
}
func (m *PlayerJoined) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerJoined.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerJoined proto.InternalMessageInfo

func (m *PlayerJoined) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type PlayerYou struct {
	PlayerId             int32    `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerYou) Reset()         { *m = PlayerYou{} }
func (m *PlayerYou) String() string { return proto.CompactTextString(m) }
func (*PlayerYou) ProtoMessage()    {}
func (*PlayerYou) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{5}
}

func (m *PlayerYou) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerYou.Unmarshal(m, b)
}
func (m *PlayerYou) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerYou.Marshal(b, m, deterministic)
}
func (m *PlayerYou) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerYou.Merge(m, src)
}
func (m *PlayerYou) XXX_Size() int {
	return xxx_messageInfo_PlayerYou.Size(m)
}
func (m *PlayerYou) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerYou.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerYou proto.InternalMessageInfo

func (m *PlayerYou) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type EntitySpawn struct {
	EntityId             int32    `protobuf:"varint,1,opt,name=entityId,proto3" json:"entityId,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	W                    int32    `protobuf:"varint,4,opt,name=w,proto3" json:"w,omitempty"`
	H                    int32    `protobuf:"varint,5,opt,name=h,proto3" json:"h,omitempty"`
	Texture              string   `protobuf:"bytes,6,opt,name=texture,proto3" json:"texture,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntitySpawn) Reset()         { *m = EntitySpawn{} }
func (m *EntitySpawn) String() string { return proto.CompactTextString(m) }
func (*EntitySpawn) ProtoMessage()    {}
func (*EntitySpawn) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{6}
}

func (m *EntitySpawn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntitySpawn.Unmarshal(m, b)
}
func (m *EntitySpawn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntitySpawn.Marshal(b, m, deterministic)
}
func (m *EntitySpawn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntitySpawn.Merge(m, src)
}
func (m *EntitySpawn) XXX_Size() int {
	return xxx_messageInfo_EntitySpawn.Size(m)
}
func (m *EntitySpawn) XXX_DiscardUnknown() {
	xxx_messageInfo_EntitySpawn.DiscardUnknown(m)
}

var xxx_messageInfo_EntitySpawn proto.InternalMessageInfo

func (m *EntitySpawn) GetEntityId() int32 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *EntitySpawn) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *EntitySpawn) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *EntitySpawn) GetW() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *EntitySpawn) GetH() int32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *EntitySpawn) GetTexture() string {
	if m != nil {
		return m.Texture
	}
	return ""
}

type Tile struct {
	Col                  int32    `protobuf:"varint,1,opt,name=col,proto3" json:"col,omitempty"`
	Row                  int32    `protobuf:"varint,2,opt,name=row,proto3" json:"row,omitempty"`
	Tex                  string   `protobuf:"bytes,3,opt,name=tex,proto3" json:"tex,omitempty"`
	FlipH                bool     `protobuf:"varint,4,opt,name=flipH,proto3" json:"flipH,omitempty"`
	FlipV                bool     `protobuf:"varint,5,opt,name=flipV,proto3" json:"flipV,omitempty"`
	Rot                  int32    `protobuf:"varint,6,opt,name=rot,proto3" json:"rot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tile) Reset()         { *m = Tile{} }
func (m *Tile) String() string { return proto.CompactTextString(m) }
func (*Tile) ProtoMessage()    {}
func (*Tile) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{7}
}

func (m *Tile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tile.Unmarshal(m, b)
}
func (m *Tile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tile.Marshal(b, m, deterministic)
}
func (m *Tile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tile.Merge(m, src)
}
func (m *Tile) XXX_Size() int {
	return xxx_messageInfo_Tile.Size(m)
}
func (m *Tile) XXX_DiscardUnknown() {
	xxx_messageInfo_Tile.DiscardUnknown(m)
}

var xxx_messageInfo_Tile proto.InternalMessageInfo

func (m *Tile) GetCol() int32 {
	if m != nil {
		return m.Col
	}
	return 0
}

func (m *Tile) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *Tile) GetTex() string {
	if m != nil {
		return m.Tex
	}
	return ""
}

func (m *Tile) GetFlipH() bool {
	if m != nil {
		return m.FlipH
	}
	return false
}

func (m *Tile) GetFlipV() bool {
	if m != nil {
		return m.FlipV
	}
	return false
}

func (m *Tile) GetRot() int32 {
	if m != nil {
		return m.Rot
	}
	return 0
}

type Grid struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Tiles                []*Tile  `protobuf:"bytes,2,rep,name=tiles,proto3" json:"tiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Grid) Reset()         { *m = Grid{} }
func (m *Grid) String() string { return proto.CompactTextString(m) }
func (*Grid) ProtoMessage()    {}
func (*Grid) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{8}
}

func (m *Grid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grid.Unmarshal(m, b)
}
func (m *Grid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grid.Marshal(b, m, deterministic)
}
func (m *Grid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grid.Merge(m, src)
}
func (m *Grid) XXX_Size() int {
	return xxx_messageInfo_Grid.Size(m)
}
func (m *Grid) XXX_DiscardUnknown() {
	xxx_messageInfo_Grid.DiscardUnknown(m)
}

var xxx_messageInfo_Grid proto.InternalMessageInfo

func (m *Grid) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Grid) GetTiles() []*Tile {
	if m != nil {
		return m.Tiles
	}
	return nil
}

type ConnectionResponse struct {
	ServerVersion        string   `protobuf:"bytes,1,opt,name=serverVersion,proto3" json:"serverVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionResponse) Reset()         { *m = ConnectionResponse{} }
func (m *ConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectionResponse) ProtoMessage()    {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{9}
}

func (m *ConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionResponse.Unmarshal(m, b)
}
func (m *ConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionResponse.Marshal(b, m, deterministic)
}
func (m *ConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionResponse.Merge(m, src)
}
func (m *ConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectionResponse.Size(m)
}
func (m *ConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionResponse proto.InternalMessageInfo

func (m *ConnectionResponse) GetServerVersion() string {
	if m != nil {
		return m.ServerVersion
	}
	return ""
}

type NoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoResponse) Reset()         { *m = NoResponse{} }
func (m *NoResponse) String() string { return proto.CompactTextString(m) }
func (*NoResponse) ProtoMessage()    {}
func (*NoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f34367b8a54972c, []int{10}
}

func (m *NoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoResponse.Unmarshal(m, b)
}
func (m *NoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoResponse.Marshal(b, m, deterministic)
}
func (m *NoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoResponse.Merge(m, src)
}
func (m *NoResponse) XXX_Size() int {
	return xxx_messageInfo_NoResponse.Size(m)
}
func (m *NoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MoveResponse)(nil), "greyvarproto.MoveResponse")
	proto.RegisterType((*EntityMoved)(nil), "greyvarproto.EntityMoved")
	proto.RegisterType((*PlayerQuit)(nil), "greyvarproto.PlayerQuit")
	proto.RegisterType((*PlayerAlreadyHere)(nil), "greyvarproto.PlayerAlreadyHere")
	proto.RegisterType((*PlayerJoined)(nil), "greyvarproto.PlayerJoined")
	proto.RegisterType((*PlayerYou)(nil), "greyvarproto.PlayerYou")
	proto.RegisterType((*EntitySpawn)(nil), "greyvarproto.EntitySpawn")
	proto.RegisterType((*Tile)(nil), "greyvarproto.Tile")
	proto.RegisterType((*Grid)(nil), "greyvarproto.Grid")
	proto.RegisterType((*ConnectionResponse)(nil), "greyvarproto.ConnectionResponse")
	proto.RegisterType((*NoResponse)(nil), "greyvarproto.NoResponse")
}

func init() { proto.RegisterFile("server_to_client_messages.proto", fileDescriptor_5f34367b8a54972c) }

var fileDescriptor_5f34367b8a54972c = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0x71, 0xfe, 0x35, 0x99, 0xb8, 0xd0, 0x8a, 0x1e, 0x44, 0x2f, 0x0d, 0xa2, 0x50, 0x37,
	0x87, 0x14, 0xda, 0x5b, 0x6f, 0xa5, 0xb4, 0x4d, 0x0b, 0x2d, 0x5d, 0xef, 0x12, 0xd8, 0x53, 0xf0,
	0xc6, 0xb3, 0x89, 0x40, 0x91, 0x8c, 0x24, 0xc7, 0xf6, 0xee, 0x97, 0x5f, 0x24, 0xd9, 0xd9, 0xdd,
	0x8b, 0x6f, 0xfa, 0x3d, 0xde, 0xbc, 0x37, 0x36, 0x03, 0xef, 0x0c, 0xea, 0x13, 0xea, 0xad, 0x55,
	0xdb, 0x9d, 0xe0, 0x28, 0xed, 0xf6, 0x88, 0xc6, 0x64, 0x7b, 0x34, 0xab, 0x42, 0x2b, 0xab, 0x48,
	0xbc, 0xd7, 0xd8, 0x9c, 0x32, 0xed, 0x89, 0x2d, 0x21, 0xfe, 0xab, 0x4e, 0x98, 0xa2, 0x29, 0x94,
	0x34, 0x48, 0xde, 0xc2, 0xb4, 0x10, 0x59, 0x83, 0xfa, 0x77, 0x4e, 0xa3, 0x45, 0x94, 0x8c, 0xd3,
	0x33, 0xb3, 0x8f, 0x30, 0xff, 0x21, 0x2d, 0xb7, 0x8d, 0x9b, 0xc8, 0x9d, 0x15, 0x3d, 0x3e, 0x5a,
	0x3b, 0x66, 0x09, 0xc0, 0x7f, 0x3f, 0x76, 0x51, 0x72, 0xdb, 0x1b, 0xfa, 0x09, 0x5e, 0x07, 0xe7,
	0x37, 0xa1, 0x31, 0xcb, 0x9b, 0x35, 0xea, 0xfe, 0x2d, 0x96, 0x10, 0x87, 0x81, 0x3f, 0x8a, 0xcb,
	0xb0, 0x46, 0x69, 0x50, 0xcb, 0xec, 0x88, 0xde, 0x3b, 0x4b, 0xcf, 0xcc, 0x3e, 0xc0, 0x2c, 0x78,
	0xaf, 0x55, 0xd9, 0x1b, 0x7a, 0xdf, 0x7d, 0xda, 0x65, 0x91, 0x55, 0xb2, 0xef, 0xd3, 0x48, 0x0c,
	0x51, 0x4d, 0x07, 0x5e, 0x8c, 0x6a, 0x47, 0x0d, 0x1d, 0x06, 0x6a, 0x1c, 0x55, 0x74, 0x14, 0xa8,
	0x72, 0x74, 0xa0, 0xe3, 0x40, 0x07, 0x42, 0xe1, 0x85, 0xc5, 0xda, 0x96, 0x1a, 0xe9, 0xc4, 0xaf,
	0xd9, 0x21, 0xbb, 0x83, 0xd1, 0x15, 0x17, 0x48, 0x5e, 0xc1, 0x70, 0xa7, 0x44, 0x5b, 0xe8, 0x9e,
	0x4e, 0xd1, 0xaa, 0x6a, 0xdb, 0xdc, 0xd3, 0x29, 0x16, 0x6b, 0xdf, 0x38, 0x4b, 0xdd, 0x93, 0xbc,
	0x81, 0xf1, 0xad, 0xe0, 0xc5, 0xda, 0xf7, 0x4e, 0xd3, 0x00, 0x9d, 0xba, 0xf1, 0xfd, 0xad, 0xba,
	0x09, 0x79, 0xd6, 0xf7, 0xfb, 0x3c, 0xcb, 0x7e, 0xc2, 0xe8, 0x97, 0xe6, 0xb9, 0xf3, 0x5b, 0x6e,
	0x45, 0xf7, 0x0b, 0x03, 0x90, 0xc4, 0xa9, 0x02, 0x0d, 0x1d, 0x2c, 0x86, 0xc9, 0xfc, 0x33, 0x59,
	0x3d, 0xbd, 0x9d, 0x95, 0x5b, 0x3a, 0x0d, 0x06, 0xf6, 0x15, 0xc8, 0x77, 0x25, 0x25, 0xee, 0x2c,
	0x57, 0xf2, 0x7c, 0x4d, 0xef, 0xe1, 0x65, 0x38, 0xc7, 0x0d, 0x6a, 0xc3, 0x95, 0x6c, 0xd3, 0x9f,
	0x8b, 0x2c, 0x06, 0xf8, 0xa7, 0xba, 0x99, 0x9b, 0x89, 0x0f, 0xff, 0xf2, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x2f, 0x31, 0x35, 0xdb, 0xc9, 0x02, 0x00, 0x00,
}