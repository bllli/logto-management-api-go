// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bllli/logto-management-api/models/components"
)

type DeleteSamlApplicationRequest struct {
	// The unique identifier of the saml application.
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteSamlApplicationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteSamlApplicationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteSamlApplicationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
