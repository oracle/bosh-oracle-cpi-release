package action

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
)

// VMMetadata contains the BOSH metadata
// for a vm instance
type VMMetadata map[string]string

// SetVMMetadata action handles the set_vm_metadata request
type SetVMMetadata struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewSetVMMetadata creates a SetVMMetadata instance
func NewSetVMMetadata(c client.Connector, l boshlog.Logger) SetVMMetadata {
	return SetVMMetadata{connector: c, logger: l}
}

// Run updates the display name of the VM
func (sm SetVMMetadata) Run(vmCID VMCID, metadata VMMetadata) (interface{}, error) {

	if name, ok := metadata["name"]; ok {
		updater := newVMUpdater(sm.connector, sm.logger)
		return nil, updater.UpdateInstanceName(string(vmCID), name)
	}
	return nil, nil
}
