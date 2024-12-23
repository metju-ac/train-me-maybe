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

// checks if the TariffNotifications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TariffNotifications{}

// TariffNotifications struct for TariffNotifications
type TariffNotifications struct {
	// Notification header translation
	Title string `json:"title"`
	// Notification description translation
	Description string   `json:"description"`
	Content     []string `json:"content"`
}

type _TariffNotifications TariffNotifications

// NewTariffNotifications instantiates a new TariffNotifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTariffNotifications(title string, description string, content []string) *TariffNotifications {
	this := TariffNotifications{}
	this.Title = title
	this.Description = description
	this.Content = content
	return &this
}

// NewTariffNotificationsWithDefaults instantiates a new TariffNotifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTariffNotificationsWithDefaults() *TariffNotifications {
	this := TariffNotifications{}
	return &this
}

// GetTitle returns the Title field value
func (o *TariffNotifications) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TariffNotifications) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TariffNotifications) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *TariffNotifications) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TariffNotifications) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TariffNotifications) SetDescription(v string) {
	o.Description = v
}

// GetContent returns the Content field value
func (o *TariffNotifications) GetContent() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *TariffNotifications) GetContentOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content, true
}

// SetContent sets field value
func (o *TariffNotifications) SetContent(v []string) {
	o.Content = v
}

func (o TariffNotifications) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TariffNotifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *TariffNotifications) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"description",
		"content",
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

	varTariffNotifications := _TariffNotifications{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTariffNotifications)

	if err != nil {
		return err
	}

	*o = TariffNotifications(varTariffNotifications)

	return err
}

type NullableTariffNotifications struct {
	value *TariffNotifications
	isSet bool
}

func (v NullableTariffNotifications) Get() *TariffNotifications {
	return v.value
}

func (v *NullableTariffNotifications) Set(val *TariffNotifications) {
	v.value = val
	v.isSet = true
}

func (v NullableTariffNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableTariffNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTariffNotifications(val *TariffNotifications) *NullableTariffNotifications {
	return &NullableTariffNotifications{value: val, isSet: true}
}

func (v NullableTariffNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTariffNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
