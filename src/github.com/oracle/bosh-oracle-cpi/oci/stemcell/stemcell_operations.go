package stemcell

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/models"
)

type stemcellOperations struct {
	connector client.Connector
	logger    boshlog.Logger
}

func (so stemcellOperations) DeleteStemcell(stemcellID string) error {

	cs := so.connector.CoreSevice()
	p := compute.NewDeleteImageParams().WithImageID(stemcellID)
	_, err := cs.Compute.DeleteImage(p)
	return err

}

func (so stemcellOperations) CreateStemcell(sourceURI string, customImageName string) (stemcellID string, err error) {

	cid := so.connector.CompartmentId()

	ci := models.CreateImageDetails{
		CompartmentID: &cid,
		DisplayName:   customImageName,
		ImageSourceDetails: &models.ImageSourceViaObjectStorageURIDetails{
			SourceURI: &sourceURI,
		},
	}

	cs := so.connector.CoreSevice()
	p := compute.NewCreateImageParams().WithCreateImageDetails(&ci)
	ok, err := cs.Compute.CreateImage(p)

	if err != nil {
		return "", fmt.Errorf("Unable to create image from source %s. Reason: %s", sourceURI, oci.CoreModelErrorMsg(err))
	}

	var image *models.Image
	waiter := imageAvailableWaiter{
		connector: so.connector,
		logger:    so.logger,
		imageProvisionedHandler: func(i *models.Image) {
			image = i
		},
	}

	if err = waiter.WaitFor(ok.Payload); err != nil {
		return "", err
	}
	return *image.ID, nil
}

func (so stemcellOperations) FindStemcell(imageOCID string) (stemcellID string, err error) {

	image, err := queryImage(so.connector, imageOCID)

	if err != nil {
		return "", err
	}
	return *image.ID, nil
}

func queryImage(connector client.Connector, imageOCID string) (*models.Image, error) {

	p := compute.NewGetImageParams().WithImageID(imageOCID)
	image, err := connector.CoreSevice().Compute.GetImage(p)
	if err != nil {
		return nil, fmt.Errorf("Error finding image %s. Reason:%s", imageOCID, oci.CoreModelErrorMsg(err))
	}
	return image.Payload, nil
}
