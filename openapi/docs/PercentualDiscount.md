# PercentualDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Percentage** | **int32** | Size of a percentual discount | 
**Amount** | **float32** | Discount sum in account currency | 
**Passengers** | **int32** | Maximal number of passengers | 
**FromLocation** | Pointer to **string** | Default departure country/city/station translations | [optional] 
**ToLocation** | Pointer to **string** | Default arrival country/city/station translations | [optional] 
**DateFrom** | Pointer to **time.Time** |  | [optional] 
**DateTo** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPercentualDiscount

`func NewPercentualDiscount(id int64, percentage int32, amount float32, passengers int32, ) *PercentualDiscount`

NewPercentualDiscount instantiates a new PercentualDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPercentualDiscountWithDefaults

`func NewPercentualDiscountWithDefaults() *PercentualDiscount`

NewPercentualDiscountWithDefaults instantiates a new PercentualDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PercentualDiscount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PercentualDiscount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PercentualDiscount) SetId(v int64)`

SetId sets Id field to given value.


### GetPercentage

`func (o *PercentualDiscount) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PercentualDiscount) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PercentualDiscount) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetAmount

`func (o *PercentualDiscount) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PercentualDiscount) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PercentualDiscount) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetPassengers

`func (o *PercentualDiscount) GetPassengers() int32`

GetPassengers returns the Passengers field if non-nil, zero value otherwise.

### GetPassengersOk

`func (o *PercentualDiscount) GetPassengersOk() (*int32, bool)`

GetPassengersOk returns a tuple with the Passengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengers

`func (o *PercentualDiscount) SetPassengers(v int32)`

SetPassengers sets Passengers field to given value.


### GetFromLocation

`func (o *PercentualDiscount) GetFromLocation() string`

GetFromLocation returns the FromLocation field if non-nil, zero value otherwise.

### GetFromLocationOk

`func (o *PercentualDiscount) GetFromLocationOk() (*string, bool)`

GetFromLocationOk returns a tuple with the FromLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLocation

`func (o *PercentualDiscount) SetFromLocation(v string)`

SetFromLocation sets FromLocation field to given value.

### HasFromLocation

`func (o *PercentualDiscount) HasFromLocation() bool`

HasFromLocation returns a boolean if a field has been set.

### GetToLocation

`func (o *PercentualDiscount) GetToLocation() string`

GetToLocation returns the ToLocation field if non-nil, zero value otherwise.

### GetToLocationOk

`func (o *PercentualDiscount) GetToLocationOk() (*string, bool)`

GetToLocationOk returns a tuple with the ToLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToLocation

`func (o *PercentualDiscount) SetToLocation(v string)`

SetToLocation sets ToLocation field to given value.

### HasToLocation

`func (o *PercentualDiscount) HasToLocation() bool`

HasToLocation returns a boolean if a field has been set.

### GetDateFrom

`func (o *PercentualDiscount) GetDateFrom() time.Time`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *PercentualDiscount) GetDateFromOk() (*time.Time, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *PercentualDiscount) SetDateFrom(v time.Time)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *PercentualDiscount) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *PercentualDiscount) GetDateTo() time.Time`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *PercentualDiscount) GetDateToOk() (*time.Time, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *PercentualDiscount) SetDateTo(v time.Time)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *PercentualDiscount) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


