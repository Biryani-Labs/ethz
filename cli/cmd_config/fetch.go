package cmdconfig

import "github.com/Biryani-Labs/ethz/pkg/schema"

type FetchConfig struct {
	schema.CliBlueprintName
	MarketplaceBranch string `arg:"" help:"Name of the branch from where to fetch the config"`
	MarketplaceName   string `arg:"" help:"Name of the blueprint to fetch from marketplace"`
}
