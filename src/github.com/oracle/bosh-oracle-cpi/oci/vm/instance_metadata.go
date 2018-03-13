package vm

import (
	"encoding/json"
	"github.com/oracle/bosh-oracle-cpi/registry"
	"strings"
)

// WARNING:
// Do not change these constant values
// without changing the agent.json file generation in the
// stemcell configuration.
const metadataAuthorizedKeys = "ssh_authorized_keys"
const metatdataAgentSettingsKeyName = "bosh_agent_settings"
const metaDataAgentUserDataKeyName = "bosh_agent_userdata"

type MetadataEntry interface {
	Key() string
	Value() string
}

type InstanceMetadata []MetadataEntry

// UserData is CPI representation of user data expected by bosh-agent.
// Bosh agent reads UserData from the metadata service configured on
// the agent. UserData mirrors
// https://godoc.org/github.com/cloudfoundry/bosh-agent/infrastructure#UserDataContentsType
type UserData struct {
	Server   UserDataServerName        `json:"server"`
	Registry UserDataRegistryEndpoint  `json:"registry"`
	DNS      UserDataDNSItems          `json:"dns,omitempty"`
	Networks registry.NetworksSettings `json:"networks"`
}

type UserDataServerName struct {
	Name string `json:"name"`
}

type UserDataRegistryEndpoint struct {
	Endpoint string `json:"endpoint"`
}

type UserDataDNSItems struct {
	NameServer []string `json:"nameserver,omitempty"`
}

type AgentSettingsMetaData struct {
	Settings registry.AgentSettings
}

type SSHKeys []string

func NewUserData(name string, registryEndpoint string, dnsNames []string, networks registry.NetworksSettings) MetadataEntry {
	return UserData{
		Server: UserDataServerName{
			Name: name,
		},
		Registry: UserDataRegistryEndpoint{
			Endpoint: registryEndpoint,
		},
		DNS: UserDataDNSItems{
			NameServer: dnsNames,
		},
		Networks: networks,
	}
}

func NewSSHKeys(in []string) MetadataEntry {
	return SSHKeys(in)
}

func NewAgentSettingsMetadata(settings registry.AgentSettings) MetadataEntry {
	return AgentSettingsMetaData{Settings: settings}
}

func (k SSHKeys) Key() string {
	return metadataAuthorizedKeys
}
func (k SSHKeys) Value() string {
	return strings.Join(k, "\n")
}
func (u UserData) Key() string {
	return metaDataAgentUserDataKeyName
}
func (u UserData) Value() string {

	s, _ := u.AsJSONString()
	return s
}

func (u UserData) AsJSONString() (s string, err error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (asm AgentSettingsMetaData) Key() string {
	return metatdataAgentSettingsKeyName
}
func (asm AgentSettingsMetaData) Value() string {
	s, _ := asm.Settings.AsJSONString()
	return s
}

func (md InstanceMetadata) AsMap() map[string]string {
	mdmap := make(map[string]string, len(md))
	for _, entry := range md {
		mdmap[entry.Key()] = entry.Value()
	}
	return mdmap
}
