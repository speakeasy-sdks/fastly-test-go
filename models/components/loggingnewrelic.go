// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// LoggingNewrelicFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingNewrelicFormatVersion int64

const (
	LoggingNewrelicFormatVersionOne LoggingNewrelicFormatVersion = 1
	LoggingNewrelicFormatVersionTwo LoggingNewrelicFormatVersion = 2
)

func (e LoggingNewrelicFormatVersion) ToPointer() *LoggingNewrelicFormatVersion {
	return &e
}

func (e *LoggingNewrelicFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingNewrelicFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicFormatVersion: %v", v)
	}
}

// LoggingNewrelicPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingNewrelicPlacement string

const (
	LoggingNewrelicPlacementNone     LoggingNewrelicPlacement = "none"
	LoggingNewrelicPlacementWafDebug LoggingNewrelicPlacement = "waf_debug"
)

func (e LoggingNewrelicPlacement) ToPointer() *LoggingNewrelicPlacement {
	return &e
}

func (e *LoggingNewrelicPlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		*e = LoggingNewrelicPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicPlacement: %v", v)
	}
}

// LoggingNewrelicRegion - The region to which to stream logs.
type LoggingNewrelicRegion string

const (
	LoggingNewrelicRegionUs LoggingNewrelicRegion = "US"
	LoggingNewrelicRegionEu LoggingNewrelicRegion = "EU"
)

func (e LoggingNewrelicRegion) ToPointer() *LoggingNewrelicRegion {
	return &e
}

func (e *LoggingNewrelicRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = LoggingNewrelicRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingNewrelicRegion: %v", v)
	}
}

type LoggingNewrelic struct {
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that New Relic Logs can ingest.
	Format *string `default:"{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingNewrelicFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingNewrelicPlacement `form:"name=placement"`
	// The region to which to stream logs.
	Region *LoggingNewrelicRegion `default:"US" form:"name=region"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// The Insert API key from the Account page of your New Relic account. Required.
	Token *string `form:"name=token"`
}

func (l LoggingNewrelic) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingNewrelic) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingNewrelic) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingNewrelic) GetFormatVersion() *LoggingNewrelicFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingNewrelic) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingNewrelic) GetPlacement() *LoggingNewrelicPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingNewrelic) GetRegion() *LoggingNewrelicRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingNewrelic) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingNewrelic) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
