# RouteSeatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionId** | **int64** |  | 
**FixedSeatReservation** | **bool** | TRUE &#x3D;&gt; seat reservation; FALSE &#x3D;&gt; seat choosed while boarding | 
**Vehicles** | Pointer to [**[]Vehicle**](Vehicle.md) |  | [optional] 
**SelectedSeats** | Pointer to [**[]SelectedSeat**](SelectedSeat.md) |  | [optional] 

## Methods

### NewRouteSeatsResponse

`func NewRouteSeatsResponse(sectionId int64, fixedSeatReservation bool, ) *RouteSeatsResponse`

NewRouteSeatsResponse instantiates a new RouteSeatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteSeatsResponseWithDefaults

`func NewRouteSeatsResponseWithDefaults() *RouteSeatsResponse`

NewRouteSeatsResponseWithDefaults instantiates a new RouteSeatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionId

`func (o *RouteSeatsResponse) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *RouteSeatsResponse) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *RouteSeatsResponse) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.


### GetFixedSeatReservation

`func (o *RouteSeatsResponse) GetFixedSeatReservation() bool`

GetFixedSeatReservation returns the FixedSeatReservation field if non-nil, zero value otherwise.

### GetFixedSeatReservationOk

`func (o *RouteSeatsResponse) GetFixedSeatReservationOk() (*bool, bool)`

GetFixedSeatReservationOk returns a tuple with the FixedSeatReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedSeatReservation

`func (o *RouteSeatsResponse) SetFixedSeatReservation(v bool)`

SetFixedSeatReservation sets FixedSeatReservation field to given value.


### GetVehicles

`func (o *RouteSeatsResponse) GetVehicles() []Vehicle`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *RouteSeatsResponse) GetVehiclesOk() (*[]Vehicle, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *RouteSeatsResponse) SetVehicles(v []Vehicle)`

SetVehicles sets Vehicles field to given value.

### HasVehicles

`func (o *RouteSeatsResponse) HasVehicles() bool`

HasVehicles returns a boolean if a field has been set.

### GetSelectedSeats

`func (o *RouteSeatsResponse) GetSelectedSeats() []SelectedSeat`

GetSelectedSeats returns the SelectedSeats field if non-nil, zero value otherwise.

### GetSelectedSeatsOk

`func (o *RouteSeatsResponse) GetSelectedSeatsOk() (*[]SelectedSeat, bool)`

GetSelectedSeatsOk returns a tuple with the SelectedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeats

`func (o *RouteSeatsResponse) SetSelectedSeats(v []SelectedSeat)`

SetSelectedSeats sets SelectedSeats field to given value.

### HasSelectedSeats

`func (o *RouteSeatsResponse) HasSelectedSeats() bool`

HasSelectedSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


