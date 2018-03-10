package action

// DiskCID is a identifier for a persistent disk.
type DiskCID string

// StemcellCID is an identifier for an image
type StemcellCID string

// VMCID is an identifier for an instance
type VMCID string

// StemcellInfo contains info about stemcells
// supported by this CPI
type StemcellInfo struct {
	Formats []string `json:"stemcell_formats"`
}
