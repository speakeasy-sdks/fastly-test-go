// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// LoggingGcsCompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type LoggingGcsCompressionCodec string

const (
	LoggingGcsCompressionCodecZstd   LoggingGcsCompressionCodec = "zstd"
	LoggingGcsCompressionCodecSnappy LoggingGcsCompressionCodec = "snappy"
	LoggingGcsCompressionCodecGzip   LoggingGcsCompressionCodec = "gzip"
)

func (e LoggingGcsCompressionCodec) ToPointer() *LoggingGcsCompressionCodec {
	return &e
}

func (e *LoggingGcsCompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = LoggingGcsCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsCompressionCodec: %v", v)
	}
}

// LoggingGcsFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingGcsFormatVersion int64

const (
	LoggingGcsFormatVersionOne LoggingGcsFormatVersion = 1
	LoggingGcsFormatVersionTwo LoggingGcsFormatVersion = 2
)

func (e LoggingGcsFormatVersion) ToPointer() *LoggingGcsFormatVersion {
	return &e
}

func (e *LoggingGcsFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingGcsFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsFormatVersion: %v", v)
	}
}

// LoggingGcsMessageType - How the message should be formatted.
type LoggingGcsMessageType string

const (
	LoggingGcsMessageTypeClassic LoggingGcsMessageType = "classic"
	LoggingGcsMessageTypeLoggly  LoggingGcsMessageType = "loggly"
	LoggingGcsMessageTypeLogplex LoggingGcsMessageType = "logplex"
	LoggingGcsMessageTypeBlank   LoggingGcsMessageType = "blank"
)

func (e LoggingGcsMessageType) ToPointer() *LoggingGcsMessageType {
	return &e
}

func (e *LoggingGcsMessageType) UnmarshalJSON(data []byte) error {
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
		*e = LoggingGcsMessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsMessageType: %v", v)
	}
}

// LoggingGcsPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingGcsPlacement string

const (
	LoggingGcsPlacementNone                   LoggingGcsPlacement = "none"
	LoggingGcsPlacementWafDebug               LoggingGcsPlacement = "waf_debug"
	LoggingGcsPlacementLessThanNilGreaterThan LoggingGcsPlacement = "<nil>"
)

func (e LoggingGcsPlacement) ToPointer() *LoggingGcsPlacement {
	return &e
}

func (e *LoggingGcsPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingGcsPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGcsPlacement: %v", v)
	}
}

type LoggingGcs struct {
	// The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided.
	AccountName *string `form:"name=account_name"`
	// The name of the GCS bucket.
	BucketName *string `form:"name=bucket_name"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *LoggingGcsCompressionCodec `form:"name=compression_codec"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingGcsFormatVersion `default:"2" form:"name=format_version"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `default:"0" form:"name=gzip_level"`
	// How the message should be formatted.
	MessageType *LoggingGcsMessageType `default:"classic" form:"name=message_type"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// The path to upload logs to.
	Path *string `default:"null" form:"name=path"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `default:"3600" form:"name=period"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingGcsPlacement `form:"name=placement"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `form:"name=project_id"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `default:"null" form:"name=public_key"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified.
	SecretKey *string `form:"name=secret_key"`
	// Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified.
	User *string `form:"name=user"`
}

func (l LoggingGcs) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingGcs) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingGcs) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *LoggingGcs) GetBucketName() *string {
	if o == nil {
		return nil
	}
	return o.BucketName
}

func (o *LoggingGcs) GetCompressionCodec() *LoggingGcsCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingGcs) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingGcs) GetFormatVersion() *LoggingGcsFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingGcs) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingGcs) GetMessageType() *LoggingGcsMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingGcs) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingGcs) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingGcs) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingGcs) GetPlacement() *LoggingGcsPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingGcs) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *LoggingGcs) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingGcs) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingGcs) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingGcs) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
