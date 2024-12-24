# Vehicle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Code** | Pointer to **string** | Vehicle code tag (BUS &#x3D;&gt; SPZ, VAGON &#x3D;&gt; code) | [optional] 
**Type** | Pointer to [**VehicleType**](VehicleType.md) |  | [optional] 
**Number** | **int32** |  | 
**SeatClasses** | [**[]VehicleSeatClass**](VehicleSeatClass.md) | Available services in this vehicle | 
**Standard** | **string** | Vehicle standard code tag | 
**Notifications** | Pointer to **[]string** | Additional informations relating to whole vehicle. These informations are visible, but wont requiring confirmation. | [optional] 
**CateringEnabled** | Pointer to **bool** | Indicates if catering is enabled in current vehicle (i.e. it is possible for a customer to place catering order). | [optional] 
**Decks** | [**[]VehicleDeck**](VehicleDeck.md) |  | 

## Methods

### NewVehicle

`func NewVehicle(id int64, number int32, seatClasses []VehicleSeatClass, standard string, decks []VehicleDeck, ) *Vehicle`

NewVehicle instantiates a new Vehicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleWithDefaults

`func NewVehicleWithDefaults() *Vehicle`

NewVehicleWithDefaults instantiates a new Vehicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vehicle) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vehicle) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vehicle) SetId(v int64)`

SetId sets Id field to given value.


### GetCode

`func (o *Vehicle) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Vehicle) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Vehicle) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Vehicle) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *Vehicle) GetType() VehicleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Vehicle) GetTypeOk() (*VehicleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Vehicle) SetType(v VehicleType)`

SetType sets Type field to given value.

### HasType

`func (o *Vehicle) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *Vehicle) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Vehicle) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Vehicle) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetSeatClasses

`func (o *Vehicle) GetSeatClasses() []VehicleSeatClass`

GetSeatClasses returns the SeatClasses field if non-nil, zero value otherwise.

### GetSeatClassesOk

`func (o *Vehicle) GetSeatClassesOk() (*[]VehicleSeatClass, bool)`

GetSeatClassesOk returns a tuple with the SeatClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClasses

`func (o *Vehicle) SetSeatClasses(v []VehicleSeatClass)`

SetSeatClasses sets SeatClasses field to given value.


### GetStandard

`func (o *Vehicle) GetStandard() string`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *Vehicle) GetStandardOk() (*string, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *Vehicle) SetStandard(v string)`

SetStandard sets Standard field to given value.


### GetNotifications

`func (o *Vehicle) GetNotifications() []string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Vehicle) GetNotificationsOk() (*[]string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Vehicle) SetNotifications(v []string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Vehicle) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetCateringEnabled

`func (o *Vehicle) GetCateringEnabled() bool`

GetCateringEnabled returns the CateringEnabled field if non-nil, zero value otherwise.

### GetCateringEnabledOk

`func (o *Vehicle) GetCateringEnabledOk() (*bool, bool)`

GetCateringEnabledOk returns a tuple with the CateringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCateringEnabled

`func (o *Vehicle) SetCateringEnabled(v bool)`

SetCateringEnabled sets CateringEnabled field to given value.

### HasCateringEnabled

`func (o *Vehicle) HasCateringEnabled() bool`

HasCateringEnabled returns a boolean if a field has been set.

### GetDecks

`func (o *Vehicle) GetDecks() []VehicleDeck`

GetDecks returns the Decks field if non-nil, zero value otherwise.

### GetDecksOk

`func (o *Vehicle) GetDecksOk() (*[]VehicleDeck, bool)`

GetDecksOk returns a tuple with the Decks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecks

`func (o *Vehicle) SetDecks(v []VehicleDeck)`

SetDecks sets Decks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


