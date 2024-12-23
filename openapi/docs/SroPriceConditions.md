# SroPriceConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptions** | [**SroPriceConditionsDescriptions**](SroPriceConditionsDescriptions.md) |  | 
**RefundToOriginalSourcePossible** | **bool** |  | 
**CancelCharge** | **float32** |  | 
**CancelCharges** | [**[]SroCancelCharge**](SroCancelCharge.md) |  | 

## Methods

### NewSroPriceConditions

`func NewSroPriceConditions(descriptions SroPriceConditionsDescriptions, refundToOriginalSourcePossible bool, cancelCharge float32, cancelCharges []SroCancelCharge, ) *SroPriceConditions`

NewSroPriceConditions instantiates a new SroPriceConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroPriceConditionsWithDefaults

`func NewSroPriceConditionsWithDefaults() *SroPriceConditions`

NewSroPriceConditionsWithDefaults instantiates a new SroPriceConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptions

`func (o *SroPriceConditions) GetDescriptions() SroPriceConditionsDescriptions`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *SroPriceConditions) GetDescriptionsOk() (*SroPriceConditionsDescriptions, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *SroPriceConditions) SetDescriptions(v SroPriceConditionsDescriptions)`

SetDescriptions sets Descriptions field to given value.


### GetRefundToOriginalSourcePossible

`func (o *SroPriceConditions) GetRefundToOriginalSourcePossible() bool`

GetRefundToOriginalSourcePossible returns the RefundToOriginalSourcePossible field if non-nil, zero value otherwise.

### GetRefundToOriginalSourcePossibleOk

`func (o *SroPriceConditions) GetRefundToOriginalSourcePossibleOk() (*bool, bool)`

GetRefundToOriginalSourcePossibleOk returns a tuple with the RefundToOriginalSourcePossible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundToOriginalSourcePossible

`func (o *SroPriceConditions) SetRefundToOriginalSourcePossible(v bool)`

SetRefundToOriginalSourcePossible sets RefundToOriginalSourcePossible field to given value.


### GetCancelCharge

`func (o *SroPriceConditions) GetCancelCharge() float32`

GetCancelCharge returns the CancelCharge field if non-nil, zero value otherwise.

### GetCancelChargeOk

`func (o *SroPriceConditions) GetCancelChargeOk() (*float32, bool)`

GetCancelChargeOk returns a tuple with the CancelCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCharge

`func (o *SroPriceConditions) SetCancelCharge(v float32)`

SetCancelCharge sets CancelCharge field to given value.


### GetCancelCharges

`func (o *SroPriceConditions) GetCancelCharges() []SroCancelCharge`

GetCancelCharges returns the CancelCharges field if non-nil, zero value otherwise.

### GetCancelChargesOk

`func (o *SroPriceConditions) GetCancelChargesOk() (*[]SroCancelCharge, bool)`

GetCancelChargesOk returns a tuple with the CancelCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelCharges

`func (o *SroPriceConditions) SetCancelCharges(v []SroCancelCharge)`

SetCancelCharges sets CancelCharges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


