# PaymentSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSuffix** | Pointer to **string** | Last four digits of the payment source account number. | [optional] 
**AccountToken** | Pointer to **string** | Unique identifier of the credit account receiving the payment. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the payment source was created on Marqeta&#39;s credit platform, in UTC. | [optional] 
**LastModifiedTime** | Pointer to **time.Time** | Date and time when the payment source was last updated on Marqeta&#39;s credit platform, in UTC. | [optional] 
**Name** | Pointer to **string** | Name of the individual or business who owns the payment source. | [optional] 
**Owner** | Pointer to **string** | Type of payment source owner. | [optional] 
**RoutingNumber** | Pointer to **string** | Routing number of the payment source. | [optional] 
**SourceType** | Pointer to **string** | Type of payment source. | [optional] 
**Status** | Pointer to [**PaymentSourceStatusEnum**](PaymentSourceStatusEnum.md) |  | [optional] 
**Token** | Pointer to **string** | Unique identifier of the payment source. | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user making the payment. | [optional] 
**VerificationNotes** | Pointer to **string** | Additional information on the verification (for example, an external verification identifier that&#39;s outside Marqeta&#39;s credit platform). | [optional] 
**VerificationStatus** | Pointer to **string** | Status of the verification for the payment source. | [optional] 

## Methods

### NewPaymentSourceResponse

`func NewPaymentSourceResponse() *PaymentSourceResponse`

NewPaymentSourceResponse instantiates a new PaymentSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSourceResponseWithDefaults

`func NewPaymentSourceResponseWithDefaults() *PaymentSourceResponse`

NewPaymentSourceResponseWithDefaults instantiates a new PaymentSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSuffix

`func (o *PaymentSourceResponse) GetAccountSuffix() string`

GetAccountSuffix returns the AccountSuffix field if non-nil, zero value otherwise.

### GetAccountSuffixOk

`func (o *PaymentSourceResponse) GetAccountSuffixOk() (*string, bool)`

GetAccountSuffixOk returns a tuple with the AccountSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSuffix

`func (o *PaymentSourceResponse) SetAccountSuffix(v string)`

SetAccountSuffix sets AccountSuffix field to given value.

### HasAccountSuffix

`func (o *PaymentSourceResponse) HasAccountSuffix() bool`

HasAccountSuffix returns a boolean if a field has been set.

### GetAccountToken

`func (o *PaymentSourceResponse) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *PaymentSourceResponse) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *PaymentSourceResponse) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *PaymentSourceResponse) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetCreatedTime

`func (o *PaymentSourceResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *PaymentSourceResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *PaymentSourceResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *PaymentSourceResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *PaymentSourceResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *PaymentSourceResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *PaymentSourceResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *PaymentSourceResponse) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetName

`func (o *PaymentSourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentSourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentSourceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentSourceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *PaymentSourceResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PaymentSourceResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PaymentSourceResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PaymentSourceResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRoutingNumber

`func (o *PaymentSourceResponse) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *PaymentSourceResponse) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *PaymentSourceResponse) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *PaymentSourceResponse) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.

### GetSourceType

`func (o *PaymentSourceResponse) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PaymentSourceResponse) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PaymentSourceResponse) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PaymentSourceResponse) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentSourceResponse) GetStatus() PaymentSourceStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentSourceResponse) GetStatusOk() (*PaymentSourceStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentSourceResponse) SetStatus(v PaymentSourceStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentSourceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *PaymentSourceResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaymentSourceResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaymentSourceResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaymentSourceResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserToken

`func (o *PaymentSourceResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *PaymentSourceResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *PaymentSourceResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *PaymentSourceResponse) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetVerificationNotes

`func (o *PaymentSourceResponse) GetVerificationNotes() string`

GetVerificationNotes returns the VerificationNotes field if non-nil, zero value otherwise.

### GetVerificationNotesOk

`func (o *PaymentSourceResponse) GetVerificationNotesOk() (*string, bool)`

GetVerificationNotesOk returns a tuple with the VerificationNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationNotes

`func (o *PaymentSourceResponse) SetVerificationNotes(v string)`

SetVerificationNotes sets VerificationNotes field to given value.

### HasVerificationNotes

`func (o *PaymentSourceResponse) HasVerificationNotes() bool`

HasVerificationNotes returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *PaymentSourceResponse) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PaymentSourceResponse) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PaymentSourceResponse) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *PaymentSourceResponse) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


