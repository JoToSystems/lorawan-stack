// Copyright © 2019 The Things Industries B.V.

package store

const (
	// NOTE: please keep this sorted
	allowRegistrationsField = "allow_registrations"
	billingField            = "billing"
	billingIdentifiersField = "billing_identifiers"
	configurationField      = "configuration"
	maxApplicationsField    = "max_applications"
	maxClientsField         = "max_clients"
	maxEndDevicesField      = "max_end_devices"
	maxGatewaysField        = "max_gateways"
	maxOrganizationsField   = "max_organizations"
	maxUsersField           = "max_users"
)

var entityQuotasFields = []string{
	maxApplicationsField,
	maxClientsField,
	maxEndDevicesField,
	maxGatewaysField,
	maxOrganizationsField,
	maxUsersField,
}
