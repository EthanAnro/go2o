// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: message_service.proto

package proto

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

var File_message_service_proto protoreflect.FileDescriptor

var file_message_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xc5, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0c, 0x2e,
	0x53, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x10, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x13, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0c, 0x2e, 0x53, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x0e,
	0x2e, 0x53, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x53, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x06,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x69, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_message_service_proto_goTypes = []interface{}{
	(*String)(nil),                   // 0: String
	(*SendMessageRequest)(nil),       // 1: SendMessageRequest
	(*Empty)(nil),                    // 2: Empty
	(*SNotifyItem)(nil),              // 3: SNotifyItem
	(*Int64)(nil),                    // 4: Int64
	(*SMailTemplate)(nil),            // 5: SMailTemplate
	(*SendSiteMessageRequest)(nil),   // 6: SendSiteMessageRequest
	(*Result)(nil),                   // 7: Result
	(*NotifyItemListResponse)(nil),   // 8: NotifyItemListResponse
	(*MailTemplateListResponse)(nil), // 9: MailTemplateListResponse
}
var file_message_service_proto_depIdxs = []int32{
	0, // 0: MessageService.GetNotifyItem:input_type -> String
	1, // 1: MessageService.SendPhoneMessage:input_type -> SendMessageRequest
	2, // 2: MessageService.GetAllNotifyItem:input_type -> Empty
	3, // 3: MessageService.SaveNotifyItem:input_type -> SNotifyItem
	4, // 4: MessageService.GetMailTemplate:input_type -> Int64
	5, // 5: MessageService.SaveMailTemplate:input_type -> SMailTemplate
	2, // 6: MessageService.GetMailTemplates:input_type -> Empty
	4, // 7: MessageService.DeleteMailTemplate:input_type -> Int64
	6, // 8: MessageService.SendSiteMessage:input_type -> SendSiteMessageRequest
	3, // 9: MessageService.GetNotifyItem:output_type -> SNotifyItem
	7, // 10: MessageService.SendPhoneMessage:output_type -> Result
	8, // 11: MessageService.GetAllNotifyItem:output_type -> NotifyItemListResponse
	7, // 12: MessageService.SaveNotifyItem:output_type -> Result
	5, // 13: MessageService.GetMailTemplate:output_type -> SMailTemplate
	7, // 14: MessageService.SaveMailTemplate:output_type -> Result
	9, // 15: MessageService.GetMailTemplates:output_type -> MailTemplateListResponse
	7, // 16: MessageService.DeleteMailTemplate:output_type -> Result
	7, // 17: MessageService.SendSiteMessage:output_type -> Result
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_service_proto_init() }
func file_message_service_proto_init() {
	if File_message_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_message_dto_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_service_proto_goTypes,
		DependencyIndexes: file_message_service_proto_depIdxs,
	}.Build()
	File_message_service_proto = out.File
	file_message_service_proto_rawDesc = nil
	file_message_service_proto_goTypes = nil
	file_message_service_proto_depIdxs = nil
}
