// Code generated by protoc-gen-go.
// source: proto/pecho.proto
// DO NOT EDIT!

package pecho

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type PEchoRequest struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PEchoRequest) Reset()         { *m = PEchoRequest{} }
func (m *PEchoRequest) String() string { return proto.CompactTextString(m) }
func (*PEchoRequest) ProtoMessage()    {}

func (m *PEchoRequest) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type PEchoResponse struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PEchoResponse) Reset()         { *m = PEchoResponse{} }
func (m *PEchoResponse) String() string { return proto.CompactTextString(m) }
func (*PEchoResponse) ProtoMessage()    {}

func (m *PEchoResponse) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
}