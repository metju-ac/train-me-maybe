# CancelCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Cancel charge description. | [optional] 
**Amount** | **float32** | Amount of money in chosen currency taken from ticket/addon/... which will not be refunded. | 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] [default to EUR]
**Percent** | Pointer to **int32** | Number of percents of original price that will be withheld for cancel. | [optional] 

## Methods

### NewCancelCharge

`func NewCancelCharge(amount float32, ) *CancelCharge`

NewCancelCharge instantiates a new CancelCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelChargeWithDefaults

`func NewCancelChargeWithDefaults() *CancelCharge`

NewCancelChargeWithDefaults instantiates a new CancelCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CancelCharge) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CancelCharge) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CancelCharge) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CancelCharge) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAmount

`func (o *CancelCharge) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CancelCharge) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CancelCharge) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CancelCharge) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CancelCharge) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CancelCharge) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CancelCharge) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPercent

`func (o *CancelCharge) GetPercent() int32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *CancelCharge) GetPercentOk() (*int32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *CancelCharge) SetPercent(v int32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *CancelCharge) HasPercent() bool`

HasPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


