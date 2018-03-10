package action

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/oracle/bosh-oracle-cpi/oci/vm"
	"github.com/oracle/bosh-oracle-cpi/registry"
)

// Networks configured in the environment
type Networks map[string]*Network

// Network properties
type Network struct {
	Type            string                 `json:"type,omitempty"`
	IP              string                 `json:"ip,omitempty"`
	Gateway         string                 `json:"gateway,omitempty"`
	Netmask         string                 `json:"netmask,omitempty"`
	DNS             []string               `json:"dns,omitempty"`
	DHCP            bool                   `json:"use_dhcp,omitempty"`
	Default         []string               `json:"default,omitempty"`
	CloudProperties NetworkCloudProperties `json:"cloud_properties,omitempty"`
}

// AsRegistryNetworks converts the networks map to network settings
// structure expected by the agent registry
func (ns Networks) AsRegistryNetworks() registry.NetworksSettings {
	networksSettings := registry.NetworksSettings{}

	for netName, network := range ns {
		networksSettings[netName] = network.AsRegistryNetwork()
	}
	return networksSettings
}

// AsNetworkConfiguration converts the networks map to vm.Networks
// suitable for use with the vm.Creator
func (ns Networks) AsNetworkConfiguration() vm.Networks {

	networks := []vm.NetworkConfiguration{}
	for _, n := range ns {
		networks = append(networks, vm.NetworkConfiguration{
			VcnName:    n.CloudProperties.VcnName,
			SubnetName: n.CloudProperties.SubnetName,
			PrivateIP:  n.IP})
	}
	return networks
}

// AsRegistryNetwork converts a single network to network setting structure
// expected by the agent registry
func (n Network) AsRegistryNetwork() registry.NetworkSetting {
	return registry.NetworkSetting{
		Type:          n.Type,
		IP:            n.IP,
		Gateway:       n.Gateway,
		Netmask:       n.Netmask,
		UseDHCP:       true,
		DNS:           n.DNS,
		Default:       n.Default,
		Resolved:      false,
		Preconfigured: true,
	}
}

// isDynamic returns true if the network is configured
// as a "dynamic" network
func (n Network) isDynamic() bool {
	// The network type property is not always known
	// (Our caller cli v2 or director gobbles it and doesn't pass it down to us).
	// So we guess -- dynamic configuration typically doesn't request a static private
	// IP and we use that as an indicator

	return n.Type == "dynamic" || n.IP == ""
}

// isStatic returns true if the network is configured
// as a "manual" network
func (n Network) isStatic() bool {
	return n.Type == "manual" || n.IP != ""
}

// isVip returns true if the network is configured
// as a "vip" network
func (n Network) isVip() bool {
	return n.Type == "vip"
}

// FirstDynamic returns the first "dynamic" network in the networks map.  It returns nil
// if none exist.
func (ns Networks) FirstDynamic() *Network {
	for _, n := range ns {
		if n.isDynamic() {
			return n
		}
	}
	return nil
}

// FirstStatic returns the first "manual" network in the networks map.  It returns nil
// if none exist.
func (ns Networks) FirstStatic() *Network {
	for _, n := range ns {
		if n.isStatic() {
			return n
		}
	}
	return nil
}

// AllDynamic returns true if all the configured networks are dynamic, otherwise false.
func (ns Networks) AllDynamic() bool {
	for _, n := range ns {
		if !n.isDynamic() {
			return false
		}
	}
	return true
}

// First returns the first network in the map. Returns nil if the map is empty
func (ns Networks) First() *Network {
	for _, n := range ns {
		return n
	}
	return nil
}

func (n *Network) validate() error {
	if n.CloudProperties.SubnetName == "" {
		return bosherr.Error("Missing subnet name from network definition ")
	}
	if n.CloudProperties.VcnName == "" {
		return bosherr.Error("Missing vcn name from network definition ")
	}
	return nil
}
