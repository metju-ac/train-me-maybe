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

// checks if the SroPriceConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SroPriceConditions{}

// SroPriceConditions struct for SroPriceConditions
type SroPriceConditions struct {
	Descriptions                   SroPriceConditionsDescriptions `json:"descriptions"`
	RefundToOriginalSourcePossible bool                           `json:"refundToOriginalSourcePossible"`
	CancelCharge                   float32                        `json:"cancelCharge"`
	CancelCharges                  []SroCancelCharge              `json:"cancelCharges"`
}

type _SroPriceConditions SroPriceConditions

// NewSroPriceConditions instantiates a new SroPriceConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSroPriceConditions(descriptions SroPriceConditionsDescriptions, refundToOriginalSourcePossible bool, cancelCharge float32, cancelCharges []SroCancelCharge) *SroPriceConditions {
	this := SroPriceConditions{}
	this.Descriptions = descriptions
	this.RefundToOriginalSourcePossible = refundToOriginalSourcePossible
	this.CancelCharge = cancelCharge
	this.CancelCharges = cancelCharges
	return &this
}

// NewSroPriceConditionsWithDefaults instantiates a new SroPriceConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSroPriceConditionsWithDefaults() *SroPriceConditions {
	this := SroPriceConditions{}
	return &this
}

// GetDescriptions returns the Descriptions field value
func (o *SroPriceConditions) GetDescriptions() SroPriceConditionsDescriptions {
	if o == nil {
		var ret SroPriceConditionsDescriptions
		return ret
	}

	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value
// and a boolean to check if the value has been set.
func (o *SroPriceConditions) GetDescriptionsOk() (*SroPriceConditionsDescriptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Descriptions, true
}

// SetDescriptions sets field value
func (o *SroPriceConditions) SetDescriptions(v SroPriceConditionsDescriptions) {
	o.Descriptions = v
}

// GetRefundToOriginalSourcePossible returns the RefundToOriginalSourcePossible field value
func (o *SroPriceConditions) GetRefundToOriginalSourcePossible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RefundToOriginalSourcePossible
}

// GetRefundToOriginalSourcePossibleOk returns a tuple with the RefundToOriginalSourcePossible field value
// and a boolean to check if the value has been set.
func (o *SroPriceConditions) GetRefundToOriginalSourcePossibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefundToOriginalSourcePossible, true
}

// SetRefundToOriginalSourcePossible sets field value
func (o *SroPriceConditions) SetRefundToOriginalSourcePossible(v bool) {
	o.RefundToOriginalSourcePossible = v
}

// GetCancelCharge returns the CancelCharge field value
func (o *SroPriceConditions) GetCancelCharge() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CancelCharge
}

// GetCancelChargeOk returns a tuple with the CancelCharge field value
// and a boolean to check if the value has been set.
func (o *SroPriceConditions) GetCancelChargeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelCharge, true
}

// SetCancelCharge sets field value
func (o *SroPriceConditions) SetCancelCharge(v float32) {
	o.CancelCharge = v
}

// GetCancelCharges returns the CancelCharges field value
func (o *SroPriceConditions) GetCancelCharges() []SroCancelCharge {
	if o == nil {
		var ret []SroCancelCharge
		return ret
	}

	return o.CancelCharges
}

// GetCancelChargesOk returns a tuple with the CancelCharges field value
// and a boolean to check if the value has been set.
func (o *SroPriceConditions) GetCancelChargesOk() ([]SroCancelCharge, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelCharges, true
}

// SetCancelCharges sets field value
func (o *SroPriceConditions) SetCancelCharges(v []SroCancelCharge) {
	o.CancelCharges = v
}

func (o SroPriceConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SroPriceConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descriptions"] = o.Descriptions
	toSerialize["refundToOriginalSourcePossible"] = o.RefundToOriginalSourcePossible
	toSerialize["cancelCharge"] = o.CancelCharge
	toSerialize["cancelCharges"] = o.CancelCharges
	return toSerialize, nil
}

func (o *SroPriceConditions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"descriptions",
		"refundToOriginalSourcePossible",
		"cancelCharge",
		"cancelCharges",
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

	varSroPriceConditions := _SroPriceConditions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSroPriceConditions)

	if err != nil {
		return err
	}

	*o = SroPriceConditions(varSroPriceConditions)

	return err
}

type NullableSroPriceConditions struct {
	value *SroPriceConditions
	isSet bool
}

func (v NullableSroPriceConditions) Get() *SroPriceConditions {
	return v.value
}

func (v *NullableSroPriceConditions) Set(val *SroPriceConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableSroPriceConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableSroPriceConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSroPriceConditions(val *SroPriceConditions) *NullableSroPriceConditions {
	return &NullableSroPriceConditions{value: val, isSet: true}
}

func (v NullableSroPriceConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSroPriceConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
