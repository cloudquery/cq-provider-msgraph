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
	//	//go:embed migrations/*.sql //todo uncomment when first migration file added
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
			"ad.applications": ad.Applications(),
			"ad.groups":       ad.Groups(),
			"ad.users":        ad.Users(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}
}
