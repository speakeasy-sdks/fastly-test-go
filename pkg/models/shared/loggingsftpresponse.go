// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

// LoggingSftpResponseCompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type LoggingSftpResponseCompressionCodec string

const (
	LoggingSftpResponseCompressionCodecZstd   LoggingSftpResponseCompressionCodec = "zstd"
	LoggingSftpResponseCompressionCodecSnappy LoggingSftpResponseCompressionCodec = "snappy"
	LoggingSftpResponseCompressionCodecGzip   LoggingSftpResponseCompressionCodec = "gzip"
)

func (e LoggingSftpResponseCompressionCodec) ToPointer() *LoggingSftpResponseCompressionCodec {
	return &e
}

func (e *LoggingSftpResponseCompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSftpResponseCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSftpResponseCompressionCodec: %v", v)
	}
}

// LoggingSftpResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingSftpResponseFormatVersion int64

const (
	LoggingSftpResponseFormatVersionOne LoggingSftpResponseFormatVersion = 1
	LoggingSftpResponseFormatVersionTwo LoggingSftpResponseFormatVersion = 2
)

func (e LoggingSftpResponseFormatVersion) ToPointer() *LoggingSftpResponseFormatVersion {
	return &e
}

func (e *LoggingSftpResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingSftpResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSftpResponseFormatVersion: %v", v)
	}
}

// LoggingSftpResponseMessageType - How the message should be formatted.
type LoggingSftpResponseMessageType string

const (
	LoggingSftpResponseMessageTypeClassic LoggingSftpResponseMessageType = "classic"
	LoggingSftpResponseMessageTypeLoggly  LoggingSftpResponseMessageType = "loggly"
	LoggingSftpResponseMessageTypeLogplex LoggingSftpResponseMessageType = "logplex"
	LoggingSftpResponseMessageTypeBlank   LoggingSftpResponseMessageType = "blank"
)

func (e LoggingSftpResponseMessageType) ToPointer() *LoggingSftpResponseMessageType {
	return &e
}

func (e *LoggingSftpResponseMessageType) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSftpResponseMessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSftpResponseMessageType: %v", v)
	}
}

// LoggingSftpResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingSftpResponsePlacement string

const (
	LoggingSftpResponsePlacementNone                   LoggingSftpResponsePlacement = "none"
	LoggingSftpResponsePlacementWafDebug               LoggingSftpResponsePlacement = "waf_debug"
	LoggingSftpResponsePlacementLessThanNilGreaterThan LoggingSftpResponsePlacement = "<nil>"
)

func (e LoggingSftpResponsePlacement) ToPointer() *LoggingSftpResponsePlacement {
	return &e
}

func (e *LoggingSftpResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSftpResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSftpResponsePlacement: %v", v)
	}
}

type LoggingSftpResponse struct {
	// A hostname or IPv4 address.
	Address *string `json:"address,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *LoggingSftpResponseCompressionCodec `json:"compression_codec,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingSftpResponseFormatVersion `default:"2" json:"format_version"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `default:"0" json:"gzip_level"`
	// How the message should be formatted.
	MessageType *LoggingSftpResponseMessageType `default:"classic" json:"message_type"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// The password for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference.
	Password *string `json:"password,omitempty"`
	// The path to upload logs to.
	Path *string `default:"null" json:"path"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `default:"3600" json:"period"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingSftpResponsePlacement `json:"placement,omitempty"`
	// The port number.
	Port *int64 `default:"22" json:"port"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `default:"null" json:"public_key"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	// The SSH private key for the server. If both `password` and `secret_key` are passed, `secret_key` will be used in preference.
	SecretKey *string `default:"null" json:"secret_key"`
	ServiceID *string `json:"service_id,omitempty"`
	// A list of host keys for all hosts we can connect to over SFTP.
	SSHKnownHosts *string `json:"ssh_known_hosts,omitempty"`
	// A timestamp format
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The username for the server.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (l LoggingSftpResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingSftpResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingSftpResponse) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *LoggingSftpResponse) GetCompressionCodec() *LoggingSftpResponseCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingSftpResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingSftpResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingSftpResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingSftpResponse) GetFormatVersion() *LoggingSftpResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingSftpResponse) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingSftpResponse) GetMessageType() *LoggingSftpResponseMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingSftpResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingSftpResponse) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *LoggingSftpResponse) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingSftpResponse) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingSftpResponse) GetPlacement() *LoggingSftpResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingSftpResponse) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *LoggingSftpResponse) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingSftpResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingSftpResponse) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingSftpResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingSftpResponse) GetSSHKnownHosts() *string {
	if o == nil {
		return nil
	}
	return o.SSHKnownHosts
}

func (o *LoggingSftpResponse) GetTimestampFormat() *string {
	if o == nil {
		return nil
	}
	return o.TimestampFormat
}

func (o *LoggingSftpResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingSftpResponse) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *LoggingSftpResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
