# \TicketsAPI

All URIs are relative to *https://brn-qa-ybus-privapi.sa.cz/affiliate*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTicket**](TicketsAPI.md#CancelTicket) | **Put** /tickets/{accountCode}/{ticketId}/cancel | Cancel ticket by ID.
[**CreateTickets**](TicketsAPI.md#CreateTickets) | **Post** /tickets/create | Create new ticket(s).
[**GetTicketById**](TicketsAPI.md#GetTicketById) | **Get** /tickets/{accountCode}/{ticketId} | Get ticket by ID.
[**GetTicketQrCode**](TicketsAPI.md#GetTicketQrCode) | **Get** /tickets/{accountCode}/{ticketId}/qrcode | Get tickets&#39; QR code.
[**GetTicketQrCodePng**](TicketsAPI.md#GetTicketQrCodePng) | **Get** /tickets/{accountCode}/{ticketId}/qrcode/png | Get QR code image for ticket.
[**Print**](TicketsAPI.md#Print) | **Get** /tickets/{accountCode}/{ticketId}/print | Prints ticket.



## CancelTicket

> SuccessResponse CancelTicket(ctx, accountCode, ticketId).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Cancel ticket by ID.

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
	accountCode := "accountCode_example" // string | Unique code for the customer's account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint `tickets/create`. This endpoint will return `token` in response, which you should use in the step 2.   2. Authenticate user's token using endpoint `users/authenticate`, which will return `accountCode` in response. 
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	request := *openapiclient.NewCancelTicketRequest("AB42CAB42CAB42CAB42CAB42CAB42C5F") // CancelTicketRequest | 
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.CancelTicket(context.Background(), accountCode, ticketId).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.CancelTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelTicket`: SuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.CancelTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountCode** | **string** | Unique code for the customer&#39;s account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint &#x60;tickets/create&#x60;. This endpoint will return &#x60;token&#x60; in response, which you should use in the step 2.   2. Authenticate user&#39;s token using endpoint &#x60;users/authenticate&#x60;, which will return &#x60;accountCode&#x60; in response.  | 
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CreateTickets

> CreateTicketResponseUnregistered CreateTickets(ctx).TicketsRequest(ticketsRequest).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Execute()

Create new ticket(s).

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
	ticketsRequest := *openapiclient.NewCreateUnregisteredTicketRequest(false, []openapiclient.CreateTicketRequest{*openapiclient.NewCreateTicketRequest(*openapiclient.NewCreateTicketRouteRequest("3804436614,3862631913", "C1", "3870611089<3088864001-3741302011-3863407000>", []openapiclient.CreateTicketSectionRequest{*openapiclient.NewCreateTicketSectionRequest(*openapiclient.NewSimpleSection(int64(3804436614), int64(17797001), int64(10204002)))}), []openapiclient.Passenger{*openapiclient.NewPassenger("REGULAR")})}) // CreateUnregisteredTicketRequest | Ticket(s) to create.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.CreateTickets(context.Background()).TicketsRequest(ticketsRequest).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.CreateTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTickets`: CreateTicketResponseUnregistered
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.CreateTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ticketsRequest** | [**CreateUnregisteredTicketRequest**](CreateUnregisteredTicketRequest.md) | Ticket(s) to create. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]

### Return type

[**CreateTicketResponseUnregistered**](CreateTicketResponseUnregistered.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketById

> Ticket GetTicketById(ctx, accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get ticket by ID.



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
	accountCode := "accountCode_example" // string | Unique code for the customer's account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint `tickets/create`. This endpoint will return `token` in response, which you should use in the step 2.   2. Authenticate user's token using endpoint `users/authenticate`, which will return `accountCode` in response. 
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.GetTicketById(context.Background(), accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.GetTicketById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTicketById`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.GetTicketById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountCode** | **string** | Unique code for the customer&#39;s account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint &#x60;tickets/create&#x60;. This endpoint will return &#x60;token&#x60; in response, which you should use in the step 2.   2. Authenticate user&#39;s token using endpoint &#x60;users/authenticate&#x60;, which will return &#x60;accountCode&#x60; in response.  | 
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**Ticket**](Ticket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketQrCode

> QrCodeTicket GetTicketQrCode(ctx, accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get tickets' QR code.



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
	accountCode := "accountCode_example" // string | Unique code for the customer's account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint `tickets/create`. This endpoint will return `token` in response, which you should use in the step 2.   2. Authenticate user's token using endpoint `users/authenticate`, which will return `accountCode` in response. 
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.GetTicketQrCode(context.Background(), accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.GetTicketQrCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTicketQrCode`: QrCodeTicket
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.GetTicketQrCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountCode** | **string** | Unique code for the customer&#39;s account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint &#x60;tickets/create&#x60;. This endpoint will return &#x60;token&#x60; in response, which you should use in the step 2.   2. Authenticate user&#39;s token using endpoint &#x60;users/authenticate&#x60;, which will return &#x60;accountCode&#x60; in response.  | 
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketQrCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**QrCodeTicket**](QrCodeTicket.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicketQrCodePng

> *os.File GetTicketQrCodePng(ctx, accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get QR code image for ticket.



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
	accountCode := "accountCode_example" // string | Unique code for the customer's account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint `tickets/create`. This endpoint will return `token` in response, which you should use in the step 2.   2. Authenticate user's token using endpoint `users/authenticate`, which will return `accountCode` in response. 
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.GetTicketQrCodePng(context.Background(), accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.GetTicketQrCodePng``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTicketQrCodePng`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.GetTicketQrCodePng`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountCode** | **string** | Unique code for the customer&#39;s account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint &#x60;tickets/create&#x60;. This endpoint will return &#x60;token&#x60; in response, which you should use in the step 2.   2. Authenticate user&#39;s token using endpoint &#x60;users/authenticate&#x60;, which will return &#x60;accountCode&#x60; in response.  | 
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketQrCodePngRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Print

> string Print(ctx, accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Prints ticket.



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
	accountCode := "accountCode_example" // string | Unique code for the customer's account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint `tickets/create`. This endpoint will return `token` in response, which you should use in the step 2.   2. Authenticate user's token using endpoint `users/authenticate`, which will return `accountCode` in response. 
	ticketId := int64(789) // int64 | Unique identifier for the ticket.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.Print(context.Background(), accountCode, ticketId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.Print``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Print`: string
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.Print`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountCode** | **string** | Unique code for the customer&#39;s account.  To obtain account code you should go through these two steps:   1. Create new ticket using endpoint &#x60;tickets/create&#x60;. This endpoint will return &#x60;token&#x60; in response, which you should use in the step 2.   2. Authenticate user&#39;s token using endpoint &#x60;users/authenticate&#x60;, which will return &#x60;accountCode&#x60; in response.  | 
**ticketId** | **int64** | Unique identifier for the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

