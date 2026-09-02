package property

// Key is the single flat property key type shared by the Entity, Skill and Item
// vocabularies (see architecture/property_key_vocabulary.md).
type Key string

// String returns the key's string value.
func (k Key) String() string {
	return string(k)
}

const (
	// all these are counters
	HP       Key = "HP"       // Absence means 10
	Movement Key = "Movement" // Absence means 3; reset at end of turn
	SP       Key = "SP"       // Absence means 10
	MP       Key = "MP"       // Absence means 10

	Attack     Key = "Attack"     // Absence means 1, basic attack
	Defense    Key = "Defense"    // Absence means 0, basic defense
	JumpHeight Key = "JumpHeight" // Absence means 2
	TeamID     Key = "TeamID"     // Team affiliation

	// Flags
	// Absence means -1 (not dying), when HP reaches 0, IsDying is set to 3, and at each end of entity's turn, will be reduced by one. When it reaches 0, the entity is removed from game.
	IsDying  Key = "IsDying"
	HasMoved Key = "HasMoved" // Absence means false, can move multiple times up to acting, once acted, can't move. Trap triggered or reaction can also mark the target as having moved.
	HasActed Key = "HasActed" // Absence means false, will be set after the target has attacked or used a skill.

	// Buffs and other applied, inherited properties ... ?
	AttackRange Key = "AttackRange" // Absence means 1, Basic attack range. Altered with items, mostly.

	// @spec-link [[mechanic_temporary_entity_system]]
	// EntityDuration tracks how many turns a temporary entity lives. 0 means permanent.
	// NOTE: distinct from the skill-scope Duration used for buff duration.
	EntityDuration Key = "EntityDuration"
	// ExpiresWithCaster: if true, entity's owned positional effects are removed when their caster entity dies.
	// @spec-link [[mechanic_effect_caster_tracking]]
	ExpiresWithCaster Key = "ExpiresWithCaster"
	// WalkThrough: if true, other entities may occupy the same cell as this entity.
	// @spec-link [[mechanic_multi_entity_cell_system]]
	WalkThrough Key = "WalkThrough"
	// Invisible: if true, the entity is not sent in client-facing state snapshots.
	Invisible Key = "Invisible"
	// AIArchetype stores the archetype slug for controller-driven AI entities
	// (e.g. "fighter", "ranger", "support", "sneak").
	// @spec-link [[upsilonbattle:mechanic_ai_controller_archetypes]]
	AIArchetype Key = "AIArchetype"

	// (counters) Absence means 0,0, can have overshield (when applied through healing and buffs...) Max shield is only used at initialisation of battle. Allowed to twice HP Max
	Shield Key = "Shield"
	Poison Key = "Poison" // Absence means 0, Poisoned state. Sum of all poison damage taken per turn. Mostly a temporary status. Negative PoisonPower will cure; /2 each turn, remove at <1
	Stun   Key = "Stun"   // Absence means 0, Stunned state. Sum of all stun taken per turn. Mostly a temporary status. Negative StunPower will cure; /2 each turn, remove at <1
)

const (
	Behavior Key = "Behavior" // property.Skill broad category: Direct, Reaction, Passive, Counter, Trap; Absence means Direct

	Range              Key = "Range"              // Range of the property.Skill // Absence means 1
	Zone               Key = "Zone"               // Area of Effect // Absence means 1 tile effect
	TargetNumber       Key = "TargetNumber"       // Absence means all targets within the targeted zone.
	Accuracy           Key = "Accuracy"           // Absence means 100%
	Dodge              Key = "Dodge"              // Absence means 0%
	Parry              Key = "Parry"              // Absence means 0%
	TargetType         Key = "TargetType"         // Entity, Tile, Both, Self
	TargetingMechanics Key = "TargetingMechanics" // Anywhere, Line of Sight, and maybe other mechanics later.

	DamageScale        Key = "DamageScale"        // Absence means 100%. RENAMED from Damage (vocabulary §7): percentage scaling, not a flat addend.
	Heal               Key = "Heal"               // Absence means 0
	ShieldPower        Key = "ShieldPower"        // Absence means 0 , can be negative or positive. VALUE RENAMED from "Shield" (vocabulary §7).
	StunPower          Key = "StunPower"          // Absence means 0 , can be negative or positive. VALUE RENAMED from "Stun" (vocabulary §7).
	StunChance         Key = "StunChance"         // Absence means 0%
	CriticalChance     Key = "CriticalChance"     // Absence means 0%
	CriticalMultiplier Key = "CriticalMultiplier" // Absence means 100%
	Duration           Key = "Duration"           // Absence means 0
	PoisonPower        Key = "PoisonPower"        // Absence means 0 , can be negative or positive. VALUE RENAMED from "Poison" (vocabulary §7).
	PoisonChance       Key = "PoisonChance"       // Absence means 0%

	Delay        Key = "Delay"        // Absence means 500
	Channeling   Key = "Channeling"   // Absence means 0
	HPLeech      Key = "HPLeech"      // Absence means 0
	MPLeech      Key = "MPLeech"      // Absence means 0
	SPLeech      Key = "SPLeech"      // Absence means 0
	MovementCost Key = "MovementCost" // Absence means 0. RENAMED from MvtCost (vocabulary §7).

	Cooldown Key = "Cooldown" // Absence means 3 turns. Special note: Cool down is stored as a counter, minValue represent initial cooldown at battle start. MaxValue represent the cooldown value when used.

	// @spec-link [[mech_movement_reposition]]
	// RepositionSubject indicates who the movement skill displaces: "Self" (caster — dash/teleport)
	// or "Target" (the targeted entity — push/pull/kick). Absence means no reposition.
	RepositionSubject Key = "RepositionSubject"
	// RepositionDistance is the number of tiles the subject is displaced along the casting ray
	// (caster→target). Positive moves along the ray (dash forward / push away from caster);
	// negative moves against it (pull toward caster). Absence (or 0) means no reposition.
	RepositionDistance Key = "RepositionDistance"

	// @spec-link [[mech_trigger_system]]
	// TriggerType defines when a positional effect fires. Value is a TriggerTypeValue string.
	TriggerType Key = "TriggerType"
	// RemoveOnTrigger: if true, the positional effect is consumed after firing once.
	RemoveOnTrigger Key = "RemoveOnTrigger"
	// TriggerCount: how many times the effect can fire (0 = unlimited).
	TriggerCount Key = "TriggerCount"
)

var SkillTargetingProperties = map[Key]bool{
	Range:              true,
	Zone:               true,
	TargetNumber:       true,
	Accuracy:           true,
	Dodge:              true,
	Parry:              true,
	TargetType:         true,
	TargetingMechanics: true,
}

var SkillEffectProperties = map[Key]bool{
	DamageScale:        true,
	Heal:               true,
	ShieldPower:        true,
	StunPower:          true,
	StunChance:         true,
	CriticalChance:     true,
	CriticalMultiplier: true,
	Duration:           true,
	PoisonPower:        true,
	PoisonChance:       true,
}

var SkillCostProperties = map[Key]bool{
	Delay:      true,
	Channeling: true,
	HPLeech:    true,
	MPLeech:    true,
	SPLeech:    true,
	Cooldown:   true,
}

const (
	Durability       Key = "Durability"       // Absence means 0: invulnerable
	Weight           Key = "Weight"           // Absence means 0: no weight
	ItemType         Key = "ItemType"         // Absence means None (out of Wearable, Consumable, Usable, Throwable, Ammunitions and None)
	ArmorRating      Key = "ArmorRating"      // Absence means 0: no armor (only for Wearable). VALUE RENAMED from "Armor" (vocabulary §7).
	WeaponType       Key = "WeaponType"       // Absence means 0: no weapon type (only for Wearable)
	ArmorType        Key = "ArmorType"        // Absence means 0: no armor type (only for Wearable)
	ToolType         Key = "ToolType"         // Absence means 0: no tool type (only for Wearable)
	WeaponRange      Key = "WeaponRange"      // Absence means 0: no weapon range (only for Wearable)
	WeaponBaseDamage Key = "WeaponBaseDamage" // Absence means 0: no weapon base damage (only for Wearable)
	Stackable        Key = "Stackable"        // Absence means 0: not stackable
	StackSize        Key = "StackSize"        // Absence means 0: no stack size
	Effect           Key = "Effect"           // Absence means nil: No effect. Effects are Skills. (except None)
	ItemValue        Key = "ItemValue"        // Absence means 0: no value. RENAMED from Value (vocabulary §7).
)

// PropertyToString converts a Key into its string representation.
func PropertyToString(p Key) string {
	return p.String()
}
