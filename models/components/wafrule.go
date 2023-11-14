// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
)

// Publisher - Rule publisher.
type Publisher string

const (
	PublisherFastly    Publisher = "fastly"
	PublisherTrustwave Publisher = "trustwave"
	PublisherOwasp     Publisher = "owasp"
)

func (e Publisher) ToPointer() *Publisher {
	return &e
}

func (e *Publisher) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fastly":
		fallthrough
	case "trustwave":
		fallthrough
	case "owasp":
		*e = Publisher(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Publisher: %v", v)
	}
}

// WafRuleType - The rule's [type](https://docs.fastly.com/en/guides/managing-rules-on-the-fastly-waf#understanding-the-types-of-rules).
type WafRuleType string

const (
	WafRuleTypeStrict    WafRuleType = "strict"
	WafRuleTypeScore     WafRuleType = "score"
	WafRuleTypeThreshold WafRuleType = "threshold"
)

func (e WafRuleType) ToPointer() *WafRuleType {
	return &e
}

func (e *WafRuleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "strict":
		fallthrough
	case "score":
		fallthrough
	case "threshold":
		*e = WafRuleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WafRuleType: %v", v)
	}
}

type WafRuleAttributes struct {
	// Corresponding ModSecurity rule ID.
	ModsecRuleID *int64 `json:"modsec_rule_id,omitempty"`
	// Rule publisher.
	Publisher *Publisher `json:"publisher,omitempty"`
	// The rule's [type](https://docs.fastly.com/en/guides/managing-rules-on-the-fastly-waf#understanding-the-types-of-rules).
	Type *WafRuleType `json:"type,omitempty"`
}

func (o *WafRuleAttributes) GetModsecRuleID() *int64 {
	if o == nil {
		return nil
	}
	return o.ModsecRuleID
}

func (o *WafRuleAttributes) GetPublisher() *Publisher {
	if o == nil {
		return nil
	}
	return o.Publisher
}

func (o *WafRuleAttributes) GetType() *WafRuleType {
	if o == nil {
		return nil
	}
	return o.Type
}

type WafRule struct {
	Attributes *WafRuleAttributes `json:"attributes,omitempty"`
	ID         *string            `json:"id,omitempty"`
	// Resource type.
	Type *TypeWafRule `default:"waf_rule" json:"type"`
}

func (w WafRule) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafRule) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *WafRule) GetAttributes() *WafRuleAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafRule) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WafRule) GetType() *TypeWafRule {
	if o == nil {
		return nil
	}
	return o.Type
}