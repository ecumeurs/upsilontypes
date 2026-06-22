---
id: entity_skill_template
status: STABLE
layer: ARCHITECTURE
version: 2.0
dependents:
  - [[upsilonapi:api_skill_template_admin_crud]]
  - [[upsilonapi:api_skill_template_browse]]
  - [[entity_character_skill_inventory]]
type: ENTITY
priority: 5
tags: [skills, templates, catalog, iss-086]
human_name: Skill Template Entity
parents:
  - [[shared:req_tech_debt_backlog]]
---

# Skill Template Entity

## INTENT
To define the canonical skill template — the immutable blueprint from which character skills are instantiated. One template row per distinct skill type in the registry.

## THE RULE / LOGIC
**`skill_templates` table fields:**
- **id:** UUID primary key.
- **name:** Display name (e.g. "Fire Bolt").
- **behavior:** VARCHAR with CHECK constraint `IN ('Direct','Reaction','Passive','Counter','Trap')`. Maps to Go `def.BehaviorType*` constants.
- **targeting:** JSON map — engine targeting constraints (range, area, etc.).
- **costs:** JSON map — skill activation costs (MP, SP, etc.).
- **effect:** JSON map — outcome applied on resolution.
- **grade:** VARCHAR with CHECK constraint `IN ('I','II','III','IV','V')`. Tier indicator used by `[[upsilonbattle:rule_skill_grading_system]]`.
- **weight_positive / weight_negative:** Integers. Lottery weights for the roll system — positive biases toward selection, negative biases away.
- **available:** Boolean. `false` hides from player catalog and excludes from roll lottery.
- **version:** Varchar (default '1.0') for future migrations.

**Snapshot model:** At acquisition (`skill_roll`), the full template row is copied into `character_skills.instance_data` as frozen JSON. Post-acquisition template edits do not affect existing inventory entries.

**D11 exotic link:** `shop_items.skill_template_id` is a nullable FK to this table. When set, the bridge instantiates the template snapshot into the entity's `EquippedSkills` at arena init.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[entity_skill_template]]`
- **Laravel Model:** `App\Models\SkillTemplate`
- **Migration:** `*_create_skill_templates_table.php`
- **Resource:** `App\Http\Resources\SkillTemplateResource`

## EXPECTATION
- `behavior` and `grade` values are enforced by DB CHECK constraints.
- `weight_positive` and `weight_negative` are positive integers.
- `available=false` excludes from player listing and roll lottery.
- Snapshot in `character_skills.instance_data` is immutable after acquisition.
