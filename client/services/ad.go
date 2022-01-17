package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
)

//go:generate mockgen -destination=./mocks/ad_applications.go -package=mocks . ADApplicationsClient
type ADApplicationsClient interface {
	List(ctx context.Context, filter string) (result graphrbac.ApplicationListResultPage, err error)
}

//go:generate mockgen -destination=./mocks/ad_groups.go -package=mocks . ADGroupsClient
type ADGroupsClient interface {
	List(ctx context.Context, filter string) (result graphrbac.GroupListResultPage, err error)
}

//go:generate mockgen -destination=./mocks/ad_service_principals.go -package=mocks . ADServicePrinicpals
type ADServicePrinicpals interface {
	List(ctx context.Context, filter string) (result graphrbac.ServicePrincipalListResultPage, err error)
}

//go:generate mockgen -destination=./mocks/ad_users.go -package=mocks . ADUsersClient
type ADUsersClient interface {
	List(ctx context.Context, filter string, expand string) (result graphrbac.UserListResultPage, err error)
}

type AD struct {
	Applications      ADApplicationsClient
	Groups            ADGroupsClient
	ServicePrincipals ADServicePrinicpals
	Users             ADUsersClient
}

func NewADClient(tenantId string, auth autorest.Authorizer) AD {
	apps := graphrbac.NewApplicationsClient(tenantId)
	apps.Authorizer = auth
	groups := graphrbac.NewGroupsClient(tenantId)
	groups.Authorizer = auth
	users := graphrbac.NewUsersClient(tenantId)
	users.Authorizer = auth
	sp := graphrbac.NewServicePrincipalsClient(tenantId)
	sp.Authorizer = auth
	return AD{
		Applications:      apps,
		Groups:            groups,
		ServicePrincipals: sp,
		Users:             users,
	}
}
