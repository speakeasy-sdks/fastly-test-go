// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

type SchemasUserResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	CustomerID *string    `json:"customer_id,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The alphanumeric string identifying a email login.
	EmailHash *string `json:"email_hash,omitempty"`
	ID        *string `json:"id,omitempty"`
	// Indicates that the user has limited access to the customer's services.
	LimitServices *bool `json:"limit_services,omitempty"`
	// Indicates whether the is account is locked for editing or not.
	Locked *bool   `json:"locked,omitempty"`
	Login  *string `json:"login,omitempty"`
	// The real life name of the user.
	Name *string `json:"name,omitempty"`
	// Indicates if a new password is required at next login.
	RequireNewPassword *bool `json:"require_new_password,omitempty"`
	// The permissions role assigned to the user. Can be `user`, `billing`, `engineer`, or `superuser`.
	Role *RoleUser `json:"role,omitempty"`
	// Indicates if 2FA is enabled on the user.
	TwoFactorAuthEnabled *bool `json:"two_factor_auth_enabled,omitempty"`
	// Indicates if 2FA is required by the user's customer account.
	TwoFactorSetupRequired *bool `json:"two_factor_setup_required,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (s SchemasUserResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SchemasUserResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SchemasUserResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *SchemasUserResponse) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *SchemasUserResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *SchemasUserResponse) GetEmailHash() *string {
	if o == nil {
		return nil
	}
	return o.EmailHash
}

func (o *SchemasUserResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SchemasUserResponse) GetLimitServices() *bool {
	if o == nil {
		return nil
	}
	return o.LimitServices
}

func (o *SchemasUserResponse) GetLocked() *bool {
	if o == nil {
		return nil
	}
	return o.Locked
}

func (o *SchemasUserResponse) GetLogin() *string {
	if o == nil {
		return nil
	}
	return o.Login
}

func (o *SchemasUserResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SchemasUserResponse) GetRequireNewPassword() *bool {
	if o == nil {
		return nil
	}
	return o.RequireNewPassword
}

func (o *SchemasUserResponse) GetRole() *RoleUser {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *SchemasUserResponse) GetTwoFactorAuthEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.TwoFactorAuthEnabled
}

func (o *SchemasUserResponse) GetTwoFactorSetupRequired() *bool {
	if o == nil {
		return nil
	}
	return o.TwoFactorSetupRequired
}

func (o *SchemasUserResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
