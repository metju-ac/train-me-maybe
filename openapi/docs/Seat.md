# Seat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Number of seat in vehicle | [optional] 
**SeatClass** | Pointer to **string** |  | [optional] 
**SeatConstraint** | Pointer to **string** | Notification which needs to be confirmed by customer before continue in reservation | [optional] 
**SeatNotes** | Pointer to **[]string** | Supplementary informations which are shown but isnt required to be confirmed | [optional] 

## Methods

### NewSeat

`func NewSeat() *Seat`

NewSeat instantiates a new Seat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeatWithDefaults

`func NewSeatWithDefaults() *Seat`

NewSeatWithDefaults instantiates a new Seat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Seat) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Seat) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Seat) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Seat) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetSeatClass

`func (o *Seat) GetSeatClass() string`

GetSeatClass returns the SeatClass field if non-nil, zero value otherwise.

### GetSeatClassOk

`func (o *Seat) GetSeatClassOk() (*string, bool)`

GetSeatClassOk returns a tuple with the SeatClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClass

`func (o *Seat) SetSeatClass(v string)`

SetSeatClass sets SeatClass field to given value.

### HasSeatClass

`func (o *Seat) HasSeatClass() bool`

HasSeatClass returns a boolean if a field has been set.

### GetSeatConstraint

`func (o *Seat) GetSeatConstraint() string`

GetSeatConstraint returns the SeatConstraint field if non-nil, zero value otherwise.

### GetSeatConstraintOk

`func (o *Seat) GetSeatConstraintOk() (*string, bool)`

GetSeatConstraintOk returns a tuple with the SeatConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatConstraint

`func (o *Seat) SetSeatConstraint(v string)`

SetSeatConstraint sets SeatConstraint field to given value.

### HasSeatConstraint

`func (o *Seat) HasSeatConstraint() bool`

HasSeatConstraint returns a boolean if a field has been set.

### GetSeatNotes

`func (o *Seat) GetSeatNotes() []string`

GetSeatNotes returns the SeatNotes field if non-nil, zero value otherwise.

### GetSeatNotesOk

`func (o *Seat) GetSeatNotesOk() (*[]string, bool)`

GetSeatNotesOk returns a tuple with the SeatNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatNotes

`func (o *Seat) SetSeatNotes(v []string)`

SetSeatNotes sets SeatNotes field to given value.

### HasSeatNotes

`func (o *Seat) HasSeatNotes() bool`

HasSeatNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


