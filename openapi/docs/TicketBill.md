# TicketBill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**Currency** | [**Currency**](Currency.md) |  | [default to EUR]
**Label** | **string** |  | 

## Methods

### NewTicketBill

`func NewTicketBill(amount float32, currency Currency, label string, ) *TicketBill`

NewTicketBill instantiates a new TicketBill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketBillWithDefaults

`func NewTicketBillWithDefaults() *TicketBill`

NewTicketBillWithDefaults instantiates a new TicketBill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TicketBill) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TicketBill) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TicketBill) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *TicketBill) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TicketBill) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TicketBill) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetLabel

`func (o *TicketBill) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TicketBill) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TicketBill) SetLabel(v string)`

SetLabel sets Label field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


