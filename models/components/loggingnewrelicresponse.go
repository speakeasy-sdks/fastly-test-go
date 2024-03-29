// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// LoggingNewrelicResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingNewrelicResponseFormatVersion int64

const (
	LoggingNewrelicResponseFormatVersionOne LoggingNewrelicResponseFormatVersion = 1
	LoggingNewrelicResponseFormatVersionTwo LoggingNewrelicResponseFormatVersion = 2
)

func (e LoggingNewrelicResponseFormatVersion) ToPointer() *LoggingNewrelicResponseFormatVersion {
	return &e
}

func (e *LoggingNewrelicResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingNewrelicResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicResponseFormatVersion: %v", v)
	}
}

// LoggingNewrelicResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingNewrelicResponsePlacement string

const (
	LoggingNewrelicResponsePlacementNone     LoggingNewrelicResponsePlacement = "none"
	LoggingNewrelicResponsePlacementWafDebug LoggingNewrelicResponsePlacement = "waf_debug"
)

func (e LoggingNewrelicResponsePlacement) ToPointer() *LoggingNewrelicResponsePlacement {
	return &e
}

func (e *LoggingNewrelicResponsePlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		*e = LoggingNewrelicResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicResponsePlacement: %v", v)
	}
}

// LoggingNewrelicResponseRegion - The region to which to stream logs.
type LoggingNewrelicResponseRegion string

const (
	LoggingNewrelicResponseRegionUs LoggingNewrelicResponseRegion = "US"
	LoggingNewrelicResponseRegionEu LoggingNewrelicResponseRegion = "EU"
)

func (e LoggingNewrelicResponseRegion) ToPointer() *LoggingNewrelicResponseRegion {
	return &e
}

func (e *LoggingNewrelicResponseRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = LoggingNewrelicResponseRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicResponseRegion: %v", v)
	}
}

type LoggingNewrelicResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that New Relic Logs can ingest.
	Format *string `default:"{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingNewrelicResponseFormatVersion `default:"2" json:"format_version"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingNewrelicResponsePlacement `json:"placement,omitempty"`
	// The region to which to stream logs.
	Region *LoggingNewrelicResponseRegion `default:"US" json:"region"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// The Insert API key from the Account page of your New Relic account. Required.
	Token *string `json:"token,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (l LoggingNewrelicResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingNewrelicResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingNewrelicResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingNewrelicResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingNewrelicResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingNewrelicResponse) GetFormatVersion() *LoggingNewrelicResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingNewrelicResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingNewrelicResponse) GetPlacement() *LoggingNewrelicResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingNewrelicResponse) GetRegion() *LoggingNewrelicResponseRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingNewrelicResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingNewrelicResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingNewrelicResponse) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingNewrelicResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingNewrelicResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
