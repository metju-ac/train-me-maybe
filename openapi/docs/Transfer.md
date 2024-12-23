# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromStationId** | **int64** | Station ID from which customer needs to transfer | 
**ToStationId** | **int64** | Station ID to which customer needs to transfer | 
**Type** | [**TransferType**](TransferType.md) |  | 
**CalculatedTransferTime** | [**TimePeriod**](TimePeriod.md) |  | 
**DeterminedTransferTime** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**Description** | Pointer to **string** | Transfer specification | [optional] 

## Methods

### NewTransfer

`func NewTransfer(fromStationId int64, toStationId int64, type_ TransferType, calculatedTransferTime TimePeriod, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromStationId

`func (o *Transfer) GetFromStationId() int64`

GetFromStationId returns the FromStationId field if non-nil, zero value otherwise.

### GetFromStationIdOk

`func (o *Transfer) GetFromStationIdOk() (*int64, bool)`

GetFromStationIdOk returns a tuple with the FromStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStationId

`func (o *Transfer) SetFromStationId(v int64)`

SetFromStationId sets FromStationId field to given value.


### GetToStationId

`func (o *Transfer) GetToStationId() int64`

GetToStationId returns the ToStationId field if non-nil, zero value otherwise.

### GetToStationIdOk

`func (o *Transfer) GetToStationIdOk() (*int64, bool)`

GetToStationIdOk returns a tuple with the ToStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStationId

`func (o *Transfer) SetToStationId(v int64)`

SetToStationId sets ToStationId field to given value.


### GetType

`func (o *Transfer) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transfer) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transfer) SetType(v TransferType)`

SetType sets Type field to given value.


### GetCalculatedTransferTime

`func (o *Transfer) GetCalculatedTransferTime() TimePeriod`

GetCalculatedTransferTime returns the CalculatedTransferTime field if non-nil, zero value otherwise.

### GetCalculatedTransferTimeOk

`func (o *Transfer) GetCalculatedTransferTimeOk() (*TimePeriod, bool)`

GetCalculatedTransferTimeOk returns a tuple with the CalculatedTransferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedTransferTime

`func (o *Transfer) SetCalculatedTransferTime(v TimePeriod)`

SetCalculatedTransferTime sets CalculatedTransferTime field to given value.


### GetDeterminedTransferTime

`func (o *Transfer) GetDeterminedTransferTime() TimePeriod`

GetDeterminedTransferTime returns the DeterminedTransferTime field if non-nil, zero value otherwise.

### GetDeterminedTransferTimeOk

`func (o *Transfer) GetDeterminedTransferTimeOk() (*TimePeriod, bool)`

GetDeterminedTransferTimeOk returns a tuple with the DeterminedTransferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterminedTransferTime

`func (o *Transfer) SetDeterminedTransferTime(v TimePeriod)`

SetDeterminedTransferTime sets DeterminedTransferTime field to given value.

### HasDeterminedTransferTime

`func (o *Transfer) HasDeterminedTransferTime() bool`

HasDeterminedTransferTime returns a boolean if a field has been set.

### GetDescription

`func (o *Transfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transfer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transfer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


