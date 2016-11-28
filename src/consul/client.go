package consulcli

///  UNDER CONSTRUCTION, NOT USED FOR NOW ///

// import (
// 	"fmt"
// 	consul "github.com/hashicorp/consul/api"
// )

// type Client interface {
// 	// Get a Service from consul
// 	Service(string, string) ([]*consul.ServiceEntry, *consul.QueryMeta, error)

// 	// Register a service with local agent
// 	Register(string, string, int) error

// 	// Deregister a service with local agent
// 	Deregister(string) error

// 	// Get list of nodes in Consul cluster
// 	ListNodes() ([]*consul.Node, *consul.QueryMeta, error)

// 	// Get cluster members
// 	ListMembers() ([]*consul.AgentMember, error)
// }

// type client struct {
// 	consul *consul.Client
// }

// // NewConsulClient returns a Client interface for given Consul address
// func NewConsulClient(consul_address string) (Client, error) {
// 	config := consul.DefaultConfig()
// 	config.Address = consul_address
// 	new_cli, err := consul.NewClient(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &client{consul: new_cli}, nil
// }

// // Register a service with Consul local agent
// func (c *client) Register(name string, address string, port int) error {
// 	reg := &consul.AgentServiceRegistration{
// 		ID:      name,
// 		Address: address,
// 		Name:    name,
// 		Port:    port,
// 	}
// 	return c.consul.Agent().ServiceRegister(reg)
// }

// // Deregister a service with Consul local agent
// func (c *client) Deregister(id string) error {
// 	return c.consul.Agent().ServiceDeregister(id)
// }

// // Return Health ServiceEntry
// func (c *client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
// 	passingOnly := true
// 	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)

// 	if len(addrs) == 0 && err == nil {
// 		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
// 	}
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	return addrs, meta, nil
// }

// func (c *client) ListMembers() ([]*consul.AgentMember, error) {
// 	wan := false
// 	members, err := c.consul.Agent().Members(wan)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return members, nil
// }

// func (c *client) ListNodes() ([]*consul.Node, *consul.QueryMeta, error) {

// 	nodes, meta, err := c.consul.Catalog().Nodes(nil)

// 	if len(nodes) == 0 && err == nil {
// 		return nil, nil, fmt.Errorf("There are no nodes in cluster!")
// 	}

// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	return nodes, meta, nil
// }
