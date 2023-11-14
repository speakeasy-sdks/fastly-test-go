// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"time"
)

// LoggingSplunkResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingSplunkResponseFormatVersion int64

const (
	LoggingSplunkResponseFormatVersionOne LoggingSplunkResponseFormatVersion = 1
	LoggingSplunkResponseFormatVersionTwo LoggingSplunkResponseFormatVersion = 2
)

func (e LoggingSplunkResponseFormatVersion) ToPointer() *LoggingSplunkResponseFormatVersion {
	return &e
}

func (e *LoggingSplunkResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingSplunkResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSplunkResponseFormatVersion: %v", v)
	}
}

// LoggingSplunkResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingSplunkResponsePlacement string

const (
	LoggingSplunkResponsePlacementNone                   LoggingSplunkResponsePlacement = "none"
	LoggingSplunkResponsePlacementWafDebug               LoggingSplunkResponsePlacement = "waf_debug"
	LoggingSplunkResponsePlacementLessThanNilGreaterThan LoggingSplunkResponsePlacement = "<nil>"
)

func (e LoggingSplunkResponsePlacement) ToPointer() *LoggingSplunkResponsePlacement {
	return &e
}

func (e *LoggingSplunkResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSplunkResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSplunkResponsePlacement: %v", v)
	}
}

type LoggingSplunkResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingSplunkResponseFormatVersion `default:"2" json:"format_version"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingSplunkResponsePlacement `json:"placement,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` for unbounded.
	RequestMaxBytes *int64 `default:"0" json:"request_max_bytes"`
	// The maximum number of logs sent in one request. Defaults `0` for unbounded.
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
	// A Splunk token for use in posting logs over HTTP to your collector.
	Token *string `json:"token,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The URL to post logs to.
	URL *string `json:"url,omitempty"`
	// Whether to use TLS.
	UseTLS  *LoggingUseTLS `default:"0" json:"use_tls"`
	Version *int64         `json:"version,omitempty"`
}

func (l LoggingSplunkResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingSplunkResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingSplunkResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingSplunkResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingSplunkResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingSplunkResponse) GetFormatVersion() *LoggingSplunkResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingSplunkResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingSplunkResponse) GetPlacement() *LoggingSplunkResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingSplunkResponse) GetRequestMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxBytes
}

func (o *LoggingSplunkResponse) GetRequestMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxEntries
}

func (o *LoggingSplunkResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingSplunkResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingSplunkResponse) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *LoggingSplunkResponse) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *LoggingSplunkResponse) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *LoggingSplunkResponse) GetTLSHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSHostname
}

func (o *LoggingSplunkResponse) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingSplunkResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingSplunkResponse) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *LoggingSplunkResponse) GetUseTLS() *LoggingUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}

func (o *LoggingSplunkResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
