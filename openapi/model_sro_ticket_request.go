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

// checks if the SroTicketRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SroTicketRequest{}

// SroTicketRequest struct for SroTicketRequest
type SroTicketRequest struct {
	RouteId       string               `json:"routeId"`
	PriceSourceId string               `json:"priceSourceId"`
	SeatClass     string               `json:"seatClass"`
	Sections      []SroTicketSections  `json:"sections"`
	Passengers    []SroTicketPassenger `json:"passengers"`
}

type _SroTicketRequest SroTicketRequest

// NewSroTicketRequest instantiates a new SroTicketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSroTicketRequest(routeId string, priceSourceId string, seatClass string, sections []SroTicketSections, passengers []SroTicketPassenger) *SroTicketRequest {
	this := SroTicketRequest{}
	this.RouteId = routeId
	this.PriceSourceId = priceSourceId
	this.SeatClass = seatClass
	this.Sections = sections
	this.Passengers = passengers
	return &this
}

// NewSroTicketRequestWithDefaults instantiates a new SroTicketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSroTicketRequestWithDefaults() *SroTicketRequest {
	this := SroTicketRequest{}
	return &this
}

// GetRouteId returns the RouteId field value
func (o *SroTicketRequest) GetRouteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value
// and a boolean to check if the value has been set.
func (o *SroTicketRequest) GetRouteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteId, true
}

// SetRouteId sets field value
func (o *SroTicketRequest) SetRouteId(v string) {
	o.RouteId = v
}

// GetPriceSourceId returns the PriceSourceId field value
func (o *SroTicketRequest) GetPriceSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PriceSourceId
}

// GetPriceSourceIdOk returns a tuple with the PriceSourceId field value
// and a boolean to check if the value has been set.
func (o *SroTicketRequest) GetPriceSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceSourceId, true
}

// SetPriceSourceId sets field value
func (o *SroTicketRequest) SetPriceSourceId(v string) {
	o.PriceSourceId = v
}

// GetSeatClass returns the SeatClass field value
func (o *SroTicketRequest) GetSeatClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SeatClass
}

// GetSeatClassOk returns a tuple with the SeatClass field value
// and a boolean to check if the value has been set.
func (o *SroTicketRequest) GetSeatClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeatClass, true
}

// SetSeatClass sets field value
func (o *SroTicketRequest) SetSeatClass(v string) {
	o.SeatClass = v
}

// GetSections returns the Sections field value
func (o *SroTicketRequest) GetSections() []SroTicketSections {
	if o == nil {
		var ret []SroTicketSections
		return ret
	}

	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value
// and a boolean to check if the value has been set.
func (o *SroTicketRequest) GetSectionsOk() ([]SroTicketSections, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sections, true
}

// SetSections sets field value
func (o *SroTicketRequest) SetSections(v []SroTicketSections) {
	o.Sections = v
}

// GetPassengers returns the Passengers field value
func (o *SroTicketRequest) GetPassengers() []SroTicketPassenger {
	if o == nil {
		var ret []SroTicketPassenger
		return ret
	}

	return o.Passengers
}

// GetPassengersOk returns a tuple with the Passengers field value
// and a boolean to check if the value has been set.
func (o *SroTicketRequest) GetPassengersOk() ([]SroTicketPassenger, bool) {
	if o == nil {
		return nil, false
	}
	return o.Passengers, true
}

// SetPassengers sets field value
func (o *SroTicketRequest) SetPassengers(v []SroTicketPassenger) {
	o.Passengers = v
}

func (o SroTicketRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SroTicketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["routeId"] = o.RouteId
	toSerialize["priceSourceId"] = o.PriceSourceId
	toSerialize["seatClass"] = o.SeatClass
	toSerialize["sections"] = o.Sections
	toSerialize["passengers"] = o.Passengers
	return toSerialize, nil
}

func (o *SroTicketRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"routeId",
		"priceSourceId",
		"seatClass",
		"sections",
		"passengers",
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

	varSroTicketRequest := _SroTicketRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSroTicketRequest)

	if err != nil {
		return err
	}

	*o = SroTicketRequest(varSroTicketRequest)

	return err
}

type NullableSroTicketRequest struct {
	value *SroTicketRequest
	isSet bool
}

func (v NullableSroTicketRequest) Get() *SroTicketRequest {
	return v.value
}

func (v *NullableSroTicketRequest) Set(val *SroTicketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSroTicketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSroTicketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSroTicketRequest(val *SroTicketRequest) *NullableSroTicketRequest {
	return &NullableSroTicketRequest{value: val, isSet: true}
}

func (v NullableSroTicketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSroTicketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
