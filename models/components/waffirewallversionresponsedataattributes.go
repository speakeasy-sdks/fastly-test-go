// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// LastDeploymentStatus - The status of the last deployment of this firewall version.
type LastDeploymentStatus string

const (
	LastDeploymentStatusPending    LastDeploymentStatus = "pending"
	LastDeploymentStatusInProgress LastDeploymentStatus = "in progress"
	LastDeploymentStatusCompleted  LastDeploymentStatus = "completed"
	LastDeploymentStatusFailed     LastDeploymentStatus = "failed"
)

func (e LastDeploymentStatus) ToPointer() *LastDeploymentStatus {
	return &e
}

func (e *LastDeploymentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "in progress":
		fallthrough
	case "completed":
		fallthrough
	case "failed":
		*e = LastDeploymentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LastDeploymentStatus: %v", v)
	}
}

type WafFirewallVersionResponseDataAttributes struct {
	// Whether a specific firewall version is currently deployed.
	Active *bool `json:"active,omitempty"`
	// The number of active Fastly rules set to block.
	ActiveRulesFastlyBlockCount *int64 `json:"active_rules_fastly_block_count,omitempty"`
	// The number of active Fastly rules set to log.
	ActiveRulesFastlyLogCount *int64 `json:"active_rules_fastly_log_count,omitempty"`
	// The number of active Fastly rules set to score.
	ActiveRulesFastlyScoreCount *int64 `json:"active_rules_fastly_score_count,omitempty"`
	// The number of active OWASP rules set to block.
	ActiveRulesOwaspBlockCount *int64 `json:"active_rules_owasp_block_count,omitempty"`
	// The number of active OWASP rules set to log.
	ActiveRulesOwaspLogCount *int64 `json:"active_rules_owasp_log_count,omitempty"`
	// The number of active OWASP rules set to score.
	ActiveRulesOwaspScoreCount *int64 `json:"active_rules_owasp_score_count,omitempty"`
	// The number of active Trustwave rules set to block.
	ActiveRulesTrustwaveBlockCount *int64 `json:"active_rules_trustwave_block_count,omitempty"`
	// The number of active Trustwave rules set to log.
	ActiveRulesTrustwaveLogCount *int64 `json:"active_rules_trustwave_log_count,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Time-stamp (GMT) indicating when the firewall version was last deployed.
	DeployedAt *string `json:"deployed_at,omitempty"`
	// Contains error message if the firewall version fails to deploy.
	Error *string `json:"error,omitempty"`
	// The status of the last deployment of this firewall version.
	LastDeploymentStatus *LastDeploymentStatus `json:"last_deployment_status,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (w WafFirewallVersionResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafFirewallVersionResponseDataAttributes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafFirewallVersionResponseDataAttributes) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyBlockCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesFastlyBlockCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyLogCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesFastlyLogCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesFastlyScoreCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesFastlyScoreCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspBlockCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesOwaspBlockCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspLogCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesOwaspLogCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesOwaspScoreCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesOwaspScoreCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveBlockCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesTrustwaveBlockCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetActiveRulesTrustwaveLogCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveRulesTrustwaveLogCount
}

func (o *WafFirewallVersionResponseDataAttributes) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WafFirewallVersionResponseDataAttributes) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *WafFirewallVersionResponseDataAttributes) GetDeployedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeployedAt
}

func (o *WafFirewallVersionResponseDataAttributes) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *WafFirewallVersionResponseDataAttributes) GetLastDeploymentStatus() *LastDeploymentStatus {
	if o == nil {
		return nil
	}
	return o.LastDeploymentStatus
}

func (o *WafFirewallVersionResponseDataAttributes) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
