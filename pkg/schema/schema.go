package schema

type Config struct {
	Ssh SshConfig `json:"ssh"`
}

type SshConfig struct {
	Username string `json:"user"`
	Hostname string `json:"hostname"`
}
