# SroTicketBookingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCode** | **string** | Ticket number that customer need during checkin on board. | 
**SroTickets** | [**[]SroTicket**](SroTicket.md) | Information about the newly created ticket. | 

## Methods

### NewSroTicketBookingResponse

`func NewSroTicketBookingResponse(accountCode string, sroTickets []SroTicket, ) *SroTicketBookingResponse`

NewSroTicketBookingResponse instantiates a new SroTicketBookingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSroTicketBookingResponseWithDefaults

`func NewSroTicketBookingResponseWithDefaults() *SroTicketBookingResponse`

NewSroTicketBookingResponseWithDefaults instantiates a new SroTicketBookingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCode

`func (o *SroTicketBookingResponse) GetAccountCode() string`

GetAccountCode returns the AccountCode field if non-nil, zero value otherwise.

### GetAccountCodeOk

`func (o *SroTicketBookingResponse) GetAccountCodeOk() (*string, bool)`

GetAccountCodeOk returns a tuple with the AccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCode

`func (o *SroTicketBookingResponse) SetAccountCode(v string)`

SetAccountCode sets AccountCode field to given value.


### GetSroTickets

`func (o *SroTicketBookingResponse) GetSroTickets() []SroTicket`

GetSroTickets returns the SroTickets field if non-nil, zero value otherwise.

### GetSroTicketsOk

`func (o *SroTicketBookingResponse) GetSroTicketsOk() (*[]SroTicket, bool)`

GetSroTicketsOk returns a tuple with the SroTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSroTickets

`func (o *SroTicketBookingResponse) SetSroTickets(v []SroTicket)`

SetSroTickets sets SroTickets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


