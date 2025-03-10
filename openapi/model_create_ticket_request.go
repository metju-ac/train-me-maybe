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

// checks if the CreateTicketRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTicketRequest{}

// CreateTicketRequest TODO: passenger by měl mít příznak isInsuarence + data...
type CreateTicketRequest struct {
	Route          CreateTicketRouteRequest `json:"route"`
	SelectedAddons []SelectedAddon          `json:"selectedAddons,omitempty"`
	Passengers     []Passenger              `json:"passengers"`
	// Flat rate discount from fare price (does not apply on addons and charges). Applies first (before percentual discount)
	CodeDiscount *string `json:"codeDiscount,omitempty"`
	// Percentual discount from fare price (does not apply on addons and charges). Applies as seconds (after flat rate discount)
	PercentualDiscountIds []int64 `json:"percentualDiscountIds,omitempty"`
}

type _CreateTicketRequest CreateTicketRequest

// NewCreateTicketRequest instantiates a new CreateTicketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTicketRequest(route CreateTicketRouteRequest, passengers []Passenger) *CreateTicketRequest {
	this := CreateTicketRequest{}
	this.Route = route
	this.Passengers = passengers
	return &this
}

// NewCreateTicketRequestWithDefaults instantiates a new CreateTicketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTicketRequestWithDefaults() *CreateTicketRequest {
	this := CreateTicketRequest{}
	return &this
}

// GetRoute returns the Route field value
func (o *CreateTicketRequest) GetRoute() CreateTicketRouteRequest {
	if o == nil {
		var ret CreateTicketRouteRequest
		return ret
	}

	return o.Route
}

// GetRouteOk returns a tuple with the Route field value
// and a boolean to check if the value has been set.
func (o *CreateTicketRequest) GetRouteOk() (*CreateTicketRouteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Route, true
}

// SetRoute sets field value
func (o *CreateTicketRequest) SetRoute(v CreateTicketRouteRequest) {
	o.Route = v
}

// GetSelectedAddons returns the SelectedAddons field value if set, zero value otherwise.
func (o *CreateTicketRequest) GetSelectedAddons() []SelectedAddon {
	if o == nil || IsNil(o.SelectedAddons) {
		var ret []SelectedAddon
		return ret
	}
	return o.SelectedAddons
}

// GetSelectedAddonsOk returns a tuple with the SelectedAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTicketRequest) GetSelectedAddonsOk() ([]SelectedAddon, bool) {
	if o == nil || IsNil(o.SelectedAddons) {
		return nil, false
	}
	return o.SelectedAddons, true
}

// HasSelectedAddons returns a boolean if a field has been set.
func (o *CreateTicketRequest) HasSelectedAddons() bool {
	if o != nil && !IsNil(o.SelectedAddons) {
		return true
	}

	return false
}

// SetSelectedAddons gets a reference to the given []SelectedAddon and assigns it to the SelectedAddons field.
func (o *CreateTicketRequest) SetSelectedAddons(v []SelectedAddon) {
	o.SelectedAddons = v
}

// GetPassengers returns the Passengers field value
func (o *CreateTicketRequest) GetPassengers() []Passenger {
	if o == nil {
		var ret []Passenger
		return ret
	}

	return o.Passengers
}

// GetPassengersOk returns a tuple with the Passengers field value
// and a boolean to check if the value has been set.
func (o *CreateTicketRequest) GetPassengersOk() ([]Passenger, bool) {
	if o == nil {
		return nil, false
	}
	return o.Passengers, true
}

// SetPassengers sets field value
func (o *CreateTicketRequest) SetPassengers(v []Passenger) {
	o.Passengers = v
}

// GetCodeDiscount returns the CodeDiscount field value if set, zero value otherwise.
func (o *CreateTicketRequest) GetCodeDiscount() string {
	if o == nil || IsNil(o.CodeDiscount) {
		var ret string
		return ret
	}
	return *o.CodeDiscount
}

// GetCodeDiscountOk returns a tuple with the CodeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTicketRequest) GetCodeDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.CodeDiscount) {
		return nil, false
	}
	return o.CodeDiscount, true
}

// HasCodeDiscount returns a boolean if a field has been set.
func (o *CreateTicketRequest) HasCodeDiscount() bool {
	if o != nil && !IsNil(o.CodeDiscount) {
		return true
	}

	return false
}

// SetCodeDiscount gets a reference to the given string and assigns it to the CodeDiscount field.
func (o *CreateTicketRequest) SetCodeDiscount(v string) {
	o.CodeDiscount = &v
}

// GetPercentualDiscountIds returns the PercentualDiscountIds field value if set, zero value otherwise.
func (o *CreateTicketRequest) GetPercentualDiscountIds() []int64 {
	if o == nil || IsNil(o.PercentualDiscountIds) {
		var ret []int64
		return ret
	}
	return o.PercentualDiscountIds
}

// GetPercentualDiscountIdsOk returns a tuple with the PercentualDiscountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTicketRequest) GetPercentualDiscountIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PercentualDiscountIds) {
		return nil, false
	}
	return o.PercentualDiscountIds, true
}

// HasPercentualDiscountIds returns a boolean if a field has been set.
func (o *CreateTicketRequest) HasPercentualDiscountIds() bool {
	if o != nil && !IsNil(o.PercentualDiscountIds) {
		return true
	}

	return false
}

// SetPercentualDiscountIds gets a reference to the given []int64 and assigns it to the PercentualDiscountIds field.
func (o *CreateTicketRequest) SetPercentualDiscountIds(v []int64) {
	o.PercentualDiscountIds = v
}

func (o CreateTicketRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTicketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["route"] = o.Route
	if !IsNil(o.SelectedAddons) {
		toSerialize["selectedAddons"] = o.SelectedAddons
	}
	toSerialize["passengers"] = o.Passengers
	if !IsNil(o.CodeDiscount) {
		toSerialize["codeDiscount"] = o.CodeDiscount
	}
	if !IsNil(o.PercentualDiscountIds) {
		toSerialize["percentualDiscountIds"] = o.PercentualDiscountIds
	}
	return toSerialize, nil
}

func (o *CreateTicketRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"route",
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

	varCreateTicketRequest := _CreateTicketRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTicketRequest)

	if err != nil {
		return err
	}

	*o = CreateTicketRequest(varCreateTicketRequest)

	return err
}

type NullableCreateTicketRequest struct {
	value *CreateTicketRequest
	isSet bool
}

func (v NullableCreateTicketRequest) Get() *CreateTicketRequest {
	return v.value
}

func (v *NullableCreateTicketRequest) Set(val *CreateTicketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTicketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTicketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTicketRequest(val *CreateTicketRequest) *NullableCreateTicketRequest {
	return &NullableCreateTicketRequest{value: val, isSet: true}
}

func (v NullableCreateTicketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTicketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
