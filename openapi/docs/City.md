# City

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the city. | 
**Name** | **string** | City name. | 
**Aliases** | Pointer to **[]string** | City nicknames or aliases. | [optional] 
**StationsTypes** | Pointer to **[]string** | Types of stations available in a city. | [optional] 
**Stations** | [**[]Station**](Station.md) |  | 

## Methods

### NewCity

`func NewCity(id int64, name string, stations []Station, ) *City`

NewCity instantiates a new City object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCityWithDefaults

`func NewCityWithDefaults() *City`

NewCityWithDefaults instantiates a new City object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *City) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *City) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *City) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *City) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *City) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *City) SetName(v string)`

SetName sets Name field to given value.


### GetAliases

`func (o *City) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *City) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *City) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *City) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetStationsTypes

`func (o *City) GetStationsTypes() []string`

GetStationsTypes returns the StationsTypes field if non-nil, zero value otherwise.

### GetStationsTypesOk

`func (o *City) GetStationsTypesOk() (*[]string, bool)`

GetStationsTypesOk returns a tuple with the StationsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationsTypes

`func (o *City) SetStationsTypes(v []string)`

SetStationsTypes sets StationsTypes field to given value.

### HasStationsTypes

`func (o *City) HasStationsTypes() bool`

HasStationsTypes returns a boolean if a field has been set.

### GetStations

`func (o *City) GetStations() []Station`

GetStations returns the Stations field if non-nil, zero value otherwise.

### GetStationsOk

`func (o *City) GetStationsOk() (*[]Station, bool)`

GetStationsOk returns a tuple with the Stations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStations

`func (o *City) SetStations(v []Station)`

SetStations sets Stations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


