package action

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/bmc/client"
)

// HasVM action handles the has_vm request
type HasVM struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewHasVM creates a HasVM instance
func NewHasVM(c client.Connector, l boshlog.Logger) HasVM {
	return HasVM{connector: c, logger: l}
}

// Run queries BMC to determine if the given vm exists
func (hv HasVM) Run(vmCID VMCID) (bool, error) {

	id := string(vmCID)
	instance, err := newVMFinder(hv.connector, hv.logger).FindInstance(id)
	return instance != nil, err
}
