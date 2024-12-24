# CreateTicketResponseUnregistered

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | Newly created user account token which is linked to ticket(s) | 
**Tickets** | [**[]Ticket**](Ticket.md) |  | 

## Methods

### NewCreateTicketResponseUnregistered

`func NewCreateTicketResponseUnregistered(token string, tickets []Ticket, ) *CreateTicketResponseUnregistered`

NewCreateTicketResponseUnregistered instantiates a new CreateTicketResponseUnregistered object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketResponseUnregisteredWithDefaults

`func NewCreateTicketResponseUnregisteredWithDefaults() *CreateTicketResponseUnregistered`

NewCreateTicketResponseUnregisteredWithDefaults instantiates a new CreateTicketResponseUnregistered object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CreateTicketResponseUnregistered) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateTicketResponseUnregistered) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateTicketResponseUnregistered) SetToken(v string)`

SetToken sets Token field to given value.


### GetTickets

`func (o *CreateTicketResponseUnregistered) GetTickets() []Ticket`

GetTickets returns the Tickets field if non-nil, zero value otherwise.

### GetTicketsOk

`func (o *CreateTicketResponseUnregistered) GetTicketsOk() (*[]Ticket, bool)`

GetTicketsOk returns a tuple with the Tickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickets

`func (o *CreateTicketResponseUnregistered) SetTickets(v []Ticket)`

SetTickets sets Tickets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


