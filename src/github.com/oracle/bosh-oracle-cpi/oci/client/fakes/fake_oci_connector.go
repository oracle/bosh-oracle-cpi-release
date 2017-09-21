package fakes

import (
	"github.com/oracle/bosh-oracle-cpi/config"
	"github.com/oracle/bosh-oracle-cpi/registry"
	cclient "oracle/oci/core/client"
)

type FakeConnector struct {
	AgentOptionsResult registry.AgentOptions
}

func (fc *FakeConnector) Connect() error {
	return nil
}

func (fc *FakeConnector) CoreSevice() *cclient.CoreServices {
	return nil
}

func (fc *FakeConnector) Tenancy() string {
	return "fake-tenancy"
}

func (fc *FakeConnector) CompartmentId() string {
	return "fake-compartment-id"
}

func (fc *FakeConnector) AuthorizedKeys() []string {
	return []string{"ssh-rsa-fake"}
}
func (fc *FakeConnector) AgentOptions() registry.AgentOptions {
	return fc.AgentOptionsResult
}

func (fc *FakeConnector) AgentRegistryEndpoint() string {
	return "fake-agent-registry-endpoint"
}

func (fc *FakeConnector) SSHTunnelConfig() config.SSHTunnel {
	return config.SSHTunnel{}
}

func (fc *FakeConnector) SSHConfig() config.SSHConfig {
	return config.SSHConfig{}
}
