package vm

import (
	"errors"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"time"
)

type instanceTerminatedWaiter struct {
	connector client.Connector
	logger    boshlog.Logger

	deletedHandler func(id string)
}

func (w *instanceTerminatedWaiter) WaitFor(ocid string) (err error) {

	getInstanceState := func() (bool, error) {

		p := compute.NewGetInstanceParams().WithInstanceID(ocid)
		r, err := w.connector.CoreSevice().Compute.GetInstance(p)
		if err != nil {
			return false, err
		}
		instance := r.Payload
		switch *instance.LifecycleState {
		case "STARTING", "PROVISIONING", "CREATING_IMAGE", "STOPPING", "STOPPED", "TERMINATING":
			return true, fmt.Errorf("Waiting. Current state %s", *instance.LifecycleState)
		case "TERMINATED":
			if w.deletedHandler != nil {
				w.deletedHandler(ocid)
			}
			return true, nil
		}
		return false, errors.New("Unknown instance state")
	}

	retryable := boshretry.NewRetryable(getInstanceState)
	retryStrategy := boshretry.NewAttemptRetryStrategy(12, 10*time.Second, retryable, w.logger)

	w.logger.Debug(logTag, "Waiting for instance to reach TERMINATED state...")
	if err := retryStrategy.Try(); err != nil {
		w.logger.Debug(logTag, "Error waiting for instance to reach TERMINATED state %v. Giving up", err)
		return err
	}
	w.logger.Debug(logTag, "Reached TERMINATED state")
	return nil
}
