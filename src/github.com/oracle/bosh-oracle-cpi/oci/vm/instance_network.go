package vm

import (
	"fmt"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/virtual_network"

	"oracle/oci/core/models"
)

type NetworkConfiguration struct {
	VcnName    string
	SubnetName string
	PrivateIP  string

	// Queried/Cached id of Vcn
	vcnId string
	// Queried/Cached id of Subnet
	subnetId string
}

func (n *NetworkConfiguration) subnetID(connector client.Connector) (string, error) {

	if n.subnetId != "" {
		return n.subnetId, nil
	}
	_, err := n.vcnID(connector)
	if err != nil {
		return "", err
	}

	p := virtual_network.NewListSubnetsParams()
	p.WithCompartmentID(connector.CompartmentId()).WithVcnID(n.vcnId)
	response, err := connector.CoreSevice().VirtualNetwork.ListSubnets(p)
	if err != nil {
		return "", err
	}
	for _, s := range response.Payload {
		if s.DisplayName == n.SubnetName {
			return *s.ID, nil
		}
	}
	return "", fmt.Errorf("Unable to find ID of subnet %s", n.SubnetName)
}

// VcnID queries the OCID of a vcn from the compute service
func (n *NetworkConfiguration) vcnID(connector client.Connector) (string, error) {

	if n.vcnId != "" {
		return n.vcnId, nil
	}
	req := virtual_network.NewListVcnsParams()
	req.WithCompartmentID(connector.CompartmentId())
	res, err := connector.CoreSevice().VirtualNetwork.ListVcns(req)
	if err != nil {
		return "", err
	}

	for _, v := range res.Payload {
		if v.DisplayName == n.VcnName {
			n.vcnId = *v.ID
			return n.vcnId, nil
		}
	}
	return "", fmt.Errorf("Error finding VcnID of VCN %s", n.VcnName)
}

func (n *NetworkConfiguration) newCreateVnicDetail(connector client.Connector, vnicName string) (models.CreateVnicDetails, error) {

	s, err := n.subnetID(connector)
	if err != nil {
		return models.CreateVnicDetails{}, err
	}

	return models.CreateVnicDetails{
		PrivateIP:   n.PrivateIP,
		SubnetID:    &s,
		DisplayName: vnicName}, nil
}

func (n *NetworkConfiguration) validate() error {

	if n.VcnName == "" {
		return fmt.Errorf(" Missing VCN name")
	}
	if n.SubnetName == "" {
		return fmt.Errorf("Missing subnet name")
	}
	return nil
}
