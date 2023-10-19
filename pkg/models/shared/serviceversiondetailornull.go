// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
	"time"
)

// ServiceVersionDetailOrNullSettings - List of default settings for this service.
type ServiceVersionDetailOrNullSettings struct {
	// The default host name for the version.
	GeneralDefaultHost *string `json:"general.default_host,omitempty"`
	// The default time-to-live (TTL) for the version.
	GeneralDefaultTTL *int64 `json:"general.default_ttl,omitempty"`
	// Enables serving a stale object if there is an error.
	GeneralStaleIfError *bool `default:"false" json:"general.stale_if_error"`
	// The default time-to-live (TTL) for serving the stale object for the version.
	GeneralStaleIfErrorTTL *int64 `default:"43200" json:"general.stale_if_error_ttl"`
}

func (s ServiceVersionDetailOrNullSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceVersionDetailOrNullSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceVersionDetailOrNullSettings) GetGeneralDefaultHost() *string {
	if o == nil {
		return nil
	}
	return o.GeneralDefaultHost
}

func (o *ServiceVersionDetailOrNullSettings) GetGeneralDefaultTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.GeneralDefaultTTL
}

func (o *ServiceVersionDetailOrNullSettings) GetGeneralStaleIfError() *bool {
	if o == nil {
		return nil
	}
	return o.GeneralStaleIfError
}

func (o *ServiceVersionDetailOrNullSettings) GetGeneralStaleIfErrorTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.GeneralStaleIfErrorTTL
}

type ServiceVersionDetailOrNullWordpress struct {
}

type ServiceVersionDetailOrNull struct {
	// Whether this is the active version or not.
	Active *bool `default:"false" json:"active"`
	// List of backends associated to this service.
	Backends []BackendResponse `json:"backends,omitempty"`
	// List of cache settings associated to this service.
	CacheSettings []CacheSettingResponse `json:"cache_settings,omitempty"`
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty"`
	// List of conditions associated to this service.
	Conditions []ConditionResponse `json:"conditions,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Unused at this time.
	Deployed *bool `json:"deployed,omitempty"`
	// List of directors associated to this service.
	Directors []Director `json:"directors,omitempty"`
	// List of domains associated to this service.
	Domains []DomainResponse `json:"domains,omitempty"`
	// List of gzip rules associated to this service.
	Gzips []GzipResponse `json:"gzips,omitempty"`
	// List of headers associated to this service.
	Headers []HeaderResponse `json:"headers,omitempty"`
	// List of healthchecks associated to this service.
	Healthchecks []HealthcheckResponse `json:"healthchecks,omitempty"`
	// Whether this version is locked or not. Objects can not be added or edited on locked versions.
	Locked *bool `default:"false" json:"locked"`
	// The number of this version.
	Number *int64 `json:"number,omitempty"`
	// List of request settings for this service.
	RequestSettings []RequestSettingsResponse `json:"request_settings,omitempty"`
	// List of response objects for this service.
	ResponseObjects []ResponseObjectResponse `json:"response_objects,omitempty"`
	ServiceID       *string                  `json:"service_id,omitempty"`
	// List of default settings for this service.
	Settings *ServiceVersionDetailOrNullSettings `json:"settings,omitempty"`
	// List of VCL snippets for this service.
	Snippets []SchemasSnippetResponse `json:"snippets,omitempty"`
	// Unused at this time.
	Staging *bool `default:"false" json:"staging"`
	// Unused at this time.
	Testing *bool `default:"false" json:"testing"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// List of VCL files for this service.
	Vcls []SchemasVclResponse `json:"vcls,omitempty"`
	// A list of Wordpress rules with this service.
	Wordpress []ServiceVersionDetailOrNullWordpress `json:"wordpress,omitempty"`
}

func (s ServiceVersionDetailOrNull) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceVersionDetailOrNull) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceVersionDetailOrNull) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ServiceVersionDetailOrNull) GetBackends() []BackendResponse {
	if o == nil {
		return nil
	}
	return o.Backends
}

func (o *ServiceVersionDetailOrNull) GetCacheSettings() []CacheSettingResponse {
	if o == nil {
		return nil
	}
	return o.CacheSettings
}

func (o *ServiceVersionDetailOrNull) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *ServiceVersionDetailOrNull) GetConditions() []ConditionResponse {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *ServiceVersionDetailOrNull) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ServiceVersionDetailOrNull) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *ServiceVersionDetailOrNull) GetDeployed() *bool {
	if o == nil {
		return nil
	}
	return o.Deployed
}

func (o *ServiceVersionDetailOrNull) GetDirectors() []Director {
	if o == nil {
		return nil
	}
	return o.Directors
}

func (o *ServiceVersionDetailOrNull) GetDomains() []DomainResponse {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *ServiceVersionDetailOrNull) GetGzips() []GzipResponse {
	if o == nil {
		return nil
	}
	return o.Gzips
}

func (o *ServiceVersionDetailOrNull) GetHeaders() []HeaderResponse {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *ServiceVersionDetailOrNull) GetHealthchecks() []HealthcheckResponse {
	if o == nil {
		return nil
	}
	return o.Healthchecks
}

func (o *ServiceVersionDetailOrNull) GetLocked() *bool {
	if o == nil {
		return nil
	}
	return o.Locked
}

func (o *ServiceVersionDetailOrNull) GetNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *ServiceVersionDetailOrNull) GetRequestSettings() []RequestSettingsResponse {
	if o == nil {
		return nil
	}
	return o.RequestSettings
}

func (o *ServiceVersionDetailOrNull) GetResponseObjects() []ResponseObjectResponse {
	if o == nil {
		return nil
	}
	return o.ResponseObjects
}

func (o *ServiceVersionDetailOrNull) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *ServiceVersionDetailOrNull) GetSettings() *ServiceVersionDetailOrNullSettings {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *ServiceVersionDetailOrNull) GetSnippets() []SchemasSnippetResponse {
	if o == nil {
		return nil
	}
	return o.Snippets
}

func (o *ServiceVersionDetailOrNull) GetStaging() *bool {
	if o == nil {
		return nil
	}
	return o.Staging
}

func (o *ServiceVersionDetailOrNull) GetTesting() *bool {
	if o == nil {
		return nil
	}
	return o.Testing
}

func (o *ServiceVersionDetailOrNull) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ServiceVersionDetailOrNull) GetVcls() []SchemasVclResponse {
	if o == nil {
		return nil
	}
	return o.Vcls
}

func (o *ServiceVersionDetailOrNull) GetWordpress() []ServiceVersionDetailOrNullWordpress {
	if o == nil {
		return nil
	}
	return o.Wordpress
}
