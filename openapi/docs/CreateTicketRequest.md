# CreateTicketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Route** | [**CreateTicketRouteRequest**](CreateTicketRouteRequest.md) |  | 
**SelectedAddons** | Pointer to [**[]SelectedAddon**](SelectedAddon.md) |  | [optional] 
**Passengers** | [**[]Passenger**](Passenger.md) |  | 
**CodeDiscount** | Pointer to **string** | Flat rate discount from fare price (does not apply on addons and charges). Applies first (before percentual discount) | [optional] 
**PercentualDiscountIds** | Pointer to **[]int64** | Percentual discount from fare price (does not apply on addons and charges). Applies as seconds (after flat rate discount) | [optional] 

## Methods

### NewCreateTicketRequest

`func NewCreateTicketRequest(route CreateTicketRouteRequest, passengers []Passenger, ) *CreateTicketRequest`

NewCreateTicketRequest instantiates a new CreateTicketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTicketRequestWithDefaults

`func NewCreateTicketRequestWithDefaults() *CreateTicketRequest`

NewCreateTicketRequestWithDefaults instantiates a new CreateTicketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoute

`func (o *CreateTicketRequest) GetRoute() CreateTicketRouteRequest`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *CreateTicketRequest) GetRouteOk() (*CreateTicketRouteRequest, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *CreateTicketRequest) SetRoute(v CreateTicketRouteRequest)`

SetRoute sets Route field to given value.


### GetSelectedAddons

`func (o *CreateTicketRequest) GetSelectedAddons() []SelectedAddon`

GetSelectedAddons returns the SelectedAddons field if non-nil, zero value otherwise.

### GetSelectedAddonsOk

`func (o *CreateTicketRequest) GetSelectedAddonsOk() (*[]SelectedAddon, bool)`

GetSelectedAddonsOk returns a tuple with the SelectedAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedAddons

`func (o *CreateTicketRequest) SetSelectedAddons(v []SelectedAddon)`

SetSelectedAddons sets SelectedAddons field to given value.

### HasSelectedAddons

`func (o *CreateTicketRequest) HasSelectedAddons() bool`

HasSelectedAddons returns a boolean if a field has been set.

### GetPassengers

`func (o *CreateTicketRequest) GetPassengers() []Passenger`

GetPassengers returns the Passengers field if non-nil, zero value otherwise.

### GetPassengersOk

`func (o *CreateTicketRequest) GetPassengersOk() (*[]Passenger, bool)`

GetPassengersOk returns a tuple with the Passengers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengers

`func (o *CreateTicketRequest) SetPassengers(v []Passenger)`

SetPassengers sets Passengers field to given value.


### GetCodeDiscount

`func (o *CreateTicketRequest) GetCodeDiscount() string`

GetCodeDiscount returns the CodeDiscount field if non-nil, zero value otherwise.

### GetCodeDiscountOk

`func (o *CreateTicketRequest) GetCodeDiscountOk() (*string, bool)`

GetCodeDiscountOk returns a tuple with the CodeDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeDiscount

`func (o *CreateTicketRequest) SetCodeDiscount(v string)`

SetCodeDiscount sets CodeDiscount field to given value.

### HasCodeDiscount

`func (o *CreateTicketRequest) HasCodeDiscount() bool`

HasCodeDiscount returns a boolean if a field has been set.

### GetPercentualDiscountIds

`func (o *CreateTicketRequest) GetPercentualDiscountIds() []int64`

GetPercentualDiscountIds returns the PercentualDiscountIds field if non-nil, zero value otherwise.

### GetPercentualDiscountIdsOk

`func (o *CreateTicketRequest) GetPercentualDiscountIdsOk() (*[]int64, bool)`

GetPercentualDiscountIdsOk returns a tuple with the PercentualDiscountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentualDiscountIds

`func (o *CreateTicketRequest) SetPercentualDiscountIds(v []int64)`

SetPercentualDiscountIds sets PercentualDiscountIds field to given value.

### HasPercentualDiscountIds

`func (o *CreateTicketRequest) HasPercentualDiscountIds() bool`

HasPercentualDiscountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


