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

// checks if the City type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &City{}

// City struct for City
type City struct {
	// Unique identifier for the city.
	Id int64 `json:"id"`
	// City name.
	Name string `json:"name"`
	// City nicknames or aliases.
	Aliases []string `json:"aliases,omitempty"`
	// Types of stations available in a city.
	StationsTypes []string  `json:"stationsTypes,omitempty"`
	Stations      []Station `json:"stations"`
}

type _City City

// NewCity instantiates a new City object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCity(id int64, name string, stations []Station) *City {
	this := City{}
	this.Id = id
	this.Name = name
	this.Stations = stations
	return &this
}

// NewCityWithDefaults instantiates a new City object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCityWithDefaults() *City {
	this := City{}
	return &this
}

// GetId returns the Id field value
func (o *City) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *City) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *City) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *City) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *City) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *City) SetName(v string) {
	o.Name = v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *City) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *City) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *City) SetAliases(v []string) {
	o.Aliases = v
}

// GetStationsTypes returns the StationsTypes field value if set, zero value otherwise.
func (o *City) GetStationsTypes() []string {
	if o == nil || IsNil(o.StationsTypes) {
		var ret []string
		return ret
	}
	return o.StationsTypes
}

// GetStationsTypesOk returns a tuple with the StationsTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *City) GetStationsTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.StationsTypes) {
		return nil, false
	}
	return o.StationsTypes, true
}

// HasStationsTypes returns a boolean if a field has been set.
func (o *City) HasStationsTypes() bool {
	if o != nil && !IsNil(o.StationsTypes) {
		return true
	}

	return false
}

// SetStationsTypes gets a reference to the given []string and assigns it to the StationsTypes field.
func (o *City) SetStationsTypes(v []string) {
	o.StationsTypes = v
}

// GetStations returns the Stations field value
func (o *City) GetStations() []Station {
	if o == nil {
		var ret []Station
		return ret
	}

	return o.Stations
}

// GetStationsOk returns a tuple with the Stations field value
// and a boolean to check if the value has been set.
func (o *City) GetStationsOk() ([]Station, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stations, true
}

// SetStations sets field value
func (o *City) SetStations(v []Station) {
	o.Stations = v
}

func (o City) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o City) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.StationsTypes) {
		toSerialize["stationsTypes"] = o.StationsTypes
	}
	toSerialize["stations"] = o.Stations
	return toSerialize, nil
}

func (o *City) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"stations",
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

	varCity := _City{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCity)

	if err != nil {
		return err
	}

	*o = City(varCity)

	return err
}

type NullableCity struct {
	value *City
	isSet bool
}

func (v NullableCity) Get() *City {
	return v.value
}

func (v *NullableCity) Set(val *City) {
	v.value = val
	v.isSet = true
}

func (v NullableCity) IsSet() bool {
	return v.isSet
}

func (v *NullableCity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCity(val *City) *NullableCity {
	return &NullableCity{value: val, isSet: true}
}

func (v NullableCity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
