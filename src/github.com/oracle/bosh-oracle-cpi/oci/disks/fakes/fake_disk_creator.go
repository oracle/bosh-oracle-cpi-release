package fakes

import "github.com/oracle/bosh-oracle-cpi/oci/resource"

type FakeDiskCreator struct {
	CreateVolumeCalled   bool
	CreateVolumeLocation resource.Location
	CreateVolumeResult   *resource.Volume
	CreateVolumeError    error
}

func (f *FakeDiskCreator) CreateVolume(name string, sizeinMB int64) (*resource.Volume, error) {

	f.CreateVolumeCalled = true
	return f.CreateVolumeResult, f.CreateVolumeError

}
