// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// LoggingOpenstackResponseCompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type LoggingOpenstackResponseCompressionCodec string

const (
	LoggingOpenstackResponseCompressionCodecZstd   LoggingOpenstackResponseCompressionCodec = "zstd"
	LoggingOpenstackResponseCompressionCodecSnappy LoggingOpenstackResponseCompressionCodec = "snappy"
	LoggingOpenstackResponseCompressionCodecGzip   LoggingOpenstackResponseCompressionCodec = "gzip"
)

func (e LoggingOpenstackResponseCompressionCodec) ToPointer() *LoggingOpenstackResponseCompressionCodec {
	return &e
}

func (e *LoggingOpenstackResponseCompressionCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zstd":
		fallthrough
	case "snappy":
		fallthrough
	case "gzip":
		*e = LoggingOpenstackResponseCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingOpenstackResponseCompressionCodec: %v", v)
	}
}

// LoggingOpenstackResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingOpenstackResponseFormatVersion int64

const (
	LoggingOpenstackResponseFormatVersionOne LoggingOpenstackResponseFormatVersion = 1
	LoggingOpenstackResponseFormatVersionTwo LoggingOpenstackResponseFormatVersion = 2
)

func (e LoggingOpenstackResponseFormatVersion) ToPointer() *LoggingOpenstackResponseFormatVersion {
	return &e
}

func (e *LoggingOpenstackResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingOpenstackResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingOpenstackResponseFormatVersion: %v", v)
	}
}

// LoggingOpenstackResponseMessageType - How the message should be formatted.
type LoggingOpenstackResponseMessageType string

const (
	LoggingOpenstackResponseMessageTypeClassic LoggingOpenstackResponseMessageType = "classic"
	LoggingOpenstackResponseMessageTypeLoggly  LoggingOpenstackResponseMessageType = "loggly"
	LoggingOpenstackResponseMessageTypeLogplex LoggingOpenstackResponseMessageType = "logplex"
	LoggingOpenstackResponseMessageTypeBlank   LoggingOpenstackResponseMessageType = "blank"
)

func (e LoggingOpenstackResponseMessageType) ToPointer() *LoggingOpenstackResponseMessageType {
	return &e
}

func (e *LoggingOpenstackResponseMessageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "classic":
		fallthrough
	case "loggly":
		fallthrough
	case "logplex":
		fallthrough
	case "blank":
		*e = LoggingOpenstackResponseMessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingOpenstackResponseMessageType: %v", v)
	}
}

// LoggingOpenstackResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingOpenstackResponsePlacement string

const (
	LoggingOpenstackResponsePlacementNone     LoggingOpenstackResponsePlacement = "none"
	LoggingOpenstackResponsePlacementWafDebug LoggingOpenstackResponsePlacement = "waf_debug"
)

func (e LoggingOpenstackResponsePlacement) ToPointer() *LoggingOpenstackResponsePlacement {
	return &e
}

func (e *LoggingOpenstackResponsePlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		*e = LoggingOpenstackResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingOpenstackResponsePlacement: %v", v)
	}
}

type LoggingOpenstackResponse struct {
	// Your OpenStack account access key.
	AccessKey *string `json:"access_key,omitempty"`
	// The name of your OpenStack container.
	BucketName *string `json:"bucket_name,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *LoggingOpenstackResponseCompressionCodec `json:"compression_codec,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t \"%r\" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingOpenstackResponseFormatVersion `default:"2" json:"format_version"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `default:"0" json:"gzip_level"`
	// How the message should be formatted.
	MessageType *LoggingOpenstackResponseMessageType `default:"classic" json:"message_type"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// The path to upload logs to.
	Path *string `default:"null" json:"path"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `default:"3600" json:"period"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingOpenstackResponsePlacement `json:"placement,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `default:"null" json:"public_key"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// A timestamp format
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Your OpenStack auth url.
	URL *string `json:"url,omitempty"`
	// The username for your OpenStack account.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (l LoggingOpenstackResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingOpenstackResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingOpenstackResponse) GetAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AccessKey
}

func (o *LoggingOpenstackResponse) GetBucketName() *string {
	if o == nil {
		return nil
	}
	return o.BucketName
}

func (o *LoggingOpenstackResponse) GetCompressionCodec() *LoggingOpenstackResponseCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingOpenstackResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingOpenstackResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingOpenstackResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingOpenstackResponse) GetFormatVersion() *LoggingOpenstackResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingOpenstackResponse) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingOpenstackResponse) GetMessageType() *LoggingOpenstackResponseMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingOpenstackResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingOpenstackResponse) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingOpenstackResponse) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingOpenstackResponse) GetPlacement() *LoggingOpenstackResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingOpenstackResponse) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingOpenstackResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingOpenstackResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingOpenstackResponse) GetTimestampFormat() *string {
	if o == nil {
		return nil
	}
	return o.TimestampFormat
}

func (o *LoggingOpenstackResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingOpenstackResponse) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *LoggingOpenstackResponse) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *LoggingOpenstackResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
