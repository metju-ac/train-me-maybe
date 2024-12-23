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
	"bytes"
	"fmt"
)

// checks if the SroPriceConditionsDescriptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SroPriceConditionsDescriptions{}

// SroPriceConditionsDescriptions struct for SroPriceConditionsDescriptions
type SroPriceConditionsDescriptions struct {
	CoolingOff []string `json:"coolingOff"`
	DepartureCancel []string `json:"departureCancel"`
	Expiration []string `json:"expiration"`
}

type _SroPriceConditionsDescriptions SroPriceConditionsDescriptions

// NewSroPriceConditionsDescriptions instantiates a new SroPriceConditionsDescriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSroPriceConditionsDescriptions(coolingOff []string, departureCancel []string, expiration []string) *SroPriceConditionsDescriptions {
	this := SroPriceConditionsDescriptions{}
	this.CoolingOff = coolingOff
	this.DepartureCancel = departureCancel
	this.Expiration = expiration
	return &this
}

// NewSroPriceConditionsDescriptionsWithDefaults instantiates a new SroPriceConditionsDescriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSroPriceConditionsDescriptionsWithDefaults() *SroPriceConditionsDescriptions {
	this := SroPriceConditionsDescriptions{}
	return &this
}

// GetCoolingOff returns the CoolingOff field value
func (o *SroPriceConditionsDescriptions) GetCoolingOff() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CoolingOff
}

// GetCoolingOffOk returns a tuple with the CoolingOff field value
// and a boolean to check if the value has been set.
func (o *SroPriceConditionsDescriptions) GetCoolingOffOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoolingOff, true
}

// SetCoolingOff sets field value
func (o *SroPriceConditionsDescriptions) SetCoolingOff(v []string) {
	o.CoolingOff = v
}

// GetDepartureCancel returns the DepartureCancel field value
func (o *SroPriceConditionsDescriptions) GetDepartureCancel() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DepartureCancel
}

// GetDepartureCancelOk returns a tuple with the DepartureCancel field value
// and a boolean to check if the value has been set.
func (o *SroPriceConditionsDescriptions) GetDepartureCancelOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepartureCancel, true
}

// SetDepartureCancel sets field value
func (o *SroPriceConditionsDescriptions) SetDepartureCancel(v []string) {
	o.DepartureCancel = v
}

// GetExpiration returns the Expiration field value
func (o *SroPriceConditionsDescriptions) GetExpiration() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *SroPriceConditionsDescriptions) GetExpirationOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration, true
}

// SetExpiration sets field value
func (o *SroPriceConditionsDescriptions) SetExpiration(v []string) {
	o.Expiration = v
}

func (o SroPriceConditionsDescriptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SroPriceConditionsDescriptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coolingOff"] = o.CoolingOff
	toSerialize["departureCancel"] = o.DepartureCancel
	toSerialize["expiration"] = o.Expiration
	return toSerialize, nil
}

func (o *SroPriceConditionsDescriptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"coolingOff",
		"departureCancel",
		"expiration",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSroPriceConditionsDescriptions := _SroPriceConditionsDescriptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSroPriceConditionsDescriptions)

	if err != nil {
		return err
	}

	*o = SroPriceConditionsDescriptions(varSroPriceConditionsDescriptions)

	return err
}

type NullableSroPriceConditionsDescriptions struct {
	value *SroPriceConditionsDescriptions
	isSet bool
}

func (v NullableSroPriceConditionsDescriptions) Get() *SroPriceConditionsDescriptions {
	return v.value
}

func (v *NullableSroPriceConditionsDescriptions) Set(val *SroPriceConditionsDescriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSroPriceConditionsDescriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSroPriceConditionsDescriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSroPriceConditionsDescriptions(val *SroPriceConditionsDescriptions) *NullableSroPriceConditionsDescriptions {
	return &NullableSroPriceConditionsDescriptions{value: val, isSet: true}
}

func (v NullableSroPriceConditionsDescriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSroPriceConditionsDescriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


