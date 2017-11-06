package vm

import (
	"fmt"
)

type Networks []NetworkConfiguration

func (n Networks) primary() NetworkConfiguration {
	// Assume first network to be the primary
	return n[0]
}

func (n Networks) privateIPs() []string {
	ips := []string{}
	for _, ip := range n {
		if ip.PrivateIP != "" {
			ips = append(ips, ip.PrivateIP)
		}
	}
	return ips
}

func (n Networks) hasSecondaries() bool {
	return len(n) > 1
}

func (n Networks) validate() error {
	if len(n) == 0 {
		return fmt.Errorf("Invalid network configuration. Must have atleast one subnet.")
	}

	for _, n := range n {
		if err := n.validate(); err != nil {
			return err
		}
	}
	return nil
}

func (n Networks) secondaries() []NetworkConfiguration {
	return n[1:]
}
