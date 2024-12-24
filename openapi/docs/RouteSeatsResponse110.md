# RouteSeatsResponse110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionId** | **int64** |  | 
**FixedSeatReservation** | **bool** | TRUE &#x3D;&gt; seat reservation; FALSE &#x3D;&gt; seat choosed while boarding | 
**Vehicles** | Pointer to [**[]Vehicle110**](Vehicle110.md) |  | [optional] 
**SelectedSeats** | Pointer to [**[]SelectedSeat100**](SelectedSeat100.md) |  | [optional] 

## Methods

### NewRouteSeatsResponse110

`func NewRouteSeatsResponse110(sectionId int64, fixedSeatReservation bool, ) *RouteSeatsResponse110`

NewRouteSeatsResponse110 instantiates a new RouteSeatsResponse110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteSeatsResponse110WithDefaults

`func NewRouteSeatsResponse110WithDefaults() *RouteSeatsResponse110`

NewRouteSeatsResponse110WithDefaults instantiates a new RouteSeatsResponse110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionId

`func (o *RouteSeatsResponse110) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *RouteSeatsResponse110) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *RouteSeatsResponse110) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.


### GetFixedSeatReservation

`func (o *RouteSeatsResponse110) GetFixedSeatReservation() bool`

GetFixedSeatReservation returns the FixedSeatReservation field if non-nil, zero value otherwise.

### GetFixedSeatReservationOk

`func (o *RouteSeatsResponse110) GetFixedSeatReservationOk() (*bool, bool)`

GetFixedSeatReservationOk returns a tuple with the FixedSeatReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedSeatReservation

`func (o *RouteSeatsResponse110) SetFixedSeatReservation(v bool)`

SetFixedSeatReservation sets FixedSeatReservation field to given value.


### GetVehicles

`func (o *RouteSeatsResponse110) GetVehicles() []Vehicle110`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *RouteSeatsResponse110) GetVehiclesOk() (*[]Vehicle110, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *RouteSeatsResponse110) SetVehicles(v []Vehicle110)`

SetVehicles sets Vehicles field to given value.

### HasVehicles

`func (o *RouteSeatsResponse110) HasVehicles() bool`

HasVehicles returns a boolean if a field has been set.

### GetSelectedSeats

`func (o *RouteSeatsResponse110) GetSelectedSeats() []SelectedSeat100`

GetSelectedSeats returns the SelectedSeats field if non-nil, zero value otherwise.

### GetSelectedSeatsOk

`func (o *RouteSeatsResponse110) GetSelectedSeatsOk() (*[]SelectedSeat100, bool)`

GetSelectedSeatsOk returns a tuple with the SelectedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeats

`func (o *RouteSeatsResponse110) SetSelectedSeats(v []SelectedSeat100)`

SetSelectedSeats sets SelectedSeats field to given value.

### HasSelectedSeats

`func (o *RouteSeatsResponse110) HasSelectedSeats() bool`

HasSelectedSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


