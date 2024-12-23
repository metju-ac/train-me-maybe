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

// checks if the PassengersInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PassengersInfo{}

// PassengersInfo Summary of all informations needed for work with passengers.
type PassengersInfo struct {
	Passengers []Passenger `json:"passengers"`
	// Required fields for 1st passenger (returns only if tickets is editable)
	FirstPassengerData []PersonalDataType `json:"firstPassengerData"`
	// Required fields for all others passengers (returns only if tickets is editable)
	OtherPassengersData []PersonalDataType `json:"otherPassengersData"`
	// Administrative charge for user information change (in ticket currency)
	ChangeCharge float32 `json:"changeCharge"`
}

type _PassengersInfo PassengersInfo

// NewPassengersInfo instantiates a new PassengersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPassengersInfo(passengers []Passenger, firstPassengerData []PersonalDataType, otherPassengersData []PersonalDataType, changeCharge float32) *PassengersInfo {
	this := PassengersInfo{}
	this.Passengers = passengers
	this.FirstPassengerData = firstPassengerData
	this.OtherPassengersData = otherPassengersData
	this.ChangeCharge = changeCharge
	return &this
}

// NewPassengersInfoWithDefaults instantiates a new PassengersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPassengersInfoWithDefaults() *PassengersInfo {
	this := PassengersInfo{}
	return &this
}

// GetPassengers returns the Passengers field value
func (o *PassengersInfo) GetPassengers() []Passenger {
	if o == nil {
		var ret []Passenger
		return ret
	}

	return o.Passengers
}

// GetPassengersOk returns a tuple with the Passengers field value
// and a boolean to check if the value has been set.
func (o *PassengersInfo) GetPassengersOk() ([]Passenger, bool) {
	if o == nil {
		return nil, false
	}
	return o.Passengers, true
}

// SetPassengers sets field value
func (o *PassengersInfo) SetPassengers(v []Passenger) {
	o.Passengers = v
}

// GetFirstPassengerData returns the FirstPassengerData field value
func (o *PassengersInfo) GetFirstPassengerData() []PersonalDataType {
	if o == nil {
		var ret []PersonalDataType
		return ret
	}

	return o.FirstPassengerData
}

// GetFirstPassengerDataOk returns a tuple with the FirstPassengerData field value
// and a boolean to check if the value has been set.
func (o *PassengersInfo) GetFirstPassengerDataOk() ([]PersonalDataType, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstPassengerData, true
}

// SetFirstPassengerData sets field value
func (o *PassengersInfo) SetFirstPassengerData(v []PersonalDataType) {
	o.FirstPassengerData = v
}

// GetOtherPassengersData returns the OtherPassengersData field value
func (o *PassengersInfo) GetOtherPassengersData() []PersonalDataType {
	if o == nil {
		var ret []PersonalDataType
		return ret
	}

	return o.OtherPassengersData
}

// GetOtherPassengersDataOk returns a tuple with the OtherPassengersData field value
// and a boolean to check if the value has been set.
func (o *PassengersInfo) GetOtherPassengersDataOk() ([]PersonalDataType, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherPassengersData, true
}

// SetOtherPassengersData sets field value
func (o *PassengersInfo) SetOtherPassengersData(v []PersonalDataType) {
	o.OtherPassengersData = v
}

// GetChangeCharge returns the ChangeCharge field value
func (o *PassengersInfo) GetChangeCharge() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ChangeCharge
}

// GetChangeChargeOk returns a tuple with the ChangeCharge field value
// and a boolean to check if the value has been set.
func (o *PassengersInfo) GetChangeChargeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeCharge, true
}

// SetChangeCharge sets field value
func (o *PassengersInfo) SetChangeCharge(v float32) {
	o.ChangeCharge = v
}

func (o PassengersInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PassengersInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["passengers"] = o.Passengers
	toSerialize["firstPassengerData"] = o.FirstPassengerData
	toSerialize["otherPassengersData"] = o.OtherPassengersData
	toSerialize["changeCharge"] = o.ChangeCharge
	return toSerialize, nil
}

func (o *PassengersInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"passengers",
		"firstPassengerData",
		"otherPassengersData",
		"changeCharge",
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

	varPassengersInfo := _PassengersInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPassengersInfo)

	if err != nil {
		return err
	}

	*o = PassengersInfo(varPassengersInfo)

	return err
}

type NullablePassengersInfo struct {
	value *PassengersInfo
	isSet bool
}

func (v NullablePassengersInfo) Get() *PassengersInfo {
	return v.value
}

func (v *NullablePassengersInfo) Set(val *PassengersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePassengersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePassengersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePassengersInfo(val *PassengersInfo) *NullablePassengersInfo {
	return &NullablePassengersInfo{value: val, isSet: true}
}

func (v NullablePassengersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePassengersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
