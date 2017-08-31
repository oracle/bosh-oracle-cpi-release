package disks

import (
	"errors"
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"
	"github.com/oracle/bosh-oracle-cpi/bmc/resource"
	"oracle/baremetal/core/client/compute"
	"oracle/baremetal/core/models"
)

type diskAttacherDetacher struct {
	connector client.Connector
	logger    boshlog.Logger
	iscsiAdm  IscsiNodeAdministrator
}

func NewAttacherDetacher(c client.Connector, l boshlog.Logger, adm IscsiNodeAdministrator) AttacherDetacher {
	return &diskAttacherDetacher{connector: c, logger: l, iscsiAdm: adm}
}

func (ad *diskAttacherDetacher) AttachVolumeToInstance(v *resource.Volume, in *resource.Instance) error {

	var details models.AttachVolumeDetails
	details = &models.AttachIScsiVolumeDetails{}

	id := in.ID()
	vid := v.ID()
	details.SetInstanceID(&id)
	details.SetVolumeID(&vid)

	ad.logger.Debug(diskOperationsLogTag, "Attaching Volume %s", vid)

	req := compute.NewAttachVolumeParams().WithAttachVolumeDetails(details)
	res, err := ad.connector.CoreSevice().Compute.AttachVolume(req)

	if err != nil {
		ad.logger.Error(diskOperationsLogTag, "Error attaching volume %v", err)
		return err
	}
	attachment := res.Payload
	iscsi, ok := attachment.(*models.IScsiVolumeAttachment)
	if !ok {
		return errors.New(fmt.Sprintf("Unexpected attachment type %v", attachment))
	}

	if err := v.EnsureAttached(ad.connector, ad.logger, iscsi); err != nil {
		ad.logger.Error(diskOperationsLogTag, "Error attaching volume %v", err)
		return err
	}

	return ad.connectAttachment(v)
}

func (ad *diskAttacherDetacher) DetachVolumeFromInstance(v *resource.Volume) error {

	req := compute.NewDetachVolumeParams().WithVolumeAttachmentID(v.AttachmentID())
	_, err := ad.connector.CoreSevice().Compute.DetachVolume(req)
	if err != nil {
		ad.logger.Error(diskOperationsLogTag, "Error detaching volume %v", err)
		return err
	}

	return v.EnsureDetached(ad.connector, ad.logger)
}

func (ad *diskAttacherDetacher) connectAttachment(v *resource.Volume) error {

	path, err := ad.iscsiAdm.RunAttachmentCommands(v.AttachmentIQN(), v.AttachmentIP(), v.AttachmentPort())
	if err != nil {
		return err
	}
	v.SetDevicePath(path)
	return nil

}
