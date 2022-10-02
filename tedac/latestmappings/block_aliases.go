package latestmappings

import (
	_ "embed"

	"github.com/sandertv/gophertunnel/minecraft/nbt"
)

var (
	//go:embed block_aliases.nbt
	blockAliasesData []byte
	// aliasMappings maps from a legacy block name alias to an updated name.
	blockAliasMappings = map[string]string{}
	// reverseAliasMappings maps from an updated block name to a legacy block name alias.
	reverseBlockAliasMappings = map[string]string{}
)

// UpdatedNameFromAlias returns the updated name of a block from a legacy alias. If no alias was found, the
// second return value will be false.
func UpdatedBlockNameFromAlias(name string) (string, bool) {
	if updated, ok := blockAliasMappings[name]; ok {
		return updated, true
	}
	return name, false
}

// AliasFromUpdatedName returns the legacy alias of a block from an updated name. If no alias was found, the
// second return value will be false.
func AliasFromUpdatedBlockName(name string) (string, bool) {
	if alias, ok := reverseBlockAliasMappings[name]; ok {
		return alias, true
	}
	return name, false
}

// init creates conversions for each legacy and alias entry.
func init() {
	if err := nbt.Unmarshal(blockAliasesData, &blockAliasMappings); err != nil {
		panic(err)
	}
	for alias, name := range blockAliasMappings {
		reverseBlockAliasMappings[name] = alias
	}
}
