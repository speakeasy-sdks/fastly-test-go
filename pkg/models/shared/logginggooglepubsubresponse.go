// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingGooglePubsubResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingGooglePubsubResponseFormatVersion int64

const (
	LoggingGooglePubsubResponseFormatVersionOne LoggingGooglePubsubResponseFormatVersion = 1
	LoggingGooglePubsubResponseFormatVersionTwo LoggingGooglePubsubResponseFormatVersion = 2
)

func (e LoggingGooglePubsubResponseFormatVersion) ToPointer() *LoggingGooglePubsubResponseFormatVersion {
	return &e
}

func (e *LoggingGooglePubsubResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingGooglePubsubResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGooglePubsubResponseFormatVersion: %v", v)
	}
}

// LoggingGooglePubsubResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingGooglePubsubResponsePlacement string

const (
	LoggingGooglePubsubResponsePlacementNone                   LoggingGooglePubsubResponsePlacement = "none"
	LoggingGooglePubsubResponsePlacementWafDebug               LoggingGooglePubsubResponsePlacement = "waf_debug"
	LoggingGooglePubsubResponsePlacementLessThanNilGreaterThan LoggingGooglePubsubResponsePlacement = "<nil>"
)

func (e LoggingGooglePubsubResponsePlacement) ToPointer() *LoggingGooglePubsubResponsePlacement {
	return &e
}

func (e *LoggingGooglePubsubResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingGooglePubsubResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGooglePubsubResponsePlacement: %v", v)
	}
}

// LoggingGooglePubsubResponse - OK
type LoggingGooglePubsubResponse struct {
	// The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided.
	AccountName *string `json:"account_name,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingGooglePubsubResponseFormatVersion `json:"format_version,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingGooglePubsubResponsePlacement `json:"placement,omitempty"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `json:"project_id,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	// Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified.
	SecretKey *string `json:"secret_key,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// The Google Cloud Pub/Sub topic to which logs will be published. Required.
	Topic *string `json:"topic,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (o *LoggingGooglePubsubResponse) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *LoggingGooglePubsubResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingGooglePubsubResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingGooglePubsubResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingGooglePubsubResponse) GetFormatVersion() *LoggingGooglePubsubResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingGooglePubsubResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingGooglePubsubResponse) GetPlacement() *LoggingGooglePubsubResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingGooglePubsubResponse) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *LoggingGooglePubsubResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingGooglePubsubResponse) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingGooglePubsubResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingGooglePubsubResponse) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}

func (o *LoggingGooglePubsubResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingGooglePubsubResponse) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *LoggingGooglePubsubResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
