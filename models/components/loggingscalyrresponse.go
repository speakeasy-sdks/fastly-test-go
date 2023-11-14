// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
	"time"
)

// LoggingScalyrResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingScalyrResponseFormatVersion int64

const (
	LoggingScalyrResponseFormatVersionOne LoggingScalyrResponseFormatVersion = 1
	LoggingScalyrResponseFormatVersionTwo LoggingScalyrResponseFormatVersion = 2
)

func (e LoggingScalyrResponseFormatVersion) ToPointer() *LoggingScalyrResponseFormatVersion {
	return &e
}

func (e *LoggingScalyrResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingScalyrResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingScalyrResponseFormatVersion: %v", v)
	}
}

// LoggingScalyrResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingScalyrResponsePlacement string

const (
	LoggingScalyrResponsePlacementNone                   LoggingScalyrResponsePlacement = "none"
	LoggingScalyrResponsePlacementWafDebug               LoggingScalyrResponsePlacement = "waf_debug"
	LoggingScalyrResponsePlacementLessThanNilGreaterThan LoggingScalyrResponsePlacement = "<nil>"
)

func (e LoggingScalyrResponsePlacement) ToPointer() *LoggingScalyrResponsePlacement {
	return &e
}

func (e *LoggingScalyrResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingScalyrResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingScalyrResponsePlacement: %v", v)
	}
}

// LoggingScalyrResponseRegion - The region that log data will be sent to.
type LoggingScalyrResponseRegion string

const (
	LoggingScalyrResponseRegionUs LoggingScalyrResponseRegion = "US"
	LoggingScalyrResponseRegionEu LoggingScalyrResponseRegion = "EU"
)

func (e LoggingScalyrResponseRegion) ToPointer() *LoggingScalyrResponseRegion {
	return &e
}

func (e *LoggingScalyrResponseRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = LoggingScalyrResponseRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingScalyrResponseRegion: %v", v)
	}
}

type LoggingScalyrResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingScalyrResponseFormatVersion `default:"2" json:"format_version"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingScalyrResponsePlacement `json:"placement,omitempty"`
	// The name of the logfile within Scalyr.
	ProjectID *string `default:"logplex" json:"project_id"`
	// The region that log data will be sent to.
	Region *LoggingScalyrResponseRegion `default:"US" json:"region"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// The token to use for authentication ([https://www.scalyr.com/keys](https://www.scalyr.com/keys)).
	Token *string `json:"token,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (l LoggingScalyrResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingScalyrResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingScalyrResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingScalyrResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingScalyrResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingScalyrResponse) GetFormatVersion() *LoggingScalyrResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingScalyrResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingScalyrResponse) GetPlacement() *LoggingScalyrResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingScalyrResponse) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *LoggingScalyrResponse) GetRegion() *LoggingScalyrResponseRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingScalyrResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingScalyrResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingScalyrResponse) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingScalyrResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingScalyrResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}