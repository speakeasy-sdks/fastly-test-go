// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
	"time"
)

// PoolResponseType - What type of load balance group to use.
type PoolResponseType string

const (
	PoolResponseTypeRandom PoolResponseType = "random"
	PoolResponseTypeHash   PoolResponseType = "hash"
	PoolResponseTypeClient PoolResponseType = "client"
)

func (e PoolResponseType) ToPointer() *PoolResponseType {
	return &e
}

func (e *PoolResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "random":
		fallthrough
	case "hash":
		fallthrough
	case "client":
		*e = PoolResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PoolResponseType: %v", v)
	}
}

// PoolResponseUseTLS - Whether to use TLS.
type PoolResponseUseTLS int64

const (
	PoolResponseUseTLSZero PoolResponseUseTLS = 0
	PoolResponseUseTLSOne  PoolResponseUseTLS = 1
)

func (e PoolResponseUseTLS) ToPointer() *PoolResponseUseTLS {
	return &e
}

func (e *PoolResponseUseTLS) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = PoolResponseUseTLS(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PoolResponseUseTLS: %v", v)
	}
}

type PoolResponse struct {
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty"`
	// How long to wait for a timeout in milliseconds. Optional.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// How long to wait for the first byte in milliseconds. Optional.
	FirstByteTimeout *int64 `json:"first_byte_timeout,omitempty"`
	// Name of the healthcheck to use with this pool. Can be empty and could be reused across multiple backend and pools.
	Healthcheck *string `json:"healthcheck,omitempty"`
	ID          *string `json:"id,omitempty"`
	// Maximum number of connections. Optional.
	MaxConnDefault *int64 `default:"200" json:"max_conn_default"`
	// Maximum allowed TLS version on connections to this server. Optional.
	MaxTLSVersion *int64 `json:"max_tls_version,omitempty"`
	// Minimum allowed TLS version on connections to this server. Optional.
	MinTLSVersion *int64 `json:"min_tls_version,omitempty"`
	// Name for the Pool.
	Name *string `json:"name,omitempty"`
	// The hostname to [override the Host header](https://docs.fastly.com/en/guides/specifying-an-override-host). Defaults to `null` meaning no override of the Host header will occur. This setting can also be added to a Server definition. If the field is set on a Server definition it will override the Pool setting.
	OverrideHost *string `default:"null" json:"override_host"`
	// Percentage of capacity (`0-100`) that needs to be operationally available for a pool to be considered up.
	Quorum *int64 `default:"75" json:"quorum"`
	// Condition which, if met, will select this configuration during a request. Optional.
	RequestCondition *string `json:"request_condition,omitempty"`
	ServiceID        *string `json:"service_id,omitempty"`
	// Selected POP to serve as a shield for the servers. Defaults to `null` meaning no origin shielding if not set. Refer to the [POPs API endpoint](/reference/api/utils/pops/) to get a list of available POPs used for shielding.
	Shield *string `default:"null" json:"shield"`
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert *string `default:"null" json:"tls_ca_cert"`
	// The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
	TLSCertHostname *string `default:"null" json:"tls_cert_hostname"`
	// Be strict on checking TLS certs. Optional.
	TLSCheckCert *int64 `json:"tls_check_cert,omitempty"`
	// List of OpenSSL ciphers (see the [openssl.org manpages](https://www.openssl.org/docs/man1.1.1/man1/ciphers.html) for details). Optional.
	TLSCiphers *string `json:"tls_ciphers,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert *string `default:"null" json:"tls_client_cert"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey *string `default:"null" json:"tls_client_key"`
	// SNI hostname. Optional.
	TLSSniHostname *string `json:"tls_sni_hostname,omitempty"`
	// What type of load balance group to use.
	Type *PoolResponseType `json:"type,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Whether to use TLS.
	UseTLS  *PoolResponseUseTLS `default:"0" json:"use_tls"`
	Version *int64              `json:"version,omitempty"`
}

func (p PoolResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PoolResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PoolResponse) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *PoolResponse) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *PoolResponse) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PoolResponse) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *PoolResponse) GetFirstByteTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.FirstByteTimeout
}

func (o *PoolResponse) GetHealthcheck() *string {
	if o == nil {
		return nil
	}
	return o.Healthcheck
}

func (o *PoolResponse) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PoolResponse) GetMaxConnDefault() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxConnDefault
}

func (o *PoolResponse) GetMaxTLSVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTLSVersion
}

func (o *PoolResponse) GetMinTLSVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.MinTLSVersion
}

func (o *PoolResponse) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PoolResponse) GetOverrideHost() *string {
	if o == nil {
		return nil
	}
	return o.OverrideHost
}

func (o *PoolResponse) GetQuorum() *int64 {
	if o == nil {
		return nil
	}
	return o.Quorum
}

func (o *PoolResponse) GetRequestCondition() *string {
	if o == nil {
		return nil
	}
	return o.RequestCondition
}

func (o *PoolResponse) GetServiceID() *string {
	if o == nil {
		return nil
	}
	return o.ServiceID
}

func (o *PoolResponse) GetShield() *string {
	if o == nil {
		return nil
	}
	return o.Shield
}

func (o *PoolResponse) GetTLSCaCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSCaCert
}

func (o *PoolResponse) GetTLSCertHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSCertHostname
}

func (o *PoolResponse) GetTLSCheckCert() *int64 {
	if o == nil {
		return nil
	}
	return o.TLSCheckCert
}

func (o *PoolResponse) GetTLSCiphers() *string {
	if o == nil {
		return nil
	}
	return o.TLSCiphers
}

func (o *PoolResponse) GetTLSClientCert() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientCert
}

func (o *PoolResponse) GetTLSClientKey() *string {
	if o == nil {
		return nil
	}
	return o.TLSClientKey
}

func (o *PoolResponse) GetTLSSniHostname() *string {
	if o == nil {
		return nil
	}
	return o.TLSSniHostname
}

func (o *PoolResponse) GetType() *PoolResponseType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PoolResponse) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *PoolResponse) GetUseTLS() *PoolResponseUseTLS {
	if o == nil {
		return nil
	}
	return o.UseTLS
}

func (o *PoolResponse) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}
