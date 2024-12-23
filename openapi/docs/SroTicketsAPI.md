# \SroTicketsAPI

All URIs are relative to *https://brn-qa-ybus-privapi.sa.cz/affiliate*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSroTicket**](SroTicketsAPI.md#CancelSroTicket) | **Delete** /tickets/{accountCode}/RJ_SRO/{ticketId} | Cancel seat reservation only ticket by ID.
[**CancelSroTicketByTxToken**](SroTicketsAPI.md#CancelSroTicketByTxToken) | **Delete** /tickets/RJ_SRO/{txToken} | Cancel seat reservation only ticket by txToken.
[**CreateUnregisteredSroTickets**](SroTicketsAPI.md#CreateUnregisteredSroTickets) | **Post** /tickets/RJ_SRO/unregistered | Create seat reservation.
[**GetBookedSroTickets**](SroTicketsAPI.md#GetBookedSroTickets) | **Get** /tickets/RJ_SRO/{txToken} | Get seat reservation by softbooking operation ID.
[**GetSroTicketById**](SroTicketsAPI.md#GetSroTicketById) | **Get** /tickets/RJ_SRO/{accountCode}/{ticketId} | Get seat reservation by ID.



## CancelSroTicket

> SuccessResponse CancelSroTicket(ctx, accountCode, ticketId).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XTxToken(xTxToken).Execute()

Cancel seat reservation only ticket by ID.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	accountCode := "accountCode_example" // string | Unique code for the customer's account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint `tickets/create`. This endpoint will return `token` in response, which you should use in the step 2.   2. Authenticate user's token using endpoint `users/authenticate`, which will return `accountCode` in response. 
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	request := *openapiclient.NewCancelTicketRequest("AB42CAB42CAB42CAB42CAB42CAB42C5F") // CancelTicketRequest | 
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xTxToken := "xTxToken_example" // string | Token (hash) identifier of transaction. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SroTicketsAPI.CancelSroTicket(context.Background(), accountCode, ticketId).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XTxToken(xTxToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SroTicketsAPI.CancelSroTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSroTicket`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `SroTicketsAPI.CancelSroTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountCode** | **string** | Unique code for the customer&#39;s account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint &#x60;tickets/create&#x60;. This endpoint will return &#x60;token&#x60; in response, which you should use in the step 2.   2. Authenticate user&#39;s token using endpoint &#x60;users/authenticate&#x60;, which will return &#x60;accountCode&#x60; in response.  | 
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSroTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CancelTicketRequest**](CancelTicketRequest.md) |  | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xTxToken** | **string** | Token (hash) identifier of transaction. | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelSroTicketByTxToken

> SuccessResponse CancelSroTicketByTxToken(ctx, txToken).XTxToken(xTxToken).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Cancel seat reservation only ticket by txToken.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	xTxToken := "xTxToken_example" // string | Internal ID of the operation - used for check status 204.
	txToken := "txToken_example" // string | Internal ID of the original booking operation that will be canceled.
	request := *openapiclient.NewCancelTicketRequest("AB42CAB42CAB42CAB42CAB42CAB42C5F") // CancelTicketRequest | 
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SroTicketsAPI.CancelSroTicketByTxToken(context.Background(), txToken).XTxToken(xTxToken).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SroTicketsAPI.CancelSroTicketByTxToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSroTicketByTxToken`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `SroTicketsAPI.CancelSroTicketByTxToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txToken** | **string** | Internal ID of the original booking operation that will be canceled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSroTicketByTxTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTxToken** | **string** | Internal ID of the operation - used for check status 204. | 

 **request** | [**CancelTicketRequest**](CancelTicketRequest.md) |  | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUnregisteredSroTickets

> Object CreateUnregisteredSroTickets(ctx).XTxToken(xTxToken).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Execute()

Create seat reservation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	xTxToken := "xTxToken_example" // string | Internal ID of the operation - used for check status 204.
	request := *openapiclient.NewCreateSroTicketsRequest(true, []openapiclient.SroTicketRequest{*openapiclient.NewSroTicketRequest("3804436614", "PHN-BHN_ONLINE_RJ_SRO_449356704_C0_CZK_NORMAL_35_1084776538", "C0", []openapiclient.SroTicketSections{*openapiclient.NewSroTicketSections(*openapiclient.NewSroTicketSection(int64(3804436614), int64(3804436614), int64(10204002)), []openapiclient.SroTicketSelectedSeat{*openapiclient.NewSroTicketSelectedSeat(int64(3804436614), int32(7), int32(15))})}, []openapiclient.SroTicketPassenger{*openapiclient.NewSroTicketPassenger(int64(3895574004), "First", "Passenger", "420777666777", "first.passenger@email.com", time.Now(), "Address_example", "45678877", "100", "90")})}, "ExternalId_example") // CreateSroTicketsRequest | Object with seat reservations that need to be created.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SroTicketsAPI.CreateUnregisteredSroTickets(context.Background()).XTxToken(xTxToken).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SroTicketsAPI.CreateUnregisteredSroTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUnregisteredSroTickets`: Object
	fmt.Fprintf(os.Stdout, "Response from `SroTicketsAPI.CreateUnregisteredSroTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUnregisteredSroTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTxToken** | **string** | Internal ID of the operation - used for check status 204. | 
 **request** | [**CreateSroTicketsRequest**](CreateSroTicketsRequest.md) | Object with seat reservations that need to be created. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]

### Return type

[**Object**](Object.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookedSroTickets

> []SroTicketBookingResponse GetBookedSroTickets(ctx, txToken).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get seat reservation by softbooking operation ID.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	txToken := "txToken_example" // string | Internal ID of the soft booking operation that will be canceled.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SroTicketsAPI.GetBookedSroTickets(context.Background(), txToken).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SroTicketsAPI.GetBookedSroTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookedSroTickets`: []SroTicketBookingResponse
	fmt.Fprintf(os.Stdout, "Response from `SroTicketsAPI.GetBookedSroTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**txToken** | **string** | Internal ID of the soft booking operation that will be canceled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookedSroTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**[]SroTicketBookingResponse**](SroTicketBookingResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSroTicketById

> SroTicket GetSroTicketById(ctx, accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get seat reservation by ID.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
)

func main() {
	accountCode := "accountCode_example" // string | Unique code for the customer's account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint `tickets/create`. This endpoint will return `token` in response, which you should use in the step 2.   2. Authenticate user's token using endpoint `users/authenticate`, which will return `accountCode` in response. 
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SroTicketsAPI.GetSroTicketById(context.Background(), accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SroTicketsAPI.GetSroTicketById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSroTicketById`: SroTicket
	fmt.Fprintf(os.Stdout, "Response from `SroTicketsAPI.GetSroTicketById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountCode** | **string** | Unique code for the customer&#39;s account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint &#x60;tickets/create&#x60;. This endpoint will return &#x60;token&#x60; in response, which you should use in the step 2.   2. Authenticate user&#39;s token using endpoint &#x60;users/authenticate&#x60;, which will return &#x60;accountCode&#x60; in response.  | 
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSroTicketByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**SroTicket**](SroTicket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

