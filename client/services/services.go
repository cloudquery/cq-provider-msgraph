package services

import "github.com/Azure/go-autorest/autorest"

type Services struct {
	AD AD
}

func InitServices(tenantId string, auth autorest.Authorizer) Services {
	return Services{
		AD: NewADClient(tenantId, auth),
	}
}
