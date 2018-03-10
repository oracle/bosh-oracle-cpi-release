package action

import (
	"fmt"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"github.com/oracle/bosh-oracle-cpi/registry"
)

const diskActionsLogTag = "diskActions"

// AttachDisk action handles the attach_disk request to attach
// a persistent disk to a vm instance
type AttachDisk struct {
	connector      client.Connector
	logger         boshlog.Logger
	registryClient registry.Client
}

// NewAttachDisk creates an AttachDisk instance
func NewAttachDisk(c client.Connector, l boshlog.Logger, r registry.Client) AttachDisk {
	return AttachDisk{connector: c, logger: l, registryClient: r}

}

// Run implements the Action handler
func (ad AttachDisk) Run(vmCID VMCID, diskCID DiskCID) (interface{}, error) {

	in, err := newVMFinder(ad.connector, ad.logger).FindInstance(string(vmCID))

	if err != nil {
		return nil, bosherr.WrapError(err, "Unable to find VM")
	}

	loc := resource.NewLocation(in.Location().AvailabilityDomain(), ad.connector.CompartmentId())

	attacher, err := newAttacherDetacherForInstance(in, ad.connector, ad.logger)

	if err != nil {
		return nil, bosherr.WrapError(err, "Error creating attacher")
	}

	vol, err := newDiskFinder(ad.connector, ad.logger, loc).FindVolume(string(diskCID))
	if err != nil {
		return nil, bosherr.WrapError(err, "Unable to find Volume")
	}

	err = attacher.AttachVolumeToInstance(vol, in)
	if err != nil || !vol.IsAttached() {
		if err == nil {
			err = fmt.Errorf("Volume not attached %v", *vol)
		}
		return nil, bosherr.WrapError(err, "Error attaching volume")
	}

	// Read VM agent settings
	agentSettings, err := ad.registryClient.Fetch(string(vmCID))
	if err != nil {
		return nil, bosherr.WrapErrorf(err, "Attaching disk '%s' to vm '%s'", diskCID, vmCID)
	}

	// Update VM agent settings
	newAgentSettings := agentSettings.AttachPersistentDisk(string(diskCID), vol.DevicePath())
	if err = ad.registryClient.Update(string(vmCID), newAgentSettings); err != nil {
		return nil, bosherr.WrapErrorf(err, "Attaching disk '%s' to vm '%s'", diskCID, vmCID)
	}
	return nil, nil
}
