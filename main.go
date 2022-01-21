package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/beiri22/terraform-provider-inwx/inwx"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: inwx.Provider,
	})
}
