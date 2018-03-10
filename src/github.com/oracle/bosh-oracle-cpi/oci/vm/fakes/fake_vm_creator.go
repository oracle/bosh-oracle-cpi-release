package fakes

import (
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"github.com/oracle/bosh-oracle-cpi/oci/vm"
)

type FakeVMCreator struct {
	Configuration vm.InstanceConfiguration
	Metadata      vm.InstanceMetadata

	CreateInstanceResult *resource.Instance
	CreateInstanceError  error
	CreateInstanceCalled bool
}

func (f *FakeVMCreator) CreateInstance(icfg vm.InstanceConfiguration,
	md vm.InstanceMetadata) (*resource.Instance, error) {

	f.CreateInstanceCalled = true
	f.Configuration = icfg
	f.Metadata = md
	return f.CreateInstanceResult, f.CreateInstanceError
}
