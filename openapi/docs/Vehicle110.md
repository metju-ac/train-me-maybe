# Vehicle110

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | **int64** |  | 
**Code** | Pointer to **string** | Vehicle code tag (BUS &#x3D;&gt; SPZ, VAGON &#x3D;&gt; code) | [optional] 
**LayoutURL** | Pointer to **string** |  | [optional] 
**HorizontalLayoutURL** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**VehicleType**](VehicleType.md) |  | [optional] 
**VehicleStandardKey** | **string** | Vehicle standard code tag | 
**Services** | Pointer to **[]string** | Supported services icons (wifi, etc.) | [optional] 
**VehicleNumber** | **int32** |  | 
**SeatClasses** | **[]string** | Available classes in this vehicle | 
**Notifications** | Pointer to **[]string** | Additional informations relating to whole vehicle. These informations are visible, but wont requiring confirmation. | [optional] 
**FreeSeats** | [**[]Seat**](Seat.md) |  | 

## Methods

### NewVehicle110

`func NewVehicle110(vehicleId int64, vehicleStandardKey string, vehicleNumber int32, seatClasses []string, freeSeats []Seat, ) *Vehicle110`

NewVehicle110 instantiates a new Vehicle110 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicle110WithDefaults

`func NewVehicle110WithDefaults() *Vehicle110`

NewVehicle110WithDefaults instantiates a new Vehicle110 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleId

`func (o *Vehicle110) GetVehicleId() int64`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *Vehicle110) GetVehicleIdOk() (*int64, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *Vehicle110) SetVehicleId(v int64)`

SetVehicleId sets VehicleId field to given value.


### GetCode

`func (o *Vehicle110) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Vehicle110) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Vehicle110) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Vehicle110) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLayoutURL

`func (o *Vehicle110) GetLayoutURL() string`

GetLayoutURL returns the LayoutURL field if non-nil, zero value otherwise.

### GetLayoutURLOk

`func (o *Vehicle110) GetLayoutURLOk() (*string, bool)`

GetLayoutURLOk returns a tuple with the LayoutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutURL

`func (o *Vehicle110) SetLayoutURL(v string)`

SetLayoutURL sets LayoutURL field to given value.

### HasLayoutURL

`func (o *Vehicle110) HasLayoutURL() bool`

HasLayoutURL returns a boolean if a field has been set.

### GetHorizontalLayoutURL

`func (o *Vehicle110) GetHorizontalLayoutURL() string`

GetHorizontalLayoutURL returns the HorizontalLayoutURL field if non-nil, zero value otherwise.

### GetHorizontalLayoutURLOk

`func (o *Vehicle110) GetHorizontalLayoutURLOk() (*string, bool)`

GetHorizontalLayoutURLOk returns a tuple with the HorizontalLayoutURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalLayoutURL

`func (o *Vehicle110) SetHorizontalLayoutURL(v string)`

SetHorizontalLayoutURL sets HorizontalLayoutURL field to given value.

### HasHorizontalLayoutURL

`func (o *Vehicle110) HasHorizontalLayoutURL() bool`

HasHorizontalLayoutURL returns a boolean if a field has been set.

### GetType

`func (o *Vehicle110) GetType() VehicleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Vehicle110) GetTypeOk() (*VehicleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Vehicle110) SetType(v VehicleType)`

SetType sets Type field to given value.

### HasType

`func (o *Vehicle110) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVehicleStandardKey

`func (o *Vehicle110) GetVehicleStandardKey() string`

GetVehicleStandardKey returns the VehicleStandardKey field if non-nil, zero value otherwise.

### GetVehicleStandardKeyOk

`func (o *Vehicle110) GetVehicleStandardKeyOk() (*string, bool)`

GetVehicleStandardKeyOk returns a tuple with the VehicleStandardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleStandardKey

`func (o *Vehicle110) SetVehicleStandardKey(v string)`

SetVehicleStandardKey sets VehicleStandardKey field to given value.


### GetServices

`func (o *Vehicle110) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Vehicle110) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Vehicle110) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Vehicle110) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetVehicleNumber

`func (o *Vehicle110) GetVehicleNumber() int32`

GetVehicleNumber returns the VehicleNumber field if non-nil, zero value otherwise.

### GetVehicleNumberOk

`func (o *Vehicle110) GetVehicleNumberOk() (*int32, bool)`

GetVehicleNumberOk returns a tuple with the VehicleNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleNumber

`func (o *Vehicle110) SetVehicleNumber(v int32)`

SetVehicleNumber sets VehicleNumber field to given value.


### GetSeatClasses

`func (o *Vehicle110) GetSeatClasses() []string`

GetSeatClasses returns the SeatClasses field if non-nil, zero value otherwise.

### GetSeatClassesOk

`func (o *Vehicle110) GetSeatClassesOk() (*[]string, bool)`

GetSeatClassesOk returns a tuple with the SeatClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClasses

`func (o *Vehicle110) SetSeatClasses(v []string)`

SetSeatClasses sets SeatClasses field to given value.


### GetNotifications

`func (o *Vehicle110) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Vehicle110) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Vehicle110) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Vehicle110) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetFreeSeats

`func (o *Vehicle110) GetFreeSeats() []Seat`

GetFreeSeats returns the FreeSeats field if non-nil, zero value otherwise.

### GetFreeSeatsOk

`func (o *Vehicle110) GetFreeSeatsOk() (*[]Seat, bool)`

GetFreeSeatsOk returns a tuple with the FreeSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSeats

`func (o *Vehicle110) SetFreeSeats(v []Seat)`

SetFreeSeats sets FreeSeats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


