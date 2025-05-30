// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/bllli/logto-management-api/models/components"
)

type InitInteractionInteractionEvent string

const (
	InitInteractionInteractionEventSignIn         InitInteractionInteractionEvent = "SignIn"
	InitInteractionInteractionEventRegister       InitInteractionInteractionEvent = "Register"
	InitInteractionInteractionEventForgotPassword InitInteractionInteractionEvent = "ForgotPassword"
)

func (e InitInteractionInteractionEvent) ToPointer() *InitInteractionInteractionEvent {
	return &e
}
func (e *InitInteractionInteractionEvent) UnmarshalJSON(data []byte) error {
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
		*e = InitInteractionInteractionEvent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InitInteractionInteractionEvent: %v", v)
	}
}

type InitInteractionRequest struct {
	InteractionEvent InitInteractionInteractionEvent `json:"interactionEvent"`
	CaptchaToken     *string                         `json:"captchaToken,omitempty"`
}

func (o *InitInteractionRequest) GetInteractionEvent() InitInteractionInteractionEvent {
	if o == nil {
		return InitInteractionInteractionEvent("")
	}
	return o.InteractionEvent
}

func (o *InitInteractionRequest) GetCaptchaToken() *string {
	if o == nil {
		return nil
	}
	return o.CaptchaToken
}

type InitInteractionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *InitInteractionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
