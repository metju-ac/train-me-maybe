# CreateUnregisteredTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreeWithTerms** | **bool** | Agree with terms | [default to false]
**TicketRequests** | [**[]CreateTicketRequest**](CreateTicketRequest.md) |  | 
**AffiliateCode** | Pointer to **string** | Optional code of an affiliate partner. | [optional] 

## Methods

### NewCreateUnregisteredTicketRequest

`func NewCreateUnregisteredTicketRequest(agreeWithTerms bool, ticketRequests []CreateTicketRequest, ) *CreateUnregisteredTicketRequest`

NewCreateUnregisteredTicketRequest instantiates a new CreateUnregisteredTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUnregisteredTicketRequestWithDefaults

`func NewCreateUnregisteredTicketRequestWithDefaults() *CreateUnregisteredTicketRequest`

NewCreateUnregisteredTicketRequestWithDefaults instantiates a new CreateUnregisteredTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreeWithTerms

`func (o *CreateUnregisteredTicketRequest) GetAgreeWithTerms() bool`

GetAgreeWithTerms returns the AgreeWithTerms field if non-nil, zero value otherwise.

### GetAgreeWithTermsOk

`func (o *CreateUnregisteredTicketRequest) GetAgreeWithTermsOk() (*bool, bool)`

GetAgreeWithTermsOk returns a tuple with the AgreeWithTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreeWithTerms

`func (o *CreateUnregisteredTicketRequest) SetAgreeWithTerms(v bool)`

SetAgreeWithTerms sets AgreeWithTerms field to given value.


### GetTicketRequests

`func (o *CreateUnregisteredTicketRequest) GetTicketRequests() []CreateTicketRequest`

GetTicketRequests returns the TicketRequests field if non-nil, zero value otherwise.

### GetTicketRequestsOk

`func (o *CreateUnregisteredTicketRequest) GetTicketRequestsOk() (*[]CreateTicketRequest, bool)`

GetTicketRequestsOk returns a tuple with the TicketRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketRequests

`func (o *CreateUnregisteredTicketRequest) SetTicketRequests(v []CreateTicketRequest)`

SetTicketRequests sets TicketRequests field to given value.


### GetAffiliateCode

`func (o *CreateUnregisteredTicketRequest) GetAffiliateCode() string`

GetAffiliateCode returns the AffiliateCode field if non-nil, zero value otherwise.

### GetAffiliateCodeOk

`func (o *CreateUnregisteredTicketRequest) GetAffiliateCodeOk() (*string, bool)`

GetAffiliateCodeOk returns a tuple with the AffiliateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateCode

`func (o *CreateUnregisteredTicketRequest) SetAffiliateCode(v string)`

SetAffiliateCode sets AffiliateCode field to given value.

### HasAffiliateCode

`func (o *CreateUnregisteredTicketRequest) HasAffiliateCode() bool`

HasAffiliateCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


