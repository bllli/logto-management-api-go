// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/bllli/logto-management-api/models/components"
)

// GenerateBackupCodesResponseBody - Backup codes have been successfully generated.
type GenerateBackupCodesResponseBody struct {
	// The unique verification ID of the newly created BackupCode verification record. This ID is required when adding the backup codes to the user profile via the Profile API.
	VerificationID string `json:"verificationId"`
	// The generated backup codes.
	Codes []string `json:"codes"`
}

func (o *GenerateBackupCodesResponseBody) GetVerificationID() string {
	if o == nil {
		return ""
	}
	return o.VerificationID
}

func (o *GenerateBackupCodesResponseBody) GetCodes() []string {
	if o == nil {
		return []string{}
	}
	return o.Codes
}

type GenerateBackupCodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Backup codes have been successfully generated.
	Object *GenerateBackupCodesResponseBody
}

func (o *GenerateBackupCodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GenerateBackupCodesResponse) GetObject() *GenerateBackupCodesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
