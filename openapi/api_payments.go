/*
RegioJet's Affiliate API Reference

The RegioJet\\'s Affiliate API is a set of endpoints that help your application integrate with RegioJet.  The API is organized arount [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API uses standard HTTP methods, authentication, and status codes.  # Authentication Authentication to the API is performed via [HTTP Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication) for all endpoints listed in this documentation with the exception of `/users/authenticate`, which uses bearer token.  API requests without authentication will fail.  All API requests must be made over [HTTPS](https://en.wikipedia.org/wiki/HTTPS).  # Errors  RegioJet uses conventional HTTP status codes in responses to indicate the success or failure of an API request.  In general:   * `2xx` codes indicate success;   * `4xx` codes indicate an error that failed given the information provided in request.   * `5xx` codes indicate an error with RegioJet's servers.

API version: 1.1.0
Contact: developers@studentagency.cz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// PaymentsAPIService PaymentsAPI service
type PaymentsAPIService service

type ApiPaySroBookingRequest struct {
	ctx                context.Context
	ApiService         *PaymentsAPIService
	bookingToken       string
	xTxToken           *string
	xApplicationOrigin *string
	xLang              *string
}

// Internal ID of the operation - used for check status 204.
func (r ApiPaySroBookingRequest) XTxToken(xTxToken string) ApiPaySroBookingRequest {
	r.xTxToken = &xTxToken
	return r
}

// Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type
func (r ApiPaySroBookingRequest) XApplicationOrigin(xApplicationOrigin string) ApiPaySroBookingRequest {
	r.xApplicationOrigin = &xApplicationOrigin
	return r
}

// A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |
func (r ApiPaySroBookingRequest) XLang(xLang string) ApiPaySroBookingRequest {
	r.xLang = &xLang
	return r
}

func (r ApiPaySroBookingRequest) Execute() (*SuccessResponse, *http.Response, error) {
	return r.ApiService.PaySroBookingExecute(r)
}

/*
PaySroBooking Method for PaySroBooking

Marks a tickets as paid for given booking token.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bookingToken Internal ID of the soft booking operation that will be paid.
	@return ApiPaySroBookingRequest
*/
func (a *PaymentsAPIService) PaySroBooking(ctx context.Context, bookingToken string) ApiPaySroBookingRequest {
	return ApiPaySroBookingRequest{
		ApiService:   a,
		ctx:          ctx,
		bookingToken: bookingToken,
	}
}

// Execute executes the request
//
//	@return SuccessResponse
func (a *PaymentsAPIService) PaySroBookingExecute(r ApiPaySroBookingRequest) (*SuccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SuccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsAPIService.PaySroBooking")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payments/RJ_SRO/{bookingToken}/pay"
	localVarPath = strings.Replace(localVarPath, "{"+"bookingToken"+"}", url.PathEscape(parameterValueToString(r.bookingToken, "bookingToken")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.bookingToken) < 8 {
		return localVarReturnValue, nil, reportError("bookingToken must have at least 8 elements")
	}
	if strlen(r.bookingToken) > 10 {
		return localVarReturnValue, nil, reportError("bookingToken must have less than 10 elements")
	}
	if r.xTxToken == nil {
		return localVarReturnValue, nil, reportError("xTxToken is required and must be specified")
	}
	if strlen(*r.xTxToken) < 8 {
		return localVarReturnValue, nil, reportError("xTxToken must have at least 8 elements")
	}
	if strlen(*r.xTxToken) > 10 {
		return localVarReturnValue, nil, reportError("xTxToken must have less than 10 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xApplicationOrigin != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Application-Origin", r.xApplicationOrigin, "", "")
	}
	if r.xLang != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Lang", r.xLang, "", "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-TxToken", r.xTxToken, "", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v BadRequestResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPayTicketRequest struct {
	ctx                context.Context
	ApiService         *PaymentsAPIService
	ticketId           int64
	xApplicationOrigin *string
	xLang              *string
}

// Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type
func (r ApiPayTicketRequest) XApplicationOrigin(xApplicationOrigin string) ApiPayTicketRequest {
	r.xApplicationOrigin = &xApplicationOrigin
	return r
}

// A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |
func (r ApiPayTicketRequest) XLang(xLang string) ApiPayTicketRequest {
	r.xLang = &xLang
	return r
}

func (r ApiPayTicketRequest) Execute() (*http.Response, error) {
	return r.ApiService.PayTicketExecute(r)
}

/*
PayTicket Method for PayTicket

Pays a ticket.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ticketId Unique identifier for the ticket.
	@return ApiPayTicketRequest
*/
func (a *PaymentsAPIService) PayTicket(ctx context.Context, ticketId int64) ApiPayTicketRequest {
	return ApiPayTicketRequest{
		ApiService: a,
		ctx:        ctx,
		ticketId:   ticketId,
	}
}

// Execute executes the request
func (a *PaymentsAPIService) PayTicketExecute(r ApiPayTicketRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentsAPIService.PayTicket")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payments/{ticketId}/pay"
	localVarPath = strings.Replace(localVarPath, "{"+"ticketId"+"}", url.PathEscape(parameterValueToString(r.ticketId, "ticketId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xApplicationOrigin != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Application-Origin", r.xApplicationOrigin, "", "")
	}
	if r.xLang != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Lang", r.xLang, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse4001
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
