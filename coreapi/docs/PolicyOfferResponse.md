# PolicyOfferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | Date and time when the offer policy was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Description** | Pointer to **string** | Description of the offer policy. | [optional] 
**Name** | Pointer to **string** | Name of the offer policy. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the offer policy. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | Date and time when the offer policy was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 

## Methods

### NewPolicyOfferResponse

`func NewPolicyOfferResponse() *PolicyOfferResponse`

NewPolicyOfferResponse instantiates a new PolicyOfferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyOfferResponseWithDefaults

`func NewPolicyOfferResponseWithDefaults() *PolicyOfferResponse`

NewPolicyOfferResponseWithDefaults instantiates a new PolicyOfferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *PolicyOfferResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PolicyOfferResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PolicyOfferResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PolicyOfferResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyOfferResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyOfferResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyOfferResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyOfferResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PolicyOfferResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyOfferResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyOfferResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyOfferResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *PolicyOfferResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PolicyOfferResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PolicyOfferResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PolicyOfferResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *PolicyOfferResponse) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *PolicyOfferResponse) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *PolicyOfferResponse) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *PolicyOfferResponse) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


