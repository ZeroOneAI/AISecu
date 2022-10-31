package mongodb

import "go.mongodb.org/mongo-driver/mongo/options"

type ControllerGen struct {
	dbProtocol         string
	dbHostname         string
	dbPort             string
	databaseName       string
	authCredential     *options.Credential
	dependencyEndpoint string
}

func (c *ControllerGen) WithDB(protocol, hostname, port, databaseName string) *ControllerGen {
	c.dbProtocol = protocol
	c.dbHostname = hostname
	c.dbPort = port
	c.databaseName = databaseName
	return c
}

func (c *ControllerGen) WithDBAuth(credential *options.Credential) *ControllerGen {
	c.authCredential = credential
	return c
}

func (c *ControllerGen) WithDependencyManager(endpoint string) *ControllerGen {
	c.dependencyEndpoint = endpoint
	return c
}

func (c *ControllerGen) Create() (*Controller, error) {
	return NewController(c.dbProtocol, c.dbHostname, c.dbPort, c.databaseName, c.authCredential, c.dependencyEndpoint)
}
