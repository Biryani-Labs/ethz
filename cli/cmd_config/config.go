package cmdconfig

type ConfigCmd struct {
	SSH   SSHConfig   `cmd:"ssh" help:"Configure SSH settings"`
	Fetch FetchConfig `cmd:"fetch" help:"Fetches configuration for a blueprint from marketplace (configured for github)"`
}
