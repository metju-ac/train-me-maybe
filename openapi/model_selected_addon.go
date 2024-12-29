/*
RegioJet's Affiliate API Reference

The RegioJet\\'s Affiliate API is a set of endpoints that help your application integrate with RegioJet.  The API is organized arount [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). Our API uses standard HTTP methods, authentication, and status codes.  # Authentication Authentication to the API is performed via [HTTP Basic Auth](https://en.wikipedia.org/wiki/Basic_access_authentication) for all endpoints listed in this documentation with the exception of `/users/authenticate`, which uses bearer token.  API requests without authentication will fail.  All API requests must be made over [HTTPS](https://en.wikipedia.org/wiki/HTTPS).  # Errors  RegioJet uses conventional HTTP status codes in responses to indicate the success or failure of an API request.  In general:   * `2xx` codes indicate success;   * `4xx` codes indicate an error that failed given the information provided in request.   * `5xx` codes indicate an error with RegioJet's servers.

API version: 1.1.0
Contact: developers@studentagency.cz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SelectedAddon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectedAddon{}

// SelectedAddon struct for SelectedAddon
type SelectedAddon struct {
	AddonId  int64    `json:"addonId"`
	Count    int64    `json:"count"`
	Price    float32  `json:"price"`
	Currency Currency `json:"currency"`
	// Hash (MD5) of cancellation terms, agreed on by the customer
	ConditionsHash string `json:"conditionsHash"`
}

type _SelectedAddon SelectedAddon

// NewSelectedAddon instantiates a new SelectedAddon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectedAddon(addonId int64, count int64, price float32, currency Currency, conditionsHash string) *SelectedAddon {
	this := SelectedAddon{}
	this.AddonId = addonId
	this.Count = count
	this.Price = price
	this.Currency = currency
	this.ConditionsHash = conditionsHash
	return &this
}

// NewSelectedAddonWithDefaults instantiates a new SelectedAddon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectedAddonWithDefaults() *SelectedAddon {
	this := SelectedAddon{}
	var currency Currency = EUR
	this.Currency = currency
	return &this
}

// GetAddonId returns the AddonId field value
func (o *SelectedAddon) GetAddonId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AddonId
}

// GetAddonIdOk returns a tuple with the AddonId field value
// and a boolean to check if the value has been set.
func (o *SelectedAddon) GetAddonIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddonId, true
}

// SetAddonId sets field value
func (o *SelectedAddon) SetAddonId(v int64) {
	o.AddonId = v
}

// GetCount returns the Count field value
func (o *SelectedAddon) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SelectedAddon) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SelectedAddon) SetCount(v int64) {
	o.Count = v
}

// GetPrice returns the Price field value
func (o *SelectedAddon) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *SelectedAddon) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *SelectedAddon) SetPrice(v float32) {
	o.Price = v
}

// GetCurrency returns the Currency field value
func (o *SelectedAddon) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SelectedAddon) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SelectedAddon) SetCurrency(v Currency) {
	o.Currency = v
}

// GetConditionsHash returns the ConditionsHash field value
func (o *SelectedAddon) GetConditionsHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionsHash
}

// GetConditionsHashOk returns a tuple with the ConditionsHash field value
// and a boolean to check if the value has been set.
func (o *SelectedAddon) GetConditionsHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionsHash, true
}

// SetConditionsHash sets field value
func (o *SelectedAddon) SetConditionsHash(v string) {
	o.ConditionsHash = v
}

func (o SelectedAddon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectedAddon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addonId"] = o.AddonId
	toSerialize["count"] = o.Count
	toSerialize["price"] = o.Price
	toSerialize["currency"] = o.Currency
	toSerialize["conditionsHash"] = o.ConditionsHash
	return toSerialize, nil
}

func (o *SelectedAddon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addonId",
		"count",
		"price",
		"currency",
		"conditionsHash",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSelectedAddon := _SelectedAddon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSelectedAddon)

	if err != nil {
		return err
	}

	*o = SelectedAddon(varSelectedAddon)

	return err
}

type NullableSelectedAddon struct {
	value *SelectedAddon
	isSet bool
}

func (v NullableSelectedAddon) Get() *SelectedAddon {
	return v.value
}

func (v *NullableSelectedAddon) Set(val *SelectedAddon) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedAddon) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedAddon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedAddon(val *SelectedAddon) *NullableSelectedAddon {
	return &NullableSelectedAddon{value: val, isSet: true}
}

func (v NullableSelectedAddon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedAddon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
