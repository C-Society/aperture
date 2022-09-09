// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aperture/policy/language/v1/fluxmeter.proto

package languagev1

import (
	v1 "github.com/fluxninja/aperture/api/gen/proto/go/aperture/common/selector/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// FluxMeter gathers metrics for the traffic that matches its selector.
//
// Example of a selector that creates a histogram metric for all HTTP requests
// to particular service:
// ```yaml
// selector:
//   service: myservice.mynamespace.svc.cluster.local
//   control_point:
//     traffic: ingress
// ```
type FluxMeter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What latency should we measure in the histogram created by this FluxMeter.
	//
	// * For traffic control points, fluxmeter will measure the duration of the
	//   whole http transaction (including sending request and receiving
	//   response).
	// * For feature control points, fluxmeter will measure execution of the span
	//   associated with particular feature. What contributes to the span's
	//   duration is entirely up to the user code that uses Aperture library.
	Selector *v1.Selector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Latency histogram buckets (in ms) for this FluxMeter.
	HistogramBuckets []float64 `protobuf:"fixed64,2,rep,packed,name=histogram_buckets,json=histogramBuckets,proto3" json:"histogram_buckets,omitempty" default:"[5.0,10.0,25.0,50.0,100.0,250.0,500.0,1000.0,2500.0,5000.0,10000.0]"` // @gotags: default:"[5.0,10.0,25.0,50.0,100.0,250.0,500.0,1000.0,2500.0,5000.0,10000.0]"
	// Key of the attribute in accesss log or span from which the metric for this flux meter is read.
	//
	// :::info
	// For list of available attributes in Envoy access logs, refer
	// [Envoy Filter](/get-started/istio.md#envoy-filter)
	// :::
	//
	AttributeKey string `protobuf:"bytes,3,opt,name=attribute_key,json=attributeKey,proto3" json:"attribute_key,omitempty" default:"duration_millis"` // @gotags: default:"duration_millis"
}

func (x *FluxMeter) Reset() {
	*x = FluxMeter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_language_v1_fluxmeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FluxMeter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FluxMeter) ProtoMessage() {}

func (x *FluxMeter) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_language_v1_fluxmeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FluxMeter.ProtoReflect.Descriptor instead.
func (*FluxMeter) Descriptor() ([]byte, []int) {
	return file_aperture_policy_language_v1_fluxmeter_proto_rawDescGZIP(), []int{0}
}

func (x *FluxMeter) GetSelector() *v1.Selector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *FluxMeter) GetHistogramBuckets() []float64 {
	if x != nil {
		return x.HistogramBuckets
	}
	return nil
}

func (x *FluxMeter) GetAttributeKey() string {
	if x != nil {
		return x.AttributeKey
	}
	return ""
}

var File_aperture_policy_language_v1_fluxmeter_proto protoreflect.FileDescriptor

var file_aperture_policy_language_v1_fluxmeter_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c,
	0x75, 0x78, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61,
	0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x09, 0x46, 0x6c, 0x75, 0x78, 0x4d,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x01, 0x42, 0x5b, 0x92, 0x41, 0x58, 0x82, 0x03, 0x55, 0x0a, 0x0c, 0x78, 0x2d, 0x67,
	0x6f, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x45, 0x1a, 0x43, 0x5b, 0x35, 0x2e,
	0x30, 0x2c, 0x31, 0x30, 0x2e, 0x30, 0x2c, 0x32, 0x35, 0x2e, 0x30, 0x2c, 0x35, 0x30, 0x2e, 0x30,
	0x2c, 0x31, 0x30, 0x30, 0x2e, 0x30, 0x2c, 0x32, 0x35, 0x30, 0x2e, 0x30, 0x2c, 0x35, 0x30, 0x30,
	0x2e, 0x30, 0x2c, 0x31, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x2c, 0x32, 0x35, 0x30, 0x30, 0x2e, 0x30,
	0x2c, 0x35, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x2c, 0x31, 0x30, 0x30, 0x30, 0x30, 0x2e, 0x30, 0x5d,
	0x52, 0x10, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0x92, 0x41, 0x24, 0x82, 0x03,
	0x21, 0x0a, 0x0c, 0x78, 0x2d, 0x67, 0x6f, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x11, 0x1a, 0x0f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x42, 0x97, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x46, 0x6c, 0x75, 0x78, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2f, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x50, 0x4c, 0xaa, 0x02, 0x1b, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1b, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x27, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x41, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x3a, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_aperture_policy_language_v1_fluxmeter_proto_rawDescOnce sync.Once
	file_aperture_policy_language_v1_fluxmeter_proto_rawDescData = file_aperture_policy_language_v1_fluxmeter_proto_rawDesc
)

func file_aperture_policy_language_v1_fluxmeter_proto_rawDescGZIP() []byte {
	file_aperture_policy_language_v1_fluxmeter_proto_rawDescOnce.Do(func() {
		file_aperture_policy_language_v1_fluxmeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_policy_language_v1_fluxmeter_proto_rawDescData)
	})
	return file_aperture_policy_language_v1_fluxmeter_proto_rawDescData
}

var file_aperture_policy_language_v1_fluxmeter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aperture_policy_language_v1_fluxmeter_proto_goTypes = []interface{}{
	(*FluxMeter)(nil),   // 0: aperture.policy.language.v1.FluxMeter
	(*v1.Selector)(nil), // 1: aperture.common.selector.v1.Selector
}
var file_aperture_policy_language_v1_fluxmeter_proto_depIdxs = []int32{
	1, // 0: aperture.policy.language.v1.FluxMeter.selector:type_name -> aperture.common.selector.v1.Selector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aperture_policy_language_v1_fluxmeter_proto_init() }
func file_aperture_policy_language_v1_fluxmeter_proto_init() {
	if File_aperture_policy_language_v1_fluxmeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aperture_policy_language_v1_fluxmeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FluxMeter); i {
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
			RawDescriptor: file_aperture_policy_language_v1_fluxmeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aperture_policy_language_v1_fluxmeter_proto_goTypes,
		DependencyIndexes: file_aperture_policy_language_v1_fluxmeter_proto_depIdxs,
		MessageInfos:      file_aperture_policy_language_v1_fluxmeter_proto_msgTypes,
	}.Build()
	File_aperture_policy_language_v1_fluxmeter_proto = out.File
	file_aperture_policy_language_v1_fluxmeter_proto_rawDesc = nil
	file_aperture_policy_language_v1_fluxmeter_proto_goTypes = nil
	file_aperture_policy_language_v1_fluxmeter_proto_depIdxs = nil
}