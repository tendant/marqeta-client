# AccountUsageUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aprs** | Pointer to [**[]AprScheduleUpdateReq**](AprScheduleUpdateReq.md) | Contains one or more annual percentage rates (APRs) associated with the credit account. | [optional] 
**Type** | [**BalanceType**](BalanceType.md) |  | 

## Methods

### NewAccountUsageUpdateReq

`func NewAccountUsageUpdateReq(type_ BalanceType, ) *AccountUsageUpdateReq`

NewAccountUsageUpdateReq instantiates a new AccountUsageUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUsageUpdateReqWithDefaults

`func NewAccountUsageUpdateReqWithDefaults() *AccountUsageUpdateReq`

NewAccountUsageUpdateReqWithDefaults instantiates a new AccountUsageUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAprs

`func (o *AccountUsageUpdateReq) GetAprs() []AprScheduleUpdateReq`

GetAprs returns the Aprs field if non-nil, zero value otherwise.

### GetAprsOk

`func (o *AccountUsageUpdateReq) GetAprsOk() (*[]AprScheduleUpdateReq, bool)`

GetAprsOk returns a tuple with the Aprs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprs

`func (o *AccountUsageUpdateReq) SetAprs(v []AprScheduleUpdateReq)`

SetAprs sets Aprs field to given value.

### HasAprs

`func (o *AccountUsageUpdateReq) HasAprs() bool`

HasAprs returns a boolean if a field has been set.

### GetType

`func (o *AccountUsageUpdateReq) GetType() BalanceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountUsageUpdateReq) GetTypeOk() (*BalanceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountUsageUpdateReq) SetType(v BalanceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


