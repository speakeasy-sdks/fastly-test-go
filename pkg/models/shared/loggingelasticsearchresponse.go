// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// LoggingElasticsearchResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingElasticsearchResponseFormatVersion int64

const (
	LoggingElasticsearchResponseFormatVersionOne LoggingElasticsearchResponseFormatVersion = 1
	LoggingElasticsearchResponseFormatVersionTwo LoggingElasticsearchResponseFormatVersion = 2
)

func (e LoggingElasticsearchResponseFormatVersion) ToPointer() *LoggingElasticsearchResponseFormatVersion {
	return &e
}

func (e *LoggingElasticsearchResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingElasticsearchResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingElasticsearchResponseFormatVersion: %v", v)
	}
}

// LoggingElasticsearchResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingElasticsearchResponsePlacement string

const (
	LoggingElasticsearchResponsePlacementNone                   LoggingElasticsearchResponsePlacement = "none"
	LoggingElasticsearchResponsePlacementWafDebug               LoggingElasticsearchResponsePlacement = "waf_debug"
	LoggingElasticsearchResponsePlacementLessThanNilGreaterThan LoggingElasticsearchResponsePlacement = "<nil>"
)

func (e LoggingElasticsearchResponsePlacement) ToPointer() *LoggingElasticsearchResponsePlacement {
	return &e
}

func (e *LoggingElasticsearchResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingElasticsearchResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingElasticsearchResponsePlacement: %v", v)
	}
}

// LoggingElasticsearchResponse - OK
type LoggingElasticsearchResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Elasticsearch can ingest.
	Format *string `json:"format,omitempty"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingElasticsearchResponseFormatVersion `json:"format_version,omitempty"`
	// The name of the Elasticsearch index to send documents (logs) to. The index must follow the Elasticsearch [index format rules](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html). We support [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) interpolated variables inside braces prefixed with a pound symbol. For example, `#{%F}` will interpolate as `YYYY-MM-DD` with today's date.
	Index *string `json:"index,omitempty"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Basic Auth password.
	Password *string `json:"password,omitempty"`
	// The ID of the Elasticsearch ingest pipeline to apply pre-process transformations to before indexing. Learn more about creating a pipeline in the [Elasticsearch docs](https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html).
	Pipeline *string `json:"pipeline,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingElasticsearchResponsePlacement `json:"placement,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` for unbounded.
	RequestMaxBytes *int64 `json:"request_max_bytes,omitempty"`
	// The maximum number of logs sent in one request. Defaults `0` for unbounded.
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
	// The URL to stream logs to. Must use HTTPS.
	URL *string `json:"url,omitempty"`
	// Basic Auth username.
	User    *string `json:"user,omitempty"`
	Version *int64  `json:"version,omitempty"`
}
