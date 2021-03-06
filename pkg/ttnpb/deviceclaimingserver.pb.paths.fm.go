// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var ClaimEndDeviceRequestFieldPathsNested = []string{
	"invalidate_authentication_code",
	"source_device",
	"source_device.authenticated_identifiers",
	"source_device.authenticated_identifiers.authentication_code",
	"source_device.authenticated_identifiers.dev_eui",
	"source_device.authenticated_identifiers.join_eui",
	"source_device.qr_code",
	"target_application_ids",
	"target_application_ids.application_id",
	"target_application_server_address",
	"target_application_server_id",
	"target_application_server_kek_label",
	"target_device_id",
	"target_net_id",
	"target_network_server_address",
	"target_network_server_kek_label",
}

var ClaimEndDeviceRequestFieldPathsTopLevel = []string{
	"invalidate_authentication_code",
	"source_device",
	"target_application_ids",
	"target_application_server_address",
	"target_application_server_id",
	"target_application_server_kek_label",
	"target_device_id",
	"target_net_id",
	"target_network_server_address",
	"target_network_server_kek_label",
}
var AuthorizeApplicationRequestFieldPathsNested = []string{
	"api_key",
	"application_ids",
	"application_ids.application_id",
}

var AuthorizeApplicationRequestFieldPathsTopLevel = []string{
	"api_key",
	"application_ids",
}
var ClaimEndDeviceRequest_AuthenticatedIdentifiersFieldPathsNested = []string{
	"authentication_code",
	"dev_eui",
	"join_eui",
}

var ClaimEndDeviceRequest_AuthenticatedIdentifiersFieldPathsTopLevel = []string{
	"authentication_code",
	"dev_eui",
	"join_eui",
}
