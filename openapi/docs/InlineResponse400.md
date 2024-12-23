# InlineResponse400

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message. | 
**ErrorFields** | Pointer to [**[]InlineResponse400ErrorFields**](InlineResponse400ErrorFields.md) |  | [optional] 

## Methods

### NewInlineResponse400

`func NewInlineResponse400(message string, ) *InlineResponse400`

NewInlineResponse400 instantiates a new InlineResponse400 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse400WithDefaults

`func NewInlineResponse400WithDefaults() *InlineResponse400`

NewInlineResponse400WithDefaults instantiates a new InlineResponse400 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InlineResponse400) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse400) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse400) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorFields

`func (o *InlineResponse400) GetErrorFields() []InlineResponse400ErrorFields`

GetErrorFields returns the ErrorFields field if non-nil, zero value otherwise.

### GetErrorFieldsOk

`func (o *InlineResponse400) GetErrorFieldsOk() (*[]InlineResponse400ErrorFields, bool)`

GetErrorFieldsOk returns a tuple with the ErrorFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFields

`func (o *InlineResponse400) SetErrorFields(v []InlineResponse400ErrorFields)`

SetErrorFields sets ErrorFields field to given value.

### HasErrorFields

`func (o *InlineResponse400) HasErrorFields() bool`

HasErrorFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


