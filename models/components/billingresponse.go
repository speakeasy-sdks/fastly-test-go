// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)

type BillingResponseSchemasRegions struct {
	AdditionalProperties map[string]BillingResponseRegions `additionalProperties:"true" json:"-"`
	Cost                 *float64                          `json:"cost,omitempty"`
}

func (b BillingResponseSchemasRegions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillingResponseSchemasRegions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillingResponseSchemasRegions) GetAdditionalProperties() map[string]BillingResponseRegions {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *BillingResponseSchemasRegions) GetCost() *float64 {
	if o == nil {
		return nil
	}
	return o.Cost
}

// BillingResponseSchemasStatus - What the current status of this invoice can be.
type BillingResponseSchemasStatus string

const (
	BillingResponseSchemasStatusPending     BillingResponseSchemasStatus = "Pending"
	BillingResponseSchemasStatusOutstanding BillingResponseSchemasStatus = "Outstanding"
	BillingResponseSchemasStatusPaid        BillingResponseSchemasStatus = "Paid"
	BillingResponseSchemasStatusMtd         BillingResponseSchemasStatus = "MTD"
)

func (e BillingResponseSchemasStatus) ToPointer() *BillingResponseSchemasStatus {
	return &e
}

func (e *BillingResponseSchemasStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Outstanding":
		fallthrough
	case "Paid":
		fallthrough
	case "MTD":
		*e = BillingResponseSchemasStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillingResponseSchemasStatus: %v", v)
	}
}

type BillingResponseStatus struct {
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	SentAt *time.Time `json:"sent_at,omitempty"`
	// What the current status of this invoice can be.
	Status *BillingResponseSchemasStatus `json:"status,omitempty"`
}

func (b BillingResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillingResponseStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillingResponseStatus) GetSentAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.SentAt
}

func (o *BillingResponseStatus) GetStatus() *BillingResponseSchemasStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type BillingResponseExtras struct {
	// The name of this extra cost.
	Name *string `json:"name,omitempty"`
	// Recurring monthly cost in USD. Not present if $0.0.
	Recurring *float64 `json:"recurring,omitempty"`
	// Initial set up cost in USD. Not present if $0.0 or this is not the month the extra was added.
	Setup *float64 `json:"setup,omitempty"`
}

func (o *BillingResponseExtras) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *BillingResponseExtras) GetRecurring() *float64 {
	if o == nil {
		return nil
	}
	return o.Recurring
}

func (o *BillingResponseExtras) GetSetup() *float64 {
	if o == nil {
		return nil
	}
	return o.Setup
}

// BillingResponseTotal - Complete summary of the billing information.
type BillingResponseTotal struct {
	// The total amount of bandwidth used this month (See bandwidth_units for measurement).
	Bandwidth *float64 `json:"bandwidth,omitempty"`
	// The cost of the bandwidth used this month in USD.
	BandwidthCost *float64 `json:"bandwidth_cost,omitempty"`
	// Bandwidth measurement units based on billing plan.
	BandwidthUnits *string `json:"bandwidth_units,omitempty"`
	// The final amount to be paid.
	Cost *float64 `json:"cost,omitempty"`
	// Total incurred cost plus extras cost.
	CostBeforeDiscount *float64 `json:"cost_before_discount,omitempty"`
	// Calculated discount rate.
	Discount *float64 `json:"discount,omitempty"`
	// A list of any extras for this invoice.
	Extras []BillingResponseExtras `json:"extras,omitempty"`
	// Total cost of all extras.
	ExtrasCost *float64 `json:"extras_cost,omitempty"`
	// The total cost of bandwidth and requests used this month.
	IncurredCost *float64 `json:"incurred_cost,omitempty"`
	// How much over the plan minimum has been incurred.
	Overage *float64 `json:"overage,omitempty"`
	// The short code the plan this invoice was generated under.
	PlanCode *string `json:"plan_code,omitempty"`
	// The minimum cost of this plan.
	PlanMinimum *float64 `json:"plan_minimum,omitempty"`
	// The name of the plan this invoice was generated under.
	PlanName *string `json:"plan_name,omitempty"`
	// The total number of requests used this month.
	Requests *float64 `json:"requests,omitempty"`
	// The cost of the requests used this month.
	RequestsCost *float64 `json:"requests_cost,omitempty"`
	// Payment terms. Almost always Net15.
	Terms *string `json:"terms,omitempty"`
}

func (o *BillingResponseTotal) GetBandwidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Bandwidth
}

func (o *BillingResponseTotal) GetBandwidthCost() *float64 {
	if o == nil {
		return nil
	}
	return o.BandwidthCost
}

func (o *BillingResponseTotal) GetBandwidthUnits() *string {
	if o == nil {
		return nil
	}
	return o.BandwidthUnits
}

func (o *BillingResponseTotal) GetCost() *float64 {
	if o == nil {
		return nil
	}
	return o.Cost
}

func (o *BillingResponseTotal) GetCostBeforeDiscount() *float64 {
	if o == nil {
		return nil
	}
	return o.CostBeforeDiscount
}

func (o *BillingResponseTotal) GetDiscount() *float64 {
	if o == nil {
		return nil
	}
	return o.Discount
}

func (o *BillingResponseTotal) GetExtras() []BillingResponseExtras {
	if o == nil {
		return nil
	}
	return o.Extras
}

func (o *BillingResponseTotal) GetExtrasCost() *float64 {
	if o == nil {
		return nil
	}
	return o.ExtrasCost
}

func (o *BillingResponseTotal) GetIncurredCost() *float64 {
	if o == nil {
		return nil
	}
	return o.IncurredCost
}

func (o *BillingResponseTotal) GetOverage() *float64 {
	if o == nil {
		return nil
	}
	return o.Overage
}

func (o *BillingResponseTotal) GetPlanCode() *string {
	if o == nil {
		return nil
	}
	return o.PlanCode
}

func (o *BillingResponseTotal) GetPlanMinimum() *float64 {
	if o == nil {
		return nil
	}
	return o.PlanMinimum
}

func (o *BillingResponseTotal) GetPlanName() *string {
	if o == nil {
		return nil
	}
	return o.PlanName
}

func (o *BillingResponseTotal) GetRequests() *float64 {
	if o == nil {
		return nil
	}
	return o.Requests
}

func (o *BillingResponseTotal) GetRequestsCost() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestsCost
}

func (o *BillingResponseTotal) GetTerms() *string {
	if o == nil {
		return nil
	}
	return o.Terms
}

type BillingResponse struct {
	CustomerID *string `json:"customer_id,omitempty"`
	// Date and time in ISO 8601 format.
	EndTime   *time.Time                `json:"end_time,omitempty"`
	InvoiceID *string                   `json:"invoice_id,omitempty"`
	LineItems []BillingResponseLineItem `json:"line_items,omitempty"`
	// Breakdown of regional data for products that are region based.
	Regions map[string]BillingResponseSchemasRegions `json:"regions,omitempty"`
	// Date and time in ISO 8601 format.
	StartTime *time.Time             `json:"start_time,omitempty"`
	Status    *BillingResponseStatus `json:"status,omitempty"`
	// Complete summary of the billing information.
	Total *BillingResponseTotal `json:"total,omitempty"`
	// The current state of our third-party billing vendor. One of `up` or `down`.
	VendorState *string `json:"vendor_state,omitempty"`
}

func (b BillingResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillingResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillingResponse) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *BillingResponse) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *BillingResponse) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *BillingResponse) GetLineItems() []BillingResponseLineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *BillingResponse) GetRegions() map[string]BillingResponseSchemasRegions {
	if o == nil {
		return nil
	}
	return o.Regions
}

func (o *BillingResponse) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *BillingResponse) GetStatus() *BillingResponseStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *BillingResponse) GetTotal() *BillingResponseTotal {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *BillingResponse) GetVendorState() *string {
	if o == nil {
		return nil
	}
	return o.VendorState
}
