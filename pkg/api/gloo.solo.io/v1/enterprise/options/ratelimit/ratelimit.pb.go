// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/ratelimit/ratelimit.proto

package ratelimit

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v1alpha1 "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Basic rate-limiting API
type IngressRateLimit struct {
	AuthorizedLimits     *v1alpha1.RateLimit `protobuf:"bytes,1,opt,name=authorized_limits,json=authorizedLimits,proto3" json:"authorized_limits,omitempty"`
	AnonymousLimits      *v1alpha1.RateLimit `protobuf:"bytes,2,opt,name=anonymous_limits,json=anonymousLimits,proto3" json:"anonymous_limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IngressRateLimit) Reset()         { *m = IngressRateLimit{} }
func (m *IngressRateLimit) String() string { return proto.CompactTextString(m) }
func (*IngressRateLimit) ProtoMessage()    {}
func (*IngressRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5d71ce45d166ff5, []int{0}
}
func (m *IngressRateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngressRateLimit.Unmarshal(m, b)
}
func (m *IngressRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngressRateLimit.Marshal(b, m, deterministic)
}
func (m *IngressRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressRateLimit.Merge(m, src)
}
func (m *IngressRateLimit) XXX_Size() int {
	return xxx_messageInfo_IngressRateLimit.Size(m)
}
func (m *IngressRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_IngressRateLimit proto.InternalMessageInfo

func (m *IngressRateLimit) GetAuthorizedLimits() *v1alpha1.RateLimit {
	if m != nil {
		return m.AuthorizedLimits
	}
	return nil
}

func (m *IngressRateLimit) GetAnonymousLimits() *v1alpha1.RateLimit {
	if m != nil {
		return m.AnonymousLimits
	}
	return nil
}

type Settings struct {
	RatelimitServerRef *core.ResourceRef `protobuf:"bytes,1,opt,name=ratelimit_server_ref,json=ratelimitServerRef,proto3" json:"ratelimit_server_ref,omitempty"`
	RequestTimeout     *time.Duration    `protobuf:"bytes,2,opt,name=request_timeout,json=requestTimeout,proto3,stdduration" json:"request_timeout,omitempty"`
	DenyOnFail         bool              `protobuf:"varint,3,opt,name=deny_on_fail,json=denyOnFail,proto3" json:"deny_on_fail,omitempty"`
	// Set this is set to true if you would like to rate limit traffic before applying external auth to it.
	// *Note*: When this is true, you will lose some features like being able to rate limit a request based on its auth state
	RateLimitBeforeAuth  bool     `protobuf:"varint,9,opt,name=rate_limit_before_auth,json=rateLimitBeforeAuth,proto3" json:"rate_limit_before_auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5d71ce45d166ff5, []int{1}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetRatelimitServerRef() *core.ResourceRef {
	if m != nil {
		return m.RatelimitServerRef
	}
	return nil
}

func (m *Settings) GetRequestTimeout() *time.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *Settings) GetDenyOnFail() bool {
	if m != nil {
		return m.DenyOnFail
	}
	return false
}

func (m *Settings) GetRateLimitBeforeAuth() bool {
	if m != nil {
		return m.RateLimitBeforeAuth
	}
	return false
}

// API based on Envoy's rate-limit service API. (reference here: https://github.com/lyft/ratelimit#configuration)
// Sample configuration below:
//
// descriptors:
//- key: account_id
//  descriptors:
//  - key: plan
//    value: BASIC
//    rateLimit:
//      requestsPerUnit: 1
//      unit: MINUTE
//  - key: plan
//    value: PLUS
//    rateLimit:
//      requestsPerUnit: 20
//      unit: MINUTE
type ServiceSettings struct {
	Descriptors          []*v1alpha1.Descriptor    `protobuf:"bytes,1,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	SetDescriptors       []*v1alpha1.SetDescriptor `protobuf:"bytes,2,rep,name=set_descriptors,json=setDescriptors,proto3" json:"set_descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ServiceSettings) Reset()         { *m = ServiceSettings{} }
func (m *ServiceSettings) String() string { return proto.CompactTextString(m) }
func (*ServiceSettings) ProtoMessage()    {}
func (*ServiceSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5d71ce45d166ff5, []int{2}
}
func (m *ServiceSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSettings.Unmarshal(m, b)
}
func (m *ServiceSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSettings.Marshal(b, m, deterministic)
}
func (m *ServiceSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSettings.Merge(m, src)
}
func (m *ServiceSettings) XXX_Size() int {
	return xxx_messageInfo_ServiceSettings.Size(m)
}
func (m *ServiceSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSettings proto.InternalMessageInfo

func (m *ServiceSettings) GetDescriptors() []*v1alpha1.Descriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *ServiceSettings) GetSetDescriptors() []*v1alpha1.SetDescriptor {
	if m != nil {
		return m.SetDescriptors
	}
	return nil
}

// A list of references to `RateLimitConfig` resources.
// Each resource represents a rate limit policy that will be independently enforced.
type RateLimitConfigRefs struct {
	Refs                 []*RateLimitConfigRef `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RateLimitConfigRefs) Reset()         { *m = RateLimitConfigRefs{} }
func (m *RateLimitConfigRefs) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfigRefs) ProtoMessage()    {}
func (*RateLimitConfigRefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5d71ce45d166ff5, []int{3}
}
func (m *RateLimitConfigRefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfigRefs.Unmarshal(m, b)
}
func (m *RateLimitConfigRefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfigRefs.Marshal(b, m, deterministic)
}
func (m *RateLimitConfigRefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfigRefs.Merge(m, src)
}
func (m *RateLimitConfigRefs) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfigRefs.Size(m)
}
func (m *RateLimitConfigRefs) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfigRefs.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfigRefs proto.InternalMessageInfo

func (m *RateLimitConfigRefs) GetRefs() []*RateLimitConfigRef {
	if m != nil {
		return m.Refs
	}
	return nil
}

// A reference to a `RateLimitConfig` resource.
type RateLimitConfigRef struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitConfigRef) Reset()         { *m = RateLimitConfigRef{} }
func (m *RateLimitConfigRef) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfigRef) ProtoMessage()    {}
func (*RateLimitConfigRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5d71ce45d166ff5, []int{4}
}
func (m *RateLimitConfigRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfigRef.Unmarshal(m, b)
}
func (m *RateLimitConfigRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfigRef.Marshal(b, m, deterministic)
}
func (m *RateLimitConfigRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfigRef.Merge(m, src)
}
func (m *RateLimitConfigRef) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfigRef.Size(m)
}
func (m *RateLimitConfigRef) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfigRef.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfigRef proto.InternalMessageInfo

func (m *RateLimitConfigRef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RateLimitConfigRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Use this field if you want to inline the Envoy rate limits for this VirtualHost.
// Note that this does not configure the rate limit server. If you are running Gloo Enterprise, you need to
// specify the server configuration via the appropriate field in the Gloo `Settings` resource. If you are
// running a custom rate limit server you need to configure it yourself.
type RateLimitVhostExtension struct {
	// Define individual rate limits here. Each rate limit will be evaluated, if any rate limit
	// would be throttled, the entire request returns a 429 (gets throttled)
	RateLimits           []*v1alpha1.RateLimitActions `protobuf:"bytes,1,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitVhostExtension) Reset()         { *m = RateLimitVhostExtension{} }
func (m *RateLimitVhostExtension) String() string { return proto.CompactTextString(m) }
func (*RateLimitVhostExtension) ProtoMessage()    {}
func (*RateLimitVhostExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5d71ce45d166ff5, []int{5}
}
func (m *RateLimitVhostExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitVhostExtension.Unmarshal(m, b)
}
func (m *RateLimitVhostExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitVhostExtension.Marshal(b, m, deterministic)
}
func (m *RateLimitVhostExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitVhostExtension.Merge(m, src)
}
func (m *RateLimitVhostExtension) XXX_Size() int {
	return xxx_messageInfo_RateLimitVhostExtension.Size(m)
}
func (m *RateLimitVhostExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitVhostExtension.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitVhostExtension proto.InternalMessageInfo

func (m *RateLimitVhostExtension) GetRateLimits() []*v1alpha1.RateLimitActions {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

// Use this field if you want to inline the Envoy rate limits for this Route.
// Note that this does not configure the rate limit server. If you are running Gloo Enterprise, you need to
// specify the server configuration via the appropriate field in the Gloo `Settings` resource. If you are
// running a custom rate limit server you need to configure it yourself.
type RateLimitRouteExtension struct {
	// Whether or not to include rate limits as defined on the VirtualHost in addition to rate limits on the Route.
	IncludeVhRateLimits bool `protobuf:"varint,1,opt,name=include_vh_rate_limits,json=includeVhRateLimits,proto3" json:"include_vh_rate_limits,omitempty"`
	// Define individual rate limits here. Each rate limit will be evaluated, if any rate limit
	// would be throttled, the entire request returns a 429 (gets throttled)
	RateLimits           []*v1alpha1.RateLimitActions `protobuf:"bytes,2,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitRouteExtension) Reset()         { *m = RateLimitRouteExtension{} }
func (m *RateLimitRouteExtension) String() string { return proto.CompactTextString(m) }
func (*RateLimitRouteExtension) ProtoMessage()    {}
func (*RateLimitRouteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5d71ce45d166ff5, []int{6}
}
func (m *RateLimitRouteExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRouteExtension.Unmarshal(m, b)
}
func (m *RateLimitRouteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRouteExtension.Marshal(b, m, deterministic)
}
func (m *RateLimitRouteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRouteExtension.Merge(m, src)
}
func (m *RateLimitRouteExtension) XXX_Size() int {
	return xxx_messageInfo_RateLimitRouteExtension.Size(m)
}
func (m *RateLimitRouteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRouteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRouteExtension proto.InternalMessageInfo

func (m *RateLimitRouteExtension) GetIncludeVhRateLimits() bool {
	if m != nil {
		return m.IncludeVhRateLimits
	}
	return false
}

func (m *RateLimitRouteExtension) GetRateLimits() []*v1alpha1.RateLimitActions {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func init() {
	proto.RegisterType((*IngressRateLimit)(nil), "ratelimit.options.gloo.solo.io.IngressRateLimit")
	proto.RegisterType((*Settings)(nil), "ratelimit.options.gloo.solo.io.Settings")
	proto.RegisterType((*ServiceSettings)(nil), "ratelimit.options.gloo.solo.io.ServiceSettings")
	proto.RegisterType((*RateLimitConfigRefs)(nil), "ratelimit.options.gloo.solo.io.RateLimitConfigRefs")
	proto.RegisterType((*RateLimitConfigRef)(nil), "ratelimit.options.gloo.solo.io.RateLimitConfigRef")
	proto.RegisterType((*RateLimitVhostExtension)(nil), "ratelimit.options.gloo.solo.io.RateLimitVhostExtension")
	proto.RegisterType((*RateLimitRouteExtension)(nil), "ratelimit.options.gloo.solo.io.RateLimitRouteExtension")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/ratelimit/ratelimit.proto", fileDescriptor_c5d71ce45d166ff5)
}

var fileDescriptor_c5d71ce45d166ff5 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x4b, 0xdc, 0x4c,
	0x14, 0xc7, 0x89, 0x2e, 0x0f, 0xee, 0xec, 0x83, 0xeb, 0x33, 0x8a, 0x8f, 0x4a, 0xb1, 0xdb, 0xa5,
	0x50, 0x6f, 0x4c, 0x58, 0xf7, 0x13, 0xf8, 0x52, 0xb1, 0xb4, 0x52, 0x3a, 0x16, 0x2f, 0x2c, 0x25,
	0xcc, 0x66, 0x4f, 0x92, 0xc1, 0xec, 0x9c, 0x74, 0x66, 0xb2, 0xd5, 0x7e, 0x12, 0xbf, 0x40, 0xa1,
	0x57, 0xbd, 0xee, 0xb7, 0x29, 0xf4, 0x3b, 0x14, 0x7a, 0x59, 0x32, 0x79, 0x5b, 0x15, 0xb5, 0xf4,
	0x26, 0x99, 0x9c, 0x73, 0xfe, 0xbf, 0xf9, 0xe7, 0x9c, 0x64, 0xc8, 0x59, 0x24, 0x4c, 0x9c, 0x8d,
	0xdc, 0x00, 0x27, 0x9e, 0xc6, 0x04, 0xb7, 0x05, 0x16, 0x77, 0x9e, 0x0a, 0xed, 0xf1, 0x54, 0x78,
	0x51, 0x82, 0x58, 0x5c, 0xa6, 0x03, 0x0f, 0xa4, 0x01, 0x95, 0x2a, 0xa1, 0xc1, 0xc3, 0xd4, 0x08,
	0x94, 0xda, 0x53, 0xdc, 0x40, 0x22, 0x26, 0xc2, 0x34, 0x2b, 0x37, 0x55, 0x68, 0x90, 0x6e, 0x36,
	0x81, 0xb2, 0xd8, 0xcd, 0x39, 0x6e, 0xce, 0x77, 0x05, 0x6e, 0x0c, 0xaf, 0x6f, 0x94, 0x57, 0x6f,
	0xdb, 0x72, 0x50, 0xde, 0x74, 0xc0, 0x93, 0x34, 0xe6, 0x83, 0x9b, 0xd0, 0x8d, 0x75, 0x2b, 0x3a,
	0x17, 0xc6, 0x6a, 0xa6, 0x03, 0x4f, 0x41, 0x58, 0xa6, 0x36, 0x23, 0xc4, 0x28, 0x01, 0xcf, 0x3e,
	0x8d, 0xb2, 0xd0, 0xfb, 0xa8, 0x78, 0x9a, 0x82, 0xd2, 0x77, 0xe5, 0xc7, 0x99, 0xe2, 0xb9, 0xaf,
	0x32, 0xbf, 0x12, 0x61, 0x84, 0x76, 0xe9, 0xe5, 0xab, 0x32, 0x4a, 0xe1, 0xc2, 0x14, 0x41, 0xb8,
	0x28, 0x4d, 0xf4, 0xbf, 0x3a, 0x64, 0xe9, 0x85, 0x8c, 0x14, 0x68, 0xcd, 0xb8, 0x81, 0x57, 0xb9,
	0x3f, 0x7a, 0x4c, 0xfe, 0xe3, 0x99, 0x89, 0x51, 0x89, 0x4f, 0x30, 0xf6, 0xad, 0x67, 0xbd, 0xe6,
	0xf4, 0x9c, 0xad, 0xce, 0x4e, 0xcf, 0x6d, 0x5e, 0x83, 0xa7, 0xa2, 0xea, 0x80, 0x5b, 0x8b, 0xd9,
	0x52, 0x23, 0xb5, 0x01, 0x4d, 0x5f, 0x92, 0x25, 0x2e, 0x51, 0x5e, 0x4e, 0x30, 0xd3, 0x15, 0x6d,
	0xee, 0x0f, 0x69, 0xdd, 0x5a, 0x59, 0xc0, 0xfa, 0xbf, 0x1c, 0xb2, 0x70, 0x02, 0xc6, 0x08, 0x19,
	0xe5, 0xe4, 0x95, 0x1a, 0xe0, 0x6b, 0x50, 0x53, 0x50, 0xbe, 0x82, 0xb0, 0xf4, 0xba, 0xee, 0x06,
	0xa8, 0xa0, 0x81, 0x82, 0xc6, 0x4c, 0x05, 0xc0, 0x20, 0x64, 0xb4, 0x96, 0x9d, 0x58, 0x15, 0x83,
	0x90, 0x1e, 0x91, 0xae, 0x82, 0x0f, 0x19, 0x68, 0xe3, 0x1b, 0x31, 0x01, 0xcc, 0x4c, 0xe9, 0x72,
	0xdd, 0x2d, 0xda, 0xed, 0x56, 0xed, 0x76, 0x0f, 0xca, 0x76, 0xef, 0xb5, 0xae, 0xbe, 0x3f, 0x76,
	0xd8, 0x62, 0xa9, 0x7b, 0x5b, 0xc8, 0x68, 0x8f, 0xfc, 0x3b, 0x06, 0x79, 0xe9, 0xa3, 0xf4, 0x43,
	0x2e, 0x92, 0xb5, 0xf9, 0x9e, 0xb3, 0xb5, 0xc0, 0x48, 0x1e, 0x7b, 0x2d, 0x0f, 0xb9, 0x48, 0xe8,
	0x90, 0xac, 0xe6, 0x0e, 0x8a, 0x6e, 0xf8, 0x23, 0x08, 0x51, 0x81, 0x9f, 0x37, 0x6e, 0xad, 0x6d,
	0x6b, 0x97, 0x55, 0xd5, 0x81, 0x3d, 0x9b, 0xdb, 0xcd, 0x4c, 0xdc, 0xff, 0xec, 0x90, 0x6e, 0x6e,
	0x57, 0x04, 0x50, 0x77, 0x60, 0x9f, 0x74, 0xc6, 0xa0, 0x03, 0x25, 0x52, 0x83, 0x2a, 0x1f, 0xd2,
	0xfc, 0x56, 0x67, 0xe7, 0xc9, 0x1d, 0x6d, 0x3d, 0xa8, 0x2b, 0xd9, 0xac, 0x8a, 0x1e, 0x93, 0xae,
	0x06, 0xe3, 0xcf, 0x82, 0xe6, 0x2c, 0xe8, 0xe9, 0x1d, 0xa0, 0x13, 0x30, 0x33, 0xac, 0x45, 0x3d,
	0xfb, 0xa8, 0xfb, 0xef, 0xc9, 0x72, 0x3d, 0xc0, 0x7d, 0x94, 0xa1, 0x88, 0x18, 0x84, 0x9a, 0x1e,
	0x92, 0x96, 0x82, 0xb0, 0xf2, 0xb8, 0xe3, 0xde, 0xff, 0x4f, 0xb9, 0xb7, 0x11, 0xcc, 0xea, 0xfb,
	0x87, 0x84, 0xde, 0xce, 0x51, 0x4a, 0x5a, 0x92, 0x4f, 0xc0, 0x8e, 0xbe, 0xcd, 0xec, 0x9a, 0x3e,
	0x22, 0xed, 0xfc, 0xae, 0x53, 0x1e, 0x80, 0x9d, 0x65, 0x9b, 0x35, 0x81, 0x7e, 0x40, 0xfe, 0xaf,
	0x39, 0xa7, 0x31, 0x6a, 0xf3, 0xfc, 0xc2, 0x80, 0xd4, 0x02, 0x25, 0x3d, 0x22, 0x9d, 0x66, 0x3c,
	0x95, 0xe3, 0x67, 0x0f, 0x7d, 0xac, 0xbb, 0x81, 0x7d, 0x0d, 0x46, 0xea, 0xe1, 0xe9, 0xfe, 0x95,
	0x33, 0xb3, 0x0b, 0xc3, 0xcc, 0x40, 0xb3, 0xcb, 0x90, 0xac, 0x0a, 0x19, 0x24, 0xd9, 0x18, 0xfc,
	0x69, 0xec, 0x5f, 0xdf, 0xd0, 0x7e, 0x04, 0x65, 0xf6, 0x34, 0xae, 0x09, 0xfa, 0xa6, 0xb5, 0xb9,
	0xbf, 0xb6, 0xb6, 0xf7, 0xee, 0xdb, 0xcf, 0x96, 0xf3, 0xe5, 0xc7, 0xa6, 0x73, 0xf6, 0xe6, 0xde,
	0xa3, 0x33, 0x3d, 0x8f, 0xea, 0xe3, 0xb3, 0x22, 0x3f, 0x74, 0x82, 0x8e, 0xfe, 0xb1, 0xff, 0xca,
	0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x9b, 0x81, 0x50, 0x96, 0x05, 0x00, 0x00,
}

func (this *IngressRateLimit) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IngressRateLimit)
	if !ok {
		that2, ok := that.(IngressRateLimit)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AuthorizedLimits.Equal(that1.AuthorizedLimits) {
		return false
	}
	if !this.AnonymousLimits.Equal(that1.AnonymousLimits) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.RatelimitServerRef.Equal(that1.RatelimitServerRef) {
		return false
	}
	if this.RequestTimeout != nil && that1.RequestTimeout != nil {
		if *this.RequestTimeout != *that1.RequestTimeout {
			return false
		}
	} else if this.RequestTimeout != nil {
		return false
	} else if that1.RequestTimeout != nil {
		return false
	}
	if this.DenyOnFail != that1.DenyOnFail {
		return false
	}
	if this.RateLimitBeforeAuth != that1.RateLimitBeforeAuth {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSettings)
	if !ok {
		that2, ok := that.(ServiceSettings)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Descriptors) != len(that1.Descriptors) {
		return false
	}
	for i := range this.Descriptors {
		if !this.Descriptors[i].Equal(that1.Descriptors[i]) {
			return false
		}
	}
	if len(this.SetDescriptors) != len(that1.SetDescriptors) {
		return false
	}
	for i := range this.SetDescriptors {
		if !this.SetDescriptors[i].Equal(that1.SetDescriptors[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RateLimitConfigRefs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitConfigRefs)
	if !ok {
		that2, ok := that.(RateLimitConfigRefs)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Refs) != len(that1.Refs) {
		return false
	}
	for i := range this.Refs {
		if !this.Refs[i].Equal(that1.Refs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RateLimitConfigRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitConfigRef)
	if !ok {
		that2, ok := that.(RateLimitConfigRef)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RateLimitVhostExtension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitVhostExtension)
	if !ok {
		that2, ok := that.(RateLimitVhostExtension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.RateLimits) != len(that1.RateLimits) {
		return false
	}
	for i := range this.RateLimits {
		if !this.RateLimits[i].Equal(that1.RateLimits[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RateLimitRouteExtension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RateLimitRouteExtension)
	if !ok {
		that2, ok := that.(RateLimitRouteExtension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.IncludeVhRateLimits != that1.IncludeVhRateLimits {
		return false
	}
	if len(this.RateLimits) != len(that1.RateLimits) {
		return false
	}
	for i := range this.RateLimits {
		if !this.RateLimits[i].Equal(that1.RateLimits[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
