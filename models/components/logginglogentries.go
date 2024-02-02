// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// LoggingLogentriesFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingLogentriesFormatVersion int64

const (
	LoggingLogentriesFormatVersionOne LoggingLogentriesFormatVersion = 1
	LoggingLogentriesFormatVersionTwo LoggingLogentriesFormatVersion = 2
)

func (e LoggingLogentriesFormatVersion) ToPointer() *LoggingLogentriesFormatVersion {
	return &e
}

func (e *LoggingLogentriesFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingLogentriesFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingLogentriesFormatVersion: %v", v)
	}
}

// LoggingLogentriesPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingLogentriesPlacement string

const (
	LoggingLogentriesPlacementNone                   LoggingLogentriesPlacement = "none"
	LoggingLogentriesPlacementWafDebug               LoggingLogentriesPlacement = "waf_debug"
	LoggingLogentriesPlacementLessThanNilGreaterThan LoggingLogentriesPlacement = "<nil>"
)

func (e LoggingLogentriesPlacement) ToPointer() *LoggingLogentriesPlacement {
	return &e
}

func (e *LoggingLogentriesPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingLogentriesPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingLogentriesPlacement: %v", v)
	}
}

// LoggingLogentriesRegion - The region to which to stream logs.
type LoggingLogentriesRegion string

const (
	LoggingLogentriesRegionUs  LoggingLogentriesRegion = "US"
	LoggingLogentriesRegionUs2 LoggingLogentriesRegion = "US-2"
	LoggingLogentriesRegionUs3 LoggingLogentriesRegion = "US-3"
	LoggingLogentriesRegionEu  LoggingLogentriesRegion = "EU"
	LoggingLogentriesRegionCa  LoggingLogentriesRegion = "CA"
	LoggingLogentriesRegionAu  LoggingLogentriesRegion = "AU"
	LoggingLogentriesRegionAp  LoggingLogentriesRegion = "AP"
)

func (e LoggingLogentriesRegion) ToPointer() *LoggingLogentriesRegion {
	return &e
}

func (e *LoggingLogentriesRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "US-2":
		fallthrough
	case "US-3":
		fallthrough
	case "EU":
		fallthrough
	case "CA":
		fallthrough
	case "AU":
		fallthrough
	case "AP":
		*e = LoggingLogentriesRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingLogentriesRegion: %v", v)
	}
}

type LoggingLogentries struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingLogentriesFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingLogentriesPlacement `form:"name=placement"`
	// The port number.
	Port *int64 `default:"20000" form:"name=port"`
	// The region to which to stream logs.
	Region *LoggingLogentriesRegion `form:"name=region"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// Use token based authentication ([https://logentries.com/doc/input-token/](https://logentries.com/doc/input-token/)).
	Token *string `form:"name=token"`
	// Whether to use TLS.
	UseTLS *LoggingUseTLS `default:"0" form:"name=use_tls"`
}

func (l LoggingLogentries) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingLogentries) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingLogentries) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingLogentries) GetFormatVersion() *LoggingLogentriesFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingLogentries) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingLogentries) GetPlacement() *LoggingLogentriesPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingLogentries) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *LoggingLogentries) GetRegion() *LoggingLogentriesRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingLogentries) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingLogentries) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingLogentries) GetUseTLS() *LoggingUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}
