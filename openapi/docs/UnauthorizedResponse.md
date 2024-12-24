# UnauthorizedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message. | 
**ErrorFields** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewUnauthorizedResponse

`func NewUnauthorizedResponse(message string, ) *UnauthorizedResponse`

NewUnauthorizedResponse instantiates a new UnauthorizedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedResponseWithDefaults

`func NewUnauthorizedResponseWithDefaults() *UnauthorizedResponse`

NewUnauthorizedResponseWithDefaults instantiates a new UnauthorizedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UnauthorizedResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorFields

`func (o *UnauthorizedResponse) GetErrorFields() []map[string]interface{}`

GetErrorFields returns the ErrorFields field if non-nil, zero value otherwise.

### GetErrorFieldsOk

`func (o *UnauthorizedResponse) GetErrorFieldsOk() (*[]map[string]interface{}, bool)`

GetErrorFieldsOk returns a tuple with the ErrorFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFields

`func (o *UnauthorizedResponse) SetErrorFields(v []map[string]interface{})`

SetErrorFields sets ErrorFields field to given value.

### HasErrorFields

`func (o *UnauthorizedResponse) HasErrorFields() bool`

HasErrorFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


