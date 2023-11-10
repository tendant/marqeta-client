# ValidationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserValidationResponse**](UserValidationResponse.md) |  | 

## Methods

### NewValidationsResponse

`func NewValidationsResponse(user UserValidationResponse, ) *ValidationsResponse`

NewValidationsResponse instantiates a new ValidationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationsResponseWithDefaults

`func NewValidationsResponseWithDefaults() *ValidationsResponse`

NewValidationsResponseWithDefaults instantiates a new ValidationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ValidationsResponse) GetUser() UserValidationResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ValidationsResponse) GetUserOk() (*UserValidationResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ValidationsResponse) SetUser(v UserValidationResponse)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


