package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/xaque208/terraform-provider-tunnelbroker/tunnelbroker"
)

func main() {
	providerserver.Serve(context.Background(), tunnelbroker.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/xaque208/tunnelbroker",
	})
}
