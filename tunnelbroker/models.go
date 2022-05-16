package tunnelbroker

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Tunnel is a specific instance of an HE tunnel.
type Tunnel struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Description types.String `tfsdk:"description"`
	ServerV4    types.String `tfsdk:"serverv4"`
	ServerV6    types.String `tfsdk:"serverv6"`
	ClientV4    types.String `tfsdk:"clientv4"`
	ClientV6    types.String `tfsdk:"clientv6"`
	Routed64    types.String `tfsdk:"routed64"`
	Routed48    types.String `tfsdk:"routed48"`
}
