# SroRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Comma-separated list of route section identifiers. | 
**ConnectionCode** | **string** | Connection code, as recognized by an external entity (e.g. OneTicket), not bound to any format. | 
**DepartureStationId** | **int64** | Numeric and unique representation of the departure station. | 
**DepartureTime** | **time.Time** | Departure time in the ISO 8601 format. | 
**ArrivalStationId** | Pointer to **int64** | Numeric and unique representation of the arrival station. | [optional] 
**ArrivalTime** | **time.Time** | Arrival time in the ISO 8601 format. | 
**VehicleTypes** | [**[]VehicleType**](VehicleType.md) | Array of vehicle types. | 
**TransfersCount** | Pointer to **int32** | Number of transfers. | [optional] 
**PriceFrom** | **float32** | Price of the least expensive route. | 
**PriceTo** | Pointer to **float32** | Price of the most expensive route. | [optional] 
**PricesCount** | **int32** | Number of returned prices. | 
**Notices** | **bool** | Are there any traffic limitations or extraordinary events on the route? | 
**Support** | **bool** | Does this line have a support connection? | 
**NationalTrip** | Pointer to **bool** | Is this trip international? | [optional] 
**Delay** | Pointer to **string** | Information about the first delay on the route in the &#39;HH:MM \\h&#39; format. | [optional] 
**TravelTime** | Pointer to **string** | Travel time in the &#39;HH:MM \\h&#39; format. | [optional] 
**VehicleStandardKey** | Pointer to **string** | Vehicle standard code. | [optional] 

## Methods

### NewSroRoute

`func NewSroRoute(id string, connectionCode string, departureStationId int64, departureTime time.Time, arrivalTime time.Time, vehicleTypes []VehicleType, priceFrom float32, pricesCount int32, notices bool, support bool, ) *SroRoute`

NewSroRoute instantiates a new SroRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroRouteWithDefaults

`func NewSroRouteWithDefaults() *SroRoute`

NewSroRouteWithDefaults instantiates a new SroRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SroRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SroRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SroRoute) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionCode

`func (o *SroRoute) GetConnectionCode() string`

GetConnectionCode returns the ConnectionCode field if non-nil, zero value otherwise.

### GetConnectionCodeOk

`func (o *SroRoute) GetConnectionCodeOk() (*string, bool)`

GetConnectionCodeOk returns a tuple with the ConnectionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCode

`func (o *SroRoute) SetConnectionCode(v string)`

SetConnectionCode sets ConnectionCode field to given value.


### GetDepartureStationId

`func (o *SroRoute) GetDepartureStationId() int64`

GetDepartureStationId returns the DepartureStationId field if non-nil, zero value otherwise.

### GetDepartureStationIdOk

`func (o *SroRoute) GetDepartureStationIdOk() (*int64, bool)`

GetDepartureStationIdOk returns a tuple with the DepartureStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureStationId

`func (o *SroRoute) SetDepartureStationId(v int64)`

SetDepartureStationId sets DepartureStationId field to given value.


### GetDepartureTime

`func (o *SroRoute) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *SroRoute) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *SroRoute) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.


### GetArrivalStationId

`func (o *SroRoute) GetArrivalStationId() int64`

GetArrivalStationId returns the ArrivalStationId field if non-nil, zero value otherwise.

### GetArrivalStationIdOk

`func (o *SroRoute) GetArrivalStationIdOk() (*int64, bool)`

GetArrivalStationIdOk returns a tuple with the ArrivalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalStationId

`func (o *SroRoute) SetArrivalStationId(v int64)`

SetArrivalStationId sets ArrivalStationId field to given value.

### HasArrivalStationId

`func (o *SroRoute) HasArrivalStationId() bool`

HasArrivalStationId returns a boolean if a field has been set.

### GetArrivalTime

`func (o *SroRoute) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *SroRoute) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *SroRoute) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.


### GetVehicleTypes

`func (o *SroRoute) GetVehicleTypes() []VehicleType`

GetVehicleTypes returns the VehicleTypes field if non-nil, zero value otherwise.

### GetVehicleTypesOk

`func (o *SroRoute) GetVehicleTypesOk() (*[]VehicleType, bool)`

GetVehicleTypesOk returns a tuple with the VehicleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleTypes

`func (o *SroRoute) SetVehicleTypes(v []VehicleType)`

SetVehicleTypes sets VehicleTypes field to given value.


### GetTransfersCount

`func (o *SroRoute) GetTransfersCount() int32`

GetTransfersCount returns the TransfersCount field if non-nil, zero value otherwise.

### GetTransfersCountOk

`func (o *SroRoute) GetTransfersCountOk() (*int32, bool)`

GetTransfersCountOk returns a tuple with the TransfersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersCount

`func (o *SroRoute) SetTransfersCount(v int32)`

SetTransfersCount sets TransfersCount field to given value.

### HasTransfersCount

`func (o *SroRoute) HasTransfersCount() bool`

HasTransfersCount returns a boolean if a field has been set.

### GetPriceFrom

`func (o *SroRoute) GetPriceFrom() float32`

GetPriceFrom returns the PriceFrom field if non-nil, zero value otherwise.

### GetPriceFromOk

`func (o *SroRoute) GetPriceFromOk() (*float32, bool)`

GetPriceFromOk returns a tuple with the PriceFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFrom

`func (o *SroRoute) SetPriceFrom(v float32)`

SetPriceFrom sets PriceFrom field to given value.


### GetPriceTo

`func (o *SroRoute) GetPriceTo() float32`

GetPriceTo returns the PriceTo field if non-nil, zero value otherwise.

### GetPriceToOk

`func (o *SroRoute) GetPriceToOk() (*float32, bool)`

GetPriceToOk returns a tuple with the PriceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTo

`func (o *SroRoute) SetPriceTo(v float32)`

SetPriceTo sets PriceTo field to given value.

### HasPriceTo

`func (o *SroRoute) HasPriceTo() bool`

HasPriceTo returns a boolean if a field has been set.

### GetPricesCount

`func (o *SroRoute) GetPricesCount() int32`

GetPricesCount returns the PricesCount field if non-nil, zero value otherwise.

### GetPricesCountOk

`func (o *SroRoute) GetPricesCountOk() (*int32, bool)`

GetPricesCountOk returns a tuple with the PricesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesCount

`func (o *SroRoute) SetPricesCount(v int32)`

SetPricesCount sets PricesCount field to given value.


### GetNotices

`func (o *SroRoute) GetNotices() bool`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *SroRoute) GetNoticesOk() (*bool, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *SroRoute) SetNotices(v bool)`

SetNotices sets Notices field to given value.


### GetSupport

`func (o *SroRoute) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *SroRoute) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *SroRoute) SetSupport(v bool)`

SetSupport sets Support field to given value.


### GetNationalTrip

`func (o *SroRoute) GetNationalTrip() bool`

GetNationalTrip returns the NationalTrip field if non-nil, zero value otherwise.

### GetNationalTripOk

`func (o *SroRoute) GetNationalTripOk() (*bool, bool)`

GetNationalTripOk returns a tuple with the NationalTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalTrip

`func (o *SroRoute) SetNationalTrip(v bool)`

SetNationalTrip sets NationalTrip field to given value.

### HasNationalTrip

`func (o *SroRoute) HasNationalTrip() bool`

HasNationalTrip returns a boolean if a field has been set.

### GetDelay

`func (o *SroRoute) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *SroRoute) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *SroRoute) SetDelay(v string)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *SroRoute) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetTravelTime

`func (o *SroRoute) GetTravelTime() string`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *SroRoute) GetTravelTimeOk() (*string, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *SroRoute) SetTravelTime(v string)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *SroRoute) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.

### GetVehicleStandardKey

`func (o *SroRoute) GetVehicleStandardKey() string`

GetVehicleStandardKey returns the VehicleStandardKey field if non-nil, zero value otherwise.

### GetVehicleStandardKeyOk

`func (o *SroRoute) GetVehicleStandardKeyOk() (*string, bool)`

GetVehicleStandardKeyOk returns a tuple with the VehicleStandardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleStandardKey

`func (o *SroRoute) SetVehicleStandardKey(v string)`

SetVehicleStandardKey sets VehicleStandardKey field to given value.

### HasVehicleStandardKey

`func (o *SroRoute) HasVehicleStandardKey() bool`

HasVehicleStandardKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


