# Passenger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**FirstName** | Pointer to **string** | First name | [optional] 
**Surname** | Pointer to **string** | Surname | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | Email are always required after first passenger if not pre-filled in user account (or customer isnt logged in) | [optional] 
**DateOfBirth** | Pointer to **string** |  | [optional] 
**Tariff** | **string** |  | 
**Amount** | Pointer to **float32** | Basic price (ticket purchase) for tariff set on ticket with ticket currency (doesnt count with discounts) | [optional] 
**MoneyBack** | Pointer to **float32** | Final amount for passenger which will be refunded if canceled. | [optional] 

## Methods

### NewPassenger

`func NewPassenger(tariff string, ) *Passenger`

NewPassenger instantiates a new Passenger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassengerWithDefaults

`func NewPassengerWithDefaults() *Passenger`

NewPassengerWithDefaults instantiates a new Passenger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Passenger) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Passenger) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Passenger) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Passenger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *Passenger) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Passenger) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Passenger) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Passenger) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetSurname

`func (o *Passenger) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *Passenger) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *Passenger) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *Passenger) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetPhone

`func (o *Passenger) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Passenger) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Passenger) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Passenger) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *Passenger) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Passenger) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Passenger) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Passenger) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *Passenger) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Passenger) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Passenger) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Passenger) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetTariff

`func (o *Passenger) GetTariff() string`

GetTariff returns the Tariff field if non-nil, zero value otherwise.

### GetTariffOk

`func (o *Passenger) GetTariffOk() (*string, bool)`

GetTariffOk returns a tuple with the Tariff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariff

`func (o *Passenger) SetTariff(v string)`

SetTariff sets Tariff field to given value.


### GetAmount

`func (o *Passenger) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Passenger) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Passenger) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Passenger) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMoneyBack

`func (o *Passenger) GetMoneyBack() float32`

GetMoneyBack returns the MoneyBack field if non-nil, zero value otherwise.

### GetMoneyBackOk

`func (o *Passenger) GetMoneyBackOk() (*float32, bool)`

GetMoneyBackOk returns a tuple with the MoneyBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyBack

`func (o *Passenger) SetMoneyBack(v float32)`

SetMoneyBack sets MoneyBack field to given value.

### HasMoneyBack

`func (o *Passenger) HasMoneyBack() bool`

HasMoneyBack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


