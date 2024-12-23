# PassengersDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstPassengerData** | [**[]PersonalDataType**](PersonalDataType.md) | Required fields for 1st passenger | 
**OtherPassengersData** | [**[]PersonalDataType**](PersonalDataType.md) | Required fields for all other passengers | 

## Methods

### NewPassengersDataResponse

`func NewPassengersDataResponse(firstPassengerData []PersonalDataType, otherPassengersData []PersonalDataType, ) *PassengersDataResponse`

NewPassengersDataResponse instantiates a new PassengersDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassengersDataResponseWithDefaults

`func NewPassengersDataResponseWithDefaults() *PassengersDataResponse`

NewPassengersDataResponseWithDefaults instantiates a new PassengersDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstPassengerData

`func (o *PassengersDataResponse) GetFirstPassengerData() []PersonalDataType`

GetFirstPassengerData returns the FirstPassengerData field if non-nil, zero value otherwise.

### GetFirstPassengerDataOk

`func (o *PassengersDataResponse) GetFirstPassengerDataOk() (*[]PersonalDataType, bool)`

GetFirstPassengerDataOk returns a tuple with the FirstPassengerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPassengerData

`func (o *PassengersDataResponse) SetFirstPassengerData(v []PersonalDataType)`

SetFirstPassengerData sets FirstPassengerData field to given value.


### GetOtherPassengersData

`func (o *PassengersDataResponse) GetOtherPassengersData() []PersonalDataType`

GetOtherPassengersData returns the OtherPassengersData field if non-nil, zero value otherwise.

### GetOtherPassengersDataOk

`func (o *PassengersDataResponse) GetOtherPassengersDataOk() (*[]PersonalDataType, bool)`

GetOtherPassengersDataOk returns a tuple with the OtherPassengersData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherPassengersData

`func (o *PassengersDataResponse) SetOtherPassengersData(v []PersonalDataType)`

SetOtherPassengersData sets OtherPassengersData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


