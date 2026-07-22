// Package economyv1 holds the plain wire DTOs of the upsiloneconomy service —
// the contracts its clients (hub gateway, auth GDPR fan-out) marshal against.
// Transport isolation: no HTTP types here, structs only; payloads travel
// inside the standard platform envelope's data field. Money is int64 base
// units everywhere — no floats in money paths.
package economyv1

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Purchase failure reasons carried in envelope meta.reason.
const (
	ReasonInsufficientCredits = "insufficient_credits"
	ReasonQuantityCap         = "quantity_cap"
	ReasonItemUnavailable     = "item_unavailable"
)

// Wallet is one user's balance. Wallets are lazily created at first touch
// with the platform default balance; a read never fails for a live user.
//
// @spec-link [[upsiloneconomy:contract_economy_service]]
type Wallet struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance int64     `json:"balance"`
}

// WalletsBatchRequest resolves several balances in one call (admin listings,
// user-view composition).
type WalletsBatchRequest struct {
	IDs []uuid.UUID `json:"ids"`
}

// WalletsResponse carries wallet collections; unknown ids answer the default
// balance (lazy-create semantics), never an error.
type WalletsResponse struct {
	Wallets []Wallet `json:"wallets"`
}

// AwardRequest applies one credit grant. IdempotencyKey is enforced by a
// unique ledger index: replays answer Applied=false and change nothing, so
// durable-retry callers get exactly-once effect.
//
// @spec-link [[upsiloneconomy:contract_economy_service]]
type AwardRequest struct {
	IdempotencyKey string     `json:"idempotency_key"`
	PlayerID       uuid.UUID  `json:"player_id"`
	Amount         int64      `json:"amount"`
	Source         string     `json:"source"`
	MatchID        *uuid.UUID `json:"match_id,omitempty"`
}

// AwardResponse reports whether this call applied the grant (false = the
// idempotency key was already consumed).
type AwardResponse struct {
	Applied bool `json:"applied"`
}

// ShopItem is one catalog entry as it crosses the boundary. SkillTemplateID
// is a bare UUID reference into the hub's vocabulary — never a join.
//
// @spec-link [[upsilontypes:entity_shop_item]]
// @spec-link [[upsiloneconomy:contract_economy_service]]
type ShopItem struct {
	ID              uuid.UUID       `json:"id"`
	Name            string          `json:"name"`
	Type            *string         `json:"type"`
	Slot            string          `json:"slot"`
	Properties      json.RawMessage `json:"properties"`
	Cost            int64           `json:"cost"`
	Available       bool            `json:"available"`
	SkillTemplateID *uuid.UUID      `json:"skill_template_id"`
	Version         string          `json:"version"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// ShopItemCreate carries a validated admin create payload; nil Available
// takes the catalog default (true).
type ShopItemCreate struct {
	Name            string          `json:"name"`
	Type            *string         `json:"type"`
	Slot            string          `json:"slot"`
	Properties      json.RawMessage `json:"properties"`
	Cost            int64           `json:"cost"`
	Available       *bool           `json:"available"`
	SkillTemplateID *uuid.UUID      `json:"skill_template_id"`
}

// ShopItemUpdate carries a validated admin partial update; nil fields keep
// their current value. SetSkillTemplate distinguishes "leave alone" from
// "set to the (possibly nil) SkillTemplateID".
type ShopItemUpdate struct {
	Name             *string         `json:"name"`
	Type             *string         `json:"type"`
	Slot             *string         `json:"slot"`
	Properties       json.RawMessage `json:"properties,omitempty"`
	Cost             *int64          `json:"cost"`
	Available        *bool           `json:"available"`
	SetSkillTemplate bool            `json:"set_skill_template"`
	SkillTemplateID  *uuid.UUID      `json:"skill_template_id"`
}

// InventoryItem is one owned-item row with its catalog entry loaded — the
// shape the hub's equipment ownership check consumes.
type InventoryItem struct {
	ID          uuid.UUID `json:"id"`
	PlayerID    uuid.UUID `json:"player_id"`
	Quantity    int       `json:"quantity"`
	PurchasedAt time.Time `json:"purchased_at"`
	Item        ShopItem  `json:"item"`
}

// PurchaseRequest executes an atomic purchase for a player.
type PurchaseRequest struct {
	PlayerID   uuid.UUID `json:"player_id"`
	ShopItemID uuid.UUID `json:"shop_item_id"`
	Quantity   int       `json:"quantity"`
}

// PurchaseResponse answers a completed purchase: the balance after debit and
// the upserted inventory row. Domain rejections travel as envelope errors
// with meta.reason instead.
type PurchaseResponse struct {
	Credits int64         `json:"credits"`
	Item    InventoryItem `json:"item"`
}

// PurgeRequest is the idempotent GDPR wallet closure, called by the auth
// service's durable termination fan-out.
//
// @spec-link [[upsiloneconomy:contract_economy_service]]
type PurgeRequest struct {
	UserID         uuid.UUID `json:"user_id"`
	IdempotencyKey string    `json:"idempotency_key"`
}
