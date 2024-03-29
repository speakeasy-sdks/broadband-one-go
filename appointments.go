// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package broadbandone

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"openapi/pkg/models/operations"
	"openapi/pkg/models/sdkerrors"
	"openapi/pkg/models/shared"
	"openapi/pkg/utils"
	"strings"
)

type appointments struct {
	sdkConfiguration sdkConfiguration
}

func newAppointments(sdkConfig sdkConfiguration) *appointments {
	return &appointments{
		sdkConfiguration: sdkConfig,
	}
}

// Schedule - Create appointment
// <p>Reserves an engineering appointment and returns an appointment ID for use in subsequent API calls.</p>      	 <p>This API supports TLS mutual authentication.</p>
//
//	<details>
//	<summary><b>Sandbox responses</b></summary>
//
//	 <p>The sandbox supports a number of success and failure scenarios. Use the parameters below in the sandbox API call to generate the corresponding responses.</p>
//	 <p>Success scenarios responses are simulated based on the parameters mentioned in the table.</p>
//
//	 <p>Back-end error scenarios responses are simulated based on APIGW-Tracking-Header.</p>
//
//	 <details>
//	 <summary><b>Success scenarios</b></summary>
//
//	 | Parameters| Sandbox response invoked | Description |
//	 | ----------- | ------------------- | ------------------ |
//	 | <i>name</i>='DirectoryNumber', <i>value</i>='2277264355'. | 201 | Retrieve product appointments based on 'DirectoryNumber'. |
//	 | <i>id</i>='R0123457890'. | 201 | Retrieve get product appointments based on product place ID. |
//
//
//	 </details>
//
//	 <details>
//	 <summary><b>Error scenarios</b></summary>
//
//	 | APIGW-Tracking-Header | Sandbox response invoked | Description |
//	 | ----------- | ------------ | ------------ |
//	 | E40001 | 400, TimeSlotNotSupported | The product does not support appointments in that slot. |
//	 | E40301 | 403, ExchangeLineAssignedToAnotherCP | The exchange line is in use　by another communication provider. |
//
//
//	 </details>
//	 </details>
func (s *appointments) Schedule(ctx context.Context, request operations.CreateAppointmentRequest, security operations.CreateAppointmentSecurity) (*operations.CreateAppointmentResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/appointmentManagement/v4/appointment"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "PostAppointmentRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := utils.ConfigureSecurityClient(s.sdkConfiguration.DefaultClient, security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateAppointmentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PostAppointmentResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostAppointmentResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.CreateAppointment400ApplicationJSONAnyOf = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 405:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.FailureResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.FailureResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.CreateAppointment403ApplicationJSONAnyOf = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 415:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.CreateAppointment415ApplicationJSONAnyOf = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// Search - Appointment availability
// <p>Find available engineer appointment slots based on a location, product, type of work and 1st available date.</p>
//
// <p>This API supports TLS mutual authentication.</p>
//
//	<details>
//	<summary><b>Sandbox responses</b></summary>
//
//	<p>The sandbox supports a number of success and failure scenarios. Use the parameters below in the sandbox API call to generate the corresponding responses.</p>
//	<p>Success scenarios responses are simulated based on the parameters mentioned in the table.</p>
//
//	<p>Back-end error scenarios responses are simulated based on APIGW-Tracking-Header.</p>
//
//	<details>
//	<summary><b>Success scenarios</b></summary>
//
//	| Parameters| Sandbox response invoked | Description |
//	| ----------- | ------------------- | ------------------ |
//	| <i>name</i>='DirectoryNumber', <i>value</i>='2277264355'.| 201 | Retrieve product appointments based on 'DirectoryNumber'. |
//	| <i>id</i>='R0123457890'. | 201 | Retrieve product appointments based on product place ID. |
//
//
//	</details>
//
//	<details>
//	<summary><b>Error scenarios</b></summary>
//
//	| APIGW-Tracking-Header | Sandbox response invoked | Description |
//	| ----------- | ------------ | ------------ |
//	| E40001 | 400, TimeSlotNotSupported | The product does not support appointments in that slot. |
//	| E40301 | 403, ExchangeLineAssignedToAnotherCP | The exchange line is in use by another communication provider. |
//
//
//	</details>
//	</details>
func (s *appointments) Search(ctx context.Context, request operations.PostSearchAppointmentRequest) (*operations.PostSearchAppointmentResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/appointmentManagement/v4/searchTimeSlot"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "PostSearchAppointmentRequest", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostSearchAppointmentResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PostSearchAppointmentResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostSearchAppointmentResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostSearchAppointment400ApplicationJSONAnyOf = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 405:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode == 500:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.FailureResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.FailureResponse = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 403:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostSearchAppointment403ApplicationJSONAnyOf = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 415:
		res.Headers = httpRes.Header

		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PostSearchAppointment415ApplicationJSONAnyOf = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
