package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"oracle/baremetal/core/client/compute"
)

type Finder interface {
	FindInstance(instanceID string) (*resource.Instance, error)
}

type FinderFactory func(client.Connector, boshlog.Logger) Finder

type finder struct {
	connector client.Connector
	logger    boshlog.Logger
}

func NewFinder(c client.Connector, l boshlog.Logger) Finder {
	return &finder{connector: c, logger: l}
}

func (f *finder) FindInstance(instanceID string) (*resource.Instance, error) {

	f.logger.Debug(logTag, "Looking up details of VM %s", instanceID)
	p := compute.NewGetInstanceParams().WithInstanceID(instanceID)

	r, err := f.connector.CoreSevice().Compute.GetInstance(p)
	if err != nil {
		f.logger.Debug(logTag, "Error GetInstance %s", err.Error())
		return nil, err
	}
	loc := resource.NewLocation("", "", *r.Payload.AvailabilityDomain,
		*r.Payload.CompartmentID)
	return resource.NewInstance(*r.Payload.ID, loc), nil

}
