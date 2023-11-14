// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// LoggingFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingFormatVersion int64

const (
	LoggingFormatVersionOne LoggingFormatVersion = 1
	LoggingFormatVersionTwo LoggingFormatVersion = 2
)

func (e LoggingFormatVersion) ToPointer() *LoggingFormatVersion {
	return &e
}

func (e *LoggingFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingFormatVersion: %v", v)
	}
}