// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingKafkaResponseAuthMethod - SASL authentication method.
type LoggingKafkaResponseAuthMethod string

const (
	LoggingKafkaResponseAuthMethodPlain       LoggingKafkaResponseAuthMethod = "plain"
	LoggingKafkaResponseAuthMethodScramSha256 LoggingKafkaResponseAuthMethod = "scram-sha-256"
	LoggingKafkaResponseAuthMethodScramSha512 LoggingKafkaResponseAuthMethod = "scram-sha-512"
)

func (e LoggingKafkaResponseAuthMethod) ToPointer() *LoggingKafkaResponseAuthMethod {
	return &e
}

func (e *LoggingKafkaResponseAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "plain":
		fallthrough
	case "scram-sha-256":
		fallthrough
	case "scram-sha-512":
		*e = LoggingKafkaResponseAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaResponseAuthMethod: %v", v)
	}
}

// LoggingKafkaResponseCompressionCodec - The codec used for compression of your logs.
type LoggingKafkaResponseCompressionCodec string

const (
	LoggingKafkaResponseCompressionCodecGzip                   LoggingKafkaResponseCompressionCodec = "gzip"
	LoggingKafkaResponseCompressionCodecSnappy                 LoggingKafkaResponseCompressionCodec = "snappy"
	LoggingKafkaResponseCompressionCodecLz4                    LoggingKafkaResponseCompressionCodec = "lz4"
	LoggingKafkaResponseCompressionCodecLessThanNilGreaterThan LoggingKafkaResponseCompressionCodec = "<nil>"
)

func (e LoggingKafkaResponseCompressionCodec) ToPointer() *LoggingKafkaResponseCompressionCodec {
	return &e
}

func (e *LoggingKafkaResponseCompressionCodec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gzip":
		fallthrough
	case "snappy":
		fallthrough
	case "lz4":
		fallthrough
	case "<nil>":
		*e = LoggingKafkaResponseCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaResponseCompressionCodec: %v", v)
	}
}

// LoggingKafkaResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingKafkaResponseFormatVersion int64

const (
	LoggingKafkaResponseFormatVersionOne LoggingKafkaResponseFormatVersion = 1
	LoggingKafkaResponseFormatVersionTwo LoggingKafkaResponseFormatVersion = 2
)

func (e LoggingKafkaResponseFormatVersion) ToPointer() *LoggingKafkaResponseFormatVersion {
	return &e
}

func (e *LoggingKafkaResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingKafkaResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaResponseFormatVersion: %v", v)
	}
}

// LoggingKafkaResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingKafkaResponsePlacement string

const (
	LoggingKafkaResponsePlacementNone                   LoggingKafkaResponsePlacement = "none"
	LoggingKafkaResponsePlacementWafDebug               LoggingKafkaResponsePlacement = "waf_debug"
	LoggingKafkaResponsePlacementLessThanNilGreaterThan LoggingKafkaResponsePlacement = "<nil>"
)

func (e LoggingKafkaResponsePlacement) ToPointer() *LoggingKafkaResponsePlacement {
	return &e
}

func (e *LoggingKafkaResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingKafkaResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaResponsePlacement: %v", v)
	}
}

// LoggingKafkaResponseRequiredAcks - The number of acknowledgements a leader must receive before a write is considered successful.
type LoggingKafkaResponseRequiredAcks int64

const (
	LoggingKafkaResponseRequiredAcksOne    LoggingKafkaResponseRequiredAcks = 1
	LoggingKafkaResponseRequiredAcksZero   LoggingKafkaResponseRequiredAcks = 0
	LoggingKafkaResponseRequiredAcksMinus1 LoggingKafkaResponseRequiredAcks = -1
)

func (e LoggingKafkaResponseRequiredAcks) ToPointer() *LoggingKafkaResponseRequiredAcks {
	return &e
}

func (e *LoggingKafkaResponseRequiredAcks) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 0:
		fallthrough
	case -1:
		*e = LoggingKafkaResponseRequiredAcks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaResponseRequiredAcks: %v", v)
	}
}

// LoggingKafkaResponse - OK
type LoggingKafkaResponse struct {
	// SASL authentication method.
	AuthMethod *LoggingKafkaResponseAuthMethod `json:"auth_method,omitempty"`
	// A comma-separated list of IP addresses or hostnames of Kafka brokers. Required.
	Brokers *string `json:"brokers,omitempty"`
	// The codec used for compression of your logs.
	CompressionCodec *LoggingKafkaResponseCompressionCodec `json:"compression_codec,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingKafkaResponseFormatVersion `json:"format_version,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Enables parsing of key=value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers).
	ParseLogKeyvals *bool `json:"parse_log_keyvals,omitempty"`
	// SASL password.
	Password *string `json:"password,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingKafkaResponsePlacement `json:"placement,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (no limit).
	RequestMaxBytes *int64 `json:"request_max_bytes,omitempty"`
	// The number of acknowledgements a leader must receive before a write is considered successful.
	RequiredAcks *LoggingKafkaResponseRequiredAcks `json:"required_acks,omitempty"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert *string `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert *string `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey *string `json:"tls_client_key,omitempty"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TLSHostname *string `json:"tls_hostname,omitempty"`
	// The Kafka topic to send logs to. Required.
	Topic *string `json:"topic,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Whether to use TLS.
	UseTLS *LoggingUseTLS `json:"use_tls,omitempty"`
	// SASL user.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (o *LoggingKafkaResponse) GetAuthMethod() *LoggingKafkaResponseAuthMethod {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *LoggingKafkaResponse) GetBrokers() *string {
	if o == nil {
		return nil
	}
	return o.Brokers
}

func (o *LoggingKafkaResponse) GetCompressionCodec() *LoggingKafkaResponseCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingKafkaResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingKafkaResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingKafkaResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingKafkaResponse) GetFormatVersion() *LoggingKafkaResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingKafkaResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingKafkaResponse) GetParseLogKeyvals() *bool {
	if o == nil {
		return nil
	}
	return o.ParseLogKeyvals
}

func (o *LoggingKafkaResponse) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *LoggingKafkaResponse) GetPlacement() *LoggingKafkaResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingKafkaResponse) GetRequestMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxBytes
}

func (o *LoggingKafkaResponse) GetRequiredAcks() *LoggingKafkaResponseRequiredAcks {
	if o == nil {
		return nil
	}
	return o.RequiredAcks
}

func (o *LoggingKafkaResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingKafkaResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingKafkaResponse) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *LoggingKafkaResponse) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *LoggingKafkaResponse) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *LoggingKafkaResponse) GetTLSHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSHostname
}

func (o *LoggingKafkaResponse) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}

func (o *LoggingKafkaResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingKafkaResponse) GetUseTLS() *LoggingUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}

func (o *LoggingKafkaResponse) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *LoggingKafkaResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
