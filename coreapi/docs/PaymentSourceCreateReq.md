# PaymentSourceCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Account number of the payment source. | 
**AccountToken** | **string** | Unique identifier of the credit account receiving the payment. | 
**Name** | **string** | Name of the individual or business who owns the payment source. | 
**Owner** | Pointer to **string** | Type of payment source owner. | [optional] 
**RoutingNumber** | **string** | Routing number of the payment source. | 
**SourceType** | **string** | Type of payment source. | 
**Token** | Pointer to **string** | Unique identifier of the payment source. | [optional] 
**UserToken** | **string** | Unique identifier of the user making the payment. | 
**VerificationNotes** | Pointer to **string** | Additional information on the verification. | [optional] 
**VerificationOverride** | **bool** | Whether to override the verification process. | 

## Methods

### NewPaymentSourceCreateReq

`func NewPaymentSourceCreateReq(accountNumber string, accountToken string, name string, routingNumber string, sourceType string, userToken string, verificationOverride bool, ) *PaymentSourceCreateReq`

NewPaymentSourceCreateReq instantiates a new PaymentSourceCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSourceCreateReqWithDefaults

`func NewPaymentSourceCreateReqWithDefaults() *PaymentSourceCreateReq`

NewPaymentSourceCreateReqWithDefaults instantiates a new PaymentSourceCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *PaymentSourceCreateReq) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentSourceCreateReq) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentSourceCreateReq) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountToken

`func (o *PaymentSourceCreateReq) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PaymentSourceCreateReq) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PaymentSourceCreateReq) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetName

`func (o *PaymentSourceCreateReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentSourceCreateReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentSourceCreateReq) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *PaymentSourceCreateReq) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PaymentSourceCreateReq) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PaymentSourceCreateReq) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PaymentSourceCreateReq) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *PaymentSourceCreateReq) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentSourceCreateReq) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentSourceCreateReq) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetSourceType

`func (o *PaymentSourceCreateReq) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PaymentSourceCreateReq) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PaymentSourceCreateReq) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetToken

`func (o *PaymentSourceCreateReq) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentSourceCreateReq) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentSourceCreateReq) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentSourceCreateReq) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *PaymentSourceCreateReq) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *PaymentSourceCreateReq) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *PaymentSourceCreateReq) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.


### GetVerificationNotes

`func (o *PaymentSourceCreateReq) GetVerificationNotes() string`

GetVerificationNotes returns the VerificationNotes field if non-nil, zero value otherwise.

### GetVerificationNotesOk

`func (o *PaymentSourceCreateReq) GetVerificationNotesOk() (*string, bool)`

GetVerificationNotesOk returns a tuple with the VerificationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNotes

`func (o *PaymentSourceCreateReq) SetVerificationNotes(v string)`

SetVerificationNotes sets VerificationNotes field to given value.

### HasVerificationNotes

`func (o *PaymentSourceCreateReq) HasVerificationNotes() bool`

HasVerificationNotes returns a boolean if a field has been set.

### GetVerificationOverride

`func (o *PaymentSourceCreateReq) GetVerificationOverride() bool`

GetVerificationOverride returns the VerificationOverride field if non-nil, zero value otherwise.

### GetVerificationOverrideOk

`func (o *PaymentSourceCreateReq) GetVerificationOverrideOk() (*bool, bool)`

GetVerificationOverrideOk returns a tuple with the VerificationOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationOverride

`func (o *PaymentSourceCreateReq) SetVerificationOverride(v bool)`

SetVerificationOverride sets VerificationOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


