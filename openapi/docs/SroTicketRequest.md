# SroTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteId** | **string** |  | 
**PriceSourceId** | **string** |  | 
**SeatClass** | **string** |  | 
**Sections** | [**[]SroTicketSections**](SroTicketSections.md) |  | 
**Passengers** | [**[]SroTicketPassenger**](SroTicketPassenger.md) |  | 

## Methods

### NewSroTicketRequest

`func NewSroTicketRequest(routeId string, priceSourceId string, seatClass string, sections []SroTicketSections, passengers []SroTicketPassenger, ) *SroTicketRequest`

NewSroTicketRequest instantiates a new SroTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroTicketRequestWithDefaults

`func NewSroTicketRequestWithDefaults() *SroTicketRequest`

NewSroTicketRequestWithDefaults instantiates a new SroTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteId

`func (o *SroTicketRequest) GetRouteId() string`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *SroTicketRequest) GetRouteIdOk() (*string, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *SroTicketRequest) SetRouteId(v string)`

SetRouteId sets RouteId field to given value.


### GetPriceSourceId

`func (o *SroTicketRequest) GetPriceSourceId() string`

GetPriceSourceId returns the PriceSourceId field if non-nil, zero value otherwise.

### GetPriceSourceIdOk

`func (o *SroTicketRequest) GetPriceSourceIdOk() (*string, bool)`

GetPriceSourceIdOk returns a tuple with the PriceSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSourceId

`func (o *SroTicketRequest) SetPriceSourceId(v string)`

SetPriceSourceId sets PriceSourceId field to given value.


### GetSeatClass

`func (o *SroTicketRequest) GetSeatClass() string`

GetSeatClass returns the SeatClass field if non-nil, zero value otherwise.

### GetSeatClassOk

`func (o *SroTicketRequest) GetSeatClassOk() (*string, bool)`

GetSeatClassOk returns a tuple with the SeatClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClass

`func (o *SroTicketRequest) SetSeatClass(v string)`

SetSeatClass sets SeatClass field to given value.


### GetSections

`func (o *SroTicketRequest) GetSections() []SroTicketSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SroTicketRequest) GetSectionsOk() (*[]SroTicketSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SroTicketRequest) SetSections(v []SroTicketSections)`

SetSections sets Sections field to given value.


### GetPassengers

`func (o *SroTicketRequest) GetPassengers() []SroTicketPassenger`

GetPassengers returns the Passengers field if non-nil, zero value otherwise.

### GetPassengersOk

`func (o *SroTicketRequest) GetPassengersOk() (*[]SroTicketPassenger, bool)`

GetPassengersOk returns a tuple with the Passengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengers

`func (o *SroTicketRequest) SetPassengers(v []SroTicketPassenger)`

SetPassengers sets Passengers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


