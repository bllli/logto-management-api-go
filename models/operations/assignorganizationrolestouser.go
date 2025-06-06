// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bllli/logto-management-api/models/components"
)

type AssignOrganizationRolesToUserRequestBody struct {
	// An array of organization role IDs to assign to the user. User existed roles assignment will be ignored.
	OrganizationRoleIds []string `json:"organizationRoleIds,omitempty"`
	// An array of organization role names to assign to the user. User existed roles assignment will be ignored.
	OrganizationRoleNames []string `json:"organizationRoleNames,omitempty"`
}

func (o *AssignOrganizationRolesToUserRequestBody) GetOrganizationRoleIds() []string {
	if o == nil {
		return nil
	}
	return o.OrganizationRoleIds
}

func (o *AssignOrganizationRolesToUserRequestBody) GetOrganizationRoleNames() []string {
	if o == nil {
		return nil
	}
	return o.OrganizationRoleNames
}

type AssignOrganizationRolesToUserRequest struct {
	// The unique identifier of the organization.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The unique identifier of the user.
	UserID      string                                   `pathParam:"style=simple,explode=false,name=userId"`
	RequestBody AssignOrganizationRolesToUserRequestBody `request:"mediaType=application/json"`
}

func (o *AssignOrganizationRolesToUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AssignOrganizationRolesToUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *AssignOrganizationRolesToUserRequest) GetRequestBody() AssignOrganizationRolesToUserRequestBody {
	if o == nil {
		return AssignOrganizationRolesToUserRequestBody{}
	}
	return o.RequestBody
}

type AssignOrganizationRolesToUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *AssignOrganizationRolesToUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
