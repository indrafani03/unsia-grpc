// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: city_service.proto

package cities

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

var File_city_service_proto protoreflect.FileDescriptor

var file_city_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x12, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xf3, 0x01, 0x0a, 0x0d, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x2e,
	0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x63, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x63,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x0c, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x69, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e,
	0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x63, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x49,
	0x64, 0x1a, 0x11, 0x2e, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x4d, 0x79, 0x42, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x62, 0x2f, 0x63, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x3b, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_city_service_proto_goTypes = []any{
	(*Id)(nil),           // 0: cities.Id
	(*EmptyMessage)(nil), // 1: cities.EmptyMessage
	(*CityInput)(nil),    // 2: cities.CityInput
	(*City)(nil),         // 3: cities.City
	(*CitiesStream)(nil), // 4: cities.CitiesStream
	(*MyBoolean)(nil),    // 5: cities.MyBoolean
}
var file_city_service_proto_depIdxs = []int32{
	0, // 0: cities.CitiesService.GetCity:input_type -> cities.Id
	1, // 1: cities.CitiesService.GetCities:input_type -> cities.EmptyMessage
	2, // 2: cities.CitiesService.Create:input_type -> cities.CityInput
	3, // 3: cities.CitiesService.Update:input_type -> cities.City
	0, // 4: cities.CitiesService.Delete:input_type -> cities.Id
	3, // 5: cities.CitiesService.GetCity:output_type -> cities.City
	4, // 6: cities.CitiesService.GetCities:output_type -> cities.CitiesStream
	3, // 7: cities.CitiesService.Create:output_type -> cities.City
	3, // 8: cities.CitiesService.Update:output_type -> cities.City
	5, // 9: cities.CitiesService.Delete:output_type -> cities.MyBoolean
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_city_service_proto_init() }
func file_city_service_proto_init() {
	if File_city_service_proto != nil {
		return
	}
	file_city_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_city_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_city_service_proto_goTypes,
		DependencyIndexes: file_city_service_proto_depIdxs,
	}.Build()
	File_city_service_proto = out.File
	file_city_service_proto_rawDesc = nil
	file_city_service_proto_goTypes = nil
	file_city_service_proto_depIdxs = nil
}
