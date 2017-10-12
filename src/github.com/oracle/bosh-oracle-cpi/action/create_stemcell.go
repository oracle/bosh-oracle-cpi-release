package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
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

// Run extracts the image URL from the properties and delegates to
// StemcellCreator for creating an image
func (cs CreateStemcell) Run(_ string, cloudProps StemcellCloudProperties) (StemcellCID, error) {

	if cloudProps.ImageSourceURL == "" {
		return "", bosherr.Error("ImageSourceURL must be specified in the stemcell manifest")
	}
	if cloudProps.ImageOCID != "" {
		return "", bosherr.Error("Image OCID light stemcells are not supported anymore. Please use a newer light stemcell")
	}

	c := newStemcellCreator(cs.connector, cs.logger)
	id, err := c.CreateStemcell(cloudProps.ImageSourceURL, cloudProps.Name+"-"+cloudProps.Version)
	if err != nil {
		return "", bosherr.WrapError(err, "Error creating stemcell")
	}
	return StemcellCID(id), nil

}
