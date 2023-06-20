// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingHoneycombResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingHoneycombResponseFormatVersion int64

const (
	LoggingHoneycombResponseFormatVersionOne LoggingHoneycombResponseFormatVersion = 1
	LoggingHoneycombResponseFormatVersionTwo LoggingHoneycombResponseFormatVersion = 2
)

func (e LoggingHoneycombResponseFormatVersion) ToPointer() *LoggingHoneycombResponseFormatVersion {
	return &e
}

func (e *LoggingHoneycombResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingHoneycombResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHoneycombResponseFormatVersion: %v", v)
	}
}

// LoggingHoneycombResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingHoneycombResponsePlacement string

const (
	LoggingHoneycombResponsePlacementNone                   LoggingHoneycombResponsePlacement = "none"
	LoggingHoneycombResponsePlacementWafDebug               LoggingHoneycombResponsePlacement = "waf_debug"
	LoggingHoneycombResponsePlacementLessThanNilGreaterThan LoggingHoneycombResponsePlacement = "<nil>"
)

func (e LoggingHoneycombResponsePlacement) ToPointer() *LoggingHoneycombResponsePlacement {
	return &e
}

func (e *LoggingHoneycombResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingHoneycombResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHoneycombResponsePlacement: %v", v)
	}
}

// LoggingHoneycombResponse - OK
type LoggingHoneycombResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The Honeycomb Dataset you want to log to.
	Dataset *string `json:"dataset,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Honeycomb can ingest.
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingHoneycombResponseFormatVersion `json:"format_version,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingHoneycombResponsePlacement `json:"placement,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// The Write Key from the Account page of your Honeycomb account.
	Token *string `json:"token,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}
