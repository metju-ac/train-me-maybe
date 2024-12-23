# PriceClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeatClassKey** | **string** | Seat class key. | 
**Conditions** | [**PriceConditions**](PriceConditions.md) |  | 
**Services** | **[]string** | Service icons for wifi, steward, etc. | 
**FreeSeatsCount** | **int32** |  | 
**Price** | **float32** |  | 
**CreditPrice** | **float32** |  | 
**PriceSource** | **string** | Pricing ID - used for price, services or terms confirmation after route search. | 
**CustomerNotifications** | **[]string** | Customer notifications | 
**ActionPrice** | Pointer to [**ActionPrice**](ActionPrice.md) |  | [optional] 
**Tariffs** | **[]string** |  | 
**TariffNotifications** | Pointer to [**TariffNotifications**](TariffNotifications.md) |  | [optional] 
**Bookable** | **bool** | &#x60;true&#x60; if there are free seats in class available for reservation, otherwise &#x60;false&#x60;. | 

## Methods

### NewPriceClass

`func NewPriceClass(seatClassKey string, conditions PriceConditions, services []string, freeSeatsCount int32, price float32, creditPrice float32, priceSource string, customerNotifications []string, tariffs []string, bookable bool, ) *PriceClass`

NewPriceClass instantiates a new PriceClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceClassWithDefaults

`func NewPriceClassWithDefaults() *PriceClass`

NewPriceClassWithDefaults instantiates a new PriceClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeatClassKey

`func (o *PriceClass) GetSeatClassKey() string`

GetSeatClassKey returns the SeatClassKey field if non-nil, zero value otherwise.

### GetSeatClassKeyOk

`func (o *PriceClass) GetSeatClassKeyOk() (*string, bool)`

GetSeatClassKeyOk returns a tuple with the SeatClassKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClassKey

`func (o *PriceClass) SetSeatClassKey(v string)`

SetSeatClassKey sets SeatClassKey field to given value.


### GetConditions

`func (o *PriceClass) GetConditions() PriceConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PriceClass) GetConditionsOk() (*PriceConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PriceClass) SetConditions(v PriceConditions)`

SetConditions sets Conditions field to given value.


### GetServices

`func (o *PriceClass) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *PriceClass) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *PriceClass) SetServices(v []string)`

SetServices sets Services field to given value.


### GetFreeSeatsCount

`func (o *PriceClass) GetFreeSeatsCount() int32`

GetFreeSeatsCount returns the FreeSeatsCount field if non-nil, zero value otherwise.

### GetFreeSeatsCountOk

`func (o *PriceClass) GetFreeSeatsCountOk() (*int32, bool)`

GetFreeSeatsCountOk returns a tuple with the FreeSeatsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSeatsCount

`func (o *PriceClass) SetFreeSeatsCount(v int32)`

SetFreeSeatsCount sets FreeSeatsCount field to given value.


### GetPrice

`func (o *PriceClass) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceClass) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceClass) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetCreditPrice

`func (o *PriceClass) GetCreditPrice() float32`

GetCreditPrice returns the CreditPrice field if non-nil, zero value otherwise.

### GetCreditPriceOk

`func (o *PriceClass) GetCreditPriceOk() (*float32, bool)`

GetCreditPriceOk returns a tuple with the CreditPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPrice

`func (o *PriceClass) SetCreditPrice(v float32)`

SetCreditPrice sets CreditPrice field to given value.


### GetPriceSource

`func (o *PriceClass) GetPriceSource() string`

GetPriceSource returns the PriceSource field if non-nil, zero value otherwise.

### GetPriceSourceOk

`func (o *PriceClass) GetPriceSourceOk() (*string, bool)`

GetPriceSourceOk returns a tuple with the PriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSource

`func (o *PriceClass) SetPriceSource(v string)`

SetPriceSource sets PriceSource field to given value.


### GetCustomerNotifications

`func (o *PriceClass) GetCustomerNotifications() []string`

GetCustomerNotifications returns the CustomerNotifications field if non-nil, zero value otherwise.

### GetCustomerNotificationsOk

`func (o *PriceClass) GetCustomerNotificationsOk() (*[]string, bool)`

GetCustomerNotificationsOk returns a tuple with the CustomerNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNotifications

`func (o *PriceClass) SetCustomerNotifications(v []string)`

SetCustomerNotifications sets CustomerNotifications field to given value.


### GetActionPrice

`func (o *PriceClass) GetActionPrice() ActionPrice`

GetActionPrice returns the ActionPrice field if non-nil, zero value otherwise.

### GetActionPriceOk

`func (o *PriceClass) GetActionPriceOk() (*ActionPrice, bool)`

GetActionPriceOk returns a tuple with the ActionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPrice

`func (o *PriceClass) SetActionPrice(v ActionPrice)`

SetActionPrice sets ActionPrice field to given value.

### HasActionPrice

`func (o *PriceClass) HasActionPrice() bool`

HasActionPrice returns a boolean if a field has been set.

### GetTariffs

`func (o *PriceClass) GetTariffs() []string`

GetTariffs returns the Tariffs field if non-nil, zero value otherwise.

### GetTariffsOk

`func (o *PriceClass) GetTariffsOk() (*[]string, bool)`

GetTariffsOk returns a tuple with the Tariffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffs

`func (o *PriceClass) SetTariffs(v []string)`

SetTariffs sets Tariffs field to given value.


### GetTariffNotifications

`func (o *PriceClass) GetTariffNotifications() TariffNotifications`

GetTariffNotifications returns the TariffNotifications field if non-nil, zero value otherwise.

### GetTariffNotificationsOk

`func (o *PriceClass) GetTariffNotificationsOk() (*TariffNotifications, bool)`

GetTariffNotificationsOk returns a tuple with the TariffNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffNotifications

`func (o *PriceClass) SetTariffNotifications(v TariffNotifications)`

SetTariffNotifications sets TariffNotifications field to given value.

### HasTariffNotifications

`func (o *PriceClass) HasTariffNotifications() bool`

HasTariffNotifications returns a boolean if a field has been set.

### GetBookable

`func (o *PriceClass) GetBookable() bool`

GetBookable returns the Bookable field if non-nil, zero value otherwise.

### GetBookableOk

`func (o *PriceClass) GetBookableOk() (*bool, bool)`

GetBookableOk returns a tuple with the Bookable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookable

`func (o *PriceClass) SetBookable(v bool)`

SetBookable sets Bookable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


