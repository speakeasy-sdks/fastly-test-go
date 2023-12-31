// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type BulkUpdateConfigStoreItemOp string

const (
	BulkUpdateConfigStoreItemOpCreate BulkUpdateConfigStoreItemOp = "create"
	BulkUpdateConfigStoreItemOpUpdate BulkUpdateConfigStoreItemOp = "update"
	BulkUpdateConfigStoreItemOpDelete BulkUpdateConfigStoreItemOp = "delete"
	BulkUpdateConfigStoreItemOpUpsert BulkUpdateConfigStoreItemOp = "upsert"
)

func (e BulkUpdateConfigStoreItemOp) ToPointer() *BulkUpdateConfigStoreItemOp {
	return &e
}

func (e *BulkUpdateConfigStoreItemOp) UnmarshalJSON(data []byte) error {
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
		*e = BulkUpdateConfigStoreItemOp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BulkUpdateConfigStoreItemOp: %v", v)
	}
}

type BulkUpdateConfigStoreItem struct {
	// Item key, maximum 256 characters.
	ItemKey *string `json:"item_key,omitempty"`
	// Item value, maximum 8000 characters.
	ItemValue *string                      `json:"item_value,omitempty"`
	Op        *BulkUpdateConfigStoreItemOp `json:"op,omitempty"`
}

func (o *BulkUpdateConfigStoreItem) GetItemKey() *string {
	if o == nil {
		return nil
	}
	return o.ItemKey
}

func (o *BulkUpdateConfigStoreItem) GetItemValue() *string {
	if o == nil {
		return nil
	}
	return o.ItemValue
}

func (o *BulkUpdateConfigStoreItem) GetOp() *BulkUpdateConfigStoreItemOp {
	if o == nil {
		return nil
	}
	return o.Op
}
