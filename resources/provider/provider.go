package provider

import (
	"embed"
	"github.com/cloudquery/cq-provider-msgraph/resources/services/ad"

	// CHANGEME: change the following to your own package
	"github.com/cloudquery/cq-provider-msgraph/client"

	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

var (
	//	//go:embed migrations/*.sql //todo uncomment on first migration
	azureMigrations embed.FS
	Version         = "Development"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Version:    Version,
		Name:       "msgraph",
		Configure:  client.Configure,
		Migrations: azureMigrations,
		ResourceMap: map[string]*schema.Table{
			"ad.applications": ad.AdApplications(),
			"ad.groups":       ad.AdGroups(),
			"ad.users":        ad.AdUsers(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
