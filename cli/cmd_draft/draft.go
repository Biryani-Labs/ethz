package cmddraft

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Biryani-Labs/ethz/common/logs"
	"github.com/Biryani-Labs/ethz/common/utils"
	"github.com/Biryani-Labs/ethz/config"
	"github.com/Biryani-Labs/ethz/constants"
	"github.com/Biryani-Labs/ethz/pkg/schema"
)

type DraftCmd struct {
	BlueprintName string `arg:"" help:"Name of the blueprint you would like to create"`
	Delete        bool   `optional:"" help:"Delete the blueprint if it already exists"`
}

func (draft *DraftCmd) Run() error {
	blueprintPath := config.LocateInHomePath(draft.BlueprintName)
	if _, err := os.Stat(blueprintPath); os.IsNotExist(err) {
		if err := createBlueprintDirectoryStructure(blueprintPath); err != nil {
			return logs.Error(err, "error creating blueprint directory structure")
		}

		config := &schema.Config{}
		if err := utils.BlueprintWriteJsonFile(filepath.Join(blueprintPath, constants.BlueprintFile), config); err != nil {
			logs.Error(err, "error saving blueprint config")
		}

		logs.Info(fmt.Sprintf("New blueprint '%s' has been created", draft.BlueprintName))
	} else if draft.Delete {
		if err := os.RemoveAll(blueprintPath); err != nil {
			return logs.Error(err, "error deleting blueprint")
		}
		logs.Info(fmt.Sprintf("Blueprint '%s' has been deleted", draft.BlueprintName))
	} else {
		return logs.Error(nil, fmt.Sprintf("blueprint '%s' already exists", draft.BlueprintName))
	}

	return nil
}

func createBlueprintDirectoryStructure(blueprintPath string) error {
	if err := os.MkdirAll(filepath.Join(blueprintPath, "services"), 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(blueprintPath, "configs"), 0755); err != nil {
		return err
	}

	return nil
}
