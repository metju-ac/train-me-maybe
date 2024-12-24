# Tariff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Tariff&#39;s key that is used in responses from other endpoints. | [optional] 
**Value** | Pointer to **string** | Tariff description. | [optional] 

## Methods

### NewTariff

`func NewTariff() *Tariff`

NewTariff instantiates a new Tariff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTariffWithDefaults

`func NewTariffWithDefaults() *Tariff`

NewTariffWithDefaults instantiates a new Tariff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Tariff) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Tariff) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Tariff) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Tariff) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *Tariff) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Tariff) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Tariff) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Tariff) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


