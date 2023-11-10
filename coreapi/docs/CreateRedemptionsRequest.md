# CreateRedemptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | Amount of the reward redemption. | 
**Destination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional] 
**Note** | Pointer to **string** | A note explaining why the reward is being redeemed. | [optional] 
**ReceivingAccountToken** | Pointer to **string** | Unique identifier of the external account receiving the reward redemption. This token is equivalent to the &lt;&lt;/core-api/payment-sources, payment source&gt;&gt; token. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the reward redemption. | [optional] 
**Type** | [**RedemptionType**](RedemptionType.md) |  | 

## Methods

### NewCreateRedemptionsRequest

`func NewCreateRedemptionsRequest(amount float32, type_ RedemptionType, ) *CreateRedemptionsRequest`

NewCreateRedemptionsRequest instantiates a new CreateRedemptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRedemptionsRequestWithDefaults

`func NewCreateRedemptionsRequestWithDefaults() *CreateRedemptionsRequest`

NewCreateRedemptionsRequestWithDefaults instantiates a new CreateRedemptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateRedemptionsRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateRedemptionsRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateRedemptionsRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDestination

`func (o *CreateRedemptionsRequest) GetDestination() DestinationType`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CreateRedemptionsRequest) GetDestinationOk() (*DestinationType, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CreateRedemptionsRequest) SetDestination(v DestinationType)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *CreateRedemptionsRequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetNote

`func (o *CreateRedemptionsRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateRedemptionsRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateRedemptionsRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateRedemptionsRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReceivingAccountToken

`func (o *CreateRedemptionsRequest) GetReceivingAccountToken() string`

GetReceivingAccountToken returns the ReceivingAccountToken field if non-nil, zero value otherwise.

### GetReceivingAccountTokenOk

`func (o *CreateRedemptionsRequest) GetReceivingAccountTokenOk() (*string, bool)`

GetReceivingAccountTokenOk returns a tuple with the ReceivingAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingAccountToken

`func (o *CreateRedemptionsRequest) SetReceivingAccountToken(v string)`

SetReceivingAccountToken sets ReceivingAccountToken field to given value.

### HasReceivingAccountToken

`func (o *CreateRedemptionsRequest) HasReceivingAccountToken() bool`

HasReceivingAccountToken returns a boolean if a field has been set.

### GetToken

`func (o *CreateRedemptionsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateRedemptionsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateRedemptionsRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateRedemptionsRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *CreateRedemptionsRequest) GetType() RedemptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRedemptionsRequest) GetTypeOk() (*RedemptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRedemptionsRequest) SetType(v RedemptionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


