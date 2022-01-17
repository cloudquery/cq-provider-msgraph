package client

import (
	"context"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

// here usually you will have custom column resolvers
func ResolveAzureTannantId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("tenant_id", client.TenantId)
}
