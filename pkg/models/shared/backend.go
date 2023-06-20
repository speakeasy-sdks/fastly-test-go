// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Backend struct {
	// A hostname, IPv4, or IPv6 address for the backend. This is the preferred way to specify the location of your backend.
	Address *string `json:"address,omitempty" form:"name=address"`
	// Whether or not this backend should be automatically load balanced. If true, all backends with this setting that don't have a `request_condition` will be selected based on their `weight`.
	AutoLoadbalance *bool `json:"auto_loadbalance,omitempty" form:"name=auto_loadbalance"`
	// Maximum duration in milliseconds that Fastly will wait while receiving no data on a download from a backend. If exceeded, the response received so far will be considered complete and the fetch will end. May be set at runtime using `bereq.between_bytes_timeout`.
	BetweenBytesTimeout *int64 `json:"between_bytes_timeout,omitempty" form:"name=between_bytes_timeout"`
	// Unused.
	ClientCert *string `json:"client_cert,omitempty" form:"name=client_cert"`
	// A freeform descriptive note.
	Comment *string `json:"comment,omitempty" form:"name=comment"`
	// Maximum duration in milliseconds to wait for a connection to this backend to be established. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.connect_timeout`.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty" form:"name=connect_timeout"`
	// Maximum duration in milliseconds to wait for the server response to begin after a TCP connection is established and the request has been sent. If exceeded, the connection is aborted and a synthethic `503` response will be presented instead. May be set at runtime using `bereq.first_byte_timeout`.
	FirstByteTimeout *int64 `json:"first_byte_timeout,omitempty" form:"name=first_byte_timeout"`
	// The name of the healthcheck to use with this backend.
	Healthcheck *string `json:"healthcheck,omitempty" form:"name=healthcheck"`
	// The hostname of the backend. May be used as an alternative to `address` to set the backend location.
	Hostname *string `json:"hostname,omitempty" form:"name=hostname"`
	// IPv4 address of the backend. May be used as an alternative to `address` to set the backend location.
	Ipv4 *string `json:"ipv4,omitempty" form:"name=ipv4"`
	// IPv6 address of the backend. May be used as an alternative to `address` to set the backend location.
	Ipv6 *string `json:"ipv6,omitempty" form:"name=ipv6"`
	// How long in seconds to keep a persistent connection to the backend between requests.
	KeepaliveTime *int64 `json:"keepalive_time,omitempty" form:"name=keepalive_time"`
	// Maximum number of concurrent connections this backend will accept.
	MaxConn *int64 `json:"max_conn,omitempty" form:"name=max_conn"`
	// Maximum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated.
	MaxTLSVersion *string `json:"max_tls_version,omitempty" form:"name=max_tls_version"`
	// Minimum allowed TLS version on SSL connections to this backend. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated.
	MinTLSVersion *string `json:"min_tls_version,omitempty" form:"name=min_tls_version"`
	// The name of the backend.
	Name *string `json:"name,omitempty" form:"name=name"`
	// If set, will replace the client-supplied HTTP `Host` header on connections to this backend. Applied after VCL has been processed, so this setting will take precedence over changing `bereq.http.Host` in VCL.
	OverrideHost *string `json:"override_host,omitempty" form:"name=override_host"`
	// Port on which the backend server is listening for connections from Fastly. Setting `port` to 80 or 443 will also set `use_ssl` automatically (to false and true respectively), unless explicitly overridden by setting `use_ssl` in the same request.
	Port *int64 `json:"port,omitempty" form:"name=port"`
	// Name of a Condition, which if satisfied, will select this backend during a request. If set, will override any `auto_loadbalance` setting. By default, the first backend added to a service is selected for all requests.
	RequestCondition *string `json:"request_condition,omitempty" form:"name=request_condition"`
	// Identifier of the POP to use as a [shield](https://docs.fastly.com/en/guides/shielding).
	Shield *string `json:"shield,omitempty" form:"name=shield"`
	// CA certificate attached to origin.
	SslCaCert *string `json:"ssl_ca_cert,omitempty" form:"name=ssl_ca_cert"`
	// Overrides `ssl_hostname`, but only for cert verification. Does not affect SNI at all.
	SslCertHostname *string `json:"ssl_cert_hostname,omitempty" form:"name=ssl_cert_hostname"`
	// Be strict on checking SSL certs.
	SslCheckCert *bool `json:"ssl_check_cert,omitempty" form:"name=ssl_check_cert"`
	// List of [OpenSSL ciphers](https://www.openssl.org/docs/manmaster/man1/ciphers.html) to support for connections to this origin. If your backend server is not able to negotiate a connection meeting this constraint, a synthetic `503` error response will be generated.
	SslCiphers *string `json:"ssl_ciphers,omitempty" form:"name=ssl_ciphers"`
	// Client certificate attached to origin.
	SslClientCert *string `json:"ssl_client_cert,omitempty" form:"name=ssl_client_cert"`
	// Client key attached to origin.
	SslClientKey *string `json:"ssl_client_key,omitempty" form:"name=ssl_client_key"`
	// Use `ssl_cert_hostname` and `ssl_sni_hostname` to configure certificate validation.
	//
	// Deprecated: this field will be removed in a future release, please migrate away from it as soon as possible.
	SslHostname *string `json:"ssl_hostname,omitempty" form:"name=ssl_hostname"`
	// Overrides `ssl_hostname`, but only for SNI in the handshake. Does not affect cert validation at all.
	SslSniHostname *string `json:"ssl_sni_hostname,omitempty" form:"name=ssl_sni_hostname"`
	// Whether or not to require TLS for connections to this backend.
	UseSsl *bool `json:"use_ssl,omitempty" form:"name=use_ssl"`
	// Weight used to load balance this backend against others. May be any positive integer. If `auto_loadbalance` is true, the chance of this backend being selected is equal to its own weight over the sum of all weights for backends that have `auto_loadbalance` set to true.
	Weight *int64 `json:"weight,omitempty" form:"name=weight"`
}
