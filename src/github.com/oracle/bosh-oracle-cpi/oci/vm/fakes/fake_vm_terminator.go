package fakes

type FakeVMTerminator struct {
	TerminateInstanceError  error
	TerminateInstanceCalled bool
	TerminatedInstance      string
}

func (f *FakeVMTerminator) TerminateInstance(instanceID string) error {
	f.TerminateInstanceCalled = true
	f.TerminatedInstance = instanceID
	return f.TerminateInstanceError
}
