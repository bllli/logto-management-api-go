// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/bllli/logto-management-api/models/components"
)

// UpdateInteractionEventInteractionEvent - The type of the interaction event. Only `SignIn` and `Register` are supported.
type UpdateInteractionEventInteractionEvent string

const (
	UpdateInteractionEventInteractionEventSignIn         UpdateInteractionEventInteractionEvent = "SignIn"
	UpdateInteractionEventInteractionEventRegister       UpdateInteractionEventInteractionEvent = "Register"
	UpdateInteractionEventInteractionEventForgotPassword UpdateInteractionEventInteractionEvent = "ForgotPassword"
)

func (e UpdateInteractionEventInteractionEvent) ToPointer() *UpdateInteractionEventInteractionEvent {
	return &e
}
func (e *UpdateInteractionEventInteractionEvent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SignIn":
		fallthrough
	case "Register":
		fallthrough
	case "ForgotPassword":
		*e = UpdateInteractionEventInteractionEvent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateInteractionEventInteractionEvent: %v", v)
	}
}

type UpdateInteractionEventRequest struct {
	// The type of the interaction event. Only `SignIn` and `Register` are supported.
	InteractionEvent UpdateInteractionEventInteractionEvent `json:"interactionEvent"`
}

func (o *UpdateInteractionEventRequest) GetInteractionEvent() UpdateInteractionEventInteractionEvent {
	if o == nil {
		return UpdateInteractionEventInteractionEvent("")
	}
	return o.InteractionEvent
}

type UpdateInteractionEventResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateInteractionEventResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
