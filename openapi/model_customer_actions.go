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

// checks if the CustomerActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerActions{}

// CustomerActions Defines which actions can be executed with current ticket state
type CustomerActions struct {
	ShowDetail bool `json:"showDetail"`
	Pay bool `json:"pay"`
	PayRemaining bool `json:"payRemaining"`
	Evaluate bool `json:"evaluate"`
	Cancel bool `json:"cancel"`
	Storno bool `json:"storno"`
	Rebook bool `json:"rebook"`
	EditPassengers bool `json:"editPassengers"`
	AdditionalServices bool `json:"additionalServices"`
	SentToMail bool `json:"sentToMail"`
	PrintTicket bool `json:"printTicket"`
	PrintInvoice bool `json:"printInvoice"`
}

type _CustomerActions CustomerActions

// NewCustomerActions instantiates a new CustomerActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerActions(showDetail bool, pay bool, payRemaining bool, evaluate bool, cancel bool, storno bool, rebook bool, editPassengers bool, additionalServices bool, sentToMail bool, printTicket bool, printInvoice bool) *CustomerActions {
	this := CustomerActions{}
	this.ShowDetail = showDetail
	this.Pay = pay
	this.PayRemaining = payRemaining
	this.Evaluate = evaluate
	this.Cancel = cancel
	this.Storno = storno
	this.Rebook = rebook
	this.EditPassengers = editPassengers
	this.AdditionalServices = additionalServices
	this.SentToMail = sentToMail
	this.PrintTicket = printTicket
	this.PrintInvoice = printInvoice
	return &this
}

// NewCustomerActionsWithDefaults instantiates a new CustomerActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerActionsWithDefaults() *CustomerActions {
	this := CustomerActions{}
	var showDetail bool = false
	this.ShowDetail = showDetail
	var pay bool = false
	this.Pay = pay
	var payRemaining bool = false
	this.PayRemaining = payRemaining
	var evaluate bool = false
	this.Evaluate = evaluate
	var cancel bool = false
	this.Cancel = cancel
	var storno bool = false
	this.Storno = storno
	var rebook bool = false
	this.Rebook = rebook
	var editPassengers bool = false
	this.EditPassengers = editPassengers
	var additionalServices bool = false
	this.AdditionalServices = additionalServices
	var sentToMail bool = false
	this.SentToMail = sentToMail
	var printTicket bool = false
	this.PrintTicket = printTicket
	var printInvoice bool = false
	this.PrintInvoice = printInvoice
	return &this
}

// GetShowDetail returns the ShowDetail field value
func (o *CustomerActions) GetShowDetail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowDetail
}

// GetShowDetailOk returns a tuple with the ShowDetail field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetShowDetailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowDetail, true
}

// SetShowDetail sets field value
func (o *CustomerActions) SetShowDetail(v bool) {
	o.ShowDetail = v
}

// GetPay returns the Pay field value
func (o *CustomerActions) GetPay() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pay
}

// GetPayOk returns a tuple with the Pay field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetPayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pay, true
}

// SetPay sets field value
func (o *CustomerActions) SetPay(v bool) {
	o.Pay = v
}

// GetPayRemaining returns the PayRemaining field value
func (o *CustomerActions) GetPayRemaining() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PayRemaining
}

// GetPayRemainingOk returns a tuple with the PayRemaining field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetPayRemainingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayRemaining, true
}

// SetPayRemaining sets field value
func (o *CustomerActions) SetPayRemaining(v bool) {
	o.PayRemaining = v
}

// GetEvaluate returns the Evaluate field value
func (o *CustomerActions) GetEvaluate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Evaluate
}

// GetEvaluateOk returns a tuple with the Evaluate field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetEvaluateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Evaluate, true
}

// SetEvaluate sets field value
func (o *CustomerActions) SetEvaluate(v bool) {
	o.Evaluate = v
}

// GetCancel returns the Cancel field value
func (o *CustomerActions) GetCancel() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetCancelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cancel, true
}

// SetCancel sets field value
func (o *CustomerActions) SetCancel(v bool) {
	o.Cancel = v
}

// GetStorno returns the Storno field value
func (o *CustomerActions) GetStorno() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Storno
}

// GetStornoOk returns a tuple with the Storno field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetStornoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storno, true
}

// SetStorno sets field value
func (o *CustomerActions) SetStorno(v bool) {
	o.Storno = v
}

// GetRebook returns the Rebook field value
func (o *CustomerActions) GetRebook() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rebook
}

// GetRebookOk returns a tuple with the Rebook field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetRebookOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rebook, true
}

// SetRebook sets field value
func (o *CustomerActions) SetRebook(v bool) {
	o.Rebook = v
}

// GetEditPassengers returns the EditPassengers field value
func (o *CustomerActions) GetEditPassengers() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EditPassengers
}

// GetEditPassengersOk returns a tuple with the EditPassengers field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetEditPassengersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditPassengers, true
}

// SetEditPassengers sets field value
func (o *CustomerActions) SetEditPassengers(v bool) {
	o.EditPassengers = v
}

// GetAdditionalServices returns the AdditionalServices field value
func (o *CustomerActions) GetAdditionalServices() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AdditionalServices
}

// GetAdditionalServicesOk returns a tuple with the AdditionalServices field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetAdditionalServicesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalServices, true
}

// SetAdditionalServices sets field value
func (o *CustomerActions) SetAdditionalServices(v bool) {
	o.AdditionalServices = v
}

// GetSentToMail returns the SentToMail field value
func (o *CustomerActions) GetSentToMail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SentToMail
}

// GetSentToMailOk returns a tuple with the SentToMail field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetSentToMailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SentToMail, true
}

// SetSentToMail sets field value
func (o *CustomerActions) SetSentToMail(v bool) {
	o.SentToMail = v
}

// GetPrintTicket returns the PrintTicket field value
func (o *CustomerActions) GetPrintTicket() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrintTicket
}

// GetPrintTicketOk returns a tuple with the PrintTicket field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetPrintTicketOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrintTicket, true
}

// SetPrintTicket sets field value
func (o *CustomerActions) SetPrintTicket(v bool) {
	o.PrintTicket = v
}

// GetPrintInvoice returns the PrintInvoice field value
func (o *CustomerActions) GetPrintInvoice() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PrintInvoice
}

// GetPrintInvoiceOk returns a tuple with the PrintInvoice field value
// and a boolean to check if the value has been set.
func (o *CustomerActions) GetPrintInvoiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrintInvoice, true
}

// SetPrintInvoice sets field value
func (o *CustomerActions) SetPrintInvoice(v bool) {
	o.PrintInvoice = v
}

func (o CustomerActions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["showDetail"] = o.ShowDetail
	toSerialize["pay"] = o.Pay
	toSerialize["payRemaining"] = o.PayRemaining
	toSerialize["evaluate"] = o.Evaluate
	toSerialize["cancel"] = o.Cancel
	toSerialize["storno"] = o.Storno
	toSerialize["rebook"] = o.Rebook
	toSerialize["editPassengers"] = o.EditPassengers
	toSerialize["additionalServices"] = o.AdditionalServices
	toSerialize["sentToMail"] = o.SentToMail
	toSerialize["printTicket"] = o.PrintTicket
	toSerialize["printInvoice"] = o.PrintInvoice
	return toSerialize, nil
}

func (o *CustomerActions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"showDetail",
		"pay",
		"payRemaining",
		"evaluate",
		"cancel",
		"storno",
		"rebook",
		"editPassengers",
		"additionalServices",
		"sentToMail",
		"printTicket",
		"printInvoice",
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

	varCustomerActions := _CustomerActions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerActions)

	if err != nil {
		return err
	}

	*o = CustomerActions(varCustomerActions)

	return err
}

type NullableCustomerActions struct {
	value *CustomerActions
	isSet bool
}

func (v NullableCustomerActions) Get() *CustomerActions {
	return v.value
}

func (v *NullableCustomerActions) Set(val *CustomerActions) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerActions) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerActions(val *CustomerActions) *NullableCustomerActions {
	return &NullableCustomerActions{value: val, isSet: true}
}

func (v NullableCustomerActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


