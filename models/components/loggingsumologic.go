// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// LoggingSumologicFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingSumologicFormatVersion int64

const (
	LoggingSumologicFormatVersionOne LoggingSumologicFormatVersion = 1
	LoggingSumologicFormatVersionTwo LoggingSumologicFormatVersion = 2
)

func (e LoggingSumologicFormatVersion) ToPointer() *LoggingSumologicFormatVersion {
	return &e
}

func (e *LoggingSumologicFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingSumologicFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSumologicFormatVersion: %v", v)
	}
}

// LoggingSumologicPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingSumologicPlacement string

const (
	LoggingSumologicPlacementNone                   LoggingSumologicPlacement = "none"
	LoggingSumologicPlacementWafDebug               LoggingSumologicPlacement = "waf_debug"
	LoggingSumologicPlacementLessThanNilGreaterThan LoggingSumologicPlacement = "<nil>"
)

func (e LoggingSumologicPlacement) ToPointer() *LoggingSumologicPlacement {
	return &e
}

func (e *LoggingSumologicPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSumologicPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSumologicPlacement: %v", v)
	}
}

type LoggingSumologic struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t \"%r\" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingSumologicFormatVersion `default:"2" form:"name=format_version"`
	// How the message should be formatted.
	MessageType *LoggingMessageType `default:"classic" form:"name=message_type"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingSumologicPlacement `form:"name=placement"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// The URL to post logs to.
	URL *string `form:"name=url"`
}

func (l LoggingSumologic) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingSumologic) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingSumologic) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingSumologic) GetFormatVersion() *LoggingSumologicFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingSumologic) GetMessageType() *LoggingMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingSumologic) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingSumologic) GetPlacement() *LoggingSumologicPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingSumologic) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingSumologic) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}
