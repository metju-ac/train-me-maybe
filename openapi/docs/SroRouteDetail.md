# SroRouteDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | {section0.id},{section1.id}, ... | 
**PriceClasses** | [**[]SroPriceClass**](SroPriceClass.md) |  | 

## Methods

### NewSroRouteDetail

`func NewSroRouteDetail(id string, priceClasses []SroPriceClass, ) *SroRouteDetail`

NewSroRouteDetail instantiates a new SroRouteDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroRouteDetailWithDefaults

`func NewSroRouteDetailWithDefaults() *SroRouteDetail`

NewSroRouteDetailWithDefaults instantiates a new SroRouteDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SroRouteDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SroRouteDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SroRouteDetail) SetId(v string)`

SetId sets Id field to given value.


### GetPriceClasses

`func (o *SroRouteDetail) GetPriceClasses() []SroPriceClass`

GetPriceClasses returns the PriceClasses field if non-nil, zero value otherwise.

### GetPriceClassesOk

`func (o *SroRouteDetail) GetPriceClassesOk() (*[]SroPriceClass, bool)`

GetPriceClassesOk returns a tuple with the PriceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceClasses

`func (o *SroRouteDetail) SetPriceClasses(v []SroPriceClass)`

SetPriceClasses sets PriceClasses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


