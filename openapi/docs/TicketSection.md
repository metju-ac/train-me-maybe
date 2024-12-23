# TicketSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Section** | [**Section**](Section.md) |  | 
**FixedSeatReservation** | **bool** | TRUE &#x3D;&gt; seat reservation; FALSE &#x3D;&gt; seat choosed while boarding | 
**SelectedSeats** | [**[]SelectedSeat100**](SelectedSeat100.md) |  | 
**Vehicles** | Pointer to [**[]Vehicle110**](Vehicle110.md) |  | [optional] 

## Methods

### NewTicketSection

`func NewTicketSection(section Section, fixedSeatReservation bool, selectedSeats []SelectedSeat100, ) *TicketSection`

NewTicketSection instantiates a new TicketSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketSectionWithDefaults

`func NewTicketSectionWithDefaults() *TicketSection`

NewTicketSectionWithDefaults instantiates a new TicketSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSection

`func (o *TicketSection) GetSection() Section`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *TicketSection) GetSectionOk() (*Section, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *TicketSection) SetSection(v Section)`

SetSection sets Section field to given value.


### GetFixedSeatReservation

`func (o *TicketSection) GetFixedSeatReservation() bool`

GetFixedSeatReservation returns the FixedSeatReservation field if non-nil, zero value otherwise.

### GetFixedSeatReservationOk

`func (o *TicketSection) GetFixedSeatReservationOk() (*bool, bool)`

GetFixedSeatReservationOk returns a tuple with the FixedSeatReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedSeatReservation

`func (o *TicketSection) SetFixedSeatReservation(v bool)`

SetFixedSeatReservation sets FixedSeatReservation field to given value.


### GetSelectedSeats

`func (o *TicketSection) GetSelectedSeats() []SelectedSeat100`

GetSelectedSeats returns the SelectedSeats field if non-nil, zero value otherwise.

### GetSelectedSeatsOk

`func (o *TicketSection) GetSelectedSeatsOk() (*[]SelectedSeat100, bool)`

GetSelectedSeatsOk returns a tuple with the SelectedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeats

`func (o *TicketSection) SetSelectedSeats(v []SelectedSeat100)`

SetSelectedSeats sets SelectedSeats field to given value.


### GetVehicles

`func (o *TicketSection) GetVehicles() []Vehicle110`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *TicketSection) GetVehiclesOk() (*[]Vehicle110, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *TicketSection) SetVehicles(v []Vehicle110)`

SetVehicles sets Vehicles field to given value.

### HasVehicles

`func (o *TicketSection) HasVehicles() bool`

HasVehicles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


