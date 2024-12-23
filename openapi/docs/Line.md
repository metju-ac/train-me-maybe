# Line

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **int64** | Unique section identifier. | 
**Code** | Pointer to **string** | Line&#39;s connection code. | [optional] 
**From** | **string** | Line&#39;s departure city. | 
**To** | **string** | Line&#39;s arrival city. | 
**LineGroupCode** | Pointer to **string** | Line group code. | [optional] 
**LineNumber** | Pointer to **int32** | Line number. | [optional] 

## Methods

### NewLine

`func NewLine(connectionId int64, from string, to string, ) *Line`

NewLine instantiates a new Line object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineWithDefaults

`func NewLineWithDefaults() *Line`

NewLineWithDefaults instantiates a new Line object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *Line) GetConnectionId() int64`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Line) GetConnectionIdOk() (*int64, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Line) SetConnectionId(v int64)`

SetConnectionId sets ConnectionId field to given value.


### GetCode

`func (o *Line) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Line) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Line) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Line) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFrom

`func (o *Line) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Line) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Line) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *Line) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Line) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Line) SetTo(v string)`

SetTo sets To field to given value.


### GetLineGroupCode

`func (o *Line) GetLineGroupCode() string`

GetLineGroupCode returns the LineGroupCode field if non-nil, zero value otherwise.

### GetLineGroupCodeOk

`func (o *Line) GetLineGroupCodeOk() (*string, bool)`

GetLineGroupCodeOk returns a tuple with the LineGroupCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroupCode

`func (o *Line) SetLineGroupCode(v string)`

SetLineGroupCode sets LineGroupCode field to given value.

### HasLineGroupCode

`func (o *Line) HasLineGroupCode() bool`

HasLineGroupCode returns a boolean if a field has been set.

### GetLineNumber

`func (o *Line) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *Line) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *Line) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *Line) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


