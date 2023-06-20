// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingSumologicResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingSumologicResponseFormatVersion int64

const (
	LoggingSumologicResponseFormatVersionOne LoggingSumologicResponseFormatVersion = 1
	LoggingSumologicResponseFormatVersionTwo LoggingSumologicResponseFormatVersion = 2
)

func (e LoggingSumologicResponseFormatVersion) ToPointer() *LoggingSumologicResponseFormatVersion {
	return &e
}

func (e *LoggingSumologicResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingSumologicResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSumologicResponseFormatVersion: %v", v)
	}
}

// LoggingSumologicResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingSumologicResponsePlacement string

const (
	LoggingSumologicResponsePlacementNone                   LoggingSumologicResponsePlacement = "none"
	LoggingSumologicResponsePlacementWafDebug               LoggingSumologicResponsePlacement = "waf_debug"
	LoggingSumologicResponsePlacementLessThanNilGreaterThan LoggingSumologicResponsePlacement = "<nil>"
)

func (e LoggingSumologicResponsePlacement) ToPointer() *LoggingSumologicResponsePlacement {
	return &e
}

func (e *LoggingSumologicResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSumologicResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSumologicResponsePlacement: %v", v)
	}
}

// LoggingSumologicResponse - OK
type LoggingSumologicResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingSumologicResponseFormatVersion `json:"format_version,omitempty"`
	// How the message should be formatted.
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingSumologicResponsePlacement `json:"placement,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The URL to post logs to.
	URL     *string `json:"url,omitempty"`
	Version *int64  `json:"version,omitempty"`
}
