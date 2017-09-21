package disks

import (
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/models"

	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci"
)

type diskFinder struct {
	connector client.Connector
	logger    boshlog.Logger
	location  resource.Location
}

func NewFinder(c client.Connector, l boshlog.Logger, loc resource.Location) Finder {
	return &diskFinder{connector: c, logger: l, location: loc}
}

type FinderFactory func(client.Connector, boshlog.Logger, resource.Location) Finder

func (f *diskFinder) FindVolume(volumeID string) (*resource.Volume, error) {

	var volume *resource.Volume
	acceptor := func(v *models.Volume) {
		volume = resource.NewVolume(*v.DisplayName, *v.ID)
	}
	waiter := &volumeAvailableWaiter{connector: f.connector, logger: f.logger,
		availableHandler: acceptor}

	v, err := waiter.queryVolume(volumeID)
	if err != nil {
		return nil, err
	}
	if err = waiter.WaitFor(v); err != nil {
		return nil, err
	}
	return f.findAttachment(volume)

}

func (f *diskFinder) findAttachment(v *resource.Volume) (*resource.Volume, error) {

	p := compute.NewListVolumeAttachmentsParams()
	vid := v.ID()
	p.WithCompartmentID(f.location.CompartmentID()).WithVolumeID(&vid)

	res, err := f.connector.CoreSevice().Compute.ListVolumeAttachments(p)
	if err != nil {
		return v, err
	}
	if res.Payload != nil && len(res.Payload) > 0 {
		if attachment, ok := res.Payload[0].(*models.IScsiVolumeAttachment); ok {
			v.SetAttachment(attachment)
		}
	}
	return v, nil
}

func (f *diskFinder) FindAllAttachedVolumes(instanceID string) ([]*resource.Volume, error) {

	p := compute.NewListVolumeAttachmentsParams().WithInstanceID(&instanceID).WithCompartmentID(f.location.CompartmentID())

	res, err := f.connector.CoreSevice().Compute.ListVolumeAttachments(p)
	if err != nil {
		return nil, fmt.Errorf("Error finding attachments for instance %s. Reason: %s", instanceID, oci.CoreModelErrorMsg(err))
	}
	volumes := []*resource.Volume{}
	for _, va := range res.Payload {
		if sva, ok := va.(*models.IScsiVolumeAttachment); ok {
			v := resource.NewVolume(sva.DisplayName(), *sva.VolumeID())
			v.SetAttachment(sva)
			volumes = append(volumes, v)
		}
	}
	return volumes, nil
}
