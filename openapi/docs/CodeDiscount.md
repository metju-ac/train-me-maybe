# CodeDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Flat rate discount ID | 
**Code** | **string** | Five-digit discount code | 
**Discount** | **float32** | Flat rate discount sum in ticket currency | 

## Methods

### NewCodeDiscount

`func NewCodeDiscount(id int64, code string, discount float32, ) *CodeDiscount`

NewCodeDiscount instantiates a new CodeDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeDiscountWithDefaults

`func NewCodeDiscountWithDefaults() *CodeDiscount`

NewCodeDiscountWithDefaults instantiates a new CodeDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CodeDiscount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeDiscount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeDiscount) SetId(v int64)`

SetId sets Id field to given value.


### GetCode

`func (o *CodeDiscount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CodeDiscount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CodeDiscount) SetCode(v string)`

SetCode sets Code field to given value.


### GetDiscount

`func (o *CodeDiscount) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *CodeDiscount) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *CodeDiscount) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


