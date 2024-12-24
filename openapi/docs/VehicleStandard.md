# VehicleStandard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Vehicle standard key that will be used in responses from other endpoints. | 
**Name** | **string** | Vehicle standard name. | 
**Description** | **string** | Vehicle standard description. | 
**ImageUrl** | **string** | URL to vehicle standard thumbnail image. | 
**SupportImageUrl** | **string** | URL to support vehicle thumbnail image. | 
**GalleryUrl** | Pointer to **string** | URL to publicly available information about vehicle standard. | [optional] 

## Methods

### NewVehicleStandard

`func NewVehicleStandard(key string, name string, description string, imageUrl string, supportImageUrl string, ) *VehicleStandard`

NewVehicleStandard instantiates a new VehicleStandard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleStandardWithDefaults

`func NewVehicleStandardWithDefaults() *VehicleStandard`

NewVehicleStandardWithDefaults instantiates a new VehicleStandard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *VehicleStandard) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VehicleStandard) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VehicleStandard) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *VehicleStandard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleStandard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleStandard) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VehicleStandard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VehicleStandard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VehicleStandard) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageUrl

`func (o *VehicleStandard) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *VehicleStandard) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *VehicleStandard) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetSupportImageUrl

`func (o *VehicleStandard) GetSupportImageUrl() string`

GetSupportImageUrl returns the SupportImageUrl field if non-nil, zero value otherwise.

### GetSupportImageUrlOk

`func (o *VehicleStandard) GetSupportImageUrlOk() (*string, bool)`

GetSupportImageUrlOk returns a tuple with the SupportImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportImageUrl

`func (o *VehicleStandard) SetSupportImageUrl(v string)`

SetSupportImageUrl sets SupportImageUrl field to given value.


### GetGalleryUrl

`func (o *VehicleStandard) GetGalleryUrl() string`

GetGalleryUrl returns the GalleryUrl field if non-nil, zero value otherwise.

### GetGalleryUrlOk

`func (o *VehicleStandard) GetGalleryUrlOk() (*string, bool)`

GetGalleryUrlOk returns a tuple with the GalleryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalleryUrl

`func (o *VehicleStandard) SetGalleryUrl(v string)`

SetGalleryUrl sets GalleryUrl field to given value.

### HasGalleryUrl

`func (o *VehicleStandard) HasGalleryUrl() bool`

HasGalleryUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


