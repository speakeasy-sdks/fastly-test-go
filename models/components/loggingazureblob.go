// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// CompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type CompressionCodec string

const (
	CompressionCodecZstd   CompressionCodec = "zstd"
	CompressionCodecSnappy CompressionCodec = "snappy"
	CompressionCodecGzip   CompressionCodec = "gzip"
)

func (e CompressionCodec) ToPointer() *CompressionCodec {
	return &e
}

func (e *CompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = CompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CompressionCodec: %v", v)
	}
}

// FormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type FormatVersion int64

const (
	FormatVersionOne FormatVersion = 1
	FormatVersionTwo FormatVersion = 2
)

func (e FormatVersion) ToPointer() *FormatVersion {
	return &e
}

func (e *FormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = FormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FormatVersion: %v", v)
	}
}

// MessageType - How the message should be formatted.
type MessageType string

const (
	MessageTypeClassic MessageType = "classic"
	MessageTypeLoggly  MessageType = "loggly"
	MessageTypeLogplex MessageType = "logplex"
	MessageTypeBlank   MessageType = "blank"
)

func (e MessageType) ToPointer() *MessageType {
	return &e
}

func (e *MessageType) UnmarshalJSON(data []byte) error {
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
		*e = MessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MessageType: %v", v)
	}
}

// Placement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type Placement string

const (
	PlacementNone                   Placement = "none"
	PlacementWafDebug               Placement = "waf_debug"
	PlacementLessThanNilGreaterThan Placement = "<nil>"
)

func (e Placement) ToPointer() *Placement {
	return &e
}

func (e *Placement) UnmarshalJSON(data []byte) error {
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
		*e = Placement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Placement: %v", v)
	}
}

type LoggingAzureblob struct {
	// The unique Azure Blob Storage namespace in which your data objects are stored. Required.
	AccountName *string `form:"name=account_name"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *CompressionCodec `form:"name=compression_codec"`
	// The name of the Azure Blob Storage container in which to store logs. Required.
	Container *string `form:"name=container"`
	// The maximum number of bytes for each uploaded file. A value of 0 can be used to indicate there is no limit on the size of uploaded files, otherwise the minimum value is 1048576 bytes (1 MiB.)
	FileMaxBytes *int64 `form:"name=file_max_bytes"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *FormatVersion `default:"2" form:"name=format_version"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `default:"0" form:"name=gzip_level"`
	// How the message should be formatted.
	MessageType *MessageType `default:"classic" form:"name=message_type"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// The path to upload logs to.
	Path *string `default:"null" form:"name=path"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `default:"3600" form:"name=period"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *Placement `form:"name=placement"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `default:"null" form:"name=public_key"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// The Azure shared access signature providing write access to the blob service objects. Be sure to update your token before it expires or the logging functionality will not work. Required.
	SasToken *string `form:"name=sas_token"`
}

func (l LoggingAzureblob) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingAzureblob) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingAzureblob) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *LoggingAzureblob) GetCompressionCodec() *CompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingAzureblob) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *LoggingAzureblob) GetFileMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.FileMaxBytes
}

func (o *LoggingAzureblob) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingAzureblob) GetFormatVersion() *FormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingAzureblob) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingAzureblob) GetMessageType() *MessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingAzureblob) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingAzureblob) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingAzureblob) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingAzureblob) GetPlacement() *Placement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingAzureblob) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingAzureblob) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingAzureblob) GetSasToken() *string {
	if o == nil {
		return nil
	}
	return o.SasToken
}
