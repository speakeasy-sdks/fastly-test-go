// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
	"time"
)

// LoggingSyslogResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingSyslogResponseFormatVersion int64

const (
	LoggingSyslogResponseFormatVersionOne LoggingSyslogResponseFormatVersion = 1
	LoggingSyslogResponseFormatVersionTwo LoggingSyslogResponseFormatVersion = 2
)

func (e LoggingSyslogResponseFormatVersion) ToPointer() *LoggingSyslogResponseFormatVersion {
	return &e
}

func (e *LoggingSyslogResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingSyslogResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSyslogResponseFormatVersion: %v", v)
	}
}

// LoggingSyslogResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingSyslogResponsePlacement string

const (
	LoggingSyslogResponsePlacementNone                   LoggingSyslogResponsePlacement = "none"
	LoggingSyslogResponsePlacementWafDebug               LoggingSyslogResponsePlacement = "waf_debug"
	LoggingSyslogResponsePlacementLessThanNilGreaterThan LoggingSyslogResponsePlacement = "<nil>"
)

func (e LoggingSyslogResponsePlacement) ToPointer() *LoggingSyslogResponsePlacement {
	return &e
}

func (e *LoggingSyslogResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSyslogResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSyslogResponsePlacement: %v", v)
	}
}

type LoggingSyslogResponse struct {
	// A hostname or IPv4 address.
	Address *string `json:"address,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingSyslogResponseFormatVersion `default:"2" json:"format_version"`
	// The hostname used for the syslog endpoint.
	Hostname *string `json:"hostname,omitempty"`
	// The IPv4 address used for the syslog endpoint.
	Ipv4 *string `json:"ipv4,omitempty"`
	// How the message should be formatted.
	MessageType *LoggingMessageType `default:"classic" json:"message_type"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingSyslogResponsePlacement `json:"placement,omitempty"`
	// The port number.
	Port *int64 `default:"514" json:"port"`
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
	// Whether to prepend each message with a specific token.
	Token *string `default:"null" json:"token"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Whether to use TLS.
	UseTLS  *LoggingUseTLS `default:"0" json:"use_tls"`
	Version *int64         `json:"version,omitempty"`
}

func (l LoggingSyslogResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingSyslogResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingSyslogResponse) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *LoggingSyslogResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingSyslogResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingSyslogResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingSyslogResponse) GetFormatVersion() *LoggingSyslogResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingSyslogResponse) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *LoggingSyslogResponse) GetIpv4() *string {
	if o == nil {
		return nil
	}
	return o.Ipv4
}

func (o *LoggingSyslogResponse) GetMessageType() *LoggingMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingSyslogResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingSyslogResponse) GetPlacement() *LoggingSyslogResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingSyslogResponse) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *LoggingSyslogResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingSyslogResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingSyslogResponse) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *LoggingSyslogResponse) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *LoggingSyslogResponse) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *LoggingSyslogResponse) GetTLSHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSHostname
}

func (o *LoggingSyslogResponse) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingSyslogResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingSyslogResponse) GetUseTLS() *LoggingUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}

func (o *LoggingSyslogResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
