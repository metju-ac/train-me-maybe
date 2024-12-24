# SimpleSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SectionId** | **int64** |  | 
**FromStationId** | **int64** | Valid station ID (city ID unsupported at this level) | 
**ToStationId** | **int64** | Valid station ID (city ID unsupported at this level) | 

## Methods

### NewSimpleSection

`func NewSimpleSection(sectionId int64, fromStationId int64, toStationId int64, ) *SimpleSection`

NewSimpleSection instantiates a new SimpleSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleSectionWithDefaults

`func NewSimpleSectionWithDefaults() *SimpleSection`

NewSimpleSectionWithDefaults instantiates a new SimpleSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSectionId

`func (o *SimpleSection) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *SimpleSection) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *SimpleSection) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.


### GetFromStationId

`func (o *SimpleSection) GetFromStationId() int64`

GetFromStationId returns the FromStationId field if non-nil, zero value otherwise.

### GetFromStationIdOk

`func (o *SimpleSection) GetFromStationIdOk() (*int64, bool)`

GetFromStationIdOk returns a tuple with the FromStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStationId

`func (o *SimpleSection) SetFromStationId(v int64)`

SetFromStationId sets FromStationId field to given value.


### GetToStationId

`func (o *SimpleSection) GetToStationId() int64`

GetToStationId returns the ToStationId field if non-nil, zero value otherwise.

### GetToStationIdOk

`func (o *SimpleSection) GetToStationIdOk() (*int64, bool)`

GetToStationIdOk returns a tuple with the ToStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToStationId

`func (o *SimpleSection) SetToStationId(v int64)`

SetToStationId sets ToStationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


