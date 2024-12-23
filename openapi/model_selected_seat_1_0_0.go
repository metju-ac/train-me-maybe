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

// checks if the SelectedSeat100 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectedSeat100{}

// SelectedSeat100 struct for SelectedSeat100
type SelectedSeat100 struct {
	SectionId int64 `json:"sectionId"`
	VehicleNumber int32 `json:"vehicleNumber"`
	SeatIndex int32 `json:"seatIndex"`
}

type _SelectedSeat100 SelectedSeat100

// NewSelectedSeat100 instantiates a new SelectedSeat100 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectedSeat100(sectionId int64, vehicleNumber int32, seatIndex int32) *SelectedSeat100 {
	this := SelectedSeat100{}
	this.SectionId = sectionId
	this.VehicleNumber = vehicleNumber
	this.SeatIndex = seatIndex
	return &this
}

// NewSelectedSeat100WithDefaults instantiates a new SelectedSeat100 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectedSeat100WithDefaults() *SelectedSeat100 {
	this := SelectedSeat100{}
	return &this
}

// GetSectionId returns the SectionId field value
func (o *SelectedSeat100) GetSectionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
func (o *SelectedSeat100) GetSectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectionId, true
}

// SetSectionId sets field value
func (o *SelectedSeat100) SetSectionId(v int64) {
	o.SectionId = v
}

// GetVehicleNumber returns the VehicleNumber field value
func (o *SelectedSeat100) GetVehicleNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VehicleNumber
}

// GetVehicleNumberOk returns a tuple with the VehicleNumber field value
// and a boolean to check if the value has been set.
func (o *SelectedSeat100) GetVehicleNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VehicleNumber, true
}

// SetVehicleNumber sets field value
func (o *SelectedSeat100) SetVehicleNumber(v int32) {
	o.VehicleNumber = v
}

// GetSeatIndex returns the SeatIndex field value
func (o *SelectedSeat100) GetSeatIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatIndex
}

// GetSeatIndexOk returns a tuple with the SeatIndex field value
// and a boolean to check if the value has been set.
func (o *SelectedSeat100) GetSeatIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeatIndex, true
}

// SetSeatIndex sets field value
func (o *SelectedSeat100) SetSeatIndex(v int32) {
	o.SeatIndex = v
}

func (o SelectedSeat100) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectedSeat100) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sectionId"] = o.SectionId
	toSerialize["vehicleNumber"] = o.VehicleNumber
	toSerialize["seatIndex"] = o.SeatIndex
	return toSerialize, nil
}

func (o *SelectedSeat100) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sectionId",
		"vehicleNumber",
		"seatIndex",
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

	varSelectedSeat100 := _SelectedSeat100{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSelectedSeat100)

	if err != nil {
		return err
	}

	*o = SelectedSeat100(varSelectedSeat100)

	return err
}

type NullableSelectedSeat100 struct {
	value *SelectedSeat100
	isSet bool
}

func (v NullableSelectedSeat100) Get() *SelectedSeat100 {
	return v.value
}

func (v *NullableSelectedSeat100) Set(val *SelectedSeat100) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedSeat100) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedSeat100) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedSeat100(val *SelectedSeat100) *NullableSelectedSeat100 {
	return &NullableSelectedSeat100{value: val, isSet: true}
}

func (v NullableSelectedSeat100) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedSeat100) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


