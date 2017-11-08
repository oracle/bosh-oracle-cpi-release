package fakes

type FakeVMUpdater struct {
	UpdateInstanceID     string
	UpdatedName          string
	UpdateInstanceError  error
	UpdateInstanceCalled bool
}

func (f *FakeVMUpdater) UpdateInstanceName(instanceID string, name string) error {

	f.UpdateInstanceCalled = true
	f.UpdateInstanceID = instanceID
	f.UpdatedName = name
	return f.UpdateInstanceError
}
