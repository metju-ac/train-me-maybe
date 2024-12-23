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

// checks if the CreateTicketResponseUnregistered type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTicketResponseUnregistered{}

// CreateTicketResponseUnregistered struct for CreateTicketResponseUnregistered
type CreateTicketResponseUnregistered struct {
	// Newly created user account token which is linked to ticket(s)
	Token string `json:"token"`
	Tickets []Ticket `json:"tickets"`
}

type _CreateTicketResponseUnregistered CreateTicketResponseUnregistered

// NewCreateTicketResponseUnregistered instantiates a new CreateTicketResponseUnregistered object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTicketResponseUnregistered(token string, tickets []Ticket) *CreateTicketResponseUnregistered {
	this := CreateTicketResponseUnregistered{}
	this.Token = token
	this.Tickets = tickets
	return &this
}

// NewCreateTicketResponseUnregisteredWithDefaults instantiates a new CreateTicketResponseUnregistered object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTicketResponseUnregisteredWithDefaults() *CreateTicketResponseUnregistered {
	this := CreateTicketResponseUnregistered{}
	return &this
}

// GetToken returns the Token field value
func (o *CreateTicketResponseUnregistered) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CreateTicketResponseUnregistered) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CreateTicketResponseUnregistered) SetToken(v string) {
	o.Token = v
}

// GetTickets returns the Tickets field value
func (o *CreateTicketResponseUnregistered) GetTickets() []Ticket {
	if o == nil {
		var ret []Ticket
		return ret
	}

	return o.Tickets
}

// GetTicketsOk returns a tuple with the Tickets field value
// and a boolean to check if the value has been set.
func (o *CreateTicketResponseUnregistered) GetTicketsOk() ([]Ticket, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tickets, true
}

// SetTickets sets field value
func (o *CreateTicketResponseUnregistered) SetTickets(v []Ticket) {
	o.Tickets = v
}

func (o CreateTicketResponseUnregistered) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTicketResponseUnregistered) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["tickets"] = o.Tickets
	return toSerialize, nil
}

func (o *CreateTicketResponseUnregistered) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
		"tickets",
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

	varCreateTicketResponseUnregistered := _CreateTicketResponseUnregistered{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTicketResponseUnregistered)

	if err != nil {
		return err
	}

	*o = CreateTicketResponseUnregistered(varCreateTicketResponseUnregistered)

	return err
}

type NullableCreateTicketResponseUnregistered struct {
	value *CreateTicketResponseUnregistered
	isSet bool
}

func (v NullableCreateTicketResponseUnregistered) Get() *CreateTicketResponseUnregistered {
	return v.value
}

func (v *NullableCreateTicketResponseUnregistered) Set(val *CreateTicketResponseUnregistered) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTicketResponseUnregistered) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTicketResponseUnregistered) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTicketResponseUnregistered(val *CreateTicketResponseUnregistered) *NullableCreateTicketResponseUnregistered {
	return &NullableCreateTicketResponseUnregistered{value: val, isSet: true}
}

func (v NullableCreateTicketResponseUnregistered) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTicketResponseUnregistered) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


