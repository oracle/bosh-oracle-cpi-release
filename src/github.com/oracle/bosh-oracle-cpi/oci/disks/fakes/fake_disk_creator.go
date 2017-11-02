package fakes

import "github.com/oracle/bosh-oracle-cpi/oci/resource"

type FakeDiskCreator struct {
	CreateVolumeCalled   bool
	CreateVolumeLocation resource.Location
	CreateVolumeResult   *resource.Volume
	CreateVolumeSize     int64
	CreateVolumeError    error
}

func (f *FakeDiskCreator) CreateVolume(name string, sizeinMB int64) (*resource.Volume, error) {

	f.CreateVolumeCalled = true
	f.CreateVolumeSize = sizeinMB
	return f.CreateVolumeResult, f.CreateVolumeError

}
