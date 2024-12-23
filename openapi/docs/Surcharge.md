# Surcharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **float32** | Surcharge amount (is not included in basic route price or even in &#x60;priceClass&#x60;). | 
**Notification** | **string** | Notification which needs to be confirmed by customer. | 

## Methods

### NewSurcharge

`func NewSurcharge(price float32, notification string, ) *Surcharge`

NewSurcharge instantiates a new Surcharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurchargeWithDefaults

`func NewSurchargeWithDefaults() *Surcharge`

NewSurchargeWithDefaults instantiates a new Surcharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *Surcharge) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Surcharge) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Surcharge) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetNotification

`func (o *Surcharge) GetNotification() string`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *Surcharge) GetNotificationOk() (*string, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *Surcharge) SetNotification(v string)`

SetNotification sets Notification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


