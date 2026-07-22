// Package authv1 holds the plain wire DTOs of the upsilonauth service —
// the contracts other services (hub gateway, future games) marshal against.
// Transport isolation: no HTTP types here, structs only; payloads travel
// inside the standard platform envelope's data field.
package authv1

import (
	"time"

	"github.com/google/uuid"
)

// User is the account as it crosses the service boundary: auth-owned fields
// only. Wallet balance, play stats and reroll count are other services'
// truth and never appear here. PasswordHash never crosses the wire.
//
// @spec-link [[upsilonauth:contract_auth_service]]
type User struct {
	ID          uuid.UUID  `json:"id"`
	AccountName string     `json:"account_name"`
	Email       string     `json:"email"`
	Role        string     `json:"role"`
	FullAddress *string    `json:"full_address"`
	BirthDate   *time.Time `json:"birth_date"`
	// DeletedAt reports soft-deletion for with_trashed admin flows; live
	// lookups always answer DeletedAt == nil.
	DeletedAt *time.Time `json:"deleted_at"`
	// UpdatedAt feeds the admin registry's cursor pagination.
	UpdatedAt time.Time `json:"updated_at"`
}

// Token is the metadata of one opaque personal access token; the plaintext
// exists only in issuance and introspection-renewal responses.
type Token struct {
	ID        int64      `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
}

// IntrospectRequest asks the trust authority to resolve one bearer plaintext.
type IntrospectRequest struct {
	Token string `json:"token"`
}

// IntrospectResponse answers an introspection. Active=false carries nothing
// else (unknown, expired, revoked and deleted-owner are indistinguishable by
// design). RenewedToken is set when the authority performed sliding renewal
// during this introspection; the caller must relay it to the end client.
//
// @spec-link [[upsilonauth:contract_auth_service]]
type IntrospectResponse struct {
	Active       bool    `json:"active"`
	User         *User   `json:"user,omitempty"`
	Token        *Token  `json:"token,omitempty"`
	RenewedToken *string `json:"renewed_token,omitempty"`
}

// UsersBatchRequest resolves several user ids in one call (matchmaking
// assembly, admin composition).
type UsersBatchRequest struct {
	IDs []uuid.UUID `json:"ids"`
}

// UsersResponse carries user collections (batch resolution, admin registry
// pages). Unknown ids are omitted, never errored.
type UsersResponse struct {
	Users []User `json:"users"`
}

// CountAdminsResponse answers the last-admin guard count.
type CountAdminsResponse struct {
	Count int64 `json:"count"`
}

// AccountPush is the durable account-lifecycle notification auth pushes to
// the hub's player read model on every account mutation (register, rename,
// anonymize, soft-delete). Consumers upsert idempotently: a push older than
// the stored UpdatedAt is a no-op.
//
// @spec-link [[upsilonauth:contract_auth_service]]
type AccountPush struct {
	UserID      uuid.UUID  `json:"user_id"`
	AccountName string     `json:"account_name"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
