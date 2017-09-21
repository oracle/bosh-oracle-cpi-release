package disks

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"

	"oracle/oci/core/client/blockstorage"
	"oracle/oci/core/models"
)

type diskCreator struct {
	connector client.Connector
	logger    boshlog.Logger
	location  resource.Location
}

func NewCreator(c client.Connector, l boshlog.Logger, loc resource.Location) Creator {
	return &diskCreator{connector: c, logger: l, location: loc}
}

type CreatorFactory func(client.Connector, boshlog.Logger, resource.Location) Creator

func (dc *diskCreator) CreateVolume(name string, sizeinMB int64) (*resource.Volume, error) {

	ad := dc.location.AvailabilityDomain()
	cid := dc.location.CompartmentID()
	details := models.CreateVolumeDetails{AvailabilityDomain: &ad, CompartmentID: &cid,
		DisplayName: name, SizeInMBs: sizeinMB}

	p := blockstorage.NewCreateVolumeParams().WithCreateVolumeDetails(&details)

	res, err := dc.connector.CoreSevice().Blockstorage.CreateVolume(p)

	if err != nil {
		dc.logger.Error(diskOperationsLogTag, "Error creating volume %v", err)
		return nil, err
	}

	var volume *resource.Volume
	if err = dc.waitUntilProvisioned(res.Payload, func(v *models.Volume) {
		volume = resource.NewVolume(*v.DisplayName, *v.ID)
	}); err != nil {
		return nil, err
	}
	dc.logger.Debug(diskOperationsLogTag, "Created volume %v", volume)
	return volume, nil
}

func (dc *diskCreator) waitUntilProvisioned(v *models.Volume, acceptor func(*models.Volume)) error {
	waiter := &volumeAvailableWaiter{connector: dc.connector, logger: dc.logger,
		availableHandler: acceptor}
	return waiter.WaitFor(v)
}
