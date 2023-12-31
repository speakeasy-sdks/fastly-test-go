// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// BillingRegion - the region used for billing
type BillingRegion string

const (
	BillingRegionAfrica       BillingRegion = "Africa"
	BillingRegionAustralia    BillingRegion = "Australia"
	BillingRegionAsia         BillingRegion = "Asia"
	BillingRegionEurope       BillingRegion = "Europe"
	BillingRegionIndia        BillingRegion = "India"
	BillingRegionNorthAmerica BillingRegion = "North America"
	BillingRegionSouthKorea   BillingRegion = "South Korea"
	BillingRegionSouthAmerica BillingRegion = "South America"
)

func (e BillingRegion) ToPointer() *BillingRegion {
	return &e
}

func (e *BillingRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Africa":
		fallthrough
	case "Australia":
		fallthrough
	case "Asia":
		fallthrough
	case "Europe":
		fallthrough
	case "India":
		fallthrough
	case "North America":
		fallthrough
	case "South Korea":
		fallthrough
	case "South America":
		*e = BillingRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillingRegion: %v", v)
	}
}

// Coordinates - the geographic location of the POP
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (o *Coordinates) GetLatitude() float64 {
	if o == nil {
		return 0.0
	}
	return o.Latitude
}

func (o *Coordinates) GetLongitude() float64 {
	if o == nil {
		return 0.0
	}
	return o.Longitude
}

type PopRegion string

const (
	PopRegionApac         PopRegion = "APAC"
	PopRegionAsia         PopRegion = "Asia"
	PopRegionAfWest       PopRegion = "AF-West"
	PopRegionEuCentral    PopRegion = "EU-Central"
	PopRegionEuEast       PopRegion = "EU-East"
	PopRegionEuWest       PopRegion = "EU-West"
	PopRegionMiddleEast   PopRegion = "Middle-East"
	PopRegionNorthAmerica PopRegion = "North-America"
	PopRegionSaSouth      PopRegion = "SA-South"
	PopRegionSaEast       PopRegion = "SA-East"
	PopRegionSaWest       PopRegion = "SA-West"
	PopRegionSaNorth      PopRegion = "SA-North"
	PopRegionSouthAfrica  PopRegion = "South-Africa"
	PopRegionSouthAmerica PopRegion = "South-America"
	PopRegionUsCentral    PopRegion = "US-Central"
	PopRegionUsEast       PopRegion = "US-East"
	PopRegionUsWest       PopRegion = "US-West"
	PopRegionAsiaSouth    PopRegion = "Asia-South"
)

func (e PopRegion) ToPointer() *PopRegion {
	return &e
}

func (e *PopRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "APAC":
		fallthrough
	case "Asia":
		fallthrough
	case "AF-West":
		fallthrough
	case "EU-Central":
		fallthrough
	case "EU-East":
		fallthrough
	case "EU-West":
		fallthrough
	case "Middle-East":
		fallthrough
	case "North-America":
		fallthrough
	case "SA-South":
		fallthrough
	case "SA-East":
		fallthrough
	case "SA-West":
		fallthrough
	case "SA-North":
		fallthrough
	case "South-Africa":
		fallthrough
	case "South-America":
		fallthrough
	case "US-Central":
		fallthrough
	case "US-East":
		fallthrough
	case "US-West":
		fallthrough
	case "Asia-South":
		*e = PopRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PopRegion: %v", v)
	}
}

// StatsRegion - the region used for stats reporting
type StatsRegion string

const (
	StatsRegionSouthamericaStd StatsRegion = "southamerica_std"
	StatsRegionAfricaStd       StatsRegion = "africa_std"
	StatsRegionAnzac           StatsRegion = "anzac"
	StatsRegionAsia            StatsRegion = "asia"
	StatsRegionEurope          StatsRegion = "europe"
	StatsRegionUsa             StatsRegion = "usa"
	StatsRegionAsiaIndia       StatsRegion = "asia_india"
	StatsRegionAsiaSouthkorea  StatsRegion = "asia_southkorea"
)

func (e StatsRegion) ToPointer() *StatsRegion {
	return &e
}

func (e *StatsRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "southamerica_std":
		fallthrough
	case "africa_std":
		fallthrough
	case "anzac":
		fallthrough
	case "asia":
		fallthrough
	case "europe":
		fallthrough
	case "usa":
		fallthrough
	case "asia_india":
		fallthrough
	case "asia_southkorea":
		*e = StatsRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StatsRegion: %v", v)
	}
}

type Pop struct {
	// the region used for billing
	BillingRegion BillingRegion `json:"billing_region"`
	// the three-letter code for the [POP](https://developer.fastly.com/learning/concepts/pop/)
	Code string `json:"code"`
	// the geographic location of the POP
	Coordinates *Coordinates `json:"coordinates,omitempty"`
	Group       string       `json:"group"`
	// the name of the POP
	Name   string    `json:"name"`
	Region PopRegion `json:"region"`
	// the name of the [shield code](https://developer.fastly.com/learning/concepts/shielding/#choosing-a-shield-location) if this POP is suitable for shielding
	Shield *string `json:"shield,omitempty"`
	// the region used for stats reporting
	StatsRegion StatsRegion `json:"stats_region"`
}

func (o *Pop) GetBillingRegion() BillingRegion {
	if o == nil {
		return BillingRegion("")
	}
	return o.BillingRegion
}

func (o *Pop) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *Pop) GetCoordinates() *Coordinates {
	if o == nil {
		return nil
	}
	return o.Coordinates
}

func (o *Pop) GetGroup() string {
	if o == nil {
		return ""
	}
	return o.Group
}

func (o *Pop) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Pop) GetRegion() PopRegion {
	if o == nil {
		return PopRegion("")
	}
	return o.Region
}

func (o *Pop) GetShield() *string {
	if o == nil {
		return nil
	}
	return o.Shield
}

func (o *Pop) GetStatsRegion() StatsRegion {
	if o == nil {
		return StatsRegion("")
	}
	return o.StatsRegion
}
