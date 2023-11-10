# DigitalWalletToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressVerification** | Pointer to [**AddressVerification**](AddressVerification.md) |  | [optional] 
**CardToken** | Pointer to **string** | Unique identifier of the card. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the digital wallet token object was created, in UTC. | [optional] 
**Device** | Pointer to [**Device**](Device.md) |  | [optional] 
**FulfillmentStatus** | Pointer to **string** | Digital wallet token&#39;s provisioning status.  For fulfillment status descriptions, see &lt;&lt;/core-api/digital-wallets-management#postDigitalwallettokentransitions, Create digital wallet token transition&gt;&gt;. | [optional] 
**IssuerEligibilityDecision** | Pointer to **string** | The Marqeta platform&#39;s decision as to whether the digital wallet token should be provisioned.  * *0000* – The token should be provisioned.  * *token.activation.verification.required* – Provisioning is pending; further action is required for completion.  For all other values, check the value of the &#x60;fulfillment_status&#x60; field to definitively ascertain the provisioning outcome.  *NOTE:* The value &#x60;invalid.cid&#x60; indicates an invalid CVV2 number. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the digital wallet token object was last modified, in UTC. | [optional] 
**Metadata** | Pointer to [**DigitalWalletTokenMetadata**](DigitalWalletTokenMetadata.md) |  | [optional] 
**State** | Pointer to **string** | State of the digital wallet token.  For state descriptions, see &lt;&lt;/developer-guides/managing-the-digital-wallet-token-lifecycle#_transitioning_token_states, Transitioning Token States&gt;&gt;. | [optional] 
**StateReason** | Pointer to **string** | Reason why the digital wallet token transitioned to its current state. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the digital wallet token. | [optional] 
**TokenServiceProvider** | Pointer to [**TokenServiceProvider**](TokenServiceProvider.md) |  | [optional] 
**User** | Pointer to [**UserCardHolderResponse**](UserCardHolderResponse.md) |  | [optional] 
**WalletProviderProfile** | Pointer to [**WalletProviderProfile**](WalletProviderProfile.md) |  | [optional] 

## Methods

### NewDigitalWalletToken

`func NewDigitalWalletToken() *DigitalWalletToken`

NewDigitalWalletToken instantiates a new DigitalWalletToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletTokenWithDefaults

`func NewDigitalWalletTokenWithDefaults() *DigitalWalletToken`

NewDigitalWalletTokenWithDefaults instantiates a new DigitalWalletToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressVerification

`func (o *DigitalWalletToken) GetAddressVerification() AddressVerification`

GetAddressVerification returns the AddressVerification field if non-nil, zero value otherwise.

### GetAddressVerificationOk

`func (o *DigitalWalletToken) GetAddressVerificationOk() (*AddressVerification, bool)`

GetAddressVerificationOk returns a tuple with the AddressVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerification

`func (o *DigitalWalletToken) SetAddressVerification(v AddressVerification)`

SetAddressVerification sets AddressVerification field to given value.

### HasAddressVerification

`func (o *DigitalWalletToken) HasAddressVerification() bool`

HasAddressVerification returns a boolean if a field has been set.

### GetCardToken

`func (o *DigitalWalletToken) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *DigitalWalletToken) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *DigitalWalletToken) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.

### HasCardToken

`func (o *DigitalWalletToken) HasCardToken() bool`

HasCardToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DigitalWalletToken) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DigitalWalletToken) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DigitalWalletToken) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DigitalWalletToken) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDevice

`func (o *DigitalWalletToken) GetDevice() Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DigitalWalletToken) GetDeviceOk() (*Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DigitalWalletToken) SetDevice(v Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DigitalWalletToken) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *DigitalWalletToken) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *DigitalWalletToken) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *DigitalWalletToken) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *DigitalWalletToken) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### GetIssuerEligibilityDecision

`func (o *DigitalWalletToken) GetIssuerEligibilityDecision() string`

GetIssuerEligibilityDecision returns the IssuerEligibilityDecision field if non-nil, zero value otherwise.

### GetIssuerEligibilityDecisionOk

`func (o *DigitalWalletToken) GetIssuerEligibilityDecisionOk() (*string, bool)`

GetIssuerEligibilityDecisionOk returns a tuple with the IssuerEligibilityDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerEligibilityDecision

`func (o *DigitalWalletToken) SetIssuerEligibilityDecision(v string)`

SetIssuerEligibilityDecision sets IssuerEligibilityDecision field to given value.

### HasIssuerEligibilityDecision

`func (o *DigitalWalletToken) HasIssuerEligibilityDecision() bool`

HasIssuerEligibilityDecision returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *DigitalWalletToken) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *DigitalWalletToken) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *DigitalWalletToken) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *DigitalWalletToken) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMetadata

`func (o *DigitalWalletToken) GetMetadata() DigitalWalletTokenMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DigitalWalletToken) GetMetadataOk() (*DigitalWalletTokenMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DigitalWalletToken) SetMetadata(v DigitalWalletTokenMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DigitalWalletToken) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetState

`func (o *DigitalWalletToken) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DigitalWalletToken) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DigitalWalletToken) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DigitalWalletToken) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *DigitalWalletToken) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *DigitalWalletToken) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *DigitalWalletToken) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *DigitalWalletToken) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetToken

`func (o *DigitalWalletToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DigitalWalletToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DigitalWalletToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DigitalWalletToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenServiceProvider

`func (o *DigitalWalletToken) GetTokenServiceProvider() TokenServiceProvider`

GetTokenServiceProvider returns the TokenServiceProvider field if non-nil, zero value otherwise.

### GetTokenServiceProviderOk

`func (o *DigitalWalletToken) GetTokenServiceProviderOk() (*TokenServiceProvider, bool)`

GetTokenServiceProviderOk returns a tuple with the TokenServiceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenServiceProvider

`func (o *DigitalWalletToken) SetTokenServiceProvider(v TokenServiceProvider)`

SetTokenServiceProvider sets TokenServiceProvider field to given value.

### HasTokenServiceProvider

`func (o *DigitalWalletToken) HasTokenServiceProvider() bool`

HasTokenServiceProvider returns a boolean if a field has been set.

### GetUser

`func (o *DigitalWalletToken) GetUser() UserCardHolderResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DigitalWalletToken) GetUserOk() (*UserCardHolderResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DigitalWalletToken) SetUser(v UserCardHolderResponse)`

SetUser sets User field to given value.

### HasUser

`func (o *DigitalWalletToken) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWalletProviderProfile

`func (o *DigitalWalletToken) GetWalletProviderProfile() WalletProviderProfile`

GetWalletProviderProfile returns the WalletProviderProfile field if non-nil, zero value otherwise.

### GetWalletProviderProfileOk

`func (o *DigitalWalletToken) GetWalletProviderProfileOk() (*WalletProviderProfile, bool)`

GetWalletProviderProfileOk returns a tuple with the WalletProviderProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletProviderProfile

`func (o *DigitalWalletToken) SetWalletProviderProfile(v WalletProviderProfile)`

SetWalletProviderProfile sets WalletProviderProfile field to given value.

### HasWalletProviderProfile

`func (o *DigitalWalletToken) HasWalletProviderProfile() bool`

HasWalletProviderProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


