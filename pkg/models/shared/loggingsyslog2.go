// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"encoding/json"
	"fmt"
)

// LoggingSyslogFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingSyslogFormatVersion int64

const (
	LoggingSyslogFormatVersionOne LoggingSyslogFormatVersion = 1
	LoggingSyslogFormatVersionTwo LoggingSyslogFormatVersion = 2
)

func (e LoggingSyslogFormatVersion) ToPointer() *LoggingSyslogFormatVersion {
	return &e
}

func (e *LoggingSyslogFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingSyslogFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSyslogFormatVersion: %v", v)
	}
}

// LoggingSyslogPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingSyslogPlacement string

const (
	LoggingSyslogPlacementNone                   LoggingSyslogPlacement = "none"
	LoggingSyslogPlacementWafDebug               LoggingSyslogPlacement = "waf_debug"
	LoggingSyslogPlacementLessThanNilGreaterThan LoggingSyslogPlacement = "<nil>"
)

func (e LoggingSyslogPlacement) ToPointer() *LoggingSyslogPlacement {
	return &e
}

func (e *LoggingSyslogPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSyslogPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSyslogPlacement: %v", v)
	}
}

type LoggingSyslog2 struct {
	// A hostname or IPv4 address.
	Address *string `form:"name=address"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingSyslogFormatVersion `default:"2" form:"name=format_version"`
	// The hostname used for the syslog endpoint.
	Hostname *string `form:"name=hostname"`
	// The IPv4 address used for the syslog endpoint.
	Ipv4 *string `form:"name=ipv4"`
	// How the message should be formatted.
	MessageType *LoggingMessageType `default:"classic" form:"name=message_type"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingSyslogPlacement `form:"name=placement"`
	// The port number.
	Port *int64 `default:"514" form:"name=port"`
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
	// Whether to prepend each message with a specific token.
	Token *string `default:"null" form:"name=token"`
	// Whether to use TLS.
	UseTLS *LoggingUseTLS `default:"0" form:"name=use_tls"`
}

func (l LoggingSyslog2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingSyslog2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingSyslog2) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *LoggingSyslog2) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingSyslog2) GetFormatVersion() *LoggingSyslogFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingSyslog2) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *LoggingSyslog2) GetIpv4() *string {
	if o == nil {
		return nil
	}
	return o.Ipv4
}

func (o *LoggingSyslog2) GetMessageType() *LoggingMessageType {
	if o == nil {
		return nil
	}
	return o.MessageType
}

func (o *LoggingSyslog2) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingSyslog2) GetPlacement() *LoggingSyslogPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingSyslog2) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *LoggingSyslog2) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingSyslog2) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *LoggingSyslog2) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *LoggingSyslog2) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *LoggingSyslog2) GetTLSHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSHostname
}

func (o *LoggingSyslog2) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingSyslog2) GetUseTLS() *LoggingUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}
