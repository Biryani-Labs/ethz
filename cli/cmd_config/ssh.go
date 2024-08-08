package cmdconfig

import "github.com/Biryani-Labs/ezeth/pkg/schema"

type SSHConfig struct {
	Host     HostSSHConfig     `cmd:"host" help:"Configure SSH Host variable"`
	Username UsernameSSHConfig `cmd:"user" help:"Configure SSH Username variable"`
}

type HostSSHConfig struct {
	schema.CliBlueprintName
	Hostname string `arg:"" help:"Hostname for SSH connection"`
}

type UsernameSSHConfig struct {
	schema.CliBlueprintName
	Username string `arg:"" help:"Username for SSH connection"`
}

func (host *HostSSHConfig) Run() error {
	return nil
}

func (username *UsernameSSHConfig) Run() error {
	return nil
}
