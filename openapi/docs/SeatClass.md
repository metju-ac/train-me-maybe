# SeatClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Seat class key. Used in responses of other endpoints. | 
**VehicleClass** | Pointer to **string** | Vehicle class. | [optional] 
**Title** | **string** | Seat class title. | 
**Description** | **string** | Full seat class description. | 
**ImageUrl** | Pointer to **string** | URL of a seat class thumbnail image. | [optional] 
**GalleryUrl** | Pointer to **string** | URL of a publicly accessible complete information about a seat class. | [optional] 

## Methods

### NewSeatClass

`func NewSeatClass(key string, title string, description string, ) *SeatClass`

NewSeatClass instantiates a new SeatClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeatClassWithDefaults

`func NewSeatClassWithDefaults() *SeatClass`

NewSeatClassWithDefaults instantiates a new SeatClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SeatClass) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SeatClass) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SeatClass) SetKey(v string)`

SetKey sets Key field to given value.


### GetVehicleClass

`func (o *SeatClass) GetVehicleClass() string`

GetVehicleClass returns the VehicleClass field if non-nil, zero value otherwise.

### GetVehicleClassOk

`func (o *SeatClass) GetVehicleClassOk() (*string, bool)`

GetVehicleClassOk returns a tuple with the VehicleClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleClass

`func (o *SeatClass) SetVehicleClass(v string)`

SetVehicleClass sets VehicleClass field to given value.

### HasVehicleClass

`func (o *SeatClass) HasVehicleClass() bool`

HasVehicleClass returns a boolean if a field has been set.

### GetTitle

`func (o *SeatClass) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeatClass) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeatClass) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SeatClass) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeatClass) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeatClass) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageUrl

`func (o *SeatClass) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *SeatClass) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *SeatClass) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *SeatClass) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetGalleryUrl

`func (o *SeatClass) GetGalleryUrl() string`

GetGalleryUrl returns the GalleryUrl field if non-nil, zero value otherwise.

### GetGalleryUrlOk

`func (o *SeatClass) GetGalleryUrlOk() (*string, bool)`

GetGalleryUrlOk returns a tuple with the GalleryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalleryUrl

`func (o *SeatClass) SetGalleryUrl(v string)`

SetGalleryUrl sets GalleryUrl field to given value.

### HasGalleryUrl

`func (o *SeatClass) HasGalleryUrl() bool`

HasGalleryUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


