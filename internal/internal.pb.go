// Code generated by protoc-gen-gogo.
// source: internal/internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Bitmap
	Chunk
	Pair
	QueryRequest
	QueryResponse
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Bitmap struct {
	Chunks           []*Chunk `protobuf:"bytes,1,rep" json:"Chunks,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Bitmap) Reset()         { *m = Bitmap{} }
func (m *Bitmap) String() string { return proto.CompactTextString(m) }
func (*Bitmap) ProtoMessage()    {}

func (m *Bitmap) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type Chunk struct {
	Key              *uint64  `protobuf:"varint,1,req" json:"Key,omitempty"`
	Value            []uint64 `protobuf:"varint,2,rep" json:"Value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}

func (m *Chunk) GetKey() uint64 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

func (m *Chunk) GetValue() []uint64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Pair struct {
	Key              *uint64 `protobuf:"varint,1,req" json:"Key,omitempty"`
	Count            *uint64 `protobuf:"varint,2,req" json:"Count,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}

func (m *Pair) GetKey() uint64 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

func (m *Pair) GetCount() uint64 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type QueryRequest struct {
	DB               *string  `protobuf:"bytes,1,req" json:"DB,omitempty"`
	Query            *string  `protobuf:"bytes,2,req" json:"Query,omitempty"`
	Slices           []uint64 `protobuf:"varint,3,rep" json:"Slices,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}

func (m *QueryRequest) GetDB() string {
	if m != nil && m.DB != nil {
		return *m.DB
	}
	return ""
}

func (m *QueryRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *QueryRequest) GetSlices() []uint64 {
	if m != nil {
		return m.Slices
	}
	return nil
}

type QueryResponse struct {
	Err              *string `protobuf:"bytes,1,opt" json:"Err,omitempty"`
	Bitmap           *Bitmap `protobuf:"bytes,2,opt" json:"Bitmap,omitempty"`
	N                *uint64 `protobuf:"varint,3,opt" json:"N,omitempty"`
	Pairs            []*Pair `protobuf:"bytes,4,rep" json:"Pairs,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}

func (m *QueryResponse) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

func (m *QueryResponse) GetBitmap() *Bitmap {
	if m != nil {
		return m.Bitmap
	}
	return nil
}

func (m *QueryResponse) GetN() uint64 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return 0
}

func (m *QueryResponse) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

func init() {
}