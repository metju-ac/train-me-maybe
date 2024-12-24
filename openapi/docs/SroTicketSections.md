# SroTicketSections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Section** | [**SroTicketSection**](SroTicketSection.md) |  | 
**SelectedSeats** | [**[]SroTicketSelectedSeat**](SroTicketSelectedSeat.md) |  | 

## Methods

### NewSroTicketSections

`func NewSroTicketSections(section SroTicketSection, selectedSeats []SroTicketSelectedSeat, ) *SroTicketSections`

NewSroTicketSections instantiates a new SroTicketSections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroTicketSectionsWithDefaults

`func NewSroTicketSectionsWithDefaults() *SroTicketSections`

NewSroTicketSectionsWithDefaults instantiates a new SroTicketSections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSection

`func (o *SroTicketSections) GetSection() SroTicketSection`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *SroTicketSections) GetSectionOk() (*SroTicketSection, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *SroTicketSections) SetSection(v SroTicketSection)`

SetSection sets Section field to given value.


### GetSelectedSeats

`func (o *SroTicketSections) GetSelectedSeats() []SroTicketSelectedSeat`

GetSelectedSeats returns the SelectedSeats field if non-nil, zero value otherwise.

### GetSelectedSeatsOk

`func (o *SroTicketSections) GetSelectedSeatsOk() (*[]SroTicketSelectedSeat, bool)`

GetSelectedSeatsOk returns a tuple with the SelectedSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSeats

`func (o *SroTicketSections) SetSelectedSeats(v []SroTicketSelectedSeat)`

SetSelectedSeats sets SelectedSeats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


