package client

import (
	"context"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

// ResolveTenantId resolver tenant in which client is running
func ResolveTenantId(_ context.Context, meta schema.ClientMeta, r *schema.Resource, _ schema.Column) error {
	client := meta.(*Client)
	return r.Set("tenant_id", client.TenantId)
}
