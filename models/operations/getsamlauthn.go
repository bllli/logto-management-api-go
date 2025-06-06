// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bllli/logto-management-api/models/components"
)

type GetSamlAuthnRequest struct {
	// The ID of the SAML application.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The SAML request message.
	SAMLRequest string `queryParam:"style=form,explode=true,name=SAMLRequest"`
	// The signature of the request.
	Signature *string `queryParam:"style=form,explode=true,name=Signature"`
	// The signature algorithm.
	SigAlg *string `queryParam:"style=form,explode=true,name=SigAlg"`
	// The relay state parameter.
	RelayState *string `queryParam:"style=form,explode=true,name=RelayState"`
}

func (o *GetSamlAuthnRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetSamlAuthnRequest) GetSAMLRequest() string {
	if o == nil {
		return ""
	}
	return o.SAMLRequest
}

func (o *GetSamlAuthnRequest) GetSignature() *string {
	if o == nil {
		return nil
	}
	return o.Signature
}

func (o *GetSamlAuthnRequest) GetSigAlg() *string {
	if o == nil {
		return nil
	}
	return o.SigAlg
}

func (o *GetSamlAuthnRequest) GetRelayState() *string {
	if o == nil {
		return nil
	}
	return o.RelayState
}

type GetSamlAuthnResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetSamlAuthnResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
