# CardholderAuthenticationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirerExemption** | Pointer to **[]string** | Indicates 3D Secure authentication exemptions from the acquirer. This array is returned if it is included in the transaction data from the card network. | [optional] 
**AuthenticationMethod** | Pointer to **string** | Specifies the 3D Secure authentication method. | [optional] 
**AuthenticationStatus** | Pointer to **string** | Specifies the status of the 3D Secure authentication:  * &#x60;ATTEMPTED&#x60;: Indicates that 3D Secure authentication was requested and processed by the card network. * &#x60;DATA_SHARE_EXEMPTION&#x60;: Indicates that 3D Secure authentication was for information only and exempted. * &#x60;EXEMPTED&#x60;: Indicates that 3D Secure authentication was requested but the challenge was exempted. * &#x60;EXEMPTION_CLAIMED&#x60;: Indicates that 3D Secure authentication was exempted because acquirer transaction risk analysis (TRA) was already performed. * &#x60;SCA_EXEMPTION&#x60;: Indicates that 3D Secure authentication was exempted because strong customer authentication (SCA) was already performed. * &#x60;SUCCESSFUL&#x60;: Indicates that 3D Secure authentication was successful. * &#x60;SUCCESSFUL_NON_PAYMENT&#x60;: Indicates that 3D Secure non-payment authentication was successful. * &#x60;THREEDS_REQUESTER_DATA_SHARE_EXEMPTION&#x60;: Status is deprecated and will be removed from a future release of the Marqeta platform. After June 2023, use &#x60;DATA_SHARE_EXEMPTION&#x60; instead. * &#x60;THREEDS_REQUESTER_SCA_EXEMPTION&#x60;: Status is deprecated and will be removed in a June 2023 release of the Marqeta platform. After June 2023, use &#x60;SCA_EXEMPTION&#x60; instead. * &#x60;THREEDS_REQUESTER_TRA_EXEMPTION&#x60;: Status is deprecated and will be removed in a June 2023 release of the Marqeta platform. After June 2023, use &#x60;EXEMPTION_CLAIMED&#x60; instead. * &#x60;UNAVAILABLE&#x60;: ** For Visa transactions, this status indicates that 3D Secure authentication was requested, but Marqeta&#39;s access control server (ACS) was not available. ** For Mastercard transactions, this status indicates that 3D Secure authentication was not available. | [optional] 
**CavvAuthenticationAmount** | Pointer to **string** | Authentication amount from the cardholder authentication verification value (CAVV) used to validate an authorization. This field is returned if it is included in the transaction data from the card network. | [optional] 
**ElectronicCommerceIndicator** | Pointer to **string** | Status of the 3D Secure authentication attempt, as provided by a transaction participant.  * &#x60;authentication_attempted&#x60;: Merchant attempted to authenticate, but either the issuer or the cardholder does not participate in 3D Secure. * &#x60;authentication_successful&#x60;: Cardholder authentication successful. * &#x60;no_authentication&#x60;: Non-authenticated e-commerce transaction. | [optional] 
**IssuerExemption** | Pointer to **string** | Indicates a 3D Secure authentication exemption from the issuer This field is returned if it is included in the transaction data from the card network. | [optional] 
**RawCavvData** | Pointer to **string** | Raw cardholder authentication verification value provided by the card network. This field is returned if it is included in the transaction data from the card network. | [optional] 
**ThreeDsMessageVersion** | Pointer to **string** | Specifies the 3D Secure message version used for authentication. | [optional] 
**VerificationResult** | Pointer to **string** | Result of cardholder authentication verification value (CAVV) or accountholder authentication value (AAV) verification performed by the card network.  * &#x60;failed&#x60; * &#x60;not_present&#x60; * &#x60;not_provided&#x60; * &#x60;not_verified&#x60; * &#x60;not_verified_authentication_outage&#x60; * &#x60;verified&#x60; * &#x60;verified_amount_checked&#x60; * &#x60;verified_amount_greater_than_20_percent&#x60;: For Mastercard AAV verification, indicates that the original authentication amount and final authorization amount are mismatched, and that the final authorization amount exceeds the original authentication amount by more than 20%. This 20% margin falls outside Mastercard&#39;s suggested tolerance for what a European cardholder might reasonably expect when the total transaction amount is not known in advance. * &#x60;verified_amount_less_than_20_percent&#x60;: For Mastercard AAV verification, indicates that the original authentication amount and final authorization amount are mismatched, and that the final authorization amount exceeds the original authentication amount by 20% or less. This 20% margin falls within Mastercard&#39;s suggested tolerance for what a European cardholder might reasonably expect when the total transaction amount is not known in advance. | [optional] 
**VerificationValueCreatedBy** | Pointer to **string** | Transaction participant who determined the verification result. | [optional] 

## Methods

### NewCardholderAuthenticationData

`func NewCardholderAuthenticationData() *CardholderAuthenticationData`

NewCardholderAuthenticationData instantiates a new CardholderAuthenticationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardholderAuthenticationDataWithDefaults

`func NewCardholderAuthenticationDataWithDefaults() *CardholderAuthenticationData`

NewCardholderAuthenticationDataWithDefaults instantiates a new CardholderAuthenticationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquirerExemption

`func (o *CardholderAuthenticationData) GetAcquirerExemption() []string`

GetAcquirerExemption returns the AcquirerExemption field if non-nil, zero value otherwise.

### GetAcquirerExemptionOk

`func (o *CardholderAuthenticationData) GetAcquirerExemptionOk() (*[]string, bool)`

GetAcquirerExemptionOk returns a tuple with the AcquirerExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerExemption

`func (o *CardholderAuthenticationData) SetAcquirerExemption(v []string)`

SetAcquirerExemption sets AcquirerExemption field to given value.

### HasAcquirerExemption

`func (o *CardholderAuthenticationData) HasAcquirerExemption() bool`

HasAcquirerExemption returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *CardholderAuthenticationData) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *CardholderAuthenticationData) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *CardholderAuthenticationData) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *CardholderAuthenticationData) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.

### GetAuthenticationStatus

`func (o *CardholderAuthenticationData) GetAuthenticationStatus() string`

GetAuthenticationStatus returns the AuthenticationStatus field if non-nil, zero value otherwise.

### GetAuthenticationStatusOk

`func (o *CardholderAuthenticationData) GetAuthenticationStatusOk() (*string, bool)`

GetAuthenticationStatusOk returns a tuple with the AuthenticationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationStatus

`func (o *CardholderAuthenticationData) SetAuthenticationStatus(v string)`

SetAuthenticationStatus sets AuthenticationStatus field to given value.

### HasAuthenticationStatus

`func (o *CardholderAuthenticationData) HasAuthenticationStatus() bool`

HasAuthenticationStatus returns a boolean if a field has been set.

### GetCavvAuthenticationAmount

`func (o *CardholderAuthenticationData) GetCavvAuthenticationAmount() string`

GetCavvAuthenticationAmount returns the CavvAuthenticationAmount field if non-nil, zero value otherwise.

### GetCavvAuthenticationAmountOk

`func (o *CardholderAuthenticationData) GetCavvAuthenticationAmountOk() (*string, bool)`

GetCavvAuthenticationAmountOk returns a tuple with the CavvAuthenticationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavvAuthenticationAmount

`func (o *CardholderAuthenticationData) SetCavvAuthenticationAmount(v string)`

SetCavvAuthenticationAmount sets CavvAuthenticationAmount field to given value.

### HasCavvAuthenticationAmount

`func (o *CardholderAuthenticationData) HasCavvAuthenticationAmount() bool`

HasCavvAuthenticationAmount returns a boolean if a field has been set.

### GetElectronicCommerceIndicator

`func (o *CardholderAuthenticationData) GetElectronicCommerceIndicator() string`

GetElectronicCommerceIndicator returns the ElectronicCommerceIndicator field if non-nil, zero value otherwise.

### GetElectronicCommerceIndicatorOk

`func (o *CardholderAuthenticationData) GetElectronicCommerceIndicatorOk() (*string, bool)`

GetElectronicCommerceIndicatorOk returns a tuple with the ElectronicCommerceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicCommerceIndicator

`func (o *CardholderAuthenticationData) SetElectronicCommerceIndicator(v string)`

SetElectronicCommerceIndicator sets ElectronicCommerceIndicator field to given value.

### HasElectronicCommerceIndicator

`func (o *CardholderAuthenticationData) HasElectronicCommerceIndicator() bool`

HasElectronicCommerceIndicator returns a boolean if a field has been set.

### GetIssuerExemption

`func (o *CardholderAuthenticationData) GetIssuerExemption() string`

GetIssuerExemption returns the IssuerExemption field if non-nil, zero value otherwise.

### GetIssuerExemptionOk

`func (o *CardholderAuthenticationData) GetIssuerExemptionOk() (*string, bool)`

GetIssuerExemptionOk returns a tuple with the IssuerExemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerExemption

`func (o *CardholderAuthenticationData) SetIssuerExemption(v string)`

SetIssuerExemption sets IssuerExemption field to given value.

### HasIssuerExemption

`func (o *CardholderAuthenticationData) HasIssuerExemption() bool`

HasIssuerExemption returns a boolean if a field has been set.

### GetRawCavvData

`func (o *CardholderAuthenticationData) GetRawCavvData() string`

GetRawCavvData returns the RawCavvData field if non-nil, zero value otherwise.

### GetRawCavvDataOk

`func (o *CardholderAuthenticationData) GetRawCavvDataOk() (*string, bool)`

GetRawCavvDataOk returns a tuple with the RawCavvData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCavvData

`func (o *CardholderAuthenticationData) SetRawCavvData(v string)`

SetRawCavvData sets RawCavvData field to given value.

### HasRawCavvData

`func (o *CardholderAuthenticationData) HasRawCavvData() bool`

HasRawCavvData returns a boolean if a field has been set.

### GetThreeDsMessageVersion

`func (o *CardholderAuthenticationData) GetThreeDsMessageVersion() string`

GetThreeDsMessageVersion returns the ThreeDsMessageVersion field if non-nil, zero value otherwise.

### GetThreeDsMessageVersionOk

`func (o *CardholderAuthenticationData) GetThreeDsMessageVersionOk() (*string, bool)`

GetThreeDsMessageVersionOk returns a tuple with the ThreeDsMessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsMessageVersion

`func (o *CardholderAuthenticationData) SetThreeDsMessageVersion(v string)`

SetThreeDsMessageVersion sets ThreeDsMessageVersion field to given value.

### HasThreeDsMessageVersion

`func (o *CardholderAuthenticationData) HasThreeDsMessageVersion() bool`

HasThreeDsMessageVersion returns a boolean if a field has been set.

### GetVerificationResult

`func (o *CardholderAuthenticationData) GetVerificationResult() string`

GetVerificationResult returns the VerificationResult field if non-nil, zero value otherwise.

### GetVerificationResultOk

`func (o *CardholderAuthenticationData) GetVerificationResultOk() (*string, bool)`

GetVerificationResultOk returns a tuple with the VerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationResult

`func (o *CardholderAuthenticationData) SetVerificationResult(v string)`

SetVerificationResult sets VerificationResult field to given value.

### HasVerificationResult

`func (o *CardholderAuthenticationData) HasVerificationResult() bool`

HasVerificationResult returns a boolean if a field has been set.

### GetVerificationValueCreatedBy

`func (o *CardholderAuthenticationData) GetVerificationValueCreatedBy() string`

GetVerificationValueCreatedBy returns the VerificationValueCreatedBy field if non-nil, zero value otherwise.

### GetVerificationValueCreatedByOk

`func (o *CardholderAuthenticationData) GetVerificationValueCreatedByOk() (*string, bool)`

GetVerificationValueCreatedByOk returns a tuple with the VerificationValueCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationValueCreatedBy

`func (o *CardholderAuthenticationData) SetVerificationValueCreatedBy(v string)`

SetVerificationValueCreatedBy sets VerificationValueCreatedBy field to given value.

### HasVerificationValueCreatedBy

`func (o *CardholderAuthenticationData) HasVerificationValueCreatedBy() bool`

HasVerificationValueCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


