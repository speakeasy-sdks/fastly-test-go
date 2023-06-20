// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type BillingEstimateResponseLinesLine struct {
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

type BillingEstimateResponseLines struct {
	Line *BillingEstimateResponseLinesLine `json:"line,omitempty"`
}

type BillingEstimateResponseRegionsTiers struct {
	Name  *string  `json:"name,omitempty"`
	Price *float64 `json:"price,omitempty"`
	Total *float64 `json:"total,omitempty"`
	Units *float64 `json:"units,omitempty"`
}

type BillingEstimateResponseRegions struct {
	Tiers []BillingEstimateResponseRegionsTiers `json:"tiers,omitempty"`
	Total *float64                              `json:"total,omitempty"`
}

// BillingEstimateResponseStatusStatus - What the current status of this invoice can be.
type BillingEstimateResponseStatusStatus string

const (
	BillingEstimateResponseStatusStatusPending     BillingEstimateResponseStatusStatus = "Pending"
	BillingEstimateResponseStatusStatusOutstanding BillingEstimateResponseStatusStatus = "Outstanding"
	BillingEstimateResponseStatusStatusPaid        BillingEstimateResponseStatusStatus = "Paid"
	BillingEstimateResponseStatusStatusMtd         BillingEstimateResponseStatusStatus = "MTD"
)

func (e BillingEstimateResponseStatusStatus) ToPointer() *BillingEstimateResponseStatusStatus {
	return &e
}

func (e *BillingEstimateResponseStatusStatus) UnmarshalJSON(data []byte) error {
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
		*e = BillingEstimateResponseStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillingEstimateResponseStatusStatus: %v", v)
	}
}

type BillingEstimateResponseStatus struct {
	// Deprecated: this field will be removed in a future release, please migrate away from it as soon as possible.
	SentAt *time.Time `json:"sent_at,omitempty"`
	// What the current status of this invoice can be.
	Status *BillingEstimateResponseStatusStatus `json:"status,omitempty"`
}

type BillingEstimateResponseTotalExtras struct {
	// The name of this extra cost.
	Name *string `json:"name,omitempty"`
	// Recurring monthly cost in USD. Not present if $0.0.
	Recurring *float64 `json:"recurring,omitempty"`
	// Initial set up cost in USD. Not present if $0.0 or this is not the month the extra was added.
	Setup *float64 `json:"setup,omitempty"`
}

// BillingEstimateResponseTotal - Complete summary of the billing information.
type BillingEstimateResponseTotal struct {
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
	Extras []BillingEstimateResponseTotalExtras `json:"extras,omitempty"`
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

// BillingEstimateResponse - OK
type BillingEstimateResponse struct {
	CustomerID *string `json:"customer_id,omitempty"`
	// Date and time in ISO 8601 format.
	EndTime   *time.Time                     `json:"end_time,omitempty"`
	InvoiceID *string                        `json:"invoice_id,omitempty"`
	Lines     []BillingEstimateResponseLines `json:"lines,omitempty"`
	// Breakdown of regional data for products that are region based.
	Regions map[string]map[string]BillingEstimateResponseRegions `json:"regions,omitempty"`
	// Date and time in ISO 8601 format.
	StartTime *time.Time                     `json:"start_time,omitempty"`
	Status    *BillingEstimateResponseStatus `json:"status,omitempty"`
	// Complete summary of the billing information.
	Total *BillingEstimateResponseTotal `json:"total,omitempty"`
	// The current state of our third-party billing vendor. One of `up` or `down`.
	VendorState *string `json:"vendor_state,omitempty"`
}
