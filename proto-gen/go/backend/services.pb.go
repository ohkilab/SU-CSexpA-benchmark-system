// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: backend/services.proto

package backend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_backend_services_proto protoreflect.FileDescriptor

var file_backend_services_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xa2, 0x02, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x50, 0x6f,
	0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x11,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa8, 0x01, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09,
	0x50, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5e, 0x0a, 0x17, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69,
	0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x50, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_backend_services_proto_goTypes = []interface{}{
	(*GetRankingRequest)(nil),               // 0: GetRankingRequest
	(*PostSubmitRequest)(nil),               // 1: PostSubmitRequest
	(*GetSubmitRequest)(nil),                // 2: GetSubmitRequest
	(*ListSubmitsRequest)(nil),              // 3: ListSubmitsRequest
	(*PostLoginRequest)(nil),                // 4: PostLoginRequest
	(*PingUnaryRequest)(nil),                // 5: PingUnaryRequest
	(*PingServerSideStreamingRequest)(nil),  // 6: PingServerSideStreamingRequest
	(*GetRankingResponse)(nil),              // 7: GetRankingResponse
	(*PostSubmitResponse)(nil),              // 8: PostSubmitResponse
	(*GetSubmitResponse)(nil),               // 9: GetSubmitResponse
	(*ListSubmitsResponse)(nil),             // 10: ListSubmitsResponse
	(*PostLoginResponse)(nil),               // 11: PostLoginResponse
	(*PingUnaryResponse)(nil),               // 12: PingUnaryResponse
	(*PingServerSideStreamingResponse)(nil), // 13: PingServerSideStreamingResponse
}
var file_backend_services_proto_depIdxs = []int32{
	0,  // 0: BackendService.GetRanking:input_type -> GetRankingRequest
	1,  // 1: BackendService.PostSubmit:input_type -> PostSubmitRequest
	2,  // 2: BackendService.GetSubmit:input_type -> GetSubmitRequest
	3,  // 3: BackendService.ListSubmits:input_type -> ListSubmitsRequest
	4,  // 4: BackendService.PostLogin:input_type -> PostLoginRequest
	5,  // 5: HealthcheckService.PingUnary:input_type -> PingUnaryRequest
	6,  // 6: HealthcheckService.PingServerSideStreaming:input_type -> PingServerSideStreamingRequest
	7,  // 7: BackendService.GetRanking:output_type -> GetRankingResponse
	8,  // 8: BackendService.PostSubmit:output_type -> PostSubmitResponse
	9,  // 9: BackendService.GetSubmit:output_type -> GetSubmitResponse
	10, // 10: BackendService.ListSubmits:output_type -> ListSubmitsResponse
	11, // 11: BackendService.PostLogin:output_type -> PostLoginResponse
	12, // 12: HealthcheckService.PingUnary:output_type -> PingUnaryResponse
	13, // 13: HealthcheckService.PingServerSideStreaming:output_type -> PingServerSideStreamingResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_backend_services_proto_init() }
func file_backend_services_proto_init() {
	if File_backend_services_proto != nil {
		return
	}
	file_backend_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_backend_services_proto_goTypes,
		DependencyIndexes: file_backend_services_proto_depIdxs,
	}.Build()
	File_backend_services_proto = out.File
	file_backend_services_proto_rawDesc = nil
	file_backend_services_proto_goTypes = nil
	file_backend_services_proto_depIdxs = nil
}
