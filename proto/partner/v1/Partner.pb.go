// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/partner/v1/Partner.proto

package prixa_partner_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
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
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	AppIds               []string             `protobuf:"bytes,4,rep,name=appIds,proto3" json:"appIds,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *PartnerResponseData) GetAppIds() []string {
	if m != nil {
		return m.AppIds
	}
	return nil
}

func (m *PartnerResponseData) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PartnerResponseData) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
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

type PageProperty struct {
	PageNo               int32    `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	TotalPages           int32    `protobuf:"varint,2,opt,name=totalPages,proto3" json:"totalPages,omitempty"`
	TotalRecords         int32    `protobuf:"varint,3,opt,name=totalRecords,proto3" json:"totalRecords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageProperty) Reset()         { *m = PageProperty{} }
func (m *PageProperty) String() string { return proto.CompactTextString(m) }
func (*PageProperty) ProtoMessage()    {}
func (*PageProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{9}
}

func (m *PageProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageProperty.Unmarshal(m, b)
}
func (m *PageProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageProperty.Marshal(b, m, deterministic)
}
func (m *PageProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageProperty.Merge(m, src)
}
func (m *PageProperty) XXX_Size() int {
	return xxx_messageInfo_PageProperty.Size(m)
}
func (m *PageProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_PageProperty.DiscardUnknown(m)
}

var xxx_messageInfo_PageProperty proto.InternalMessageInfo

func (m *PageProperty) GetPageNo() int32 {
	if m != nil {
		return m.PageNo
	}
	return 0
}

func (m *PageProperty) GetTotalPages() int32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func (m *PageProperty) GetTotalRecords() int32 {
	if m != nil {
		return m.TotalRecords
	}
	return 0
}

type ListPartnersRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPartnersRequest) Reset()         { *m = ListPartnersRequest{} }
func (m *ListPartnersRequest) String() string { return proto.CompactTextString(m) }
func (*ListPartnersRequest) ProtoMessage()    {}
func (*ListPartnersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{10}
}

func (m *ListPartnersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPartnersRequest.Unmarshal(m, b)
}
func (m *ListPartnersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPartnersRequest.Marshal(b, m, deterministic)
}
func (m *ListPartnersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPartnersRequest.Merge(m, src)
}
func (m *ListPartnersRequest) XXX_Size() int {
	return xxx_messageInfo_ListPartnersRequest.Size(m)
}
func (m *ListPartnersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPartnersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPartnersRequest proto.InternalMessageInfo

func (m *ListPartnersRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ListPartnersResponse struct {
	Data                 []*PartnerResponseData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Page                 *PageProperty          `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListPartnersResponse) Reset()         { *m = ListPartnersResponse{} }
func (m *ListPartnersResponse) String() string { return proto.CompactTextString(m) }
func (*ListPartnersResponse) ProtoMessage()    {}
func (*ListPartnersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{11}
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

func (m *ListPartnersResponse) GetData() []*PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListPartnersResponse) GetPage() *PageProperty {
	if m != nil {
		return m.Page
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
	proto.RegisterType((*PageProperty)(nil), "prixa.partner.v1.PageProperty")
	proto.RegisterType((*ListPartnersRequest)(nil), "prixa.partner.v1.ListPartnersRequest")
	proto.RegisterType((*ListPartnersResponse)(nil), "prixa.partner.v1.ListPartnersResponse")
}

func init() { proto.RegisterFile("proto/partner/v1/Partner.proto", fileDescriptor_c6f7b43f209e465d) }

var fileDescriptor_c6f7b43f209e465d = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0x66, 0xaf, 0x49, 0x20, 0xd3, 0x56, 0xed, 0x36, 0x6a, 0x3c, 0x35, 0x0d, 0x57, 0x6d, 0x6b,
	0xa1, 0x77, 0x34, 0x82, 0xa8, 0x08, 0xa2, 0x16, 0xa4, 0x20, 0x5a, 0x4e, 0x85, 0x82, 0x4f, 0xdb,
	0xdc, 0x7a, 0x9e, 0x26, 0xb7, 0xe7, 0xed, 0xe6, 0xaa, 0x58, 0x5f, 0x84, 0x82, 0x8f, 0x82, 0x0f,
	0xfe, 0x30, 0x7f, 0x40, 0xa1, 0xf8, 0x23, 0x7c, 0x94, 0xdb, 0xdb, 0x24, 0xbd, 0xbb, 0x25, 0x55,
	0xf3, 0xb6, 0xbb, 0x33, 0xdf, 0xcc, 0x37, 0x5f, 0xe6, 0xbb, 0x40, 0x2b, 0x8a, 0x99, 0x60, 0x4e,
	0x44, 0x62, 0x11, 0xd2, 0xd8, 0x49, 0x36, 0x9d, 0x9d, 0xec, 0x68, 0xcb, 0x00, 0x3e, 0x17, 0xc5,
	0xc1, 0x07, 0x62, 0xab, 0xb8, 0x9d, 0x6c, 0x9a, 0x4b, 0x3e, 0x63, 0x7e, 0x8f, 0x3a, 0x32, 0xbe,
	0x37, 0x78, 0xed, 0x88, 0xa0, 0x4f, 0xb9, 0x20, 0xfd, 0x28, 0x83, 0x98, 0x57, 0x54, 0x02, 0x89,
	0x02, 0x87, 0x84, 0x21, 0x13, 0x44, 0x04, 0x2c, 0xe4, 0x2a, 0x7a, 0xcb, 0x0f, 0xc4, 0x9b, 0xc1,
	0x9e, 0xdd, 0x65, 0x7d, 0xa7, 0xbf, 0x1f, 0x88, 0x77, 0x6c, 0xdf, 0xf1, 0xd9, 0x86, 0x0c, 0x6e,
	0x24, 0xa4, 0x17, 0x78, 0x44, 0xb0, 0x98, 0x3b, 0xa3, 0x63, 0x86, 0xb3, 0x7e, 0x23, 0x58, 0x54,
	0xd4, 0x5c, 0xca, 0x23, 0x16, 0x72, 0xba, 0x45, 0x04, 0xc1, 0x97, 0xc0, 0x08, 0xbc, 0x26, 0x6a,
	0xa3, 0xb5, 0xfa, 0xc3, 0xfa, 0xf1, 0xd1, 0x52, 0x75, 0x17, 0x7d, 0x43, 0x15, 0xd7, 0x08, 0x3c,
	0x6c, 0x42, 0x25, 0x24, 0x7d, 0xda, 0x34, 0x64, 0xb0, 0x76, 0x7c, 0xb4, 0x64, 0xec, 0x22, 0x57,
	0xbe, 0xe1, 0x16, 0xd4, 0xb8, 0x20, 0x62, 0xc0, 0x9b, 0x33, 0xb9, 0xa8, 0x7a, 0xc5, 0x17, 0xa0,
	0x46, 0xa2, 0x68, 0xdb, 0xe3, 0xcd, 0x4a, 0x7b, 0x66, 0xad, 0xee, 0xaa, 0x1b, 0xbe, 0x0d, 0xf5,
	0x6e, 0x4c, 0x89, 0xa0, 0xde, 0x03, 0xd1, 0xac, 0xb6, 0xd1, 0xda, 0x6c, 0xc7, 0xb4, 0xb3, 0x81,
	0xed, 0xa1, 0x22, 0xf6, 0x8b, 0xa1, 0x22, 0xee, 0x38, 0x39, 0x45, 0x0e, 0x22, 0x4f, 0x21, 0x6b,
	0xa7, 0x23, 0x47, 0xc9, 0x56, 0x07, 0x1a, 0x8f, 0x64, 0x99, 0xd1, 0xfc, 0xef, 0x07, 0x94, 0x8b,
	0xd1, 0x7c, 0xa8, 0x3c, 0x9f, 0xe5, 0xc2, 0xf9, 0x02, 0x26, 0xd3, 0x0c, 0xdf, 0x81, 0x8a, 0x47,
	0x04, 0x91, 0xa0, 0xd9, 0xce, 0x75, 0xbb, 0xf8, 0xfb, 0xda, 0x1a, 0x91, 0x5d, 0x09, 0xb1, 0xee,
	0x43, 0x63, 0x8b, 0xf6, 0x68, 0x89, 0xc7, 0x2a, 0xd4, 0x15, 0x7e, 0x5b, 0xf3, 0x4b, 0x8c, 0x63,
	0x29, 0xa9, 0x42, 0x81, 0xe9, 0x49, 0xbd, 0x82, 0xc6, 0x4b, 0xa9, 0xd4, 0xdf, 0x8b, 0x93, 0x27,
	0x6c, 0x4c, 0x26, 0x5c, 0x28, 0x3e, 0x3d, 0xe1, 0x7b, 0xb0, 0xf0, 0x98, 0x8a, 0xff, 0x95, 0xf0,
	0x19, 0xe0, 0x93, 0xe8, 0xe9, 0xe9, 0xbc, 0x85, 0xb9, 0x1d, 0xe2, 0xd3, 0x9d, 0x98, 0x45, 0x34,
	0x16, 0x1f, 0xd3, 0xc5, 0x8f, 0x88, 0x4f, 0x9f, 0x32, 0x59, 0xac, 0xea, 0xaa, 0x1b, 0x6e, 0x01,
	0x08, 0x26, 0x48, 0x2f, 0x4d, 0xe6, 0x52, 0xb4, 0xaa, 0x7b, 0xe2, 0x05, 0x5b, 0x30, 0x27, 0x6f,
	0x2e, 0xed, 0xb2, 0xd8, 0xcb, 0x6c, 0x55, 0x75, 0x73, 0x6f, 0xd6, 0x0d, 0x58, 0x7c, 0x12, 0xf0,
	0x21, 0x7b, 0x3e, 0x1c, 0x1e, 0x43, 0x25, 0x6d, 0xa2, 0x1a, 0xca, 0xb3, 0x75, 0x88, 0xa0, 0x91,
	0xcf, 0x2d, 0x8d, 0x3a, 0xf3, 0x8f, 0xa3, 0xe2, 0x8e, 0xea, 0x63, 0x48, 0x95, 0x5a, 0x3a, 0xe8,
	0x58, 0x88, 0x8c, 0x47, 0xe7, 0x47, 0x15, 0xce, 0xa8, 0x8a, 0xcf, 0x69, 0x9c, 0x04, 0x5d, 0x8a,
	0x0f, 0x60, 0x3e, 0x67, 0x2d, 0xbc, 0x52, 0xae, 0xa4, 0xf3, 0xab, 0xb9, 0x7a, 0x6a, 0x5e, 0x46,
	0xd9, 0x32, 0xbf, 0xfc, 0xfc, 0xf5, 0xdd, 0x68, 0x58, 0x67, 0xe5, 0x37, 0x34, 0xd9, 0x1c, 0x7e,
	0x9e, 0xef, 0xa2, 0x75, 0x7c, 0x88, 0x60, 0x3e, 0x67, 0x22, 0x5d, 0x7b, 0x9d, 0x4d, 0x75, 0xed,
	0xb5, 0x6e, 0xb4, 0x96, 0x65, 0xfb, 0xab, 0xeb, 0x97, 0x0b, 0xed, 0x9d, 0x4f, 0xa3, 0x3d, 0xfc,
	0x8c, 0xbf, 0x22, 0x98, 0xcf, 0x79, 0x43, 0xc7, 0x43, 0xe7, 0x4c, 0x1d, 0x0f, 0xad, 0xc9, 0xac,
	0x15, 0xc9, 0xa3, 0x6d, 0x4e, 0xe2, 0x91, 0x4a, 0x72, 0x00, 0x30, 0xf6, 0x04, 0x5e, 0x2e, 0x97,
	0x2f, 0xf9, 0xcd, 0xbc, 0x36, 0x39, 0x29, 0x2f, 0x04, 0x9e, 0x28, 0x44, 0x02, 0x73, 0x27, 0x17,
	0x15, 0x6b, 0x56, 0x52, 0xb3, 0xf4, 0xe6, 0xca, 0x69, 0x69, 0x8a, 0xc3, 0x45, 0xc9, 0x61, 0x01,
	0x17, 0x77, 0x61, 0xaf, 0x26, 0xff, 0x34, 0x6e, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x48, 0x0b,
	0x62, 0xf1, 0xc2, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PartnerServiceClient is the client API for PartnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PartnerServiceClient interface {
	CreatePartner(ctx context.Context, in *CreatePartnerRequest, opts ...grpc.CallOption) (*CreatePartnerResponse, error)
	DeletePartner(ctx context.Context, in *DeletePartnerRequest, opts ...grpc.CallOption) (*DeletePartnerResponse, error)
	UpdatePartner(ctx context.Context, in *UpdatePartnerRequest, opts ...grpc.CallOption) (*UpdatePartnerResponse, error)
	GetPartner(ctx context.Context, in *GetPartnerRequest, opts ...grpc.CallOption) (*GetPartnerResponse, error)
	ListPartners(ctx context.Context, in *ListPartnersRequest, opts ...grpc.CallOption) (*ListPartnersResponse, error)
}

type partnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerServiceClient(cc grpc.ClientConnInterface) PartnerServiceClient {
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

func (c *partnerServiceClient) ListPartners(ctx context.Context, in *ListPartnersRequest, opts ...grpc.CallOption) (*ListPartnersResponse, error) {
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
	ListPartners(context.Context, *ListPartnersRequest) (*ListPartnersResponse, error)
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
func (*UnimplementedPartnerServiceServer) ListPartners(ctx context.Context, req *ListPartnersRequest) (*ListPartnersResponse, error) {
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
	in := new(ListPartnersRequest)
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
		return srv.(PartnerServiceServer).ListPartners(ctx, req.(*ListPartnersRequest))
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
