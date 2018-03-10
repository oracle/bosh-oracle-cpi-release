package action

// Info action handles the info() request
type Info struct{}

// lightStemcellFormat name
// Must match the name produced (and stored in the stemcell manifest) by bosh-oracle-light-stemcell-builder
const lightStemcellFormat string = "oracle-light"

//NewInfo creates a new Info instance
func NewInfo() Info {
	return Info{}
}

// Run returns information about the stemcell(s)
// supported by this CPI.
func (i Info) Run() (StemcellInfo, error) {

	return StemcellInfo{
		Formats: []string{
			lightStemcellFormat,
		},
	}, nil
}
