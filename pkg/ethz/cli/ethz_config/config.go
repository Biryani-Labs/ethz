package ethzconfig

import "github.com/Biryani-Labs/ethz/pkg/schema"

func SSHUpdateConfigHostname(hostname string, config *schema.Config) {
	config.Ssh.Hostname = hostname
}

func SSHUpdateConfigUsername(username string, config *schema.Config) {
	config.Ssh.Username = username
}

func SSHUpdateConfigPort(port string, config *schema.Config) {
	config.Ssh.Port = port
}

func CPUUpdateConfigHostname(cpu float64, config *schema.Config) {
	config.System.CPU = cpu
}

func RAMUpdateConfigHostname(ram float64, config *schema.Config) {
	config.System.RAM = ram
}

func StorageUpdateConfigHostname(storage float64, config *schema.Config) {
	config.System.Storage = storage
}
