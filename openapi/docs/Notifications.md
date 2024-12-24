# Notifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Newsletter** | **bool** |  | [default to false]
**ReservationChange** | **bool** |  | [default to false]
**RouteRatingSurvey** | **bool** |  | [default to false]

## Methods

### NewNotifications

`func NewNotifications(newsletter bool, reservationChange bool, routeRatingSurvey bool, ) *Notifications`

NewNotifications instantiates a new Notifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsWithDefaults

`func NewNotificationsWithDefaults() *Notifications`

NewNotificationsWithDefaults instantiates a new Notifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewsletter

`func (o *Notifications) GetNewsletter() bool`

GetNewsletter returns the Newsletter field if non-nil, zero value otherwise.

### GetNewsletterOk

`func (o *Notifications) GetNewsletterOk() (*bool, bool)`

GetNewsletterOk returns a tuple with the Newsletter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsletter

`func (o *Notifications) SetNewsletter(v bool)`

SetNewsletter sets Newsletter field to given value.


### GetReservationChange

`func (o *Notifications) GetReservationChange() bool`

GetReservationChange returns the ReservationChange field if non-nil, zero value otherwise.

### GetReservationChangeOk

`func (o *Notifications) GetReservationChangeOk() (*bool, bool)`

GetReservationChangeOk returns a tuple with the ReservationChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationChange

`func (o *Notifications) SetReservationChange(v bool)`

SetReservationChange sets ReservationChange field to given value.


### GetRouteRatingSurvey

`func (o *Notifications) GetRouteRatingSurvey() bool`

GetRouteRatingSurvey returns the RouteRatingSurvey field if non-nil, zero value otherwise.

### GetRouteRatingSurveyOk

`func (o *Notifications) GetRouteRatingSurveyOk() (*bool, bool)`

GetRouteRatingSurveyOk returns a tuple with the RouteRatingSurvey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteRatingSurvey

`func (o *Notifications) SetRouteRatingSurvey(v bool)`

SetRouteRatingSurvey sets RouteRatingSurvey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


