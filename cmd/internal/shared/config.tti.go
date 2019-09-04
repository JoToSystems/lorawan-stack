// Copyright © 2019 The Things Industries B.V.

package shared

import "go.thethings.network/lorawan-stack/pkg/tenant"

// DefaultTenancyConfig is the default tenancy configuration.
var DefaultTenancyConfig = tenant.Config{
	DefaultID: "default",
}