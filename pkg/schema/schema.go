package schema

type Config struct {
	Ssh    SshConfig    `json:"ssh"`
	System SystemConfig `json:"system"`
}

type SshConfig struct {
	Username string `json:"user"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
}

type SystemConfig struct {
	CPU     float64 `json:"cpu"`
	RAM     float64 `json:"ram"`
	Storage float64 `json:"storage"`
}
