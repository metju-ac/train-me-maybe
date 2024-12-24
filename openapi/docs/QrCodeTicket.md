# QrCodeTicket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCodeVersion** | Pointer to **int32** | Version of QR code | [optional] 
**QrCodeTicketType** | Pointer to [**QrCodeTicketType**](QrCodeTicketType.md) |  | [optional] 
**TicketId** | Pointer to **int64** | If there is used discount there is as well ticket ID which is linked to it. | [optional] 

## Methods

### NewQrCodeTicket

`func NewQrCodeTicket() *QrCodeTicket`

NewQrCodeTicket instantiates a new QrCodeTicket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQrCodeTicketWithDefaults

`func NewQrCodeTicketWithDefaults() *QrCodeTicket`

NewQrCodeTicketWithDefaults instantiates a new QrCodeTicket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCodeVersion

`func (o *QrCodeTicket) GetQrCodeVersion() int32`

GetQrCodeVersion returns the QrCodeVersion field if non-nil, zero value otherwise.

### GetQrCodeVersionOk

`func (o *QrCodeTicket) GetQrCodeVersionOk() (*int32, bool)`

GetQrCodeVersionOk returns a tuple with the QrCodeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeVersion

`func (o *QrCodeTicket) SetQrCodeVersion(v int32)`

SetQrCodeVersion sets QrCodeVersion field to given value.

### HasQrCodeVersion

`func (o *QrCodeTicket) HasQrCodeVersion() bool`

HasQrCodeVersion returns a boolean if a field has been set.

### GetQrCodeTicketType

`func (o *QrCodeTicket) GetQrCodeTicketType() QrCodeTicketType`

GetQrCodeTicketType returns the QrCodeTicketType field if non-nil, zero value otherwise.

### GetQrCodeTicketTypeOk

`func (o *QrCodeTicket) GetQrCodeTicketTypeOk() (*QrCodeTicketType, bool)`

GetQrCodeTicketTypeOk returns a tuple with the QrCodeTicketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeTicketType

`func (o *QrCodeTicket) SetQrCodeTicketType(v QrCodeTicketType)`

SetQrCodeTicketType sets QrCodeTicketType field to given value.

### HasQrCodeTicketType

`func (o *QrCodeTicket) HasQrCodeTicketType() bool`

HasQrCodeTicketType returns a boolean if a field has been set.

### GetTicketId

`func (o *QrCodeTicket) GetTicketId() int64`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *QrCodeTicket) GetTicketIdOk() (*int64, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *QrCodeTicket) SetTicketId(v int64)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *QrCodeTicket) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


