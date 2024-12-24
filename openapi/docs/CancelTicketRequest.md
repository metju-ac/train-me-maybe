# CancelTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlHash** | **string** | Hash (MD5) of cancellation terms, agreed on by the customer | 
**RefundToOriginalSource** | Pointer to **bool** | TRUE if refund money after ticket cancel should return to original source, otherwise FALSE | [optional] 

## Methods

### NewCancelTicketRequest

`func NewCancelTicketRequest(controlHash string, ) *CancelTicketRequest`

NewCancelTicketRequest instantiates a new CancelTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelTicketRequestWithDefaults

`func NewCancelTicketRequestWithDefaults() *CancelTicketRequest`

NewCancelTicketRequestWithDefaults instantiates a new CancelTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlHash

`func (o *CancelTicketRequest) GetControlHash() string`

GetControlHash returns the ControlHash field if non-nil, zero value otherwise.

### GetControlHashOk

`func (o *CancelTicketRequest) GetControlHashOk() (*string, bool)`

GetControlHashOk returns a tuple with the ControlHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlHash

`func (o *CancelTicketRequest) SetControlHash(v string)`

SetControlHash sets ControlHash field to given value.


### GetRefundToOriginalSource

`func (o *CancelTicketRequest) GetRefundToOriginalSource() bool`

GetRefundToOriginalSource returns the RefundToOriginalSource field if non-nil, zero value otherwise.

### GetRefundToOriginalSourceOk

`func (o *CancelTicketRequest) GetRefundToOriginalSourceOk() (*bool, bool)`

GetRefundToOriginalSourceOk returns a tuple with the RefundToOriginalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundToOriginalSource

`func (o *CancelTicketRequest) SetRefundToOriginalSource(v bool)`

SetRefundToOriginalSource sets RefundToOriginalSource field to given value.

### HasRefundToOriginalSource

`func (o *CancelTicketRequest) HasRefundToOriginalSource() bool`

HasRefundToOriginalSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


