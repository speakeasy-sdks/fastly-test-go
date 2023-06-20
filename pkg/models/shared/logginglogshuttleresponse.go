// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingLogshuttleResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingLogshuttleResponseFormatVersion int64

const (
	LoggingLogshuttleResponseFormatVersionOne LoggingLogshuttleResponseFormatVersion = 1
	LoggingLogshuttleResponseFormatVersionTwo LoggingLogshuttleResponseFormatVersion = 2
)

func (e LoggingLogshuttleResponseFormatVersion) ToPointer() *LoggingLogshuttleResponseFormatVersion {
	return &e
}

func (e *LoggingLogshuttleResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingLogshuttleResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingLogshuttleResponseFormatVersion: %v", v)
	}
}

// LoggingLogshuttleResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingLogshuttleResponsePlacement string

const (
	LoggingLogshuttleResponsePlacementNone                   LoggingLogshuttleResponsePlacement = "none"
	LoggingLogshuttleResponsePlacementWafDebug               LoggingLogshuttleResponsePlacement = "waf_debug"
	LoggingLogshuttleResponsePlacementLessThanNilGreaterThan LoggingLogshuttleResponsePlacement = "<nil>"
)

func (e LoggingLogshuttleResponsePlacement) ToPointer() *LoggingLogshuttleResponsePlacement {
	return &e
}

func (e *LoggingLogshuttleResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingLogshuttleResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingLogshuttleResponsePlacement: %v", v)
	}
}

// LoggingLogshuttleResponse - OK
type LoggingLogshuttleResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingLogshuttleResponseFormatVersion `json:"format_version,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingLogshuttleResponsePlacement `json:"placement,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// The data authentication token associated with this endpoint.
	Token *string `json:"token,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The URL to stream logs to.
	URL     *string `json:"url,omitempty"`
	Version *int64  `json:"version,omitempty"`
}
