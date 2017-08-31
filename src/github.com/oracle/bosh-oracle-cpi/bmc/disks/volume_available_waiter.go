package disks

import (
	"errors"
	"fmt"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"oracle/baremetal/core/client/blockstorage"
	"oracle/baremetal/core/models"
	"time"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
)

type volumeAvailableWaiter struct {
	connector client.Connector
	logger    boshlog.Logger

	availableHandler func(*models.Volume)
}

func (w *volumeAvailableWaiter) WaitFor(vol *models.Volume) error {

	getVolumeState := func() (bool, error) {

		switch *vol.LifecycleState {
		case models.VolumeLifecycleStatePROVISIONING,
			models.VolumeLifecycleStateRESTORING:
			var err error
			vol, err = w.queryVolume(*vol.ID)
			if err != nil {
				return false, err
			}
			return true, errors.New("Waiting")

		case models.VolumeLifecycleStateAVAILABLE:
			w.availableHandler(vol)
			return true, nil

		case models.VolumeLifecycleStateTERMINATING,
			models.VolumeLifecycleStateTERMINATED:
			return false, VolumeTerminatedError{*vol.ID}

		case models.VolumeLifecycleStateFAULTY:
			return false, VolumeFaultyError{*vol.ID}
		default:
			return false, errors.New(fmt.Sprintf("Unknown volume lifecycle state %s", *vol.LifecycleState))

		}
	}

	retryable := boshretry.NewRetryable(getVolumeState)
	retryStrategy := boshretry.NewUnlimitedRetryStrategy(500*time.Millisecond, retryable, w.logger)

	w.logger.Debug(diskOperationsLogTag, "Waiting for volume to be provisioned...")
	if err := retryStrategy.Try(); err != nil {
		w.logger.Debug(diskOperationsLogTag, "Error waiting %v", err)
		return err
	}
	w.logger.Debug(diskOperationsLogTag, "Done")
	return nil

}

func (w *volumeAvailableWaiter) queryVolume(ocid string) (*models.Volume, error) {

	p := blockstorage.NewGetVolumeParams().WithVolumeID(ocid)
	res, err := w.connector.CoreSevice().Blockstorage.GetVolume(p)
	if err != nil {
		w.logger.Error(diskOperationsLogTag, "Error finding volume %v", err)
		return nil, err
	}
	return res.Payload, nil

}
