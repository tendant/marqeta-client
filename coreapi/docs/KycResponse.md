# KycResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | Pointer to **string** | The business account holder on which the identity check was performed. | [optional] 
**CreatedTime** | **time.Time** | Time when the KYC verification was performed. | 
**LastModifiedTime** | **time.Time** | Time when the KYC verification was last updated. | 
**ManualOverride** | Pointer to **bool** | If &#x60;true&#x60;, the user account holder is designated as having passed a verification check without Marqeta performing the check.  This override is used when verification is performed through another mechanism, such as with an alternative KYC provider or directly with the account holder. | [optional] [readonly] [default to false]
**Notes** | Pointer to **string** | Notes pertaining to a manual override. This field is included in the response only when the &#x60;manual_override&#x60; field is set to &#x60;true&#x60;. | [optional] [readonly] 
**ReferenceId** | Pointer to **string** | If you verified the account holder&#39;s identity by performing a KYC verification outside of the Marqeta platform, you can use the &#x60;reference_id&#x60; field to store the reference number returned by that KYC provider. This field is included in the response only when the &#x60;manual_override&#x60; field is set to &#x60;true&#x60;.  *NOTE:* When you submit a KYC verification request with &#x60;manual_override&#x3D;false&#x60;, the Marqeta platform performs the verification through one of its KYC providers. If the KYC provider responds with a reference identifier, that identifier is passed to you by way of this field in the response. | [optional] 
**Result** | Pointer to [**Result**](Result.md) |  | [optional] 
**Token** | Pointer to **string** | The unique identifier of the identity check. | [optional] 
**UserToken** | Pointer to **string** | The user account holder on which the identity check was performed. | [optional] 

## Methods

### NewKycResponse

`func NewKycResponse(createdTime time.Time, lastModifiedTime time.Time, ) *KycResponse`

NewKycResponse instantiates a new KycResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycResponseWithDefaults

`func NewKycResponseWithDefaults() *KycResponse`

NewKycResponseWithDefaults instantiates a new KycResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *KycResponse) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *KycResponse) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *KycResponse) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *KycResponse) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *KycResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *KycResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *KycResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetLastModifiedTime

`func (o *KycResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *KycResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *KycResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetManualOverride

`func (o *KycResponse) GetManualOverride() bool`

GetManualOverride returns the ManualOverride field if non-nil, zero value otherwise.

### GetManualOverrideOk

`func (o *KycResponse) GetManualOverrideOk() (*bool, bool)`

GetManualOverrideOk returns a tuple with the ManualOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverride

`func (o *KycResponse) SetManualOverride(v bool)`

SetManualOverride sets ManualOverride field to given value.

### HasManualOverride

`func (o *KycResponse) HasManualOverride() bool`

HasManualOverride returns a boolean if a field has been set.

### GetNotes

`func (o *KycResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *KycResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *KycResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *KycResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReferenceId

`func (o *KycResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KycResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KycResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KycResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetResult

`func (o *KycResponse) GetResult() Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *KycResponse) GetResultOk() (*Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *KycResponse) SetResult(v Result)`

SetResult sets Result field to given value.

### HasResult

`func (o *KycResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetToken

`func (o *KycResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KycResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KycResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KycResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *KycResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *KycResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *KycResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *KycResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


