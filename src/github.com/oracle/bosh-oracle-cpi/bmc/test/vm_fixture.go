package test

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"github.com/oracle/bosh-oracle-cpi/bmc/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"testing"
)

type VMFixture struct {
	connectionFixture *ConnectionFixture
	creator           vm.Creator
	finder            vm.Finder
	terminator        vm.Terminator
	instance          *resource.Instance
}

func NewVMFixture() *VMFixture {
	return &VMFixture{connectionFixture: NewConnectionFixture()}
}

func (vf *VMFixture) Setup(t *testing.T) error {
	vf.connectionFixture.Setup(t)

	vf.creator = vm.NewCreator(vf.Connector(),
		vf.connectionFixture.Logger(), vf.connectionFixture.VCN(),
		vf.connectionFixture.Subnet(), vf.connectionFixture.AD())
	vf.terminator = vm.NewTerminator(vf.Connector(), vf.Logger())
	vf.finder = vm.NewFinder(vf.Connector(), vf.Logger())

	icfg := vmStandard12config
	icfg.ImageId = vf.connectionFixture.StemcellImageID()

	agentSettings := registry.NewAgentSettings("vm-fixture-agent-id", icfg.Name,
		registry.NetworksSettings{"test-network": manualNetworkNoIp},
		nil,
		vf.connectionFixture.Connector().AgentOptions())

	in, err := vf.creator.CreateInstance(icfg,
		vm.InstanceMetadata{
			vm.NewSSHKeys(vf.connectionFixture.Connector().AuthorizedKeys()),
			vm.NewAgentSettingsMetadata(agentSettings),
		})

	if err != nil {
		t.Fatalf("Error creating instance %v", err)
	}
	vf.instance = in
	return nil
}

func (vf *VMFixture) TearDown(t *testing.T) error {
	vf.terminator.TerminateInstance(vf.instance.ID())
	return vf.connectionFixture.TearDown(t)
}

func (vf *VMFixture) Connector() client.Connector {
	return vf.connectionFixture.Connector()
}
func (vf *VMFixture) Logger() boshlog.Logger {
	return vf.connectionFixture.Logger()
}

func (vf *VMFixture) Instance() *resource.Instance {
	return vf.instance
}
func (vf *VMFixture) Creator() vm.Creator {
	return vf.creator
}

func (vf *VMFixture) Terminator() vm.Terminator {
	return vf.terminator
}

func (vf *VMFixture) Finder() vm.Finder {
	return vf.finder
}
