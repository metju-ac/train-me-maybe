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
	"time"
)

// checks if the SimpleRoute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleRoute{}

// SimpleRoute struct for SimpleRoute
type SimpleRoute struct {
	// Unique identifier of a route. Consists of unique sections identifiers separated by commas.
	Id string `json:"id"`
	// Unique identifier of departure station.
	DepartureStationId *int64 `json:"departureStationId,omitempty"`
	// Departure time.
	DepartureTime time.Time `json:"departureTime"`
	// Unique identifier of arrival station.
	ArrivalStationId *int64 `json:"arrivalStationId,omitempty"`
	// Arrival time.
	ArrivalTime time.Time `json:"arrivalTime"`
	// Vehicle type.
	VehicleTypes []VehicleType `json:"vehicleTypes"`
	// Number of transfers between stations.
	TransfersCount *int32 `json:"transfersCount,omitempty"`
	// Number of available free seats through all sections.
	FreeSeatsCount int32 `json:"freeSeatsCount"`
	// Minimum ticket price for open account.
	PriceFrom float32 `json:"priceFrom"`
	// Maximum ticket price for open account.
	PriceTo *float32 `json:"priceTo,omitempty"`
	// Minimum ticket price for RegioJet Pay.
	CreditPriceFrom float32 `json:"creditPriceFrom"`
	// Maximum ticket price for RegioJet Pay.
	CreditPriceTo *float32 `json:"creditPriceTo,omitempty"`
	// Number of available prices.
	PricesCount int32 `json:"pricesCount"`
	// `true` if any price is action price, otherwise `false`.
	ActionPrice bool `json:"actionPrice"`
	// `true` if there is surcharge on this route, otherwise `false`.
	Surcharge bool `json:"surcharge"`
	// Notice of any extraordinary events on route or traffic problems.
	Notices bool `json:"notices"`
	// `true` if this route (or its part) have support (e.g. in form of additional railroad cars), otherwise `false`.
	Support bool `json:"support"`
	// `true` if route is interstate, `false` if route is international.
	NationalTrip *bool `json:"nationalTrip,omitempty"`
	// `true` if at least one class have enough free seats for reservation, otherwise `false`.
	Bookable bool `json:"bookable"`
	// Textual information about the first delay on the route.
	Delay *string `json:"delay,omitempty"`
	// Textual information about route's travel time.
	TravelTime *string `json:"travelTime,omitempty"`
	// Vehicle standard key. Detailed data about every available vehicle standard can be obtained from `/consts/vehicleStandards`.
	VehicleStandardKey *string `json:"vehicleStandardKey,omitempty"`
}

type _SimpleRoute SimpleRoute

// NewSimpleRoute instantiates a new SimpleRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleRoute(id string, departureTime time.Time, arrivalTime time.Time, vehicleTypes []VehicleType, freeSeatsCount int32, priceFrom float32, creditPriceFrom float32, pricesCount int32, actionPrice bool, surcharge bool, notices bool, support bool, bookable bool) *SimpleRoute {
	this := SimpleRoute{}
	this.Id = id
	this.DepartureTime = departureTime
	this.ArrivalTime = arrivalTime
	this.VehicleTypes = vehicleTypes
	this.FreeSeatsCount = freeSeatsCount
	this.PriceFrom = priceFrom
	this.CreditPriceFrom = creditPriceFrom
	this.PricesCount = pricesCount
	this.ActionPrice = actionPrice
	this.Surcharge = surcharge
	this.Notices = notices
	this.Support = support
	this.Bookable = bookable
	return &this
}

// NewSimpleRouteWithDefaults instantiates a new SimpleRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleRouteWithDefaults() *SimpleRoute {
	this := SimpleRoute{}
	return &this
}

// GetId returns the Id field value
func (o *SimpleRoute) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimpleRoute) SetId(v string) {
	o.Id = v
}

// GetDepartureStationId returns the DepartureStationId field value if set, zero value otherwise.
func (o *SimpleRoute) GetDepartureStationId() int64 {
	if o == nil || IsNil(o.DepartureStationId) {
		var ret int64
		return ret
	}
	return *o.DepartureStationId
}

// GetDepartureStationIdOk returns a tuple with the DepartureStationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetDepartureStationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DepartureStationId) {
		return nil, false
	}
	return o.DepartureStationId, true
}

// HasDepartureStationId returns a boolean if a field has been set.
func (o *SimpleRoute) HasDepartureStationId() bool {
	if o != nil && !IsNil(o.DepartureStationId) {
		return true
	}

	return false
}

// SetDepartureStationId gets a reference to the given int64 and assigns it to the DepartureStationId field.
func (o *SimpleRoute) SetDepartureStationId(v int64) {
	o.DepartureStationId = &v
}

// GetDepartureTime returns the DepartureTime field value
func (o *SimpleRoute) GetDepartureTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DepartureTime
}

// GetDepartureTimeOk returns a tuple with the DepartureTime field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetDepartureTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepartureTime, true
}

// SetDepartureTime sets field value
func (o *SimpleRoute) SetDepartureTime(v time.Time) {
	o.DepartureTime = v
}

// GetArrivalStationId returns the ArrivalStationId field value if set, zero value otherwise.
func (o *SimpleRoute) GetArrivalStationId() int64 {
	if o == nil || IsNil(o.ArrivalStationId) {
		var ret int64
		return ret
	}
	return *o.ArrivalStationId
}

// GetArrivalStationIdOk returns a tuple with the ArrivalStationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetArrivalStationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ArrivalStationId) {
		return nil, false
	}
	return o.ArrivalStationId, true
}

// HasArrivalStationId returns a boolean if a field has been set.
func (o *SimpleRoute) HasArrivalStationId() bool {
	if o != nil && !IsNil(o.ArrivalStationId) {
		return true
	}

	return false
}

// SetArrivalStationId gets a reference to the given int64 and assigns it to the ArrivalStationId field.
func (o *SimpleRoute) SetArrivalStationId(v int64) {
	o.ArrivalStationId = &v
}

// GetArrivalTime returns the ArrivalTime field value
func (o *SimpleRoute) GetArrivalTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ArrivalTime
}

// GetArrivalTimeOk returns a tuple with the ArrivalTime field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetArrivalTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArrivalTime, true
}

// SetArrivalTime sets field value
func (o *SimpleRoute) SetArrivalTime(v time.Time) {
	o.ArrivalTime = v
}

// GetVehicleTypes returns the VehicleTypes field value
func (o *SimpleRoute) GetVehicleTypes() []VehicleType {
	if o == nil {
		var ret []VehicleType
		return ret
	}

	return o.VehicleTypes
}

// GetVehicleTypesOk returns a tuple with the VehicleTypes field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetVehicleTypesOk() ([]VehicleType, bool) {
	if o == nil {
		return nil, false
	}
	return o.VehicleTypes, true
}

// SetVehicleTypes sets field value
func (o *SimpleRoute) SetVehicleTypes(v []VehicleType) {
	o.VehicleTypes = v
}

// GetTransfersCount returns the TransfersCount field value if set, zero value otherwise.
func (o *SimpleRoute) GetTransfersCount() int32 {
	if o == nil || IsNil(o.TransfersCount) {
		var ret int32
		return ret
	}
	return *o.TransfersCount
}

// GetTransfersCountOk returns a tuple with the TransfersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetTransfersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TransfersCount) {
		return nil, false
	}
	return o.TransfersCount, true
}

// HasTransfersCount returns a boolean if a field has been set.
func (o *SimpleRoute) HasTransfersCount() bool {
	if o != nil && !IsNil(o.TransfersCount) {
		return true
	}

	return false
}

// SetTransfersCount gets a reference to the given int32 and assigns it to the TransfersCount field.
func (o *SimpleRoute) SetTransfersCount(v int32) {
	o.TransfersCount = &v
}

// GetFreeSeatsCount returns the FreeSeatsCount field value
func (o *SimpleRoute) GetFreeSeatsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FreeSeatsCount
}

// GetFreeSeatsCountOk returns a tuple with the FreeSeatsCount field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetFreeSeatsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeSeatsCount, true
}

// SetFreeSeatsCount sets field value
func (o *SimpleRoute) SetFreeSeatsCount(v int32) {
	o.FreeSeatsCount = v
}

// GetPriceFrom returns the PriceFrom field value
func (o *SimpleRoute) GetPriceFrom() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PriceFrom
}

// GetPriceFromOk returns a tuple with the PriceFrom field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetPriceFromOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceFrom, true
}

// SetPriceFrom sets field value
func (o *SimpleRoute) SetPriceFrom(v float32) {
	o.PriceFrom = v
}

// GetPriceTo returns the PriceTo field value if set, zero value otherwise.
func (o *SimpleRoute) GetPriceTo() float32 {
	if o == nil || IsNil(o.PriceTo) {
		var ret float32
		return ret
	}
	return *o.PriceTo
}

// GetPriceToOk returns a tuple with the PriceTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetPriceToOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceTo) {
		return nil, false
	}
	return o.PriceTo, true
}

// HasPriceTo returns a boolean if a field has been set.
func (o *SimpleRoute) HasPriceTo() bool {
	if o != nil && !IsNil(o.PriceTo) {
		return true
	}

	return false
}

// SetPriceTo gets a reference to the given float32 and assigns it to the PriceTo field.
func (o *SimpleRoute) SetPriceTo(v float32) {
	o.PriceTo = &v
}

// GetCreditPriceFrom returns the CreditPriceFrom field value
func (o *SimpleRoute) GetCreditPriceFrom() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreditPriceFrom
}

// GetCreditPriceFromOk returns a tuple with the CreditPriceFrom field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetCreditPriceFromOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditPriceFrom, true
}

// SetCreditPriceFrom sets field value
func (o *SimpleRoute) SetCreditPriceFrom(v float32) {
	o.CreditPriceFrom = v
}

// GetCreditPriceTo returns the CreditPriceTo field value if set, zero value otherwise.
func (o *SimpleRoute) GetCreditPriceTo() float32 {
	if o == nil || IsNil(o.CreditPriceTo) {
		var ret float32
		return ret
	}
	return *o.CreditPriceTo
}

// GetCreditPriceToOk returns a tuple with the CreditPriceTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetCreditPriceToOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditPriceTo) {
		return nil, false
	}
	return o.CreditPriceTo, true
}

// HasCreditPriceTo returns a boolean if a field has been set.
func (o *SimpleRoute) HasCreditPriceTo() bool {
	if o != nil && !IsNil(o.CreditPriceTo) {
		return true
	}

	return false
}

// SetCreditPriceTo gets a reference to the given float32 and assigns it to the CreditPriceTo field.
func (o *SimpleRoute) SetCreditPriceTo(v float32) {
	o.CreditPriceTo = &v
}

// GetPricesCount returns the PricesCount field value
func (o *SimpleRoute) GetPricesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PricesCount
}

// GetPricesCountOk returns a tuple with the PricesCount field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetPricesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricesCount, true
}

// SetPricesCount sets field value
func (o *SimpleRoute) SetPricesCount(v int32) {
	o.PricesCount = v
}

// GetActionPrice returns the ActionPrice field value
func (o *SimpleRoute) GetActionPrice() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ActionPrice
}

// GetActionPriceOk returns a tuple with the ActionPrice field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetActionPriceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionPrice, true
}

// SetActionPrice sets field value
func (o *SimpleRoute) SetActionPrice(v bool) {
	o.ActionPrice = v
}

// GetSurcharge returns the Surcharge field value
func (o *SimpleRoute) GetSurcharge() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Surcharge
}

// GetSurchargeOk returns a tuple with the Surcharge field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetSurchargeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Surcharge, true
}

// SetSurcharge sets field value
func (o *SimpleRoute) SetSurcharge(v bool) {
	o.Surcharge = v
}

// GetNotices returns the Notices field value
func (o *SimpleRoute) GetNotices() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Notices
}

// GetNoticesOk returns a tuple with the Notices field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetNoticesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notices, true
}

// SetNotices sets field value
func (o *SimpleRoute) SetNotices(v bool) {
	o.Notices = v
}

// GetSupport returns the Support field value
func (o *SimpleRoute) GetSupport() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Support
}

// GetSupportOk returns a tuple with the Support field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetSupportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Support, true
}

// SetSupport sets field value
func (o *SimpleRoute) SetSupport(v bool) {
	o.Support = v
}

// GetNationalTrip returns the NationalTrip field value if set, zero value otherwise.
func (o *SimpleRoute) GetNationalTrip() bool {
	if o == nil || IsNil(o.NationalTrip) {
		var ret bool
		return ret
	}
	return *o.NationalTrip
}

// GetNationalTripOk returns a tuple with the NationalTrip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetNationalTripOk() (*bool, bool) {
	if o == nil || IsNil(o.NationalTrip) {
		return nil, false
	}
	return o.NationalTrip, true
}

// HasNationalTrip returns a boolean if a field has been set.
func (o *SimpleRoute) HasNationalTrip() bool {
	if o != nil && !IsNil(o.NationalTrip) {
		return true
	}

	return false
}

// SetNationalTrip gets a reference to the given bool and assigns it to the NationalTrip field.
func (o *SimpleRoute) SetNationalTrip(v bool) {
	o.NationalTrip = &v
}

// GetBookable returns the Bookable field value
func (o *SimpleRoute) GetBookable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bookable
}

// GetBookableOk returns a tuple with the Bookable field value
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetBookableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bookable, true
}

// SetBookable sets field value
func (o *SimpleRoute) SetBookable(v bool) {
	o.Bookable = v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *SimpleRoute) GetDelay() string {
	if o == nil || IsNil(o.Delay) {
		var ret string
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetDelayOk() (*string, bool) {
	if o == nil || IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *SimpleRoute) HasDelay() bool {
	if o != nil && !IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given string and assigns it to the Delay field.
func (o *SimpleRoute) SetDelay(v string) {
	o.Delay = &v
}

// GetTravelTime returns the TravelTime field value if set, zero value otherwise.
func (o *SimpleRoute) GetTravelTime() string {
	if o == nil || IsNil(o.TravelTime) {
		var ret string
		return ret
	}
	return *o.TravelTime
}

// GetTravelTimeOk returns a tuple with the TravelTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetTravelTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TravelTime) {
		return nil, false
	}
	return o.TravelTime, true
}

// HasTravelTime returns a boolean if a field has been set.
func (o *SimpleRoute) HasTravelTime() bool {
	if o != nil && !IsNil(o.TravelTime) {
		return true
	}

	return false
}

// SetTravelTime gets a reference to the given string and assigns it to the TravelTime field.
func (o *SimpleRoute) SetTravelTime(v string) {
	o.TravelTime = &v
}

// GetVehicleStandardKey returns the VehicleStandardKey field value if set, zero value otherwise.
func (o *SimpleRoute) GetVehicleStandardKey() string {
	if o == nil || IsNil(o.VehicleStandardKey) {
		var ret string
		return ret
	}
	return *o.VehicleStandardKey
}

// GetVehicleStandardKeyOk returns a tuple with the VehicleStandardKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleRoute) GetVehicleStandardKeyOk() (*string, bool) {
	if o == nil || IsNil(o.VehicleStandardKey) {
		return nil, false
	}
	return o.VehicleStandardKey, true
}

// HasVehicleStandardKey returns a boolean if a field has been set.
func (o *SimpleRoute) HasVehicleStandardKey() bool {
	if o != nil && !IsNil(o.VehicleStandardKey) {
		return true
	}

	return false
}

// SetVehicleStandardKey gets a reference to the given string and assigns it to the VehicleStandardKey field.
func (o *SimpleRoute) SetVehicleStandardKey(v string) {
	o.VehicleStandardKey = &v
}

func (o SimpleRoute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.DepartureStationId) {
		toSerialize["departureStationId"] = o.DepartureStationId
	}
	toSerialize["departureTime"] = o.DepartureTime
	if !IsNil(o.ArrivalStationId) {
		toSerialize["arrivalStationId"] = o.ArrivalStationId
	}
	toSerialize["arrivalTime"] = o.ArrivalTime
	toSerialize["vehicleTypes"] = o.VehicleTypes
	if !IsNil(o.TransfersCount) {
		toSerialize["transfersCount"] = o.TransfersCount
	}
	toSerialize["freeSeatsCount"] = o.FreeSeatsCount
	toSerialize["priceFrom"] = o.PriceFrom
	if !IsNil(o.PriceTo) {
		toSerialize["priceTo"] = o.PriceTo
	}
	toSerialize["creditPriceFrom"] = o.CreditPriceFrom
	if !IsNil(o.CreditPriceTo) {
		toSerialize["creditPriceTo"] = o.CreditPriceTo
	}
	toSerialize["pricesCount"] = o.PricesCount
	toSerialize["actionPrice"] = o.ActionPrice
	toSerialize["surcharge"] = o.Surcharge
	toSerialize["notices"] = o.Notices
	toSerialize["support"] = o.Support
	if !IsNil(o.NationalTrip) {
		toSerialize["nationalTrip"] = o.NationalTrip
	}
	toSerialize["bookable"] = o.Bookable
	if !IsNil(o.Delay) {
		toSerialize["delay"] = o.Delay
	}
	if !IsNil(o.TravelTime) {
		toSerialize["travelTime"] = o.TravelTime
	}
	if !IsNil(o.VehicleStandardKey) {
		toSerialize["vehicleStandardKey"] = o.VehicleStandardKey
	}
	return toSerialize, nil
}

func (o *SimpleRoute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"departureTime",
		"arrivalTime",
		"vehicleTypes",
		"freeSeatsCount",
		"priceFrom",
		"creditPriceFrom",
		"pricesCount",
		"actionPrice",
		"surcharge",
		"notices",
		"support",
		"bookable",
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

	varSimpleRoute := _SimpleRoute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimpleRoute)

	if err != nil {
		return err
	}

	*o = SimpleRoute(varSimpleRoute)

	return err
}

type NullableSimpleRoute struct {
	value *SimpleRoute
	isSet bool
}

func (v NullableSimpleRoute) Get() *SimpleRoute {
	return v.value
}

func (v *NullableSimpleRoute) Set(val *SimpleRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleRoute(val *SimpleRoute) *NullableSimpleRoute {
	return &NullableSimpleRoute{value: val, isSet: true}
}

func (v NullableSimpleRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
