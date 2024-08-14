package cmdconfig

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Biryani-Labs/ethz/common/logs"
	"github.com/Biryani-Labs/ethz/common/ssh"
	"github.com/Biryani-Labs/ethz/common/utils"
	"github.com/Biryani-Labs/ethz/config"
	"github.com/Biryani-Labs/ethz/constants"
	ethzconfig "github.com/Biryani-Labs/ethz/pkg/ethz/cli/ethz_config"
	"github.com/Biryani-Labs/ethz/pkg/schema"
)

type SystemConfig struct {
	Cpu     CpuConfig         `cmd:"cpu" help:"Configure System CPU variable - vCPU"`
	Ram     RamConfig         `cmd:"ram" help:"Configure System RAM variable - GB"`
	Storage StorageConfig     `cmd:"storage" help:"Configure System Storage variable - GB (/)"`
	Fetch   SystemFetchConfig `cmd:"fetch" help:"Auto Configure and fetch the details from the server"`
}

type CpuConfig struct {
	schema.CliBlueprintName
	Cpu float64 `arg:"" help:"CPU Size for system Config"`
}

type RamConfig struct {
	schema.CliBlueprintName
	Ram float64 `arg:"" help:"RAM Size for system Config"`
}

type StorageConfig struct {
	schema.CliBlueprintName
	Storage float64 `arg:"" help:"Storage Size for system Config"`
}

type SystemFetchConfig struct {
	schema.CliBlueprintName
}

func updateBlueprintFile(blueprintName string, updateFunc func(*schema.Config) error) error {
	configPath := filepath.Join(config.LocateInHomePath(blueprintName), constants.BlueprintFile)
	configFile, err := utils.BlueprintReadJsonFile(configPath)
	if err != nil {
		return fmt.Errorf("unable to read the configuration file: %w", err)
	}

	if err := updateFunc(configFile); err != nil {
		return fmt.Errorf("failed to update configuration: %w", err)
	}

	if err := utils.BlueprintWriteJsonFile(configPath, configFile); err != nil {
		return fmt.Errorf("unable to update the blueprint file: %w", err)
	}

	return nil
}

func (cpu *CpuConfig) Run() error {
	err := updateBlueprintFile(cpu.BlueprintName, func(configFile *schema.Config) error {
		ethzconfig.CPUUpdateConfigHostname(cpu.Cpu, configFile)
		return nil
	})
	if err != nil {
		return err
	}
	logs.Info("CPU details successfully updated.")
	return nil
}

func (ram *RamConfig) Run() error {
	err := updateBlueprintFile(ram.BlueprintName, func(configFile *schema.Config) error {
		ethzconfig.RAMUpdateConfigHostname(ram.Ram, configFile)
		return nil
	})
	if err != nil {
		return err
	}
	logs.Info("RAM details successfully updated.")
	return nil
}

func (storage *StorageConfig) Run() error {
	err := updateBlueprintFile(storage.BlueprintName, func(configFile *schema.Config) error {
		ethzconfig.StorageUpdateConfigHostname(storage.Storage, configFile)
		return nil
	})
	if err != nil {
		return err
	}
	logs.Info("Storage details successfully updated.")
	return nil
}

func (fetch *SystemFetchConfig) Run() error {
	configPath := filepath.Join(config.LocateInHomePath(fetch.BlueprintName), constants.BlueprintFile)
	cfg, err := utils.BlueprintReadJsonFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read blueprint file: %w", err)
	}

	password := ""
	if config.ASK_PASS {
		password = utils.ReadPassword()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	client, err := ssh.GetSSHConnection(ctx, cfg.Ssh.Username, cfg.Ssh.Hostname, cfg.Ssh.Port, password)
	if err != nil {
		return fmt.Errorf("failed to establish SSH connection: %w", err)
	}
	defer client.Close()

	commands := []string{
		"grep -c processor /proc/cpuinfo",
		"free -h | awk '/^Mem:/{print substr($2, 1, length($2)-2)}'",
		"df -h / | awk 'NR==2 {print substr($2, 1, length($2)-1)}'",
	}

	results := ssh.ExecSSHCommandInOrder(client, commands)

	for i, result := range results {
		if result.Error != nil {
			return fmt.Errorf("failed to execute command %d: %w", i+1, result.Error)
		}
	}

	updateFuncs := []struct {
		parse  func(string) (float64, error)
		update func(float64, *schema.Config)
		name   string
	}{
		{
			parse:  func(s string) (float64, error) { return strconv.ParseFloat(s, 64) },
			update: ethzconfig.CPUUpdateConfigHostname,
			name:   "CPU",
		},
		{
			parse:  func(s string) (float64, error) { return strconv.ParseFloat(s, 64) },
			update: ethzconfig.RAMUpdateConfigHostname,
			name:   "RAM",
		},
		{
			parse:  func(s string) (float64, error) { return strconv.ParseFloat(s, 64) },
			update: ethzconfig.StorageUpdateConfigHostname,
			name:   "Storage",
		},
	}

	for i, result := range results {
		value, err := updateFuncs[i].parse(strings.TrimSpace(result.Output))
		if err != nil {
			return fmt.Errorf("failed to parse %s value: %w", updateFuncs[i].name, err)
		}
		updateFuncs[i].update(value, cfg)
		logs.Info("%s value updated: %v", updateFuncs[i].name, value)
	}

	if err := utils.BlueprintWriteJsonFile(configPath, cfg); err != nil {
		return fmt.Errorf("failed to write updated configuration: %w", err)
	}

	logs.Info("System configuration successfully updated.")
	return nil
}
