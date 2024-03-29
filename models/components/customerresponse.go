// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// CustomerResponseBillingNetworkType - Customer's current network revenue type.
type CustomerResponseBillingNetworkType string

const (
	CustomerResponseBillingNetworkTypePublic  CustomerResponseBillingNetworkType = "public"
	CustomerResponseBillingNetworkTypePrivate CustomerResponseBillingNetworkType = "private"
)

func (e CustomerResponseBillingNetworkType) ToPointer() *CustomerResponseBillingNetworkType {
	return &e
}

func (e *CustomerResponseBillingNetworkType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		*e = CustomerResponseBillingNetworkType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CustomerResponseBillingNetworkType: %v", v)
	}
}

type CustomerResponse struct {
	// The alphanumeric string representing the primary billing contact.
	BillingContactID *string `json:"billing_contact_id,omitempty"`
	// Customer's current network revenue type.
	BillingNetworkType *CustomerResponseBillingNetworkType `json:"billing_network_type,omitempty"`
	// Used for adding purchased orders to customer's account.
	BillingRef *string `json:"billing_ref,omitempty"`
	// Whether this customer can view or edit wordpress.
	CanConfigureWordpress *bool `json:"can_configure_wordpress,omitempty"`
	// Whether this customer can reset passwords.
	CanResetPasswords *bool `json:"can_reset_passwords,omitempty"`
	// Whether this customer can upload VCL.
	CanUploadVcl *bool `json:"can_upload_vcl,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Specifies whether 2FA is forced or not forced on the customer account. Logs out non-2FA users once 2FA is force enabled.
	Force2fa *bool `json:"force_2fa,omitempty"`
	// Specifies whether SSO is forced or not forced on the customer account.
	ForceSso *bool `json:"force_sso,omitempty"`
	// Specifies whether the account has access or does not have access to the account panel.
	HasAccountPanel *bool `json:"has_account_panel,omitempty"`
	// Specifies whether the account has access or does not have access to the improved events.
	HasImprovedEvents *bool `json:"has_improved_events,omitempty"`
	// Whether this customer can view or edit the SSL config.
	HasImprovedSslConfig *bool `json:"has_improved_ssl_config,omitempty"`
	// Specifies whether the account has enabled or not enabled openstack logging.
	HasOpenstackLogging *bool `json:"has_openstack_logging,omitempty"`
	// Specifies whether the account can edit PCI for a service.
	HasPci *bool `json:"has_pci,omitempty"`
	// Specifies whether PCI passwords are required for the account.
	HasPciPasswords *bool   `json:"has_pci_passwords,omitempty"`
	ID              *string `json:"id,omitempty"`
	// The range of IP addresses authorized to access the customer account.
	IPWhitelist *string `json:"ip_whitelist,omitempty"`
	// The alphanumeric string identifying the account's legal contact.
	LegalContactID *string `json:"legal_contact_id,omitempty"`
	// The name of the customer, generally the company name.
	Name *string `json:"name,omitempty"`
	// The alphanumeric string identifying the account owner.
	OwnerID *string `json:"owner_id,omitempty"`
	// The phone number associated with the account.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The postal address associated with the account.
	PostalAddress *string `json:"postal_address,omitempty"`
	// The pricing plan this customer is under.
	PricingPlan *string `json:"pricing_plan,omitempty"`
	// The alphanumeric string identifying the pricing plan.
	PricingPlanID *string `json:"pricing_plan_id,omitempty"`
	// The alphanumeric string identifying the account's security contact.
	SecurityContactID *string `json:"security_contact_id,omitempty"`
	// The alphanumeric string identifying the account's technical contact.
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (c CustomerResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CustomerResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CustomerResponse) GetBillingContactID() *string {
	if o == nil {
		return nil
	}
	return o.BillingContactID
}

func (o *CustomerResponse) GetBillingNetworkType() *CustomerResponseBillingNetworkType {
	if o == nil {
		return nil
	}
	return o.BillingNetworkType
}

func (o *CustomerResponse) GetBillingRef() *string {
	if o == nil {
		return nil
	}
	return o.BillingRef
}

func (o *CustomerResponse) GetCanConfigureWordpress() *bool {
	if o == nil {
		return nil
	}
	return o.CanConfigureWordpress
}

func (o *CustomerResponse) GetCanResetPasswords() *bool {
	if o == nil {
		return nil
	}
	return o.CanResetPasswords
}

func (o *CustomerResponse) GetCanUploadVcl() *bool {
	if o == nil {
		return nil
	}
	return o.CanUploadVcl
}

func (o *CustomerResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CustomerResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *CustomerResponse) GetForce2fa() *bool {
	if o == nil {
		return nil
	}
	return o.Force2fa
}

func (o *CustomerResponse) GetForceSso() *bool {
	if o == nil {
		return nil
	}
	return o.ForceSso
}

func (o *CustomerResponse) GetHasAccountPanel() *bool {
	if o == nil {
		return nil
	}
	return o.HasAccountPanel
}

func (o *CustomerResponse) GetHasImprovedEvents() *bool {
	if o == nil {
		return nil
	}
	return o.HasImprovedEvents
}

func (o *CustomerResponse) GetHasImprovedSslConfig() *bool {
	if o == nil {
		return nil
	}
	return o.HasImprovedSslConfig
}

func (o *CustomerResponse) GetHasOpenstackLogging() *bool {
	if o == nil {
		return nil
	}
	return o.HasOpenstackLogging
}

func (o *CustomerResponse) GetHasPci() *bool {
	if o == nil {
		return nil
	}
	return o.HasPci
}

func (o *CustomerResponse) GetHasPciPasswords() *bool {
	if o == nil {
		return nil
	}
	return o.HasPciPasswords
}

func (o *CustomerResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CustomerResponse) GetIPWhitelist() *string {
	if o == nil {
		return nil
	}
	return o.IPWhitelist
}

func (o *CustomerResponse) GetLegalContactID() *string {
	if o == nil {
		return nil
	}
	return o.LegalContactID
}

func (o *CustomerResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CustomerResponse) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *CustomerResponse) GetPhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}

func (o *CustomerResponse) GetPostalAddress() *string {
	if o == nil {
		return nil
	}
	return o.PostalAddress
}

func (o *CustomerResponse) GetPricingPlan() *string {
	if o == nil {
		return nil
	}
	return o.PricingPlan
}

func (o *CustomerResponse) GetPricingPlanID() *string {
	if o == nil {
		return nil
	}
	return o.PricingPlanID
}

func (o *CustomerResponse) GetSecurityContactID() *string {
	if o == nil {
		return nil
	}
	return o.SecurityContactID
}

func (o *CustomerResponse) GetTechnicalContactID() *string {
	if o == nil {
		return nil
	}
	return o.TechnicalContactID
}

func (o *CustomerResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
