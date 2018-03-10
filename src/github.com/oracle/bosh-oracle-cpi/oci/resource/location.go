package resource

import (
	"fmt"
	"github.com/oracle/bosh-oracle-cpi/oci/client"

	"github.com/oracle/bosh-oracle-cpi/oci"
	"oracle/oci/core/client/compute"
	"oracle/oci/core/client/virtual_network"
	"oracle/oci/core/models"
)

type Location struct {
	availabilityDomain string
	compartmentId      string
}

func NewLocation(ad string, compartmentId string) Location {
	return Location{availabilityDomain: ad, compartmentId: compartmentId}
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
		private[i] = *v.PrivateIP
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
			instanceID, oci.CoreModelErrorMsg(err))
	}

	// Get the Vnic for each attachment
	if len(r.Payload) == 0 {
		return nil, fmt.Errorf("No Vnic Attachments found for VM %s", instanceID)
	}

	vnics := []*models.Vnic{}
	for _, attachment := range r.Payload {

		switch *attachment.LifecycleState {
		case "ATTACHED":
			req := virtual_network.NewGetVnicParams().WithVnicID(attachment.VnicID)
			res, err := connector.CoreSevice().VirtualNetwork.GetVnic(req)
			if err != nil {
				return nil, fmt.Errorf("Error finding Vnic for attachment %s. Reason:%s",
					*attachment.ID, oci.CoreModelErrorMsg(err))
			}
			vnics = append(vnics, res.Payload)

		case "DETACHED", "DETACHING", "ATTACHING":
		}
	}
	return vnics, nil
}

func (loc Location) CompartmentID() string {
	return loc.compartmentId
}

func (loc Location) AvailabilityDomain() string {
	return loc.availabilityDomain
}
