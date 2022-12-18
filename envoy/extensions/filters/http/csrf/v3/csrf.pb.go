// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: envoy/extensions/filters/http/csrf/v3/csrf.proto

package csrfv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// CSRF filter config.
type CsrfPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the % of requests for which the CSRF filter is enabled.
	//
	// If :ref:`runtime_key <envoy_v3_api_field_config.core.v3.RuntimeFractionalPercent.runtime_key>` is specified,
	// Envoy will lookup the runtime key to get the percentage of requests to filter.
	//
	// .. note::
	//
	//   This field defaults to 100/:ref:`HUNDRED
	//   <envoy_v3_api_enum_type.v3.FractionalPercent.DenominatorType>`.
	FilterEnabled *v3.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	// Specifies that CSRF policies will be evaluated and tracked, but not enforced.
	//
	// This is intended to be used when ``filter_enabled`` is off and will be ignored otherwise.
	//
	// If :ref:`runtime_key <envoy_v3_api_field_config.core.v3.RuntimeFractionalPercent.runtime_key>` is specified,
	// Envoy will lookup the runtime key to get the percentage of requests for which it will evaluate
	// and track the request's ``Origin`` and ``Destination`` to determine if it's valid, but will not
	// enforce any policies.
	ShadowEnabled *v3.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	// Specifies additional source origins that will be allowed in addition to
	// the destination origin.
	//
	// More information on how this can be configured via runtime can be found
	// :ref:`here <csrf-configuration>`.
	AdditionalOrigins []*v31.StringMatcher `protobuf:"bytes,3,rep,name=additional_origins,json=additionalOrigins,proto3" json:"additional_origins,omitempty"`
}

func (x *CsrfPolicy) Reset() {
	*x = CsrfPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_csrf_v3_csrf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsrfPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsrfPolicy) ProtoMessage() {}

func (x *CsrfPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_csrf_v3_csrf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsrfPolicy.ProtoReflect.Descriptor instead.
func (*CsrfPolicy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescGZIP(), []int{0}
}

func (x *CsrfPolicy) GetFilterEnabled() *v3.RuntimeFractionalPercent {
	if x != nil {
		return x.FilterEnabled
	}
	return nil
}

func (x *CsrfPolicy) GetShadowEnabled() *v3.RuntimeFractionalPercent {
	if x != nil {
		return x.ShadowEnabled
	}
	return nil
}

func (x *CsrfPolicy) GetAdditionalOrigins() []*v31.StringMatcher {
	if x != nil {
		return x.AdditionalOrigins
	}
	return nil
}

var File_envoy_extensions_filters_http_csrf_v3_csrf_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDesc = []byte{
	0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x73, 0x72, 0x66, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x73, 0x72, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x63, 0x73, 0x72, 0x66, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x0a, 0x43, 0x73,
	0x72, 0x66, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x5f, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x55, 0x0a, 0x0e, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x52, 0x0d, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x53, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x73, 0x3a, 0x32, 0x9a, 0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x73, 0x72, 0x66, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x73, 0x72, 0x66, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x9f, 0x01, 0x0a, 0x33, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x73, 0x72, 0x66, 0x2e, 0x76,
	0x33, 0x42, 0x09, 0x43, 0x73, 0x72, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x73, 0x72, 0x66, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x73, 0x72,
	0x66, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescData = file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDesc
)

func file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDescData
}

var file_envoy_extensions_filters_http_csrf_v3_csrf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_csrf_v3_csrf_proto_goTypes = []interface{}{
	(*CsrfPolicy)(nil),                  // 0: envoy.extensions.filters.http.csrf.v3.CsrfPolicy
	(*v3.RuntimeFractionalPercent)(nil), // 1: envoy.config.core.v3.RuntimeFractionalPercent
	(*v31.StringMatcher)(nil),           // 2: envoy.type.matcher.v3.StringMatcher
}
var file_envoy_extensions_filters_http_csrf_v3_csrf_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.http.csrf.v3.CsrfPolicy.filter_enabled:type_name -> envoy.config.core.v3.RuntimeFractionalPercent
	1, // 1: envoy.extensions.filters.http.csrf.v3.CsrfPolicy.shadow_enabled:type_name -> envoy.config.core.v3.RuntimeFractionalPercent
	2, // 2: envoy.extensions.filters.http.csrf.v3.CsrfPolicy.additional_origins:type_name -> envoy.type.matcher.v3.StringMatcher
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_csrf_v3_csrf_proto_init() }
func file_envoy_extensions_filters_http_csrf_v3_csrf_proto_init() {
	if File_envoy_extensions_filters_http_csrf_v3_csrf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_csrf_v3_csrf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsrfPolicy); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_csrf_v3_csrf_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_csrf_v3_csrf_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_csrf_v3_csrf_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_csrf_v3_csrf_proto = out.File
	file_envoy_extensions_filters_http_csrf_v3_csrf_proto_rawDesc = nil
	file_envoy_extensions_filters_http_csrf_v3_csrf_proto_goTypes = nil
	file_envoy_extensions_filters_http_csrf_v3_csrf_proto_depIdxs = nil
}
