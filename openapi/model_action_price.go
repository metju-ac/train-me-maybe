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

// checks if the ActionPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionPrice{}

// ActionPrice struct for ActionPrice
type ActionPrice struct {
	// Unique identifier of action.
	Id int64 `json:"id"`
	// Action price code.
	Code string `json:"code"`
	// Action price name.
	Name string `json:"name"`
	// Publicly available URL of an action.
	Url string `json:"url"`
	// Action price description.
	Description *string `json:"description,omitempty"`
	// Flag that indicates wheter icon should be shown with action price name.
	ShowIcon bool `json:"showIcon"`
}

type _ActionPrice ActionPrice

// NewActionPrice instantiates a new ActionPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionPrice(id int64, code string, name string, url string, showIcon bool) *ActionPrice {
	this := ActionPrice{}
	this.Id = id
	this.Code = code
	this.Name = name
	this.Url = url
	this.ShowIcon = showIcon
	return &this
}

// NewActionPriceWithDefaults instantiates a new ActionPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionPriceWithDefaults() *ActionPrice {
	this := ActionPrice{}
	return &this
}

// GetId returns the Id field value
func (o *ActionPrice) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActionPrice) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActionPrice) SetId(v int64) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *ActionPrice) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ActionPrice) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ActionPrice) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *ActionPrice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActionPrice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActionPrice) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *ActionPrice) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ActionPrice) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ActionPrice) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ActionPrice) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionPrice) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ActionPrice) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ActionPrice) SetDescription(v string) {
	o.Description = &v
}

// GetShowIcon returns the ShowIcon field value
func (o *ActionPrice) GetShowIcon() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowIcon
}

// GetShowIconOk returns a tuple with the ShowIcon field value
// and a boolean to check if the value has been set.
func (o *ActionPrice) GetShowIconOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowIcon, true
}

// SetShowIcon sets field value
func (o *ActionPrice) SetShowIcon(v bool) {
	o.ShowIcon = v
}

func (o ActionPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["showIcon"] = o.ShowIcon
	return toSerialize, nil
}

func (o *ActionPrice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"code",
		"name",
		"url",
		"showIcon",
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

	varActionPrice := _ActionPrice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionPrice)

	if err != nil {
		return err
	}

	*o = ActionPrice(varActionPrice)

	return err
}

type NullableActionPrice struct {
	value *ActionPrice
	isSet bool
}

func (v NullableActionPrice) Get() *ActionPrice {
	return v.value
}

func (v *NullableActionPrice) Set(val *ActionPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableActionPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableActionPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionPrice(val *ActionPrice) *NullableActionPrice {
	return &NullableActionPrice{value: val, isSet: true}
}

func (v NullableActionPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


