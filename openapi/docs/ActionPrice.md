# ActionPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier of action. | 
**Code** | **string** | Action price code. | 
**Name** | **string** | Action price name. | 
**Url** | **string** | Publicly available URL of an action. | 
**Description** | Pointer to **string** | Action price description. | [optional] 
**ShowIcon** | **bool** | Flag that indicates wheter icon should be shown with action price name. | 

## Methods

### NewActionPrice

`func NewActionPrice(id int64, code string, name string, url string, showIcon bool, ) *ActionPrice`

NewActionPrice instantiates a new ActionPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPriceWithDefaults

`func NewActionPriceWithDefaults() *ActionPrice`

NewActionPriceWithDefaults instantiates a new ActionPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionPrice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPrice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPrice) SetId(v int64)`

SetId sets Id field to given value.


### GetCode

`func (o *ActionPrice) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ActionPrice) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ActionPrice) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *ActionPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionPrice) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *ActionPrice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ActionPrice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ActionPrice) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *ActionPrice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionPrice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionPrice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionPrice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShowIcon

`func (o *ActionPrice) GetShowIcon() bool`

GetShowIcon returns the ShowIcon field if non-nil, zero value otherwise.

### GetShowIconOk

`func (o *ActionPrice) GetShowIconOk() (*bool, bool)`

GetShowIconOk returns a tuple with the ShowIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIcon

`func (o *ActionPrice) SetShowIcon(v bool)`

SetShowIcon sets ShowIcon field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


