package action

type Info struct{}

// lightStemcellFormat name
// Must match the name produced (and stored in the stemcell manifest) by bosh-oracle-light-stemcell-builder
const lightStemcellFormat string = "oracle-light"

func NewInfo() Info {
	return Info{}
}

func (i Info) Run() (StemcellInfo, error) {

	return StemcellInfo{
		Formats: []string{
			lightStemcellFormat,
		},
	}, nil
}
