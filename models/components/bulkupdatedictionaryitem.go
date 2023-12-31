// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type BulkUpdateDictionaryItemOp string

const (
	BulkUpdateDictionaryItemOpCreate BulkUpdateDictionaryItemOp = "create"
	BulkUpdateDictionaryItemOpUpdate BulkUpdateDictionaryItemOp = "update"
	BulkUpdateDictionaryItemOpDelete BulkUpdateDictionaryItemOp = "delete"
	BulkUpdateDictionaryItemOpUpsert BulkUpdateDictionaryItemOp = "upsert"
)

func (e BulkUpdateDictionaryItemOp) ToPointer() *BulkUpdateDictionaryItemOp {
	return &e
}

func (e *BulkUpdateDictionaryItemOp) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "create":
		fallthrough
	case "update":
		fallthrough
	case "delete":
		fallthrough
	case "upsert":
		*e = BulkUpdateDictionaryItemOp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BulkUpdateDictionaryItemOp: %v", v)
	}
}

type BulkUpdateDictionaryItem struct {
	// Item key, maximum 256 characters.
	ItemKey *string `json:"item_key,omitempty"`
	// Item value, maximum 8000 characters.
	ItemValue *string                     `json:"item_value,omitempty"`
	Op        *BulkUpdateDictionaryItemOp `json:"op,omitempty"`
}

func (o *BulkUpdateDictionaryItem) GetItemKey() *string {
	if o == nil {
		return nil
	}
	return o.ItemKey
}

func (o *BulkUpdateDictionaryItem) GetItemValue() *string {
	if o == nil {
		return nil
	}
	return o.ItemValue
}

func (o *BulkUpdateDictionaryItem) GetOp() *BulkUpdateDictionaryItemOp {
	if o == nil {
		return nil
	}
	return o.Op
}
