// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// LoggingDatadogResponseFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingDatadogResponseFormatVersion int64

const (
	LoggingDatadogResponseFormatVersionOne LoggingDatadogResponseFormatVersion = 1
	LoggingDatadogResponseFormatVersionTwo LoggingDatadogResponseFormatVersion = 2
)

func (e LoggingDatadogResponseFormatVersion) ToPointer() *LoggingDatadogResponseFormatVersion {
	return &e
}

func (e *LoggingDatadogResponseFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingDatadogResponseFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingDatadogResponseFormatVersion: %v", v)
	}
}

// LoggingDatadogResponsePlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingDatadogResponsePlacement string

const (
	LoggingDatadogResponsePlacementNone                   LoggingDatadogResponsePlacement = "none"
	LoggingDatadogResponsePlacementWafDebug               LoggingDatadogResponsePlacement = "waf_debug"
	LoggingDatadogResponsePlacementLessThanNilGreaterThan LoggingDatadogResponsePlacement = "<nil>"
)

func (e LoggingDatadogResponsePlacement) ToPointer() *LoggingDatadogResponsePlacement {
	return &e
}

func (e *LoggingDatadogResponsePlacement) UnmarshalJSON(data []byte) error {
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
		*e = LoggingDatadogResponsePlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingDatadogResponsePlacement: %v", v)
	}
}

// LoggingDatadogResponseRegion - The region that log data will be sent to.
type LoggingDatadogResponseRegion string

const (
	LoggingDatadogResponseRegionUs LoggingDatadogResponseRegion = "US"
	LoggingDatadogResponseRegionEu LoggingDatadogResponseRegion = "EU"
)

func (e LoggingDatadogResponseRegion) ToPointer() *LoggingDatadogResponseRegion {
	return &e
}

func (e *LoggingDatadogResponseRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = LoggingDatadogResponseRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingDatadogResponseRegion: %v", v)
	}
}

type LoggingDatadogResponse struct {
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Datadog can ingest.
	//
	Format *string `default:"{\"ddsource\":\"fastly\",\"service\":\"%{req.service_id}V\",\"date\":\"%{begin:%Y-%m-%dT%H:%M:%S%Z}t\",\"time_start\":\"%{begin:%Y-%m-%dT%H:%M:%S%Z}t\",\"time_end\":\"%{end:%Y-%m-%dT%H:%M:%S%Z}t\",\"http\":{\"request_time_ms\":\"%D\",\"method\":\"%m\",\"url\":\"%{json.escape(req.url)}V\",\"useragent\":\"%{User-Agent}i\",\"referer\":\"%{Referer}i\",\"protocol\":\"%H\",\"request_x_forwarded_for\":\"%{X-Forwarded-For}i\",\"status_code\":\"%s\"},\"network\":{\"client\":{\"ip\":\"%h\",\"name\":\"%{client.as.name}V\",\"number\":\"%{client.as.number}V\",\"connection_speed\":\"%{client.geo.conn_speed}V\"},\"destination\":{\"ip\":\"%A\"},\"geoip\":{\"geo_city\":\"%{client.geo.city.utf8}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"geo_continent_code\":\"%{client.geo.continent_code}V\",\"geo_region\":\"%{client.geo.region}V\"},\"bytes_written\":\"%B\",\"bytes_read\":\"%{req.body_bytes_read}V\"},\"host\":\"%{Fastly-Orig-Host}i\",\"origin_host\":\"%v\",\"is_ipv6\":\"%{if(req.is_ipv6, \\\"true\\\", \\\"false\\\")}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"tls_client_protocol\":\"%{json.escape(tls.client.protocol)}V\",\"tls_client_servername\":\"%{json.escape(tls.client.servername)}V\",\"tls_client_cipher\":\"%{json.escape(tls.client.cipher)}V\",\"tls_client_cipher_sha\":\"%{json.escape(tls.client.ciphers_sha)}V\",\"tls_client_tlsexts_sha\":\"%{json.escape(tls.client.tlsexts_sha)}V\",\"is_h2\":\"%{if(fastly_info.is_h2, \\\"true\\\", \\\"false\\\")}V\",\"is_h2_push\":\"%{if(fastly_info.h2.is_push, \\\"true\\\", \\\"false\\\")}V\",\"h2_stream_id\":\"%{fastly_info.h2.stream_id}V\",\"request_accept_content\":\"%{Accept}i\",\"request_accept_language\":\"%{Accept-Language}i\",\"request_accept_encoding\":\"%{Accept-Encoding}i\",\"request_accept_charset\":\"%{Accept-Charset}i\",\"request_connection\":\"%{Connection}i\",\"request_dnt\":\"%{DNT}i\",\"request_forwarded\":\"%{Forwarded}i\",\"request_via\":\"%{Via}i\",\"request_cache_control\":\"%{Cache-Control}i\",\"request_x_requested_with\":\"%{X-Requested-With}i\",\"request_x_att_device_id\":\"%{X-ATT-Device-Id}i\",\"content_type\":\"%{Content-Type}o\",\"is_cacheable\":\"%{if(fastly_info.state~\\\"^(HIT|MISS)$\\\", \\\"true\\\", \\\"false\\\")}V\",\"response_age\":\"%{Age}o\",\"response_cache_control\":\"%{Cache-Control}o\",\"response_expires\":\"%{Expires}o\",\"response_last_modified\":\"%{Last-Modified}o\",\"response_tsv\":\"%{TSV}o\",\"server_datacenter\":\"%{server.datacenter}V\",\"req_header_size\":\"%{req.header_bytes_read}V\",\"resp_header_size\":\"%{resp.header_bytes_written}V\",\"socket_cwnd\":\"%{client.socket.cwnd}V\",\"socket_nexthop\":\"%{client.socket.nexthop}V\",\"socket_tcpi_rcv_mss\":\"%{client.socket.tcpi_rcv_mss}V\",\"socket_tcpi_snd_mss\":\"%{client.socket.tcpi_snd_mss}V\",\"socket_tcpi_rtt\":\"%{client.socket.tcpi_rtt}V\",\"socket_tcpi_rttvar\":\"%{client.socket.tcpi_rttvar}V\",\"socket_tcpi_rcv_rtt\":\"%{client.socket.tcpi_rcv_rtt}V\",\"socket_tcpi_rcv_space\":\"%{client.socket.tcpi_rcv_space}V\",\"socket_tcpi_last_data_sent\":\"%{client.socket.tcpi_last_data_sent}V\",\"socket_tcpi_total_retrans\":\"%{client.socket.tcpi_total_retrans}V\",\"socket_tcpi_delta_retrans\":\"%{client.socket.tcpi_delta_retrans}V\",\"socket_ploss\":\"%{client.socket.ploss}V\"}" json:"format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingDatadogResponseFormatVersion `default:"2" json:"format_version"`
	// The name for the real-time logging configuration.
	Name *string `json:"name,omitempty"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingDatadogResponsePlacement `json:"placement,omitempty"`
	// The region that log data will be sent to.
	Region *LoggingDatadogResponseRegion `default:"US" json:"region"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `json:"response_condition,omitempty"`
	ServiceID         *string `json:"service_id,omitempty"`
	// The API key from your Datadog account. Required.
	Token *string `json:"token,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Version   *int64     `json:"version,omitempty"`
}

func (l LoggingDatadogResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingDatadogResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingDatadogResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LoggingDatadogResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *LoggingDatadogResponse) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingDatadogResponse) GetFormatVersion() *LoggingDatadogResponseFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingDatadogResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingDatadogResponse) GetPlacement() *LoggingDatadogResponsePlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingDatadogResponse) GetRegion() *LoggingDatadogResponseRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingDatadogResponse) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingDatadogResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *LoggingDatadogResponse) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *LoggingDatadogResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *LoggingDatadogResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
