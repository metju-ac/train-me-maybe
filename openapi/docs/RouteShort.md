# RouteShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of a route. Consists of unique sections identifiers separated by commas. | 
**MainSectionId** | **int64** | Main section of the route. Only for this section can customer select seat class or vehicle standard. | 
**DepartureStationId** | **int64** |  | 
**DepartureStationName** | Pointer to **string** |  | [optional] 
**DepartureCityId** | **int64** |  | 
**DepartureCityName** | Pointer to **string** |  | [optional] 
**DepartureTime** | **time.Time** |  | 
**ArrivalStationId** | **int64** |  | 
**ArrivalStationName** | Pointer to **string** |  | [optional] 
**ArrivalCityId** | **int64** |  | 
**ArrivalCityName** | Pointer to **string** |  | [optional] 
**ArrivalTime** | **time.Time** |  | 
**FreeSeatsCount** | **int32** | Number of available free seats through all sections. | 
**PriceFrom** | Pointer to **float32** |  | [optional] 
**PriceTo** | Pointer to **float32** |  | [optional] 
**CreditPriceFrom** | Pointer to **float32** |  | [optional] 
**CreditPriceTo** | Pointer to **float32** |  | [optional] 
**VehicleTypes** | Pointer to [**[]VehicleType**](VehicleType.md) |  | [optional] 
**Surcharge** | Pointer to [**Surcharge**](Surcharge.md) |  | [optional] 
**Sections** | [**[]Section**](Section.md) |  | 
**Notices** | Pointer to **bool** | Notice of any extraordinarily events on route / traffic limitation | [optional] 
**TransfersInfo** | Pointer to [**TransfersInfo**](TransfersInfo.md) |  | [optional] 
**NationalTrip** | Pointer to **bool** | TRUE &#x3D;&gt; national, FALSE &#x3D;&gt; international | [optional] 
**Bookable** | **bool** | TRUE if at least one class have enough free seats for reservation, otherwise FALSE | 
**Delay** | Pointer to **string** | Textual information about the first delay on the route | [optional] 
**TravelTime** | Pointer to **string** | Textual information about the travel time on the route | [optional] 

## Methods

### NewRouteShort

`func NewRouteShort(id string, mainSectionId int64, departureStationId int64, departureCityId int64, departureTime time.Time, arrivalStationId int64, arrivalCityId int64, arrivalTime time.Time, freeSeatsCount int32, sections []Section, bookable bool, ) *RouteShort`

NewRouteShort instantiates a new RouteShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteShortWithDefaults

`func NewRouteShortWithDefaults() *RouteShort`

NewRouteShortWithDefaults instantiates a new RouteShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteShort) SetId(v string)`

SetId sets Id field to given value.


### GetMainSectionId

`func (o *RouteShort) GetMainSectionId() int64`

GetMainSectionId returns the MainSectionId field if non-nil, zero value otherwise.

### GetMainSectionIdOk

`func (o *RouteShort) GetMainSectionIdOk() (*int64, bool)`

GetMainSectionIdOk returns a tuple with the MainSectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainSectionId

`func (o *RouteShort) SetMainSectionId(v int64)`

SetMainSectionId sets MainSectionId field to given value.


### GetDepartureStationId

`func (o *RouteShort) GetDepartureStationId() int64`

GetDepartureStationId returns the DepartureStationId field if non-nil, zero value otherwise.

### GetDepartureStationIdOk

`func (o *RouteShort) GetDepartureStationIdOk() (*int64, bool)`

GetDepartureStationIdOk returns a tuple with the DepartureStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureStationId

`func (o *RouteShort) SetDepartureStationId(v int64)`

SetDepartureStationId sets DepartureStationId field to given value.


### GetDepartureStationName

`func (o *RouteShort) GetDepartureStationName() string`

GetDepartureStationName returns the DepartureStationName field if non-nil, zero value otherwise.

### GetDepartureStationNameOk

`func (o *RouteShort) GetDepartureStationNameOk() (*string, bool)`

GetDepartureStationNameOk returns a tuple with the DepartureStationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureStationName

`func (o *RouteShort) SetDepartureStationName(v string)`

SetDepartureStationName sets DepartureStationName field to given value.

### HasDepartureStationName

`func (o *RouteShort) HasDepartureStationName() bool`

HasDepartureStationName returns a boolean if a field has been set.

### GetDepartureCityId

`func (o *RouteShort) GetDepartureCityId() int64`

GetDepartureCityId returns the DepartureCityId field if non-nil, zero value otherwise.

### GetDepartureCityIdOk

`func (o *RouteShort) GetDepartureCityIdOk() (*int64, bool)`

GetDepartureCityIdOk returns a tuple with the DepartureCityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureCityId

`func (o *RouteShort) SetDepartureCityId(v int64)`

SetDepartureCityId sets DepartureCityId field to given value.


### GetDepartureCityName

`func (o *RouteShort) GetDepartureCityName() string`

GetDepartureCityName returns the DepartureCityName field if non-nil, zero value otherwise.

### GetDepartureCityNameOk

`func (o *RouteShort) GetDepartureCityNameOk() (*string, bool)`

GetDepartureCityNameOk returns a tuple with the DepartureCityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureCityName

`func (o *RouteShort) SetDepartureCityName(v string)`

SetDepartureCityName sets DepartureCityName field to given value.

### HasDepartureCityName

`func (o *RouteShort) HasDepartureCityName() bool`

HasDepartureCityName returns a boolean if a field has been set.

### GetDepartureTime

`func (o *RouteShort) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *RouteShort) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *RouteShort) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.


### GetArrivalStationId

`func (o *RouteShort) GetArrivalStationId() int64`

GetArrivalStationId returns the ArrivalStationId field if non-nil, zero value otherwise.

### GetArrivalStationIdOk

`func (o *RouteShort) GetArrivalStationIdOk() (*int64, bool)`

GetArrivalStationIdOk returns a tuple with the ArrivalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalStationId

`func (o *RouteShort) SetArrivalStationId(v int64)`

SetArrivalStationId sets ArrivalStationId field to given value.


### GetArrivalStationName

`func (o *RouteShort) GetArrivalStationName() string`

GetArrivalStationName returns the ArrivalStationName field if non-nil, zero value otherwise.

### GetArrivalStationNameOk

`func (o *RouteShort) GetArrivalStationNameOk() (*string, bool)`

GetArrivalStationNameOk returns a tuple with the ArrivalStationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalStationName

`func (o *RouteShort) SetArrivalStationName(v string)`

SetArrivalStationName sets ArrivalStationName field to given value.

### HasArrivalStationName

`func (o *RouteShort) HasArrivalStationName() bool`

HasArrivalStationName returns a boolean if a field has been set.

### GetArrivalCityId

`func (o *RouteShort) GetArrivalCityId() int64`

GetArrivalCityId returns the ArrivalCityId field if non-nil, zero value otherwise.

### GetArrivalCityIdOk

`func (o *RouteShort) GetArrivalCityIdOk() (*int64, bool)`

GetArrivalCityIdOk returns a tuple with the ArrivalCityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalCityId

`func (o *RouteShort) SetArrivalCityId(v int64)`

SetArrivalCityId sets ArrivalCityId field to given value.


### GetArrivalCityName

`func (o *RouteShort) GetArrivalCityName() string`

GetArrivalCityName returns the ArrivalCityName field if non-nil, zero value otherwise.

### GetArrivalCityNameOk

`func (o *RouteShort) GetArrivalCityNameOk() (*string, bool)`

GetArrivalCityNameOk returns a tuple with the ArrivalCityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalCityName

`func (o *RouteShort) SetArrivalCityName(v string)`

SetArrivalCityName sets ArrivalCityName field to given value.

### HasArrivalCityName

`func (o *RouteShort) HasArrivalCityName() bool`

HasArrivalCityName returns a boolean if a field has been set.

### GetArrivalTime

`func (o *RouteShort) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *RouteShort) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *RouteShort) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.


### GetFreeSeatsCount

`func (o *RouteShort) GetFreeSeatsCount() int32`

GetFreeSeatsCount returns the FreeSeatsCount field if non-nil, zero value otherwise.

### GetFreeSeatsCountOk

`func (o *RouteShort) GetFreeSeatsCountOk() (*int32, bool)`

GetFreeSeatsCountOk returns a tuple with the FreeSeatsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSeatsCount

`func (o *RouteShort) SetFreeSeatsCount(v int32)`

SetFreeSeatsCount sets FreeSeatsCount field to given value.


### GetPriceFrom

`func (o *RouteShort) GetPriceFrom() float32`

GetPriceFrom returns the PriceFrom field if non-nil, zero value otherwise.

### GetPriceFromOk

`func (o *RouteShort) GetPriceFromOk() (*float32, bool)`

GetPriceFromOk returns a tuple with the PriceFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFrom

`func (o *RouteShort) SetPriceFrom(v float32)`

SetPriceFrom sets PriceFrom field to given value.

### HasPriceFrom

`func (o *RouteShort) HasPriceFrom() bool`

HasPriceFrom returns a boolean if a field has been set.

### GetPriceTo

`func (o *RouteShort) GetPriceTo() float32`

GetPriceTo returns the PriceTo field if non-nil, zero value otherwise.

### GetPriceToOk

`func (o *RouteShort) GetPriceToOk() (*float32, bool)`

GetPriceToOk returns a tuple with the PriceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTo

`func (o *RouteShort) SetPriceTo(v float32)`

SetPriceTo sets PriceTo field to given value.

### HasPriceTo

`func (o *RouteShort) HasPriceTo() bool`

HasPriceTo returns a boolean if a field has been set.

### GetCreditPriceFrom

`func (o *RouteShort) GetCreditPriceFrom() float32`

GetCreditPriceFrom returns the CreditPriceFrom field if non-nil, zero value otherwise.

### GetCreditPriceFromOk

`func (o *RouteShort) GetCreditPriceFromOk() (*float32, bool)`

GetCreditPriceFromOk returns a tuple with the CreditPriceFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPriceFrom

`func (o *RouteShort) SetCreditPriceFrom(v float32)`

SetCreditPriceFrom sets CreditPriceFrom field to given value.

### HasCreditPriceFrom

`func (o *RouteShort) HasCreditPriceFrom() bool`

HasCreditPriceFrom returns a boolean if a field has been set.

### GetCreditPriceTo

`func (o *RouteShort) GetCreditPriceTo() float32`

GetCreditPriceTo returns the CreditPriceTo field if non-nil, zero value otherwise.

### GetCreditPriceToOk

`func (o *RouteShort) GetCreditPriceToOk() (*float32, bool)`

GetCreditPriceToOk returns a tuple with the CreditPriceTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPriceTo

`func (o *RouteShort) SetCreditPriceTo(v float32)`

SetCreditPriceTo sets CreditPriceTo field to given value.

### HasCreditPriceTo

`func (o *RouteShort) HasCreditPriceTo() bool`

HasCreditPriceTo returns a boolean if a field has been set.

### GetVehicleTypes

`func (o *RouteShort) GetVehicleTypes() []VehicleType`

GetVehicleTypes returns the VehicleTypes field if non-nil, zero value otherwise.

### GetVehicleTypesOk

`func (o *RouteShort) GetVehicleTypesOk() (*[]VehicleType, bool)`

GetVehicleTypesOk returns a tuple with the VehicleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleTypes

`func (o *RouteShort) SetVehicleTypes(v []VehicleType)`

SetVehicleTypes sets VehicleTypes field to given value.

### HasVehicleTypes

`func (o *RouteShort) HasVehicleTypes() bool`

HasVehicleTypes returns a boolean if a field has been set.

### GetSurcharge

`func (o *RouteShort) GetSurcharge() Surcharge`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *RouteShort) GetSurchargeOk() (*Surcharge, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *RouteShort) SetSurcharge(v Surcharge)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *RouteShort) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetSections

`func (o *RouteShort) GetSections() []Section`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *RouteShort) GetSectionsOk() (*[]Section, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *RouteShort) SetSections(v []Section)`

SetSections sets Sections field to given value.


### GetNotices

`func (o *RouteShort) GetNotices() bool`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *RouteShort) GetNoticesOk() (*bool, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *RouteShort) SetNotices(v bool)`

SetNotices sets Notices field to given value.

### HasNotices

`func (o *RouteShort) HasNotices() bool`

HasNotices returns a boolean if a field has been set.

### GetTransfersInfo

`func (o *RouteShort) GetTransfersInfo() TransfersInfo`

GetTransfersInfo returns the TransfersInfo field if non-nil, zero value otherwise.

### GetTransfersInfoOk

`func (o *RouteShort) GetTransfersInfoOk() (*TransfersInfo, bool)`

GetTransfersInfoOk returns a tuple with the TransfersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersInfo

`func (o *RouteShort) SetTransfersInfo(v TransfersInfo)`

SetTransfersInfo sets TransfersInfo field to given value.

### HasTransfersInfo

`func (o *RouteShort) HasTransfersInfo() bool`

HasTransfersInfo returns a boolean if a field has been set.

### GetNationalTrip

`func (o *RouteShort) GetNationalTrip() bool`

GetNationalTrip returns the NationalTrip field if non-nil, zero value otherwise.

### GetNationalTripOk

`func (o *RouteShort) GetNationalTripOk() (*bool, bool)`

GetNationalTripOk returns a tuple with the NationalTrip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalTrip

`func (o *RouteShort) SetNationalTrip(v bool)`

SetNationalTrip sets NationalTrip field to given value.

### HasNationalTrip

`func (o *RouteShort) HasNationalTrip() bool`

HasNationalTrip returns a boolean if a field has been set.

### GetBookable

`func (o *RouteShort) GetBookable() bool`

GetBookable returns the Bookable field if non-nil, zero value otherwise.

### GetBookableOk

`func (o *RouteShort) GetBookableOk() (*bool, bool)`

GetBookableOk returns a tuple with the Bookable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookable

`func (o *RouteShort) SetBookable(v bool)`

SetBookable sets Bookable field to given value.


### GetDelay

`func (o *RouteShort) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *RouteShort) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *RouteShort) SetDelay(v string)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *RouteShort) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetTravelTime

`func (o *RouteShort) GetTravelTime() string`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *RouteShort) GetTravelTimeOk() (*string, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *RouteShort) SetTravelTime(v string)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *RouteShort) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


