# BundleTransitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleToken** | Pointer to **string** | Unique identifier of the bundle. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the bundle was changed the status on Marqeta&#39;s credit platform, in UTC. | [optional] 
**OriginalStatus** | Pointer to [**BundleResourceStatus**](BundleResourceStatus.md) |  | [optional] 
**Status** | Pointer to [**BundleResourceStatus**](BundleResourceStatus.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the bundle transition. | [optional] 

## Methods

### NewBundleTransitionResponse

`func NewBundleTransitionResponse() *BundleTransitionResponse`

NewBundleTransitionResponse instantiates a new BundleTransitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleTransitionResponseWithDefaults

`func NewBundleTransitionResponseWithDefaults() *BundleTransitionResponse`

NewBundleTransitionResponseWithDefaults instantiates a new BundleTransitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleToken

`func (o *BundleTransitionResponse) GetBundleToken() string`

GetBundleToken returns the BundleToken field if non-nil, zero value otherwise.

### GetBundleTokenOk

`func (o *BundleTransitionResponse) GetBundleTokenOk() (*string, bool)`

GetBundleTokenOk returns a tuple with the BundleToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleToken

`func (o *BundleTransitionResponse) SetBundleToken(v string)`

SetBundleToken sets BundleToken field to given value.

### HasBundleToken

`func (o *BundleTransitionResponse) HasBundleToken() bool`

HasBundleToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BundleTransitionResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BundleTransitionResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BundleTransitionResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BundleTransitionResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetOriginalStatus

`func (o *BundleTransitionResponse) GetOriginalStatus() BundleResourceStatus`

GetOriginalStatus returns the OriginalStatus field if non-nil, zero value otherwise.

### GetOriginalStatusOk

`func (o *BundleTransitionResponse) GetOriginalStatusOk() (*BundleResourceStatus, bool)`

GetOriginalStatusOk returns a tuple with the OriginalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStatus

`func (o *BundleTransitionResponse) SetOriginalStatus(v BundleResourceStatus)`

SetOriginalStatus sets OriginalStatus field to given value.

### HasOriginalStatus

`func (o *BundleTransitionResponse) HasOriginalStatus() bool`

HasOriginalStatus returns a boolean if a field has been set.

### GetStatus

`func (o *BundleTransitionResponse) GetStatus() BundleResourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BundleTransitionResponse) GetStatusOk() (*BundleResourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BundleTransitionResponse) SetStatus(v BundleResourceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BundleTransitionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *BundleTransitionResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BundleTransitionResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BundleTransitionResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *BundleTransitionResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


