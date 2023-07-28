// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SchemasWafFirewallVersionDataAttributes struct {
	// Allowed HTTP versions.
	AllowedHTTPVersions *string `json:"allowed_http_versions,omitempty"`
	// A space-separated list of HTTP method names.
	AllowedMethods *string `json:"allowed_methods,omitempty"`
	// Allowed request content types.
	AllowedRequestContentType *string `json:"allowed_request_content_type,omitempty"`
	// Allowed request content type charset.
	AllowedRequestContentTypeCharset *string `json:"allowed_request_content_type_charset,omitempty"`
	// The maximum allowed length of an argument.
	ArgLength *int64 `json:"arg_length,omitempty"`
	// The maximum allowed argument name length.
	ArgNameLength *int64 `json:"arg_name_length,omitempty"`
	// The maximum allowed size of all files (in bytes).
	CombinedFileSizes *int64 `json:"combined_file_sizes,omitempty"`
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty"`
	// Score value to add for critical anomalies.
	CriticalAnomalyScore *int64 `json:"critical_anomaly_score,omitempty"`
	// CRS validate UTF8 encoding.
	CrsValidateUTF8Encoding *bool `json:"crs_validate_utf8_encoding,omitempty"`
	// Score value to add for error anomalies.
	ErrorAnomalyScore *int64 `json:"error_anomaly_score,omitempty"`
	// A space-separated list of country codes in ISO 3166-1 (two-letter) format.
	HighRiskCountryCodes *string `json:"high_risk_country_codes,omitempty"`
	// HTTP violation threshold.
	HTTPViolationScoreThreshold *int64 `json:"http_violation_score_threshold,omitempty"`
	// Inbound anomaly threshold.
	InboundAnomalyScoreThreshold *int64 `json:"inbound_anomaly_score_threshold,omitempty"`
	// Local file inclusion attack threshold.
	LfiScoreThreshold *int64 `json:"lfi_score_threshold,omitempty"`
	// Whether a specific firewall version is locked from being modified.
	Locked *bool `json:"locked,omitempty"`
	// The maximum allowed file size, in bytes.
	MaxFileSize *int64 `json:"max_file_size,omitempty"`
	// The maximum number of arguments allowed.
	MaxNumArgs *int64 `json:"max_num_args,omitempty"`
	// Score value to add for notice anomalies.
	NoticeAnomalyScore *int64 `json:"notice_anomaly_score,omitempty"`
	Number             *int64 `json:"number,omitempty"`
	// The configured paranoia level.
	ParanoiaLevel *int64 `json:"paranoia_level,omitempty"`
	// PHP injection threshold.
	PhpInjectionScoreThreshold *int64 `json:"php_injection_score_threshold,omitempty"`
	// Remote code execution threshold.
	RceScoreThreshold *int64 `json:"rce_score_threshold,omitempty"`
	// A space-separated list of allowed file extensions.
	RestrictedExtensions *string `json:"restricted_extensions,omitempty"`
	// A space-separated list of allowed header names.
	RestrictedHeaders *string `json:"restricted_headers,omitempty"`
	// Remote file inclusion attack threshold.
	RfiScoreThreshold *int64 `json:"rfi_score_threshold,omitempty"`
	// Session fixation attack threshold.
	SessionFixationScoreThreshold *int64 `json:"session_fixation_score_threshold,omitempty"`
	// SQL injection attack threshold.
	SQLInjectionScoreThreshold *int64 `json:"sql_injection_score_threshold,omitempty"`
	// The maximum size of argument names and values.
	TotalArgLength *int64 `json:"total_arg_length,omitempty"`
	// Score value to add for warning anomalies.
	WarningAnomalyScore *int64 `json:"warning_anomaly_score,omitempty"`
	// XSS attack threshold.
	XSSScoreThreshold *int64 `json:"xss_score_threshold,omitempty"`
}

func (o *SchemasWafFirewallVersionDataAttributes) GetAllowedHTTPVersions() *string {
	if o == nil {
		return nil
	}
	return o.AllowedHTTPVersions
}

func (o *SchemasWafFirewallVersionDataAttributes) GetAllowedMethods() *string {
	if o == nil {
		return nil
	}
	return o.AllowedMethods
}

func (o *SchemasWafFirewallVersionDataAttributes) GetAllowedRequestContentType() *string {
	if o == nil {
		return nil
	}
	return o.AllowedRequestContentType
}

func (o *SchemasWafFirewallVersionDataAttributes) GetAllowedRequestContentTypeCharset() *string {
	if o == nil {
		return nil
	}
	return o.AllowedRequestContentTypeCharset
}

func (o *SchemasWafFirewallVersionDataAttributes) GetArgLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ArgLength
}

func (o *SchemasWafFirewallVersionDataAttributes) GetArgNameLength() *int64 {
	if o == nil {
		return nil
	}
	return o.ArgNameLength
}

func (o *SchemasWafFirewallVersionDataAttributes) GetCombinedFileSizes() *int64 {
	if o == nil {
		return nil
	}
	return o.CombinedFileSizes
}

func (o *SchemasWafFirewallVersionDataAttributes) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *SchemasWafFirewallVersionDataAttributes) GetCriticalAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.CriticalAnomalyScore
}

func (o *SchemasWafFirewallVersionDataAttributes) GetCrsValidateUTF8Encoding() *bool {
	if o == nil {
		return nil
	}
	return o.CrsValidateUTF8Encoding
}

func (o *SchemasWafFirewallVersionDataAttributes) GetErrorAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.ErrorAnomalyScore
}

func (o *SchemasWafFirewallVersionDataAttributes) GetHighRiskCountryCodes() *string {
	if o == nil {
		return nil
	}
	return o.HighRiskCountryCodes
}

func (o *SchemasWafFirewallVersionDataAttributes) GetHTTPViolationScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.HTTPViolationScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetInboundAnomalyScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.InboundAnomalyScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetLfiScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LfiScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetLocked() *bool {
	if o == nil {
		return nil
	}
	return o.Locked
}

func (o *SchemasWafFirewallVersionDataAttributes) GetMaxFileSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxFileSize
}

func (o *SchemasWafFirewallVersionDataAttributes) GetMaxNumArgs() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxNumArgs
}

func (o *SchemasWafFirewallVersionDataAttributes) GetNoticeAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.NoticeAnomalyScore
}

func (o *SchemasWafFirewallVersionDataAttributes) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *SchemasWafFirewallVersionDataAttributes) GetParanoiaLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.ParanoiaLevel
}

func (o *SchemasWafFirewallVersionDataAttributes) GetPhpInjectionScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.PhpInjectionScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetRceScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.RceScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetRestrictedExtensions() *string {
	if o == nil {
		return nil
	}
	return o.RestrictedExtensions
}

func (o *SchemasWafFirewallVersionDataAttributes) GetRestrictedHeaders() *string {
	if o == nil {
		return nil
	}
	return o.RestrictedHeaders
}

func (o *SchemasWafFirewallVersionDataAttributes) GetRfiScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.RfiScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetSessionFixationScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.SessionFixationScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetSQLInjectionScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.SQLInjectionScoreThreshold
}

func (o *SchemasWafFirewallVersionDataAttributes) GetTotalArgLength() *int64 {
	if o == nil {
		return nil
	}
	return o.TotalArgLength
}

func (o *SchemasWafFirewallVersionDataAttributes) GetWarningAnomalyScore() *int64 {
	if o == nil {
		return nil
	}
	return o.WarningAnomalyScore
}

func (o *SchemasWafFirewallVersionDataAttributes) GetXSSScoreThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.XSSScoreThreshold
}

type SchemasWafFirewallVersionData struct {
	Attributes *SchemasWafFirewallVersionDataAttributes `json:"attributes,omitempty"`
	// Resource type.
	Type *TypeWafFirewallVersion `json:"type,omitempty"`
}

func (o *SchemasWafFirewallVersionData) GetAttributes() *SchemasWafFirewallVersionDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *SchemasWafFirewallVersionData) GetType() *TypeWafFirewallVersion {
	if o == nil {
		return nil
	}
	return o.Type
}
