// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// LoggingCloudfilesResponseCompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type LoggingCloudfilesResponseCompressionCodec string

const (
	LoggingCloudfilesResponseCompressionCodecZstd   LoggingCloudfilesResponseCompressionCodec = "zstd"
	LoggingCloudfilesResponseCompressionCodecSnappy LoggingCloudfilesResponseCompressionCodec = "snappy"
	LoggingCloudfilesResponseCompressionCodecGzip   LoggingCloudfilesResponseCompressionCodec = "gzip"
)

func (e LoggingCloudfilesResponseCompressionCodec) ToPointer() *LoggingCloudfilesResponseCompressionCodec {
	return &e
}

func (e *LoggingCloudfilesResponseCompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = LoggingCloudfilesResponseCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingCloudfilesResponseCompressionCodec: %v", v)
	}
}

// LoggingCloudfilesResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingCloudfilesResponseFormatVersion int64

const (
	LoggingCloudfilesResponseFormatVersionOne LoggingCloudfilesResponseFormatVersion = 1
	LoggingCloudfilesResponseFormatVersionTwo LoggingCloudfilesResponseFormatVersion = 2
)

func (e LoggingCloudfilesResponseFormatVersion) ToPointer() *LoggingCloudfilesResponseFormatVersion {
	return &e
}

func (e *LoggingCloudfilesResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingCloudfilesResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingCloudfilesResponseFormatVersion: %v", v)
	}
}

// LoggingCloudfilesResponseMessageType - How the message should be formatted.
type LoggingCloudfilesResponseMessageType string

const (
	LoggingCloudfilesResponseMessageTypeClassic LoggingCloudfilesResponseMessageType = "classic"
	LoggingCloudfilesResponseMessageTypeLoggly  LoggingCloudfilesResponseMessageType = "loggly"
	LoggingCloudfilesResponseMessageTypeLogplex LoggingCloudfilesResponseMessageType = "logplex"
	LoggingCloudfilesResponseMessageTypeBlank   LoggingCloudfilesResponseMessageType = "blank"
)

func (e LoggingCloudfilesResponseMessageType) ToPointer() *LoggingCloudfilesResponseMessageType {
	return &e
}

func (e *LoggingCloudfilesResponseMessageType) UnmarshalJSON(data []byte) error {
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
		*e = LoggingCloudfilesResponseMessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingCloudfilesResponseMessageType: %v", v)
	}
}

// LoggingCloudfilesResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingCloudfilesResponsePlacement string

const (
	LoggingCloudfilesResponsePlacementNone     LoggingCloudfilesResponsePlacement = "none"
	LoggingCloudfilesResponsePlacementWafDebug LoggingCloudfilesResponsePlacement = "waf_debug"
)

func (e LoggingCloudfilesResponsePlacement) ToPointer() *LoggingCloudfilesResponsePlacement {
	return &e
}

func (e *LoggingCloudfilesResponsePlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		*e = LoggingCloudfilesResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingCloudfilesResponsePlacement: %v", v)
	}
}

// LoggingCloudfilesResponseRegion - The region to stream logs to.
type LoggingCloudfilesResponseRegion string

const (
	LoggingCloudfilesResponseRegionDfw LoggingCloudfilesResponseRegion = "DFW"
	LoggingCloudfilesResponseRegionOrd LoggingCloudfilesResponseRegion = "ORD"
	LoggingCloudfilesResponseRegionIad LoggingCloudfilesResponseRegion = "IAD"
	LoggingCloudfilesResponseRegionLon LoggingCloudfilesResponseRegion = "LON"
	LoggingCloudfilesResponseRegionSyd LoggingCloudfilesResponseRegion = "SYD"
	LoggingCloudfilesResponseRegionHkg LoggingCloudfilesResponseRegion = "HKG"
)

func (e LoggingCloudfilesResponseRegion) ToPointer() *LoggingCloudfilesResponseRegion {
	return &e
}

func (e *LoggingCloudfilesResponseRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DFW":
		fallthrough
	case "ORD":
		fallthrough
	case "IAD":
		fallthrough
	case "LON":
		fallthrough
	case "SYD":
		fallthrough
	case "HKG":
		*e = LoggingCloudfilesResponseRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingCloudfilesResponseRegion: %v", v)
	}
}

type LoggingCloudfilesResponse struct {
	// Your Cloud Files account access key.
	AccessKey *string `json:"access_key,omitempty"`
	// The name of your Cloud Files container.
	BucketName *string `json:"bucket_name,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *LoggingCloudfilesResponseCompressionCodec `json:"compression_codec,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t \"%r\" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingCloudfilesResponseFormatVersion `default:"2" json:"format_version"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `default:"0" json:"gzip_level"`
	// How the message should be formatted.
	MessageType *LoggingCloudfilesResponseMessageType `default:"classic" json:"message_type"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// The path to upload logs to.
	Path *string `default:"null" json:"path"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `default:"3600" json:"period"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingCloudfilesResponsePlacement `json:"placement,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `default:"null" json:"public_key"`
	// The region to stream logs to.
	Region *LoggingCloudfilesResponseRegion `json:"region,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// A timestamp format
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The username for your Cloud Files account.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (l LoggingCloudfilesResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingCloudfilesResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingCloudfilesResponse) GetAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AccessKey
}

func (o *LoggingCloudfilesResponse) GetBucketName() *string {
	if o == nil {
		return nil
	}
	return o.BucketName
}

func (o *LoggingCloudfilesResponse) GetCompressionCodec() *LoggingCloudfilesResponseCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingCloudfilesResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingCloudfilesResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingCloudfilesResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingCloudfilesResponse) GetFormatVersion() *LoggingCloudfilesResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingCloudfilesResponse) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingCloudfilesResponse) GetMessageType() *LoggingCloudfilesResponseMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingCloudfilesResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingCloudfilesResponse) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingCloudfilesResponse) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingCloudfilesResponse) GetPlacement() *LoggingCloudfilesResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingCloudfilesResponse) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingCloudfilesResponse) GetRegion() *LoggingCloudfilesResponseRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingCloudfilesResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingCloudfilesResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingCloudfilesResponse) GetTimestampFormat() *string {
	if o == nil {
		return nil
	}
	return o.TimestampFormat
}

func (o *LoggingCloudfilesResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingCloudfilesResponse) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *LoggingCloudfilesResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
