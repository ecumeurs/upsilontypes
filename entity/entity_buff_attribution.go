package entity

import "github.com/ecumeurs/upsilontypes/property"

// BuffContribution pairs a single buff block's contribution to a requested
// property key with the full owning TemporaryProperties block it came from.
//
// No json tags: this type is engine-internal only. TemporaryProperties
// carries OriginEntityID/OriginSkillID attribution data, and the project's
// foe-loadout privacy guarantee is enforced by a per-recipient masking
// funnel in upsilonhub. Keeping this type unserializable means nothing can
// leak it to a client without deliberate new serialization work being
// written elsewhere.
type BuffContribution struct {
	// Buff is the full owning block: OriginEntityID, OriginSkillID,
	// Duration, Forever, and the rest of its Properties.
	Buff property.TemporaryProperties
	// Value is this block's contribution to the requested key.
	Value property.Property
}

// GetBuffAttributionFor returns the entity's natural (unbuffed) value for
// name as base, plus one BuffContribution per buff block in e.Buffs that
// carries name. Unlike GetBuffsFor, which composes and discards the owning
// block, each contribution here retains the full TemporaryProperties it
// came from, so a caller can tell who applied a buff, from which skill, and
// for how long, not just how much it is worth.
//
// Read-only: no persistence, no mutation of e.Buffs. Must never be used as
// the read half of a read-modify-write onto base state — use
// GetBaseProperty for that (see rule_entity_property_write_isolation).
//
// An unbuffed key is a normal state: base is still resolved and contribs is
// an empty (non-nil) slice; there is no error return and no panic.
// @spec-link [[mechanic_buff_attribution_accessor]]
func (e Entity) GetBuffAttributionFor(name property.Key) (base property.Property, contribs []BuffContribution) {
	base = e.getBasePropertyOrDefault(name)

	nname := property.PropertyToString(name)

	contribs = make([]BuffContribution, 0)
	for _, buff := range e.Buffs {
		value, found := buff.Properties[nname]
		if !found {
			continue
		}
		contribs = append(contribs, BuffContribution{
			Buff:  buff,
			Value: value,
		})
	}

	return base, contribs
}
