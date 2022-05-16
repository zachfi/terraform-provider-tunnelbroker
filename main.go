package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/xaque208/terraform-provider-tunnelbroker/tunnelbroker"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	providerserver.Serve(context.Background(), tunnelbroker.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/xaque208/tunnelbroker",
	})
}
