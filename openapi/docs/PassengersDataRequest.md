# PassengersDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sections** | [**[]SimpleSection**](SimpleSection.md) |  | 
**Tariffs** | **[]string** |  | 
**SeatClass** | **string** |  | 
**PriceSource** | **string** | Pricing ID - used for confirmation that price, services or conditions werent changed | 

## Methods

### NewPassengersDataRequest

`func NewPassengersDataRequest(sections []SimpleSection, tariffs []string, seatClass string, priceSource string, ) *PassengersDataRequest`

NewPassengersDataRequest instantiates a new PassengersDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassengersDataRequestWithDefaults

`func NewPassengersDataRequestWithDefaults() *PassengersDataRequest`

NewPassengersDataRequestWithDefaults instantiates a new PassengersDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSections

`func (o *PassengersDataRequest) GetSections() []SimpleSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *PassengersDataRequest) GetSectionsOk() (*[]SimpleSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *PassengersDataRequest) SetSections(v []SimpleSection)`

SetSections sets Sections field to given value.


### GetTariffs

`func (o *PassengersDataRequest) GetTariffs() []string`

GetTariffs returns the Tariffs field if non-nil, zero value otherwise.

### GetTariffsOk

`func (o *PassengersDataRequest) GetTariffsOk() (*[]string, bool)`

GetTariffsOk returns a tuple with the Tariffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffs

`func (o *PassengersDataRequest) SetTariffs(v []string)`

SetTariffs sets Tariffs field to given value.


### GetSeatClass

`func (o *PassengersDataRequest) GetSeatClass() string`

GetSeatClass returns the SeatClass field if non-nil, zero value otherwise.

### GetSeatClassOk

`func (o *PassengersDataRequest) GetSeatClassOk() (*string, bool)`

GetSeatClassOk returns a tuple with the SeatClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatClass

`func (o *PassengersDataRequest) SetSeatClass(v string)`

SetSeatClass sets SeatClass field to given value.


### GetPriceSource

`func (o *PassengersDataRequest) GetPriceSource() string`

GetPriceSource returns the PriceSource field if non-nil, zero value otherwise.

### GetPriceSourceOk

`func (o *PassengersDataRequest) GetPriceSourceOk() (*string, bool)`

GetPriceSourceOk returns a tuple with the PriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSource

`func (o *PassengersDataRequest) SetPriceSource(v string)`

SetPriceSource sets PriceSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


