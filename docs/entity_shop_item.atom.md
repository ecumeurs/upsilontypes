---
id: entity_shop_item
status: STABLE
type: ENTITY
layer: ARCHITECTURE
priority: 5
version: 2.0
tags: [shop, items, catalog, iss-074]
human_name: Shop Item Catalog Entity
parents:
  - [[upsilonapi:domain_credit_economy]]
  - [[upsilonbattle:mec_credit_spending_shop]]
dependents: []
---

# New Atom

## INTENT
To define the shop catalog entity — the immutable list of items players can purchase with credits. One row per purchasable item; the V2.0 seed contains exactly three rows (Basic Armor, Basic Sword, Swift Boots).

## THE RULE / LOGIC
**Shop Item Entity:**

**Core Fields (`shop_items` table):**
- **id:** UUID primary key (deterministic in V2.0 seed for test stability).
- **name:** Varchar display name.
- **type:** Varchar item category tag (e.g. "armor", "weapon"); informational.
- **slot:** Varchar with CHECK constraint `slot IN ('armor','utility','weapon')`. The 3-slot binding target.
- **properties:** JSON map keyed by engine `ItemProperties` enum (e.g. `{"ArmorRating":5}`). Materialised at battle init via `def.ItemProperty()` factory.
- **cost:** Integer credit cost per unit.
- **available:** Boolean (default `true`). When `false`, hidden from player `GET /v1/shop/items` but existing inventory rows remain valid.
- **skill_template_id:** Nullable UUID FK → `skill_templates.id`. When set, this is a D11 exotic item — the bridge appends the linked template's snapshot to the entity's `EquippedSkills` at arena init with `origin='item:{inv_item_id}'`.
- **version:** Varchar catalog generation tag.

**V2.0 Seed (deterministic):**
- Basic Armor — slot=armor, properties={ArmorRating:5}, cost=200.
- Basic Sword — slot=weapon, properties={WeaponBaseDamage:5, WeaponType:"One-Handed Melee", WeaponRange:1}, cost=300.
- Swift Boots — slot=utility, properties={Movement:1}, cost=150.

**Admin mutations:** Full CRUD via `[[api_shop_item_admin_crud]]` (ISS-086). Includes availability toggling and exotic item creation.

## TECHNICAL INTERFACE
- **Code Tag:** `@spec-link [[entity_shop_item]]`
- **Laravel Model:** `App\Models\ShopItem`
- **Migration:** `*_create_item_system_tables.php`
- **Seeder:** `Database\Seeders\ShopItemsSeeder`
- **API:** `[[api_shop_browse]]`

## EXPECTATION
- `php artisan migrate:fresh --seed` produces exactly three rows.
- Deterministic UUIDs allow test scenarios to reference items by ID.
- `slot` value is enforced by CHECK; invalid inserts fail.
- `cost` is a positive integer; `properties` deserialises to a valid item property map.
