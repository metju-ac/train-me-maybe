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

// checks if the CreateUnregisteredTicketRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUnregisteredTicketRequest{}

// CreateUnregisteredTicketRequest struct for CreateUnregisteredTicketRequest
type CreateUnregisteredTicketRequest struct {
	// Agree with terms
	AgreeWithTerms bool                  `json:"agreeWithTerms"`
	TicketRequests []CreateTicketRequest `json:"ticketRequests"`
	// Optional code of an affiliate partner.
	AffiliateCode *string `json:"affiliateCode,omitempty"`
}

type _CreateUnregisteredTicketRequest CreateUnregisteredTicketRequest

// NewCreateUnregisteredTicketRequest instantiates a new CreateUnregisteredTicketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUnregisteredTicketRequest(agreeWithTerms bool, ticketRequests []CreateTicketRequest) *CreateUnregisteredTicketRequest {
	this := CreateUnregisteredTicketRequest{}
	this.AgreeWithTerms = agreeWithTerms
	this.TicketRequests = ticketRequests
	return &this
}

// NewCreateUnregisteredTicketRequestWithDefaults instantiates a new CreateUnregisteredTicketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUnregisteredTicketRequestWithDefaults() *CreateUnregisteredTicketRequest {
	this := CreateUnregisteredTicketRequest{}
	var agreeWithTerms bool = false
	this.AgreeWithTerms = agreeWithTerms
	return &this
}

// GetAgreeWithTerms returns the AgreeWithTerms field value
func (o *CreateUnregisteredTicketRequest) GetAgreeWithTerms() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgreeWithTerms
}

// GetAgreeWithTermsOk returns a tuple with the AgreeWithTerms field value
// and a boolean to check if the value has been set.
func (o *CreateUnregisteredTicketRequest) GetAgreeWithTermsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgreeWithTerms, true
}

// SetAgreeWithTerms sets field value
func (o *CreateUnregisteredTicketRequest) SetAgreeWithTerms(v bool) {
	o.AgreeWithTerms = v
}

// GetTicketRequests returns the TicketRequests field value
func (o *CreateUnregisteredTicketRequest) GetTicketRequests() []CreateTicketRequest {
	if o == nil {
		var ret []CreateTicketRequest
		return ret
	}

	return o.TicketRequests
}

// GetTicketRequestsOk returns a tuple with the TicketRequests field value
// and a boolean to check if the value has been set.
func (o *CreateUnregisteredTicketRequest) GetTicketRequestsOk() ([]CreateTicketRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.TicketRequests, true
}

// SetTicketRequests sets field value
func (o *CreateUnregisteredTicketRequest) SetTicketRequests(v []CreateTicketRequest) {
	o.TicketRequests = v
}

// GetAffiliateCode returns the AffiliateCode field value if set, zero value otherwise.
func (o *CreateUnregisteredTicketRequest) GetAffiliateCode() string {
	if o == nil || IsNil(o.AffiliateCode) {
		var ret string
		return ret
	}
	return *o.AffiliateCode
}

// GetAffiliateCodeOk returns a tuple with the AffiliateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUnregisteredTicketRequest) GetAffiliateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AffiliateCode) {
		return nil, false
	}
	return o.AffiliateCode, true
}

// HasAffiliateCode returns a boolean if a field has been set.
func (o *CreateUnregisteredTicketRequest) HasAffiliateCode() bool {
	if o != nil && !IsNil(o.AffiliateCode) {
		return true
	}

	return false
}

// SetAffiliateCode gets a reference to the given string and assigns it to the AffiliateCode field.
func (o *CreateUnregisteredTicketRequest) SetAffiliateCode(v string) {
	o.AffiliateCode = &v
}

func (o CreateUnregisteredTicketRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUnregisteredTicketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agreeWithTerms"] = o.AgreeWithTerms
	toSerialize["ticketRequests"] = o.TicketRequests
	if !IsNil(o.AffiliateCode) {
		toSerialize["affiliateCode"] = o.AffiliateCode
	}
	return toSerialize, nil
}

func (o *CreateUnregisteredTicketRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agreeWithTerms",
		"ticketRequests",
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

	varCreateUnregisteredTicketRequest := _CreateUnregisteredTicketRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUnregisteredTicketRequest)

	if err != nil {
		return err
	}

	*o = CreateUnregisteredTicketRequest(varCreateUnregisteredTicketRequest)

	return err
}

type NullableCreateUnregisteredTicketRequest struct {
	value *CreateUnregisteredTicketRequest
	isSet bool
}

func (v NullableCreateUnregisteredTicketRequest) Get() *CreateUnregisteredTicketRequest {
	return v.value
}

func (v *NullableCreateUnregisteredTicketRequest) Set(val *CreateUnregisteredTicketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUnregisteredTicketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUnregisteredTicketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUnregisteredTicketRequest(val *CreateUnregisteredTicketRequest) *NullableCreateUnregisteredTicketRequest {
	return &NullableCreateUnregisteredTicketRequest{value: val, isSet: true}
}

func (v NullableCreateUnregisteredTicketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUnregisteredTicketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
