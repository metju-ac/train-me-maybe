# \ConstsAPI

All URIs are relative to *https://brn-qa-ybus-privapi.sa.cz/affiliate*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActionPrices**](ConstsAPI.md#GetActionPrices) | **Get** /consts/actionPrices | Returns list of action prices.
[**GetLanguageDictionary**](ConstsAPI.md#GetLanguageDictionary) | **Get** /consts/translations/{language} | Get translations for the selected language.
[**GetLocations**](ConstsAPI.md#GetLocations) | **Get** /consts/locations | List all locations served by RegioJet.
[**GetSeatClasses**](ConstsAPI.md#GetSeatClasses) | **Get** /consts/seatClasses | List all seat classes provided by RegioJet.
[**GetStandards**](ConstsAPI.md#GetStandards) | **Get** /consts/vehicleStandards | Get all existing vehicle standards.
[**GetTariffs**](ConstsAPI.md#GetTariffs) | **Get** /consts/tariffs | Get all possible tariffs.



## GetActionPrices

> []ActionPrice GetActionPrices(ctx).XApplicationOrigin(xApplicationOrigin).XLang(xLang).ActiveOnly(activeOnly).Execute()

Returns list of action prices.

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
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	activeOnly := true // bool | A flag that indicates whether response should contain only active action prices. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstsAPI.GetActionPrices(context.Background()).XApplicationOrigin(xApplicationOrigin).XLang(xLang).ActiveOnly(activeOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstsAPI.GetActionPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActionPrices`: []ActionPrice
	fmt.Fprintf(os.Stdout, "Response from `ConstsAPI.GetActionPrices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActionPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **activeOnly** | **bool** | A flag that indicates whether response should contain only active action prices. | [default to true]

### Return type

[**[]ActionPrice**](ActionPrice.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLanguageDictionary

> map[string]string GetLanguageDictionary(ctx, language).XApplicationOrigin(xApplicationOrigin).Execute()

Get translations for the selected language.



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
	language := "language_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Very similar to X-Lang header of other endpoints, but its value is taken from the path instead of a header.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             | 
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstsAPI.GetLanguageDictionary(context.Background(), language).XApplicationOrigin(xApplicationOrigin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstsAPI.GetLanguageDictionary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanguageDictionary`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `ConstsAPI.GetLanguageDictionary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**language** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Very similar to X-Lang header of other endpoints, but its value is taken from the path instead of a header.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguageDictionaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]

### Return type

**map[string]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocations

> []Country GetLocations(ctx).XApplicationOrigin(xApplicationOrigin).XLang(xLang).StationType(stationType).Execute()

List all locations served by RegioJet.



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
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	stationType := "stationType_example" // string | Only returns locations of the given type. One of `BUS_STATION`, or `TRAIN_STATION`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstsAPI.GetLocations(context.Background()).XApplicationOrigin(xApplicationOrigin).XLang(xLang).StationType(stationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstsAPI.GetLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocations`: []Country
	fmt.Fprintf(os.Stdout, "Response from `ConstsAPI.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **stationType** | **string** | Only returns locations of the given type. One of &#x60;BUS_STATION&#x60;, or &#x60;TRAIN_STATION&#x60;. | 

### Return type

[**[]Country**](Country.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeatClasses

> []SeatClass GetSeatClasses(ctx).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

List all seat classes provided by RegioJet.



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
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstsAPI.GetSeatClasses(context.Background()).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstsAPI.GetSeatClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeatClasses`: []SeatClass
	fmt.Fprintf(os.Stdout, "Response from `ConstsAPI.GetSeatClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeatClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**[]SeatClass**](SeatClass.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStandards

> []VehicleStandard GetStandards(ctx).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get all existing vehicle standards.



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
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstsAPI.GetStandards(context.Background()).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstsAPI.GetStandards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStandards`: []VehicleStandard
	fmt.Fprintf(os.Stdout, "Response from `ConstsAPI.GetStandards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStandardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**[]VehicleStandard**](VehicleStandard.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTariffs

> []Tariff GetTariffs(ctx).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get all possible tariffs.



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
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstsAPI.GetTariffs(context.Background()).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstsAPI.GetTariffs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTariffs`: []Tariff
	fmt.Fprintf(os.Stdout, "Response from `ConstsAPI.GetTariffs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTariffsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**[]Tariff**](Tariff.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

