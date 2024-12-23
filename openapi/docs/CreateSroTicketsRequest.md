# CreateSroTicketsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreeWithTerms** | **bool** |  | 
**TicketRequests** | [**[]SroTicketRequest**](SroTicketRequest.md) |  | 
**ExternalId** | **string** |  | 

## Methods

### NewCreateSroTicketsRequest

`func NewCreateSroTicketsRequest(agreeWithTerms bool, ticketRequests []SroTicketRequest, externalId string, ) *CreateSroTicketsRequest`

NewCreateSroTicketsRequest instantiates a new CreateSroTicketsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSroTicketsRequestWithDefaults

`func NewCreateSroTicketsRequestWithDefaults() *CreateSroTicketsRequest`

NewCreateSroTicketsRequestWithDefaults instantiates a new CreateSroTicketsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreeWithTerms

`func (o *CreateSroTicketsRequest) GetAgreeWithTerms() bool`

GetAgreeWithTerms returns the AgreeWithTerms field if non-nil, zero value otherwise.

### GetAgreeWithTermsOk

`func (o *CreateSroTicketsRequest) GetAgreeWithTermsOk() (*bool, bool)`

GetAgreeWithTermsOk returns a tuple with the AgreeWithTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeWithTerms

`func (o *CreateSroTicketsRequest) SetAgreeWithTerms(v bool)`

SetAgreeWithTerms sets AgreeWithTerms field to given value.


### GetTicketRequests

`func (o *CreateSroTicketsRequest) GetTicketRequests() []SroTicketRequest`

GetTicketRequests returns the TicketRequests field if non-nil, zero value otherwise.

### GetTicketRequestsOk

`func (o *CreateSroTicketsRequest) GetTicketRequestsOk() (*[]SroTicketRequest, bool)`

GetTicketRequestsOk returns a tuple with the TicketRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketRequests

`func (o *CreateSroTicketsRequest) SetTicketRequests(v []SroTicketRequest)`

SetTicketRequests sets TicketRequests field to given value.


### GetExternalId

`func (o *CreateSroTicketsRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateSroTicketsRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateSroTicketsRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


