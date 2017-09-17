package vm

import (
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/baremetal/core/client/compute"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
)

type Terminator interface {
	TerminateInstance(instanceID string) error
}
type TerminatorFactory func(client.Connector, boshlog.Logger) Terminator

type terminator struct {
	connector client.Connector
	logger    boshlog.Logger
}

func NewTerminator(c client.Connector, l boshlog.Logger) Terminator {
	return &terminator{connector: c, logger: l}
}

func (t *terminator) TerminateInstance(instanceID string) error {

	t.logger.Info(logTag, "Deleting VM %s...", instanceID)

	p := compute.NewTerminateInstanceParams().WithInstanceID(instanceID)
	_, err := t.connector.CoreSevice().Compute.TerminateInstance(p)
	if err != nil {
		t.logger.Error(logTag, "Ignoring error deleting instance %s", err.Error())
	}
	t.logger.Info(logTag, "Done")
	return nil

}
