package test

import (
	"github.com/oracle/bosh-oracle-cpi/oci/disks"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"testing"
)

func Test_DiskOpsNewFactory(t *testing.T) {

	state := NewConnectionFixture()

	state.Setup(t)
	defer state.TearDown(t)

	loc := resource.NewLocation(state.VCN(), state.Subnet(), state.AD(), state.Connector().CompartmentId())

	assertNotNil(t, disks.NewCreator(state.Connector(), state.Logger(), loc), " Null creator")
	assertNotNil(t, disks.NewFinder(state.Connector(), state.Logger(), loc), " Null Finder")
	assertNotNil(t, disks.NewTerminator(state.Connector(), state.Logger()), " Null terminator")
}

func Test_DiskOpsCreateDeleteVolume(t *testing.T) {

	state := NewDiskFactoryFixture()
	state.Setup(t)
	defer state.TearDown(t)

	var newVolumeID string

	t.Run("Create", func(t *testing.T) {
		v, err := state.Creator().CreateVolume("test-volume", 50*volume1GB)
		if err != nil {
			t.Fatalf("Create volume %v", err)
		}
		assertNotNil(t, v.ID(), "Invalid volume ID")
		newVolumeID = v.ID()
	})
	t.Run("Delete", func(t *testing.T) {
		assertNotNil(t, newVolumeID, "Nil volume ID")
		err := state.Terminator().DeleteVolume(newVolumeID)
		if err != nil {
			t.Fatalf("Delete volume %v", err)
		}
	})

}

func Test_DiskOpsFindVolume(t *testing.T) {

	state := NewDiskFixture()
	state.Setup(t)
	defer state.TearDown(t)

	v, err := state.DiskFactory.Finder().FindVolume(state.VolumeID())
	if err != nil {
		t.Fatalf("Find volume %v", err)
	}
	assertNotNil(t, v.ID(), "Empty volume ID from find volume")
}

func Test_DiskOpsAttachDisk(t *testing.T) {

	state := NewDiskAndVMFixture()
	state.Setup(t)
	defer state.TearDown(t)

	attacher, err := disks.NewAttacherDetacherForInstance(state.Instance(),
		state.diskFixture.DiskFactory.Connection.Connector(),
		state.diskFixture.DiskFactory.Connection.Logger())

	if err != nil {
		t.Fatalf("Creating attacher %v", err)
	}

	// Capture instance and volume
	instance := state.Instance()
	volume := state.Volume()
	defer attacher.DetachVolumeFromInstance(volume)
	err = attacher.AttachVolumeToInstance(volume, instance)

	if err != nil {
		t.Fatalf("Attaching volume %v, %T", volume, err)
	}
	if !volume.IsAttached() {
		t.Fail()
	}

}

func Test_DiskOpsFindAttachedDisk(t *testing.T) {

	state := NewAttachedDiskFixture()
	state.Setup(t)
	defer state.TearDown(t)

	finder := state.diskFixture.DiskFactory.Finder()
	volumes, err := finder.FindAllAttachedVolumes(state.Instance().ID())
	if err != nil {
		t.Fatalf("FindAllAttachedVolumes %v", err)
	}

	if volumes == nil || len(volumes) != 1 || volumes[0] == nil {
		t.Fatalf("Unexpected volumes %v", volumes)
	}

	if volumes[0].ID() != state.Volume().ID() {
		t.Fatalf("FindAllAttachedVolumes failure. Expected %s Actual %s", state.Volume().ID(), volumes[0].ID())
	}

}
