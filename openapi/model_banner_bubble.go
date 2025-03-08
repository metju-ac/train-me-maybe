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

// checks if the BannerBubble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BannerBubble{}

// BannerBubble Bubble picture with possibility of redirection to more informations
type BannerBubble struct {
	Id int64 `json:"id"`
	// Extra information for bubble.
	Text string `json:"text"`
	// Redirection to more information.
	Url *string `json:"url,omitempty"`
	// Bubble picture URL address.
	ImageUrl string `json:"imageUrl"`
}

type _BannerBubble BannerBubble

// NewBannerBubble instantiates a new BannerBubble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBannerBubble(id int64, text string, imageUrl string) *BannerBubble {
	this := BannerBubble{}
	this.Id = id
	this.Text = text
	this.ImageUrl = imageUrl
	return &this
}

// NewBannerBubbleWithDefaults instantiates a new BannerBubble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBannerBubbleWithDefaults() *BannerBubble {
	this := BannerBubble{}
	return &this
}

// GetId returns the Id field value
func (o *BannerBubble) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BannerBubble) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BannerBubble) SetId(v int64) {
	o.Id = v
}

// GetText returns the Text field value
func (o *BannerBubble) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *BannerBubble) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *BannerBubble) SetText(v string) {
	o.Text = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BannerBubble) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BannerBubble) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BannerBubble) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BannerBubble) SetUrl(v string) {
	o.Url = &v
}

// GetImageUrl returns the ImageUrl field value
func (o *BannerBubble) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *BannerBubble) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *BannerBubble) SetImageUrl(v string) {
	o.ImageUrl = v
}

func (o BannerBubble) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BannerBubble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["text"] = o.Text
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	toSerialize["imageUrl"] = o.ImageUrl
	return toSerialize, nil
}

func (o *BannerBubble) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"text",
		"imageUrl",
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

	varBannerBubble := _BannerBubble{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBannerBubble)

	if err != nil {
		return err
	}

	*o = BannerBubble(varBannerBubble)

	return err
}

type NullableBannerBubble struct {
	value *BannerBubble
	isSet bool
}

func (v NullableBannerBubble) Get() *BannerBubble {
	return v.value
}

func (v *NullableBannerBubble) Set(val *BannerBubble) {
	v.value = val
	v.isSet = true
}

func (v NullableBannerBubble) IsSet() bool {
	return v.isSet
}

func (v *NullableBannerBubble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBannerBubble(val *BannerBubble) *NullableBannerBubble {
	return &NullableBannerBubble{value: val, isSet: true}
}

func (v NullableBannerBubble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBannerBubble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
