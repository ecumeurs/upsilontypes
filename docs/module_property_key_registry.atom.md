---
id: module_property_key_registry
status: DRAFT
priority: 5
parents:
  - [[shared:rule_stat_taxonomy]]
type: MODULE
version: 1.0
tags: property,registry,iss-140,iss-143,iss-145,iss-147
dependents: []
human_name: Unified Property Key Registry
layer: ARCHITECTURE
---

# Unified Property Key Registry

## INTENT
To replace the three independent property key namespaces in upsilontypes/property (EntityProperties, SkillProperties, ItemProperties — three named string types sharing no common ancestor but interface{}) with one flat property.Key key space backed by a declarative, table-driven registry in package upsilontypes/property/def. Scope (Entity/Skill/Item) demotes from a Go type to a metadata field on each registry entry, so a property can be single- or multi-scoped without a duplicate declaration. This single change structurally dissolves four filed defects at once: ISS-140 (the bridge silently drops unrecognized skill properties because the type-switch resolver has no reject path), ISS-143 (the propertyAliasMap loses its reason to exist once there is one key space to alias against), ISS-145 Defect 1 (four combat modifiers — Accuracy, Dodge, CriticalChance, CriticalMultiplier — are unreachable as entity properties because they only exist in the Skill namespace), and ISS-147 (the string values "Poison", "Stun", "Shield" collide across the Entity and Skill namespaces with different meanings, because storage is a flat map[string]Property keyed by the stringified name).

## THE RULE / LOGIC
**One flat key space, one registry.** Constants remain declared in package `property` as a single `property.Key` string type; the registry that maps each key to its behavior (`property/def`) is a set of ~7 theme-grouped `map[property.Key]Entry` tables (vitals, movement, combat-core, status effects, skill targeting, skill effect, skill cost, item), not 64 individual files and not codegen. Declaring the same key twice is a Go compile error (duplicate map key), which is the strongest available guard against the ISS-147 collision class recurring.

**Invariant: constant identifier == string value, for every registry entry.** Today exactly 4 of 64 keys violate this (`ShieldPower`/`StunPower`/`PoisonPower` carry the values `"Shield"`/`"Stun"`/`"Poison"`; `ArmorRating` carries `"Armor"`) and three of those four violations are the literal ISS-147 collision. The registry restores the invariant for all 64 entries via a 7-entry rename map (full per-key detail in the frozen specification, `architecture/property_key_vocabulary.md` §1 and §7). This is registry-walkable and therefore testable: a conformance test asserts `entry.Key == property.Key(<identifier>)` for every entry, turning a future divergence into a failing test rather than a silent production defect.

**Composition is a declared per-entry field with exactly four values**, not the two originally assumed. `Add` — `base + buff`, applies to Int/IntCounter/Float (IntCounter composes both `Value` and `MaxValue`). `And` — `base && buff`, applies to Bool. `Replace` — the buff's value overwrites the base, applies to the Zone kind only. `None` — `ApplyBuff` returns the base unchanged and the buff is ignored, applies to String and Effect kinds. `ApplyBuff` remains the executor; a table-driven conformance test asserts each entry's observed `ApplyBuff` behavior matches its declared Composition rule.

**Default-when-absent is a distinct concept from the starting loadout.** The registry declares only the value the resolver returns when a key is not present on a carrier (default-when-absent). `def.PropertiesForCharacter()`, the list a real character is created with, is a separate, deliberately-divergent concern for some keys (e.g. `Attack` defaults to `1` when absent but a new character starts with `3`) and is not folded into the registry by this decision.

**Scope moves from a Go type to entry metadata**, giving the resolver three distinguishable outcomes where today there are two: unknown key (not in the registry — reject), known-but-wrong-scope (in the registry, but not declared for the scope being resolved — reject), and accepted. A key may declare more than one scope; four keys (`Accuracy`, `Dodge`, `CriticalChance`, `CriticalMultiplier`) become dual Skill+Entity scoped by this decision, resolving ISS-145 Defect 1.

**Dual Item+Entity scope — eleven keys, encoding the ISS-142 buffability ruling (user, 2026-08-27, reconfirmed 2026-09-01).** Separately from the Skill+Entity widening above, eleven keys are declared with both `Item` and `Entity` scope: `HP`, `Movement`, `SP`, `MP`, `Attack`, `Defense`, `JumpHeight` and `AttackRange` are baseline entity attributes ruled legitimately item-grantable — an equipped item may buff them via `applyItemAsBuff` (`upsilonapi/bridge/bridge_start.go`); `ArmorRating`, `WeaponRange` and `WeaponBaseDamage` are authored on items but legitimately read directly off an Entity, because `applyItemAsBuff` flattens every equipped item into a `Forever:true` entity buff — equipment is never a live, independently-queried layer. `Shield`, `Poison` and `Stun` are deliberately excluded from this widening (already applied directly by the effect applicator; a parallel item-buff path would reproduce the ISS-147 collision class), as are the nine flag/plumbing keys `TeamID`, `IsDying`, `HasMoved`, `HasActed`, `EntityDuration`, `ExpiresWithCaster`, `WalkThrough`, `Invisible` and `AIArchetype` — buffing those from an item would be an exploit, not a feature. A registry conformance test (`registry_buffability_test.go`) pins this exact partition.

**Supersession.** This atom replaces `[[upsilontypes:mech_entity_properties]]` (MODULE/ARCHITECTURE) and its two IMPLEMENTATION children `[[upsilontypes:mech_entity_properties_item_properties]]` and `[[upsilontypes:mech_entity_properties_skill_properties]]`, whose subject matter — three independent Entity/Skill/Item property namespaces — is exactly the split this registry dissolves. Those three atoms are demoted to DRAFT and marked superseded in their own LOGIC sections rather than deleted (2026-08-30).

**Governance — backward-compatibility waiver, user-approved 2026-08-30.** `[[upsilontypes:contract_types_contract]]` (STABLE) states: "Stability: changes to core types must be backwards compatible or strictly versioned to avoid breaking the multi-stack integration." This decision deletes three exported type names (`EntityProperties`, `SkillProperties`, `ItemProperties`) and renames seven property string values — explicitly not backwards compatible. The user explicitly waived this clause for this change on 2026-08-30, because: the application is not live; all three consumers of the property package (upsilonbattle, upsilonapi, upsilontypes itself) are in-workspace and move together under the `go.work` multi-module workspace as one coordinated, atomic change; there is no external consumer of these types; and the database is disposable (reseeding is cheap, no migration constraint). The waiver applies to this change only, not as a standing exception to the contract clause.

**Governance — the Logic Isolation clause is not violated by the per-entry Composition field.** The same CONTRACT also states: "Logic Isolation: this project must contain only data structures and basic validation/serialization logic. No game mechanics." Declaring a Composition field per registry entry (Add/And/Replace/None, above) could read as mechanics moving into upsilontypes. It is not new mechanics: composition already lives in this exact package today — `defaultproperty.go`'s `DefaultIntProperty.ApplyBuff` already performs additive composition (base plus buff), with five comparable `ApplyBuff` implementations already present in `defaultproperty.go` before this decision. Declaring `Composition: CompositionAdd` (etc.) on a registry entry documents behavior the package already executes; it does not introduce a new mechanic.

## TECHNICAL INTERFACE
- **Code Tag (once implemented):** `@spec-link [[module_property_key_registry]]` — not yet injected; no source code exists for this decision at atom creation time (2026-08-30, pre-code capture per ATD Workflow E).
- **Expected package layout:** constants as `property.Key` in package `upsilontypes/property` (replacing `propertyenum.go`'s three separate string types); the registry itself as ~7 theme-grouped `map[property.Key]Entry` tables in package `upsilontypes/property/def`.
- **Full per-key specification:** `architecture/property_key_vocabulary.md` — the frozen table of all 64 keys (scope, kind, default-when-absent, composition, min info level), the complete 7-entry rename map, and 7 recorded implementation traps. That document is the data appendix this atom does not restate; this atom states the rule the registry itself must obey, not its row-by-row content.

## EXPECTATION
- Every registry entry satisfies `entry.Key == property.Key(<its own identifier>)` — walkable and asserted by a conformance test; closes the ISS-147 collision class structurally.
- For every registry entry, a freshly constructed default value's `Name(GameMaster)` equals the registry key it is stored under; closes the `Name()`-desync trap where `ZoneProperty`/`EffectProperty` currently return hardcoded literals.
- `ApplyBuff` behavior for every entry matches its declared Composition rule (Add/And/Replace/None), asserted by a table-driven conformance test.
- Resolving a property key against a scope produces exactly one of three outcomes — unknown key, known-but-wrong-scope, or accepted — never a silent nil and never a cross-scope fallthrough (the ISS-147 defect path).
- `def.PropertiesForCharacter()` (starting loadout) and the registry's default-when-absent values remain independently declared and are never collapsed into a single source.
- Duplicate key declarations across the ~7 theme-grouped registry files fail to compile (Go duplicate-map-key error), not merely to lint.
