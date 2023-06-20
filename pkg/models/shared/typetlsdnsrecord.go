// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// TypeTLSDNSRecord - Resource type
type TypeTLSDNSRecord string

const (
	TypeTLSDNSRecordDNSRecord TypeTLSDNSRecord = "dns_record"
)

func (e TypeTLSDNSRecord) ToPointer() *TypeTLSDNSRecord {
	return &e
}

func (e *TypeTLSDNSRecord) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dns_record":
		*e = TypeTLSDNSRecord(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TypeTLSDNSRecord: %v", v)
	}
}
