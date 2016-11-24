package consulcli

import (
	serf "github.com/hashicorp/serf/client"
)

type Client interface {
	// Get cluster members
	ListAllMembers() ([]serf.Member, error)
	ListMembers(map[string]string, string) ([]serf.Member, error)
}

type client struct {
	serf *serf.RPCClient
}

// NewConsulClient returns a Client interface for given Serf
func NewSerfClient(serf_address string) (Client, error) {
	serfcli, err := serf.NewRPCClient(serf_address)
	if err != nil {
		return nil, err
	}
	return &client{serf: serfcli}, nil
}

func (c *client) ListAllMembers() ([]serf.Member, error) {
	members, err := c.serf.Members()
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (c *client) ListMembers(tags map[string]string, status string) ([]serf.Member, error) {
	members, err := c.serf.MembersFiltered(tags, status, "")
	if err != nil {
		return nil, err
	}
	return members, nil
}
