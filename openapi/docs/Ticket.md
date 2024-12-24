# Ticket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the ticket | 
**RouteId** | **string** | route id (section0.id,section1.id, ... sectionx.id) | 
**Price** | **float32** | Final price for ticket, addons etc. | 
**Unpaid** | **float32** | Final price to be paid | 
**Currency** | [**Currency**](Currency.md) |  | [default to EUR]
**State** | [**TicketState**](TicketState.md) |  | 
**SeatClassKey** | **string** |  | 
**Conditions** | [**PriceConditions**](PriceConditions.md) |  | 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**ExpirateAt** | Pointer to [**TimePeriod**](TimePeriod.md) |  | [optional] 
**CustomerNotifications** | Pointer to **[]string** |  | [optional] 
**CustomerActions** | [**CustomerActions**](CustomerActions.md) |  | 
**RouteSections** | [**[]TicketSection**](TicketSection.md) |  | 
**Addons** | Pointer to [**[]OrderedAddon**](OrderedAddon.md) |  | [optional] 
**PaymentId** | Pointer to **int64** | Payment ID (groupTransactionID). Available only with paid ticket. Necessary for further action with ticket (for example: print invoice) | [optional] 
**Bills** | Pointer to [**[]TicketBill**](TicketBill.md) |  | [optional] 
**UsedCodeDiscount** | Pointer to [**CodeDiscount**](CodeDiscount.md) |  | [optional] 
**UsedPercentualDiscounts** | Pointer to [**[]PercentualDiscount**](PercentualDiscount.md) | Applied procentual discounts | [optional] 
**TransfersInfo** | Pointer to [**TransfersInfo**](TransfersInfo.md) |  | [optional] 
**Surcharge** | Pointer to **float32** | Total count of all irreversible surcharges in current currency | [optional] 
**CancelChargeSum** | Pointer to **float32** | Total count of all charges and surcharge | [optional] 
**CancelMoneyBackSum** | Pointer to **float32** | Total count of reversible amounts | [optional] 
**PassengersInfo** | [**PassengersInfo**](PassengersInfo.md) |  | 
**Delay** | Pointer to **string** | Textual information about the first delay on the route | [optional] 
**TravelTime** | Pointer to **string** | Textual information about the travel time on a given section | [optional] 
**EstimatedArrivalTime** | Pointer to **time.Time** | Estimated arrival date time (arrival + delay) | [optional] 

## Methods

### NewTicket

`func NewTicket(id int64, routeId string, price float32, unpaid float32, currency Currency, state TicketState, seatClassKey string, conditions PriceConditions, customerActions CustomerActions, routeSections []TicketSection, passengersInfo PassengersInfo, ) *Ticket`

NewTicket instantiates a new Ticket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketWithDefaults

`func NewTicketWithDefaults() *Ticket`

NewTicketWithDefaults instantiates a new Ticket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ticket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ticket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ticket) SetId(v int64)`

SetId sets Id field to given value.


### GetRouteId

`func (o *Ticket) GetRouteId() string`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *Ticket) GetRouteIdOk() (*string, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *Ticket) SetRouteId(v string)`

SetRouteId sets RouteId field to given value.


### GetPrice

`func (o *Ticket) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Ticket) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Ticket) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetUnpaid

`func (o *Ticket) GetUnpaid() float32`

GetUnpaid returns the Unpaid field if non-nil, zero value otherwise.

### GetUnpaidOk

`func (o *Ticket) GetUnpaidOk() (*float32, bool)`

GetUnpaidOk returns a tuple with the Unpaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpaid

`func (o *Ticket) SetUnpaid(v float32)`

SetUnpaid sets Unpaid field to given value.


### GetCurrency

`func (o *Ticket) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Ticket) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Ticket) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetState

`func (o *Ticket) GetState() TicketState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Ticket) GetStateOk() (*TicketState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Ticket) SetState(v TicketState)`

SetState sets State field to given value.


### GetSeatClassKey

`func (o *Ticket) GetSeatClassKey() string`

GetSeatClassKey returns the SeatClassKey field if non-nil, zero value otherwise.

### GetSeatClassKeyOk

`func (o *Ticket) GetSeatClassKeyOk() (*string, bool)`

GetSeatClassKeyOk returns a tuple with the SeatClassKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClassKey

`func (o *Ticket) SetSeatClassKey(v string)`

SetSeatClassKey sets SeatClassKey field to given value.


### GetConditions

`func (o *Ticket) GetConditions() PriceConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Ticket) GetConditionsOk() (*PriceConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Ticket) SetConditions(v PriceConditions)`

SetConditions sets Conditions field to given value.


### GetExpirationDate

`func (o *Ticket) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Ticket) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Ticket) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Ticket) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetExpirateAt

`func (o *Ticket) GetExpirateAt() TimePeriod`

GetExpirateAt returns the ExpirateAt field if non-nil, zero value otherwise.

### GetExpirateAtOk

`func (o *Ticket) GetExpirateAtOk() (*TimePeriod, bool)`

GetExpirateAtOk returns a tuple with the ExpirateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirateAt

`func (o *Ticket) SetExpirateAt(v TimePeriod)`

SetExpirateAt sets ExpirateAt field to given value.

### HasExpirateAt

`func (o *Ticket) HasExpirateAt() bool`

HasExpirateAt returns a boolean if a field has been set.

### GetCustomerNotifications

`func (o *Ticket) GetCustomerNotifications() []string`

GetCustomerNotifications returns the CustomerNotifications field if non-nil, zero value otherwise.

### GetCustomerNotificationsOk

`func (o *Ticket) GetCustomerNotificationsOk() (*[]string, bool)`

GetCustomerNotificationsOk returns a tuple with the CustomerNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotifications

`func (o *Ticket) SetCustomerNotifications(v []string)`

SetCustomerNotifications sets CustomerNotifications field to given value.

### HasCustomerNotifications

`func (o *Ticket) HasCustomerNotifications() bool`

HasCustomerNotifications returns a boolean if a field has been set.

### GetCustomerActions

`func (o *Ticket) GetCustomerActions() CustomerActions`

GetCustomerActions returns the CustomerActions field if non-nil, zero value otherwise.

### GetCustomerActionsOk

`func (o *Ticket) GetCustomerActionsOk() (*CustomerActions, bool)`

GetCustomerActionsOk returns a tuple with the CustomerActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerActions

`func (o *Ticket) SetCustomerActions(v CustomerActions)`

SetCustomerActions sets CustomerActions field to given value.


### GetRouteSections

`func (o *Ticket) GetRouteSections() []TicketSection`

GetRouteSections returns the RouteSections field if non-nil, zero value otherwise.

### GetRouteSectionsOk

`func (o *Ticket) GetRouteSectionsOk() (*[]TicketSection, bool)`

GetRouteSectionsOk returns a tuple with the RouteSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteSections

`func (o *Ticket) SetRouteSections(v []TicketSection)`

SetRouteSections sets RouteSections field to given value.


### GetAddons

`func (o *Ticket) GetAddons() []OrderedAddon`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *Ticket) GetAddonsOk() (*[]OrderedAddon, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *Ticket) SetAddons(v []OrderedAddon)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *Ticket) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetPaymentId

`func (o *Ticket) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Ticket) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Ticket) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *Ticket) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetBills

`func (o *Ticket) GetBills() []TicketBill`

GetBills returns the Bills field if non-nil, zero value otherwise.

### GetBillsOk

`func (o *Ticket) GetBillsOk() (*[]TicketBill, bool)`

GetBillsOk returns a tuple with the Bills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBills

`func (o *Ticket) SetBills(v []TicketBill)`

SetBills sets Bills field to given value.

### HasBills

`func (o *Ticket) HasBills() bool`

HasBills returns a boolean if a field has been set.

### GetUsedCodeDiscount

`func (o *Ticket) GetUsedCodeDiscount() CodeDiscount`

GetUsedCodeDiscount returns the UsedCodeDiscount field if non-nil, zero value otherwise.

### GetUsedCodeDiscountOk

`func (o *Ticket) GetUsedCodeDiscountOk() (*CodeDiscount, bool)`

GetUsedCodeDiscountOk returns a tuple with the UsedCodeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCodeDiscount

`func (o *Ticket) SetUsedCodeDiscount(v CodeDiscount)`

SetUsedCodeDiscount sets UsedCodeDiscount field to given value.

### HasUsedCodeDiscount

`func (o *Ticket) HasUsedCodeDiscount() bool`

HasUsedCodeDiscount returns a boolean if a field has been set.

### GetUsedPercentualDiscounts

`func (o *Ticket) GetUsedPercentualDiscounts() []PercentualDiscount`

GetUsedPercentualDiscounts returns the UsedPercentualDiscounts field if non-nil, zero value otherwise.

### GetUsedPercentualDiscountsOk

`func (o *Ticket) GetUsedPercentualDiscountsOk() (*[]PercentualDiscount, bool)`

GetUsedPercentualDiscountsOk returns a tuple with the UsedPercentualDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercentualDiscounts

`func (o *Ticket) SetUsedPercentualDiscounts(v []PercentualDiscount)`

SetUsedPercentualDiscounts sets UsedPercentualDiscounts field to given value.

### HasUsedPercentualDiscounts

`func (o *Ticket) HasUsedPercentualDiscounts() bool`

HasUsedPercentualDiscounts returns a boolean if a field has been set.

### GetTransfersInfo

`func (o *Ticket) GetTransfersInfo() TransfersInfo`

GetTransfersInfo returns the TransfersInfo field if non-nil, zero value otherwise.

### GetTransfersInfoOk

`func (o *Ticket) GetTransfersInfoOk() (*TransfersInfo, bool)`

GetTransfersInfoOk returns a tuple with the TransfersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfersInfo

`func (o *Ticket) SetTransfersInfo(v TransfersInfo)`

SetTransfersInfo sets TransfersInfo field to given value.

### HasTransfersInfo

`func (o *Ticket) HasTransfersInfo() bool`

HasTransfersInfo returns a boolean if a field has been set.

### GetSurcharge

`func (o *Ticket) GetSurcharge() float32`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *Ticket) GetSurchargeOk() (*float32, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *Ticket) SetSurcharge(v float32)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *Ticket) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetCancelChargeSum

`func (o *Ticket) GetCancelChargeSum() float32`

GetCancelChargeSum returns the CancelChargeSum field if non-nil, zero value otherwise.

### GetCancelChargeSumOk

`func (o *Ticket) GetCancelChargeSumOk() (*float32, bool)`

GetCancelChargeSumOk returns a tuple with the CancelChargeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelChargeSum

`func (o *Ticket) SetCancelChargeSum(v float32)`

SetCancelChargeSum sets CancelChargeSum field to given value.

### HasCancelChargeSum

`func (o *Ticket) HasCancelChargeSum() bool`

HasCancelChargeSum returns a boolean if a field has been set.

### GetCancelMoneyBackSum

`func (o *Ticket) GetCancelMoneyBackSum() float32`

GetCancelMoneyBackSum returns the CancelMoneyBackSum field if non-nil, zero value otherwise.

### GetCancelMoneyBackSumOk

`func (o *Ticket) GetCancelMoneyBackSumOk() (*float32, bool)`

GetCancelMoneyBackSumOk returns a tuple with the CancelMoneyBackSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelMoneyBackSum

`func (o *Ticket) SetCancelMoneyBackSum(v float32)`

SetCancelMoneyBackSum sets CancelMoneyBackSum field to given value.

### HasCancelMoneyBackSum

`func (o *Ticket) HasCancelMoneyBackSum() bool`

HasCancelMoneyBackSum returns a boolean if a field has been set.

### GetPassengersInfo

`func (o *Ticket) GetPassengersInfo() PassengersInfo`

GetPassengersInfo returns the PassengersInfo field if non-nil, zero value otherwise.

### GetPassengersInfoOk

`func (o *Ticket) GetPassengersInfoOk() (*PassengersInfo, bool)`

GetPassengersInfoOk returns a tuple with the PassengersInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengersInfo

`func (o *Ticket) SetPassengersInfo(v PassengersInfo)`

SetPassengersInfo sets PassengersInfo field to given value.


### GetDelay

`func (o *Ticket) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *Ticket) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *Ticket) SetDelay(v string)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *Ticket) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetTravelTime

`func (o *Ticket) GetTravelTime() string`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *Ticket) GetTravelTimeOk() (*string, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *Ticket) SetTravelTime(v string)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *Ticket) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.

### GetEstimatedArrivalTime

`func (o *Ticket) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *Ticket) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *Ticket) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *Ticket) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


