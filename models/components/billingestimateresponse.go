// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)

type Line struct {
	Amount          *float64 `json:"amount,omitempty"`
	ClientPlanID    *string  `json:"client_plan_id,omitempty"`
	ClientServiceID *string  `json:"client_service_id,omitempty"`
	Description     *string  `json:"description,omitempty"`
	PerUnitCost     *float64 `json:"per_unit_cost,omitempty"`
	PlanNo          *int64   `json:"plan_no,omitempty"`
	ServiceNo       *float64 `json:"service_no,omitempty"`
	ServiceType     *string  `json:"service_type,omitempty"`
	Units           *float64 `json:"units,omitempty"`
}

func (o *Line) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *Line) GetClientPlanID() *string {
	if o == nil {
		return nil
	}
	return o.ClientPlanID
}

func (o *Line) GetClientServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ClientServiceID
}

func (o *Line) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Line) GetPerUnitCost() *float64 {
	if o == nil {
		return nil
	}
	return o.PerUnitCost
}

func (o *Line) GetPlanNo() *int64 {
	if o == nil {
		return nil
	}
	return o.PlanNo
}

func (o *Line) GetServiceNo() *float64 {
	if o == nil {
		return nil
	}
	return o.ServiceNo
}

func (o *Line) GetServiceType() *string {
	if o == nil {
		return nil
	}
	return o.ServiceType
}

func (o *Line) GetUnits() *float64 {
	if o == nil {
		return nil
	}
	return o.Units
}

type Lines struct {
	Line *Line `json:"line,omitempty"`
}

func (o *Lines) GetLine() *Line {
	if o == nil {
		return nil
	}
	return o.Line
}

type Tiers struct {
	Name  *string  `json:"name,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Total *float64 `json:"total,omitempty"`
	Units *float64 `json:"units,omitempty"`
}

func (o *Tiers) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Tiers) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *Tiers) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *Tiers) GetUnits() *float64 {
	if o == nil {
		return nil
	}
	return o.Units
}

type Regions struct {
	Tiers []Tiers  `json:"tiers,omitempty"`
	Total *float64 `json:"total,omitempty"`
}

func (o *Regions) GetTiers() []Tiers {
	if o == nil {
		return nil
	}
	return o.Tiers
}

func (o *Regions) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}

type BillingEstimateResponseRegions struct {
	AdditionalProperties map[string]Regions `additionalProperties:"true" json:"-"`
	Cost                 *float64           `json:"cost,omitempty"`
}

func (b BillingEstimateResponseRegions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillingEstimateResponseRegions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillingEstimateResponseRegions) GetAdditionalProperties() map[string]Regions {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *BillingEstimateResponseRegions) GetCost() *float64 {
	if o == nil {
		return nil
	}
	return o.Cost
}

// BillingEstimateResponseStatus - What the current status of this invoice can be.
type BillingEstimateResponseStatus string

const (
	BillingEstimateResponseStatusPending     BillingEstimateResponseStatus = "Pending"
	BillingEstimateResponseStatusOutstanding BillingEstimateResponseStatus = "Outstanding"
	BillingEstimateResponseStatusPaid        BillingEstimateResponseStatus = "Paid"
	BillingEstimateResponseStatusMtd         BillingEstimateResponseStatus = "MTD"
)

func (e BillingEstimateResponseStatus) ToPointer() *BillingEstimateResponseStatus {
	return &e
}

func (e *BillingEstimateResponseStatus) UnmarshalJSON(data []byte) error {
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
		*e = BillingEstimateResponseStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillingEstimateResponseStatus: %v", v)
	}
}

type Status struct {
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	SentAt *time.Time `json:"sent_at,omitempty"`
	// What the current status of this invoice can be.
	Status *BillingEstimateResponseStatus `json:"status,omitempty"`
}

func (s Status) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Status) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Status) GetSentAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.SentAt
}

func (o *Status) GetStatus() *BillingEstimateResponseStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type Extras struct {
	// The name of this extra cost.
	Name *string `json:"name,omitempty"`
	// Recurring monthly cost in USD. Not present if $0.0.
	Recurring *float64 `json:"recurring,omitempty"`
	// Initial set up cost in USD. Not present if $0.0 or this is not the month the extra was added.
	Setup *float64 `json:"setup,omitempty"`
}

func (o *Extras) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Extras) GetRecurring() *float64 {
	if o == nil {
		return nil
	}
	return o.Recurring
}

func (o *Extras) GetSetup() *float64 {
	if o == nil {
		return nil
	}
	return o.Setup
}

// Total - Complete summary of the billing information.
type Total struct {
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
	Extras []Extras `json:"extras,omitempty"`
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

func (o *Total) GetBandwidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Bandwidth
}

func (o *Total) GetBandwidthCost() *float64 {
	if o == nil {
		return nil
	}
	return o.BandwidthCost
}

func (o *Total) GetBandwidthUnits() *string {
	if o == nil {
		return nil
	}
	return o.BandwidthUnits
}

func (o *Total) GetCost() *float64 {
	if o == nil {
		return nil
	}
	return o.Cost
}

func (o *Total) GetCostBeforeDiscount() *float64 {
	if o == nil {
		return nil
	}
	return o.CostBeforeDiscount
}

func (o *Total) GetDiscount() *float64 {
	if o == nil {
		return nil
	}
	return o.Discount
}

func (o *Total) GetExtras() []Extras {
	if o == nil {
		return nil
	}
	return o.Extras
}

func (o *Total) GetExtrasCost() *float64 {
	if o == nil {
		return nil
	}
	return o.ExtrasCost
}

func (o *Total) GetIncurredCost() *float64 {
	if o == nil {
		return nil
	}
	return o.IncurredCost
}

func (o *Total) GetOverage() *float64 {
	if o == nil {
		return nil
	}
	return o.Overage
}

func (o *Total) GetPlanCode() *string {
	if o == nil {
		return nil
	}
	return o.PlanCode
}

func (o *Total) GetPlanMinimum() *float64 {
	if o == nil {
		return nil
	}
	return o.PlanMinimum
}

func (o *Total) GetPlanName() *string {
	if o == nil {
		return nil
	}
	return o.PlanName
}

func (o *Total) GetRequests() *float64 {
	if o == nil {
		return nil
	}
	return o.Requests
}

func (o *Total) GetRequestsCost() *float64 {
	if o == nil {
		return nil
	}
	return o.RequestsCost
}

func (o *Total) GetTerms() *string {
	if o == nil {
		return nil
	}
	return o.Terms
}

type BillingEstimateResponse struct {
	CustomerID *string `json:"customer_id,omitempty"`
	// Date and time in ISO 8601 format.
	EndTime   *time.Time `json:"end_time,omitempty"`
	InvoiceID *string    `json:"invoice_id,omitempty"`
	Lines     []Lines    `json:"lines,omitempty"`
	// Breakdown of regional data for products that are region based.
	Regions map[string]BillingEstimateResponseRegions `json:"regions,omitempty"`
	// Date and time in ISO 8601 format.
	StartTime *time.Time `json:"start_time,omitempty"`
	Status    *Status    `json:"status,omitempty"`
	// Complete summary of the billing information.
	Total *Total `json:"total,omitempty"`
	// The current state of our third-party billing vendor. One of `up` or `down`.
	VendorState *string `json:"vendor_state,omitempty"`
}

func (b BillingEstimateResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillingEstimateResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillingEstimateResponse) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *BillingEstimateResponse) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *BillingEstimateResponse) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *BillingEstimateResponse) GetLines() []Lines {
	if o == nil {
		return nil
	}
	return o.Lines
}

func (o *BillingEstimateResponse) GetRegions() map[string]BillingEstimateResponseRegions {
	if o == nil {
		return nil
	}
	return o.Regions
}

func (o *BillingEstimateResponse) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *BillingEstimateResponse) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *BillingEstimateResponse) GetTotal() *Total {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *BillingEstimateResponse) GetVendorState() *string {
	if o == nil {
		return nil
	}
	return o.VendorState
}