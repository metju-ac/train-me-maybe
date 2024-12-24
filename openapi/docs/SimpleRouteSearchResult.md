# SimpleRouteSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | Pointer to [**[]SimpleRoute**](SimpleRoute.md) |  | [optional] 
**RoutesMessage** | Pointer to **string** | Message describing exceptional state of route search. | [optional] 
**BannerBubbles** | Pointer to [**[]BannerBubble**](BannerBubble.md) |  | [optional] 
**TextBubbles** | Pointer to [**[]TextBubble**](TextBubble.md) |  | [optional] 

## Methods

### NewSimpleRouteSearchResult

`func NewSimpleRouteSearchResult() *SimpleRouteSearchResult`

NewSimpleRouteSearchResult instantiates a new SimpleRouteSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleRouteSearchResultWithDefaults

`func NewSimpleRouteSearchResultWithDefaults() *SimpleRouteSearchResult`

NewSimpleRouteSearchResultWithDefaults instantiates a new SimpleRouteSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *SimpleRouteSearchResult) GetRoutes() []SimpleRoute`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *SimpleRouteSearchResult) GetRoutesOk() (*[]SimpleRoute, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *SimpleRouteSearchResult) SetRoutes(v []SimpleRoute)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *SimpleRouteSearchResult) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetRoutesMessage

`func (o *SimpleRouteSearchResult) GetRoutesMessage() string`

GetRoutesMessage returns the RoutesMessage field if non-nil, zero value otherwise.

### GetRoutesMessageOk

`func (o *SimpleRouteSearchResult) GetRoutesMessageOk() (*string, bool)`

GetRoutesMessageOk returns a tuple with the RoutesMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutesMessage

`func (o *SimpleRouteSearchResult) SetRoutesMessage(v string)`

SetRoutesMessage sets RoutesMessage field to given value.

### HasRoutesMessage

`func (o *SimpleRouteSearchResult) HasRoutesMessage() bool`

HasRoutesMessage returns a boolean if a field has been set.

### GetBannerBubbles

`func (o *SimpleRouteSearchResult) GetBannerBubbles() []BannerBubble`

GetBannerBubbles returns the BannerBubbles field if non-nil, zero value otherwise.

### GetBannerBubblesOk

`func (o *SimpleRouteSearchResult) GetBannerBubblesOk() (*[]BannerBubble, bool)`

GetBannerBubblesOk returns a tuple with the BannerBubbles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerBubbles

`func (o *SimpleRouteSearchResult) SetBannerBubbles(v []BannerBubble)`

SetBannerBubbles sets BannerBubbles field to given value.

### HasBannerBubbles

`func (o *SimpleRouteSearchResult) HasBannerBubbles() bool`

HasBannerBubbles returns a boolean if a field has been set.

### GetTextBubbles

`func (o *SimpleRouteSearchResult) GetTextBubbles() []TextBubble`

GetTextBubbles returns the TextBubbles field if non-nil, zero value otherwise.

### GetTextBubblesOk

`func (o *SimpleRouteSearchResult) GetTextBubblesOk() (*[]TextBubble, bool)`

GetTextBubblesOk returns a tuple with the TextBubbles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBubbles

`func (o *SimpleRouteSearchResult) SetTextBubbles(v []TextBubble)`

SetTextBubbles sets TextBubbles field to given value.

### HasTextBubbles

`func (o *SimpleRouteSearchResult) HasTextBubbles() bool`

HasTextBubbles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


