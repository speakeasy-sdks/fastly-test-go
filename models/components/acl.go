// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type ACL struct {
	// Name for the ACL. Must start with an alphanumeric character and contain only alphanumeric characters, underscores, and whitespace.
	Name *string `form:"name=name"`
}

func (o *ACL) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
