// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// WafRuleRevisionResponseDataState - The state, indicating if the revision is the most recent version of the rule.
type WafRuleRevisionResponseDataState string

const (
	WafRuleRevisionResponseDataStateLatest   WafRuleRevisionResponseDataState = "latest"
	WafRuleRevisionResponseDataStateOutdated WafRuleRevisionResponseDataState = "outdated"
)

func (e WafRuleRevisionResponseDataState) ToPointer() *WafRuleRevisionResponseDataState {
	return &e
}

func (e *WafRuleRevisionResponseDataState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "latest":
		fallthrough
	case "outdated":
		*e = WafRuleRevisionResponseDataState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafRuleRevisionResponseDataState: %v", v)
	}
}

type WafRuleRevisionResponseDataAttributes struct {
	// Message metadata for the rule.
	Message *string `json:"message,omitempty"`
	// Corresponding ModSecurity rule ID.
	ModsecRuleID *int64 `json:"modsec_rule_id,omitempty"`
	// Paranoia level for the rule.
	ParanoiaLevel *int64 `json:"paranoia_level,omitempty"`
	// Revision number.
	Revision *int64 `json:"revision,omitempty"`
	// Severity metadata for the rule.
	Severity *int64 `json:"severity,omitempty"`
	// The ModSecurity rule logic.
	Source *string `json:"source,omitempty"`
	// The state, indicating if the revision is the most recent version of the rule.
	State *WafRuleRevisionResponseDataState `json:"state,omitempty"`
	// The VCL representation of the rule logic.
	Vcl *string `json:"vcl,omitempty"`
}

func (o *WafRuleRevisionResponseDataAttributes) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *WafRuleRevisionResponseDataAttributes) GetModsecRuleID() *int64 {
	if o == nil {
		return nil
	}
	return o.ModsecRuleID
}

func (o *WafRuleRevisionResponseDataAttributes) GetParanoiaLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.ParanoiaLevel
}

func (o *WafRuleRevisionResponseDataAttributes) GetRevision() *int64 {
	if o == nil {
		return nil
	}
	return o.Revision
}

func (o *WafRuleRevisionResponseDataAttributes) GetSeverity() *int64 {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *WafRuleRevisionResponseDataAttributes) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *WafRuleRevisionResponseDataAttributes) GetState() *WafRuleRevisionResponseDataState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *WafRuleRevisionResponseDataAttributes) GetVcl() *string {
	if o == nil {
		return nil
	}
	return o.Vcl
}

type WafRuleRevisionResponseData struct {
	Attributes *WafRuleRevisionResponseDataAttributes `json:"attributes,omitempty"`
	// Alphanumeric string identifying a WAF rule revision.
	ID            *string              `json:"id,omitempty"`
	Relationships *RelationshipWafRule `json:"relationships,omitempty"`
	// Resource type.
	Type *TypeWafRuleRevision `default:"waf_rule_revision" json:"type"`
}

func (w WafRuleRevisionResponseData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafRuleRevisionResponseData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafRuleRevisionResponseData) GetAttributes() *WafRuleRevisionResponseDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafRuleRevisionResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WafRuleRevisionResponseData) GetRelationships() *RelationshipWafRule {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *WafRuleRevisionResponseData) GetType() *TypeWafRuleRevision {
	if o == nil {
		return nil
	}
	return o.Type
}
