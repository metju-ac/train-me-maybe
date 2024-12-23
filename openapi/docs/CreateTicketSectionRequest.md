# CreateTicketSectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Section** | [**SimpleSection**](SimpleSection.md) |  | 
**SelectedSeats** | Pointer to [**[]SelectedSeat100**](SelectedSeat100.md) |  | [optional] 

## Methods

### NewCreateTicketSectionRequest

`func NewCreateTicketSectionRequest(section SimpleSection, ) *CreateTicketSectionRequest`

NewCreateTicketSectionRequest instantiates a new CreateTicketSectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketSectionRequestWithDefaults

`func NewCreateTicketSectionRequestWithDefaults() *CreateTicketSectionRequest`

NewCreateTicketSectionRequestWithDefaults instantiates a new CreateTicketSectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSection

`func (o *CreateTicketSectionRequest) GetSection() SimpleSection`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *CreateTicketSectionRequest) GetSectionOk() (*SimpleSection, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *CreateTicketSectionRequest) SetSection(v SimpleSection)`

SetSection sets Section field to given value.


### GetSelectedSeats

`func (o *CreateTicketSectionRequest) GetSelectedSeats() []SelectedSeat100`

GetSelectedSeats returns the SelectedSeats field if non-nil, zero value otherwise.

### GetSelectedSeatsOk

`func (o *CreateTicketSectionRequest) GetSelectedSeatsOk() (*[]SelectedSeat100, bool)`

GetSelectedSeatsOk returns a tuple with the SelectedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeats

`func (o *CreateTicketSectionRequest) SetSelectedSeats(v []SelectedSeat100)`

SetSelectedSeats sets SelectedSeats field to given value.

### HasSelectedSeats

`func (o *CreateTicketSectionRequest) HasSelectedSeats() bool`

HasSelectedSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


