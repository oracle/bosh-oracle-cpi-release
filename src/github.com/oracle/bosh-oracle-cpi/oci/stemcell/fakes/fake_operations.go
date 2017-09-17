package fakes

type FakeCreator struct {
	CreateStemcellCalled       bool
	CreateStemcellCalledWithID string

	CreateStemcellResult string
	CreateStemcellError  error
}

type FakeDestroyer struct {
	DestroyStemcellCalled bool
	DestroyStemcellError  error
}

func (f *FakeCreator) CreateStemcell(imageOCID string) (stemcellId string, err error) {
	f.CreateStemcellCalled = true
	f.CreateStemcellCalledWithID = imageOCID
	return f.CreateStemcellResult, f.CreateStemcellError
}

func (f *FakeDestroyer) DeleteStemcell(stemcellId string) (err error) {
	f.DestroyStemcellCalled = true
	return f.DestroyStemcellError
}
