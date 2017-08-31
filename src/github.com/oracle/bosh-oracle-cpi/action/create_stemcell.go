package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/bmc/client"
)

// CreateStemcell action handles the create_stemcell method invocation
type CreateStemcell struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewCreateStemcell creates a CreateStemcell instance
func NewCreateStemcell(c client.Connector, logger boshlog.Logger) CreateStemcell {
	return CreateStemcell{connector: c, logger: logger}
}

// Run extracts the image ocid from the properties and delegates the creation
// to a StemcellCreator
func (cs CreateStemcell) Run(stemcellPath string, cloudProps StemcellCloudProperties) (StemcellCID, error) {

	c := newStemcellCreator(cs.connector, cs.logger)
	id, err := c.CreateStemcell(cloudProps.ImageOCID)
	if err != nil {
		return "", bosherr.WrapError(err, "Error creating stemcell")
	}
	return StemcellCID(id), nil

}
