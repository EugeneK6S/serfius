package serfcli

import (
	serf "github.com/hashicorp/serf/client"
)

type Client interface {
	// Get cluster members
	Close() error
	ListAllMembers() (*[]serf.Member, error)
	ListMembers(map[string]string, string, string) (*[]serf.Member, error)
	NodeLeave(string) error
}

type RPCClient struct {
	serf *serf.RPCClient
}

// NewConsulClient returns a Client interface for given Serf
func NewSerfClient(serf_address string) (Client, error) {
	serfcli, err := serf.NewRPCClient(serf_address)
	if err != nil {
		return nil, err
	}
	return &RPCClient{serf: serfcli}, nil
}

func (c *RPCClient) Close() error {
	err := c.serf.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *RPCClient) ListAllMembers() (*[]serf.Member, error) {
	members, err := c.serf.Members()
	if err != nil {
		return nil, err
	}
	return &members, nil
}

func (c *RPCClient) ListMembers(tags map[string]string, status string, name string) (*[]serf.Member, error) {
	members, err := c.serf.MembersFiltered(tags, status, name)
	if err != nil {
		return nil, err
	}
	return &members, nil
}

func (c *RPCClient) NodeLeave(node string) error {
	err := c.serf.ForceLeave(node)
	if err != nil {
		return err
	}
	return nil
}
