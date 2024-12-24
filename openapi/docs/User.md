# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountCode** | Pointer to **string** | Used as account ID for login | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Surname** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**RestrictPhoneNumbers** | Pointer to **bool** | Restrict phone number for work with sms. | [optional] [default to false]
**Email** | Pointer to **string** |  | [optional] 
**Credit** | Pointer to **float32** | Account balance | [optional] 
**CreditPrice** | Pointer to **bool** | Difference between registered credit account with price advantage and regular unregistered open account | [optional] [default to true]
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] [default to EUR]
**DefaultTariffKey** | Pointer to **string** | Set tariff which would be pre-filled after login | [optional] 
**Notifications** | Pointer to [**Notifications**](Notifications.md) |  | [optional] 
**CompanyInformation** | Pointer to **bool** |  | [optional] [default to false]
**Company** | Pointer to [**Company**](Company.md) |  | [optional] 
**ConditionsAcceptance** | Pointer to **bool** | Acceptance of transport conditions and personal data protection | [optional] [default to false]
**EmployeeNumber** | Pointer to **int32** | NULL for normal customer, internal number for Regiojet employee | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountCode

`func (o *User) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *User) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *User) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.

### HasAccountCode

`func (o *User) HasAccountCode() bool`

HasAccountCode returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetSurname

`func (o *User) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *User) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *User) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *User) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *User) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *User) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *User) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *User) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRestrictPhoneNumbers

`func (o *User) GetRestrictPhoneNumbers() bool`

GetRestrictPhoneNumbers returns the RestrictPhoneNumbers field if non-nil, zero value otherwise.

### GetRestrictPhoneNumbersOk

`func (o *User) GetRestrictPhoneNumbersOk() (*bool, bool)`

GetRestrictPhoneNumbersOk returns a tuple with the RestrictPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictPhoneNumbers

`func (o *User) SetRestrictPhoneNumbers(v bool)`

SetRestrictPhoneNumbers sets RestrictPhoneNumbers field to given value.

### HasRestrictPhoneNumbers

`func (o *User) HasRestrictPhoneNumbers() bool`

HasRestrictPhoneNumbers returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCredit

`func (o *User) GetCredit() float32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *User) GetCreditOk() (*float32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *User) SetCredit(v float32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *User) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCreditPrice

`func (o *User) GetCreditPrice() bool`

GetCreditPrice returns the CreditPrice field if non-nil, zero value otherwise.

### GetCreditPriceOk

`func (o *User) GetCreditPriceOk() (*bool, bool)`

GetCreditPriceOk returns a tuple with the CreditPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPrice

`func (o *User) SetCreditPrice(v bool)`

SetCreditPrice sets CreditPrice field to given value.

### HasCreditPrice

`func (o *User) HasCreditPrice() bool`

HasCreditPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *User) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *User) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *User) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *User) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultTariffKey

`func (o *User) GetDefaultTariffKey() string`

GetDefaultTariffKey returns the DefaultTariffKey field if non-nil, zero value otherwise.

### GetDefaultTariffKeyOk

`func (o *User) GetDefaultTariffKeyOk() (*string, bool)`

GetDefaultTariffKeyOk returns a tuple with the DefaultTariffKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTariffKey

`func (o *User) SetDefaultTariffKey(v string)`

SetDefaultTariffKey sets DefaultTariffKey field to given value.

### HasDefaultTariffKey

`func (o *User) HasDefaultTariffKey() bool`

HasDefaultTariffKey returns a boolean if a field has been set.

### GetNotifications

`func (o *User) GetNotifications() Notifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *User) GetNotificationsOk() (*Notifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *User) SetNotifications(v Notifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *User) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetCompanyInformation

`func (o *User) GetCompanyInformation() bool`

GetCompanyInformation returns the CompanyInformation field if non-nil, zero value otherwise.

### GetCompanyInformationOk

`func (o *User) GetCompanyInformationOk() (*bool, bool)`

GetCompanyInformationOk returns a tuple with the CompanyInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInformation

`func (o *User) SetCompanyInformation(v bool)`

SetCompanyInformation sets CompanyInformation field to given value.

### HasCompanyInformation

`func (o *User) HasCompanyInformation() bool`

HasCompanyInformation returns a boolean if a field has been set.

### GetCompany

`func (o *User) GetCompany() Company`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *User) GetCompanyOk() (*Company, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *User) SetCompany(v Company)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *User) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetConditionsAcceptance

`func (o *User) GetConditionsAcceptance() bool`

GetConditionsAcceptance returns the ConditionsAcceptance field if non-nil, zero value otherwise.

### GetConditionsAcceptanceOk

`func (o *User) GetConditionsAcceptanceOk() (*bool, bool)`

GetConditionsAcceptanceOk returns a tuple with the ConditionsAcceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsAcceptance

`func (o *User) SetConditionsAcceptance(v bool)`

SetConditionsAcceptance sets ConditionsAcceptance field to given value.

### HasConditionsAcceptance

`func (o *User) HasConditionsAcceptance() bool`

HasConditionsAcceptance returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *User) GetEmployeeNumber() int32`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *User) GetEmployeeNumberOk() (*int32, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *User) SetEmployeeNumber(v int32)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *User) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


