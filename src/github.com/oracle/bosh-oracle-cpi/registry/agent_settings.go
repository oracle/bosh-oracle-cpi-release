package registry

import "encoding/json"

const defaultSystemDisk = "/dev/sda"

type agentSettingsResponse struct {
	Settings string `json:"settings"`
	Status   string `json:"status"`
}

// AgentSettings holds bosh-agent settings for a particular VM.
type AgentSettings struct {
	// Agent ID
	AgentID string `json:"agent_id"`

	// Blobstore settings
	Blobstore BlobstoreSettings `json:"blobstore"`

	// Disks settings
	Disks DisksSettings `json:"disks"`

	// Environment settings
	Env EnvSettings `json:"env"`

	// Mbus URI
	Mbus string `json:"mbus"`

	// Networks settings
	Networks NetworksSettings `json:"networks"`

	// List of NTP servers
	Ntp []string `json:"ntp"`

	// VM settings
	VM VMSettings `json:"vm"`
}

// BlobstoreSettings are the Blobstore settings for a particular VM.
type BlobstoreSettings struct {
	// Blobstore provider
	Provider string `json:"provider"`

	// Blobstore options
	Options map[string]interface{} `json:"options"`
}

// DisksSettings are the Disks settings for a particular VM.
type DisksSettings struct {
	// System disk
	System string `json:"system"`

	// Ephemeral disk
	Ephemeral string `json:"ephemeral"`

	// Persistent disk
	Persistent map[string]PersistentSettings `json:"persistent"`
}

// PersistentSettings are the Persistent Disk settings for a particular VM.
type PersistentSettings struct {
	// Persistent disk ID
	ID string `json:"id"`

	// Persistent disk path
	Path string `json:"path"`
}

// EnvSettings are the Environment settings for a particular VM.
type EnvSettings map[string]interface{}

// NetworksSettings are the Networks settings for a particular VM.
type NetworksSettings map[string]NetworkSetting

// NetworkSetting is a setting for one interface for a particular VM
type NetworkSetting struct {
	Type          string   `json:"type"`
	IP            string   `json:"ip"`
	Netmask       string   `json:"netmask"`
	Gateway       string   `json:"gateway"`
	Resolved      bool     `json:"resolved"`
	UseDHCP       bool     `json:"use_dhcp"`
	Default       []string `json:"default"`
	DNS           []string `json:"dns"`
	Mac           string   `json:"mac"`
	Preconfigured bool     `json:"preconfigured"`
}

// VMSettings are the VM settings for a particular VM.
type VMSettings struct {
	// VM name
	Name string `json:"name"`
}

// NewAgentSettings creates new agent settings for a VM.
func NewAgentSettings(agentID string, vmCID string, networksSettings NetworksSettings, env EnvSettings, agentOptions AgentOptions) AgentSettings {
	agentSettings := AgentSettings{
		AgentID: agentID,
		Disks: DisksSettings{
			System:     defaultSystemDisk,
			Persistent: map[string]PersistentSettings{},
		},
		Blobstore: BlobstoreSettings{
			Provider: agentOptions.Blobstore.Provider,
			Options:  agentOptions.Blobstore.Options,
		},
		Env:      EnvSettings(env),
		Mbus:     agentOptions.Mbus,
		Networks: networksSettings,
		Ntp:      agentOptions.Ntp,
		VM: VMSettings{
			Name: vmCID,
		},
	}

	return agentSettings
}

// AttachPersistentDisk updates the agent settings in order to add an attached persistent disk.
func (as AgentSettings) AttachPersistentDisk(diskID string, path string) AgentSettings {
	persistentDiskSettings := make(map[string]PersistentSettings)
	if as.Disks.Persistent != nil {
		persistentDiskSettings = as.Disks.Persistent
	}
	persistentDiskSettings[diskID] = PersistentSettings{
		ID:   diskID,
		Path: path,
	}
	as.Disks.Persistent = persistentDiskSettings

	return as
}

// ConfigureNetworks updates the agent settings with the networks settings.
func (as AgentSettings) ConfigureNetworks(networksSettings NetworksSettings) AgentSettings {
	as.Networks = networksSettings

	return as
}

// DetachPersistentDisk updates the agent settings in order to delete an attached persistent disk.
func (as AgentSettings) DetachPersistentDisk(diskID string) AgentSettings {
	persistenDiskSettings := as.Disks.Persistent
	delete(persistenDiskSettings, diskID)
	as.Disks.Persistent = persistenDiskSettings

	return as
}

// AsJSONString returns the agent settings as a JSON string
func (as AgentSettings) AsJSONString() (s string, err error) {
	b, err := json.Marshal(as)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
