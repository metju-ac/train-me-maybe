# PassengersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passengers** | [**[]Passenger**](Passenger.md) |  | 
**FirstPassengerData** | [**[]PersonalDataType**](PersonalDataType.md) | Required fields for 1st passenger (returns only if tickets is editable) | 
**OtherPassengersData** | [**[]PersonalDataType**](PersonalDataType.md) | Required fields for all others passengers (returns only if tickets is editable) | 
**ChangeCharge** | **float32** | Administrative charge for user information change (in ticket currency) | 

## Methods

### NewPassengersInfo

`func NewPassengersInfo(passengers []Passenger, firstPassengerData []PersonalDataType, otherPassengersData []PersonalDataType, changeCharge float32, ) *PassengersInfo`

NewPassengersInfo instantiates a new PassengersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassengersInfoWithDefaults

`func NewPassengersInfoWithDefaults() *PassengersInfo`

NewPassengersInfoWithDefaults instantiates a new PassengersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassengers

`func (o *PassengersInfo) GetPassengers() []Passenger`

GetPassengers returns the Passengers field if non-nil, zero value otherwise.

### GetPassengersOk

`func (o *PassengersInfo) GetPassengersOk() (*[]Passenger, bool)`

GetPassengersOk returns a tuple with the Passengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengers

`func (o *PassengersInfo) SetPassengers(v []Passenger)`

SetPassengers sets Passengers field to given value.


### GetFirstPassengerData

`func (o *PassengersInfo) GetFirstPassengerData() []PersonalDataType`

GetFirstPassengerData returns the FirstPassengerData field if non-nil, zero value otherwise.

### GetFirstPassengerDataOk

`func (o *PassengersInfo) GetFirstPassengerDataOk() (*[]PersonalDataType, bool)`

GetFirstPassengerDataOk returns a tuple with the FirstPassengerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPassengerData

`func (o *PassengersInfo) SetFirstPassengerData(v []PersonalDataType)`

SetFirstPassengerData sets FirstPassengerData field to given value.


### GetOtherPassengersData

`func (o *PassengersInfo) GetOtherPassengersData() []PersonalDataType`

GetOtherPassengersData returns the OtherPassengersData field if non-nil, zero value otherwise.

### GetOtherPassengersDataOk

`func (o *PassengersInfo) GetOtherPassengersDataOk() (*[]PersonalDataType, bool)`

GetOtherPassengersDataOk returns a tuple with the OtherPassengersData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherPassengersData

`func (o *PassengersInfo) SetOtherPassengersData(v []PersonalDataType)`

SetOtherPassengersData sets OtherPassengersData field to given value.


### GetChangeCharge

`func (o *PassengersInfo) GetChangeCharge() float32`

GetChangeCharge returns the ChangeCharge field if non-nil, zero value otherwise.

### GetChangeChargeOk

`func (o *PassengersInfo) GetChangeChargeOk() (*float32, bool)`

GetChangeChargeOk returns a tuple with the ChangeCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeCharge

`func (o *PassengersInfo) SetChangeCharge(v float32)`

SetChangeCharge sets ChangeCharge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


