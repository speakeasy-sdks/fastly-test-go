// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"Fastly/internal/utils"
	"encoding/json"
	"fmt"
)

// LoggingSplunkFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingSplunkFormatVersion int64

const (
	LoggingSplunkFormatVersionOne LoggingSplunkFormatVersion = 1
	LoggingSplunkFormatVersionTwo LoggingSplunkFormatVersion = 2
)

func (e LoggingSplunkFormatVersion) ToPointer() *LoggingSplunkFormatVersion {
	return &e
}

func (e *LoggingSplunkFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingSplunkFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSplunkFormatVersion: %v", v)
	}
}

// LoggingSplunkPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingSplunkPlacement string

const (
	LoggingSplunkPlacementNone                   LoggingSplunkPlacement = "none"
	LoggingSplunkPlacementWafDebug               LoggingSplunkPlacement = "waf_debug"
	LoggingSplunkPlacementLessThanNilGreaterThan LoggingSplunkPlacement = "<nil>"
)

func (e LoggingSplunkPlacement) ToPointer() *LoggingSplunkPlacement {
	return &e
}

func (e *LoggingSplunkPlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingSplunkPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingSplunkPlacement: %v", v)
	}
}

type LoggingSplunk struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t "%r" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingSplunkFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingSplunkPlacement `form:"name=placement"`
	// The maximum number of bytes sent in one request. Defaults `0` for unbounded.
	RequestMaxBytes *int64 `default:"0" form:"name=request_max_bytes"`
	// The maximum number of logs sent in one request. Defaults `0` for unbounded.
	RequestMaxEntries *int64 `default:"0" form:"name=request_max_entries"`
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
	// A Splunk token for use in posting logs over HTTP to your collector.
	Token *string `form:"name=token"`
	// The URL to post logs to.
	URL *string `form:"name=url"`
	// Whether to use TLS.
	UseTLS *LoggingUseTLS `default:"0" form:"name=use_tls"`
}

func (l LoggingSplunk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingSplunk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingSplunk) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingSplunk) GetFormatVersion() *LoggingSplunkFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingSplunk) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingSplunk) GetPlacement() *LoggingSplunkPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingSplunk) GetRequestMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxBytes
}

func (o *LoggingSplunk) GetRequestMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestMaxEntries
}

func (o *LoggingSplunk) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingSplunk) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *LoggingSplunk) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *LoggingSplunk) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *LoggingSplunk) GetTLSHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSHostname
}

func (o *LoggingSplunk) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingSplunk) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *LoggingSplunk) GetUseTLS() *LoggingUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}
