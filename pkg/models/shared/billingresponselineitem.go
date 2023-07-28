// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type BillingResponseLineItem struct {
	Amount          *float64 `json:"amount,omitempty"`
	AriaInvoiceID   *string  `json:"aria_invoice_id,omitempty"`
	ClientServiceID *string  `json:"client_service_id,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	CreditCouponCode *string    `json:"credit_coupon_code,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
	Description        *string    `json:"description,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	LineNumber         *int64     `json:"line_number,omitempty"`
	PlanName           *string    `json:"plan_name,omitempty"`
	PlanNo             *float64   `json:"plan_no,omitempty"`
	RatePerUnit        *float64   `json:"rate_per_unit,omitempty"`
	RateScheduleNo     *float64   `json:"rate_schedule_no,omitempty"`
	RateScheduleTierNo *float64   `json:"rate_schedule_tier_no,omitempty"`
	ServiceName        *string    `json:"service_name,omitempty"`
	ServiceNo          *float64   `json:"service_no,omitempty"`
	Units              *float64   `json:"units,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	UsageTypeCd *string    `json:"usage_type_cd,omitempty"`
	UsageTypeNo *float64   `json:"usage_type_no,omitempty"`
}

func (o *BillingResponseLineItem) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *BillingResponseLineItem) GetAriaInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.AriaInvoiceID
}

func (o *BillingResponseLineItem) GetClientServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ClientServiceID
}

func (o *BillingResponseLineItem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BillingResponseLineItem) GetCreditCouponCode() *string {
	if o == nil {
		return nil
	}
	return o.CreditCouponCode
}

func (o *BillingResponseLineItem) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *BillingResponseLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *BillingResponseLineItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BillingResponseLineItem) GetLineNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

func (o *BillingResponseLineItem) GetPlanName() *string {
	if o == nil {
		return nil
	}
	return o.PlanName
}

func (o *BillingResponseLineItem) GetPlanNo() *float64 {
	if o == nil {
		return nil
	}
	return o.PlanNo
}

func (o *BillingResponseLineItem) GetRatePerUnit() *float64 {
	if o == nil {
		return nil
	}
	return o.RatePerUnit
}

func (o *BillingResponseLineItem) GetRateScheduleNo() *float64 {
	if o == nil {
		return nil
	}
	return o.RateScheduleNo
}

func (o *BillingResponseLineItem) GetRateScheduleTierNo() *float64 {
	if o == nil {
		return nil
	}
	return o.RateScheduleTierNo
}

func (o *BillingResponseLineItem) GetServiceName() *string {
	if o == nil {
		return nil
	}
	return o.ServiceName
}

func (o *BillingResponseLineItem) GetServiceNo() *float64 {
	if o == nil {
		return nil
	}
	return o.ServiceNo
}

func (o *BillingResponseLineItem) GetUnits() *float64 {
	if o == nil {
		return nil
	}
	return o.Units
}

func (o *BillingResponseLineItem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *BillingResponseLineItem) GetUsageTypeCd() *string {
	if o == nil {
		return nil
	}
	return o.UsageTypeCd
}

func (o *BillingResponseLineItem) GetUsageTypeNo() *float64 {
	if o == nil {
		return nil
	}
	return o.UsageTypeNo
}
