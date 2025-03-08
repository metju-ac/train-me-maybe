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

// checks if the SroTicketBookingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SroTicketBookingResponse{}

// SroTicketBookingResponse struct for SroTicketBookingResponse
type SroTicketBookingResponse struct {
	// Ticket number that customer need during checkin on board.
	AccountCode string `json:"accountCode"`
	// Information about the newly created ticket.
	SroTickets []SroTicket `json:"sroTickets"`
}

type _SroTicketBookingResponse SroTicketBookingResponse

// NewSroTicketBookingResponse instantiates a new SroTicketBookingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSroTicketBookingResponse(accountCode string, sroTickets []SroTicket) *SroTicketBookingResponse {
	this := SroTicketBookingResponse{}
	this.AccountCode = accountCode
	this.SroTickets = sroTickets
	return &this
}

// NewSroTicketBookingResponseWithDefaults instantiates a new SroTicketBookingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSroTicketBookingResponseWithDefaults() *SroTicketBookingResponse {
	this := SroTicketBookingResponse{}
	return &this
}

// GetAccountCode returns the AccountCode field value
func (o *SroTicketBookingResponse) GetAccountCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value
// and a boolean to check if the value has been set.
func (o *SroTicketBookingResponse) GetAccountCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountCode, true
}

// SetAccountCode sets field value
func (o *SroTicketBookingResponse) SetAccountCode(v string) {
	o.AccountCode = v
}

// GetSroTickets returns the SroTickets field value
func (o *SroTicketBookingResponse) GetSroTickets() []SroTicket {
	if o == nil {
		var ret []SroTicket
		return ret
	}

	return o.SroTickets
}

// GetSroTicketsOk returns a tuple with the SroTickets field value
// and a boolean to check if the value has been set.
func (o *SroTicketBookingResponse) GetSroTicketsOk() ([]SroTicket, bool) {
	if o == nil {
		return nil, false
	}
	return o.SroTickets, true
}

// SetSroTickets sets field value
func (o *SroTicketBookingResponse) SetSroTickets(v []SroTicket) {
	o.SroTickets = v
}

func (o SroTicketBookingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SroTicketBookingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountCode"] = o.AccountCode
	toSerialize["sroTickets"] = o.SroTickets
	return toSerialize, nil
}

func (o *SroTicketBookingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountCode",
		"sroTickets",
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

	varSroTicketBookingResponse := _SroTicketBookingResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSroTicketBookingResponse)

	if err != nil {
		return err
	}

	*o = SroTicketBookingResponse(varSroTicketBookingResponse)

	return err
}

type NullableSroTicketBookingResponse struct {
	value *SroTicketBookingResponse
	isSet bool
}

func (v NullableSroTicketBookingResponse) Get() *SroTicketBookingResponse {
	return v.value
}

func (v *NullableSroTicketBookingResponse) Set(val *SroTicketBookingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSroTicketBookingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSroTicketBookingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSroTicketBookingResponse(val *SroTicketBookingResponse) *NullableSroTicketBookingResponse {
	return &NullableSroTicketBookingResponse{value: val, isSet: true}
}

func (v NullableSroTicketBookingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSroTicketBookingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
