# OrderedAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**IconUrl** | **string** |  | 
**InfoUrl** | Pointer to **string** |  | [optional] 
**InfoUrlLabel** | Pointer to **string** |  | [optional] 
**Price** | **float32** |  | 
**Currency** | [**Currency**](Currency.md) |  | [default to EUR]
**Conditions** | [**PriceConditions**](PriceConditions.md) |  | 

## Methods

### NewOrderedAddon

`func NewOrderedAddon(id int64, name string, description string, iconUrl string, price float32, currency Currency, conditions PriceConditions, ) *OrderedAddon`

NewOrderedAddon instantiates a new OrderedAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderedAddonWithDefaults

`func NewOrderedAddonWithDefaults() *OrderedAddon`

NewOrderedAddonWithDefaults instantiates a new OrderedAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderedAddon) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderedAddon) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderedAddon) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *OrderedAddon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderedAddon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderedAddon) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OrderedAddon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderedAddon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderedAddon) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIconUrl

`func (o *OrderedAddon) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *OrderedAddon) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *OrderedAddon) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetInfoUrl

`func (o *OrderedAddon) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *OrderedAddon) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *OrderedAddon) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *OrderedAddon) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### GetInfoUrlLabel

`func (o *OrderedAddon) GetInfoUrlLabel() string`

GetInfoUrlLabel returns the InfoUrlLabel field if non-nil, zero value otherwise.

### GetInfoUrlLabelOk

`func (o *OrderedAddon) GetInfoUrlLabelOk() (*string, bool)`

GetInfoUrlLabelOk returns a tuple with the InfoUrlLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrlLabel

`func (o *OrderedAddon) SetInfoUrlLabel(v string)`

SetInfoUrlLabel sets InfoUrlLabel field to given value.

### HasInfoUrlLabel

`func (o *OrderedAddon) HasInfoUrlLabel() bool`

HasInfoUrlLabel returns a boolean if a field has been set.

### GetPrice

`func (o *OrderedAddon) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderedAddon) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderedAddon) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCurrency

`func (o *OrderedAddon) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderedAddon) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderedAddon) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetConditions

`func (o *OrderedAddon) GetConditions() PriceConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *OrderedAddon) GetConditionsOk() (*PriceConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *OrderedAddon) SetConditions(v PriceConditions)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


