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

// checks if the CreateTicketSectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTicketSectionRequest{}

// CreateTicketSectionRequest struct for CreateTicketSectionRequest
type CreateTicketSectionRequest struct {
	Section       SimpleSection     `json:"section"`
	SelectedSeats []SelectedSeat100 `json:"selectedSeats,omitempty"`
}

type _CreateTicketSectionRequest CreateTicketSectionRequest

// NewCreateTicketSectionRequest instantiates a new CreateTicketSectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTicketSectionRequest(section SimpleSection) *CreateTicketSectionRequest {
	this := CreateTicketSectionRequest{}
	this.Section = section
	return &this
}

// NewCreateTicketSectionRequestWithDefaults instantiates a new CreateTicketSectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTicketSectionRequestWithDefaults() *CreateTicketSectionRequest {
	this := CreateTicketSectionRequest{}
	return &this
}

// GetSection returns the Section field value
func (o *CreateTicketSectionRequest) GetSection() SimpleSection {
	if o == nil {
		var ret SimpleSection
		return ret
	}

	return o.Section
}

// GetSectionOk returns a tuple with the Section field value
// and a boolean to check if the value has been set.
func (o *CreateTicketSectionRequest) GetSectionOk() (*SimpleSection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Section, true
}

// SetSection sets field value
func (o *CreateTicketSectionRequest) SetSection(v SimpleSection) {
	o.Section = v
}

// GetSelectedSeats returns the SelectedSeats field value if set, zero value otherwise.
func (o *CreateTicketSectionRequest) GetSelectedSeats() []SelectedSeat100 {
	if o == nil || IsNil(o.SelectedSeats) {
		var ret []SelectedSeat100
		return ret
	}
	return o.SelectedSeats
}

// GetSelectedSeatsOk returns a tuple with the SelectedSeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTicketSectionRequest) GetSelectedSeatsOk() ([]SelectedSeat100, bool) {
	if o == nil || IsNil(o.SelectedSeats) {
		return nil, false
	}
	return o.SelectedSeats, true
}

// HasSelectedSeats returns a boolean if a field has been set.
func (o *CreateTicketSectionRequest) HasSelectedSeats() bool {
	if o != nil && !IsNil(o.SelectedSeats) {
		return true
	}

	return false
}

// SetSelectedSeats gets a reference to the given []SelectedSeat100 and assigns it to the SelectedSeats field.
func (o *CreateTicketSectionRequest) SetSelectedSeats(v []SelectedSeat100) {
	o.SelectedSeats = v
}

func (o CreateTicketSectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTicketSectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["section"] = o.Section
	if !IsNil(o.SelectedSeats) {
		toSerialize["selectedSeats"] = o.SelectedSeats
	}
	return toSerialize, nil
}

func (o *CreateTicketSectionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"section",
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

	varCreateTicketSectionRequest := _CreateTicketSectionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTicketSectionRequest)

	if err != nil {
		return err
	}

	*o = CreateTicketSectionRequest(varCreateTicketSectionRequest)

	return err
}

type NullableCreateTicketSectionRequest struct {
	value *CreateTicketSectionRequest
	isSet bool
}

func (v NullableCreateTicketSectionRequest) Get() *CreateTicketSectionRequest {
	return v.value
}

func (v *NullableCreateTicketSectionRequest) Set(val *CreateTicketSectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTicketSectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTicketSectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTicketSectionRequest(val *CreateTicketSectionRequest) *NullableCreateTicketSectionRequest {
	return &NullableCreateTicketSectionRequest{value: val, isSet: true}
}

func (v NullableCreateTicketSectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTicketSectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
