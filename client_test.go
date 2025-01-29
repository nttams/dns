package dns

import (
	"fmt"
	"net/netip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := NewClient("8.8.8.8:53")
	domains := []string{
		"google.com",
		"linkedin.com",
		"microsoft.com",
		"apple.com",
	}
	for _, domain := range domains {
		ips, err := client.Query(domain)
		assert.NoError(t, err)
		for _, ipStr := range ips {
			_, err := netip.ParseAddr(ipStr)
			assert.NoError(t, err)
		}
		assert.NotEmpty(t, ips, fmt.Sprintf("domain %s not found", domain))
	}
}
