package vm

import (
	"errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	"github.com/oracle/bosh-oracle-cpi/oci"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/models"
	"time"
)

const rateLimitErrorCode int = 429

type Updater interface {
	UpdateInstanceName(instanceID string, name string) error
}

type UpdaterFactory func(client.Connector, boshlog.Logger) Updater

type updater struct {
	connector client.Connector
	logger    boshlog.Logger

	updateParam *compute.UpdateInstanceParams
}

func NewUpdater(c client.Connector, l boshlog.Logger) Updater {
	return &updater{connector: c, logger: l}
}

func (u *updater) UpdateInstanceName(instanceID string, name string) error {

	updateDetails := models.UpdateInstanceDetails{
		DisplayName: name,
	}
	u.updateParam = compute.NewUpdateInstanceParams().WithInstanceID(instanceID).WithUpdateInstanceDetails(&updateDetails)

	strategy := boshretry.NewUnlimitedRetryStrategy(2*time.Second, boshretry.NewRetryable(u.tryUpdate), u.logger)
	return strategy.Try()
}

func (u *updater) tryUpdate() (bool, error) {
	_, err := u.connector.CoreSevice().Compute.UpdateInstance(u.updateParam)
	if err != nil {
		return shouldRetry(err)
	}
	return true, nil
}

func shouldRetry(modelError error) (bool, error) {
	if _, ok := modelError.(*compute.UpdateInstanceConflict); ok {
		return true, errors.New("Waiting for instance conflict to resolve")
	}
	u, ok := modelError.(*compute.UpdateInstanceDefault)
	if ok && u.Code() == rateLimitErrorCode {
		return true, errors.New("Waiting while being throttled")
	}
	return false, errors.New(oci.CoreModelErrorMsg(modelError))
}
