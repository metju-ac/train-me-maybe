# SimpleRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of a route. Consists of unique sections identifiers separated by commas. | 
**DepartureStationId** | Pointer to **int64** | Unique identifier of departure station. | [optional] 
**DepartureTime** | **time.Time** | Departure time. | 
**ArrivalStationId** | Pointer to **int64** | Unique identifier of arrival station. | [optional] 
**ArrivalTime** | **time.Time** | Arrival time. | 
**VehicleTypes** | [**[]VehicleType**](VehicleType.md) | Vehicle type. | 
**TransfersCount** | Pointer to **int32** | Number of transfers between stations. | [optional] 
**FreeSeatsCount** | **int32** | Number of available free seats through all sections. | 
**PriceFrom** | **float32** | Minimum ticket price for open account. | 
**PriceTo** | Pointer to **float32** | Maximum ticket price for open account. | [optional] 
**CreditPriceFrom** | **float32** | Minimum ticket price for RegioJet Pay. | 
**CreditPriceTo** | Pointer to **float32** | Maximum ticket price for RegioJet Pay. | [optional] 
**PricesCount** | **int32** | Number of available prices. | 
**ActionPrice** | **bool** | &#x60;true&#x60; if any price is action price, otherwise &#x60;false&#x60;. | 
**Surcharge** | **bool** | &#x60;true&#x60; if there is surcharge on this route, otherwise &#x60;false&#x60;. | 
**Notices** | **bool** | Notice of any extraordinary events on route or traffic problems. | 
**Support** | **bool** | &#x60;true&#x60; if this route (or its part) have support (e.g. in form of additional railroad cars), otherwise &#x60;false&#x60;. | 
**NationalTrip** | Pointer to **bool** | &#x60;true&#x60; if route is interstate, &#x60;false&#x60; if route is international. | [optional] 
**Bookable** | **bool** | &#x60;true&#x60; if at least one class have enough free seats for reservation, otherwise &#x60;false&#x60;. | 
**Delay** | Pointer to **string** | Textual information about the first delay on the route. | [optional] 
**TravelTime** | Pointer to **string** | Textual information about route&#39;s travel time. | [optional] 
**VehicleStandardKey** | Pointer to **string** | Vehicle standard key. Detailed data about every available vehicle standard can be obtained from &#x60;/consts/vehicleStandards&#x60;. | [optional] 

## Methods

### NewSimpleRoute

`func NewSimpleRoute(id string, departureTime time.Time, arrivalTime time.Time, vehicleTypes []VehicleType, freeSeatsCount int32, priceFrom float32, creditPriceFrom float32, pricesCount int32, actionPrice bool, surcharge bool, notices bool, support bool, bookable bool, ) *SimpleRoute`

NewSimpleRoute instantiates a new SimpleRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleRouteWithDefaults

`func NewSimpleRouteWithDefaults() *SimpleRoute`

NewSimpleRouteWithDefaults instantiates a new SimpleRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleRoute) SetId(v string)`

SetId sets Id field to given value.


### GetDepartureStationId

`func (o *SimpleRoute) GetDepartureStationId() int64`

GetDepartureStationId returns the DepartureStationId field if non-nil, zero value otherwise.

### GetDepartureStationIdOk

`func (o *SimpleRoute) GetDepartureStationIdOk() (*int64, bool)`

GetDepartureStationIdOk returns a tuple with the DepartureStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureStationId

`func (o *SimpleRoute) SetDepartureStationId(v int64)`

SetDepartureStationId sets DepartureStationId field to given value.

### HasDepartureStationId

`func (o *SimpleRoute) HasDepartureStationId() bool`

HasDepartureStationId returns a boolean if a field has been set.

### GetDepartureTime

`func (o *SimpleRoute) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *SimpleRoute) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *SimpleRoute) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.


### GetArrivalStationId

`func (o *SimpleRoute) GetArrivalStationId() int64`

GetArrivalStationId returns the ArrivalStationId field if non-nil, zero value otherwise.

### GetArrivalStationIdOk

`func (o *SimpleRoute) GetArrivalStationIdOk() (*int64, bool)`

GetArrivalStationIdOk returns a tuple with the ArrivalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalStationId

`func (o *SimpleRoute) SetArrivalStationId(v int64)`

SetArrivalStationId sets ArrivalStationId field to given value.

### HasArrivalStationId

`func (o *SimpleRoute) HasArrivalStationId() bool`

HasArrivalStationId returns a boolean if a field has been set.

### GetArrivalTime

`func (o *SimpleRoute) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *SimpleRoute) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *SimpleRoute) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.


### GetVehicleTypes

`func (o *SimpleRoute) GetVehicleTypes() []VehicleType`

GetVehicleTypes returns the VehicleTypes field if non-nil, zero value otherwise.

### GetVehicleTypesOk

`func (o *SimpleRoute) GetVehicleTypesOk() (*[]VehicleType, bool)`

GetVehicleTypesOk returns a tuple with the VehicleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleTypes

`func (o *SimpleRoute) SetVehicleTypes(v []VehicleType)`

SetVehicleTypes sets VehicleTypes field to given value.


### GetTransfersCount

`func (o *SimpleRoute) GetTransfersCount() int32`

GetTransfersCount returns the TransfersCount field if non-nil, zero value otherwise.

### GetTransfersCountOk

`func (o *SimpleRoute) GetTransfersCountOk() (*int32, bool)`

GetTransfersCountOk returns a tuple with the TransfersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersCount

`func (o *SimpleRoute) SetTransfersCount(v int32)`

SetTransfersCount sets TransfersCount field to given value.

### HasTransfersCount

`func (o *SimpleRoute) HasTransfersCount() bool`

HasTransfersCount returns a boolean if a field has been set.

### GetFreeSeatsCount

`func (o *SimpleRoute) GetFreeSeatsCount() int32`

GetFreeSeatsCount returns the FreeSeatsCount field if non-nil, zero value otherwise.

### GetFreeSeatsCountOk

`func (o *SimpleRoute) GetFreeSeatsCountOk() (*int32, bool)`

GetFreeSeatsCountOk returns a tuple with the FreeSeatsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSeatsCount

`func (o *SimpleRoute) SetFreeSeatsCount(v int32)`

SetFreeSeatsCount sets FreeSeatsCount field to given value.


### GetPriceFrom

`func (o *SimpleRoute) GetPriceFrom() float32`

GetPriceFrom returns the PriceFrom field if non-nil, zero value otherwise.

### GetPriceFromOk

`func (o *SimpleRoute) GetPriceFromOk() (*float32, bool)`

GetPriceFromOk returns a tuple with the PriceFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFrom

`func (o *SimpleRoute) SetPriceFrom(v float32)`

SetPriceFrom sets PriceFrom field to given value.


### GetPriceTo

`func (o *SimpleRoute) GetPriceTo() float32`

GetPriceTo returns the PriceTo field if non-nil, zero value otherwise.

### GetPriceToOk

`func (o *SimpleRoute) GetPriceToOk() (*float32, bool)`

GetPriceToOk returns a tuple with the PriceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTo

`func (o *SimpleRoute) SetPriceTo(v float32)`

SetPriceTo sets PriceTo field to given value.

### HasPriceTo

`func (o *SimpleRoute) HasPriceTo() bool`

HasPriceTo returns a boolean if a field has been set.

### GetCreditPriceFrom

`func (o *SimpleRoute) GetCreditPriceFrom() float32`

GetCreditPriceFrom returns the CreditPriceFrom field if non-nil, zero value otherwise.

### GetCreditPriceFromOk

`func (o *SimpleRoute) GetCreditPriceFromOk() (*float32, bool)`

GetCreditPriceFromOk returns a tuple with the CreditPriceFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPriceFrom

`func (o *SimpleRoute) SetCreditPriceFrom(v float32)`

SetCreditPriceFrom sets CreditPriceFrom field to given value.


### GetCreditPriceTo

`func (o *SimpleRoute) GetCreditPriceTo() float32`

GetCreditPriceTo returns the CreditPriceTo field if non-nil, zero value otherwise.

### GetCreditPriceToOk

`func (o *SimpleRoute) GetCreditPriceToOk() (*float32, bool)`

GetCreditPriceToOk returns a tuple with the CreditPriceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPriceTo

`func (o *SimpleRoute) SetCreditPriceTo(v float32)`

SetCreditPriceTo sets CreditPriceTo field to given value.

### HasCreditPriceTo

`func (o *SimpleRoute) HasCreditPriceTo() bool`

HasCreditPriceTo returns a boolean if a field has been set.

### GetPricesCount

`func (o *SimpleRoute) GetPricesCount() int32`

GetPricesCount returns the PricesCount field if non-nil, zero value otherwise.

### GetPricesCountOk

`func (o *SimpleRoute) GetPricesCountOk() (*int32, bool)`

GetPricesCountOk returns a tuple with the PricesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricesCount

`func (o *SimpleRoute) SetPricesCount(v int32)`

SetPricesCount sets PricesCount field to given value.


### GetActionPrice

`func (o *SimpleRoute) GetActionPrice() bool`

GetActionPrice returns the ActionPrice field if non-nil, zero value otherwise.

### GetActionPriceOk

`func (o *SimpleRoute) GetActionPriceOk() (*bool, bool)`

GetActionPriceOk returns a tuple with the ActionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPrice

`func (o *SimpleRoute) SetActionPrice(v bool)`

SetActionPrice sets ActionPrice field to given value.


### GetSurcharge

`func (o *SimpleRoute) GetSurcharge() bool`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *SimpleRoute) GetSurchargeOk() (*bool, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *SimpleRoute) SetSurcharge(v bool)`

SetSurcharge sets Surcharge field to given value.


### GetNotices

`func (o *SimpleRoute) GetNotices() bool`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *SimpleRoute) GetNoticesOk() (*bool, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *SimpleRoute) SetNotices(v bool)`

SetNotices sets Notices field to given value.


### GetSupport

`func (o *SimpleRoute) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *SimpleRoute) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *SimpleRoute) SetSupport(v bool)`

SetSupport sets Support field to given value.


### GetNationalTrip

`func (o *SimpleRoute) GetNationalTrip() bool`

GetNationalTrip returns the NationalTrip field if non-nil, zero value otherwise.

### GetNationalTripOk

`func (o *SimpleRoute) GetNationalTripOk() (*bool, bool)`

GetNationalTripOk returns a tuple with the NationalTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalTrip

`func (o *SimpleRoute) SetNationalTrip(v bool)`

SetNationalTrip sets NationalTrip field to given value.

### HasNationalTrip

`func (o *SimpleRoute) HasNationalTrip() bool`

HasNationalTrip returns a boolean if a field has been set.

### GetBookable

`func (o *SimpleRoute) GetBookable() bool`

GetBookable returns the Bookable field if non-nil, zero value otherwise.

### GetBookableOk

`func (o *SimpleRoute) GetBookableOk() (*bool, bool)`

GetBookableOk returns a tuple with the Bookable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookable

`func (o *SimpleRoute) SetBookable(v bool)`

SetBookable sets Bookable field to given value.


### GetDelay

`func (o *SimpleRoute) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *SimpleRoute) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *SimpleRoute) SetDelay(v string)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *SimpleRoute) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetTravelTime

`func (o *SimpleRoute) GetTravelTime() string`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *SimpleRoute) GetTravelTimeOk() (*string, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *SimpleRoute) SetTravelTime(v string)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *SimpleRoute) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.

### GetVehicleStandardKey

`func (o *SimpleRoute) GetVehicleStandardKey() string`

GetVehicleStandardKey returns the VehicleStandardKey field if non-nil, zero value otherwise.

### GetVehicleStandardKeyOk

`func (o *SimpleRoute) GetVehicleStandardKeyOk() (*string, bool)`

GetVehicleStandardKeyOk returns a tuple with the VehicleStandardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleStandardKey

`func (o *SimpleRoute) SetVehicleStandardKey(v string)`

SetVehicleStandardKey sets VehicleStandardKey field to given value.

### HasVehicleStandardKey

`func (o *SimpleRoute) HasVehicleStandardKey() bool`

HasVehicleStandardKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


