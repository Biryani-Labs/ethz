package cmdconfig

import (
	"fmt"
	"path/filepath"

	"github.com/Biryani-Labs/ethz/common/logs"
	"github.com/Biryani-Labs/ethz/common/utils"
	"github.com/Biryani-Labs/ethz/config"
	"github.com/Biryani-Labs/ethz/constants"
	ethzconfig "github.com/Biryani-Labs/ethz/pkg/ethz/cli/ethz_config"
	"github.com/Biryani-Labs/ethz/pkg/schema"
)

type SSHConfig struct {
	Host     HostSSHConfig     `cmd:"host" help:"Configure SSH Host variable"`
	Username UsernameSSHConfig `cmd:"user" help:"Configure SSH Username variable"`
	Port     PortSSHConfig     `cmd:"host" help:"Configure Port Host variable"`
}

type HostSSHConfig struct {
	schema.CliBlueprintName
	Hostname string `arg:"" help:"Hostname for SSH connection"`
}

type PortSSHConfig struct {
	schema.CliBlueprintName
	Port string `arg:"" help:"PortNumber for SSH connection"`
}

type UsernameSSHConfig struct {
	schema.CliBlueprintName
	Username string `arg:"" help:"Username for SSH connection"`
}

func (host *HostSSHConfig) Run() error {
	configFile, err := utils.BlueprintReadJsonFile(filepath.Join(config.LocateInHomePath(host.BlueprintName), constants.BlueprintFile))
	if err != nil {
		return fmt.Errorf("unable to read the configuration file for SSH host: %w", err)
	}

	ethzconfig.SSHUpdateConfigHostname(host.Hostname, configFile)
	if err := utils.BlueprintWriteJsonFile(filepath.Join(config.LocateInHomePath(host.BlueprintName), constants.BlueprintFile), configFile); err != nil {
		logs.Error(err, "Unable to update the blueprint file")
	}
	logs.Info("SSH hostname successfully updated.")
	return nil
}

func (username *UsernameSSHConfig) Run() error {
	configFile, err := utils.BlueprintReadJsonFile(filepath.Join(config.LocateInHomePath(username.BlueprintName), constants.BlueprintFile))
	if err != nil {
		return fmt.Errorf("unable to read the configuration file for SSH username: %w", err)
	}

	ethzconfig.SSHUpdateConfigUsername(username.Username, configFile)
	if err := utils.BlueprintWriteJsonFile(filepath.Join(config.LocateInHomePath(username.BlueprintName), constants.BlueprintFile), configFile); err != nil {
		logs.Error(err, "Unable to update the blueprint file")
	}
	logs.Info("SSH username successfully updated.")
	return nil
}

func (port *PortSSHConfig) Run() error {
	configFile, err := utils.BlueprintReadJsonFile(filepath.Join(config.LocateInHomePath(port.BlueprintName), constants.BlueprintFile))
	if err != nil {
		return fmt.Errorf("unable to read the configuration file for SSH username: %w", err)
	}

	ethzconfig.SSHUpdateConfigPort(port.Port, configFile)
	if err := utils.BlueprintWriteJsonFile(filepath.Join(config.LocateInHomePath(port.BlueprintName), constants.BlueprintFile), configFile); err != nil {
		logs.Error(err, "Unable to update the blueprint file")
	}
	logs.Info("SSH username successfully updated.")
	return nil
}
