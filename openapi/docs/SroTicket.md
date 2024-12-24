# SroTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SroTicketId** | **int64** |  | 
**Price** | **float32** | Final price for ticket, addons etc. | 
**Unpaid** | **float32** | Final price to be paid | 
**Currency** | [**Currency**](Currency.md) |  | [default to EUR]
**State** | [**TicketState**](TicketState.md) |  | 
**SeatClassKey** | **string** |  | 
**Conditions** | [**SroConditions**](SroConditions.md) |  | 
**CustomerActions** | [**CustomerActions**](CustomerActions.md) |  | 
**RouteSections** | [**[]TicketSection**](TicketSection.md) |  | 
**PaymentId** | **int64** |  | 
**Bills** | [**[]TicketBill**](TicketBill.md) |  | 
**PassengersInfo** | [**SroPassengersInfo**](SroPassengersInfo.md) |  | 
**Delay** | **string** | Textual information about the first delay on the route. | 
**TravelTime** | **string** | Textual information about the travel time on a given section. | 
**EstimatedArrivalTime** | **time.Time** |  | 
**AffiliateTicket** | **bool** | Was the ticket created by an affiliate partner? | 

## Methods

### NewSroTicket

`func NewSroTicket(sroTicketId int64, price float32, unpaid float32, currency Currency, state TicketState, seatClassKey string, conditions SroConditions, customerActions CustomerActions, routeSections []TicketSection, paymentId int64, bills []TicketBill, passengersInfo SroPassengersInfo, delay string, travelTime string, estimatedArrivalTime time.Time, affiliateTicket bool, ) *SroTicket`

NewSroTicket instantiates a new SroTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroTicketWithDefaults

`func NewSroTicketWithDefaults() *SroTicket`

NewSroTicketWithDefaults instantiates a new SroTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSroTicketId

`func (o *SroTicket) GetSroTicketId() int64`

GetSroTicketId returns the SroTicketId field if non-nil, zero value otherwise.

### GetSroTicketIdOk

`func (o *SroTicket) GetSroTicketIdOk() (*int64, bool)`

GetSroTicketIdOk returns a tuple with the SroTicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSroTicketId

`func (o *SroTicket) SetSroTicketId(v int64)`

SetSroTicketId sets SroTicketId field to given value.


### GetPrice

`func (o *SroTicket) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SroTicket) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SroTicket) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetUnpaid

`func (o *SroTicket) GetUnpaid() float32`

GetUnpaid returns the Unpaid field if non-nil, zero value otherwise.

### GetUnpaidOk

`func (o *SroTicket) GetUnpaidOk() (*float32, bool)`

GetUnpaidOk returns a tuple with the Unpaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaid

`func (o *SroTicket) SetUnpaid(v float32)`

SetUnpaid sets Unpaid field to given value.


### GetCurrency

`func (o *SroTicket) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SroTicket) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SroTicket) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetState

`func (o *SroTicket) GetState() TicketState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SroTicket) GetStateOk() (*TicketState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SroTicket) SetState(v TicketState)`

SetState sets State field to given value.


### GetSeatClassKey

`func (o *SroTicket) GetSeatClassKey() string`

GetSeatClassKey returns the SeatClassKey field if non-nil, zero value otherwise.

### GetSeatClassKeyOk

`func (o *SroTicket) GetSeatClassKeyOk() (*string, bool)`

GetSeatClassKeyOk returns a tuple with the SeatClassKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClassKey

`func (o *SroTicket) SetSeatClassKey(v string)`

SetSeatClassKey sets SeatClassKey field to given value.


### GetConditions

`func (o *SroTicket) GetConditions() SroConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SroTicket) GetConditionsOk() (*SroConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SroTicket) SetConditions(v SroConditions)`

SetConditions sets Conditions field to given value.


### GetCustomerActions

`func (o *SroTicket) GetCustomerActions() CustomerActions`

GetCustomerActions returns the CustomerActions field if non-nil, zero value otherwise.

### GetCustomerActionsOk

`func (o *SroTicket) GetCustomerActionsOk() (*CustomerActions, bool)`

GetCustomerActionsOk returns a tuple with the CustomerActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerActions

`func (o *SroTicket) SetCustomerActions(v CustomerActions)`

SetCustomerActions sets CustomerActions field to given value.


### GetRouteSections

`func (o *SroTicket) GetRouteSections() []TicketSection`

GetRouteSections returns the RouteSections field if non-nil, zero value otherwise.

### GetRouteSectionsOk

`func (o *SroTicket) GetRouteSectionsOk() (*[]TicketSection, bool)`

GetRouteSectionsOk returns a tuple with the RouteSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteSections

`func (o *SroTicket) SetRouteSections(v []TicketSection)`

SetRouteSections sets RouteSections field to given value.


### GetPaymentId

`func (o *SroTicket) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *SroTicket) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *SroTicket) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.


### GetBills

`func (o *SroTicket) GetBills() []TicketBill`

GetBills returns the Bills field if non-nil, zero value otherwise.

### GetBillsOk

`func (o *SroTicket) GetBillsOk() (*[]TicketBill, bool)`

GetBillsOk returns a tuple with the Bills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBills

`func (o *SroTicket) SetBills(v []TicketBill)`

SetBills sets Bills field to given value.


### GetPassengersInfo

`func (o *SroTicket) GetPassengersInfo() SroPassengersInfo`

GetPassengersInfo returns the PassengersInfo field if non-nil, zero value otherwise.

### GetPassengersInfoOk

`func (o *SroTicket) GetPassengersInfoOk() (*SroPassengersInfo, bool)`

GetPassengersInfoOk returns a tuple with the PassengersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengersInfo

`func (o *SroTicket) SetPassengersInfo(v SroPassengersInfo)`

SetPassengersInfo sets PassengersInfo field to given value.


### GetDelay

`func (o *SroTicket) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *SroTicket) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *SroTicket) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetTravelTime

`func (o *SroTicket) GetTravelTime() string`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *SroTicket) GetTravelTimeOk() (*string, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *SroTicket) SetTravelTime(v string)`

SetTravelTime sets TravelTime field to given value.


### GetEstimatedArrivalTime

`func (o *SroTicket) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *SroTicket) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *SroTicket) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.


### GetAffiliateTicket

`func (o *SroTicket) GetAffiliateTicket() bool`

GetAffiliateTicket returns the AffiliateTicket field if non-nil, zero value otherwise.

### GetAffiliateTicketOk

`func (o *SroTicket) GetAffiliateTicketOk() (*bool, bool)`

GetAffiliateTicketOk returns a tuple with the AffiliateTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateTicket

`func (o *SroTicket) SetAffiliateTicket(v bool)`

SetAffiliateTicket sets AffiliateTicket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


