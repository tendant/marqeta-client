# AccountUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aprs** | [**[]AprScheduleResponse**](AprScheduleResponse.md) | Contains one or more APRs associated with the type of balance on the credit account | 
**Fees** | Pointer to [**[]AccountFee**](AccountFee.md) | Contains one or more fees associated with the usage type. | [optional] 
**Type** | [**BalanceType**](BalanceType.md) |  | 

## Methods

### NewAccountUsageResponse

`func NewAccountUsageResponse(aprs []AprScheduleResponse, type_ BalanceType, ) *AccountUsageResponse`

NewAccountUsageResponse instantiates a new AccountUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUsageResponseWithDefaults

`func NewAccountUsageResponseWithDefaults() *AccountUsageResponse`

NewAccountUsageResponseWithDefaults instantiates a new AccountUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAprs

`func (o *AccountUsageResponse) GetAprs() []AprScheduleResponse`

GetAprs returns the Aprs field if non-nil, zero value otherwise.

### GetAprsOk

`func (o *AccountUsageResponse) GetAprsOk() (*[]AprScheduleResponse, bool)`

GetAprsOk returns a tuple with the Aprs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprs

`func (o *AccountUsageResponse) SetAprs(v []AprScheduleResponse)`

SetAprs sets Aprs field to given value.


### GetFees

`func (o *AccountUsageResponse) GetFees() []AccountFee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *AccountUsageResponse) GetFeesOk() (*[]AccountFee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *AccountUsageResponse) SetFees(v []AccountFee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *AccountUsageResponse) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetType

`func (o *AccountUsageResponse) GetType() BalanceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountUsageResponse) GetTypeOk() (*BalanceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountUsageResponse) SetType(v BalanceType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


