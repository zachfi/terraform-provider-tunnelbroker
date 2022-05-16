package tunnelbroker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceTunnelType struct{}

func (r resourceTunnelType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type: types.StringType,
				// When Computed is true, the provider will set value --
				// the user cannot define the value
				Computed: true,
			},
			"last_updated": {
				Type:     types.StringType,
				Computed: true,
			},
			"description": {
				Type:     types.StringType,
				Computed: true,
			},
			"serverv4": {
				Type:     types.StringType,
				Computed: true,
			},
			"serverv6": {
				Type:     types.StringType,
				Computed: true,
			},
			"clientv4": {
				Type:     types.StringType,
				Required: true,
			},
			"clientv6": {
				Type:     types.StringType,
				Computed: true,
			},
			"routed64": {
				Type:     types.StringType,
				Computed: true,
			},
			"routed48": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}, nil
}

func (r resourceTunnelType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceTunnel{
		p: *(p.(*provider)),
	}, nil
}

type resourceTunnel struct {
	p provider
}

func (r resourceTunnel) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
}

// Read resource information
func (r resourceTunnel) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state Tunnel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tunnelID := state.ID.Value

	// Get order current value
	tunnel, err := r.p.client.GetTunnel(tunnelID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading tunnel",
			"Could not read tunnelID "+tunnelID+": "+err.Error(),
		)
		return
	}

	state.ClientV4 = types.String{Value: tunnel.ClientV4}
	state.ClientV6 = types.String{Value: tunnel.ClientV6}
	state.Description = types.String{Value: tunnel.Description}
	state.Routed48 = types.String{Value: tunnel.Routed48}
	state.Routed64 = types.String{Value: tunnel.Routed64}
	state.ServerV4 = types.String{Value: tunnel.ServerV4}
	state.ServerV6 = types.String{Value: tunnel.ServerV6}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceTunnel) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Get plan values
	var plan Tunnel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get current state
	var state Tunnel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get order ID from state
	tunnelID := state.ID.Value

	// Update order by calling API
	err := r.p.client.UpdateTunnel(tunnelID, plan.ClientV4.Value)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error update order",
			"Could not update tunnelID "+tunnelID+": "+err.Error(),
		)
		return
	}

	state.ClientV4 = plan.ClientV4

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete resource
func (r resourceTunnel) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
}

// Import resource
func (r resourceTunnel) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
