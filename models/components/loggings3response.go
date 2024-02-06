// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// LoggingS3ResponseCompressionCodec - The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
type LoggingS3ResponseCompressionCodec string

const (
	LoggingS3ResponseCompressionCodecZstd   LoggingS3ResponseCompressionCodec = "zstd"
	LoggingS3ResponseCompressionCodecSnappy LoggingS3ResponseCompressionCodec = "snappy"
	LoggingS3ResponseCompressionCodecGzip   LoggingS3ResponseCompressionCodec = "gzip"
)

func (e LoggingS3ResponseCompressionCodec) ToPointer() *LoggingS3ResponseCompressionCodec {
	return &e
}

func (e *LoggingS3ResponseCompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = LoggingS3ResponseCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingS3ResponseCompressionCodec: %v", v)
	}
}

// LoggingS3ResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingS3ResponseFormatVersion int64

const (
	LoggingS3ResponseFormatVersionOne LoggingS3ResponseFormatVersion = 1
	LoggingS3ResponseFormatVersionTwo LoggingS3ResponseFormatVersion = 2
)

func (e LoggingS3ResponseFormatVersion) ToPointer() *LoggingS3ResponseFormatVersion {
	return &e
}

func (e *LoggingS3ResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingS3ResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingS3ResponseFormatVersion: %v", v)
	}
}

// LoggingS3ResponseMessageType - How the message should be formatted.
type LoggingS3ResponseMessageType string

const (
	LoggingS3ResponseMessageTypeClassic LoggingS3ResponseMessageType = "classic"
	LoggingS3ResponseMessageTypeLoggly  LoggingS3ResponseMessageType = "loggly"
	LoggingS3ResponseMessageTypeLogplex LoggingS3ResponseMessageType = "logplex"
	LoggingS3ResponseMessageTypeBlank   LoggingS3ResponseMessageType = "blank"
)

func (e LoggingS3ResponseMessageType) ToPointer() *LoggingS3ResponseMessageType {
	return &e
}

func (e *LoggingS3ResponseMessageType) UnmarshalJSON(data []byte) error {
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
		*e = LoggingS3ResponseMessageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingS3ResponseMessageType: %v", v)
	}
}

// LoggingS3ResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingS3ResponsePlacement string

const (
	LoggingS3ResponsePlacementNone                   LoggingS3ResponsePlacement = "none"
	LoggingS3ResponsePlacementWafDebug               LoggingS3ResponsePlacement = "waf_debug"
	LoggingS3ResponsePlacementLessThanNilGreaterThan LoggingS3ResponsePlacement = "<nil>"
)

func (e LoggingS3ResponsePlacement) ToPointer() *LoggingS3ResponsePlacement {
	return &e
}

func (e *LoggingS3ResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingS3ResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingS3ResponsePlacement: %v", v)
	}
}

type LoggingS3Response struct {
	// The access key for your S3 account. Not required if `iam_role` is provided.
	AccessKey *string `json:"access_key,omitempty"`
	// The access control list (ACL) specific request header. See the AWS documentation for [Access Control List (ACL) Specific Request Headers](https://docs.aws.amazon.com/AmazonS3/latest/API/mpUploadInitiate.html#initiate-mpu-acl-specific-request-headers) for more information.
	ACL *string `json:"acl,omitempty"`
	// The bucket name for S3 account.
	BucketName *string `json:"bucket_name,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *LoggingS3ResponseCompressionCodec `json:"compression_codec,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The domain of the Amazon S3 endpoint.
	Domain *string `json:"domain,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t \"%r\" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingS3ResponseFormatVersion `default:"2" json:"format_version"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel *int64 `default:"0" json:"gzip_level"`
	// The Amazon Resource Name (ARN) for the IAM role granting Fastly access to S3. Not required if `access_key` and `secret_key` are provided.
	IamRole *string `json:"iam_role,omitempty"`
	// How the message should be formatted.
	MessageType *LoggingS3ResponseMessageType `default:"classic" json:"message_type"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// The path to upload logs to.
	Path *string `default:"null" json:"path"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *int64 `default:"3600" json:"period"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingS3ResponsePlacement `json:"placement,omitempty"`
	// A PGP public key that Fastly will use to encrypt your log files before writing them to disk.
	PublicKey *string `default:"null" json:"public_key"`
	// The S3 redundancy level.
	Redundancy *string `default:"null" json:"redundancy"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	// The secret key for your S3 account. Not required if `iam_role` is provided.
	SecretKey *string `json:"secret_key,omitempty"`
	// Set this to `AES256` or `aws:kms` to enable S3 Server Side Encryption.
	ServerSideEncryption *string `default:"null" json:"server_side_encryption"`
	// Optional server-side KMS Key Id. Must be set if `server_side_encryption` is set to `aws:kms` or `AES256`.
	ServerSideEncryptionKmsKeyID *string `default:"null" json:"server_side_encryption_kms_key_id"`
	ServiceID                    *string `json:"service_id,omitempty"`
	// A timestamp format
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (l LoggingS3Response) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingS3Response) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingS3Response) GetAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AccessKey
}

func (o *LoggingS3Response) GetACL() *string {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *LoggingS3Response) GetBucketName() *string {
	if o == nil {
		return nil
	}
	return o.BucketName
}

func (o *LoggingS3Response) GetCompressionCodec() *LoggingS3ResponseCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingS3Response) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingS3Response) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingS3Response) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *LoggingS3Response) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingS3Response) GetFormatVersion() *LoggingS3ResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingS3Response) GetGzipLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.GzipLevel
}

func (o *LoggingS3Response) GetIamRole() *string {
	if o == nil {
		return nil
	}
	return o.IamRole
}

func (o *LoggingS3Response) GetMessageType() *LoggingS3ResponseMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingS3Response) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingS3Response) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *LoggingS3Response) GetPeriod() *int64 {
	if o == nil {
		return nil
	}
	return o.Period
}

func (o *LoggingS3Response) GetPlacement() *LoggingS3ResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingS3Response) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *LoggingS3Response) GetRedundancy() *string {
	if o == nil {
		return nil
	}
	return o.Redundancy
}

func (o *LoggingS3Response) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingS3Response) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingS3Response) GetServerSideEncryption() *string {
	if o == nil {
		return nil
	}
	return o.ServerSideEncryption
}

func (o *LoggingS3Response) GetServerSideEncryptionKmsKeyID() *string {
	if o == nil {
		return nil
	}
	return o.ServerSideEncryptionKmsKeyID
}

func (o *LoggingS3Response) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingS3Response) GetTimestampFormat() *string {
	if o == nil {
		return nil
	}
	return o.TimestampFormat
}

func (o *LoggingS3Response) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingS3Response) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
