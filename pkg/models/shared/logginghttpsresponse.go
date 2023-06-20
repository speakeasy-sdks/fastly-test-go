// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingHTTPSResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingHTTPSResponseFormatVersion int64

const (
	LoggingHTTPSResponseFormatVersionOne LoggingHTTPSResponseFormatVersion = 1
	LoggingHTTPSResponseFormatVersionTwo LoggingHTTPSResponseFormatVersion = 2
)

func (e LoggingHTTPSResponseFormatVersion) ToPointer() *LoggingHTTPSResponseFormatVersion {
	return &e
}

func (e *LoggingHTTPSResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingHTTPSResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHTTPSResponseFormatVersion: %v", v)
	}
}

// LoggingHTTPSResponseJSONFormat - Enforces valid JSON formatting for log entries.
type LoggingHTTPSResponseJSONFormat string

const (
	LoggingHTTPSResponseJSONFormatZero LoggingHTTPSResponseJSONFormat = "0"
	LoggingHTTPSResponseJSONFormatOne  LoggingHTTPSResponseJSONFormat = "1"
	LoggingHTTPSResponseJSONFormatTwo  LoggingHTTPSResponseJSONFormat = "2"
)

func (e LoggingHTTPSResponseJSONFormat) ToPointer() *LoggingHTTPSResponseJSONFormat {
	return &e
}

func (e *LoggingHTTPSResponseJSONFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "0":
		fallthrough
	case "1":
		fallthrough
	case "2":
		*e = LoggingHTTPSResponseJSONFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHTTPSResponseJSONFormat: %v", v)
	}
}

// LoggingHTTPSResponseMethod - HTTP method used for request.
type LoggingHTTPSResponseMethod string

const (
	LoggingHTTPSResponseMethodPost LoggingHTTPSResponseMethod = "POST"
	LoggingHTTPSResponseMethodPut  LoggingHTTPSResponseMethod = "PUT"
)

func (e LoggingHTTPSResponseMethod) ToPointer() *LoggingHTTPSResponseMethod {
	return &e
}

func (e *LoggingHTTPSResponseMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POST":
		fallthrough
	case "PUT":
		*e = LoggingHTTPSResponseMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHTTPSResponseMethod: %v", v)
	}
}

// LoggingHTTPSResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingHTTPSResponsePlacement string

const (
	LoggingHTTPSResponsePlacementNone                   LoggingHTTPSResponsePlacement = "none"
	LoggingHTTPSResponsePlacementWafDebug               LoggingHTTPSResponsePlacement = "waf_debug"
	LoggingHTTPSResponsePlacementLessThanNilGreaterThan LoggingHTTPSResponsePlacement = "<nil>"
)

func (e LoggingHTTPSResponsePlacement) ToPointer() *LoggingHTTPSResponsePlacement {
	return &e
}

func (e *LoggingHTTPSResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingHTTPSResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingHTTPSResponsePlacement: %v", v)
	}
}

// LoggingHTTPSResponse - OK
type LoggingHTTPSResponse struct {
	// Content type of the header sent with the request.
	ContentType *string `json:"content_type,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingHTTPSResponseFormatVersion `json:"format_version,omitempty"`
	// Name of the custom header sent with the request.
	HeaderName *string `json:"header_name,omitempty"`
	// Value of the custom header sent with the request.
	HeaderValue *string `json:"header_value,omitempty"`
	// Enforces valid JSON formatting for log entries.
	JSONFormat *LoggingHTTPSResponseJSONFormat `json:"json_format,omitempty"`
	// How the message should be formatted.
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// HTTP method used for request.
	Method *LoggingHTTPSResponseMethod `json:"method,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingHTTPSResponsePlacement `json:"placement,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (100MB).
	RequestMaxBytes *int64 `json:"request_max_bytes,omitempty"`
	// The maximum number of logs sent in one request. Defaults `0` (10k).
	RequestMaxEntries *int64 `json:"request_max_entries,omitempty"`
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
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The URL to send logs to. Must use HTTPS. Required.
	URL     *string `json:"url,omitempty"`
	Version *int64  `json:"version,omitempty"`
}
