package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"github.com/oracle/bosh-oracle-cpi/registry"
)

// DetachDisk action handles the detach_disk request to detach
// a persistent disk from a vm instance
type DetachDisk struct {
	connector      client.Connector
	logger         boshlog.Logger
	registryClient registry.Client
}

// NewDetachDisk creates a DetachDisk instance
func NewDetachDisk(c client.Connector, l boshlog.Logger, r registry.Client) DetachDisk {
	return DetachDisk{connector: c, logger: l, registryClient: r}
}

// Run detaches the given disk from the the given vm. It also updates the agent registry
// after the detachment is completed. An error is thrown in case the disk or vm is not found,
// there is a failure in detachment, or if the registry can't be updated successfully.
func (dd DetachDisk) Run(vmCID VMCID, diskCID DiskCID) (interface{}, error) {

	in, err := newVMFinder(dd.connector, dd.logger).FindInstance(string(vmCID))

	if err != nil {
		return nil, bosherr.WrapError(err, "Unable to find VM")
	}

	detacher, err := newAttacherDetacherForInstance(in, dd.connector, dd.logger)
	if err != nil {
		return nil, bosherr.WrapError(err, "Error creating detacher")
	}

	loc := resource.NewLocation("", "", in.Location().AvailabilityDomain(), dd.connector.CompartmentId())
	vol, err := newDiskFinder(dd.connector, dd.logger, loc).FindVolume(string(diskCID))
	if err != nil {
		return nil, bosherr.WrapError(err, "Unable to find Volume")
	}

	if err := detacher.DetachVolumeFromInstance(vol); err != nil {
		return nil, bosherr.WrapError(err, "Error detaching volume")
	}

	// Read VM agent settings
	agentSettings, err := dd.registryClient.Fetch(string(vmCID))
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Detaching disk '%s' from vm '%s", diskCID, vmCID)
	}

	// Update VM agent settings
	newAgentSettings := agentSettings.DetachPersistentDisk(string(diskCID))
	if err = dd.registryClient.Update(string(vmCID), newAgentSettings); err != nil {
		return nil, bosherr.WrapErrorf(err, "Detaching disk '%s' from vm '%s", diskCID, vmCID)
	}
	return nil, nil
}
