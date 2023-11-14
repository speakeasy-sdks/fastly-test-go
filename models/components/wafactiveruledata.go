// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

// WafActiveRuleDataStatus - Describes the behavior for the particular rule revision within this firewall version.
type WafActiveRuleDataStatus string

const (
	WafActiveRuleDataStatusLog   WafActiveRuleDataStatus = "log"
	WafActiveRuleDataStatusBlock WafActiveRuleDataStatus = "block"
	WafActiveRuleDataStatusScore WafActiveRuleDataStatus = "score"
)

func (e WafActiveRuleDataStatus) ToPointer() *WafActiveRuleDataStatus {
	return &e
}

func (e *WafActiveRuleDataStatus) UnmarshalJSON(data []byte) error {
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
		*e = WafActiveRuleDataStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafActiveRuleDataStatus: %v", v)
	}
}

type WafActiveRuleDataAttributes struct {
	// The ModSecurity rule ID of the associated rule revision.
	ModsecRuleID *int64                   `json:"modsec_rule_id,omitempty"`
	Revision     *WafRuleRevisionOrLatest `json:"revision,omitempty"`
	// Describes the behavior for the particular rule revision within this firewall version.
	Status *WafActiveRuleDataStatus `json:"status,omitempty"`
}

func (o *WafActiveRuleDataAttributes) GetModsecRuleID() *int64 {
	if o == nil {
		return nil
	}
	return o.ModsecRuleID
}

func (o *WafActiveRuleDataAttributes) GetRevision() *WafRuleRevisionOrLatest {
	if o == nil {
		return nil
	}
	return o.Revision
}

func (o *WafActiveRuleDataAttributes) GetStatus() *WafActiveRuleDataStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type WafActiveRuleData struct {
	Attributes    *WafActiveRuleDataAttributes   `json:"attributes,omitempty"`
	Relationships *RelationshipsForWafActiveRule `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafActiveRule `default:"waf_active_rule" json:"type"`
}

func (w WafActiveRuleData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafActiveRuleData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafActiveRuleData) GetAttributes() *WafActiveRuleDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafActiveRuleData) GetRelationships() *RelationshipsForWafActiveRule {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *WafActiveRuleData) GetType() *TypeWafActiveRule {
	if o == nil {
		return nil
	}
	return o.Type
}

type WafActiveRuleData1 struct {
	Attributes    *WafActiveRuleDataAttributes        `json:"attributes,omitempty"`
	Relationships *RelationshipsForWafActiveRuleInput `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafActiveRule `default:"waf_active_rule" json:"type"`
}

func (w WafActiveRuleData1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafActiveRuleData1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafActiveRuleData1) GetAttributes() *WafActiveRuleDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafActiveRuleData1) GetRelationships() *RelationshipsForWafActiveRuleInput {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *WafActiveRuleData1) GetType() *TypeWafActiveRule {
	if o == nil {
		return nil
	}
	return o.Type
}
