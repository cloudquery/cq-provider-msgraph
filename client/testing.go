package client

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
)

const TestTenantId = "testTenant"

func CreateTestClient(address string) http.Client {
	transport := http.Transport{
		DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, network, address)
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return http.Client{
		Transport: &transport,
	}
}
