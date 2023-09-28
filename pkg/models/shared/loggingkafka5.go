// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
)

// LoggingKafkaAuthMethod - SASL authentication method.
type LoggingKafkaAuthMethod string

const (
	LoggingKafkaAuthMethodPlain       LoggingKafkaAuthMethod = "plain"
	LoggingKafkaAuthMethodScramSha256 LoggingKafkaAuthMethod = "scram-sha-256"
	LoggingKafkaAuthMethodScramSha512 LoggingKafkaAuthMethod = "scram-sha-512"
)

func (e LoggingKafkaAuthMethod) ToPointer() *LoggingKafkaAuthMethod {
	return &e
}

func (e *LoggingKafkaAuthMethod) UnmarshalJSON(data []byte) error {
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
		*e = LoggingKafkaAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaAuthMethod: %v", v)
	}
}

// LoggingKafkaCompressionCodec - The codec used for compression of your logs.
type LoggingKafkaCompressionCodec string

const (
	LoggingKafkaCompressionCodecGzip                   LoggingKafkaCompressionCodec = "gzip"
	LoggingKafkaCompressionCodecSnappy                 LoggingKafkaCompressionCodec = "snappy"
	LoggingKafkaCompressionCodecLz4                    LoggingKafkaCompressionCodec = "lz4"
	LoggingKafkaCompressionCodecLessThanNilGreaterThan LoggingKafkaCompressionCodec = "<nil>"
)

func (e LoggingKafkaCompressionCodec) ToPointer() *LoggingKafkaCompressionCodec {
	return &e
}

func (e *LoggingKafkaCompressionCodec) UnmarshalJSON(data []byte) error {
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
		*e = LoggingKafkaCompressionCodec(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaCompressionCodec: %v", v)
	}
}

// LoggingKafkaFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingKafkaFormatVersion int64

const (
	LoggingKafkaFormatVersionOne LoggingKafkaFormatVersion = 1
	LoggingKafkaFormatVersionTwo LoggingKafkaFormatVersion = 2
)

func (e LoggingKafkaFormatVersion) ToPointer() *LoggingKafkaFormatVersion {
	return &e
}

func (e *LoggingKafkaFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingKafkaFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaFormatVersion: %v", v)
	}
}

// LoggingKafkaPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingKafkaPlacement string

const (
	LoggingKafkaPlacementNone                   LoggingKafkaPlacement = "none"
	LoggingKafkaPlacementWafDebug               LoggingKafkaPlacement = "waf_debug"
	LoggingKafkaPlacementLessThanNilGreaterThan LoggingKafkaPlacement = "<nil>"
)

func (e LoggingKafkaPlacement) ToPointer() *LoggingKafkaPlacement {
	return &e
}

func (e *LoggingKafkaPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingKafkaPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaPlacement: %v", v)
	}
}

// LoggingKafkaRequiredAcks - The number of acknowledgements a leader must receive before a write is considered successful.
type LoggingKafkaRequiredAcks int64

const (
	LoggingKafkaRequiredAcksOne    LoggingKafkaRequiredAcks = 1
	LoggingKafkaRequiredAcksZero   LoggingKafkaRequiredAcks = 0
	LoggingKafkaRequiredAcksMinus1 LoggingKafkaRequiredAcks = -1
)

func (e LoggingKafkaRequiredAcks) ToPointer() *LoggingKafkaRequiredAcks {
	return &e
}

func (e *LoggingKafkaRequiredAcks) UnmarshalJSON(data []byte) error {
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
		*e = LoggingKafkaRequiredAcks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingKafkaRequiredAcks: %v", v)
	}
}

type LoggingKafka5 struct {
	// SASL authentication method.
	AuthMethod *LoggingKafkaAuthMethod `form:"name=auth_method"`
	// A comma-separated list of IP addresses or hostnames of Kafka brokers. Required.
	Brokers *string `form:"name=brokers"`
	// The codec used for compression of your logs.
	CompressionCodec *LoggingKafkaCompressionCodec `form:"name=compression_codec"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingKafkaFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Enables parsing of key=value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers).
	ParseLogKeyvals *bool `form:"name=parse_log_keyvals"`
	// SASL password.
	Password *string `form:"name=password"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingKafkaPlacement `form:"name=placement"`
	// The maximum number of bytes sent in one request. Defaults `0` (no limit).
	RequestMaxBytes *int64 `default:"0" form:"name=request_max_bytes"`
	// The number of acknowledgements a leader must receive before a write is considered successful.
	RequiredAcks *LoggingKafkaRequiredAcks `default:"1" form:"name=required_acks"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert *string `default:"null" form:"name=tls_ca_cert"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert *string `default:"null" form:"name=tls_client_cert"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey *string `default:"null" form:"name=tls_client_key"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TLSHostname *string `default:"null" form:"name=tls_hostname"`
	// The Kafka topic to send logs to. Required.
	Topic *string `form:"name=topic"`
	// Whether to use TLS.
	UseTLS *LoggingUseTLS `default:"0" form:"name=use_tls"`
	// SASL user.
	User *string `form:"name=user"`
}

func (l LoggingKafka5) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingKafka5) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingKafka5) GetAuthMethod() *LoggingKafkaAuthMethod {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *LoggingKafka5) GetBrokers() *string {
	if o == nil {
		return nil
	}
	return o.Brokers
}

func (o *LoggingKafka5) GetCompressionCodec() *LoggingKafkaCompressionCodec {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *LoggingKafka5) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingKafka5) GetFormatVersion() *LoggingKafkaFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingKafka5) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingKafka5) GetParseLogKeyvals() *bool {
	if o == nil {
		return nil
	}
	return o.ParseLogKeyvals
}

func (o *LoggingKafka5) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *LoggingKafka5) GetPlacement() *LoggingKafkaPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingKafka5) GetRequestMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxBytes
}

func (o *LoggingKafka5) GetRequiredAcks() *LoggingKafkaRequiredAcks {
	if o == nil {
		return nil
	}
	return o.RequiredAcks
}

func (o *LoggingKafka5) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingKafka5) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *LoggingKafka5) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *LoggingKafka5) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *LoggingKafka5) GetTLSHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSHostname
}

func (o *LoggingKafka5) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}

func (o *LoggingKafka5) GetUseTLS() *LoggingUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}

func (o *LoggingKafka5) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
