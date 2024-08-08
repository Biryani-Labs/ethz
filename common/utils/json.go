package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Biryani-Labs/ethz/pkg/schema"
)

func BlueprintReadJsonFile(filepath string) (*schema.Config, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var config schema.Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return &config, nil
}

func BlueprintWriteJsonFile(filepath string, config *schema.Config) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}
