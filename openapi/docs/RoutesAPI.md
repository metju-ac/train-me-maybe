# \RoutesAPI

All URIs are relative to *https://brn-qa-ybus-privapi.sa.cz/affiliate*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPassengersData**](RoutesAPI.md#GetPassengersData) | **Post** /routes/{routeId}/passengersData | Get mandatory passengers data.
[**GetRouteFreeSeats**](RoutesAPI.md#GetRouteFreeSeats) | **Post** /routes/freeSeats | Get route tandem free seats grouped by vehicles.
[**GetRouteFreeSeats110**](RoutesAPI.md#GetRouteFreeSeats110) | **Post** /routes/{routeId}/freeSeats | Get list of free seats.
[**GetSimpleRouteDetail**](RoutesAPI.md#GetSimpleRouteDetail) | **Get** /routes/{routeId}/simple | Get details of the given route.
[**GetSroPassengersData**](RoutesAPI.md#GetSroPassengersData) | **Get** /routes/{routeId}/passengersData/RJ_SRO | Get mandatory passengers data.
[**GetSroRoutePrices**](RoutesAPI.md#GetSroRoutePrices) | **Get** /routes/{routeId}/prices/RJ_SRO | Get route detail.
[**SearchRoutes**](RoutesAPI.md#SearchRoutes) | **Get** /routes/search | Get collection of all routes that satisfy specified search criteria for the route.
[**SearchSR70Routes**](RoutesAPI.md#SearchSR70Routes) | **Get** /routes/{connection}/{fromLocation}/{toLocation}/{ticketType} | Search for available routes.
[**SimpleSearchRoutes**](RoutesAPI.md#SimpleSearchRoutes) | **Get** /routes/search/simple | Search for routes that satisfy specified search criteria.



## GetPassengersData

> PassengersDataResponse GetPassengersData(ctx, routeId).Filter(filter).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Execute()

Get mandatory passengers data.



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
	routeId := "routeId_example" // string | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. `section0.id,section1.id,section2.id`).
	filter := *openapiclient.NewPassengersDataRequest([]openapiclient.SimpleSection{*openapiclient.NewSimpleSection(int64(3804436614), int64(17797001), int64(10204002))}, []string{"REGULAR"}, "C1", "3870611089<3088864001-3741302011-3863407000>") // PassengersDataRequest | 
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetPassengersData(context.Background(), routeId).Filter(filter).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetPassengersData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPassengersData`: PassengersDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetPassengersData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. &#x60;section0.id,section1.id,section2.id&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPassengersDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**PassengersDataRequest**](PassengersDataRequest.md) |  | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]

### Return type

[**PassengersDataResponse**](PassengersDataResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFreeSeats

> RouteSeatsResponse GetRouteFreeSeats(ctx).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XOccupied(xOccupied).Execute()

Get route tandem free seats grouped by vehicles.

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
	request := *openapiclient.NewRouteSeatsRequest([]openapiclient.SimpleSection{*openapiclient.NewSimpleSection(int64(3804436614), int64(17797001), int64(10204002))}, []string{"REGULAR"}, "C1") // RouteSeatsRequest | Descriptions of the Route detail
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xOccupied := true // bool | Return occupied seats. To turn off have to be deactivated at administrative app and X-occupied have to be false. Default value is false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRouteFreeSeats(context.Background()).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XOccupied(xOccupied).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRouteFreeSeats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFreeSeats`: RouteSeatsResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRouteFreeSeats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFreeSeatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**RouteSeatsRequest**](RouteSeatsRequest.md) | Descriptions of the Route detail | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xOccupied** | **bool** | Return occupied seats. To turn off have to be deactivated at administrative app and X-occupied have to be false. Default value is false. | [default to false]

### Return type

[**RouteSeatsResponse**](RouteSeatsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteFreeSeats110

> RouteSeatsResponse110 GetRouteFreeSeats110(ctx, routeId).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()

Get list of free seats.



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
	routeId := "routeId_example" // string | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. `section0.id,section1.id,section2.id`).
	request := *openapiclient.NewRouteSeatsRequest([]openapiclient.SimpleSection{*openapiclient.NewSimpleSection(int64(3804436614), int64(17797001), int64(10204002))}, []string{"REGULAR"}, "C1") // RouteSeatsRequest | Route section details.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetRouteFreeSeats110(context.Background(), routeId).Request(request).XApplicationOrigin(xApplicationOrigin).XLang(xLang).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetRouteFreeSeats110``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteFreeSeats110`: RouteSeatsResponse110
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetRouteFreeSeats110`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. &#x60;section0.id,section1.id,section2.id&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteFreeSeats110Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**RouteSeatsRequest**](RouteSeatsRequest.md) | Route section details. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]

### Return type

[**RouteSeatsResponse110**](RouteSeatsResponse110.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimpleRouteDetail

> Route GetSimpleRouteDetail(ctx, routeId).FromStationId(fromStationId).ToStationId(toStationId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Tariffs(tariffs).Execute()

Get details of the given route.



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
	routeId := "routeId_example" // string | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. `section0.id,section1.id,section2.id`).
	fromStationId := int64(789) // int64 | Unique identifier for `from` station. Detailed data for identifier can be obtained from `/consts/locations` endpoint.
	toStationId := int64(789) // int64 | Unique identifier for `from` station. Detailed data for identifier can be obtained from `/consts/locations` endpoint.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")
	tariffs := []string{"Inner_example"} // []string | List of tariffs. Tariff keys can be obtained from `/consts/tariffs`. Multiple parameter instances should be used rather than multiple values when submitting multiple values. For example: `&tariffs=REGULAR&tariffs=ISIC` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetSimpleRouteDetail(context.Background(), routeId).FromStationId(fromStationId).ToStationId(toStationId).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).Tariffs(tariffs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetSimpleRouteDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimpleRouteDetail`: Route
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetSimpleRouteDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. &#x60;section0.id,section1.id,section2.id&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimpleRouteDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromStationId** | **int64** | Unique identifier for &#x60;from&#x60; station. Detailed data for identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. | 
 **toStationId** | **int64** | Unique identifier for &#x60;from&#x60; station. Detailed data for identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]
 **tariffs** | **[]string** | List of tariffs. Tariff keys can be obtained from &#x60;/consts/tariffs&#x60;. Multiple parameter instances should be used rather than multiple values when submitting multiple values. For example: &#x60;&amp;tariffs&#x3D;REGULAR&amp;tariffs&#x3D;ISIC&#x60; | 

### Return type

[**Route**](Route.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSroPassengersData

> SroPassengersDataResponse GetSroPassengersData(ctx, routeId).FromStationId(fromStationId).ToStationId(toStationId).XApplicationOrigin(xApplicationOrigin).XCurrency(xCurrency).XLang(xLang).DepartureDate(departureDate).SeatClass(seatClass).NumberOfPassengers(numberOfPassengers).Execute()

Get mandatory passengers data.



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
	routeId := "routeId_example" // string | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. `section0.id,section1.id,section2.id`).
	fromStationId := int64(17797001) // int64 | Unique identifier for `from` station. Detailed data for identifier can be obtained from `/consts/locations` endpoint.
	toStationId := int64(10204002) // int64 | Unique identifier for `from` station. Detailed data for identifier can be obtained from `/consts/locations` endpoint.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	departureDate := time.Now() // time.Time | Date and time of departure. Date is required. Time is optional. (optional)
	seatClass := int32(56) // int32 | Seat class. (optional) (default to 2)
	numberOfPassengers := int32(56) // int32 | Number of passengers. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetSroPassengersData(context.Background(), routeId).FromStationId(fromStationId).ToStationId(toStationId).XApplicationOrigin(xApplicationOrigin).XCurrency(xCurrency).XLang(xLang).DepartureDate(departureDate).SeatClass(seatClass).NumberOfPassengers(numberOfPassengers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetSroPassengersData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSroPassengersData`: SroPassengersDataResponse
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetSroPassengersData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. &#x60;section0.id,section1.id,section2.id&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSroPassengersDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromStationId** | **int64** | Unique identifier for &#x60;from&#x60; station. Detailed data for identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. | 
 **toStationId** | **int64** | Unique identifier for &#x60;from&#x60; station. Detailed data for identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **departureDate** | **time.Time** | Date and time of departure. Date is required. Time is optional. | 
 **seatClass** | **int32** | Seat class. | [default to 2]
 **numberOfPassengers** | **int32** | Number of passengers. | [default to 1]

### Return type

[**SroPassengersDataResponse**](SroPassengersDataResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/1.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSroRoutePrices

> SroRouteDetail GetSroRoutePrices(ctx, routeId).FromStationId(fromStationId).ToStationId(toStationId).XApplicationOrigin(xApplicationOrigin).XCurrency(xCurrency).XLang(xLang).DepartureDate(departureDate).SeatClass(seatClass).NumberOfPassengers(numberOfPassengers).Execute()

Get route detail.



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
	routeId := "routeId_example" // string | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. `section0.id,section1.id,section2.id`).
	fromStationId := int64(17797001) // int64 | Unique identifier for `from` station. Detailed data for identifier can be obtained from `/consts/locations` endpoint.
	toStationId := int64(10204002) // int64 | Unique identifier for `from` station. Detailed data for identifier can be obtained from `/consts/locations` endpoint.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	departureDate := time.Now() // time.Time | Date and time of departure. Date is required. Time is optional. (optional)
	seatClass := int32(56) // int32 | Seat class. (optional) (default to 2)
	numberOfPassengers := int32(56) // int32 | Number of passengers. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.GetSroRoutePrices(context.Background(), routeId).FromStationId(fromStationId).ToStationId(toStationId).XApplicationOrigin(xApplicationOrigin).XCurrency(xCurrency).XLang(xLang).DepartureDate(departureDate).SeatClass(seatClass).NumberOfPassengers(numberOfPassengers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.GetSroRoutePrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSroRoutePrices`: SroRouteDetail
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.GetSroRoutePrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **string** | Unique identifier for the route. It consists of route section identifiers separed by commas (e.g. &#x60;section0.id,section1.id,section2.id&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSroRoutePricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromStationId** | **int64** | Unique identifier for &#x60;from&#x60; station. Detailed data for identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. | 
 **toStationId** | **int64** | Unique identifier for &#x60;from&#x60; station. Detailed data for identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **departureDate** | **time.Time** | Date and time of departure. Date is required. Time is optional. | 
 **seatClass** | **int32** | Seat class. | [default to 2]
 **numberOfPassengers** | **int32** | Number of passengers. | [default to 1]

### Return type

[**SroRouteDetail**](SroRouteDetail.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRoutes

> SearchResult SearchRoutes(ctx).FromLocationId(fromLocationId).FromLocationType(fromLocationType).ToLocationId(toLocationId).ToLocationType(toLocationType).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).DepartureTime(departureTime).Tariffs(tariffs).ActionPrice(actionPrice).Execute()

Get collection of all routes that satisfy specified search criteria for the route.



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
	fromLocationId := int64(789) // int64 | 
	fromLocationType := "fromLocationType_example" // string | 
	toLocationId := int64(789) // int64 | 
	toLocationType := "toLocationType_example" // string | 
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")
	departureTime := time.Now() // time.Time | Departure date-time (optional)
	tariffs := []string{"Inner_example"} // []string | Ticket tariff (optional)
	actionPrice := "actionPrice_example" // string | Code indication of a current marketing action. Filtres searched routes on current marketing action. List of all marketing action for current route included in endpoint /consts/actionPrices. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.SearchRoutes(context.Background()).FromLocationId(fromLocationId).FromLocationType(fromLocationType).ToLocationId(toLocationId).ToLocationType(toLocationType).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).DepartureTime(departureTime).Tariffs(tariffs).ActionPrice(actionPrice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.SearchRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRoutes`: SearchResult
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.SearchRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromLocationId** | **int64** |  | 
 **fromLocationType** | **string** |  | 
 **toLocationId** | **int64** |  | 
 **toLocationType** | **string** |  | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]
 **departureTime** | **time.Time** | Departure date-time | 
 **tariffs** | **[]string** | Ticket tariff | 
 **actionPrice** | **string** | Code indication of a current marketing action. Filtres searched routes on current marketing action. List of all marketing action for current route included in endpoint /consts/actionPrices. | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSR70Routes

> []SroRoute SearchSR70Routes(ctx, connection, fromLocation, toLocation, ticketType).Departure(departure).XApplicationOrigin(xApplicationOrigin).XCurrency(xCurrency).XLang(xLang).SeatClass(seatClass).NumberOfPassengers(numberOfPassengers).Execute()

Search for available routes.



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
	departure := time.Now() // time.Time | Date and time of departure, mandatory for this request.
	connection := int64(789) // int64 | CIS connection identifier.
	fromLocation := int64(789) // int64 | Departure station numeric identifier in the SR70 format.
	toLocation := "toLocation_example" // string | Arrival station numeric identifier in the SR70 format.
	ticketType := "ticketType_example" // string | Ticket type.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	seatClass := int32(56) // int32 | Seat class. (optional) (default to 2)
	numberOfPassengers := int32(56) // int32 | Number of passengers. (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.SearchSR70Routes(context.Background(), connection, fromLocation, toLocation, ticketType).Departure(departure).XApplicationOrigin(xApplicationOrigin).XCurrency(xCurrency).XLang(xLang).SeatClass(seatClass).NumberOfPassengers(numberOfPassengers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.SearchSR70Routes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSR70Routes`: []SroRoute
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.SearchSR70Routes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connection** | **int64** | CIS connection identifier. | 
**fromLocation** | **int64** | Departure station numeric identifier in the SR70 format. | 
**toLocation** | **string** | Arrival station numeric identifier in the SR70 format. | 
**ticketType** | **string** | Ticket type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSR70RoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **departure** | **time.Time** | Date and time of departure, mandatory for this request. | 




 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **seatClass** | **int32** | Seat class. | [default to 2]
 **numberOfPassengers** | **int32** | Number of passengers. | [default to 1]

### Return type

[**[]SroRoute**](SroRoute.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimpleSearchRoutes

> SimpleRouteSearchResult SimpleSearchRoutes(ctx).FromLocationId(fromLocationId).FromLocationType(fromLocationType).ToLocationId(toLocationId).ToLocationType(toLocationType).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).XUsedDepartureFromDateTime(xUsedDepartureFromDateTime).XUsedDepartureToDateTime(xUsedDepartureToDateTime).DepartureDate(departureDate).Tariffs(tariffs).ActionPrice(actionPrice).Move(move).Execute()

Search for routes that satisfy specified search criteria.



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
	fromLocationId := int64(789) // int64 | Unique identifier for `from` location. This identifier can be obtained from `/consts/locations` endpoint. Unique identifier of either city, or station, can be used.
	fromLocationType := "fromLocationType_example" // string | Location type of unique identifier used in query parameter `fromLocationId`.
	toLocationId := int64(789) // int64 | Unique identifier for `to` location. This identifier can be obtained from `/consts/locations` endpoint. Unique identifier of either city, or station, can be used.
	toLocationType := "toLocationType_example" // string | Location type of unique identifier used in query parameter `toLocationId`.
	xApplicationOrigin := "xApplicationOrigin_example" // string | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type (optional) (default to "AFF")
	xLang := "xLang_example" // string | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  (optional) (default to "en")
	xCurrency := "xCurrency_example" // string | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  (optional) (default to "EUR")
	xUsedDepartureFromDateTime := time.Now() // time.Time | Should contain `date-time` value from previous route search. Its value is returned in response header of the same name. Used to move forward or backward in search results (i.e. for pagination). (optional)
	xUsedDepartureToDateTime := time.Now() // time.Time | Should contain `date-time` value from previous route search. Its value is returned in response header of the same name. Used to move forward or backward in search results (i.e. for pagination). (optional)
	departureDate := time.Now() // string | Departure date. (optional)
	tariffs := []string{"Inner_example"} // []string | List of tariffs. Tariff keys can be obtained from `/consts/tariffs`. Multiple parameter instances should be used rather than multiple values when submitting multiple values. For example: `&tariffs=REGULAR&tariffs=ISIC` (optional)
	actionPrice := "actionPrice_example" // string | Code indentifier of a marketing action. Filters search result based on applicability of provided marketing action to a route. List of all marketing actions for current route can be obtained from endpoint `/consts/actionPrices`. (optional)
	move := "move_example" // string | Move forward or backward in search results. Defines if next call to this endpoint will return next or previous result page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoutesAPI.SimpleSearchRoutes(context.Background()).FromLocationId(fromLocationId).FromLocationType(fromLocationType).ToLocationId(toLocationId).ToLocationType(toLocationType).XApplicationOrigin(xApplicationOrigin).XLang(xLang).XCurrency(xCurrency).XUsedDepartureFromDateTime(xUsedDepartureFromDateTime).XUsedDepartureToDateTime(xUsedDepartureToDateTime).DepartureDate(departureDate).Tariffs(tariffs).ActionPrice(actionPrice).Move(move).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutesAPI.SimpleSearchRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SimpleSearchRoutes`: SimpleRouteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `RoutesAPI.SimpleSearchRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimpleSearchRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromLocationId** | **int64** | Unique identifier for &#x60;from&#x60; location. This identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. Unique identifier of either city, or station, can be used. | 
 **fromLocationType** | **string** | Location type of unique identifier used in query parameter &#x60;fromLocationId&#x60;. | 
 **toLocationId** | **int64** | Unique identifier for &#x60;to&#x60; location. This identifier can be obtained from &#x60;/consts/locations&#x60; endpoint. Unique identifier of either city, or station, can be used. | 
 **toLocationType** | **string** | Location type of unique identifier used in query parameter &#x60;toLocationId&#x60;. | 
 **xApplicationOrigin** | **string** | Application origin - APP - Mobile application (Android / Apple) - AFF - Affiliate application which is managed by third party - CAT - Web application used to sell catering - DEV - Only for development and testing - DOT - Check-in application for ticket sales on a train or bus - NOT - Unknown application type | [default to &quot;AFF&quot;]
 **xLang** | **string** | A two-letter language code from [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes). Defines the language into which the response will be translated.  Currently supported languages:  | ISO language name | ISO 639-1 code | |-------------------|----------------| | Czech             | cs             | | German            | de             | | English           | en             | | Spanish           | es             | | French            | fr             | | Hungarian         | hu             | | Russian           | ru             | | Slovak            | sk             | | Ukrainian         | uk             | | Chinese           | zh             |  | [default to &quot;en&quot;]
 **xCurrency** | **string** | A three-letter currency code from [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217). Defines monetary unit to be used in respons.  Currently supported currencies:  | Currency     | ISO 4217 code | |--------------|---------------| | Czech koruna | CZK           | | Euro         | EUR           |  | [default to &quot;EUR&quot;]
 **xUsedDepartureFromDateTime** | **time.Time** | Should contain &#x60;date-time&#x60; value from previous route search. Its value is returned in response header of the same name. Used to move forward or backward in search results (i.e. for pagination). | 
 **xUsedDepartureToDateTime** | **time.Time** | Should contain &#x60;date-time&#x60; value from previous route search. Its value is returned in response header of the same name. Used to move forward or backward in search results (i.e. for pagination). | 
 **departureDate** | **string** | Departure date. | 
 **tariffs** | **[]string** | List of tariffs. Tariff keys can be obtained from &#x60;/consts/tariffs&#x60;. Multiple parameter instances should be used rather than multiple values when submitting multiple values. For example: &#x60;&amp;tariffs&#x3D;REGULAR&amp;tariffs&#x3D;ISIC&#x60; | 
 **actionPrice** | **string** | Code indentifier of a marketing action. Filters search result based on applicability of provided marketing action to a route. List of all marketing actions for current route can be obtained from endpoint &#x60;/consts/actionPrices&#x60;. | 
 **move** | **string** | Move forward or backward in search results. Defines if next call to this endpoint will return next or previous result page. | 

### Return type

[**SimpleRouteSearchResult**](SimpleRouteSearchResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

