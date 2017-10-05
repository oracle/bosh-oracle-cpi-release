package action

// DiskCloudProperties holds the CPI specific disk properties
type DiskCloudProperties struct {
}

// Environment used to create an instance
type Environment map[string]interface{}

// NetworkCloudProperties holds the CPI specific network properties
// defined in cloud config
type NetworkCloudProperties struct {
	VcnName    string `json:"vcn,omitempty"`
	SubnetName string `json:"subnet_name,omitempty"`
}

// StemcellCloudProperties holds the CPI specific stemcell properties
// defined in stemcell's manifest
type StemcellCloudProperties struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	// Deprecated: Light stemcells with image-ocid attribute are deprecated and
	// will be removed prior to GA.  Use stemcells with image-source-url attribute
	// instead.
	ImageOCID      string `json:"image-ocid,omitempty"`
	ImageSourceURL string `json:"image-source-url,omitempty"`
}

// VMCloudProperties holds the CPI specific properties
// defined in cloud-config for creating a instance
type VMCloudProperties struct {
	Name               string `json:"name,omitempty"`
	Shape              string `json:"instance_shape"`
	AvailabilityDomain string `json:"availability_domain"`
}
