# SelectedAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonId** | **int64** |  | 
**Count** | **int32** |  | 
**Price** | **float32** |  | 
**Currency** | [**Currency**](Currency.md) |  | [default to EUR]
**ConditionsHash** | **string** | Hash (MD5) of cancellation terms, agreed on by the customer | 

## Methods

### NewSelectedAddon

`func NewSelectedAddon(addonId int64, count int32, price float32, currency Currency, conditionsHash string, ) *SelectedAddon`

NewSelectedAddon instantiates a new SelectedAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedAddonWithDefaults

`func NewSelectedAddonWithDefaults() *SelectedAddon`

NewSelectedAddonWithDefaults instantiates a new SelectedAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonId

`func (o *SelectedAddon) GetAddonId() int64`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *SelectedAddon) GetAddonIdOk() (*int64, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *SelectedAddon) SetAddonId(v int64)`

SetAddonId sets AddonId field to given value.


### GetCount

`func (o *SelectedAddon) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SelectedAddon) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SelectedAddon) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPrice

`func (o *SelectedAddon) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SelectedAddon) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SelectedAddon) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *SelectedAddon) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SelectedAddon) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SelectedAddon) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetConditionsHash

`func (o *SelectedAddon) GetConditionsHash() string`

GetConditionsHash returns the ConditionsHash field if non-nil, zero value otherwise.

### GetConditionsHashOk

`func (o *SelectedAddon) GetConditionsHashOk() (*string, bool)`

GetConditionsHashOk returns a tuple with the ConditionsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsHash

`func (o *SelectedAddon) SetConditionsHash(v string)`

SetConditionsHash sets ConditionsHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


