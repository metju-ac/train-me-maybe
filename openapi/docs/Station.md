# Station

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique identifier for the station. | 
**Name** | **string** | Station name. | 
**Fullname** | **string** | Station full name. | 
**Aliases** | Pointer to **[]string** | Station nicknames or aliases. | [optional] 
**Address** | Pointer to **string** | Station address. | [optional] 
**StationsTypes** | Pointer to **[]string** | Station type. | [optional] 
**IataCode** | Pointer to **string** | A three-letter from [IATA airport code](https://en.wikipedia.org/wiki/IATA_airport_code). | [optional] 
**StationUrl** | Pointer to **string** | Publicly accessible station URL. | [optional] 
**StationTimeZoneCode** | Pointer to **string** | [TZ database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) time zone name in which station is located. | [optional] 
**WheelChairPlatform** | Pointer to **string** | \&quot;Name of wheelchair platform provider (if available).\&quot;  Possible providers:    | Provider name | Provider code |   |---------------|---------------|   | České dráhy   | CD            |   | RegioJet      | RJ            |  | [optional] 
**Significance** | **int32** | Station&#39;s significance within the city. (e.g. for the city of Brno, Brno main railway station is more significant than Brno - Královo Pole railway station.) | 
**Longitude** | Pointer to **float32** | The geographic longitude where station is located. | [optional] 
**Latitude** | Pointer to **float32** | The geographic latitude where station is located. | [optional] 
**ImageUrl** | Pointer to **string** | URL to station&#39;s picture (if available). | [optional] 

## Methods

### NewStation

`func NewStation(id int64, name string, fullname string, significance int32, ) *Station`

NewStation instantiates a new Station object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStationWithDefaults

`func NewStationWithDefaults() *Station`

NewStationWithDefaults instantiates a new Station object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Station) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Station) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Station) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Station) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Station) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Station) SetName(v string)`

SetName sets Name field to given value.


### GetFullname

`func (o *Station) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *Station) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *Station) SetFullname(v string)`

SetFullname sets Fullname field to given value.


### GetAliases

`func (o *Station) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Station) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Station) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Station) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAddress

`func (o *Station) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Station) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Station) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Station) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStationsTypes

`func (o *Station) GetStationsTypes() []string`

GetStationsTypes returns the StationsTypes field if non-nil, zero value otherwise.

### GetStationsTypesOk

`func (o *Station) GetStationsTypesOk() (*[]string, bool)`

GetStationsTypesOk returns a tuple with the StationsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationsTypes

`func (o *Station) SetStationsTypes(v []string)`

SetStationsTypes sets StationsTypes field to given value.

### HasStationsTypes

`func (o *Station) HasStationsTypes() bool`

HasStationsTypes returns a boolean if a field has been set.

### GetIataCode

`func (o *Station) GetIataCode() string`

GetIataCode returns the IataCode field if non-nil, zero value otherwise.

### GetIataCodeOk

`func (o *Station) GetIataCodeOk() (*string, bool)`

GetIataCodeOk returns a tuple with the IataCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIataCode

`func (o *Station) SetIataCode(v string)`

SetIataCode sets IataCode field to given value.

### HasIataCode

`func (o *Station) HasIataCode() bool`

HasIataCode returns a boolean if a field has been set.

### GetStationUrl

`func (o *Station) GetStationUrl() string`

GetStationUrl returns the StationUrl field if non-nil, zero value otherwise.

### GetStationUrlOk

`func (o *Station) GetStationUrlOk() (*string, bool)`

GetStationUrlOk returns a tuple with the StationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationUrl

`func (o *Station) SetStationUrl(v string)`

SetStationUrl sets StationUrl field to given value.

### HasStationUrl

`func (o *Station) HasStationUrl() bool`

HasStationUrl returns a boolean if a field has been set.

### GetStationTimeZoneCode

`func (o *Station) GetStationTimeZoneCode() string`

GetStationTimeZoneCode returns the StationTimeZoneCode field if non-nil, zero value otherwise.

### GetStationTimeZoneCodeOk

`func (o *Station) GetStationTimeZoneCodeOk() (*string, bool)`

GetStationTimeZoneCodeOk returns a tuple with the StationTimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationTimeZoneCode

`func (o *Station) SetStationTimeZoneCode(v string)`

SetStationTimeZoneCode sets StationTimeZoneCode field to given value.

### HasStationTimeZoneCode

`func (o *Station) HasStationTimeZoneCode() bool`

HasStationTimeZoneCode returns a boolean if a field has been set.

### GetWheelChairPlatform

`func (o *Station) GetWheelChairPlatform() string`

GetWheelChairPlatform returns the WheelChairPlatform field if non-nil, zero value otherwise.

### GetWheelChairPlatformOk

`func (o *Station) GetWheelChairPlatformOk() (*string, bool)`

GetWheelChairPlatformOk returns a tuple with the WheelChairPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWheelChairPlatform

`func (o *Station) SetWheelChairPlatform(v string)`

SetWheelChairPlatform sets WheelChairPlatform field to given value.

### HasWheelChairPlatform

`func (o *Station) HasWheelChairPlatform() bool`

HasWheelChairPlatform returns a boolean if a field has been set.

### GetSignificance

`func (o *Station) GetSignificance() int32`

GetSignificance returns the Significance field if non-nil, zero value otherwise.

### GetSignificanceOk

`func (o *Station) GetSignificanceOk() (*int32, bool)`

GetSignificanceOk returns a tuple with the Significance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificance

`func (o *Station) SetSignificance(v int32)`

SetSignificance sets Significance field to given value.


### GetLongitude

`func (o *Station) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Station) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Station) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Station) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetLatitude

`func (o *Station) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Station) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Station) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Station) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetImageUrl

`func (o *Station) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Station) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Station) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Station) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


