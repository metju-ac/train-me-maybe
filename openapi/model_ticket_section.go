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

// checks if the TicketSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TicketSection{}

// TicketSection struct for TicketSection
type TicketSection struct {
	Section Section `json:"section"`
	// TRUE => seat reservation; FALSE => seat choosed while boarding
	FixedSeatReservation bool              `json:"fixedSeatReservation"`
	SelectedSeats        []SelectedSeat100 `json:"selectedSeats"`
	Vehicles             []Vehicle110      `json:"vehicles,omitempty"`
}

type _TicketSection TicketSection

// NewTicketSection instantiates a new TicketSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicketSection(section Section, fixedSeatReservation bool, selectedSeats []SelectedSeat100) *TicketSection {
	this := TicketSection{}
	this.Section = section
	this.FixedSeatReservation = fixedSeatReservation
	this.SelectedSeats = selectedSeats
	return &this
}

// NewTicketSectionWithDefaults instantiates a new TicketSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicketSectionWithDefaults() *TicketSection {
	this := TicketSection{}
	return &this
}

// GetSection returns the Section field value
func (o *TicketSection) GetSection() Section {
	if o == nil {
		var ret Section
		return ret
	}

	return o.Section
}

// GetSectionOk returns a tuple with the Section field value
// and a boolean to check if the value has been set.
func (o *TicketSection) GetSectionOk() (*Section, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Section, true
}

// SetSection sets field value
func (o *TicketSection) SetSection(v Section) {
	o.Section = v
}

// GetFixedSeatReservation returns the FixedSeatReservation field value
func (o *TicketSection) GetFixedSeatReservation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FixedSeatReservation
}

// GetFixedSeatReservationOk returns a tuple with the FixedSeatReservation field value
// and a boolean to check if the value has been set.
func (o *TicketSection) GetFixedSeatReservationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FixedSeatReservation, true
}

// SetFixedSeatReservation sets field value
func (o *TicketSection) SetFixedSeatReservation(v bool) {
	o.FixedSeatReservation = v
}

// GetSelectedSeats returns the SelectedSeats field value
func (o *TicketSection) GetSelectedSeats() []SelectedSeat100 {
	if o == nil {
		var ret []SelectedSeat100
		return ret
	}

	return o.SelectedSeats
}

// GetSelectedSeatsOk returns a tuple with the SelectedSeats field value
// and a boolean to check if the value has been set.
func (o *TicketSection) GetSelectedSeatsOk() ([]SelectedSeat100, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedSeats, true
}

// SetSelectedSeats sets field value
func (o *TicketSection) SetSelectedSeats(v []SelectedSeat100) {
	o.SelectedSeats = v
}

// GetVehicles returns the Vehicles field value if set, zero value otherwise.
func (o *TicketSection) GetVehicles() []Vehicle110 {
	if o == nil || IsNil(o.Vehicles) {
		var ret []Vehicle110
		return ret
	}
	return o.Vehicles
}

// GetVehiclesOk returns a tuple with the Vehicles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TicketSection) GetVehiclesOk() ([]Vehicle110, bool) {
	if o == nil || IsNil(o.Vehicles) {
		return nil, false
	}
	return o.Vehicles, true
}

// HasVehicles returns a boolean if a field has been set.
func (o *TicketSection) HasVehicles() bool {
	if o != nil && !IsNil(o.Vehicles) {
		return true
	}

	return false
}

// SetVehicles gets a reference to the given []Vehicle110 and assigns it to the Vehicles field.
func (o *TicketSection) SetVehicles(v []Vehicle110) {
	o.Vehicles = v
}

func (o TicketSection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TicketSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["section"] = o.Section
	toSerialize["fixedSeatReservation"] = o.FixedSeatReservation
	toSerialize["selectedSeats"] = o.SelectedSeats
	if !IsNil(o.Vehicles) {
		toSerialize["vehicles"] = o.Vehicles
	}
	return toSerialize, nil
}

func (o *TicketSection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"section",
		"fixedSeatReservation",
		"selectedSeats",
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

	varTicketSection := _TicketSection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTicketSection)

	if err != nil {
		return err
	}

	*o = TicketSection(varTicketSection)

	return err
}

type NullableTicketSection struct {
	value *TicketSection
	isSet bool
}

func (v NullableTicketSection) Get() *TicketSection {
	return v.value
}

func (v *NullableTicketSection) Set(val *TicketSection) {
	v.value = val
	v.isSet = true
}

func (v NullableTicketSection) IsSet() bool {
	return v.isSet
}

func (v *NullableTicketSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicketSection(val *TicketSection) *NullableTicketSection {
	return &NullableTicketSection{value: val, isSet: true}
}

func (v NullableTicketSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicketSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
