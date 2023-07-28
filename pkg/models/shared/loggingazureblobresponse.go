// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingAzureblobResponseCompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type LoggingAzureblobResponseCompressionCodec string

const (
	LoggingAzureblobResponseCompressionCodecZstd   LoggingAzureblobResponseCompressionCodec = "zstd"
	LoggingAzureblobResponseCompressionCodecSnappy LoggingAzureblobResponseCompressionCodec = "snappy"
	LoggingAzureblobResponseCompressionCodecGzip   LoggingAzureblobResponseCompressionCodec = "gzip"
)

func (e LoggingAzureblobResponseCompressionCodec) ToPointer() *LoggingAzureblobResponseCompressionCodec {
	return &e
}

func (e *LoggingAzureblobResponseCompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = LoggingAzureblobResponseCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingAzureblobResponseCompressionCodec: %v", v)
	}
}

// LoggingAzureblobResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingAzureblobResponseFormatVersion int64

const (
	LoggingAzureblobResponseFormatVersionOne LoggingAzureblobResponseFormatVersion = 1
	LoggingAzureblobResponseFormatVersionTwo LoggingAzureblobResponseFormatVersion = 2
)

func (e LoggingAzureblobResponseFormatVersion) ToPointer() *LoggingAzureblobResponseFormatVersion {
	return &e
}

func (e *LoggingAzureblobResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingAzureblobResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingAzureblobResponseFormatVersion: %v", v)
	}
}

// LoggingAzureblobResponseMessageType - How the message should be formatted.
type LoggingAzureblobResponseMessageType string

const (
	LoggingAzureblobResponseMessageTypeClassic LoggingAzureblobResponseMessageType = "classic"
	LoggingAzureblobResponseMessageTypeLoggly  LoggingAzureblobResponseMessageType = "loggly"
	LoggingAzureblobResponseMessageTypeLogplex LoggingAzureblobResponseMessageType = "logplex"
	LoggingAzureblobResponseMessageTypeBlank   LoggingAzureblobResponseMessageType = "blank"
)

func (e LoggingAzureblobResponseMessageType) ToPointer() *LoggingAzureblobResponseMessageType {
	return &e
}

func (e *LoggingAzureblobResponseMessageType) UnmarshalJSON(data []byte) error {
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
		*e = LoggingAzureblobResponseMessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingAzureblobResponseMessageType: %v", v)
	}
}

// LoggingAzureblobResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingAzureblobResponsePlacement string

const (
	LoggingAzureblobResponsePlacementNone                   LoggingAzureblobResponsePlacement = "none"
	LoggingAzureblobResponsePlacementWafDebug               LoggingAzureblobResponsePlacement = "waf_debug"
	LoggingAzureblobResponsePlacementLessThanNilGreaterThan LoggingAzureblobResponsePlacement = "<nil>"
)

func (e LoggingAzureblobResponsePlacement) ToPointer() *LoggingAzureblobResponsePlacement {
	return &e
}

func (e *LoggingAzureblobResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingAzureblobResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingAzureblobResponsePlacement: %v", v)
	}
}

// LoggingAzureblobResponse - OK
type LoggingAzureblobResponse struct {
	// The unique Azure Blob Storage namespace in which your data objects are stored. Required.
	AccountName *string `json:"account_name,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *LoggingAzureblobResponseCompressionCodec `json:"compression_codec,omitempty"`
	// The name of the Azure Blob Storage container in which to store logs. Required.
	Container *string `json:"container,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
	FileMaxBytes *int64 `json:"file_max_bytes,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingAzureblobResponseFormatVersion `json:"format_version,omitempty"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `json:"gzip_level,omitempty"`
	// How the message should be formatted.
	MessageType *LoggingAzureblobResponseMessageType `json:"message_type,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// The path to upload logs to.
	Path *string `json:"path,omitempty"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `json:"period,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingAzureblobResponsePlacement `json:"placement,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `json:"public_key,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	// The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required.
	SasToken  *string `json:"sas_token,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// A timestamp format
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (o *LoggingAzureblobResponse) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *LoggingAzureblobResponse) GetCompressionCodec() *LoggingAzureblobResponseCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingAzureblobResponse) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *LoggingAzureblobResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingAzureblobResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingAzureblobResponse) GetFileMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.FileMaxBytes
}

func (o *LoggingAzureblobResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingAzureblobResponse) GetFormatVersion() *LoggingAzureblobResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingAzureblobResponse) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingAzureblobResponse) GetMessageType() *LoggingAzureblobResponseMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingAzureblobResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingAzureblobResponse) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingAzureblobResponse) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingAzureblobResponse) GetPlacement() *LoggingAzureblobResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingAzureblobResponse) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingAzureblobResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingAzureblobResponse) GetSasToken() *string {
	if o == nil {
		return nil
	}
	return o.SasToken
}

func (o *LoggingAzureblobResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingAzureblobResponse) GetTimestampFormat() *string {
	if o == nil {
		return nil
	}
	return o.TimestampFormat
}

func (o *LoggingAzureblobResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingAzureblobResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
