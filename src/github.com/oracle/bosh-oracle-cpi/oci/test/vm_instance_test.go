package test

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"github.com/oracle/bosh-oracle-cpi/oci/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"testing"
	"sync"
)

type getIPsFunc func(client.Connector, boshlog.Logger) ([]string, error)

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
		state.Logger(), state.AD())
	terminator := vm.NewTerminator(state.Connector(), state.Logger())

	deleteInstance := func() {
		if err == nil && in != nil {
			terminator.TerminateInstance(in.ID())
		}
	}
	defer deleteInstance()

	icfg := state.DefaultInstanceConfiguration()
	icfg.Name = "test-instance-with-metadata"

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

func Test_VmOpsRecreateVMWithSameIP(t *testing.T) {
	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	// Creator and Terminator
	var in *resource.Instance
	var err error

	creator := vm.NewCreator(state.Connector(),
		state.Logger(), state.AD())
	terminator := vm.NewTerminator(state.Connector(), state.Logger())

	deleteInstance := func() {
		if err == nil && in != nil {
			terminator.TerminateInstance(in.ID())
			in = nil
		}
	}
	defer deleteInstance()

	// Create a VM
	icfg := state.DefaultInstanceConfiguration()
	icfg.Name = "instance-with-auto-assigned-ip"
	in, err = creator.CreateInstance(icfg, vm.InstanceMetadata{})

	if err != nil {
		t.Fatalf("Error creating initial instance. Err: %v", err)
	}

	// Record its IP
	var ip string
	ip, err = in.PrivateIP(state.connector, state.logger)
	if err != nil {
		t.Fatalf("Error obtaining private IP. Err: %v", err)
	}

	// Delete it
	deleteInstance()

	// Recreate With same IP
	defer deleteInstance()
	icfg = state.DefaultInstanceConfiguration()
	icfg.Name = "recreated-with-deleted-ip"
	icfg.Network[0].PrivateIP = ip
	in, err = creator.CreateInstance(icfg, vm.InstanceMetadata{})

	if err != nil {
		t.Fatalf("Error re-creating instance with same-ip. Err: %v", err)
	}
	if in == nil {
		t.Fatalf("Unexpected nil instance")
	}
}

func Test_VmOpsAttachMultipleVnics(t *testing.T) {

	state := NewConnectionFixture()
	state.Setup(t)
	defer state.TearDown(t)

	// Creator and Terminator
	var in *resource.Instance
	var err error

	creator := vm.NewCreator(state.Connector(),
		state.Logger(), state.AD())
	terminator := vm.NewTerminator(state.Connector(), state.Logger())

	deleteInstance := func() {
		if err == nil && in != nil {
			terminator.TerminateInstance(in.ID())
			in = nil
		}
	}
	defer deleteInstance()

	icfg := state.DefaultInstanceConfiguration()
	icfg.Network = append(icfg.Network,
		vm.NetworkConfiguration{VcnName: state.VCN(),
			SubnetName: state.Subnet2()})
	in, err = creator.CreateInstance(icfg, vm.InstanceMetadata{})

	if err != nil {
		t.Fatalf("Error creating instance %v", err)
	}
	ipResults := map[string]getIPsFunc{
		"privateIPs": in.PrivateIPs,
		"publicIPs":  in.PublicIPs,
	}

	for n, r := range ipResults {
		ips, err := r(state.connector, state.Logger())
		if err != nil {
			t.Fatalf("Error finding IPs %s", err.Error())
		}
		if len(ips) != 2 {
			t.Errorf(" Unexpected number of IPs retuned by %s. Expected 2. Actual %d", n, len(ips))
		}
	}
}

func Test_VmOpsUpdateInstanceName(t *testing.T) {

	state := NewVMFixture()
	state.Setup(t)
	defer state.TearDown(t)

	updateInstance(t, state, nil)
}

// Test_VmOpsUpdateMultipleInstancesConcurrently tests issue #21
func Test_VmOpsUpdateMultipleInstancesConcurrently(t *testing.T) {

	state := NewVMFixtures(10)
	state.Setup(t)
	defer state.TearDown(t)

	wg := sync.WaitGroup{}
	for _, vf := range state.Fixtures() {
		wg.Add(1)
		go updateInstance(t, vf, &wg)
	}
	wg.Wait()
}

func updateInstance(t *testing.T, f *VMFixture, wg *sync.WaitGroup) {

	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	instance := f.Instance()
	updater := vm.NewUpdater(f.Connector(), f.Logger())
	err := updater.UpdateInstanceName(instance.ID(), "test-vm-renamed")
	assertIsNil(t, err, fmt.Sprintf("Unexpected failure when updating instance %s.  Error = [%v]",
		instance.ID(), err))
}
