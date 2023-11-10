# NetworkFraudView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionRiskScore** | Pointer to **int32** | Network-provided risk score for the transaction. A higher score indicates higher risk. Useful for making authorization decisions. | [optional] 
**TransactionRiskScoreReasonCode** | Pointer to **string** | _(Mastercard only)_ Unique code that describes the main driver of the &#x60;transaction_risk_score&#x60;. | [optional] 
**TransactionRiskScoreReasonDescription** | Pointer to **string** | _(Mastercard only)_ Description of the &#x60;transaction_risk_score_reason_code&#x60;. | [optional] 
**AccountRiskScore** | Pointer to **string** | _(Visa only)_ Account holder risk condition code evaluated by the card network. A higher score indicates a greater likelihood that the card number is compromised. | [optional] 
**AccountRiskScoreReasonCode** | Pointer to **string** | _(Visa only)_ Unique code that describes the main driver of the &#x60;account_risk_score&#x60;. | [optional] 

## Methods

### NewNetworkFraudView

`func NewNetworkFraudView() *NetworkFraudView`

NewNetworkFraudView instantiates a new NetworkFraudView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFraudViewWithDefaults

`func NewNetworkFraudViewWithDefaults() *NetworkFraudView`

NewNetworkFraudViewWithDefaults instantiates a new NetworkFraudView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionRiskScore

`func (o *NetworkFraudView) GetTransactionRiskScore() int32`

GetTransactionRiskScore returns the TransactionRiskScore field if non-nil, zero value otherwise.

### GetTransactionRiskScoreOk

`func (o *NetworkFraudView) GetTransactionRiskScoreOk() (*int32, bool)`

GetTransactionRiskScoreOk returns a tuple with the TransactionRiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRiskScore

`func (o *NetworkFraudView) SetTransactionRiskScore(v int32)`

SetTransactionRiskScore sets TransactionRiskScore field to given value.

### HasTransactionRiskScore

`func (o *NetworkFraudView) HasTransactionRiskScore() bool`

HasTransactionRiskScore returns a boolean if a field has been set.

### GetTransactionRiskScoreReasonCode

`func (o *NetworkFraudView) GetTransactionRiskScoreReasonCode() string`

GetTransactionRiskScoreReasonCode returns the TransactionRiskScoreReasonCode field if non-nil, zero value otherwise.

### GetTransactionRiskScoreReasonCodeOk

`func (o *NetworkFraudView) GetTransactionRiskScoreReasonCodeOk() (*string, bool)`

GetTransactionRiskScoreReasonCodeOk returns a tuple with the TransactionRiskScoreReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRiskScoreReasonCode

`func (o *NetworkFraudView) SetTransactionRiskScoreReasonCode(v string)`

SetTransactionRiskScoreReasonCode sets TransactionRiskScoreReasonCode field to given value.

### HasTransactionRiskScoreReasonCode

`func (o *NetworkFraudView) HasTransactionRiskScoreReasonCode() bool`

HasTransactionRiskScoreReasonCode returns a boolean if a field has been set.

### GetTransactionRiskScoreReasonDescription

`func (o *NetworkFraudView) GetTransactionRiskScoreReasonDescription() string`

GetTransactionRiskScoreReasonDescription returns the TransactionRiskScoreReasonDescription field if non-nil, zero value otherwise.

### GetTransactionRiskScoreReasonDescriptionOk

`func (o *NetworkFraudView) GetTransactionRiskScoreReasonDescriptionOk() (*string, bool)`

GetTransactionRiskScoreReasonDescriptionOk returns a tuple with the TransactionRiskScoreReasonDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionRiskScoreReasonDescription

`func (o *NetworkFraudView) SetTransactionRiskScoreReasonDescription(v string)`

SetTransactionRiskScoreReasonDescription sets TransactionRiskScoreReasonDescription field to given value.

### HasTransactionRiskScoreReasonDescription

`func (o *NetworkFraudView) HasTransactionRiskScoreReasonDescription() bool`

HasTransactionRiskScoreReasonDescription returns a boolean if a field has been set.

### GetAccountRiskScore

`func (o *NetworkFraudView) GetAccountRiskScore() string`

GetAccountRiskScore returns the AccountRiskScore field if non-nil, zero value otherwise.

### GetAccountRiskScoreOk

`func (o *NetworkFraudView) GetAccountRiskScoreOk() (*string, bool)`

GetAccountRiskScoreOk returns a tuple with the AccountRiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRiskScore

`func (o *NetworkFraudView) SetAccountRiskScore(v string)`

SetAccountRiskScore sets AccountRiskScore field to given value.

### HasAccountRiskScore

`func (o *NetworkFraudView) HasAccountRiskScore() bool`

HasAccountRiskScore returns a boolean if a field has been set.

### GetAccountRiskScoreReasonCode

`func (o *NetworkFraudView) GetAccountRiskScoreReasonCode() string`

GetAccountRiskScoreReasonCode returns the AccountRiskScoreReasonCode field if non-nil, zero value otherwise.

### GetAccountRiskScoreReasonCodeOk

`func (o *NetworkFraudView) GetAccountRiskScoreReasonCodeOk() (*string, bool)`

GetAccountRiskScoreReasonCodeOk returns a tuple with the AccountRiskScoreReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRiskScoreReasonCode

`func (o *NetworkFraudView) SetAccountRiskScoreReasonCode(v string)`

SetAccountRiskScoreReasonCode sets AccountRiskScoreReasonCode field to given value.

### HasAccountRiskScoreReasonCode

`func (o *NetworkFraudView) HasAccountRiskScoreReasonCode() bool`

HasAccountRiskScoreReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


