package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"github.com/oracle/bosh-oracle-cpi/registry"
)

// DeleteVM action handles the delete_vm request
type DeleteVM struct {
	connector client.Connector
	logger    boshlog.Logger
	registry  registry.Client
}

// NewDeleteVM creates a DeleteVM instance
func NewDeleteVM(c client.Connector, l boshlog.Logger, r registry.Client) DeleteVM {
	return DeleteVM{connector: c, logger: l, registry: r}
}

// Run deletes the requested VM.  Prior to deleting a vm, it detaches any block volumes
// attached to that vm. It also deletes the vm information from he agent registry
// after the VM is deleted.  These operations are not atomic, i.e if the registry can't be updated
// Run will return an error, but deleted VM is not restored.
func (dv DeleteVM) Run(vmCID VMCID) (interface{}, error) {

	var vmID = string(vmCID)

	// Detach any volume attachments
	if err := dv.detachInstanceVolumes(vmID); err != nil {
		dv.logger.Debug(logTag, "Error detaching volumes %v. Ignored", err)
	}

	// Delete instance
	if err := newVMTerminator(dv.connector, dv.logger).TerminateInstance(vmID); err != nil {
		return nil, bosherr.WrapErrorf(err, "Error deleting instance %s", vmID)
	}

	// Cleanup agent registry
	if err := dv.registry.Delete(vmID); err != nil {
		return nil, bosherr.WrapErrorf(err, "Deleting vm '%s'", vmCID)
	}
	return nil, nil
}

func (dv DeleteVM) detachInstanceVolumes(instanceID string) error {

	vm, err := newVMFinder(dv.connector, dv.logger).FindInstance(instanceID)
	if err != nil {
		return err
	}

	ad, err := newAttacherDetacherForInstance(vm, dv.connector, dv.logger)
	if err != nil {
		return err
	}
	loc := resource.NewLocation("", "", "", dv.connector.CompartmentId())
	volumes, err := newDiskFinder(dv.connector, dv.logger, loc).FindAllAttachedVolumes(instanceID)
	if err != nil {
		return err
	}

	for _, v := range volumes {
		ad.DetachVolumeFromInstance(v)
	}
	return nil

}
