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

// checks if the MessageField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageField{}

// MessageField struct for MessageField
type MessageField struct {
	// Error/invalid field title
	Key string `json:"key"`
	// Error/validation message to current field
	Value string `json:"value"`
}

type _MessageField MessageField

// NewMessageField instantiates a new MessageField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageField(key string, value string) *MessageField {
	this := MessageField{}
	this.Key = key
	this.Value = value
	return &this
}

// NewMessageFieldWithDefaults instantiates a new MessageField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageFieldWithDefaults() *MessageField {
	this := MessageField{}
	return &this
}

// GetKey returns the Key field value
func (o *MessageField) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MessageField) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MessageField) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *MessageField) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MessageField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MessageField) SetValue(v string) {
	o.Value = v
}

func (o MessageField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *MessageField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varMessageField := _MessageField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageField)

	if err != nil {
		return err
	}

	*o = MessageField(varMessageField)

	return err
}

type NullableMessageField struct {
	value *MessageField
	isSet bool
}

func (v NullableMessageField) Get() *MessageField {
	return v.value
}

func (v *NullableMessageField) Set(val *MessageField) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageField) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageField(val *MessageField) *NullableMessageField {
	return &NullableMessageField{value: val, isSet: true}
}

func (v NullableMessageField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
