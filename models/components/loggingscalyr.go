// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// LoggingScalyrFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingScalyrFormatVersion int64

const (
	LoggingScalyrFormatVersionOne LoggingScalyrFormatVersion = 1
	LoggingScalyrFormatVersionTwo LoggingScalyrFormatVersion = 2
)

func (e LoggingScalyrFormatVersion) ToPointer() *LoggingScalyrFormatVersion {
	return &e
}

func (e *LoggingScalyrFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingScalyrFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingScalyrFormatVersion: %v", v)
	}
}

// LoggingScalyrPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingScalyrPlacement string

const (
	LoggingScalyrPlacementNone                   LoggingScalyrPlacement = "none"
	LoggingScalyrPlacementWafDebug               LoggingScalyrPlacement = "waf_debug"
	LoggingScalyrPlacementLessThanNilGreaterThan LoggingScalyrPlacement = "<nil>"
)

func (e LoggingScalyrPlacement) ToPointer() *LoggingScalyrPlacement {
	return &e
}

func (e *LoggingScalyrPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingScalyrPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingScalyrPlacement: %v", v)
	}
}

// LoggingScalyrRegion - The region that log data will be sent to.
type LoggingScalyrRegion string

const (
	LoggingScalyrRegionUs LoggingScalyrRegion = "US"
	LoggingScalyrRegionEu LoggingScalyrRegion = "EU"
)

func (e LoggingScalyrRegion) ToPointer() *LoggingScalyrRegion {
	return &e
}

func (e *LoggingScalyrRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = LoggingScalyrRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingScalyrRegion: %v", v)
	}
}

type LoggingScalyr struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingScalyrFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingScalyrPlacement `form:"name=placement"`
	// The name of the logfile within Scalyr.
	ProjectID *string `default:"logplex" form:"name=project_id"`
	// The region that log data will be sent to.
	Region *LoggingScalyrRegion `default:"US" form:"name=region"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// The token to use for authentication ([https://www.scalyr.com/keys](https://www.scalyr.com/keys)).
	Token *string `form:"name=token"`
}

func (l LoggingScalyr) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingScalyr) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingScalyr) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingScalyr) GetFormatVersion() *LoggingScalyrFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingScalyr) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingScalyr) GetPlacement() *LoggingScalyrPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingScalyr) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *LoggingScalyr) GetRegion() *LoggingScalyrRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingScalyr) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingScalyr) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
