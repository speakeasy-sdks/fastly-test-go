// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type BulkUpdateDictionaryListRequest struct {
	Items []BulkUpdateDictionaryItem `json:"items,omitempty"`
}

func (o *BulkUpdateDictionaryListRequest) GetItems() []BulkUpdateDictionaryItem {
	if o == nil {
		return nil
	}
	return o.Items
}
