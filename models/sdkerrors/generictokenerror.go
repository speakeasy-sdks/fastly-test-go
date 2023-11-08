// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type GenericTokenError struct {
	Msg *string `json:"msg,omitempty"`
}

var _ error = &GenericTokenError{}

func (e *GenericTokenError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
