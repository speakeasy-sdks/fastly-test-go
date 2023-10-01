// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RelationshipTLSConfigurationsTLSConfigurationsInput struct {
	Data []RelationshipMemberTLSConfigurationInput `json:"data,omitempty"`
}

func (o *RelationshipTLSConfigurationsTLSConfigurationsInput) GetData() []RelationshipMemberTLSConfigurationInput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipTLSConfigurationsInput struct {
	TLSConfigurations *RelationshipTLSConfigurationsTLSConfigurationsInput `json:"tls_configurations,omitempty"`
}

func (o *RelationshipTLSConfigurationsInput) GetTLSConfigurations() *RelationshipTLSConfigurationsTLSConfigurationsInput {
	if o == nil {
		return nil
	}
	return o.TLSConfigurations
}

type RelationshipTLSConfigurationsTLSConfigurations struct {
	Data []RelationshipMemberTLSConfiguration `json:"data,omitempty"`
}

func (o *RelationshipTLSConfigurationsTLSConfigurations) GetData() []RelationshipMemberTLSConfiguration {
	if o == nil {
		return nil
	}
	return o.Data
}

type RelationshipTLSConfigurations struct {
	TLSConfigurations *RelationshipTLSConfigurationsTLSConfigurations `json:"tls_configurations,omitempty"`
}

func (o *RelationshipTLSConfigurations) GetTLSConfigurations() *RelationshipTLSConfigurationsTLSConfigurations {
	if o == nil {
		return nil
	}
	return o.TLSConfigurations
}
