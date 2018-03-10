package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
)

// GetDisks action handles the get_disks request
type GetDisks struct {
	connector client.Connector
	logger    boshlog.Logger
}

// NewGetDisks creates a GetDisks instance
func NewGetDisks(c client.Connector, l boshlog.Logger) GetDisks {
	return GetDisks{connector: c, logger: l}
}

// Run queries and returns the IDs of block volumes attached to the given vm
func (gd GetDisks) Run(vmCID VMCID) ([]string, error) {

	loc := resource.NewLocation("", gd.connector.CompartmentId())
	volumes, err := newDiskFinder(gd.connector, gd.logger, loc).FindAllAttachedVolumes(string(vmCID))

	if err != nil {
		return nil, bosherr.WrapError(err, "Error finding disks")
	}
	diskIds := []string{}
	for _, v := range volumes {
		diskIds = append(diskIds, v.ID())
	}
	return diskIds, nil
}
