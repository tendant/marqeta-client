# AccountFundingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Account holder&#39;s first name. | 
**LastName** | Pointer to **string** | Account holder&#39;s last name. | [optional] 
**ReceiverName** | Pointer to **string** | Recipient&#39;s name. | [optional] 
**FundingSource** | **string** | Describes the source of the funding. | 
**ReceiverAccountType** | **string** | Identifies the account type used for the funding request. | 
**TransactionPurpose** | Pointer to **string** | Identifies the purpose of the transaction. | [optional] 
**TransactionType** | **string** | Describes the type of transaction. | 

## Methods

### NewAccountFundingRequest

`func NewAccountFundingRequest(firstName string, fundingSource string, receiverAccountType string, transactionType string, ) *AccountFundingRequest`

NewAccountFundingRequest instantiates a new AccountFundingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFundingRequestWithDefaults

`func NewAccountFundingRequestWithDefaults() *AccountFundingRequest`

NewAccountFundingRequestWithDefaults instantiates a new AccountFundingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *AccountFundingRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountFundingRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountFundingRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *AccountFundingRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountFundingRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountFundingRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountFundingRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetReceiverName

`func (o *AccountFundingRequest) GetReceiverName() string`

GetReceiverName returns the ReceiverName field if non-nil, zero value otherwise.

### GetReceiverNameOk

`func (o *AccountFundingRequest) GetReceiverNameOk() (*string, bool)`

GetReceiverNameOk returns a tuple with the ReceiverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverName

`func (o *AccountFundingRequest) SetReceiverName(v string)`

SetReceiverName sets ReceiverName field to given value.

### HasReceiverName

`func (o *AccountFundingRequest) HasReceiverName() bool`

HasReceiverName returns a boolean if a field has been set.

### GetFundingSource

`func (o *AccountFundingRequest) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *AccountFundingRequest) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *AccountFundingRequest) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.


### GetReceiverAccountType

`func (o *AccountFundingRequest) GetReceiverAccountType() string`

GetReceiverAccountType returns the ReceiverAccountType field if non-nil, zero value otherwise.

### GetReceiverAccountTypeOk

`func (o *AccountFundingRequest) GetReceiverAccountTypeOk() (*string, bool)`

GetReceiverAccountTypeOk returns a tuple with the ReceiverAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverAccountType

`func (o *AccountFundingRequest) SetReceiverAccountType(v string)`

SetReceiverAccountType sets ReceiverAccountType field to given value.


### GetTransactionPurpose

`func (o *AccountFundingRequest) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *AccountFundingRequest) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *AccountFundingRequest) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *AccountFundingRequest) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetTransactionType

`func (o *AccountFundingRequest) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *AccountFundingRequest) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *AccountFundingRequest) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


