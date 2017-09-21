package resource

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshretry "github.com/cloudfoundry/bosh-utils/retrystrategy"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/models"

	"errors"
	"fmt"
	"time"
)

type Volume struct {
	ocid string
	name string

	attachmentID   string
	attachmentIQN  string
	attachmentIP   string
	attachmentPort int64
	devicePath     string
}

func NewVolume(name string, ocid string) *Volume {
	return &Volume{ocid: ocid, name: name}
}

func (v *Volume) ID() string {
	return v.ocid
}

func (v *Volume) SetAttachmentID(id string) {
	v.attachmentID = id
}
func (v *Volume) AttachmentID() string {
	return v.attachmentID
}

func (v *Volume) SetAttachmentIQN(iqn string) {
	v.attachmentIQN = iqn
}
func (v *Volume) AttachmentIQN() string {
	return v.attachmentIQN
}

func (v *Volume) SetAttachmentIP(ip string) {
	v.attachmentIP = ip
}
func (v *Volume) AttachmentIP() string {
	return v.attachmentIP
}

func (v *Volume) SetAttachmentPort(port int64) {
	v.attachmentPort = port
}
func (v *Volume) AttachmentPort() int64 {
	return v.attachmentPort
}

func (v *Volume) detached() {
	v.attachmentID = ""
	v.devicePath = ""
}

func (v *Volume) IsAttached() bool {
	return v.attachmentID != "" && v.attachmentIQN != "" && v.devicePath != ""
}

func (v *Volume) EnsureAttached(c client.Connector, l boshlog.Logger,
	attachment *models.IScsiVolumeAttachment) error {

	v.attachmentID = *attachment.ID()
	getVolumeIqn := func() (bool, error) {
		l.Debug(logTag, "Attachment state is %s ", *attachment.LifecycleState())
		switch *attachment.LifecycleState() {
		case "ATTACHED":
			v.SetAttachment(attachment)
			return true, nil
		case "ATTACHING":
			var err error
			attachment, err = v.queryAttachment(c, l)
			if err != nil {
				return false, err
			}
			return true, errors.New("Attaching")
		}
		return true, errors.New(fmt.Sprintf("Invalid attachment state %s", *attachment.LifecycleState()))
	}
	retryable := boshretry.NewRetryable(getVolumeIqn)
	retryStrategy := boshretry.NewAttemptRetryStrategy(100, 1*time.Second, retryable, l)

	l.Debug(logTag, "Waiting for volume attachment to reach ATTACHED state...")
	if err := retryStrategy.Try(); err != nil {
		l.Debug(logTag, "Error waiting for attachment %v", err)
		return err
	}
	l.Debug(logTag, "Done")
	return nil
}

func (v *Volume) SetAttachment(a *models.IScsiVolumeAttachment) {
	v.attachmentIQN = *a.Iqn
	v.attachmentIP = *a.IPV4
	v.attachmentPort = *a.Port
	v.attachmentID = *a.ID()
}

func (v *Volume) queryAttachment(c client.Connector, l boshlog.Logger) (*models.IScsiVolumeAttachment, error) {

	l.Debug(logTag, "Querying attachment state...")
	p := compute.NewGetVolumeAttachmentParams().WithVolumeAttachmentID(v.attachmentID)

	res, err := c.CoreSevice().Compute.GetVolumeAttachment(p)
	if err != nil {
		return nil, err
	}

	attachment, ok := res.Payload.(*models.IScsiVolumeAttachment)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unexpected attachment type %T", res.Payload))
	}
	return attachment, nil
}

func (v *Volume) SetDevicePath(path string) {
	v.devicePath = path
}

func (v *Volume) DevicePath() string {
	return v.devicePath
}

func (v *Volume) EnsureDetached(c client.Connector, l boshlog.Logger) error {

	attachment, err := v.queryAttachment(c, l)
	if err != nil {
		return err
	}
	getDetachmentState := func() (bool, error) {

		l.Debug(logTag, "Attachment state is %s ", *attachment.LifecycleState())

		switch *attachment.LifecycleState() {
		case "DETACHED":
			v.detached()
			return true, nil
		case "DETACHING":
			var err error
			attachment, err = v.queryAttachment(c, l)
			if err != nil {
				return false, err
			}
			return true, errors.New("Detaching")
		}
		return true, errors.New(fmt.Sprintf("Invalid detachment state %s", *attachment.LifecycleState()))
	}

	retryable := boshretry.NewRetryable(getDetachmentState)
	retryStrategy := boshretry.NewUnlimitedRetryStrategy(1*time.Second, retryable, l)

	l.Debug(logTag, "Waiting for volume attachment to reach DETACHED state...")
	if err := retryStrategy.Try(); err != nil {
		l.Debug(logTag, "Error waiting for detachment %v", err)
		return err
	}
	l.Debug(logTag, "Done")
	return nil
}
