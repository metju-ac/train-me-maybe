# TransfersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | **string** | Accompanying informations about transfers | 
**Transfers** | [**[]Transfer**](Transfer.md) |  | 

## Methods

### NewTransfersInfo

`func NewTransfersInfo(info string, transfers []Transfer, ) *TransfersInfo`

NewTransfersInfo instantiates a new TransfersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransfersInfoWithDefaults

`func NewTransfersInfoWithDefaults() *TransfersInfo`

NewTransfersInfoWithDefaults instantiates a new TransfersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *TransfersInfo) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TransfersInfo) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TransfersInfo) SetInfo(v string)`

SetInfo sets Info field to given value.


### GetTransfers

`func (o *TransfersInfo) GetTransfers() []Transfer`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *TransfersInfo) GetTransfersOk() (*[]Transfer, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *TransfersInfo) SetTransfers(v []Transfer)`

SetTransfers sets Transfers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


