package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
)

// HasDisk action handles the has_disk request
type HasDisk struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewHasDisk creates a HasDisk instance
func NewHasDisk(c client.Connector, l boshlog.Logger) HasDisk {
	return HasDisk{connector: c, logger: l}
}

// Run queries BMC to determine if the given block volume exists
func (hd HasDisk) Run(diskCID DiskCID) (bool, error) {

	loc := resource.NewLocation("", "", "", hd.connector.CompartmentId())
	vol, err := newDiskFinder(hd.connector, hd.logger, loc).FindVolume(string(diskCID))
	if err != nil {
		return false, bosherr.WrapError(err, "Error finding disk")
	}
	return vol.ID() == string(diskCID), nil
}
