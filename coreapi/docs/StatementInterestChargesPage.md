# StatementInterestChargesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | **string** | Unique identifier of the credit account on which the statement interest charge is generated. | 
**Data** | [**[]StatementInterestCharge**](StatementInterestCharge.md) | Contains one or more interest charges on a statement. | 
**StatementSummaryToken** | **string** | Unique identifier of the statement summary. | 

## Methods

### NewStatementInterestChargesPage

`func NewStatementInterestChargesPage(accountToken string, data []StatementInterestCharge, statementSummaryToken string, ) *StatementInterestChargesPage`

NewStatementInterestChargesPage instantiates a new StatementInterestChargesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementInterestChargesPageWithDefaults

`func NewStatementInterestChargesPageWithDefaults() *StatementInterestChargesPage`

NewStatementInterestChargesPageWithDefaults instantiates a new StatementInterestChargesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *StatementInterestChargesPage) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *StatementInterestChargesPage) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *StatementInterestChargesPage) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.


### GetData

`func (o *StatementInterestChargesPage) GetData() []StatementInterestCharge`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StatementInterestChargesPage) GetDataOk() (*[]StatementInterestCharge, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StatementInterestChargesPage) SetData(v []StatementInterestCharge)`

SetData sets Data field to given value.


### GetStatementSummaryToken

`func (o *StatementInterestChargesPage) GetStatementSummaryToken() string`

GetStatementSummaryToken returns the StatementSummaryToken field if non-nil, zero value otherwise.

### GetStatementSummaryTokenOk

`func (o *StatementInterestChargesPage) GetStatementSummaryTokenOk() (*string, bool)`

GetStatementSummaryTokenOk returns a tuple with the StatementSummaryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementSummaryToken

`func (o *StatementInterestChargesPage) SetStatementSummaryToken(v string)`

SetStatementSummaryToken sets StatementSummaryToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


