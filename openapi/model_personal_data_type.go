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
	"fmt"
)

// PersonalDataType the model 'PersonalDataType'
type PersonalDataType string

// List of PersonalDataType
const (
	FIRST_NAME      PersonalDataType = "FIRST_NAME"
	SURNAME         PersonalDataType = "SURNAME"
	BIRTHDAY        PersonalDataType = "BIRTHDAY"
	EMAIL           PersonalDataType = "EMAIL"
	PHONE           PersonalDataType = "PHONE"
	ZIP_CODE        PersonalDataType = "ZIP_CODE"
	PERSONAL_NUMBER PersonalDataType = "PERSONAL_NUMBER"
	STREET          PersonalDataType = "STREET"
	HOUSE_NUMBER    PersonalDataType = "HOUSE_NUMBER"
	CITY            PersonalDataType = "CITY"
)

// All allowed values of PersonalDataType enum
var AllowedPersonalDataTypeEnumValues = []PersonalDataType{
	"FIRST_NAME",
	"SURNAME",
	"BIRTHDAY",
	"EMAIL",
	"PHONE",
	"ZIP_CODE",
	"PERSONAL_NUMBER",
	"STREET",
	"HOUSE_NUMBER",
	"CITY",
}

func (v *PersonalDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersonalDataType(value)
	for _, existing := range AllowedPersonalDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersonalDataType", value)
}

// NewPersonalDataTypeFromValue returns a pointer to a valid PersonalDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersonalDataTypeFromValue(v string) (*PersonalDataType, error) {
	ev := PersonalDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersonalDataType: valid values are %v", v, AllowedPersonalDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersonalDataType) IsValid() bool {
	for _, existing := range AllowedPersonalDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonalDataType value
func (v PersonalDataType) Ptr() *PersonalDataType {
	return &v
}

type NullablePersonalDataType struct {
	value *PersonalDataType
	isSet bool
}

func (v NullablePersonalDataType) Get() *PersonalDataType {
	return v.value
}

func (v *NullablePersonalDataType) Set(val *PersonalDataType) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonalDataType) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonalDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonalDataType(val *PersonalDataType) *NullablePersonalDataType {
	return &NullablePersonalDataType{value: val, isSet: true}
}

func (v NullablePersonalDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonalDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
