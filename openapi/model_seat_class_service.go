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

// checks if the SeatClassService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeatClassService{}

// SeatClassService struct for SeatClassService
type SeatClassService struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	ImageCode string `json:"imageCode"`
}

type _SeatClassService SeatClassService

// NewSeatClassService instantiates a new SeatClassService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeatClassService(code string, name string, imageCode string) *SeatClassService {
	this := SeatClassService{}
	this.Code = code
	this.Name = name
	this.ImageCode = imageCode
	return &this
}

// NewSeatClassServiceWithDefaults instantiates a new SeatClassService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeatClassServiceWithDefaults() *SeatClassService {
	this := SeatClassService{}
	return &this
}

// GetCode returns the Code field value
func (o *SeatClassService) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SeatClassService) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SeatClassService) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *SeatClassService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SeatClassService) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SeatClassService) SetName(v string) {
	o.Name = v
}

// GetImageCode returns the ImageCode field value
func (o *SeatClassService) GetImageCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageCode
}

// GetImageCodeOk returns a tuple with the ImageCode field value
// and a boolean to check if the value has been set.
func (o *SeatClassService) GetImageCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageCode, true
}

// SetImageCode sets field value
func (o *SeatClassService) SetImageCode(v string) {
	o.ImageCode = v
}

func (o SeatClassService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeatClassService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	toSerialize["imageCode"] = o.ImageCode
	return toSerialize, nil
}

func (o *SeatClassService) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
		"imageCode",
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

	varSeatClassService := _SeatClassService{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSeatClassService)

	if err != nil {
		return err
	}

	*o = SeatClassService(varSeatClassService)

	return err
}

type NullableSeatClassService struct {
	value *SeatClassService
	isSet bool
}

func (v NullableSeatClassService) Get() *SeatClassService {
	return v.value
}

func (v *NullableSeatClassService) Set(val *SeatClassService) {
	v.value = val
	v.isSet = true
}

func (v NullableSeatClassService) IsSet() bool {
	return v.isSet
}

func (v *NullableSeatClassService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeatClassService(val *SeatClassService) *NullableSeatClassService {
	return &NullableSeatClassService{value: val, isSet: true}
}

func (v NullableSeatClassService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeatClassService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
