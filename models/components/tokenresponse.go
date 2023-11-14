// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"time"
)

// Scope - Space-delimited list of authorization scope.
type Scope string

const (
	ScopeGlobal      Scope = "global"
	ScopePurgeSelect Scope = "purge_select"
	ScopePurgeAll    Scope = "purge_all"
	ScopeGlobalRead  Scope = "global:read"
)

func (e Scope) ToPointer() *Scope {
	return &e
}

func (e *Scope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "global":
		fallthrough
	case "purge_select":
		fallthrough
	case "purge_all":
		fallthrough
	case "global:read":
		*e = Scope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Scope: %v", v)
	}
}

type TokenResponse struct {
	// Time-stamp (UTC) of when the token was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Time-stamp (UTC) of when the token will expire (optional).
	ExpiresAt *string `json:"expires_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	// IP Address of the client that last used the token.
	IP *string `json:"ip,omitempty"`
	// Time-stamp (UTC) of when the token was last used.
	LastUsedAt *string `json:"last_used_at,omitempty"`
	// Name of the token.
	Name *string `json:"name,omitempty"`
	// Space-delimited list of authorization scope.
	Scope *Scope `default:"global" json:"scope"`
	// List of alphanumeric strings identifying services (optional). If no services are specified, the token will have access to all services on the account.
	//
	Services []string `json:"services,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// User-Agent header of the client that last used the token.
	UserAgent *string `json:"user_agent,omitempty"`
	UserID    *string `json:"user_id,omitempty"`
}

func (t TokenResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TokenResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TokenResponse) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TokenResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *TokenResponse) GetExpiresAt() *string {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *TokenResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TokenResponse) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *TokenResponse) GetLastUsedAt() *string {
	if o == nil {
		return nil
	}
	return o.LastUsedAt
}

func (o *TokenResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TokenResponse) GetScope() *Scope {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *TokenResponse) GetServices() []string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *TokenResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *TokenResponse) GetUserAgent() *string {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

func (o *TokenResponse) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
