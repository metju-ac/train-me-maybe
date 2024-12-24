# BannerBubble

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Text** | **string** | Extra information for bubble. | 
**Url** | Pointer to **string** | Redirection to more information. | [optional] 
**ImageUrl** | **string** | Bubble picture URL address. | 

## Methods

### NewBannerBubble

`func NewBannerBubble(id int64, text string, imageUrl string, ) *BannerBubble`

NewBannerBubble instantiates a new BannerBubble object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannerBubbleWithDefaults

`func NewBannerBubbleWithDefaults() *BannerBubble`

NewBannerBubbleWithDefaults instantiates a new BannerBubble object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BannerBubble) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BannerBubble) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BannerBubble) SetId(v int64)`

SetId sets Id field to given value.


### GetText

`func (o *BannerBubble) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BannerBubble) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BannerBubble) SetText(v string)`

SetText sets Text field to given value.


### GetUrl

`func (o *BannerBubble) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BannerBubble) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BannerBubble) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BannerBubble) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *BannerBubble) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *BannerBubble) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *BannerBubble) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


