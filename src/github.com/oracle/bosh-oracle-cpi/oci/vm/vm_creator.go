package vm

import (
	"fmt"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"github.com/oracle/bosh-oracle-cpi/oci/resource"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/models"
)

const logTag = "VMOperations"

type InstanceConfiguration struct {
	ImageId   string
	Shape     string
	Name      string
	PrivateIP string
}

type Creator interface {
	CreateInstance(icfg InstanceConfiguration, md InstanceMetadata) (*resource.Instance, error)
}

type CreatorFactory func(client.Connector, boshlog.Logger, string, string, string) Creator

type creator struct {
	connector client.Connector
	logger    boshlog.Logger
	location  resource.Location
}

func NewCreator(c client.Connector, l boshlog.Logger, vcnName string,
	subnetName string, availabilityDomain string) Creator {

	return &creator{connector: c, logger: l,
		location: resource.NewLocation(vcnName, subnetName, availabilityDomain, c.CompartmentId()),
	}
}

func (cv *creator) CreateInstance(icfg InstanceConfiguration,
	md InstanceMetadata) (*resource.Instance, error) {

	return cv.launchInstance(icfg.Name, icfg.PrivateIP, icfg.ImageId, icfg.Shape, md)
}

func (cv *creator) launchInstance(name string, assignIP string, imageId string, shape string,
	md InstanceMetadata) (*resource.Instance, error) {

	instance := resource.Instance{}
	var assignedIPs []string

	subnetId, err := cv.location.SubnetID(cv.connector)
	if err != nil {
		return nil, err
	}

	req := cv.buildLaunchInstanceParams(name, shape, imageId, subnetId, md.AsMap())
	if assignIP != "" {
		req = cv.populateCreateVnicDetails(req, assignIP)
		assignedIPs = []string{assignIP}
	}

	cv.logLaunchingInstanceDebugMsg(req)
	res, err := cv.connector.CoreSevice().Compute.LaunchInstance(req)
	if err != nil {
		errMsg := extractMsgFromError(err)
		cv.logger.Error(logTag, "Error launching instance. Err:%v Reason: %s", err, errMsg)
		return &instance, fmt.Errorf("Error launching instance. Reason: %s", errMsg)
	}

	return resource.NewInstanceWithPrivateIPs(*res.Payload.ID, cv.location, assignedIPs), nil
}

func (cv *creator) buildLaunchInstanceParams(name string, shape string,
	imageId string, subnetId string, metadata map[string]string) *compute.LaunchInstanceParams {
	req := compute.NewLaunchInstanceParams()
	ad := cv.location.AvailabilityDomain()
	cid := cv.location.CompartmentID()
	details := models.LaunchInstanceDetails{
		AvailabilityDomain: &ad,
		DisplayName:        name,
		CompartmentID:      &cid,
		Shape:              &shape,
		ImageID:            &imageId,
		SubnetID:           subnetId,
	}
	details.Metadata = metadata

	return req.WithLaunchInstanceDetails(&details)
}

func (cv *creator) populateCreateVnicDetails(param *compute.LaunchInstanceParams,
	privateIP string) *compute.LaunchInstanceParams {

	vnicDetails := models.CreateVnicDetails{
		HostnameLabel: param.LaunchInstanceDetails.HostnameLabel,
		PrivateIP:     privateIP,
		SubnetID:      &param.LaunchInstanceDetails.SubnetID,
		DisplayName:   "bosh-assigned"}

	param.LaunchInstanceDetails.CreateVnicDetails = &vnicDetails
	return param
}

func extractMsgFromError(err error) string {
	return oci.CoreModelErrorMsg(err)
}

func (cv *creator) logLaunchingInstanceDebugMsg(p *compute.LaunchInstanceParams) {

	fmtStr := "LaunchInstance: AD:%s, Name:%s, Shape:%s\nCompartmentId:%s\nImageId=%s\n"
	args := []interface{}{*p.LaunchInstanceDetails.AvailabilityDomain,
		p.LaunchInstanceDetails.DisplayName,
		*p.LaunchInstanceDetails.Shape,
		*p.LaunchInstanceDetails.CompartmentID,
		*p.LaunchInstanceDetails.ImageID,
	}
	if p.LaunchInstanceDetails.CreateVnicDetails != nil {
		fmtStr += "Subnet:%s, IP:%s\n"
		args = append(args, *p.LaunchInstanceDetails.CreateVnicDetails.SubnetID,
			p.LaunchInstanceDetails.CreateVnicDetails.PrivateIP)

	}
	fmtStr += "Metadata:\n\t%v\n"
	args = append(args, p.LaunchInstanceDetails.Metadata)

	cv.logger.Debug(logTag, fmtStr, args...)
}
