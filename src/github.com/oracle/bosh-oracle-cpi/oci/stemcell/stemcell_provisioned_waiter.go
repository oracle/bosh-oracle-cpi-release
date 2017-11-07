package stemcell

import (
	"errors"
	"fmt"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/models"
	"time"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
)

// imageAvailableWaiter waits for a newly created
// image to be fully provisioned
type imageAvailableWaiter struct {
	connector client.Connector
	logger    boshlog.Logger

	imageProvisionedHandler func(image *models.Image)
}

func (w *imageAvailableWaiter) WaitFor(img *models.Image) error {

	getImageState := func() (bool, error) {

		switch *img.LifecycleState {
		case models.ImageLifecycleStatePROVISIONING,
			models.ImageLifecycleStateIMPORTING:
			var err error
			img, err = queryImage(w.connector, *img.ID)
			if err != nil {
				return false, err
			}
			return true, errors.New("Waiting")

		case models.ImageLifecycleStateAVAILABLE:
			w.imageProvisionedHandler(img)
			return true, nil

		case models.ImageLifecycleStateDELETED,
			models.ImageLifecycleStateDISABLED:
			return false, fmt.Errorf("Image %s is either disabled or terminated", *img.ID)
		default:
			return false, fmt.Errorf("Unknown image lifecycle state %s", *img.LifecycleState)

		}
	}

	retryable := boshretry.NewRetryable(getImageState)
	delay := 2 * time.Minute
	retryStrategy := boshretry.NewUnlimitedRetryStrategy(delay, retryable, w.logger)

	w.logger.Debug(stemCellLogTag, "Waiting for image import to complete...will check every %d minutes", int(delay.Minutes()))
	if err := retryStrategy.Try(); err != nil {
		w.logger.Debug(stemCellLogTag, "Error waiting %v", err)
		return err
	}
	w.logger.Debug(stemCellLogTag, "Done")
	return nil

}
