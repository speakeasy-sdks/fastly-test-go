// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConfigStoreItem struct {
	// Item key, maximum 256 characters.
	ItemKey *string `form:"name=item_key"`
	// Item value, maximum 8000 characters.
	ItemValue *string `form:"name=item_value"`
}

func (o *ConfigStoreItem) GetItemKey() *string {
	if o == nil {
		return nil
	}
	return o.ItemKey
}

func (o *ConfigStoreItem) GetItemValue() *string {
	if o == nil {
		return nil
	}
	return o.ItemValue
}
