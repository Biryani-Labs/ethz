package ethzconfig

import "github.com/Biryani-Labs/ethz/pkg/schema"

func SSHUpdateConfigHostname(hostname string, config *schema.Config) {
	config.Ssh.Hostname = hostname
}

func SSHUpdateConfigUsername(username string, config *schema.Config) {
	config.Ssh.Username = username
}
