// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Security struct {
	Token string `security:"scheme,type=apiKey,subtype=header,name=Fastly-Key"`
}

func (o *Security) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}