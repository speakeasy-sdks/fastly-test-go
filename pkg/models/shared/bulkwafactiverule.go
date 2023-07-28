// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// BulkWafActiveRuleAttributesStatus - Describes the behavior for the particular rule revision within this firewall version.
type BulkWafActiveRuleAttributesStatus string

const (
	BulkWafActiveRuleAttributesStatusLog   BulkWafActiveRuleAttributesStatus = "log"
	BulkWafActiveRuleAttributesStatusBlock BulkWafActiveRuleAttributesStatus = "block"
	BulkWafActiveRuleAttributesStatusScore BulkWafActiveRuleAttributesStatus = "score"
)

func (e BulkWafActiveRuleAttributesStatus) ToPointer() *BulkWafActiveRuleAttributesStatus {
	return &e
}

func (e *BulkWafActiveRuleAttributesStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "log":
		fallthrough
	case "block":
		fallthrough
	case "score":
		*e = BulkWafActiveRuleAttributesStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BulkWafActiveRuleAttributesStatus: %v", v)
	}
}

type BulkWafActiveRuleAttributes struct {
	// The ModSecurity rule ID of the associated rule revision.
	ModsecRuleID *int64                   `json:"modsec_rule_id,omitempty"`
	Revision     *WafRuleRevisionOrLatest `json:"revision,omitempty"`
	// Describes the behavior for the particular rule revision within this firewall version.
	Status *BulkWafActiveRuleAttributesStatus `json:"status,omitempty"`
}

func (o *BulkWafActiveRuleAttributes) GetModsecRuleID() *int64 {
	if o == nil {
		return nil
	}
	return o.ModsecRuleID
}

func (o *BulkWafActiveRuleAttributes) GetRevision() *WafRuleRevisionOrLatest {
	if o == nil {
		return nil
	}
	return o.Revision
}

func (o *BulkWafActiveRuleAttributes) GetStatus() *BulkWafActiveRuleAttributesStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type BulkWafActiveRuleInput struct {
	Attributes    *BulkWafActiveRuleAttributes `json:"attributes,omitempty"`
	Relationships interface{}                  `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafActiveRule `json:"type,omitempty"`
}

func (o *BulkWafActiveRuleInput) GetAttributes() *BulkWafActiveRuleAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *BulkWafActiveRuleInput) GetRelationships() interface{} {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *BulkWafActiveRuleInput) GetType() *TypeWafActiveRule {
	if o == nil {
		return nil
	}
	return o.Type
}
