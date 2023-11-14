// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/fastly-test-go/internal/utils"
)

// CertificateAuthority - The entity that issues and certifies the TLS certificates for your subscription.
type CertificateAuthority string

const (
	CertificateAuthorityLetsEncrypt CertificateAuthority = "lets-encrypt"
	CertificateAuthorityGlobalsign  CertificateAuthority = "globalsign"
)

func (e CertificateAuthority) ToPointer() *CertificateAuthority {
	return &e
}

func (e *CertificateAuthority) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lets-encrypt":
		fallthrough
	case "globalsign":
		*e = CertificateAuthority(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CertificateAuthority: %v", v)
	}
}

type TLSSubscriptionDataAttributes struct {
	// The entity that issues and certifies the TLS certificates for your subscription.
	CertificateAuthority *CertificateAuthority `json:"certificate_authority,omitempty"`
}

func (o *TLSSubscriptionDataAttributes) GetCertificateAuthority() *CertificateAuthority {
	if o == nil {
		return nil
	}
	return o.CertificateAuthority
}

type TLSSubscriptionData struct {
	Attributes    *TLSSubscriptionDataAttributes   `json:"attributes,omitempty"`
	Relationships *RelationshipsForTLSSubscription `json:"relationships,omitempty"`
	// Resource type
	Type *TypeTLSSubscription `default:"tls_subscription" json:"type"`
}

func (t TLSSubscriptionData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TLSSubscriptionData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TLSSubscriptionData) GetAttributes() *TLSSubscriptionDataAttributes {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *TLSSubscriptionData) GetRelationships() *RelationshipsForTLSSubscription {
	if o == nil {
		return nil
	}
	return o.Relationships
}

func (o *TLSSubscriptionData) GetType() *TypeTLSSubscription {
	if o == nil {
		return nil
	}
	return o.Type
}
