/*
RegioJet's Affiliate API Reference

The RegioJet\\'s Affiliate API is a set of endpoints that help your application integrate with RegioJet.  The API is organized arount [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API uses standard HTTP methods, authentication, and status codes.  # Authentication Authentication to the API is performed via [HTTP Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication) for all endpoints listed in this documentation with the exception of `/users/authenticate`, which uses bearer token.  API requests without authentication will fail.  All API requests must be made over [HTTPS](https://en.wikipedia.org/wiki/HTTPS).  # Errors  RegioJet uses conventional HTTP status codes in responses to indicate the success or failure of an API request.  In general:   * `2xx` codes indicate success;   * `4xx` codes indicate an error that failed given the information provided in request.   * `5xx` codes indicate an error with RegioJet's servers.

API version: 1.1.0
Contact: developers@studentagency.cz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	Id *int64 `json:"id,omitempty"`
	// Used as account ID for login
	AccountCode *string `json:"accountCode,omitempty"`
	FirstName   *string `json:"firstName,omitempty"`
	Surname     *string `json:"surname,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Restrict phone number for work with sms.
	RestrictPhoneNumbers *bool   `json:"restrictPhoneNumbers,omitempty"`
	Email                *string `json:"email,omitempty"`
	// Account balance
	Credit *float32 `json:"credit,omitempty"`
	// Difference between registered credit account with price advantage and regular unregistered open account
	CreditPrice *bool     `json:"creditPrice,omitempty"`
	Currency    *Currency `json:"currency,omitempty"`
	// Set tariff which would be pre-filled after login
	DefaultTariffKey   *string        `json:"defaultTariffKey,omitempty"`
	Notifications      *Notifications `json:"notifications,omitempty"`
	CompanyInformation *bool          `json:"companyInformation,omitempty"`
	Company            *Company       `json:"company,omitempty"`
	// Acceptance of transport conditions and personal data protection
	ConditionsAcceptance *bool `json:"conditionsAcceptance,omitempty"`
	// NULL for normal customer, internal number for Regiojet employee
	EmployeeNumber *int64 `json:"employeeNumber,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	var restrictPhoneNumbers bool = false
	this.RestrictPhoneNumbers = &restrictPhoneNumbers
	var creditPrice bool = true
	this.CreditPrice = &creditPrice
	var currency Currency = EUR
	this.Currency = &currency
	var companyInformation bool = false
	this.CompanyInformation = &companyInformation
	var conditionsAcceptance bool = false
	this.ConditionsAcceptance = &conditionsAcceptance
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	var restrictPhoneNumbers bool = false
	this.RestrictPhoneNumbers = &restrictPhoneNumbers
	var creditPrice bool = true
	this.CreditPrice = &creditPrice
	var currency Currency = EUR
	this.Currency = &currency
	var companyInformation bool = false
	this.CompanyInformation = &companyInformation
	var conditionsAcceptance bool = false
	this.ConditionsAcceptance = &conditionsAcceptance
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *User) SetId(v int64) {
	o.Id = &v
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *User) GetAccountCode() string {
	if o == nil || IsNil(o.AccountCode) {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCode) {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCode returns a boolean if a field has been set.
func (o *User) HasAccountCode() bool {
	if o != nil && !IsNil(o.AccountCode) {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *User) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *User) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *User) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *User) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *User) SetSurname(v string) {
	o.Surname = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *User) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *User) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *User) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetRestrictPhoneNumbers returns the RestrictPhoneNumbers field value if set, zero value otherwise.
func (o *User) GetRestrictPhoneNumbers() bool {
	if o == nil || IsNil(o.RestrictPhoneNumbers) {
		var ret bool
		return ret
	}
	return *o.RestrictPhoneNumbers
}

// GetRestrictPhoneNumbersOk returns a tuple with the RestrictPhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRestrictPhoneNumbersOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictPhoneNumbers) {
		return nil, false
	}
	return o.RestrictPhoneNumbers, true
}

// HasRestrictPhoneNumbers returns a boolean if a field has been set.
func (o *User) HasRestrictPhoneNumbers() bool {
	if o != nil && !IsNil(o.RestrictPhoneNumbers) {
		return true
	}

	return false
}

// SetRestrictPhoneNumbers gets a reference to the given bool and assigns it to the RestrictPhoneNumbers field.
func (o *User) SetRestrictPhoneNumbers(v bool) {
	o.RestrictPhoneNumbers = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *User) GetCredit() float32 {
	if o == nil || IsNil(o.Credit) {
		var ret float32
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreditOk() (*float32, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *User) HasCredit() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given float32 and assigns it to the Credit field.
func (o *User) SetCredit(v float32) {
	o.Credit = &v
}

// GetCreditPrice returns the CreditPrice field value if set, zero value otherwise.
func (o *User) GetCreditPrice() bool {
	if o == nil || IsNil(o.CreditPrice) {
		var ret bool
		return ret
	}
	return *o.CreditPrice
}

// GetCreditPriceOk returns a tuple with the CreditPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreditPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.CreditPrice) {
		return nil, false
	}
	return o.CreditPrice, true
}

// HasCreditPrice returns a boolean if a field has been set.
func (o *User) HasCreditPrice() bool {
	if o != nil && !IsNil(o.CreditPrice) {
		return true
	}

	return false
}

// SetCreditPrice gets a reference to the given bool and assigns it to the CreditPrice field.
func (o *User) SetCreditPrice(v bool) {
	o.CreditPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *User) GetCurrency() Currency {
	if o == nil || IsNil(o.Currency) {
		var ret Currency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCurrencyOk() (*Currency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *User) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given Currency and assigns it to the Currency field.
func (o *User) SetCurrency(v Currency) {
	o.Currency = &v
}

// GetDefaultTariffKey returns the DefaultTariffKey field value if set, zero value otherwise.
func (o *User) GetDefaultTariffKey() string {
	if o == nil || IsNil(o.DefaultTariffKey) {
		var ret string
		return ret
	}
	return *o.DefaultTariffKey
}

// GetDefaultTariffKeyOk returns a tuple with the DefaultTariffKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDefaultTariffKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTariffKey) {
		return nil, false
	}
	return o.DefaultTariffKey, true
}

// HasDefaultTariffKey returns a boolean if a field has been set.
func (o *User) HasDefaultTariffKey() bool {
	if o != nil && !IsNil(o.DefaultTariffKey) {
		return true
	}

	return false
}

// SetDefaultTariffKey gets a reference to the given string and assigns it to the DefaultTariffKey field.
func (o *User) SetDefaultTariffKey(v string) {
	o.DefaultTariffKey = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *User) GetNotifications() Notifications {
	if o == nil || IsNil(o.Notifications) {
		var ret Notifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNotificationsOk() (*Notifications, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *User) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given Notifications and assigns it to the Notifications field.
func (o *User) SetNotifications(v Notifications) {
	o.Notifications = &v
}

// GetCompanyInformation returns the CompanyInformation field value if set, zero value otherwise.
func (o *User) GetCompanyInformation() bool {
	if o == nil || IsNil(o.CompanyInformation) {
		var ret bool
		return ret
	}
	return *o.CompanyInformation
}

// GetCompanyInformationOk returns a tuple with the CompanyInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCompanyInformationOk() (*bool, bool) {
	if o == nil || IsNil(o.CompanyInformation) {
		return nil, false
	}
	return o.CompanyInformation, true
}

// HasCompanyInformation returns a boolean if a field has been set.
func (o *User) HasCompanyInformation() bool {
	if o != nil && !IsNil(o.CompanyInformation) {
		return true
	}

	return false
}

// SetCompanyInformation gets a reference to the given bool and assigns it to the CompanyInformation field.
func (o *User) SetCompanyInformation(v bool) {
	o.CompanyInformation = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *User) GetCompany() Company {
	if o == nil || IsNil(o.Company) {
		var ret Company
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCompanyOk() (*Company, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *User) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given Company and assigns it to the Company field.
func (o *User) SetCompany(v Company) {
	o.Company = &v
}

// GetConditionsAcceptance returns the ConditionsAcceptance field value if set, zero value otherwise.
func (o *User) GetConditionsAcceptance() bool {
	if o == nil || IsNil(o.ConditionsAcceptance) {
		var ret bool
		return ret
	}
	return *o.ConditionsAcceptance
}

// GetConditionsAcceptanceOk returns a tuple with the ConditionsAcceptance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetConditionsAcceptanceOk() (*bool, bool) {
	if o == nil || IsNil(o.ConditionsAcceptance) {
		return nil, false
	}
	return o.ConditionsAcceptance, true
}

// HasConditionsAcceptance returns a boolean if a field has been set.
func (o *User) HasConditionsAcceptance() bool {
	if o != nil && !IsNil(o.ConditionsAcceptance) {
		return true
	}

	return false
}

// SetConditionsAcceptance gets a reference to the given bool and assigns it to the ConditionsAcceptance field.
func (o *User) SetConditionsAcceptance(v bool) {
	o.ConditionsAcceptance = &v
}

// GetEmployeeNumber returns the EmployeeNumber field value if set, zero value otherwise.
func (o *User) GetEmployeeNumber() int64 {
	if o == nil || IsNil(o.EmployeeNumber) {
		var ret int64
		return ret
	}
	return *o.EmployeeNumber
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmployeeNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.EmployeeNumber) {
		return nil, false
	}
	return o.EmployeeNumber, true
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *User) HasEmployeeNumber() bool {
	if o != nil && !IsNil(o.EmployeeNumber) {
		return true
	}

	return false
}

// SetEmployeeNumber gets a reference to the given int64 and assigns it to the EmployeeNumber field.
func (o *User) SetEmployeeNumber(v int64) {
	o.EmployeeNumber = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccountCode) {
		toSerialize["accountCode"] = o.AccountCode
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.RestrictPhoneNumbers) {
		toSerialize["restrictPhoneNumbers"] = o.RestrictPhoneNumbers
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !IsNil(o.CreditPrice) {
		toSerialize["creditPrice"] = o.CreditPrice
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.DefaultTariffKey) {
		toSerialize["defaultTariffKey"] = o.DefaultTariffKey
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.CompanyInformation) {
		toSerialize["companyInformation"] = o.CompanyInformation
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.ConditionsAcceptance) {
		toSerialize["conditionsAcceptance"] = o.ConditionsAcceptance
	}
	if !IsNil(o.EmployeeNumber) {
		toSerialize["employeeNumber"] = o.EmployeeNumber
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
