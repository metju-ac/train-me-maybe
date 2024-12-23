# PriceConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Control MD5 hash transformed from conditions content (used for current terms visibility check). | 
**Descriptions** | [**PriceConditionsDescriptions**](PriceConditionsDescriptions.md) |  | 
**RefundToOriginalSourcePossible** | Pointer to **bool** | States that if its possible to refund money to original money source | [optional] 
**CancelCharge** | Pointer to **float32** | Sum of all cancel charges in chosen currency. | [optional] 
**CancelCharges** | Pointer to [**[]CancelCharge**](CancelCharge.md) |  | [optional] 

## Methods

### NewPriceConditions

`func NewPriceConditions(code string, descriptions PriceConditionsDescriptions, ) *PriceConditions`

NewPriceConditions instantiates a new PriceConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceConditionsWithDefaults

`func NewPriceConditionsWithDefaults() *PriceConditions`

NewPriceConditionsWithDefaults instantiates a new PriceConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PriceConditions) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PriceConditions) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PriceConditions) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescriptions

`func (o *PriceConditions) GetDescriptions() PriceConditionsDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *PriceConditions) GetDescriptionsOk() (*PriceConditionsDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *PriceConditions) SetDescriptions(v PriceConditionsDescriptions)`

SetDescriptions sets Descriptions field to given value.


### GetRefundToOriginalSourcePossible

`func (o *PriceConditions) GetRefundToOriginalSourcePossible() bool`

GetRefundToOriginalSourcePossible returns the RefundToOriginalSourcePossible field if non-nil, zero value otherwise.

### GetRefundToOriginalSourcePossibleOk

`func (o *PriceConditions) GetRefundToOriginalSourcePossibleOk() (*bool, bool)`

GetRefundToOriginalSourcePossibleOk returns a tuple with the RefundToOriginalSourcePossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundToOriginalSourcePossible

`func (o *PriceConditions) SetRefundToOriginalSourcePossible(v bool)`

SetRefundToOriginalSourcePossible sets RefundToOriginalSourcePossible field to given value.

### HasRefundToOriginalSourcePossible

`func (o *PriceConditions) HasRefundToOriginalSourcePossible() bool`

HasRefundToOriginalSourcePossible returns a boolean if a field has been set.

### GetCancelCharge

`func (o *PriceConditions) GetCancelCharge() float32`

GetCancelCharge returns the CancelCharge field if non-nil, zero value otherwise.

### GetCancelChargeOk

`func (o *PriceConditions) GetCancelChargeOk() (*float32, bool)`

GetCancelChargeOk returns a tuple with the CancelCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCharge

`func (o *PriceConditions) SetCancelCharge(v float32)`

SetCancelCharge sets CancelCharge field to given value.

### HasCancelCharge

`func (o *PriceConditions) HasCancelCharge() bool`

HasCancelCharge returns a boolean if a field has been set.

### GetCancelCharges

`func (o *PriceConditions) GetCancelCharges() []CancelCharge`

GetCancelCharges returns the CancelCharges field if non-nil, zero value otherwise.

### GetCancelChargesOk

`func (o *PriceConditions) GetCancelChargesOk() (*[]CancelCharge, bool)`

GetCancelChargesOk returns a tuple with the CancelCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCharges

`func (o *PriceConditions) SetCancelCharges(v []CancelCharge)`

SetCancelCharges sets CancelCharges field to given value.

### HasCancelCharges

`func (o *PriceConditions) HasCancelCharges() bool`

HasCancelCharges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


