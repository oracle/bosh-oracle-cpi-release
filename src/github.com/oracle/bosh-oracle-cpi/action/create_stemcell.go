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
// either StemcellFinder or StemcellCreator for creating an image
func (cs CreateStemcell) Run(_ string, cloudProps StemcellCloudProperties) (StemcellCID, error) {

	if cloudProps.ImageOCID != "" {
		f := newStemcellFinder(cs.connector, cs.logger)
		id, err := f.FindStemcell(cloudProps.ImageOCID)
		if err != nil {
			return "", bosherr.WrapError(err, "Error creating stemcell")
		}
		return StemcellCID(id), nil
	}

	if cloudProps.ImageSourceURL != "" {
		c := newStemcellCreator(cs.connector, cs.logger)
		id, err := c.CreateStemcell(cloudProps.ImageSourceURL, cloudProps.Name+"-"+cloudProps.Version)
		if err != nil {
			return "", bosherr.WrapError(err, "Error creating stemcell")
		}
		return StemcellCID(id), nil

	}

	return "", bosherr.Error("OCI Image OCID or an ImageSourceURL must be specified in the stemcell manifest")
}
