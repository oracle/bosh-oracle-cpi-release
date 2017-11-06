package test

import (
	"github.com/oracle/bosh-oracle-cpi/oci/disks"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"testing"
)

type DiskFactoryFixture struct {
	Connection *ConnectionFixture
	creator    disks.Creator
	terminator disks.Terminator
	finder     disks.Finder
}

type DiskFixture struct {
	DiskFactory *DiskFactoryFixture
	volume      *resource.Volume
}

func NewDiskFactoryFixture() *DiskFactoryFixture {
	return &DiskFactoryFixture{Connection: NewConnectionFixture()}
}

func NewDiskFixture() *DiskFixture {
	return &DiskFixture{DiskFactory: NewDiskFactoryFixture()}
}

func (dff *DiskFactoryFixture) Setup(t *testing.T) error {
	dff.Connection.Setup(t)

	loc := resource.NewLocation(dff.Connection.AD(), dff.Connection.Connector().CompartmentId())

	dff.creator = disks.NewCreator(dff.Connection.Connector(), dff.Connection.Logger(), loc)
	dff.terminator = disks.NewTerminator(dff.Connection.Connector(), dff.Connection.Logger())
	dff.finder = disks.NewFinder(dff.Connection.Connector(), dff.Connection.Logger(), loc)

	return nil
}

func (dff *DiskFactoryFixture) TearDown(t *testing.T) error {
	return dff.Connection.TearDown(t)
}

func (dff *DiskFactoryFixture) Creator() disks.Creator {
	return dff.creator
}

func (dff *DiskFactoryFixture) Terminator() disks.Terminator {
	return dff.terminator
}
func (dff *DiskFactoryFixture) Finder() disks.Finder {
	return dff.finder
}

func (df *DiskFixture) Setup(t *testing.T) error {
	df.DiskFactory.Setup(t)
	v, err := df.DiskFactory.creator.CreateVolume("Disk-tests-fixture", 50*volume1GB)
	if err != nil {
		t.Fatalf("Create volume error %v", err)
		return err
	}
	df.volume = v
	return nil
}

func (df *DiskFixture) TearDown(t *testing.T) error {
	if err := df.DiskFactory.terminator.DeleteVolume(df.VolumeID()); err != nil {
		t.Fatalf("TearDown: Error deleting volme %v", err)
	}

	return df.DiskFactory.TearDown(t)
}

func (df *DiskFixture) VolumeID() string {
	return df.volume.ID()
}

func (df *DiskFixture) Volume() *resource.Volume {
	return df.volume
}
