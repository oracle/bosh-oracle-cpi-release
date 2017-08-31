package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
)

// DeleteDisk action handles the delete_disk request
type DeleteDisk struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewDeleteDisk creates a DeleteDisk instance
func NewDeleteDisk(c client.Connector, l boshlog.Logger) DeleteDisk {
	return DeleteDisk{connector: c, logger: l}
}

// Run deletes a previously created persistent block volume.
func (dd DeleteDisk) Run(diskCID DiskCID) (interface{}, error) {

	t := newDiskTerminator(dd.connector, dd.logger)

	if err := t.DeleteVolume(string(diskCID)); err != nil {
		return nil, bosherr.WrapError(err, "Error deleting disk")
	}
	return nil, nil
}
