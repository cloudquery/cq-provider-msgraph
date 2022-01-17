package provider

import (
	"embed"
	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-msgraph/client"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//go:embed migrations/*.sql
	azureMigrations embed.FS
	Version         = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:     Version,
		Name:        "azure",
		Configure:   client.Configure,
		Migrations:  azureMigrations,
		ResourceMap: map[string]*schema.Table{},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
