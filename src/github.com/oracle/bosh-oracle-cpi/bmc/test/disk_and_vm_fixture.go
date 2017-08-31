package test

import (
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"testing"
)

type DiskAndVMFixture struct {
	diskFixture *DiskFixture
	vmFixture   *VMFixture
}

func NewDiskAndVMFixture() *DiskAndVMFixture {
	return &DiskAndVMFixture{diskFixture: NewDiskFixture(), vmFixture: NewVMFixture()}
}

func (dv *DiskAndVMFixture) Setup(t *testing.T) error {

	dv.vmFixture.Setup(t)
	dv.diskFixture.Setup(t)
	return nil
}
func (dv *DiskAndVMFixture) TearDown(t *testing.T) error {

	dv.vmFixture.TearDown(t)
	dv.diskFixture.TearDown(t)
	return nil
}

func (dv *DiskAndVMFixture) Instance() *resource.Instance {
	return dv.vmFixture.Instance()
}

func (dv *DiskAndVMFixture) Volume() *resource.Volume {
	return dv.diskFixture.Volume()
}
