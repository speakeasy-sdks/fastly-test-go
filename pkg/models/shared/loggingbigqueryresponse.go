// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

// LoggingBigqueryResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingBigqueryResponseFormatVersion int64

const (
	LoggingBigqueryResponseFormatVersionOne LoggingBigqueryResponseFormatVersion = 1
	LoggingBigqueryResponseFormatVersionTwo LoggingBigqueryResponseFormatVersion = 2
)

func (e LoggingBigqueryResponseFormatVersion) ToPointer() *LoggingBigqueryResponseFormatVersion {
	return &e
}

func (e *LoggingBigqueryResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingBigqueryResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingBigqueryResponseFormatVersion: %v", v)
	}
}

// LoggingBigqueryResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingBigqueryResponsePlacement string

const (
	LoggingBigqueryResponsePlacementNone                   LoggingBigqueryResponsePlacement = "none"
	LoggingBigqueryResponsePlacementWafDebug               LoggingBigqueryResponsePlacement = "waf_debug"
	LoggingBigqueryResponsePlacementLessThanNilGreaterThan LoggingBigqueryResponsePlacement = "<nil>"
)

func (e LoggingBigqueryResponsePlacement) ToPointer() *LoggingBigqueryResponsePlacement {
	return &e
}

func (e *LoggingBigqueryResponsePlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		fallthrough
	case "<nil>":
		*e = LoggingBigqueryResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingBigqueryResponsePlacement: %v", v)
	}
}

type LoggingBigqueryResponse struct {
	// The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided.
	AccountName *string `json:"account_name,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Your BigQuery dataset.
	Dataset *string `json:"dataset,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table.
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingBigqueryResponseFormatVersion `default:"2" json:"format_version"`
	// The name of the BigQuery logging object. Used as a primary key for API access.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingBigqueryResponsePlacement `json:"placement,omitempty"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `json:"project_id,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	// Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified.
	SecretKey *string `json:"secret_key,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// Your BigQuery table.
	Table *string `json:"table,omitempty"`
	// BigQuery table name suffix template. Optional.
	TemplateSuffix *string `json:"template_suffix,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (l LoggingBigqueryResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingBigqueryResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingBigqueryResponse) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *LoggingBigqueryResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingBigqueryResponse) GetDataset() *string {
	if o == nil {
		return nil
	}
	return o.Dataset
}

func (o *LoggingBigqueryResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingBigqueryResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingBigqueryResponse) GetFormatVersion() *LoggingBigqueryResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingBigqueryResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingBigqueryResponse) GetPlacement() *LoggingBigqueryResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingBigqueryResponse) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *LoggingBigqueryResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingBigqueryResponse) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingBigqueryResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingBigqueryResponse) GetTable() *string {
	if o == nil {
		return nil
	}
	return o.Table
}

func (o *LoggingBigqueryResponse) GetTemplateSuffix() *string {
	if o == nil {
		return nil
	}
	return o.TemplateSuffix
}

func (o *LoggingBigqueryResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingBigqueryResponse) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *LoggingBigqueryResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
