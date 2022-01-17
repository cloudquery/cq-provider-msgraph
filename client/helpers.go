package client

import "github.com/cloudquery/cq-provider-sdk/provider/schema"

func DeleteTenantFilter(meta schema.ClientMeta, _ *schema.Resource) []interface{} {
	client := meta.(*Client)
	return []interface{}{"tenant_id", client.TenantId}
}
