package fakes

type FakeDiskTerminator struct {
	DeleteVolumeCalled bool
	DeleteVolumeError  error
	DeleteVolumeID     string
}

func (f *FakeDiskTerminator) DeleteVolume(volumeID string) error {
	f.DeleteVolumeCalled = true
	f.DeleteVolumeID = volumeID
	return f.DeleteVolumeError
}
