package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/yannh/terraform-provider-statuspage/statuspage"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: statuspage.Provider,
	})
}
