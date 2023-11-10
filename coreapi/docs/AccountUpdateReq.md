# AccountUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**AccountConfigUpdateReq**](AccountConfigUpdateReq.md) |  | [optional] 
**CreditLimit** | Pointer to [**AccountUpdateReqCreditLimit**](AccountUpdateReqCreditLimit.md) |  | [optional] 
**LatestStatementCycleType** | Pointer to [**CycleType**](CycleType.md) |  | [optional] 
**Usages** | Pointer to [**[]AccountUsageUpdateReq**](AccountUsageUpdateReq.md) | Contains one or more &#x60;usages&#x60; objects that contain information on how a credit account is used and what types of balances are permitted on the account.  You can pass only one &#x60;usages&#x60; object per &#x60;usages.type&#x60;. | [optional] 

## Methods

### NewAccountUpdateReq

`func NewAccountUpdateReq() *AccountUpdateReq`

NewAccountUpdateReq instantiates a new AccountUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateReqWithDefaults

`func NewAccountUpdateReqWithDefaults() *AccountUpdateReq`

NewAccountUpdateReqWithDefaults instantiates a new AccountUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AccountUpdateReq) GetConfig() AccountConfigUpdateReq`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AccountUpdateReq) GetConfigOk() (*AccountConfigUpdateReq, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AccountUpdateReq) SetConfig(v AccountConfigUpdateReq)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AccountUpdateReq) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreditLimit

`func (o *AccountUpdateReq) GetCreditLimit() AccountUpdateReqCreditLimit`

GetCreditLimit returns the CreditLimit field if non-nil, zero value otherwise.

### GetCreditLimitOk

`func (o *AccountUpdateReq) GetCreditLimitOk() (*AccountUpdateReqCreditLimit, bool)`

GetCreditLimitOk returns a tuple with the CreditLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLimit

`func (o *AccountUpdateReq) SetCreditLimit(v AccountUpdateReqCreditLimit)`

SetCreditLimit sets CreditLimit field to given value.

### HasCreditLimit

`func (o *AccountUpdateReq) HasCreditLimit() bool`

HasCreditLimit returns a boolean if a field has been set.

### GetLatestStatementCycleType

`func (o *AccountUpdateReq) GetLatestStatementCycleType() CycleType`

GetLatestStatementCycleType returns the LatestStatementCycleType field if non-nil, zero value otherwise.

### GetLatestStatementCycleTypeOk

`func (o *AccountUpdateReq) GetLatestStatementCycleTypeOk() (*CycleType, bool)`

GetLatestStatementCycleTypeOk returns a tuple with the LatestStatementCycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestStatementCycleType

`func (o *AccountUpdateReq) SetLatestStatementCycleType(v CycleType)`

SetLatestStatementCycleType sets LatestStatementCycleType field to given value.

### HasLatestStatementCycleType

`func (o *AccountUpdateReq) HasLatestStatementCycleType() bool`

HasLatestStatementCycleType returns a boolean if a field has been set.

### GetUsages

`func (o *AccountUpdateReq) GetUsages() []AccountUsageUpdateReq`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *AccountUpdateReq) GetUsagesOk() (*[]AccountUsageUpdateReq, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *AccountUpdateReq) SetUsages(v []AccountUsageUpdateReq)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *AccountUpdateReq) HasUsages() bool`

HasUsages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


