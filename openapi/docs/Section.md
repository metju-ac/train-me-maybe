# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**VehicleStandardKey** | **string** |  | 
**Support** | **bool** | &#x60;true&#x60; if it is a support connnection or &#x60;false&#x60; if this is a RegioJet connection. | 
**SupportCode** | Pointer to **string** | Codes of individuals support connections on a single route. | [optional] 
**VehicleType** | **string** |  | 
**FixedSeatReservation** | **bool** | &#x60;true&#x60; if seat can be chosen during reservation or &#x60;false&#x60; if seat is chosen during boarding. | 
**Line** | [**Line**](Line.md) |  | 
**DepartureStationId** | **int64** |  | 
**DepartureStationName** | Pointer to **string** |  | [optional] 
**DepartureCityId** | **int64** |  | 
**DepartureCityName** | Pointer to **string** |  | [optional] 
**DepartureTime** | **time.Time** |  | 
**DeparturePlatform** | Pointer to **string** |  | [optional] 
**ArrivalStationId** | **int64** |  | 
**ArrivalStationName** | Pointer to **string** |  | [optional] 
**ArrivalCityId** | **int64** |  | 
**ArrivalCityName** | Pointer to **string** |  | [optional] 
**ArrivalTime** | **time.Time** |  | 
**ArrivalPlatform** | Pointer to **string** |  | [optional] 
**CarrierId** | Pointer to **int64** | Company Id | [optional] 
**FreeSeatsCount** | Pointer to **int32** |  | [optional] 
**Notices** | Pointer to **[]string** | Extraordinary event on route / traffic limitation | [optional] 
**Services** | Pointer to **[]string** | Available services on this connection | [optional] 
**Delay** | Pointer to **string** | Textual information about the delay on a given section | [optional] 
**TravelTime** | **string** | Textual information about the travel time on a given section | 

## Methods

### NewSection

`func NewSection(id int64, vehicleStandardKey string, support bool, vehicleType string, fixedSeatReservation bool, line Line, departureStationId int64, departureCityId int64, departureTime time.Time, arrivalStationId int64, arrivalCityId int64, arrivalTime time.Time, travelTime string, ) *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Section) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Section) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Section) SetId(v int64)`

SetId sets Id field to given value.


### GetVehicleStandardKey

`func (o *Section) GetVehicleStandardKey() string`

GetVehicleStandardKey returns the VehicleStandardKey field if non-nil, zero value otherwise.

### GetVehicleStandardKeyOk

`func (o *Section) GetVehicleStandardKeyOk() (*string, bool)`

GetVehicleStandardKeyOk returns a tuple with the VehicleStandardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleStandardKey

`func (o *Section) SetVehicleStandardKey(v string)`

SetVehicleStandardKey sets VehicleStandardKey field to given value.


### GetSupport

`func (o *Section) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *Section) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *Section) SetSupport(v bool)`

SetSupport sets Support field to given value.


### GetSupportCode

`func (o *Section) GetSupportCode() string`

GetSupportCode returns the SupportCode field if non-nil, zero value otherwise.

### GetSupportCodeOk

`func (o *Section) GetSupportCodeOk() (*string, bool)`

GetSupportCodeOk returns a tuple with the SupportCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCode

`func (o *Section) SetSupportCode(v string)`

SetSupportCode sets SupportCode field to given value.

### HasSupportCode

`func (o *Section) HasSupportCode() bool`

HasSupportCode returns a boolean if a field has been set.

### GetVehicleType

`func (o *Section) GetVehicleType() string`

GetVehicleType returns the VehicleType field if non-nil, zero value otherwise.

### GetVehicleTypeOk

`func (o *Section) GetVehicleTypeOk() (*string, bool)`

GetVehicleTypeOk returns a tuple with the VehicleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleType

`func (o *Section) SetVehicleType(v string)`

SetVehicleType sets VehicleType field to given value.


### GetFixedSeatReservation

`func (o *Section) GetFixedSeatReservation() bool`

GetFixedSeatReservation returns the FixedSeatReservation field if non-nil, zero value otherwise.

### GetFixedSeatReservationOk

`func (o *Section) GetFixedSeatReservationOk() (*bool, bool)`

GetFixedSeatReservationOk returns a tuple with the FixedSeatReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedSeatReservation

`func (o *Section) SetFixedSeatReservation(v bool)`

SetFixedSeatReservation sets FixedSeatReservation field to given value.


### GetLine

`func (o *Section) GetLine() Line`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Section) GetLineOk() (*Line, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Section) SetLine(v Line)`

SetLine sets Line field to given value.


### GetDepartureStationId

`func (o *Section) GetDepartureStationId() int64`

GetDepartureStationId returns the DepartureStationId field if non-nil, zero value otherwise.

### GetDepartureStationIdOk

`func (o *Section) GetDepartureStationIdOk() (*int64, bool)`

GetDepartureStationIdOk returns a tuple with the DepartureStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureStationId

`func (o *Section) SetDepartureStationId(v int64)`

SetDepartureStationId sets DepartureStationId field to given value.


### GetDepartureStationName

`func (o *Section) GetDepartureStationName() string`

GetDepartureStationName returns the DepartureStationName field if non-nil, zero value otherwise.

### GetDepartureStationNameOk

`func (o *Section) GetDepartureStationNameOk() (*string, bool)`

GetDepartureStationNameOk returns a tuple with the DepartureStationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureStationName

`func (o *Section) SetDepartureStationName(v string)`

SetDepartureStationName sets DepartureStationName field to given value.

### HasDepartureStationName

`func (o *Section) HasDepartureStationName() bool`

HasDepartureStationName returns a boolean if a field has been set.

### GetDepartureCityId

`func (o *Section) GetDepartureCityId() int64`

GetDepartureCityId returns the DepartureCityId field if non-nil, zero value otherwise.

### GetDepartureCityIdOk

`func (o *Section) GetDepartureCityIdOk() (*int64, bool)`

GetDepartureCityIdOk returns a tuple with the DepartureCityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureCityId

`func (o *Section) SetDepartureCityId(v int64)`

SetDepartureCityId sets DepartureCityId field to given value.


### GetDepartureCityName

`func (o *Section) GetDepartureCityName() string`

GetDepartureCityName returns the DepartureCityName field if non-nil, zero value otherwise.

### GetDepartureCityNameOk

`func (o *Section) GetDepartureCityNameOk() (*string, bool)`

GetDepartureCityNameOk returns a tuple with the DepartureCityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureCityName

`func (o *Section) SetDepartureCityName(v string)`

SetDepartureCityName sets DepartureCityName field to given value.

### HasDepartureCityName

`func (o *Section) HasDepartureCityName() bool`

HasDepartureCityName returns a boolean if a field has been set.

### GetDepartureTime

`func (o *Section) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *Section) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *Section) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.


### GetDeparturePlatform

`func (o *Section) GetDeparturePlatform() string`

GetDeparturePlatform returns the DeparturePlatform field if non-nil, zero value otherwise.

### GetDeparturePlatformOk

`func (o *Section) GetDeparturePlatformOk() (*string, bool)`

GetDeparturePlatformOk returns a tuple with the DeparturePlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeparturePlatform

`func (o *Section) SetDeparturePlatform(v string)`

SetDeparturePlatform sets DeparturePlatform field to given value.

### HasDeparturePlatform

`func (o *Section) HasDeparturePlatform() bool`

HasDeparturePlatform returns a boolean if a field has been set.

### GetArrivalStationId

`func (o *Section) GetArrivalStationId() int64`

GetArrivalStationId returns the ArrivalStationId field if non-nil, zero value otherwise.

### GetArrivalStationIdOk

`func (o *Section) GetArrivalStationIdOk() (*int64, bool)`

GetArrivalStationIdOk returns a tuple with the ArrivalStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalStationId

`func (o *Section) SetArrivalStationId(v int64)`

SetArrivalStationId sets ArrivalStationId field to given value.


### GetArrivalStationName

`func (o *Section) GetArrivalStationName() string`

GetArrivalStationName returns the ArrivalStationName field if non-nil, zero value otherwise.

### GetArrivalStationNameOk

`func (o *Section) GetArrivalStationNameOk() (*string, bool)`

GetArrivalStationNameOk returns a tuple with the ArrivalStationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalStationName

`func (o *Section) SetArrivalStationName(v string)`

SetArrivalStationName sets ArrivalStationName field to given value.

### HasArrivalStationName

`func (o *Section) HasArrivalStationName() bool`

HasArrivalStationName returns a boolean if a field has been set.

### GetArrivalCityId

`func (o *Section) GetArrivalCityId() int64`

GetArrivalCityId returns the ArrivalCityId field if non-nil, zero value otherwise.

### GetArrivalCityIdOk

`func (o *Section) GetArrivalCityIdOk() (*int64, bool)`

GetArrivalCityIdOk returns a tuple with the ArrivalCityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalCityId

`func (o *Section) SetArrivalCityId(v int64)`

SetArrivalCityId sets ArrivalCityId field to given value.


### GetArrivalCityName

`func (o *Section) GetArrivalCityName() string`

GetArrivalCityName returns the ArrivalCityName field if non-nil, zero value otherwise.

### GetArrivalCityNameOk

`func (o *Section) GetArrivalCityNameOk() (*string, bool)`

GetArrivalCityNameOk returns a tuple with the ArrivalCityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalCityName

`func (o *Section) SetArrivalCityName(v string)`

SetArrivalCityName sets ArrivalCityName field to given value.

### HasArrivalCityName

`func (o *Section) HasArrivalCityName() bool`

HasArrivalCityName returns a boolean if a field has been set.

### GetArrivalTime

`func (o *Section) GetArrivalTime() time.Time`

GetArrivalTime returns the ArrivalTime field if non-nil, zero value otherwise.

### GetArrivalTimeOk

`func (o *Section) GetArrivalTimeOk() (*time.Time, bool)`

GetArrivalTimeOk returns a tuple with the ArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalTime

`func (o *Section) SetArrivalTime(v time.Time)`

SetArrivalTime sets ArrivalTime field to given value.


### GetArrivalPlatform

`func (o *Section) GetArrivalPlatform() string`

GetArrivalPlatform returns the ArrivalPlatform field if non-nil, zero value otherwise.

### GetArrivalPlatformOk

`func (o *Section) GetArrivalPlatformOk() (*string, bool)`

GetArrivalPlatformOk returns a tuple with the ArrivalPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalPlatform

`func (o *Section) SetArrivalPlatform(v string)`

SetArrivalPlatform sets ArrivalPlatform field to given value.

### HasArrivalPlatform

`func (o *Section) HasArrivalPlatform() bool`

HasArrivalPlatform returns a boolean if a field has been set.

### GetCarrierId

`func (o *Section) GetCarrierId() int64`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *Section) GetCarrierIdOk() (*int64, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *Section) SetCarrierId(v int64)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *Section) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetFreeSeatsCount

`func (o *Section) GetFreeSeatsCount() int32`

GetFreeSeatsCount returns the FreeSeatsCount field if non-nil, zero value otherwise.

### GetFreeSeatsCountOk

`func (o *Section) GetFreeSeatsCountOk() (*int32, bool)`

GetFreeSeatsCountOk returns a tuple with the FreeSeatsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSeatsCount

`func (o *Section) SetFreeSeatsCount(v int32)`

SetFreeSeatsCount sets FreeSeatsCount field to given value.

### HasFreeSeatsCount

`func (o *Section) HasFreeSeatsCount() bool`

HasFreeSeatsCount returns a boolean if a field has been set.

### GetNotices

`func (o *Section) GetNotices() []string`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *Section) GetNoticesOk() (*[]string, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *Section) SetNotices(v []string)`

SetNotices sets Notices field to given value.

### HasNotices

`func (o *Section) HasNotices() bool`

HasNotices returns a boolean if a field has been set.

### GetServices

`func (o *Section) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Section) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Section) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Section) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetDelay

`func (o *Section) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *Section) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *Section) SetDelay(v string)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *Section) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetTravelTime

`func (o *Section) GetTravelTime() string`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *Section) GetTravelTimeOk() (*string, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *Section) SetTravelTime(v string)`

SetTravelTime sets TravelTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


