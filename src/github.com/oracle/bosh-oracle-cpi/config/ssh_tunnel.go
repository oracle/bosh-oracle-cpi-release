package config

// SSHTunnel holds the CPI configuration
// for creating a forward SSH tunnel
// to a newly created vm
type SSHTunnel struct {
	LocalPort   int    `json:"localPort"`
	UsePublicIP bool   `json:"usePublicIP,omitempty"`
	User        string `json:"user"`
	Duration    string `json:"duration"`
}

// IsConfigured determines if an SSHTunnel has
// been configured or not
func (c SSHTunnel) IsConfigured() bool {
	return c.LocalPort > 0
}
