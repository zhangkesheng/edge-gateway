// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/account.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type InfoResponse struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string                 `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Clients              []*InfoResponse_Client `protobuf:"bytes,3,rep,name=clients,proto3" json:"clients,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{0}
}

func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *InfoResponse) GetClients() []*InfoResponse_Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

type InfoResponse_Client struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResponse_Client) Reset()         { *m = InfoResponse_Client{} }
func (m *InfoResponse_Client) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_Client) ProtoMessage()    {}
func (*InfoResponse_Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{0, 0}
}

func (m *InfoResponse_Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse_Client.Unmarshal(m, b)
}
func (m *InfoResponse_Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse_Client.Marshal(b, m, deterministic)
}
func (m *InfoResponse_Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse_Client.Merge(m, src)
}
func (m *InfoResponse_Client) XXX_Size() int {
	return xxx_messageInfo_InfoResponse_Client.Size(m)
}
func (m *InfoResponse_Client) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse_Client.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse_Client proto.InternalMessageInfo

func (m *InfoResponse_Client) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InfoResponse_Client) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type LoginRequest struct {
	ResponseType         string   `protobuf:"bytes,1,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	RedirectUrl          string   `protobuf:"bytes,3,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	Scope                string   `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{1}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetResponseType() string {
	if m != nil {
		return m.ResponseType
	}
	return ""
}

func (m *LoginRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *LoginRequest) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

func (m *LoginRequest) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *LoginRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type LoginResponse struct {
	RedirectTo           string   `protobuf:"bytes,1,opt,name=redirect_to,json=redirectTo,proto3" json:"redirect_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{2}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetRedirectTo() string {
	if m != nil {
		return m.RedirectTo
	}
	return ""
}

type CallbackRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackRequest) Reset()         { *m = CallbackRequest{} }
func (m *CallbackRequest) String() string { return proto.CompactTextString(m) }
func (*CallbackRequest) ProtoMessage()    {}
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{3}
}

func (m *CallbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackRequest.Unmarshal(m, b)
}
func (m *CallbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackRequest.Marshal(b, m, deterministic)
}
func (m *CallbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackRequest.Merge(m, src)
}
func (m *CallbackRequest) XXX_Size() int {
	return xxx_messageInfo_CallbackRequest.Size(m)
}
func (m *CallbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackRequest proto.InternalMessageInfo

func (m *CallbackRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CallbackRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type CallbackResponse struct {
	RedirectUrl          string   `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackResponse) Reset()         { *m = CallbackResponse{} }
func (m *CallbackResponse) String() string { return proto.CompactTextString(m) }
func (*CallbackResponse) ProtoMessage()    {}
func (*CallbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{4}
}

func (m *CallbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackResponse.Unmarshal(m, b)
}
func (m *CallbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackResponse.Marshal(b, m, deterministic)
}
func (m *CallbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackResponse.Merge(m, src)
}
func (m *CallbackResponse) XXX_Size() int {
	return xxx_messageInfo_CallbackResponse.Size(m)
}
func (m *CallbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackResponse proto.InternalMessageInfo

func (m *CallbackResponse) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

type TokenRequest struct {
	GrantType            string   `protobuf:"bytes,1,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	ClientId             string   `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenRequest) Reset()         { *m = TokenRequest{} }
func (m *TokenRequest) String() string { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()    {}
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{5}
}

func (m *TokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenRequest.Unmarshal(m, b)
}
func (m *TokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenRequest.Marshal(b, m, deterministic)
}
func (m *TokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenRequest.Merge(m, src)
}
func (m *TokenRequest) XXX_Size() int {
	return xxx_messageInfo_TokenRequest.Size(m)
}
func (m *TokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenRequest proto.InternalMessageInfo

func (m *TokenRequest) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

func (m *TokenRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *TokenRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TokenRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type TokenResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	IdToken              string   `protobuf:"bytes,3,opt,name=id_token,json=idToken,proto3" json:"id_token,omitempty"`
	ExpiresIn            int64    `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenResponse) Reset()         { *m = TokenResponse{} }
func (m *TokenResponse) String() string { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()    {}
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{6}
}

func (m *TokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenResponse.Unmarshal(m, b)
}
func (m *TokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenResponse.Marshal(b, m, deterministic)
}
func (m *TokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenResponse.Merge(m, src)
}
func (m *TokenResponse) XXX_Size() int {
	return xxx_messageInfo_TokenResponse.Size(m)
}
func (m *TokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenResponse proto.InternalMessageInfo

func (m *TokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *TokenResponse) GetIdToken() string {
	if m != nil {
		return m.IdToken
	}
	return ""
}

func (m *TokenResponse) GetExpiresIn() int64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

type RefreshRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{7}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type RefreshResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshResponse) Reset()         { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()    {}
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{8}
}

func (m *RefreshResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshResponse.Unmarshal(m, b)
}
func (m *RefreshResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshResponse.Marshal(b, m, deterministic)
}
func (m *RefreshResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshResponse.Merge(m, src)
}
func (m *RefreshResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshResponse.Size(m)
}
func (m *RefreshResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshResponse proto.InternalMessageInfo

type VerifyRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{9}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyResponse struct {
	Sub                  string   `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{10}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetSub() string {
	if m != nil {
		return m.Sub
	}
	return ""
}

type LogoutRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{11}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e8030af56250e93, []int{12}
}

func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (m *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(m, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InfoResponse)(nil), "edge.account.v1.InfoResponse")
	proto.RegisterType((*InfoResponse_Client)(nil), "edge.account.v1.InfoResponse.Client")
	proto.RegisterType((*LoginRequest)(nil), "edge.account.v1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "edge.account.v1.LoginResponse")
	proto.RegisterType((*CallbackRequest)(nil), "edge.account.v1.CallbackRequest")
	proto.RegisterType((*CallbackResponse)(nil), "edge.account.v1.CallbackResponse")
	proto.RegisterType((*TokenRequest)(nil), "edge.account.v1.TokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "edge.account.v1.TokenResponse")
	proto.RegisterType((*RefreshRequest)(nil), "edge.account.v1.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "edge.account.v1.RefreshResponse")
	proto.RegisterType((*VerifyRequest)(nil), "edge.account.v1.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "edge.account.v1.VerifyResponse")
	proto.RegisterType((*LogoutRequest)(nil), "edge.account.v1.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "edge.account.v1.LogoutResponse")
}

func init() { proto.RegisterFile("v1/account.proto", fileDescriptor_9e8030af56250e93) }

var fileDescriptor_9e8030af56250e93 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x55, 0x6a, 0xa7, 0x69, 0x6e, 0xf3, 0x62, 0x54, 0x21, 0x63, 0x14, 0x92, 0x1a, 0x90, 0xba,
	0x72, 0x69, 0x2b, 0x16, 0x08, 0x81, 0x04, 0x15, 0x48, 0x91, 0x22, 0x21, 0x59, 0x81, 0x05, 0x1b,
	0xcb, 0xb1, 0x27, 0xc6, 0xaa, 0xeb, 0x31, 0xf6, 0xb8, 0x22, 0x1f, 0xc1, 0x2f, 0xf0, 0x0f, 0xfc,
	0x21, 0x9a, 0x97, 0xed, 0xc4, 0x49, 0xd8, 0xcd, 0x9c, 0x7b, 0xe6, 0x9c, 0x3b, 0xf7, 0x01, 0xa3,
	0x87, 0xab, 0x4b, 0xcf, 0xf7, 0x49, 0x91, 0x50, 0x3b, 0xcd, 0x08, 0x25, 0x68, 0x88, 0x83, 0x10,
	0xdb, 0x0a, 0x7b, 0xb8, 0x32, 0x51, 0x48, 0x48, 0x18, 0xe3, 0x4b, 0x7c, 0x9f, 0xd2, 0xb5, 0x20,
	0x59, 0x7f, 0x5b, 0xd0, 0x9b, 0x25, 0x2b, 0xe2, 0xe0, 0x3c, 0x25, 0x49, 0x8e, 0x11, 0x02, 0x3d,
	0xf1, 0xee, 0xb1, 0xd1, 0x9a, 0xb6, 0x2e, 0xba, 0x0e, 0x3f, 0x33, 0x2c, 0xc0, 0xb9, 0x6f, 0x1c,
	0x09, 0x8c, 0x9d, 0xd1, 0x7b, 0xe8, 0xf8, 0x71, 0x84, 0x13, 0x9a, 0x1b, 0xda, 0x54, 0xbb, 0x38,
	0xbd, 0x7e, 0x61, 0x6f, 0xf9, 0xd9, 0x75, 0x5d, 0xfb, 0x96, 0x93, 0x1d, 0xf5, 0xc8, 0x7c, 0x03,
	0xc7, 0x02, 0x62, 0xea, 0x74, 0x9d, 0x96, 0x8e, 0xec, 0x8c, 0x9e, 0x42, 0x57, 0x10, 0xdd, 0x28,
	0x90, 0xb6, 0x27, 0x02, 0x98, 0x05, 0xd6, 0x9f, 0x16, 0xf4, 0xe6, 0x24, 0x8c, 0x12, 0x07, 0xff,
	0x2c, 0x70, 0x4e, 0xd1, 0x73, 0xe8, 0x67, 0xd2, 0xc7, 0xad, 0x49, 0xf5, 0x14, 0xb8, 0x60, 0x92,
	0x26, 0x94, 0x0a, 0xdb, 0x8a, 0xe8, 0x1c, 0x7a, 0x19, 0x0e, 0xa2, 0x0c, 0xfb, 0xd4, 0x2d, 0xb2,
	0xd8, 0xd0, 0x78, 0xfc, 0x54, 0x61, 0x5f, 0xb3, 0x18, 0x9d, 0x41, 0x3b, 0xf7, 0x49, 0x8a, 0x0d,
	0x9d, 0xc7, 0xc4, 0x85, 0xa3, 0xd4, 0xa3, 0xd8, 0x68, 0x4b, 0x94, 0x5d, 0xac, 0x57, 0xd0, 0x97,
	0xf9, 0xc9, 0xa2, 0x4e, 0xa0, 0xd4, 0x72, 0x29, 0x91, 0xe9, 0x81, 0x82, 0x16, 0xc4, 0x7a, 0x0b,
	0xc3, 0x5b, 0x2f, 0x8e, 0x97, 0x9e, 0x7f, 0xa7, 0x3e, 0x85, 0x40, 0xf7, 0x49, 0x50, 0x96, 0x85,
	0x9d, 0x2b, 0xbb, 0xa3, 0xba, 0xdd, 0x6b, 0x18, 0x55, 0x8f, 0xa5, 0xe3, 0xf6, 0x8f, 0x5a, 0x8d,
	0x1f, 0x59, 0x14, 0x7a, 0x0b, 0x72, 0x87, 0xcb, 0x2a, 0x8e, 0x01, 0xc2, 0xcc, 0x4b, 0x68, 0xbd,
	0x84, 0x5d, 0x8e, 0xf0, 0xfa, 0xa9, 0x7c, 0x8e, 0x76, 0xe5, 0xa3, 0xd5, 0xf2, 0xd9, 0x6c, 0x9e,
	0xbe, 0xd5, 0xbc, 0x18, 0xfa, 0xd2, 0xb5, 0xca, 0xd4, 0xf3, 0x7d, 0x9c, 0xe7, 0x2e, 0x65, 0xb8,
	0xca, 0x54, 0x60, 0x9c, 0x8a, 0x9e, 0xc0, 0x49, 0x14, 0xc8, 0xb0, 0x70, 0xea, 0x44, 0x81, 0x08,
	0x8d, 0x01, 0xf0, 0xaf, 0x34, 0xca, 0x70, 0xee, 0x46, 0x09, 0x37, 0xd3, 0x9c, 0xae, 0x44, 0x66,
	0x89, 0x75, 0x03, 0x03, 0x07, 0xaf, 0x32, 0x9c, 0xff, 0x50, 0xbf, 0xfc, 0xbf, 0x9d, 0xf5, 0x08,
	0x86, 0xe5, 0x23, 0x91, 0xa4, 0xf5, 0x12, 0xfa, 0xdf, 0x70, 0x16, 0xad, 0xd6, 0x4a, 0xe6, 0x0c,
	0xda, 0xf5, 0xf7, 0xe2, 0x62, 0x59, 0x30, 0x50, 0x34, 0xf9, 0xbb, 0x11, 0x68, 0x79, 0xb1, 0x94,
	0x2c, 0x76, 0x64, 0x52, 0x73, 0x12, 0x92, 0x82, 0x1e, 0x96, 0x1a, 0xc1, 0x40, 0xd1, 0x84, 0xd4,
	0xf5, 0x6f, 0x1d, 0x3a, 0x1f, 0xc4, 0x76, 0xa1, 0x77, 0xa0, 0xb3, 0xed, 0x42, 0x8f, 0x6d, 0xb1,
	0xd3, 0x62, 0x9b, 0x97, 0xc5, 0xca, 0xfe, 0xc4, 0x96, 0xdb, 0x1c, 0x1f, 0x5c, 0x46, 0xf4, 0x19,
	0xda, 0x7c, 0x40, 0x51, 0x93, 0x57, 0x5f, 0x2c, 0xf3, 0xd9, 0xbe, 0xb0, 0xd4, 0xf9, 0x02, 0x27,
	0x6a, 0xf2, 0xd0, 0xb4, 0xc1, 0xdd, 0x9a, 0x68, 0xf3, 0xfc, 0x00, 0xa3, 0x4a, 0x4c, 0xf6, 0xb5,
	0xc1, 0xad, 0xcf, 0xea, 0x8e, 0xc4, 0x36, 0x87, 0x6a, 0x0e, 0x1d, 0xd9, 0x42, 0x34, 0x69, 0x50,
	0x37, 0x27, 0xc2, 0x9c, 0xee, 0x27, 0x48, 0xb5, 0x19, 0x1c, 0x8b, 0xb6, 0xa2, 0xa6, 0xef, 0xc6,
	0x58, 0x98, 0x93, 0xbd, 0xf1, 0x4a, 0x4a, 0xb4, 0x15, 0xed, 0xac, 0x6d, 0x35, 0x16, 0x3b, 0xa4,
	0x36, 0xe7, 0xe1, 0x63, 0xfb, 0xbb, 0xe6, 0xa5, 0xd1, 0xf2, 0x98, 0xb7, 0xfe, 0xe6, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xda, 0x59, 0x5d, 0x05, 0x01, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Info(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InfoResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Info(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/edge.account.v1.Account/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/edge.account.v1.Account/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := c.cc.Invoke(ctx, "/edge.account.v1.Account/Callback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/edge.account.v1.Account/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/edge.account.v1.Account/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/edge.account.v1.Account/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/edge.account.v1.Account/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Info(context.Context, *empty.Empty) (*InfoResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Callback(context.Context, *CallbackRequest) (*CallbackResponse, error)
	Token(context.Context, *TokenRequest) (*TokenResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) Info(ctx context.Context, req *empty.Empty) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (*UnimplementedAccountServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAccountServer) Callback(ctx context.Context, req *CallbackRequest) (*CallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}
func (*UnimplementedAccountServer) Token(ctx context.Context, req *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (*UnimplementedAccountServer) Refresh(ctx context.Context, req *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (*UnimplementedAccountServer) Verify(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (*UnimplementedAccountServer) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edge.account.v1.Account/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Info(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edge.account.v1.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edge.account.v1.Account/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edge.account.v1.Account/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edge.account.v1.Account/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edge.account.v1.Account/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edge.account.v1.Account/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "edge.account.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Account_Info_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _Account_Callback_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Account_Token_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Account_Refresh_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Account_Verify_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Account_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/account.proto",
}
