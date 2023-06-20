// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// AwsRegion - A named set of [AWS resources](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints) that's in the same geographical area.
type AwsRegion string

const (
	AwsRegionUsEast1      AwsRegion = "us-east-1"
	AwsRegionUsEast2      AwsRegion = "us-east-2"
	AwsRegionUsWest1      AwsRegion = "us-west-1"
	AwsRegionUsWest2      AwsRegion = "us-west-2"
	AwsRegionAfSouth1     AwsRegion = "af-south-1"
	AwsRegionApEast1      AwsRegion = "ap-east-1"
	AwsRegionApSouth1     AwsRegion = "ap-south-1"
	AwsRegionApNortheast3 AwsRegion = "ap-northeast-3"
	AwsRegionApNortheast2 AwsRegion = "ap-northeast-2"
	AwsRegionApSoutheast1 AwsRegion = "ap-southeast-1"
	AwsRegionApSoutheast2 AwsRegion = "ap-southeast-2"
	AwsRegionApNortheast1 AwsRegion = "ap-northeast-1"
	AwsRegionCaCentral1   AwsRegion = "ca-central-1"
	AwsRegionCnNorth1     AwsRegion = "cn-north-1"
	AwsRegionCnNorthwest1 AwsRegion = "cn-northwest-1"
	AwsRegionEuCentral1   AwsRegion = "eu-central-1"
	AwsRegionEuWest1      AwsRegion = "eu-west-1"
	AwsRegionEuWest2      AwsRegion = "eu-west-2"
	AwsRegionEuSouth1     AwsRegion = "eu-south-1"
	AwsRegionEuWest3      AwsRegion = "eu-west-3"
	AwsRegionEuNorth1     AwsRegion = "eu-north-1"
	AwsRegionMeSouth1     AwsRegion = "me-south-1"
	AwsRegionSaEast1      AwsRegion = "sa-east-1"
)

func (e AwsRegion) ToPointer() *AwsRegion {
	return &e
}

func (e *AwsRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "sa-east-1":
		*e = AwsRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AwsRegion: %v", v)
	}
}
