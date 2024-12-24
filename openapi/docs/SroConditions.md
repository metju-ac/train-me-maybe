# SroConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoolingOff** | **time.Time** | CoolingOff Time for sro ticket | 
**DepartureCancel** | **time.Time** | Departure Cancel Time for sro ticket | 
**Expiration** | **time.Time** | Expiration Time for sro ticket | 
**CancellationFee** | **float32** | Cancellation Fee price for sro ticket, addons etc. | 

## Methods

### NewSroConditions

`func NewSroConditions(coolingOff time.Time, departureCancel time.Time, expiration time.Time, cancellationFee float32, ) *SroConditions`

NewSroConditions instantiates a new SroConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroConditionsWithDefaults

`func NewSroConditionsWithDefaults() *SroConditions`

NewSroConditionsWithDefaults instantiates a new SroConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoolingOff

`func (o *SroConditions) GetCoolingOff() time.Time`

GetCoolingOff returns the CoolingOff field if non-nil, zero value otherwise.

### GetCoolingOffOk

`func (o *SroConditions) GetCoolingOffOk() (*time.Time, bool)`

GetCoolingOffOk returns a tuple with the CoolingOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolingOff

`func (o *SroConditions) SetCoolingOff(v time.Time)`

SetCoolingOff sets CoolingOff field to given value.


### GetDepartureCancel

`func (o *SroConditions) GetDepartureCancel() time.Time`

GetDepartureCancel returns the DepartureCancel field if non-nil, zero value otherwise.

### GetDepartureCancelOk

`func (o *SroConditions) GetDepartureCancelOk() (*time.Time, bool)`

GetDepartureCancelOk returns a tuple with the DepartureCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureCancel

`func (o *SroConditions) SetDepartureCancel(v time.Time)`

SetDepartureCancel sets DepartureCancel field to given value.


### GetExpiration

`func (o *SroConditions) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *SroConditions) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *SroConditions) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### GetCancellationFee

`func (o *SroConditions) GetCancellationFee() float32`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *SroConditions) GetCancellationFeeOk() (*float32, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *SroConditions) SetCancellationFee(v float32)`

SetCancellationFee sets CancellationFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


