# OrignalcreditRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**CardAcceptor** | Pointer to [**CardAcceptorModel**](CardAcceptorModel.md) |  | [optional] 
**CardToken** | **string** |  | 
**Mid** | **string** |  | 
**ScreeningScore** | Pointer to **string** |  | [optional] 
**SenderData** | Pointer to [**OriginalCreditSenderData**](OriginalCreditSenderData.md) |  | [optional] 
**TransactionPurpose** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Webhook** | Pointer to [**Webhook**](Webhook.md) |  | [optional] 

## Methods

### NewOrignalcreditRequestModel

`func NewOrignalcreditRequestModel(amount float32, cardToken string, mid string, type_ string, ) *OrignalcreditRequestModel`

NewOrignalcreditRequestModel instantiates a new OrignalcreditRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrignalcreditRequestModelWithDefaults

`func NewOrignalcreditRequestModelWithDefaults() *OrignalcreditRequestModel`

NewOrignalcreditRequestModelWithDefaults instantiates a new OrignalcreditRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OrignalcreditRequestModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrignalcreditRequestModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrignalcreditRequestModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCardAcceptor

`func (o *OrignalcreditRequestModel) GetCardAcceptor() CardAcceptorModel`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *OrignalcreditRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *OrignalcreditRequestModel) SetCardAcceptor(v CardAcceptorModel)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *OrignalcreditRequestModel) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.

### GetCardToken

`func (o *OrignalcreditRequestModel) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *OrignalcreditRequestModel) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *OrignalcreditRequestModel) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.


### GetMid

`func (o *OrignalcreditRequestModel) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *OrignalcreditRequestModel) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *OrignalcreditRequestModel) SetMid(v string)`

SetMid sets Mid field to given value.


### GetScreeningScore

`func (o *OrignalcreditRequestModel) GetScreeningScore() string`

GetScreeningScore returns the ScreeningScore field if non-nil, zero value otherwise.

### GetScreeningScoreOk

`func (o *OrignalcreditRequestModel) GetScreeningScoreOk() (*string, bool)`

GetScreeningScoreOk returns a tuple with the ScreeningScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningScore

`func (o *OrignalcreditRequestModel) SetScreeningScore(v string)`

SetScreeningScore sets ScreeningScore field to given value.

### HasScreeningScore

`func (o *OrignalcreditRequestModel) HasScreeningScore() bool`

HasScreeningScore returns a boolean if a field has been set.

### GetSenderData

`func (o *OrignalcreditRequestModel) GetSenderData() OriginalCreditSenderData`

GetSenderData returns the SenderData field if non-nil, zero value otherwise.

### GetSenderDataOk

`func (o *OrignalcreditRequestModel) GetSenderDataOk() (*OriginalCreditSenderData, bool)`

GetSenderDataOk returns a tuple with the SenderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderData

`func (o *OrignalcreditRequestModel) SetSenderData(v OriginalCreditSenderData)`

SetSenderData sets SenderData field to given value.

### HasSenderData

`func (o *OrignalcreditRequestModel) HasSenderData() bool`

HasSenderData returns a boolean if a field has been set.

### GetTransactionPurpose

`func (o *OrignalcreditRequestModel) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *OrignalcreditRequestModel) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *OrignalcreditRequestModel) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *OrignalcreditRequestModel) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetType

`func (o *OrignalcreditRequestModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrignalcreditRequestModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrignalcreditRequestModel) SetType(v string)`

SetType sets Type field to given value.


### GetWebhook

`func (o *OrignalcreditRequestModel) GetWebhook() Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *OrignalcreditRequestModel) GetWebhookOk() (*Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *OrignalcreditRequestModel) SetWebhook(v Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *OrignalcreditRequestModel) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


