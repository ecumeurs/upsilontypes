---
id: mechanic_buff_attribution_accessor
status: DRAFT
tags: entity,buffs,attribution,engine
parents:
  - [[upsilonbattle:mechanic_item_buff_application]]
layer: IMPLEMENTATION
priority: 5
dependents: []
type: MECHANIC
human_name: Buff Attribution Accessor Mechanic
version: 1.0
---

# Buff Attribution Accessor Mechanic

## INTENT
To provide a pure Go-internal read accessor on `Entity` that returns a property key's natural (unbuffed) value together with each individual contributing buff and its full owning `TemporaryProperties` block (origin entity, origin skill, duration, forever). This closes the attribution gap left by the existing buff-composition accessor `GetBuffsFor`, which returns only the composed `[]property.Property` values and discards which buff block produced each one — so today a caller can know that a value is buffed and by how much, but never by whom, from which skill, or for how long.

## THE RULE / LOGIC
**New accessor, additive to the existing composition path:**

```go
type BuffContribution struct {
    Buff  property.TemporaryProperties // full owning block: OriginEntityID, OriginSkillID, Duration, Forever
    Value property.Property            // this block's contribution to the requested key
}

func (e Entity) GetBuffAttributionFor(name property.Key) (base property.Property, contribs []BuffContribution)
```

`GetBuffAttributionFor` returns the entity's natural (unbuffed) value for the requested property key as `base`, plus one `BuffContribution` per buff block that contributed to that key. Each contribution preserves the full owning `TemporaryProperties` (`OriginEntityID`, `OriginSkillID`, `Duration`, `Forever`) instead of discarding it, which is what the existing `GetBuffsFor` (`entity.go:255`) does today.

**Additive, not a signature change.** `GetBuffsFor` is deliberately left untouched. Its five umbrella-wide call sites (one internal to `GetProperty`, three tests) need only the composed value, and buff *composition* itself does not want the owning block — only attribution does. The new accessor sits alongside `GetBuffsFor`, it does not replace or alter it.

**Read-only; not a read-modify-write primitive.** This accessor performs no persistence and no mutation of base state. It must never be used as the read half of a read-modify-write, because it exposes composed contributions rather than a value that is safe to write back onto the entity's base state — that base-state write-isolation invariant belongs to the write side of the same `Entity` property system and is unaffected by this purely additive read accessor.

**No JSON tags anywhere in this file.** `BuffContribution`, and the `property.TemporaryProperties`/`property.Property` types it embeds as used here, carry no json tags. Nothing can serialize this accessor's return value to a client without deliberate new serialization work being written elsewhere. Surfacing buff attribution to a client (e.g. showing which caster/skill applied a given buff) is explicitly out of scope for this atom and must be captured by its own BUSINESS atom written up front when that need arises — not retrofitted onto this one after the fact.

**`OriginEntityID`/`OriginSkillID` are reserved scaffolding for summons and traps, not dead fields.** These are an early breed of skill that has not seen much production use yet. `OriginSkillID` currently has exactly one occurrence umbrella-wide: its own declaration on `property.TemporaryProperties`. This accessor surfaces genuine, wired scaffolding rather than dead code: the `EntityDuration` field has a live decrementer, and `ExpiresWithCaster` has a live reader, but neither field has a production writer yet — the consumer machinery already exists and is waiting on its producer. A future reader of this accessor's output should not conclude these fields are broken merely because they are unpopulated for most buffs today.

**Known adjacent gap, not fixed by this atom.** Buff-scoped `Duration` can neither tick down (`Entity.BuffTickDown` is never called in production — the end-of-turn pass ticks skill cooldowns and the entity-scoped `EntityDuration`, but never buff-scoped duration) nor reach a client (the wire `Buff` DTO carries no `duration` field). This accessor surfaces whatever `Duration` value happens to be present on the owning block; it does not fix the underlying tick or serialization gap.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[mechanic_buff_attribution_accessor]]`, placed directly above the `GetBuffAttributionFor` method declaration — never at file or package header, per this project's surgical-placement convention.
- **File (not yet created):** `upsilontypes/entity/entity_buff_attribution.go`, package `entity`. This is a new sibling file to `upsilontypes/entity/entity.go` rather than an addition to it, because `entity.go` is at 385 LOC against the 400-LOC warn ceiling in `code_health_check.py` (which counts comments as effective LOC).
- **Type:** `BuffContribution struct { Buff property.TemporaryProperties; Value property.Property }` — `Buff` carries the full owning block (`OriginEntityID`, `OriginSkillID`, `Duration`, `Forever`); `Value` is that block's contribution to the requested key. No json tags anywhere in the file.
- **Method:** `func (e Entity) GetBuffAttributionFor(name property.Key) (base property.Property, contribs []BuffContribution)`.
- **Parent mechanic:** `[[upsilonbattle:mechanic_item_buff_application]]` — governs buff registration/composition (`RegisterBuff`, `GetProperty`) that this accessor reads from without altering.
- **Related, not modified:** `[[upsilonbattle:rule_entity_property_write_isolation]]` governs the write side of the same `Entity` property system; this atom's accessor is purely additive and read-only and does not implicate that rule's base-state isolation invariant.
- **Pre-existing gap, not created by this atom:** `RegisterBuff` (`entity.go:239`), `GetBuffsFor` (`entity.go:255`) and the `Buffs` field (`entity.go:56`) currently carry no `@spec-link` at all.

## EXPECTATION
- Calling `GetBuffAttributionFor("Armor")` on an entity with two buff blocks each contributing to Armor returns the entity's natural (unbuffed) Armor value as `base`, and a two-element `contribs` slice whose entries each carry the full originating `TemporaryProperties` (including `OriginEntityID`/`OriginSkillID`/`Duration`/`Forever`) alongside that block's `Value` contribution.
- An entity with no buffs contributing to the requested key returns `base` equal to the entity's natural value and an empty (not nil-panicking) `contribs` slice.
- The method is side-effect free: calling it twice in a row on the same entity returns identical results and leaves `Entity.Buffs` unchanged.
- Nothing produced by this method can be marshaled to JSON without new, deliberate serialization code being written elsewhere — no field on `BuffContribution`, or on the `TemporaryProperties`/`Property` values it exposes, carries a json tag.
