package resource

import (
	"fmt"
	"github.com/oracle/bosh-oracle-cpi/bmc/client"

	"oracle/baremetal/core/client/compute"
	"oracle/baremetal/core/client/virtual_network"
	"oracle/baremetal/core/models"
)

type Location struct {
	vcnName            string
	subnetName         string
	availabilityDomain string
	compartmentId      string
}

func NewLocation(vcnName string, subnetName string, ad string, compartmentId string) Location {
	return Location{vcnName: vcnName, subnetName: subnetName, availabilityDomain: ad, compartmentId: compartmentId}
}

// VcnID queries the OCID of a vcn from the compute service
func (loc Location) VcnID(connector client.Connector) (string, error) {
	req := virtual_network.NewListVcnsParams()
	req.WithCompartmentID(loc.compartmentId)
	res, err := connector.CoreSevice().VirtualNetwork.ListVcns(req)
	if err != nil {
		return "", err
	}

	for _, v := range res.Payload {
		if v.DisplayName == loc.vcnName {
			return *v.ID, nil
		}
	}
	return "", fmt.Errorf("Error finding VcnID of VCN %s", loc.vcnName)
}

// SubnetID queries the OCID of the subnet from
// the compute service
func (loc Location) SubnetID(connector client.Connector) (string, error) {

	vcnID, err := loc.VcnID(connector)
	if err != nil {
		return "", err
	}

	p := virtual_network.NewListSubnetsParams()
	p.WithCompartmentID(loc.compartmentId).WithVcnID(vcnID)
	response, err := connector.CoreSevice().VirtualNetwork.ListSubnets(p)
	if err != nil {
		return "", err
	}
	for _, s := range response.Payload {
		if s.DisplayName == loc.subnetName {
			return *s.ID, nil
		}

	}
	return "", fmt.Errorf("Unable to find ID of subnet %s", loc.subnetName)
}

func (loc Location) instanceIPs(connector client.Connector, instanceID string) (
	publicip []string, privateip []string, err error) {

	vnics, err := loc.Vnics(connector, instanceID)
	if err != nil {
		return nil, nil, err
	}
	public := make([]string, len(vnics))
	private := make([]string, len(vnics))
	for i, v := range vnics {
		public[i] = v.PublicIP
		private[i] = v.PrivateIP
	}
	return public, private, nil
}

func (loc Location) Vnics(connector client.Connector, instanceID string) ([]*models.Vnic, error) {

	// Find all VnicAttachments associated with the given instance
	p := compute.NewListVnicAttachmentsParams()
	p.WithInstanceID(&instanceID).WithCompartmentID(loc.compartmentId).WithAvailabilityDomain(&loc.availabilityDomain)
	r, err := connector.CoreSevice().Compute.ListVnicAttachments(p)
	if err != nil {
		return nil, fmt.Errorf("Error finding VnicAttachments for instance %s, %v",
			instanceID, err)
	}

	// Get the Vnic for each attachment
	if len(r.Payload) == 0 {
		return nil, fmt.Errorf("No Vnic Attachments found for VM %s", instanceID)
	}

	vnics := make([]*models.Vnic, len(r.Payload))
	for i, attachment := range r.Payload {

		req := virtual_network.NewGetVnicParams().WithVnicID(*attachment.VnicID)
		res, err := connector.CoreSevice().VirtualNetwork.GetVnic(req)
		if err != nil {
			return nil, fmt.Errorf("Error finding Vnic for attachment %s, %v",
				*attachment.ID, err)
		}
		vnics[i] = res.Payload
	}
	return vnics, nil
}

func (loc Location) CompartmentID() string {
	return loc.compartmentId
}

func (loc Location) AvailabilityDomain() string {
	return loc.availabilityDomain
}

func (loc Location) Vcn() string {
	return loc.vcnName
}

func (loc Location) Subnet() string {
	return loc.subnetName
}
