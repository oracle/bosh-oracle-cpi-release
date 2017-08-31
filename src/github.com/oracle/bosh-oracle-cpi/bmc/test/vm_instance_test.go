package test

import (
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"github.com/oracle/bosh-oracle-cpi/bmc/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"testing"
)

//  Test_LocationOfFoundInstance verifies that
// find instance populates availability domain
// and compartment ID
func Test_VmOpsLocationOfFoundInstance(t *testing.T) {

	state := NewVMFixture()
	state.Setup(t)
	defer state.TearDown(t)

	vmid := state.instance.ID()

	finder := vm.NewFinder(state.Connector(), state.Logger())

	in, err := finder.FindInstance(vmid)
	if err != nil {
		t.Fatalf(" Find failure %v", err)
	}
	loc := in.Location()

	assertEqual(t, state.connectionFixture.AD(), loc.AvailabilityDomain(), "")
	assertEqual(t, state.Connector().CompartmentId(), loc.CompartmentID(), "")
}

func Test_VmOpsCreateInstanceWithUserData(t *testing.T) {
	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	var in *resource.Instance
	var err error
	creator := vm.NewCreator(state.Connector(),
		state.Logger(), state.VCN(),
		state.Subnet(), state.AD())
	terminator := vm.NewTerminator(state.Connector(), state.Logger())

	deleteInstance := func() {
		if err == nil && in != nil {
			terminator.TerminateInstance(in.ID())
		}
	}
	defer deleteInstance()

	icfg := vmStandard12config
	icfg.Name = "test-instance"
	icfg.ImageId = state.StemcellImageID()
	in, err = creator.CreateInstance(icfg,
		vm.InstanceMetadata{
			vm.NewSSHKeys(state.connector.AuthorizedKeys()),
			vm.NewUserData(icfg.Name,
				"http://127.0.0.1:6901",
				manualNetworkNoIp.DNS,
				registry.NetworksSettings{
					"test-network": manualNetworkNoIp,
				}),
		})

	if err != nil {
		t.Fatalf("Error creating instance %v", err)
	}
	if in == nil {
		t.Fatalf("Unexpected nil instance")
	}
}
