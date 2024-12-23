# \PaymentsAPI

All URIs are relative to *https://brn-qa-ybus-privapi.sa.cz/affiliate*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaySroBooking**](PaymentsAPI.md#PaySroBooking) | **Put** /payments/RJ_SRO/{bookingToken}/pay | 
[**PayTicket**](PaymentsAPI.md#PayTicket) | **Put** /payments/{ticketId}/pay | 



## PaySroBooking

> SuccessResponse PaySroBooking(ctx, bookingToken).XTxToken(xTxToken).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookingToken := "bookingToken_example" // string | Internal ID of the soft booking operation that will be paid.
	xTxToken := "xTxToken_example" // string | Internal ID of the operation - used for check status 204.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaySroBooking(context.Background(), bookingToken).XTxToken(xTxToken).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaySroBooking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaySroBooking`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaySroBooking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookingToken** | **string** | Internal ID of the soft booking operation that will be paid. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaySroBookingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xTxToken** | **string** | Internal ID of the operation - used for check status 204. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayTicket

> PayTicket(ctx, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaymentsAPI.PayTicket(context.Background(), ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PayTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPayTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

