package fakes

import "github.com/oracle/bosh-oracle-cpi/oci/resource"

type FakeAttacherDetacher struct {
	AttachVolumeCalled bool
	AttachedVolume     *resource.Volume
	AttachedInstance   *resource.Instance
	AttachmentError    error

	DetachVolumeCalled bool
	DetachedVolume     *resource.Volume
	DetachmentError    error
}

func (f *FakeAttacherDetacher) AttachVolumeToInstance(v *resource.Volume, in *resource.Instance) error {
	f.AttachVolumeCalled = true
	f.AttachedVolume = v
	f.AttachedInstance = in
	return f.AttachmentError
}

func (f *FakeAttacherDetacher) DetachVolumeFromInstance(v *resource.Volume) error {
	f.DetachVolumeCalled = true
	f.DetachedVolume = v
	return f.DetachmentError
}
