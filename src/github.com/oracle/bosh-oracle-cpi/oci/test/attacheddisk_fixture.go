package test

import (
	"github.com/oracle/bosh-oracle-cpi/oci/disks"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"testing"
)

type AttachedDiskFixture struct {
	diskFixture *DiskFixture
	vmFixture   *VMFixture
	attacher    disks.AttacherDetacher
}

func NewAttachedDiskFixture() *AttachedDiskFixture {
	return &AttachedDiskFixture{diskFixture: NewDiskFixture(), vmFixture: NewVMFixture()}
}

func (dv *AttachedDiskFixture) Setup(t *testing.T) error {

	dv.vmFixture.Setup(t)
	dv.diskFixture.Setup(t)

	a, err := disks.NewAttacherDetacherForInstance(dv.Instance(),
		dv.diskFixture.DiskFactory.Connection.Connector(),
		dv.diskFixture.DiskFactory.Connection.Logger())
	if err != nil {
		t.Fatalf("Error creating attacher %v", err)
	}
	dv.attacher = a

	if err := dv.attacher.AttachVolumeToInstance(dv.Volume(), dv.Instance()); err != nil {
		t.Fatalf("Error attaching volume to instance %v", err)
	}
	return nil
}

func (dv *AttachedDiskFixture) TearDown(t *testing.T) error {

	dv.attacher.DetachVolumeFromInstance(dv.Volume())
	dv.vmFixture.TearDown(t)
	dv.diskFixture.TearDown(t)
	return nil
}

func (dv *AttachedDiskFixture) Instance() *resource.Instance {
	return dv.vmFixture.Instance()
}

func (dv *AttachedDiskFixture) Volume() *resource.Volume {
	return dv.diskFixture.Volume()
}
