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

// checks if the Notifications type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Notifications{}

// Notifications struct for Notifications
type Notifications struct {
	Newsletter        bool `json:"newsletter"`
	ReservationChange bool `json:"reservationChange"`
	RouteRatingSurvey bool `json:"routeRatingSurvey"`
}

type _Notifications Notifications

// NewNotifications instantiates a new Notifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifications(newsletter bool, reservationChange bool, routeRatingSurvey bool) *Notifications {
	this := Notifications{}
	this.Newsletter = newsletter
	this.ReservationChange = reservationChange
	this.RouteRatingSurvey = routeRatingSurvey
	return &this
}

// NewNotificationsWithDefaults instantiates a new Notifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsWithDefaults() *Notifications {
	this := Notifications{}
	var newsletter bool = false
	this.Newsletter = newsletter
	var reservationChange bool = false
	this.ReservationChange = reservationChange
	var routeRatingSurvey bool = false
	this.RouteRatingSurvey = routeRatingSurvey
	return &this
}

// GetNewsletter returns the Newsletter field value
func (o *Notifications) GetNewsletter() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Newsletter
}

// GetNewsletterOk returns a tuple with the Newsletter field value
// and a boolean to check if the value has been set.
func (o *Notifications) GetNewsletterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Newsletter, true
}

// SetNewsletter sets field value
func (o *Notifications) SetNewsletter(v bool) {
	o.Newsletter = v
}

// GetReservationChange returns the ReservationChange field value
func (o *Notifications) GetReservationChange() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReservationChange
}

// GetReservationChangeOk returns a tuple with the ReservationChange field value
// and a boolean to check if the value has been set.
func (o *Notifications) GetReservationChangeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReservationChange, true
}

// SetReservationChange sets field value
func (o *Notifications) SetReservationChange(v bool) {
	o.ReservationChange = v
}

// GetRouteRatingSurvey returns the RouteRatingSurvey field value
func (o *Notifications) GetRouteRatingSurvey() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RouteRatingSurvey
}

// GetRouteRatingSurveyOk returns a tuple with the RouteRatingSurvey field value
// and a boolean to check if the value has been set.
func (o *Notifications) GetRouteRatingSurveyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteRatingSurvey, true
}

// SetRouteRatingSurvey sets field value
func (o *Notifications) SetRouteRatingSurvey(v bool) {
	o.RouteRatingSurvey = v
}

func (o Notifications) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notifications) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["newsletter"] = o.Newsletter
	toSerialize["reservationChange"] = o.ReservationChange
	toSerialize["routeRatingSurvey"] = o.RouteRatingSurvey
	return toSerialize, nil
}

func (o *Notifications) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newsletter",
		"reservationChange",
		"routeRatingSurvey",
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

	varNotifications := _Notifications{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotifications)

	if err != nil {
		return err
	}

	*o = Notifications(varNotifications)

	return err
}

type NullableNotifications struct {
	value *Notifications
	isSet bool
}

func (v NullableNotifications) Get() *Notifications {
	return v.value
}

func (v *NullableNotifications) Set(val *Notifications) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifications(val *Notifications) *NullableNotifications {
	return &NullableNotifications{value: val, isSet: true}
}

func (v NullableNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
