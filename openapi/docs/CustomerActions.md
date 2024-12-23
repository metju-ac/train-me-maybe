# CustomerActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowDetail** | **bool** |  | [default to false]
**Pay** | **bool** |  | [default to false]
**PayRemaining** | **bool** |  | [default to false]
**Evaluate** | **bool** |  | [default to false]
**Cancel** | **bool** |  | [default to false]
**Storno** | **bool** |  | [default to false]
**Rebook** | **bool** |  | [default to false]
**EditPassengers** | **bool** |  | [default to false]
**AdditionalServices** | **bool** |  | [default to false]
**SentToMail** | **bool** |  | [default to false]
**PrintTicket** | **bool** |  | [default to false]
**PrintInvoice** | **bool** |  | [default to false]

## Methods

### NewCustomerActions

`func NewCustomerActions(showDetail bool, pay bool, payRemaining bool, evaluate bool, cancel bool, storno bool, rebook bool, editPassengers bool, additionalServices bool, sentToMail bool, printTicket bool, printInvoice bool, ) *CustomerActions`

NewCustomerActions instantiates a new CustomerActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerActionsWithDefaults

`func NewCustomerActionsWithDefaults() *CustomerActions`

NewCustomerActionsWithDefaults instantiates a new CustomerActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowDetail

`func (o *CustomerActions) GetShowDetail() bool`

GetShowDetail returns the ShowDetail field if non-nil, zero value otherwise.

### GetShowDetailOk

`func (o *CustomerActions) GetShowDetailOk() (*bool, bool)`

GetShowDetailOk returns a tuple with the ShowDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDetail

`func (o *CustomerActions) SetShowDetail(v bool)`

SetShowDetail sets ShowDetail field to given value.


### GetPay

`func (o *CustomerActions) GetPay() bool`

GetPay returns the Pay field if non-nil, zero value otherwise.

### GetPayOk

`func (o *CustomerActions) GetPayOk() (*bool, bool)`

GetPayOk returns a tuple with the Pay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPay

`func (o *CustomerActions) SetPay(v bool)`

SetPay sets Pay field to given value.


### GetPayRemaining

`func (o *CustomerActions) GetPayRemaining() bool`

GetPayRemaining returns the PayRemaining field if non-nil, zero value otherwise.

### GetPayRemainingOk

`func (o *CustomerActions) GetPayRemainingOk() (*bool, bool)`

GetPayRemainingOk returns a tuple with the PayRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayRemaining

`func (o *CustomerActions) SetPayRemaining(v bool)`

SetPayRemaining sets PayRemaining field to given value.


### GetEvaluate

`func (o *CustomerActions) GetEvaluate() bool`

GetEvaluate returns the Evaluate field if non-nil, zero value otherwise.

### GetEvaluateOk

`func (o *CustomerActions) GetEvaluateOk() (*bool, bool)`

GetEvaluateOk returns a tuple with the Evaluate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluate

`func (o *CustomerActions) SetEvaluate(v bool)`

SetEvaluate sets Evaluate field to given value.


### GetCancel

`func (o *CustomerActions) GetCancel() bool`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *CustomerActions) GetCancelOk() (*bool, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *CustomerActions) SetCancel(v bool)`

SetCancel sets Cancel field to given value.


### GetStorno

`func (o *CustomerActions) GetStorno() bool`

GetStorno returns the Storno field if non-nil, zero value otherwise.

### GetStornoOk

`func (o *CustomerActions) GetStornoOk() (*bool, bool)`

GetStornoOk returns a tuple with the Storno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorno

`func (o *CustomerActions) SetStorno(v bool)`

SetStorno sets Storno field to given value.


### GetRebook

`func (o *CustomerActions) GetRebook() bool`

GetRebook returns the Rebook field if non-nil, zero value otherwise.

### GetRebookOk

`func (o *CustomerActions) GetRebookOk() (*bool, bool)`

GetRebookOk returns a tuple with the Rebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebook

`func (o *CustomerActions) SetRebook(v bool)`

SetRebook sets Rebook field to given value.


### GetEditPassengers

`func (o *CustomerActions) GetEditPassengers() bool`

GetEditPassengers returns the EditPassengers field if non-nil, zero value otherwise.

### GetEditPassengersOk

`func (o *CustomerActions) GetEditPassengersOk() (*bool, bool)`

GetEditPassengersOk returns a tuple with the EditPassengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPassengers

`func (o *CustomerActions) SetEditPassengers(v bool)`

SetEditPassengers sets EditPassengers field to given value.


### GetAdditionalServices

`func (o *CustomerActions) GetAdditionalServices() bool`

GetAdditionalServices returns the AdditionalServices field if non-nil, zero value otherwise.

### GetAdditionalServicesOk

`func (o *CustomerActions) GetAdditionalServicesOk() (*bool, bool)`

GetAdditionalServicesOk returns a tuple with the AdditionalServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalServices

`func (o *CustomerActions) SetAdditionalServices(v bool)`

SetAdditionalServices sets AdditionalServices field to given value.


### GetSentToMail

`func (o *CustomerActions) GetSentToMail() bool`

GetSentToMail returns the SentToMail field if non-nil, zero value otherwise.

### GetSentToMailOk

`func (o *CustomerActions) GetSentToMailOk() (*bool, bool)`

GetSentToMailOk returns a tuple with the SentToMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentToMail

`func (o *CustomerActions) SetSentToMail(v bool)`

SetSentToMail sets SentToMail field to given value.


### GetPrintTicket

`func (o *CustomerActions) GetPrintTicket() bool`

GetPrintTicket returns the PrintTicket field if non-nil, zero value otherwise.

### GetPrintTicketOk

`func (o *CustomerActions) GetPrintTicketOk() (*bool, bool)`

GetPrintTicketOk returns a tuple with the PrintTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintTicket

`func (o *CustomerActions) SetPrintTicket(v bool)`

SetPrintTicket sets PrintTicket field to given value.


### GetPrintInvoice

`func (o *CustomerActions) GetPrintInvoice() bool`

GetPrintInvoice returns the PrintInvoice field if non-nil, zero value otherwise.

### GetPrintInvoiceOk

`func (o *CustomerActions) GetPrintInvoiceOk() (*bool, bool)`

GetPrintInvoiceOk returns a tuple with the PrintInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintInvoice

`func (o *CustomerActions) SetPrintInvoice(v bool)`

SetPrintInvoice sets PrintInvoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


