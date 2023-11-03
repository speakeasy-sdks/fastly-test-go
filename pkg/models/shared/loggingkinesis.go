// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"Fastly/pkg/utils"
)

type LoggingKinesis struct {
	// The access key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified.
	AccessKey *string `form:"name=access_key"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce valid JSON that Kinesis can ingest.
	Format *string `default:"{"timestamp":"%{begin:%Y-%m-%dT%H:%M:%S}t","time_elapsed":"%{time.elapsed.usec}V","is_tls":"%{if(req.is_ssl, \"true\", \"false\")}V","client_ip":"%{req.http.Fastly-Client-IP}V","geo_city":"%{client.geo.city}V","geo_country_code":"%{client.geo.country_code}V","request":"%{req.request}V","host":"%{req.http.Fastly-Orig-Host}V","url":"%{json.escape(req.url)}V","request_referer":"%{json.escape(req.http.Referer)}V","request_user_agent":"%{json.escape(req.http.User-Agent)}V","request_accept_language":"%{json.escape(req.http.Accept-Language)}V","request_accept_charset":"%{json.escape(req.http.Accept-Charset)}V","cache_status":"%{regsub(fastly_info.state, \"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\", \"\\2\\3\") }V"}" form:"name=format"`
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	//
	FormatVersion *LoggingFormatVersion `default:"2" form:"name=format_version"`
	// The ARN for an IAM role granting Fastly access to the target Amazon Kinesis stream. Not required if `access_key` and `secret_key` are provided.
	IamRole *string `form:"name=iam_role"`
	// The name for the real-time logging configuration.
	Name *string `form:"name=name"`
	// Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`.
	//
	Placement *LoggingPlacement `form:"name=placement"`
	// A named set of [AWS resources](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints) that's in the same geographical area.
	Region *AwsRegion `form:"name=region"`
	// The secret key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified.
	SecretKey *string `form:"name=secret_key"`
	// The Amazon Kinesis stream to send logs to. Required.
	Topic *string `form:"name=topic"`
}

func (l LoggingKinesis) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoggingKinesis) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LoggingKinesis) GetAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AccessKey
}

func (o *LoggingKinesis) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *LoggingKinesis) GetFormatVersion() *LoggingFormatVersion {
	if o == nil {
		return nil
	}
	return o.FormatVersion
}

func (o *LoggingKinesis) GetIamRole() *string {
	if o == nil {
		return nil
	}
	return o.IamRole
}

func (o *LoggingKinesis) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LoggingKinesis) GetPlacement() *LoggingPlacement {
	if o == nil {
		return nil
	}
	return o.Placement
}

func (o *LoggingKinesis) GetRegion() *AwsRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *LoggingKinesis) GetSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretKey
}

func (o *LoggingKinesis) GetTopic() *string {
	if o == nil {
		return nil
	}
	return o.Topic
}
