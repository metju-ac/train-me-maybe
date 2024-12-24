# ForbiddenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message. | 
**ErrorFields** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewForbiddenResponse

`func NewForbiddenResponse(message string, ) *ForbiddenResponse`

NewForbiddenResponse instantiates a new ForbiddenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenResponseWithDefaults

`func NewForbiddenResponseWithDefaults() *ForbiddenResponse`

NewForbiddenResponseWithDefaults instantiates a new ForbiddenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ForbiddenResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ForbiddenResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ForbiddenResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorFields

`func (o *ForbiddenResponse) GetErrorFields() []map[string]interface{}`

GetErrorFields returns the ErrorFields field if non-nil, zero value otherwise.

### GetErrorFieldsOk

`func (o *ForbiddenResponse) GetErrorFieldsOk() (*[]map[string]interface{}, bool)`

GetErrorFieldsOk returns a tuple with the ErrorFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFields

`func (o *ForbiddenResponse) SetErrorFields(v []map[string]interface{})`

SetErrorFields sets ErrorFields field to given value.

### HasErrorFields

`func (o *ForbiddenResponse) HasErrorFields() bool`

HasErrorFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


