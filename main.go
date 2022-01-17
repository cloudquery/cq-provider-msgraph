package main

import (
	"github.com/cloudquery/cq-provider-msgraph/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	p := provider.Provider()
	serve.Serve(&serve.Options{
		Name:                p.Name,
		Provider:            p,
		Logger:              nil,
		NoLogOutputOverride: false,
	})
}
