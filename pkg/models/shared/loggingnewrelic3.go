// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// LoggingNewrelicFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingNewrelicFormatVersion int64

const (
	LoggingNewrelicFormatVersionOne LoggingNewrelicFormatVersion = 1
	LoggingNewrelicFormatVersionTwo LoggingNewrelicFormatVersion = 2
)

func (e LoggingNewrelicFormatVersion) ToPointer() *LoggingNewrelicFormatVersion {
	return &e
}

func (e *LoggingNewrelicFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingNewrelicFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicFormatVersion: %v", v)
	}
}

// LoggingNewrelicPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingNewrelicPlacement string

const (
	LoggingNewrelicPlacementNone                   LoggingNewrelicPlacement = "none"
	LoggingNewrelicPlacementWafDebug               LoggingNewrelicPlacement = "waf_debug"
	LoggingNewrelicPlacementLessThanNilGreaterThan LoggingNewrelicPlacement = "<nil>"
)

func (e LoggingNewrelicPlacement) ToPointer() *LoggingNewrelicPlacement {
	return &e
}

func (e *LoggingNewrelicPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingNewrelicPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicPlacement: %v", v)
	}
}

// LoggingNewrelicRegion - The region to which to stream logs.
type LoggingNewrelicRegion string

const (
	LoggingNewrelicRegionUs LoggingNewrelicRegion = "US"
	LoggingNewrelicRegionEu LoggingNewrelicRegion = "EU"
)

func (e LoggingNewrelicRegion) ToPointer() *LoggingNewrelicRegion {
	return &e
}

func (e *LoggingNewrelicRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = LoggingNewrelicRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicRegion: %v", v)
	}
}

type LoggingNewrelic3 struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that New Relic Logs can ingest.
	Format *string `form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingNewrelicFormatVersion `form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingNewrelicPlacement `form:"name=placement"`
	// The region to which to stream logs.
	Region *LoggingNewrelicRegion `form:"name=region"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// The Insert API key from the Account page of your New Relic account. Required.
	Token *string `form:"name=token"`
}

func (o *LoggingNewrelic3) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingNewrelic3) GetFormatVersion() *LoggingNewrelicFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingNewrelic3) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingNewrelic3) GetPlacement() *LoggingNewrelicPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingNewrelic3) GetRegion() *LoggingNewrelicRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingNewrelic3) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingNewrelic3) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
