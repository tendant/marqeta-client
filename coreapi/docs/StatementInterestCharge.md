# StatementInterestCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount of interest calculated for the billing period. | [optional] 
**AprType** | Pointer to **string** | Type of APR. | [optional] 
**AprValue** | Pointer to **float32** | Annual percentage rate. | [optional] 
**BalanceSubjectToInterestRate** | Pointer to **float32** | Average daily balance used to calculate interest. | [optional] 
**BalanceType** | Pointer to **string** | Type of balance.  * &#x60;PURCHASE&#x60; - The balance on purchases. | [optional] 

## Methods

### NewStatementInterestCharge

`func NewStatementInterestCharge() *StatementInterestCharge`

NewStatementInterestCharge instantiates a new StatementInterestCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementInterestChargeWithDefaults

`func NewStatementInterestChargeWithDefaults() *StatementInterestCharge`

NewStatementInterestChargeWithDefaults instantiates a new StatementInterestCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StatementInterestCharge) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StatementInterestCharge) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StatementInterestCharge) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *StatementInterestCharge) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAprType

`func (o *StatementInterestCharge) GetAprType() string`

GetAprType returns the AprType field if non-nil, zero value otherwise.

### GetAprTypeOk

`func (o *StatementInterestCharge) GetAprTypeOk() (*string, bool)`

GetAprTypeOk returns a tuple with the AprType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprType

`func (o *StatementInterestCharge) SetAprType(v string)`

SetAprType sets AprType field to given value.

### HasAprType

`func (o *StatementInterestCharge) HasAprType() bool`

HasAprType returns a boolean if a field has been set.

### GetAprValue

`func (o *StatementInterestCharge) GetAprValue() float32`

GetAprValue returns the AprValue field if non-nil, zero value otherwise.

### GetAprValueOk

`func (o *StatementInterestCharge) GetAprValueOk() (*float32, bool)`

GetAprValueOk returns a tuple with the AprValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprValue

`func (o *StatementInterestCharge) SetAprValue(v float32)`

SetAprValue sets AprValue field to given value.

### HasAprValue

`func (o *StatementInterestCharge) HasAprValue() bool`

HasAprValue returns a boolean if a field has been set.

### GetBalanceSubjectToInterestRate

`func (o *StatementInterestCharge) GetBalanceSubjectToInterestRate() float32`

GetBalanceSubjectToInterestRate returns the BalanceSubjectToInterestRate field if non-nil, zero value otherwise.

### GetBalanceSubjectToInterestRateOk

`func (o *StatementInterestCharge) GetBalanceSubjectToInterestRateOk() (*float32, bool)`

GetBalanceSubjectToInterestRateOk returns a tuple with the BalanceSubjectToInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSubjectToInterestRate

`func (o *StatementInterestCharge) SetBalanceSubjectToInterestRate(v float32)`

SetBalanceSubjectToInterestRate sets BalanceSubjectToInterestRate field to given value.

### HasBalanceSubjectToInterestRate

`func (o *StatementInterestCharge) HasBalanceSubjectToInterestRate() bool`

HasBalanceSubjectToInterestRate returns a boolean if a field has been set.

### GetBalanceType

`func (o *StatementInterestCharge) GetBalanceType() string`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *StatementInterestCharge) GetBalanceTypeOk() (*string, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *StatementInterestCharge) SetBalanceType(v string)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *StatementInterestCharge) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


