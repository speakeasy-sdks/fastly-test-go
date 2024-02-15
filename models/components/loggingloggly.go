// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// LoggingLogglyFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingLogglyFormatVersion int64

const (
	LoggingLogglyFormatVersionOne LoggingLogglyFormatVersion = 1
	LoggingLogglyFormatVersionTwo LoggingLogglyFormatVersion = 2
)

func (e LoggingLogglyFormatVersion) ToPointer() *LoggingLogglyFormatVersion {
	return &e
}

func (e *LoggingLogglyFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingLogglyFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingLogglyFormatVersion: %v", v)
	}
}

// LoggingLogglyPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingLogglyPlacement string

const (
	LoggingLogglyPlacementNone     LoggingLogglyPlacement = "none"
	LoggingLogglyPlacementWafDebug LoggingLogglyPlacement = "waf_debug"
)

func (e LoggingLogglyPlacement) ToPointer() *LoggingLogglyPlacement {
	return &e
}

func (e *LoggingLogglyPlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		*e = LoggingLogglyPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingLogglyPlacement: %v", v)
	}
}

type LoggingLoggly struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t \"%r\" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingLogglyFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingLogglyPlacement `form:"name=placement"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// The token to use for authentication ([https://www.loggly.com/docs/customer-token-authentication-token/](https://www.loggly.com/docs/customer-token-authentication-token/)).
	Token *string `form:"name=token"`
}

func (l LoggingLoggly) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingLoggly) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingLoggly) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingLoggly) GetFormatVersion() *LoggingLogglyFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingLoggly) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingLoggly) GetPlacement() *LoggingLogglyPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingLoggly) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingLoggly) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
