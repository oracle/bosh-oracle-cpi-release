package fakes

import "github.com/oracle/bosh-oracle-cpi/oci/resource"

type FakeDiskFinder struct {
	FindVolumeCalled bool
	FindVolumeID     string
	FindVolumeResult *resource.Volume
	FindVolumeError  error

	FindAllAttachedVolumesCalled bool
	FindAllAttachedInstanceID    string
	FindAllAttachedResult        []*resource.Volume
	FindAllAttachedError         error
}

func (f *FakeDiskFinder) FindVolume(volumeID string) (*resource.Volume, error) {
	f.FindVolumeCalled = true
	f.FindVolumeID = volumeID
	return f.FindVolumeResult, f.FindVolumeError
}

func (f *FakeDiskFinder) FindAllAttachedVolumes(instanceID string) ([]*resource.Volume, error) {
	f.FindAllAttachedVolumesCalled = true
	f.FindAllAttachedInstanceID = instanceID

	return f.FindAllAttachedResult, f.FindAllAttachedError
}
