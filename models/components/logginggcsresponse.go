// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// LoggingGcsResponseCompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type LoggingGcsResponseCompressionCodec string

const (
	LoggingGcsResponseCompressionCodecZstd   LoggingGcsResponseCompressionCodec = "zstd"
	LoggingGcsResponseCompressionCodecSnappy LoggingGcsResponseCompressionCodec = "snappy"
	LoggingGcsResponseCompressionCodecGzip   LoggingGcsResponseCompressionCodec = "gzip"
)

func (e LoggingGcsResponseCompressionCodec) ToPointer() *LoggingGcsResponseCompressionCodec {
	return &e
}

func (e *LoggingGcsResponseCompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = LoggingGcsResponseCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsResponseCompressionCodec: %v", v)
	}
}

// LoggingGcsResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingGcsResponseFormatVersion int64

const (
	LoggingGcsResponseFormatVersionOne LoggingGcsResponseFormatVersion = 1
	LoggingGcsResponseFormatVersionTwo LoggingGcsResponseFormatVersion = 2
)

func (e LoggingGcsResponseFormatVersion) ToPointer() *LoggingGcsResponseFormatVersion {
	return &e
}

func (e *LoggingGcsResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingGcsResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsResponseFormatVersion: %v", v)
	}
}

// LoggingGcsResponseMessageType - How the message should be formatted.
type LoggingGcsResponseMessageType string

const (
	LoggingGcsResponseMessageTypeClassic LoggingGcsResponseMessageType = "classic"
	LoggingGcsResponseMessageTypeLoggly  LoggingGcsResponseMessageType = "loggly"
	LoggingGcsResponseMessageTypeLogplex LoggingGcsResponseMessageType = "logplex"
	LoggingGcsResponseMessageTypeBlank   LoggingGcsResponseMessageType = "blank"
)

func (e LoggingGcsResponseMessageType) ToPointer() *LoggingGcsResponseMessageType {
	return &e
}

func (e *LoggingGcsResponseMessageType) UnmarshalJSON(data []byte) error {
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
		*e = LoggingGcsResponseMessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsResponseMessageType: %v", v)
	}
}

// LoggingGcsResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingGcsResponsePlacement string

const (
	LoggingGcsResponsePlacementNone     LoggingGcsResponsePlacement = "none"
	LoggingGcsResponsePlacementWafDebug LoggingGcsResponsePlacement = "waf_debug"
)

func (e LoggingGcsResponsePlacement) ToPointer() *LoggingGcsResponsePlacement {
	return &e
}

func (e *LoggingGcsResponsePlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		*e = LoggingGcsResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsResponsePlacement: %v", v)
	}
}

type LoggingGcsResponse struct {
	// The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided.
	AccountName *string `json:"account_name,omitempty"`
	// The name of the GCS bucket.
	BucketName *string `json:"bucket_name,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *LoggingGcsResponseCompressionCodec `json:"compression_codec,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t \"%r\" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingGcsResponseFormatVersion `default:"2" json:"format_version"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `default:"0" json:"gzip_level"`
	// How the message should be formatted.
	MessageType *LoggingGcsResponseMessageType `default:"classic" json:"message_type"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// The path to upload logs to.
	Path *string `default:"null" json:"path"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `default:"3600" json:"period"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingGcsResponsePlacement `json:"placement,omitempty"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `json:"project_id,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `default:"null" json:"public_key"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	// Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified.
	SecretKey *string `json:"secret_key,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	// A timestamp format
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (l LoggingGcsResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingGcsResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingGcsResponse) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *LoggingGcsResponse) GetBucketName() *string {
	if o == nil {
		return nil
	}
	return o.BucketName
}

func (o *LoggingGcsResponse) GetCompressionCodec() *LoggingGcsResponseCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingGcsResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingGcsResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingGcsResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingGcsResponse) GetFormatVersion() *LoggingGcsResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingGcsResponse) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingGcsResponse) GetMessageType() *LoggingGcsResponseMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingGcsResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingGcsResponse) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingGcsResponse) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingGcsResponse) GetPlacement() *LoggingGcsResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingGcsResponse) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *LoggingGcsResponse) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingGcsResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingGcsResponse) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingGcsResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingGcsResponse) GetTimestampFormat() *string {
	if o == nil {
		return nil
	}
	return o.TimestampFormat
}

func (o *LoggingGcsResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingGcsResponse) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *LoggingGcsResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
