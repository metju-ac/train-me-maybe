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

// checks if the InlineResponse4003 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineResponse4003{}

// InlineResponse4003 struct for InlineResponse4003
type InlineResponse4003 struct {
	// Error message.
	Message     string                   `json:"message"`
	ErrorFields []map[string]interface{} `json:"errorFields,omitempty"`
}

type _InlineResponse4003 InlineResponse4003

// NewInlineResponse4003 instantiates a new InlineResponse4003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse4003(message string) *InlineResponse4003 {
	this := InlineResponse4003{}
	this.Message = message
	return &this
}

// NewInlineResponse4003WithDefaults instantiates a new InlineResponse4003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse4003WithDefaults() *InlineResponse4003 {
	this := InlineResponse4003{}
	return &this
}

// GetMessage returns the Message field value
func (o *InlineResponse4003) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *InlineResponse4003) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *InlineResponse4003) SetMessage(v string) {
	o.Message = v
}

// GetErrorFields returns the ErrorFields field value if set, zero value otherwise.
func (o *InlineResponse4003) GetErrorFields() []map[string]interface{} {
	if o == nil || IsNil(o.ErrorFields) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ErrorFields
}

// GetErrorFieldsOk returns a tuple with the ErrorFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse4003) GetErrorFieldsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ErrorFields) {
		return nil, false
	}
	return o.ErrorFields, true
}

// HasErrorFields returns a boolean if a field has been set.
func (o *InlineResponse4003) HasErrorFields() bool {
	if o != nil && !IsNil(o.ErrorFields) {
		return true
	}

	return false
}

// SetErrorFields gets a reference to the given []map[string]interface{} and assigns it to the ErrorFields field.
func (o *InlineResponse4003) SetErrorFields(v []map[string]interface{}) {
	o.ErrorFields = v
}

func (o InlineResponse4003) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineResponse4003) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.ErrorFields) {
		toSerialize["errorFields"] = o.ErrorFields
	}
	return toSerialize, nil
}

func (o *InlineResponse4003) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varInlineResponse4003 := _InlineResponse4003{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInlineResponse4003)

	if err != nil {
		return err
	}

	*o = InlineResponse4003(varInlineResponse4003)

	return err
}

type NullableInlineResponse4003 struct {
	value *InlineResponse4003
	isSet bool
}

func (v NullableInlineResponse4003) Get() *InlineResponse4003 {
	return v.value
}

func (v *NullableInlineResponse4003) Set(val *InlineResponse4003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse4003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse4003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse4003(val *InlineResponse4003) *NullableInlineResponse4003 {
	return &NullableInlineResponse4003{value: val, isSet: true}
}

func (v NullableInlineResponse4003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse4003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
