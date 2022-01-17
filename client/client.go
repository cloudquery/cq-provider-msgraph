package client

import (
	"context"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/yaegashi/msgraph.go/msauth"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
	"golang.org/x/oauth2"
)

type Client struct {
	graph    *msgraph.GraphServiceRequestBuilder
	logger   hclog.Logger
	TenantId string
}

func NewAzureClient(log hclog.Logger, tenantId string, graph *msgraph.GraphServiceRequestBuilder) *Client {
	return &Client{
		logger:   log,
		TenantId: tenantId,
		graph:    graph,
	}
}

func (c Client) Logger() hclog.Logger {
	return c.logger
}

// Client return msgraph client
func (c Client) Client() *msgraph.GraphServiceRequestBuilder {
	return c.graph
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	settings, err := auth.GetSettingsFromEnvironment()
	if err != nil {
		return nil, err
	}

	creds, err := settings.GetClientCredentials()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	m := msauth.NewManager()
	scopes := []string{msauth.DefaultMSGraphScope}
	ts, err := m.ClientCredentialsGrant(ctx, creds.TenantID, creds.ClientID, creds.ClientSecret, scopes)
	if err != nil {
		return nil, err
	}

	httpClient := oauth2.NewClient(ctx, ts)
	graphClient := msgraph.NewClient(httpClient)

	client := NewAzureClient(logger, creds.TenantID, graphClient)

	// Return the initialized client. It will be passed to your resources
	return client, nil
}
