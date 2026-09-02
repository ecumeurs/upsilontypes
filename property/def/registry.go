package def

import (
	"fmt"

	"github.com/ecumeurs/upsilontypes/property"
)

// Scope is a bitmask of the vocabularies (Entity/Skill/Item) an Entry belongs to.
// architecture/property_key_vocabulary.md §7 lists four keys that become dual-scope
// (Entity|Skill) in this round.
type Scope uint8

const (
	ScopeEntity Scope = 1 << iota
	ScopeSkill
	ScopeItem
)

// Kind identifies which property.Property implementation backs a registry Entry.
type Kind string

const (
	KindInt        Kind = "Int"
	KindIntCounter Kind = "IntCounter"
	KindBool       Kind = "Bool"
	KindString     Kind = "String"
	KindZone       Kind = "Zone"
	KindEffect     Kind = "Effect"
)

// Composition names how ApplyBuff folds a buff onto a base value for a given Kind.
// Deliberately a named string type, not a function value or strategy object
// (project decision 5) — see architecture/property_key_vocabulary.md §6 correction 1
// for the four observed rules.
type Composition string

const (
	CompositionAdd     Composition = "Add"
	CompositionAnd     Composition = "And"
	CompositionReplace Composition = "Replace"
	CompositionNone    Composition = "None"
)

// Entry is one row of the property key registry: everything the vocabulary
// (architecture/property_key_vocabulary.md) declares about a single property.Key.
type Entry struct {
	Key          property.Key
	Scopes       Scope // bitmask; 4 keys are Entity|Skill
	Kind         Kind
	Composition  Composition
	MinInfoLevel property.InformationLevel
	Allowed      []string                 // non-nil only for validated String kinds
	New          func() property.Property // the default-when-absent constructor
}

// registry is the merged, flat property.Key -> Entry table. It is assembled by
// mergeThemes from a list of per-theme source maps; the seven theme data files
// that populate it are written by other agents in later slices of this round, so
// the source list is left empty here.
var registry = mergeThemes(
	entityVitalsEntries,
	entityMovementEntries,
	entityCoreEntries,
	skillTargetingEntries,
	skillEffectEntries,
	skillCostTriggerEntries,
	itemEntries,
)

// mergeThemes merges a list of theme maps into a single map[property.Key]Entry,
// panicking on any duplicate key (CODING_RULE §3, crash early). Splitting the table
// across theme files loses the compile-time duplicate-key check a single map
// literal would give; this panic is the deliberate runtime backstop for that.
// @spec-link [[upsilontypes:module_property_key_registry]]
func mergeThemes(themes ...map[property.Key]Entry) map[property.Key]Entry {
	merged := make(map[property.Key]Entry)
	for _, theme := range themes {
		for k, entry := range theme {
			if _, exists := merged[k]; exists {
				panic(fmt.Sprintf("def: duplicate property key %q while merging registry themes", k))
			}
			merged[k] = entry
		}
	}
	return merged
}

// Lookup returns the registry Entry for k, and whether it was found.
// @spec-link [[upsilontypes:module_property_key_registry]]
func Lookup(k property.Key) (Entry, bool) {
	entry, ok := registry[k]
	return entry, ok
}
