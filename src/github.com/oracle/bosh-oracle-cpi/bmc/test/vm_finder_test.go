package test

import (
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"testing"
)

func Test_VmOpsFindInstanceAndIPs(t *testing.T) {

	state := NewVMFixture()
	state.Setup(t)
	defer state.TearDown(t)

	vmid := state.instance.ID()

	var fin *resource.Instance
	t.Run("Find Instance", func(t *testing.T) {
		in, err := state.Finder().FindInstance(vmid)
		if err != nil {
			t.Fatalf("Error finding vm %v", err)
		}
		assertEqual(t, vmid, in.ID(), "Mismatch in vm id")
		fin = in
	})
	t.Run("Find Instance public IP", func(t *testing.T) {
		ip, err := fin.PublicIP(state.Connector(), state.Logger())
		if err != nil {
			t.Fatalf("Error getting Public IP %v", err)
		}
		assertNotNil(t, ip, "Unexpected public IP")
	})
	t.Run("Find Instance private IP", func(t *testing.T) {
		ip, err := fin.PrivateIP(state.Connector(), state.Logger())
		if err != nil {
			t.Fatalf("Error getting Private IP %v", err)
		}
		assertNotNil(t, ip, "Unexpected private IP")
	})
}
