// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Healthcheck struct {
	// How often to run the health check in milliseconds.
	CheckInterval *int64 `form:"name=check_interval"`
	// A freeform descriptive note.
	Comment *string `form:"name=comment"`
	// The status code expected from the host.
	ExpectedResponse *int64 `form:"name=expected_response"`
	// Array of custom headers that will be added to the health check probes.
	Headers []string `form:"name=headers"`
	// Which host to check.
	Host *string `form:"name=host"`
	// Whether to use version 1.0 or 1.1 HTTP.
	HTTPVersion *string `form:"name=http_version"`
	// When loading a config, the initial number of probes to be seen as OK.
	Initial *int64 `form:"name=initial"`
	// Which HTTP method to use.
	Method *string `form:"name=method"`
	// The name of the health check.
	Name *string `form:"name=name"`
	// The path to check.
	Path *string `form:"name=path"`
	// How many health checks must succeed to be considered healthy.
	Threshold *int64 `form:"name=threshold"`
	// Timeout in milliseconds.
	Timeout *int64 `form:"name=timeout"`
	// The number of most recent health check queries to keep for this health check.
	Window *int64 `form:"name=window"`
}
