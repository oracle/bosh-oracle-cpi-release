package fakes

type FakeFinder struct {
	FindStemcellCalled       bool
	FindStemcellCalledWithID string

	FindStemcellResult string
	FindStemcellError  error
}

type FakeCreator struct {
	CreateStemcellCalled              bool
	CreateStemcellCalledWithURL       string
	CreateStemcellCalledWithImageName string

	CreateStemcellResult string
	CreateStemcellError  error
}

type FakeDestroyer struct {
	DestroyStemcellCalled bool
	DestroyStemcellError  error
}

func (f *FakeCreator) CreateStemcell(imageSourceURL string, imageName string) (stemcellId string, err error) {
	f.CreateStemcellCalled = true
	f.CreateStemcellCalledWithURL = imageSourceURL
	f.CreateStemcellCalledWithImageName = imageName
	return f.CreateStemcellResult, f.CreateStemcellError
}

func (f *FakeDestroyer) DeleteStemcell(stemcellId string) (err error) {
	f.DestroyStemcellCalled = true
	return f.DestroyStemcellError
}

func (f *FakeFinder) FindStemcell(imageOCID string) (stemcellId string, err error) {
	f.FindStemcellCalled = true
	f.FindStemcellCalledWithID = imageOCID
	return f.FindStemcellResult, f.FindStemcellError
}
