# InlineResponse4001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message. | 
**ErrorFields** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInlineResponse4001

`func NewInlineResponse4001(message string, ) *InlineResponse4001`

NewInlineResponse4001 instantiates a new InlineResponse4001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse4001WithDefaults

`func NewInlineResponse4001WithDefaults() *InlineResponse4001`

NewInlineResponse4001WithDefaults instantiates a new InlineResponse4001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InlineResponse4001) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse4001) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse4001) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorFields

`func (o *InlineResponse4001) GetErrorFields() []map[string]interface{}`

GetErrorFields returns the ErrorFields field if non-nil, zero value otherwise.

### GetErrorFieldsOk

`func (o *InlineResponse4001) GetErrorFieldsOk() (*[]map[string]interface{}, bool)`

GetErrorFieldsOk returns a tuple with the ErrorFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorFields

`func (o *InlineResponse4001) SetErrorFields(v []map[string]interface{})`

SetErrorFields sets ErrorFields field to given value.

### HasErrorFields

`func (o *InlineResponse4001) HasErrorFields() bool`

HasErrorFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


