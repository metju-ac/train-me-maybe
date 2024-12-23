# SroPriceClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeatClassKey** | **string** | Seat class code. | 
**FreeSeatsCount** | **int32** |  | 
**Price** | **float32** |  | 
**PriceSource** | **string** |  | 
**NumberOfPassengers** | **int32** |  | 
**Conditions** | [**SroPriceConditions**](SroPriceConditions.md) |  | 

## Methods

### NewSroPriceClass

`func NewSroPriceClass(seatClassKey string, freeSeatsCount int32, price float32, priceSource string, numberOfPassengers int32, conditions SroPriceConditions, ) *SroPriceClass`

NewSroPriceClass instantiates a new SroPriceClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroPriceClassWithDefaults

`func NewSroPriceClassWithDefaults() *SroPriceClass`

NewSroPriceClassWithDefaults instantiates a new SroPriceClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeatClassKey

`func (o *SroPriceClass) GetSeatClassKey() string`

GetSeatClassKey returns the SeatClassKey field if non-nil, zero value otherwise.

### GetSeatClassKeyOk

`func (o *SroPriceClass) GetSeatClassKeyOk() (*string, bool)`

GetSeatClassKeyOk returns a tuple with the SeatClassKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClassKey

`func (o *SroPriceClass) SetSeatClassKey(v string)`

SetSeatClassKey sets SeatClassKey field to given value.


### GetFreeSeatsCount

`func (o *SroPriceClass) GetFreeSeatsCount() int32`

GetFreeSeatsCount returns the FreeSeatsCount field if non-nil, zero value otherwise.

### GetFreeSeatsCountOk

`func (o *SroPriceClass) GetFreeSeatsCountOk() (*int32, bool)`

GetFreeSeatsCountOk returns a tuple with the FreeSeatsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSeatsCount

`func (o *SroPriceClass) SetFreeSeatsCount(v int32)`

SetFreeSeatsCount sets FreeSeatsCount field to given value.


### GetPrice

`func (o *SroPriceClass) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SroPriceClass) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SroPriceClass) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetPriceSource

`func (o *SroPriceClass) GetPriceSource() string`

GetPriceSource returns the PriceSource field if non-nil, zero value otherwise.

### GetPriceSourceOk

`func (o *SroPriceClass) GetPriceSourceOk() (*string, bool)`

GetPriceSourceOk returns a tuple with the PriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSource

`func (o *SroPriceClass) SetPriceSource(v string)`

SetPriceSource sets PriceSource field to given value.


### GetNumberOfPassengers

`func (o *SroPriceClass) GetNumberOfPassengers() int32`

GetNumberOfPassengers returns the NumberOfPassengers field if non-nil, zero value otherwise.

### GetNumberOfPassengersOk

`func (o *SroPriceClass) GetNumberOfPassengersOk() (*int32, bool)`

GetNumberOfPassengersOk returns a tuple with the NumberOfPassengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPassengers

`func (o *SroPriceClass) SetNumberOfPassengers(v int32)`

SetNumberOfPassengers sets NumberOfPassengers field to given value.


### GetConditions

`func (o *SroPriceClass) GetConditions() SroPriceConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SroPriceClass) GetConditionsOk() (*SroPriceConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SroPriceClass) SetConditions(v SroPriceConditions)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


