// Package seedids derives the deterministic, well-known UUIDs that glue the
// per-service seeds together: auth seeds accounts, economy seeds the shop
// catalog and the hub seeds characters/stats — independently, in any order —
// yet all three agree on every id because each id is a UUIDv5 of a stable
// name under a fixed namespace. Never persist a seed row under a random id.
package seedids

import "github.com/google/uuid"

// Namespace roots, all derived deterministically from the platform domain so
// no magic literals need to be kept in sync.
//
// @spec-link [[upsilonauth:contract_auth_service]]
// @spec-link [[upsiloneconomy:contract_economy_service]]
var (
	// Root is the platform seed namespace: UUIDv5(DNS, "seed.upsilon-hub.com").
	Root = uuid.NewSHA1(uuid.NameSpaceDNS, []byte("seed.upsilon-hub.com"))
	// accounts and shopItems are per-kind sub-namespaces under Root.
	accounts  = uuid.NewSHA1(Root, []byte("accounts"))
	shopItems = uuid.NewSHA1(Root, []byte("shop_items"))
)

// Account answers the well-known user UUID for a seeded account name.
// auth seeds the account under this id; the hub seeds characters and stats
// against the same id without ever asking auth.
func Account(accountName string) uuid.UUID {
	return uuid.NewSHA1(accounts, []byte(accountName))
}

// ShopItem answers the well-known catalog UUID for a seeded shop item name,
// so scenario scripts and cross-service fixtures can reference catalog
// entries without querying the economy service first.
func ShopItem(name string) uuid.UUID {
	return uuid.NewSHA1(shopItems, []byte(name))
}
