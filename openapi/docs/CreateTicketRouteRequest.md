# CreateTicketRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteId** | **string** | route id (section0.id,section1.id, ... sectionx.id) | 
**SeatClass** | **string** |  | 
**PriceSource** | **string** | Pricing ID - used for confirmation that price, services or conditions werent changed | 
**ActionPrice** | Pointer to [**ActionPrice**](ActionPrice.md) |  | [optional] 
**Sections** | [**[]CreateTicketSectionRequest**](CreateTicketSectionRequest.md) |  | 

## Methods

### NewCreateTicketRouteRequest

`func NewCreateTicketRouteRequest(routeId string, seatClass string, priceSource string, sections []CreateTicketSectionRequest, ) *CreateTicketRouteRequest`

NewCreateTicketRouteRequest instantiates a new CreateTicketRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketRouteRequestWithDefaults

`func NewCreateTicketRouteRequestWithDefaults() *CreateTicketRouteRequest`

NewCreateTicketRouteRequestWithDefaults instantiates a new CreateTicketRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteId

`func (o *CreateTicketRouteRequest) GetRouteId() string`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *CreateTicketRouteRequest) GetRouteIdOk() (*string, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *CreateTicketRouteRequest) SetRouteId(v string)`

SetRouteId sets RouteId field to given value.


### GetSeatClass

`func (o *CreateTicketRouteRequest) GetSeatClass() string`

GetSeatClass returns the SeatClass field if non-nil, zero value otherwise.

### GetSeatClassOk

`func (o *CreateTicketRouteRequest) GetSeatClassOk() (*string, bool)`

GetSeatClassOk returns a tuple with the SeatClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClass

`func (o *CreateTicketRouteRequest) SetSeatClass(v string)`

SetSeatClass sets SeatClass field to given value.


### GetPriceSource

`func (o *CreateTicketRouteRequest) GetPriceSource() string`

GetPriceSource returns the PriceSource field if non-nil, zero value otherwise.

### GetPriceSourceOk

`func (o *CreateTicketRouteRequest) GetPriceSourceOk() (*string, bool)`

GetPriceSourceOk returns a tuple with the PriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSource

`func (o *CreateTicketRouteRequest) SetPriceSource(v string)`

SetPriceSource sets PriceSource field to given value.


### GetActionPrice

`func (o *CreateTicketRouteRequest) GetActionPrice() ActionPrice`

GetActionPrice returns the ActionPrice field if non-nil, zero value otherwise.

### GetActionPriceOk

`func (o *CreateTicketRouteRequest) GetActionPriceOk() (*ActionPrice, bool)`

GetActionPriceOk returns a tuple with the ActionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPrice

`func (o *CreateTicketRouteRequest) SetActionPrice(v ActionPrice)`

SetActionPrice sets ActionPrice field to given value.

### HasActionPrice

`func (o *CreateTicketRouteRequest) HasActionPrice() bool`

HasActionPrice returns a boolean if a field has been set.

### GetSections

`func (o *CreateTicketRouteRequest) GetSections() []CreateTicketSectionRequest`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *CreateTicketRouteRequest) GetSectionsOk() (*[]CreateTicketSectionRequest, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *CreateTicketRouteRequest) SetSections(v []CreateTicketSectionRequest)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


