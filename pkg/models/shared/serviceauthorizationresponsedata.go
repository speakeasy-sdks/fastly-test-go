// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ServiceAuthorizationResponseDataRelationshipsUserData struct {
	ID *string `json:"id,omitempty"`
	// Resource type
	Type *TypeUser `json:"type,omitempty"`
}

func (o *ServiceAuthorizationResponseDataRelationshipsUserData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ServiceAuthorizationResponseDataRelationshipsUserData) GetType() *TypeUser {
	if o == nil {
		return nil
	}
	return o.Type
}

// ServiceAuthorizationResponseDataRelationshipsUser - The ID of the user being given access to the service.
type ServiceAuthorizationResponseDataRelationshipsUser struct {
	Data *ServiceAuthorizationResponseDataRelationshipsUserData `json:"data,omitempty"`
}

func (o *ServiceAuthorizationResponseDataRelationshipsUser) GetData() *ServiceAuthorizationResponseDataRelationshipsUserData {
	if o == nil {
		return nil
	}
	return o.Data
}

type ServiceAuthorizationResponseDataRelationships struct {
	Service *RelationshipMemberService `json:"service,omitempty"`
	// The ID of the user being given access to the service.
	User *ServiceAuthorizationResponseDataRelationshipsUser `json:"user,omitempty"`
}

func (o *ServiceAuthorizationResponseDataRelationships) GetService() *RelationshipMemberService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ServiceAuthorizationResponseDataRelationships) GetUser() *ServiceAuthorizationResponseDataRelationshipsUser {
	if o == nil {
		return nil
	}
	return o.User
}

type ServiceAuthorizationResponseData struct {
	Attributes    *Timestamps                                    `json:"attributes,omitempty"`
	ID            *string                                        `json:"id,omitempty"`
	Relationships *ServiceAuthorizationResponseDataRelationships `json:"relationships,omitempty"`
	// Resource type
	Type *TypeServiceAuthorization `json:"type,omitempty"`
}

func (o *ServiceAuthorizationResponseData) GetAttributes() *Timestamps {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *ServiceAuthorizationResponseData) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ServiceAuthorizationResponseData) GetRelationships() *ServiceAuthorizationResponseDataRelationships {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *ServiceAuthorizationResponseData) GetType() *TypeServiceAuthorization {
	if o == nil {
		return nil
	}
	return o.Type
}
