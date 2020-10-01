// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

var TenantFieldPathsNested = []string{
	"attributes",
	"billing",
	"billing.provider",
	"billing.provider.stripe",
	"billing.provider.stripe.customer_id",
	"billing.provider.stripe.plan_id",
	"billing.provider.stripe.subscription_id",
	"billing.provider.stripe.subscription_item_id",
	"capabilities",
	"configuration",
	"configuration.default_cluster",
	"configuration.default_cluster.is",
	"configuration.default_cluster.is.end_device_picture",
	"configuration.default_cluster.is.end_device_picture.disable_upload",
	"configuration.default_cluster.is.profile_picture",
	"configuration.default_cluster.is.profile_picture.disable_upload",
	"configuration.default_cluster.is.profile_picture.use_gravatar",
	"configuration.default_cluster.is.user_registration",
	"configuration.default_cluster.is.user_registration.admin_approval",
	"configuration.default_cluster.is.user_registration.admin_approval.required",
	"configuration.default_cluster.is.user_registration.contact_info_validation",
	"configuration.default_cluster.is.user_registration.contact_info_validation.required",
	"configuration.default_cluster.is.user_registration.invitation",
	"configuration.default_cluster.is.user_registration.invitation.required",
	"configuration.default_cluster.is.user_registration.invitation.token_ttl",
	"configuration.default_cluster.is.user_registration.password_requirements",
	"configuration.default_cluster.is.user_registration.password_requirements.max_length",
	"configuration.default_cluster.is.user_registration.password_requirements.min_digits",
	"configuration.default_cluster.is.user_registration.password_requirements.min_length",
	"configuration.default_cluster.is.user_registration.password_requirements.min_special",
	"configuration.default_cluster.is.user_registration.password_requirements.min_uppercase",
	"configuration.default_cluster.is.user_rights",
	"configuration.default_cluster.is.user_rights.create_applications",
	"configuration.default_cluster.is.user_rights.create_clients",
	"configuration.default_cluster.is.user_rights.create_gateways",
	"configuration.default_cluster.is.user_rights.create_organizations",
	"configuration.default_cluster.ns",
	"configuration.default_cluster.ns.cooldown_window",
	"configuration.default_cluster.ns.deduplication_window",
	"configuration.default_cluster.ns.dev_addr_prefixes",
	"configuration.default_cluster.ui",
	"configuration.default_cluster.ui.branding_base_url",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"ids.tenant_id",
	"max_applications",
	"max_clients",
	"max_end_devices",
	"max_gateways",
	"max_organizations",
	"max_users",
	"name",
	"state",
	"updated_at",
}

var TenantFieldPathsTopLevel = []string{
	"attributes",
	"billing",
	"capabilities",
	"configuration",
	"contact_info",
	"created_at",
	"description",
	"ids",
	"max_applications",
	"max_clients",
	"max_end_devices",
	"max_gateways",
	"max_organizations",
	"max_users",
	"name",
	"state",
	"updated_at",
}
var TenantsFieldPathsNested = []string{
	"tenants",
}

var TenantsFieldPathsTopLevel = []string{
	"tenants",
}
var GetTenantRequestFieldPathsNested = []string{
	"field_mask",
	"tenant_ids",
	"tenant_ids.tenant_id",
}

var GetTenantRequestFieldPathsTopLevel = []string{
	"field_mask",
	"tenant_ids",
}
var ListTenantsRequestFieldPathsNested = []string{
	"field_mask",
	"limit",
	"order",
	"page",
}

var ListTenantsRequestFieldPathsTopLevel = []string{
	"field_mask",
	"limit",
	"order",
	"page",
}
var CreateTenantRequestFieldPathsNested = []string{
	"initial_user",
	"tenant",
	"tenant.attributes",
	"tenant.billing",
	"tenant.billing.provider",
	"tenant.billing.provider.stripe",
	"tenant.billing.provider.stripe.customer_id",
	"tenant.billing.provider.stripe.plan_id",
	"tenant.billing.provider.stripe.subscription_id",
	"tenant.billing.provider.stripe.subscription_item_id",
	"tenant.capabilities",
	"tenant.configuration",
	"tenant.configuration.default_cluster",
	"tenant.configuration.default_cluster.is",
	"tenant.configuration.default_cluster.is.end_device_picture",
	"tenant.configuration.default_cluster.is.end_device_picture.disable_upload",
	"tenant.configuration.default_cluster.is.profile_picture",
	"tenant.configuration.default_cluster.is.profile_picture.disable_upload",
	"tenant.configuration.default_cluster.is.profile_picture.use_gravatar",
	"tenant.configuration.default_cluster.is.user_registration",
	"tenant.configuration.default_cluster.is.user_registration.admin_approval",
	"tenant.configuration.default_cluster.is.user_registration.admin_approval.required",
	"tenant.configuration.default_cluster.is.user_registration.contact_info_validation",
	"tenant.configuration.default_cluster.is.user_registration.contact_info_validation.required",
	"tenant.configuration.default_cluster.is.user_registration.invitation",
	"tenant.configuration.default_cluster.is.user_registration.invitation.required",
	"tenant.configuration.default_cluster.is.user_registration.invitation.token_ttl",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.max_length",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_digits",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_length",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_special",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_uppercase",
	"tenant.configuration.default_cluster.is.user_rights",
	"tenant.configuration.default_cluster.is.user_rights.create_applications",
	"tenant.configuration.default_cluster.is.user_rights.create_clients",
	"tenant.configuration.default_cluster.is.user_rights.create_gateways",
	"tenant.configuration.default_cluster.is.user_rights.create_organizations",
	"tenant.configuration.default_cluster.ns",
	"tenant.configuration.default_cluster.ns.cooldown_window",
	"tenant.configuration.default_cluster.ns.deduplication_window",
	"tenant.configuration.default_cluster.ns.dev_addr_prefixes",
	"tenant.configuration.default_cluster.ui",
	"tenant.configuration.default_cluster.ui.branding_base_url",
	"tenant.contact_info",
	"tenant.created_at",
	"tenant.description",
	"tenant.ids",
	"tenant.ids.tenant_id",
	"tenant.max_applications",
	"tenant.max_clients",
	"tenant.max_end_devices",
	"tenant.max_gateways",
	"tenant.max_organizations",
	"tenant.max_users",
	"tenant.name",
	"tenant.state",
	"tenant.updated_at",
}

var CreateTenantRequestFieldPathsTopLevel = []string{
	"initial_user",
	"tenant",
}
var UpdateTenantRequestFieldPathsNested = []string{
	"field_mask",
	"tenant",
	"tenant.attributes",
	"tenant.billing",
	"tenant.billing.provider",
	"tenant.billing.provider.stripe",
	"tenant.billing.provider.stripe.customer_id",
	"tenant.billing.provider.stripe.plan_id",
	"tenant.billing.provider.stripe.subscription_id",
	"tenant.billing.provider.stripe.subscription_item_id",
	"tenant.capabilities",
	"tenant.configuration",
	"tenant.configuration.default_cluster",
	"tenant.configuration.default_cluster.is",
	"tenant.configuration.default_cluster.is.end_device_picture",
	"tenant.configuration.default_cluster.is.end_device_picture.disable_upload",
	"tenant.configuration.default_cluster.is.profile_picture",
	"tenant.configuration.default_cluster.is.profile_picture.disable_upload",
	"tenant.configuration.default_cluster.is.profile_picture.use_gravatar",
	"tenant.configuration.default_cluster.is.user_registration",
	"tenant.configuration.default_cluster.is.user_registration.admin_approval",
	"tenant.configuration.default_cluster.is.user_registration.admin_approval.required",
	"tenant.configuration.default_cluster.is.user_registration.contact_info_validation",
	"tenant.configuration.default_cluster.is.user_registration.contact_info_validation.required",
	"tenant.configuration.default_cluster.is.user_registration.invitation",
	"tenant.configuration.default_cluster.is.user_registration.invitation.required",
	"tenant.configuration.default_cluster.is.user_registration.invitation.token_ttl",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.max_length",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_digits",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_length",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_special",
	"tenant.configuration.default_cluster.is.user_registration.password_requirements.min_uppercase",
	"tenant.configuration.default_cluster.is.user_rights",
	"tenant.configuration.default_cluster.is.user_rights.create_applications",
	"tenant.configuration.default_cluster.is.user_rights.create_clients",
	"tenant.configuration.default_cluster.is.user_rights.create_gateways",
	"tenant.configuration.default_cluster.is.user_rights.create_organizations",
	"tenant.configuration.default_cluster.ns",
	"tenant.configuration.default_cluster.ns.cooldown_window",
	"tenant.configuration.default_cluster.ns.deduplication_window",
	"tenant.configuration.default_cluster.ns.dev_addr_prefixes",
	"tenant.configuration.default_cluster.ui",
	"tenant.configuration.default_cluster.ui.branding_base_url",
	"tenant.contact_info",
	"tenant.created_at",
	"tenant.description",
	"tenant.ids",
	"tenant.ids.tenant_id",
	"tenant.max_applications",
	"tenant.max_clients",
	"tenant.max_end_devices",
	"tenant.max_gateways",
	"tenant.max_organizations",
	"tenant.max_users",
	"tenant.name",
	"tenant.state",
	"tenant.updated_at",
}

var UpdateTenantRequestFieldPathsTopLevel = []string{
	"field_mask",
	"tenant",
}
var GetTenantIdentifiersForEndDeviceEUIsRequestFieldPathsNested = []string{
	"dev_eui",
	"join_eui",
}

var GetTenantIdentifiersForEndDeviceEUIsRequestFieldPathsTopLevel = []string{
	"dev_eui",
	"join_eui",
}
var GetTenantIdentifiersForGatewayEUIRequestFieldPathsNested = []string{
	"eui",
}

var GetTenantIdentifiersForGatewayEUIRequestFieldPathsTopLevel = []string{
	"eui",
}
var GetTenantRegistryTotalsRequestFieldPathsNested = []string{
	"field_mask",
	"tenant_ids",
	"tenant_ids.tenant_id",
}

var GetTenantRegistryTotalsRequestFieldPathsTopLevel = []string{
	"field_mask",
	"tenant_ids",
}
var TenantRegistryTotalsFieldPathsNested = []string{
	"applications",
	"clients",
	"end_devices",
	"gateways",
	"organizations",
	"users",
}

var TenantRegistryTotalsFieldPathsTopLevel = []string{
	"applications",
	"clients",
	"end_devices",
	"gateways",
	"organizations",
	"users",
}