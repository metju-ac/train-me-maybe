# VehicleDeck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** |  | 
**Name** | **string** |  | 
**LayoutURL** | Pointer to **string** |  | [optional] 
**HorizontalLayoutURL** | Pointer to **string** |  | [optional] 
**FreeSeats** | [**[]Seat**](Seat.md) |  | 
**OccupiedSeats** | [**[]Seat**](Seat.md) |  | 

## Methods

### NewVehicleDeck

`func NewVehicleDeck(number int32, name string, freeSeats []Seat, occupiedSeats []Seat, ) *VehicleDeck`

NewVehicleDeck instantiates a new VehicleDeck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleDeckWithDefaults

`func NewVehicleDeckWithDefaults() *VehicleDeck`

NewVehicleDeckWithDefaults instantiates a new VehicleDeck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *VehicleDeck) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VehicleDeck) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VehicleDeck) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetName

`func (o *VehicleDeck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleDeck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleDeck) SetName(v string)`

SetName sets Name field to given value.


### GetLayoutURL

`func (o *VehicleDeck) GetLayoutURL() string`

GetLayoutURL returns the LayoutURL field if non-nil, zero value otherwise.

### GetLayoutURLOk

`func (o *VehicleDeck) GetLayoutURLOk() (*string, bool)`

GetLayoutURLOk returns a tuple with the LayoutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutURL

`func (o *VehicleDeck) SetLayoutURL(v string)`

SetLayoutURL sets LayoutURL field to given value.

### HasLayoutURL

`func (o *VehicleDeck) HasLayoutURL() bool`

HasLayoutURL returns a boolean if a field has been set.

### GetHorizontalLayoutURL

`func (o *VehicleDeck) GetHorizontalLayoutURL() string`

GetHorizontalLayoutURL returns the HorizontalLayoutURL field if non-nil, zero value otherwise.

### GetHorizontalLayoutURLOk

`func (o *VehicleDeck) GetHorizontalLayoutURLOk() (*string, bool)`

GetHorizontalLayoutURLOk returns a tuple with the HorizontalLayoutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalLayoutURL

`func (o *VehicleDeck) SetHorizontalLayoutURL(v string)`

SetHorizontalLayoutURL sets HorizontalLayoutURL field to given value.

### HasHorizontalLayoutURL

`func (o *VehicleDeck) HasHorizontalLayoutURL() bool`

HasHorizontalLayoutURL returns a boolean if a field has been set.

### GetFreeSeats

`func (o *VehicleDeck) GetFreeSeats() []Seat`

GetFreeSeats returns the FreeSeats field if non-nil, zero value otherwise.

### GetFreeSeatsOk

`func (o *VehicleDeck) GetFreeSeatsOk() (*[]Seat, bool)`

GetFreeSeatsOk returns a tuple with the FreeSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSeats

`func (o *VehicleDeck) SetFreeSeats(v []Seat)`

SetFreeSeats sets FreeSeats field to given value.


### GetOccupiedSeats

`func (o *VehicleDeck) GetOccupiedSeats() []Seat`

GetOccupiedSeats returns the OccupiedSeats field if non-nil, zero value otherwise.

### GetOccupiedSeatsOk

`func (o *VehicleDeck) GetOccupiedSeatsOk() (*[]Seat, bool)`

GetOccupiedSeatsOk returns a tuple with the OccupiedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupiedSeats

`func (o *VehicleDeck) SetOccupiedSeats(v []Seat)`

SetOccupiedSeats sets OccupiedSeats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


