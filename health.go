package storageos

import (
	"context"
	"encoding/json"

	"github.com/storageos/go-api/netutil"
	"github.com/storageos/go-api/types"
)

var (
	// HealthAPIPrefix is a partial path to the HTTP endpoint.
	HealthAPIPrefix = "health"
)

// Health returns the health of the control plane server at a given url.
func (c *Client) Health(ctx context.Context, hostname string) (*types.HealthStatus, error) {
	address, err := netutil.AddressFromNode(hostname)
	if err != nil {
		return nil, err
	}
	resp, err := c.doAddress(ctx, "GET", address, doOptions{urlpath: HealthAPIPrefix})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	status := &types.HealthStatus{}
	if err := json.NewDecoder(resp.Body).Decode(status); err != nil {
		return nil, err
	}

	return status, nil
}

// CPHealth returns the health of the control plane server at a given url.
func (c *Client) CPHealth(ctx context.Context, hostname string) (*types.CPHealthStatus, error) {
	status, err := c.Health(ctx, hostname)
	if err != nil {
		return nil, err
	}
	return status.ToCPHealthStatus(), nil
}

// DPHealth returns the health of the data plane server at a given url.
func (c *Client) DPHealth(ctx context.Context, hostname string) (*types.DPHealthStatus, error) {
	status, err := c.Health(ctx, hostname)
	if err != nil {
		return nil, err
	}
	return status.ToDPHealthStatus(), nil
}
