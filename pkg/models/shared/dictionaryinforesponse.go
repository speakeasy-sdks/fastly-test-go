// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DictionaryInfoResponse - OK
type DictionaryInfoResponse struct {
	// A hash of all the dictionary content.
	Digest *string `json:"digest,omitempty"`
	// The number of items currently in the dictionary.
	ItemCount *int64 `json:"item_count,omitempty"`
	// Timestamp (UTC) when the dictionary was last updated or an item was added or removed.
	LastUpdated *string `json:"last_updated,omitempty"`
}