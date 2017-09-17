package disks

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
)

type Creator interface {
	CreateVolume(name string, sizeinMB int64) (*resource.Volume, error)
}

type Terminator interface {
	DeleteVolume(volumeID string) error
}

type AttacherDetacher interface {
	AttachVolumeToInstance(v *resource.Volume, in *resource.Instance) error
	DetachVolumeFromInstance(v *resource.Volume) error
}

type Finder interface {
	FindVolume(volumeID string) (*resource.Volume, error)
	FindAllAttachedVolumes(instanceID string) ([]*resource.Volume, error)
}

const diskOperationsLogTag = "OCIDiskOperations"

type InstanceAttacherDetacherFactory func(*resource.Instance, client.Connector, boshlog.Logger) (AttacherDetacher, error)
type AttacherDetacherFactory func(c client.Connector, l boshlog.Logger, adm IscsiNodeAdministrator) (AttacherDetacher, error)

func NewAttacherDetacherForInstance(in *resource.Instance, c client.Connector, l boshlog.Logger) (AttacherDetacher, error) {

	var remoteIP string
	var err error

	if c.SSHConfig().UsePublicIP() {
		remoteIP, err = in.PublicIP(c, l)
	} else {
		remoteIP, err = in.PrivateIP(c, l)
	}
	if err != nil {
		l.Debug(diskOperationsLogTag, "Error getting IP %v", err)
		return nil, err
	}
	adm := NewRemoteIscsiNodeAdministrator(c.SSHConfig().RemoteUser(), remoteIP,
		c.SSHConfig().LocalIdentityKeyPath(), l)

	return NewAttacherDetacher(c, l, adm), nil

}
