// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FailureResponse - Unauthorized
//
// | Code | Message | Description |
// | ---- | ------- | ----------- |
// | 40 | Missing credentials | The <i>Authorization</i> parameter is missing. |
// | 41 | Invalid credentials | The <i>Authorization</i> parameter is not valid. |
// | 42 | Expired credentials | Renew the access token using the OAuth API and try again. |
type FailureResponse struct {
	// Represents the error code.
	Code string `json:"code"`
	// Represents the error message.
	Message string `json:"message"`
}

func (o *FailureResponse) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *FailureResponse) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
