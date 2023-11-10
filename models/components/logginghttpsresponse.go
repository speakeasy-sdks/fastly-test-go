// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
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

type LoggingHTTPSResponse struct {
	// Content type of the header sent with the request.
	ContentType *string `default:"null" json:"content_type"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingHTTPSResponseFormatVersion `default:"2" json:"format_version"`
	// Name of the custom header sent with the request.
	HeaderName *string `default:"null" json:"header_name"`
	// Value of the custom header sent with the request.
	HeaderValue *string `default:"null" json:"header_value"`
	// Enforces valid JSON formatting for log entries.
	JSONFormat *LoggingHTTPSResponseJSONFormat `json:"json_format,omitempty"`
	// How the message should be formatted.
	MessageType *LoggingMessageType `default:"classic" json:"message_type"`
	// HTTP method used for request.
	Method *LoggingHTTPSResponseMethod `default:"POST" json:"method"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingHTTPSResponsePlacement `json:"placement,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (100MB).
	RequestMaxBytes *int64 `default:"0" json:"request_max_bytes"`
	// The maximum number of logs sent in one request. Defaults `0` (10k).
	RequestMaxEntries *int64 `default:"0" json:"request_max_entries"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert *string `default:"null" json:"tls_ca_cert"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert *string `default:"null" json:"tls_client_cert"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey *string `default:"null" json:"tls_client_key"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TLSHostname *string `default:"null" json:"tls_hostname"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The URL to send logs to. Must use HTTPS. Required.
	URL     *string `json:"url,omitempty"`
	Version *int64  `json:"version,omitempty"`
}

func (l LoggingHTTPSResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingHTTPSResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingHTTPSResponse) GetContentType() *string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *LoggingHTTPSResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingHTTPSResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingHTTPSResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingHTTPSResponse) GetFormatVersion() *LoggingHTTPSResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingHTTPSResponse) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *LoggingHTTPSResponse) GetHeaderValue() *string {
	if o == nil {
		return nil
	}
	return o.HeaderValue
}

func (o *LoggingHTTPSResponse) GetJSONFormat() *LoggingHTTPSResponseJSONFormat {
	if o == nil {
		return nil
	}
	return o.JSONFormat
}

func (o *LoggingHTTPSResponse) GetMessageType() *LoggingMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingHTTPSResponse) GetMethod() *LoggingHTTPSResponseMethod {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *LoggingHTTPSResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingHTTPSResponse) GetPlacement() *LoggingHTTPSResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingHTTPSResponse) GetRequestMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxBytes
}

func (o *LoggingHTTPSResponse) GetRequestMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxEntries
}

func (o *LoggingHTTPSResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingHTTPSResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingHTTPSResponse) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *LoggingHTTPSResponse) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *LoggingHTTPSResponse) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *LoggingHTTPSResponse) GetTLSHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSHostname
}

func (o *LoggingHTTPSResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingHTTPSResponse) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *LoggingHTTPSResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
