// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/types"
	"Fastly/pkg/utils"
)

type WafFirewallVersionDataAttributesInput struct {
	// Allowed HTTP versions.
	AllowedHTTPVersions *string `default:"HTTP/1.0 HTTP/1.1 HTTP/2" json:"allowed_http_versions"`
	// A space-separated list of HTTP method names.
	AllowedMethods *string `default:"GET HEAD POST OPTIONS PUT PATCH DELETE" json:"allowed_methods"`
	// Allowed request content types.
	AllowedRequestContentType *string `default:"application/x-www-form-urlencoded|multipart/form-data|text/xml|application/xml|application/x-amf|application/json|text/plain" json:"allowed_request_content_type"`
	// Allowed request content type charset.
	AllowedRequestContentTypeCharset *string `default:"utf-8|iso-8859-1|iso-8859-15|windows-1252" json:"allowed_request_content_type_charset"`
	// The maximum allowed length of an argument.
	ArgLength *int64 `default:"400" json:"arg_length"`
	// The maximum allowed argument name length.
	ArgNameLength *int64 `default:"100" json:"arg_name_length"`
	// The maximum allowed size of all files (in bytes).
	CombinedFileSizes *int64 `default:"10000000" json:"combined_file_sizes"`
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty"`
	// Score value to add for critical anomalies.
	CriticalAnomalyScore *int64 `default:"6" json:"critical_anomaly_score"`
	// CRS validate UTF8 encoding.
	CrsValidateUTF8Encoding *bool `json:"crs_validate_utf8_encoding,omitempty"`
	// Score value to add for error anomalies.
	ErrorAnomalyScore *int64 `default:"5" json:"error_anomaly_score"`
	// A space-separated list of country codes in ISO 3166-1 (two-letter) format.
	HighRiskCountryCodes *string `json:"high_risk_country_codes,omitempty"`
	// HTTP violation threshold.
	HTTPViolationScoreThreshold *int64 `json:"http_violation_score_threshold,omitempty"`
	// Inbound anomaly threshold.
	InboundAnomalyScoreThreshold *int64 `json:"inbound_anomaly_score_threshold,omitempty"`
	// Local file inclusion attack threshold.
	LfiScoreThreshold *int64 `json:"lfi_score_threshold,omitempty"`
	// Whether a specific firewall version is locked from being modified.
	Locked *bool `default:"false" json:"locked"`
	// The maximum allowed file size, in bytes.
	MaxFileSize *int64 `default:"10000000" json:"max_file_size"`
	// The maximum number of arguments allowed.
	MaxNumArgs *int64 `default:"255" json:"max_num_args"`
	// Score value to add for notice anomalies.
	NoticeAnomalyScore *int64 `default:"4" json:"notice_anomaly_score"`
	// The configured paranoia level.
	ParanoiaLevel *int64 `default:"1" json:"paranoia_level"`
	// PHP injection threshold.
	PhpInjectionScoreThreshold *int64 `json:"php_injection_score_threshold,omitempty"`
	// Remote code execution threshold.
	RceScoreThreshold *int64 `json:"rce_score_threshold,omitempty"`
	// A space-separated list of allowed file extensions.
	RestrictedExtensions *string `default:".asa/ .asax/ .ascx/ .axd/ .backup/ .bak/ .bat/ .cdx/ .cer/ .cfg/ .cmd/ .com/ .config/ .conf/ .cs/ .csproj/ .csr/ .dat/ .db/ .dbf/ .dll/ .dos/ .htr/ .htw/ .ida/ .idc/ .idq/ .inc/ .ini/ .key/ .licx/ .lnk/ .log/ .mdb/ .old/ .pass/ .pdb/ .pol/ .printer/ .pwd/ .resources/ .resx/ .sql/ .sys/ .vb/ .vbs/ .vbproj/ .vsdisco/ .webinfo/ .xsd/ .xsx" json:"restricted_extensions"`
	// A space-separated list of allowed header names.
	RestrictedHeaders *string `default:"/proxy/ /lock-token/ /content-range/ /translate/ /if/" json:"restricted_headers"`
	// Remote file inclusion attack threshold.
	RfiScoreThreshold *int64 `json:"rfi_score_threshold,omitempty"`
	// Session fixation attack threshold.
	SessionFixationScoreThreshold *int64 `json:"session_fixation_score_threshold,omitempty"`
	// SQL injection attack threshold.
	SQLInjectionScoreThreshold *int64 `json:"sql_injection_score_threshold,omitempty"`
	// The maximum size of argument names and values.
	TotalArgLength *int64 `default:"6400" json:"total_arg_length"`
	// Score value to add for warning anomalies.
	WarningAnomalyScore *int64 `json:"warning_anomaly_score,omitempty"`
	// XSS attack threshold.
	XSSScoreThreshold *int64 `json:"xss_score_threshold,omitempty"`
}

func (w WafFirewallVersionDataAttributesInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafFirewallVersionDataAttributesInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafFirewallVersionDataAttributesInput) GetAllowedHTTPVersions() *string {
	if o == nil {
		return nil
	}
	return o.AllowedHTTPVersions
}

func (o *WafFirewallVersionDataAttributesInput) GetAllowedMethods() *string {
	if o == nil {
		return nil
	}
	return o.AllowedMethods
}

func (o *WafFirewallVersionDataAttributesInput) GetAllowedRequestContentType() *string {
	if o == nil {
		return nil
	}
	return o.AllowedRequestContentType
}

func (o *WafFirewallVersionDataAttributesInput) GetAllowedRequestContentTypeCharset() *string {
	if o == nil {
		return nil
	}
	return o.AllowedRequestContentTypeCharset
}

func (o *WafFirewallVersionDataAttributesInput) GetArgLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ArgLength
}

func (o *WafFirewallVersionDataAttributesInput) GetArgNameLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ArgNameLength
}

func (o *WafFirewallVersionDataAttributesInput) GetCombinedFileSizes() *int64 {
	if o == nil {
		return nil
	}
	return o.CombinedFileSizes
}

func (o *WafFirewallVersionDataAttributesInput) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *WafFirewallVersionDataAttributesInput) GetCriticalAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.CriticalAnomalyScore
}

func (o *WafFirewallVersionDataAttributesInput) GetCrsValidateUTF8Encoding() *bool {
	if o == nil {
		return nil
	}
	return o.CrsValidateUTF8Encoding
}

func (o *WafFirewallVersionDataAttributesInput) GetErrorAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.ErrorAnomalyScore
}

func (o *WafFirewallVersionDataAttributesInput) GetHighRiskCountryCodes() *string {
	if o == nil {
		return nil
	}
	return o.HighRiskCountryCodes
}

func (o *WafFirewallVersionDataAttributesInput) GetHTTPViolationScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPViolationScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetInboundAnomalyScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.InboundAnomalyScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetLfiScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LfiScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetLocked() *bool {
	if o == nil {
		return nil
	}
	return o.Locked
}

func (o *WafFirewallVersionDataAttributesInput) GetMaxFileSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *WafFirewallVersionDataAttributesInput) GetMaxNumArgs() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxNumArgs
}

func (o *WafFirewallVersionDataAttributesInput) GetNoticeAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.NoticeAnomalyScore
}

func (o *WafFirewallVersionDataAttributesInput) GetParanoiaLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.ParanoiaLevel
}

func (o *WafFirewallVersionDataAttributesInput) GetPhpInjectionScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.PhpInjectionScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetRceScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.RceScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetRestrictedExtensions() *string {
	if o == nil {
		return nil
	}
	return o.RestrictedExtensions
}

func (o *WafFirewallVersionDataAttributesInput) GetRestrictedHeaders() *string {
	if o == nil {
		return nil
	}
	return o.RestrictedHeaders
}

func (o *WafFirewallVersionDataAttributesInput) GetRfiScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.RfiScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetSessionFixationScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionFixationScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetSQLInjectionScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.SQLInjectionScoreThreshold
}

func (o *WafFirewallVersionDataAttributesInput) GetTotalArgLength() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalArgLength
}

func (o *WafFirewallVersionDataAttributesInput) GetWarningAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.WarningAnomalyScore
}

func (o *WafFirewallVersionDataAttributesInput) GetXSSScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.XSSScoreThreshold
}

type WafFirewallVersionDataInput struct {
	Attributes *WafFirewallVersionDataAttributesInput `json:"attributes,omitempty"`
	// Resource type.
	type_ *string `const:"waf_firewall_version" json:"type"`
}

func (w WafFirewallVersionDataInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WafFirewallVersionDataInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WafFirewallVersionDataInput) GetAttributes() *WafFirewallVersionDataAttributesInput {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *WafFirewallVersionDataInput) GetType() *string {
	return types.String("waf_firewall_version")
}
