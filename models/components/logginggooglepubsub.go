// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/v2/internal/utils"
)

// LoggingGooglePubsubFormatVersion - The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
type LoggingGooglePubsubFormatVersion int64

const (
	LoggingGooglePubsubFormatVersionOne LoggingGooglePubsubFormatVersion = 1
	LoggingGooglePubsubFormatVersionTwo LoggingGooglePubsubFormatVersion = 2
)

func (e LoggingGooglePubsubFormatVersion) ToPointer() *LoggingGooglePubsubFormatVersion {
	return &e
}

func (e *LoggingGooglePubsubFormatVersion) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		*e = LoggingGooglePubsubFormatVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGooglePubsubFormatVersion: %v", v)
	}
}

// LoggingGooglePubsubPlacement - Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
type LoggingGooglePubsubPlacement string

const (
	LoggingGooglePubsubPlacementNone     LoggingGooglePubsubPlacement = "none"
	LoggingGooglePubsubPlacementWafDebug LoggingGooglePubsubPlacement = "waf_debug"
)

func (e LoggingGooglePubsubPlacement) ToPointer() *LoggingGooglePubsubPlacement {
	return &e
}

func (e *LoggingGooglePubsubPlacement) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		fallthrough
	case "waf_debug":
		*e = LoggingGooglePubsubPlacement(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingGooglePubsubPlacement: %v", v)
	}
}

type LoggingGooglePubsub struct {
	// The name of the Google Cloud Platform service account associated with the target log collection service. Not required if `user` and `secret_key` are provided.
	AccountName *string `form:"name=account_name"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats).
	Format *string `default:"%h %l %u %t \"%r\" %&gt;s %b" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingGooglePubsubFormatVersion `default:"2" form:"name=format_version"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingGooglePubsubPlacement `form:"name=placement"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `form:"name=project_id"`
	// The name of an existing condition in the configured endpoint, or leave blank to always execute.
	ResponseCondition *string `form:"name=response_condition"`
	// Your Google Cloud Platform account secret key. The `private_key` field in your service account authentication JSON. Not required if `account_name` is specified.
	SecretKey *string `form:"name=secret_key"`
	// The Google Cloud Pub/Sub topic to which logs will be published. Required.
	Topic *string `form:"name=topic"`
	// Your Google Cloud Platform service account email address. The `client_email` field in your service account authentication JSON. Not required if `account_name` is specified.
	User *string `form:"name=user"`
}

func (l LoggingGooglePubsub) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingGooglePubsub) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingGooglePubsub) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *LoggingGooglePubsub) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingGooglePubsub) GetFormatVersion() *LoggingGooglePubsubFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingGooglePubsub) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingGooglePubsub) GetPlacement() *LoggingGooglePubsubPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingGooglePubsub) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *LoggingGooglePubsub) GetResponseCondition() *string {
	if o == nil {
		return nil
	}
	return o.ResponseCondition
}

func (o *LoggingGooglePubsub) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingGooglePubsub) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}

func (o *LoggingGooglePubsub) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
