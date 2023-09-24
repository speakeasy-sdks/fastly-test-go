// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
)

// LoggingHerokuFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingHerokuFormatVersion int64

const (
	LoggingHerokuFormatVersionOne LoggingHerokuFormatVersion = 1
	LoggingHerokuFormatVersionTwo LoggingHerokuFormatVersion = 2
)

func (e LoggingHerokuFormatVersion) ToPointer() *LoggingHerokuFormatVersion {
	return &e
}

func (e *LoggingHerokuFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingHerokuFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHerokuFormatVersion: %v", v)
	}
}

// LoggingHerokuPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingHerokuPlacement string

const (
	LoggingHerokuPlacementNone                   LoggingHerokuPlacement = "none"
	LoggingHerokuPlacementWafDebug               LoggingHerokuPlacement = "waf_debug"
	LoggingHerokuPlacementLessThanNilGreaterThan LoggingHerokuPlacement = "<nil>"
)

func (e LoggingHerokuPlacement) ToPointer() *LoggingHerokuPlacement {
	return &e
}

func (e *LoggingHerokuPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingHerokuPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHerokuPlacement: %v", v)
	}
}

type LoggingHeroku2 struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingHerokuFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingHerokuPlacement `form:"name=placement"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// The token to use for authentication ([https://devcenter.heroku.com/articles/add-on-partner-log-integration](https://devcenter.heroku.com/articles/add-on-partner-log-integration)).
	Token *string `form:"name=token"`
	// The URL to stream logs to.
	URL *string `form:"name=url"`
}

func (l LoggingHeroku2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingHeroku2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingHeroku2) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingHeroku2) GetFormatVersion() *LoggingHerokuFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingHeroku2) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingHeroku2) GetPlacement() *LoggingHerokuPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingHeroku2) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingHeroku2) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingHeroku2) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
