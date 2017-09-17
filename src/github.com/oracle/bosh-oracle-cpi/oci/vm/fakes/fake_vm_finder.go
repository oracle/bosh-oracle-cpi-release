package fakes

import (
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
)

type FakeVMFinder struct {
	FindInstanceResult *resource.Instance
	FindInstanceID     string
	FindInstanceError  error
	FindInstanceCalled bool
}

func (f *FakeVMFinder) FindInstance(instanceID string) (*resource.Instance, error) {

	f.FindInstanceCalled = true
	f.FindInstanceID = instanceID
	return f.FindInstanceResult, f.FindInstanceError
}
