# BulkIssuanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardAllocation** | **int32** | Number of cards in the order. | 
**CardProductToken** | **string** | Specifies the card product from which to create your cards. | 
**Expedite** | Pointer to **bool** | Set to &#x60;true&#x60; to request expedited processing of the order by your card fulfillment provider.  This expedited service is available for cards fulfilled by link:http://perfectplastic.com/[Perfect Plastic Printing, window&#x3D;\&quot;_blank\&quot;], link:http://www.idemia.com[IDEMIA, window&#x3D;\&quot;_blank\&quot;], and link:https://www.arroweye.com/[Arroweye Solutions, window&#x3D;\&quot;_blank\&quot;].  *NOTE:* Contact your Marqeta representative for information regarding the cost of expedited service. | [optional] [default to false]
**ExpirationOffset** | Pointer to [**ExpirationOffset**](ExpirationOffset.md) |  | [optional] 
**Fulfillment** | [**FulfillmentRequest**](FulfillmentRequest.md) |  | 
**NameLine1NumericPostfix** | Pointer to **bool** | If set to &#x60;true&#x60;, the unique numeric postfix appended to each card&#39;s token field is also appended to the card&#39;s &#x60;fulfillment.card_personalization.text.name_line_1.value&#x60; field. | [optional] [default to false]
**NameLine1RandomPostfix** | Pointer to **bool** | If set to &#x60;true&#x60;, the unique random postfix appended to each card&#39;s token field is also appended to the card&#39;s &#x60;fulfillment.card_personalization.text.name_line_1.value&#x60; field. | [optional] [default to false]
**Token** | **string** | If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | 
**UserAssociation** | Pointer to [**UserAssociation**](UserAssociation.md) |  | [optional] 

## Methods

### NewBulkIssuanceRequest

`func NewBulkIssuanceRequest(cardAllocation int32, cardProductToken string, fulfillment FulfillmentRequest, token string, ) *BulkIssuanceRequest`

NewBulkIssuanceRequest instantiates a new BulkIssuanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIssuanceRequestWithDefaults

`func NewBulkIssuanceRequestWithDefaults() *BulkIssuanceRequest`

NewBulkIssuanceRequestWithDefaults instantiates a new BulkIssuanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardAllocation

`func (o *BulkIssuanceRequest) GetCardAllocation() int32`

GetCardAllocation returns the CardAllocation field if non-nil, zero value otherwise.

### GetCardAllocationOk

`func (o *BulkIssuanceRequest) GetCardAllocationOk() (*int32, bool)`

GetCardAllocationOk returns a tuple with the CardAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAllocation

`func (o *BulkIssuanceRequest) SetCardAllocation(v int32)`

SetCardAllocation sets CardAllocation field to given value.


### GetCardProductToken

`func (o *BulkIssuanceRequest) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *BulkIssuanceRequest) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *BulkIssuanceRequest) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.


### GetExpedite

`func (o *BulkIssuanceRequest) GetExpedite() bool`

GetExpedite returns the Expedite field if non-nil, zero value otherwise.

### GetExpediteOk

`func (o *BulkIssuanceRequest) GetExpediteOk() (*bool, bool)`

GetExpediteOk returns a tuple with the Expedite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedite

`func (o *BulkIssuanceRequest) SetExpedite(v bool)`

SetExpedite sets Expedite field to given value.

### HasExpedite

`func (o *BulkIssuanceRequest) HasExpedite() bool`

HasExpedite returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *BulkIssuanceRequest) GetExpirationOffset() ExpirationOffset`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *BulkIssuanceRequest) GetExpirationOffsetOk() (*ExpirationOffset, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *BulkIssuanceRequest) SetExpirationOffset(v ExpirationOffset)`

SetExpirationOffset sets ExpirationOffset field to given value.

### HasExpirationOffset

`func (o *BulkIssuanceRequest) HasExpirationOffset() bool`

HasExpirationOffset returns a boolean if a field has been set.

### GetFulfillment

`func (o *BulkIssuanceRequest) GetFulfillment() FulfillmentRequest`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *BulkIssuanceRequest) GetFulfillmentOk() (*FulfillmentRequest, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *BulkIssuanceRequest) SetFulfillment(v FulfillmentRequest)`

SetFulfillment sets Fulfillment field to given value.


### GetNameLine1NumericPostfix

`func (o *BulkIssuanceRequest) GetNameLine1NumericPostfix() bool`

GetNameLine1NumericPostfix returns the NameLine1NumericPostfix field if non-nil, zero value otherwise.

### GetNameLine1NumericPostfixOk

`func (o *BulkIssuanceRequest) GetNameLine1NumericPostfixOk() (*bool, bool)`

GetNameLine1NumericPostfixOk returns a tuple with the NameLine1NumericPostfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLine1NumericPostfix

`func (o *BulkIssuanceRequest) SetNameLine1NumericPostfix(v bool)`

SetNameLine1NumericPostfix sets NameLine1NumericPostfix field to given value.

### HasNameLine1NumericPostfix

`func (o *BulkIssuanceRequest) HasNameLine1NumericPostfix() bool`

HasNameLine1NumericPostfix returns a boolean if a field has been set.

### GetNameLine1RandomPostfix

`func (o *BulkIssuanceRequest) GetNameLine1RandomPostfix() bool`

GetNameLine1RandomPostfix returns the NameLine1RandomPostfix field if non-nil, zero value otherwise.

### GetNameLine1RandomPostfixOk

`func (o *BulkIssuanceRequest) GetNameLine1RandomPostfixOk() (*bool, bool)`

GetNameLine1RandomPostfixOk returns a tuple with the NameLine1RandomPostfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLine1RandomPostfix

`func (o *BulkIssuanceRequest) SetNameLine1RandomPostfix(v bool)`

SetNameLine1RandomPostfix sets NameLine1RandomPostfix field to given value.

### HasNameLine1RandomPostfix

`func (o *BulkIssuanceRequest) HasNameLine1RandomPostfix() bool`

HasNameLine1RandomPostfix returns a boolean if a field has been set.

### GetToken

`func (o *BulkIssuanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BulkIssuanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BulkIssuanceRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserAssociation

`func (o *BulkIssuanceRequest) GetUserAssociation() UserAssociation`

GetUserAssociation returns the UserAssociation field if non-nil, zero value otherwise.

### GetUserAssociationOk

`func (o *BulkIssuanceRequest) GetUserAssociationOk() (*UserAssociation, bool)`

GetUserAssociationOk returns a tuple with the UserAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssociation

`func (o *BulkIssuanceRequest) SetUserAssociation(v UserAssociation)`

SetUserAssociation sets UserAssociation field to given value.

### HasUserAssociation

`func (o *BulkIssuanceRequest) HasUserAssociation() bool`

HasUserAssociation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


