# KycRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessToken** | Pointer to **string** | Specifies the business account holder on which to perform the identity check. Do not pass this field if your request includes the &#x60;user_token&#x60; field.  Send a &#x60;GET&#x60; request to &#x60;/businesses&#x60; to retrieve business tokens. | [optional] 
**ManualOverride** | Pointer to **bool** | Set to &#x60;true&#x60; to designate a user account holder as having passed a verification check without Marqeta performing the check through one of its KYC providers.  Use this override when you perform verification through another mechanism, such as with an alternative KYC provider or directly with the account holder. You must obtain explicit, written approval from Marqeta before using the &#x60;manual_override&#x60; field for KYC verification. Consult your Customer Success representative for more information. | [optional] [default to false]
**Notes** | Pointer to **string** | Notes pertaining to a manual override. This field is returned in the response only when the &#x60;manual_override&#x60; field is set to &#x60;true&#x60;. | [optional] 
**ReferenceId** | Pointer to **string** | Can be specified only if &#x60;manual_override&#x3D;true&#x60;. If you verified a user account holder&#39;s identity by performing a KYC verification outside of the Marqeta platform, you can use the &#x60;reference_id&#x60; field to store the reference number returned by that KYC provider.  *NOTE:* When you submit a KYC verification request with &#x60;manual_override&#x3D;false&#x60;, the Marqeta platform performs the verification through one of its KYC providers. If the KYC provider responds with a reference identifier, that identifier is passed to you by way of this field in the response. | [optional] 
**Token** | Pointer to **string** | The unique identifier of the identity check.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 
**UserToken** | Pointer to **string** | Specifies the user account holder on which to perform the identity check. Do not pass this field if your request includes the &#x60;business_token&#x60; field.  Send a &#x60;GET&#x60; request to &#x60;/users&#x60; to retrieve user tokens. | [optional] 

## Methods

### NewKycRequest

`func NewKycRequest() *KycRequest`

NewKycRequest instantiates a new KycRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKycRequestWithDefaults

`func NewKycRequestWithDefaults() *KycRequest`

NewKycRequestWithDefaults instantiates a new KycRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessToken

`func (o *KycRequest) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *KycRequest) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *KycRequest) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *KycRequest) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetManualOverride

`func (o *KycRequest) GetManualOverride() bool`

GetManualOverride returns the ManualOverride field if non-nil, zero value otherwise.

### GetManualOverrideOk

`func (o *KycRequest) GetManualOverrideOk() (*bool, bool)`

GetManualOverrideOk returns a tuple with the ManualOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverride

`func (o *KycRequest) SetManualOverride(v bool)`

SetManualOverride sets ManualOverride field to given value.

### HasManualOverride

`func (o *KycRequest) HasManualOverride() bool`

HasManualOverride returns a boolean if a field has been set.

### GetNotes

`func (o *KycRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *KycRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *KycRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *KycRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReferenceId

`func (o *KycRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KycRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KycRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KycRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetToken

`func (o *KycRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KycRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KycRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KycRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *KycRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *KycRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *KycRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *KycRequest) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


