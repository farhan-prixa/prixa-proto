// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/partner/v1/Partner.proto

package prixa_partner_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	v1 "github.com/prixa-ai/prixa-proto/proto/support/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PartnerResponseData struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string             `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	AppIds               *_struct.ListValue `protobuf:"bytes,4,opt,name=appIds,proto3" json:"appIds,omitempty"`
	DataTimestamp        *v1.DataTimestamp  `protobuf:"bytes,5,opt,name=dataTimestamp,proto3" json:"dataTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PartnerResponseData) Reset()         { *m = PartnerResponseData{} }
func (m *PartnerResponseData) String() string { return proto.CompactTextString(m) }
func (*PartnerResponseData) ProtoMessage()    {}
func (*PartnerResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{0}
}

func (m *PartnerResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartnerResponseData.Unmarshal(m, b)
}
func (m *PartnerResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartnerResponseData.Marshal(b, m, deterministic)
}
func (m *PartnerResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartnerResponseData.Merge(m, src)
}
func (m *PartnerResponseData) XXX_Size() int {
	return xxx_messageInfo_PartnerResponseData.Size(m)
}
func (m *PartnerResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_PartnerResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_PartnerResponseData proto.InternalMessageInfo

func (m *PartnerResponseData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PartnerResponseData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PartnerResponseData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PartnerResponseData) GetAppIds() *_struct.ListValue {
	if m != nil {
		return m.AppIds
	}
	return nil
}

func (m *PartnerResponseData) GetDataTimestamp() *v1.DataTimestamp {
	if m != nil {
		return m.DataTimestamp
	}
	return nil
}

type CreatePartnerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePartnerRequest) Reset()         { *m = CreatePartnerRequest{} }
func (m *CreatePartnerRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePartnerRequest) ProtoMessage()    {}
func (*CreatePartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{1}
}

func (m *CreatePartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePartnerRequest.Unmarshal(m, b)
}
func (m *CreatePartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePartnerRequest.Marshal(b, m, deterministic)
}
func (m *CreatePartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePartnerRequest.Merge(m, src)
}
func (m *CreatePartnerRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePartnerRequest.Size(m)
}
func (m *CreatePartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePartnerRequest proto.InternalMessageInfo

func (m *CreatePartnerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreatePartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreatePartnerResponse) Reset()         { *m = CreatePartnerResponse{} }
func (m *CreatePartnerResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePartnerResponse) ProtoMessage()    {}
func (*CreatePartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{2}
}

func (m *CreatePartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePartnerResponse.Unmarshal(m, b)
}
func (m *CreatePartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePartnerResponse.Marshal(b, m, deterministic)
}
func (m *CreatePartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePartnerResponse.Merge(m, src)
}
func (m *CreatePartnerResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePartnerResponse.Size(m)
}
func (m *CreatePartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePartnerResponse proto.InternalMessageInfo

func (m *CreatePartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeletePartnerRequest struct {
	PartnerId            string   `protobuf:"bytes,1,opt,name=partnerId,proto3" json:"partnerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePartnerRequest) Reset()         { *m = DeletePartnerRequest{} }
func (m *DeletePartnerRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePartnerRequest) ProtoMessage()    {}
func (*DeletePartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{3}
}

func (m *DeletePartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePartnerRequest.Unmarshal(m, b)
}
func (m *DeletePartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePartnerRequest.Marshal(b, m, deterministic)
}
func (m *DeletePartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePartnerRequest.Merge(m, src)
}
func (m *DeletePartnerRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePartnerRequest.Size(m)
}
func (m *DeletePartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePartnerRequest proto.InternalMessageInfo

func (m *DeletePartnerRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type DeletePartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeletePartnerResponse) Reset()         { *m = DeletePartnerResponse{} }
func (m *DeletePartnerResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePartnerResponse) ProtoMessage()    {}
func (*DeletePartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{4}
}

func (m *DeletePartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePartnerResponse.Unmarshal(m, b)
}
func (m *DeletePartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePartnerResponse.Marshal(b, m, deterministic)
}
func (m *DeletePartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePartnerResponse.Merge(m, src)
}
func (m *DeletePartnerResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePartnerResponse.Size(m)
}
func (m *DeletePartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePartnerResponse proto.InternalMessageInfo

func (m *DeletePartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdatePartnerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PartnerId            string   `protobuf:"bytes,2,opt,name=partnerId,proto3" json:"partnerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePartnerRequest) Reset()         { *m = UpdatePartnerRequest{} }
func (m *UpdatePartnerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePartnerRequest) ProtoMessage()    {}
func (*UpdatePartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{5}
}

func (m *UpdatePartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePartnerRequest.Unmarshal(m, b)
}
func (m *UpdatePartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePartnerRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePartnerRequest.Merge(m, src)
}
func (m *UpdatePartnerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePartnerRequest.Size(m)
}
func (m *UpdatePartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePartnerRequest proto.InternalMessageInfo

func (m *UpdatePartnerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePartnerRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type UpdatePartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdatePartnerResponse) Reset()         { *m = UpdatePartnerResponse{} }
func (m *UpdatePartnerResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePartnerResponse) ProtoMessage()    {}
func (*UpdatePartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{6}
}

func (m *UpdatePartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePartnerResponse.Unmarshal(m, b)
}
func (m *UpdatePartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePartnerResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePartnerResponse.Merge(m, src)
}
func (m *UpdatePartnerResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePartnerResponse.Size(m)
}
func (m *UpdatePartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePartnerResponse proto.InternalMessageInfo

func (m *UpdatePartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetPartnerRequest struct {
	PartnerId            string   `protobuf:"bytes,1,opt,name=partnerId,proto3" json:"partnerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPartnerRequest) Reset()         { *m = GetPartnerRequest{} }
func (m *GetPartnerRequest) String() string { return proto.CompactTextString(m) }
func (*GetPartnerRequest) ProtoMessage()    {}
func (*GetPartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{7}
}

func (m *GetPartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPartnerRequest.Unmarshal(m, b)
}
func (m *GetPartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPartnerRequest.Marshal(b, m, deterministic)
}
func (m *GetPartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPartnerRequest.Merge(m, src)
}
func (m *GetPartnerRequest) XXX_Size() int {
	return xxx_messageInfo_GetPartnerRequest.Size(m)
}
func (m *GetPartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPartnerRequest proto.InternalMessageInfo

func (m *GetPartnerRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type GetPartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetPartnerResponse) Reset()         { *m = GetPartnerResponse{} }
func (m *GetPartnerResponse) String() string { return proto.CompactTextString(m) }
func (*GetPartnerResponse) ProtoMessage()    {}
func (*GetPartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{8}
}

func (m *GetPartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPartnerResponse.Unmarshal(m, b)
}
func (m *GetPartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPartnerResponse.Marshal(b, m, deterministic)
}
func (m *GetPartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPartnerResponse.Merge(m, src)
}
func (m *GetPartnerResponse) XXX_Size() int {
	return xxx_messageInfo_GetPartnerResponse.Size(m)
}
func (m *GetPartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPartnerResponse proto.InternalMessageInfo

func (m *GetPartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListPartnersResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListPartnersResponse) Reset()         { *m = ListPartnersResponse{} }
func (m *ListPartnersResponse) String() string { return proto.CompactTextString(m) }
func (*ListPartnersResponse) ProtoMessage()    {}
func (*ListPartnersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{9}
}

func (m *ListPartnersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPartnersResponse.Unmarshal(m, b)
}
func (m *ListPartnersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPartnersResponse.Marshal(b, m, deterministic)
}
func (m *ListPartnersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPartnersResponse.Merge(m, src)
}
func (m *ListPartnersResponse) XXX_Size() int {
	return xxx_messageInfo_ListPartnersResponse.Size(m)
}
func (m *ListPartnersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPartnersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPartnersResponse proto.InternalMessageInfo

func (m *ListPartnersResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PartnerResponseData)(nil), "prixa.partner.v1.PartnerResponseData")
	proto.RegisterType((*CreatePartnerRequest)(nil), "prixa.partner.v1.CreatePartnerRequest")
	proto.RegisterType((*CreatePartnerResponse)(nil), "prixa.partner.v1.CreatePartnerResponse")
	proto.RegisterType((*DeletePartnerRequest)(nil), "prixa.partner.v1.DeletePartnerRequest")
	proto.RegisterType((*DeletePartnerResponse)(nil), "prixa.partner.v1.DeletePartnerResponse")
	proto.RegisterType((*UpdatePartnerRequest)(nil), "prixa.partner.v1.UpdatePartnerRequest")
	proto.RegisterType((*UpdatePartnerResponse)(nil), "prixa.partner.v1.UpdatePartnerResponse")
	proto.RegisterType((*GetPartnerRequest)(nil), "prixa.partner.v1.GetPartnerRequest")
	proto.RegisterType((*GetPartnerResponse)(nil), "prixa.partner.v1.GetPartnerResponse")
	proto.RegisterType((*ListPartnersResponse)(nil), "prixa.partner.v1.ListPartnersResponse")
}

func init() { proto.RegisterFile("proto/partner/v1/Partner.proto", fileDescriptor_c6f7b43f209e465d) }

var fileDescriptor_c6f7b43f209e465d = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0xe5, 0x90, 0x46, 0xea, 0x29, 0x29, 0x74, 0x30, 0xc5, 0x72, 0x0b, 0x0a, 0x2e, 0x94,
	0xaa, 0x8b, 0xb1, 0x1c, 0xd8, 0xc0, 0x96, 0xa0, 0x52, 0x09, 0x09, 0x30, 0x97, 0xfd, 0xb4, 0x1e,
	0x22, 0x4b, 0x89, 0x3d, 0x78, 0x8e, 0x2d, 0x10, 0x65, 0xc3, 0x0a, 0x89, 0x25, 0xef, 0xc5, 0x86,
	0x57, 0xe0, 0x41, 0x90, 0xc7, 0x97, 0xc4, 0xf6, 0xa8, 0x05, 0xa5, 0x3b, 0x67, 0xce, 0xe5, 0xff,
	0xce, 0x99, 0x3f, 0x36, 0xdc, 0x11, 0x49, 0x8c, 0xb1, 0x2b, 0x58, 0x82, 0x11, 0x4f, 0xdc, 0xcc,
	0x73, 0x5f, 0x15, 0x8f, 0x54, 0x05, 0xc8, 0x75, 0x91, 0x84, 0x9f, 0x18, 0x2d, 0xe3, 0x34, 0xf3,
	0xec, 0x9d, 0x69, 0x1c, 0x4f, 0x67, 0xdc, 0x55, 0xf1, 0x93, 0xf4, 0x83, 0xcb, 0xe7, 0x02, 0x3f,
	0x17, 0xe9, 0xf6, 0x6e, 0x3b, 0x28, 0x31, 0x49, 0x4f, 0xb1, 0x8c, 0x8e, 0x0a, 0x31, 0x99, 0x0a,
	0x11, 0x27, 0x98, 0x8b, 0xbd, 0x0d, 0xe7, 0x5c, 0x22, 0x9b, 0x8b, 0x56, 0x3d, 0x13, 0xa1, 0xcb,
	0xa2, 0x28, 0x46, 0x86, 0x61, 0x1c, 0xc9, 0x22, 0xea, 0xfc, 0x32, 0xe0, 0x46, 0x89, 0xe7, 0x73,
	0x29, 0xe2, 0x48, 0xf2, 0x09, 0x43, 0x46, 0x36, 0xa1, 0x17, 0x06, 0x96, 0x31, 0x32, 0x0e, 0xd6,
	0xfd, 0x5e, 0x18, 0x10, 0x02, 0xfd, 0x88, 0xcd, 0xb9, 0xd5, 0x53, 0x27, 0xea, 0x99, 0x6c, 0xc3,
	0x40, 0x22, 0xc3, 0x54, 0x5a, 0x57, 0xd4, 0x69, 0xf9, 0x8b, 0x8c, 0x61, 0xc0, 0x84, 0x38, 0x0e,
	0xa4, 0xd5, 0x1f, 0x19, 0x07, 0x1b, 0x63, 0x9b, 0x16, 0x08, 0xb4, 0x1a, 0x81, 0xbe, 0x08, 0x25,
	0xbe, 0x67, 0xb3, 0x94, 0xfb, 0x65, 0x26, 0x39, 0x82, 0x61, 0xc0, 0x90, 0xd5, 0xf0, 0xd6, 0x9a,
	0x2a, 0xbd, 0x4b, 0x8b, 0x65, 0x61, 0x3d, 0x54, 0xe6, 0xd1, 0xc9, 0x72, 0xa2, 0xdf, 0xac, 0x73,
	0x0e, 0xc1, 0x7c, 0x9a, 0x70, 0x86, 0xbc, 0x9e, 0xea, 0x63, 0xca, 0x25, 0xd6, 0x03, 0x18, 0x8b,
	0x01, 0x1c, 0x1f, 0x6e, 0xb6, 0x72, 0x8b, 0x0d, 0x90, 0xc7, 0xd0, 0xcf, 0xbb, 0xaa, 0xe4, 0x8d,
	0xf1, 0x7d, 0xda, 0xbe, 0x31, 0xaa, 0x59, 0x99, 0xaf, 0x4a, 0x9c, 0x47, 0x60, 0x4e, 0xf8, 0x8c,
	0x77, 0xf4, 0x77, 0x61, 0xbd, 0xac, 0x3f, 0xae, 0xf6, 0xba, 0x38, 0xc8, 0x49, 0x5a, 0x55, 0xab,
	0x93, 0x3c, 0x07, 0xf3, 0x9d, 0x08, 0xfe, 0x69, 0x13, 0x4d, 0xba, 0x9e, 0x86, 0xae, 0xd5, 0x69,
	0x75, 0x3a, 0x0f, 0xb6, 0x8e, 0x38, 0xfe, 0xd7, 0x92, 0x5e, 0x02, 0x59, 0x2e, 0x59, 0x9d, 0xe1,
	0x35, 0x98, 0xb9, 0x13, 0xcb, 0x04, 0x79, 0x09, 0x2d, 0xc7, 0x3f, 0xd6, 0x60, 0xb3, 0x8c, 0xbe,
	0xe1, 0x49, 0x16, 0x9e, 0x72, 0x72, 0x06, 0xc3, 0x86, 0xcb, 0xc8, 0x7e, 0xb7, 0xa1, 0xce, 0xb2,
	0xf6, 0x83, 0x0b, 0xf3, 0x0a, 0x79, 0xc7, 0xfe, 0xf6, 0xfb, 0xcf, 0xcf, 0x9e, 0xe9, 0x5c, 0x53,
	0x7f, 0xf2, 0xcc, 0xab, 0xde, 0x3d, 0x4f, 0x8c, 0x43, 0xf2, 0xdd, 0x80, 0x61, 0xc3, 0x5a, 0x3a,
	0x79, 0x9d, 0x63, 0x75, 0xf2, 0x5a, 0x8f, 0x3a, 0xfb, 0x4a, 0x7e, 0xe4, 0xec, 0xb4, 0xe4, 0xdd,
	0x2f, 0xf5, 0xdd, 0x7d, 0xad, 0x50, 0x1a, 0x3e, 0xd2, 0xa1, 0xe8, 0x2c, 0xab, 0x43, 0xd1, 0x1a,
	0xb2, 0x42, 0xb1, 0x2f, 0x42, 0x39, 0x03, 0x58, 0x58, 0x89, 0xec, 0x75, 0xdb, 0x77, 0xbc, 0x69,
	0xdf, 0x3b, 0x3f, 0xa9, 0x04, 0xd8, 0x53, 0x00, 0xb7, 0xc9, 0x79, 0x00, 0x64, 0x0a, 0x57, 0x97,
	0x7d, 0x47, 0xb6, 0x3b, 0x2f, 0xc8, 0x67, 0xf9, 0x07, 0xc0, 0xd6, 0xac, 0x47, 0xe7, 0x57, 0xe7,
	0x96, 0x12, 0xdd, 0x22, 0xed, 0xfb, 0x3f, 0x19, 0xa8, 0x86, 0x0f, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xf7, 0x52, 0xf5, 0xaf, 0x93, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PartnerServiceClient is the client API for PartnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PartnerServiceClient interface {
	CreatePartner(ctx context.Context, in *CreatePartnerRequest, opts ...grpc.CallOption) (*CreatePartnerResponse, error)
	DeletePartner(ctx context.Context, in *DeletePartnerRequest, opts ...grpc.CallOption) (*DeletePartnerResponse, error)
	UpdatePartner(ctx context.Context, in *UpdatePartnerRequest, opts ...grpc.CallOption) (*UpdatePartnerResponse, error)
	GetPartner(ctx context.Context, in *GetPartnerRequest, opts ...grpc.CallOption) (*GetPartnerResponse, error)
	ListPartners(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListPartnersResponse, error)
}

type partnerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPartnerServiceClient(cc *grpc.ClientConn) PartnerServiceClient {
	return &partnerServiceClient{cc}
}

func (c *partnerServiceClient) CreatePartner(ctx context.Context, in *CreatePartnerRequest, opts ...grpc.CallOption) (*CreatePartnerResponse, error) {
	out := new(CreatePartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/CreatePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) DeletePartner(ctx context.Context, in *DeletePartnerRequest, opts ...grpc.CallOption) (*DeletePartnerResponse, error) {
	out := new(DeletePartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/DeletePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) UpdatePartner(ctx context.Context, in *UpdatePartnerRequest, opts ...grpc.CallOption) (*UpdatePartnerResponse, error) {
	out := new(UpdatePartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/UpdatePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) GetPartner(ctx context.Context, in *GetPartnerRequest, opts ...grpc.CallOption) (*GetPartnerResponse, error) {
	out := new(GetPartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/GetPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) ListPartners(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListPartnersResponse, error) {
	out := new(ListPartnersResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/ListPartners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerServiceServer is the server API for PartnerService service.
type PartnerServiceServer interface {
	CreatePartner(context.Context, *CreatePartnerRequest) (*CreatePartnerResponse, error)
	DeletePartner(context.Context, *DeletePartnerRequest) (*DeletePartnerResponse, error)
	UpdatePartner(context.Context, *UpdatePartnerRequest) (*UpdatePartnerResponse, error)
	GetPartner(context.Context, *GetPartnerRequest) (*GetPartnerResponse, error)
	ListPartners(context.Context, *empty.Empty) (*ListPartnersResponse, error)
}

// UnimplementedPartnerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPartnerServiceServer struct {
}

func (*UnimplementedPartnerServiceServer) CreatePartner(ctx context.Context, req *CreatePartnerRequest) (*CreatePartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartner not implemented")
}
func (*UnimplementedPartnerServiceServer) DeletePartner(ctx context.Context, req *DeletePartnerRequest) (*DeletePartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePartner not implemented")
}
func (*UnimplementedPartnerServiceServer) UpdatePartner(ctx context.Context, req *UpdatePartnerRequest) (*UpdatePartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartner not implemented")
}
func (*UnimplementedPartnerServiceServer) GetPartner(ctx context.Context, req *GetPartnerRequest) (*GetPartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartner not implemented")
}
func (*UnimplementedPartnerServiceServer) ListPartners(ctx context.Context, req *empty.Empty) (*ListPartnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPartners not implemented")
}

func RegisterPartnerServiceServer(s *grpc.Server, srv PartnerServiceServer) {
	s.RegisterService(&_PartnerService_serviceDesc, srv)
}

func _PartnerService_CreatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).CreatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/CreatePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).CreatePartner(ctx, req.(*CreatePartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_DeletePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).DeletePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/DeletePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).DeletePartner(ctx, req.(*DeletePartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_UpdatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).UpdatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/UpdatePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).UpdatePartner(ctx, req.(*UpdatePartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_GetPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).GetPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/GetPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).GetPartner(ctx, req.(*GetPartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_ListPartners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).ListPartners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/ListPartners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).ListPartners(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PartnerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prixa.partner.v1.PartnerService",
	HandlerType: (*PartnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePartner",
			Handler:    _PartnerService_CreatePartner_Handler,
		},
		{
			MethodName: "DeletePartner",
			Handler:    _PartnerService_DeletePartner_Handler,
		},
		{
			MethodName: "UpdatePartner",
			Handler:    _PartnerService_UpdatePartner_Handler,
		},
		{
			MethodName: "GetPartner",
			Handler:    _PartnerService_GetPartner_Handler,
		},
		{
			MethodName: "ListPartners",
			Handler:    _PartnerService_ListPartners_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/partner/v1/Partner.proto",
}
