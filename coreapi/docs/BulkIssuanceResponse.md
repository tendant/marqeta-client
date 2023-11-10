# BulkIssuanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardAllocation** | **int32** | Number of cards in the order. | 
**CardFulfillmentTime** | Pointer to **time.Time** | Date and time when the bulk card order was fulfilled, in UTC.  This field is included if your bulk card order has been processed. | [optional] 
**CardProductToken** | **string** | Specifies the card product from which the cards are created. | 
**CardsProcessed** | Pointer to **int32** | Number of cards processed in the bulk card order.  This field is returned if it exists in the resource. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the resource was created, in UTC.  This field is returned if it exists in the resource. | [optional] 
**Expedite** | Pointer to **bool** | Indicates if expedited processing of this bulk card order was requested.  This field is returned if it exists in the resource. | [optional] [default to false]
**ExpirationOffset** | Pointer to [**ExpirationOffset**](ExpirationOffset.md) |  | [optional] 
**Fulfillment** | [**FulfillmentResponse**](FulfillmentResponse.md) |  | 
**NameLine1EndIndex** | Pointer to **int32** | This field is included if your bulk card order has been processed.  You can use the &#x60;name_line1_start_index&#x60; and &#x60;name_line1_end_index&#x60; fields to identify the cards and users associated with the order. For example, if the start index is \&quot;1\&quot; and the end index is \&quot;3\&quot;, the card tokens are \&quot;card-1\&quot;, \&quot;card-2\&quot;, and \&quot;card-3\&quot;, and the user tokens are \&quot;user-1\&quot;, \&quot;user-2\&quot;, and \&quot;user-3\&quot;.  See &lt;&lt;/core-api/bulk-card-orders#create_bulk_card_order, Create bulk card order&gt;&gt; for more information about the automatic generation and naming of cards and users. | [optional] 
**NameLine1StartIndex** | Pointer to **int32** | This field is included if your bulk card order has been processed.  You can use the &#x60;name_line1_start_index&#x60; and &#x60;name_line1_end_index&#x60; fields to identify the cards and users associated with the order. For example, if the start index is \&quot;1\&quot; and the end index is \&quot;3\&quot;, the card tokens are \&quot;card-1\&quot;, \&quot;card-2\&quot;, and \&quot;card-3\&quot;, and the user tokens are \&quot;user-1\&quot;, \&quot;user-2\&quot;, and \&quot;user-3\&quot;.  See &lt;&lt;/core-api/bulk-card-orders#create_bulk_card_order, Create bulk card order&gt;&gt; for more information about the automatic generation and naming of cards and users. | [optional] 
**NameLine1NumericPostfix** | Pointer to **bool** | If set to &#x60;true&#x60;, the unique numeric postfix appended to each card&#39;s token field is also appended to the card&#39;s &#x60;fulfillment.card_personalization.text.name_line_1.value&#x60; field. | [optional] [default to false]
**NameLine1RandomPostfix** | Pointer to **bool** | If set to &#x60;true&#x60;, the unique random postfix appended to each card&#39;s token field is also appended to the card&#39;s &#x60;fulfillment.card_personalization.text.name_line_1.value&#x60; field.  This field is returned if it exists in the resource. | [optional] [default to false]
**ProviderShipDate** | Pointer to **time.Time** | Date and time when the card provider shipped the bulk card order, in UTC.  This field is included if your bulk card order has shipped. | [optional] 
**ProviderShippingMethod** | Pointer to **string** | Shipping method used by the card provider. &#x60;United_Postal_Service&#x60;, for example.  This field is included if your bulk card order has shipped. | [optional] 
**ProviderTrackingNumber** | Pointer to **string** | A tracking number is included only if your card provider is Arroweye Solutions.  This field is included if your bulk card order has shipped. | [optional] 
**Token** | **string** | Unique identifier of the bulk card order. | 
**UserAssociation** | Pointer to [**UserAssociation**](UserAssociation.md) |  | [optional] 

## Methods

### NewBulkIssuanceResponse

`func NewBulkIssuanceResponse(cardAllocation int32, cardProductToken string, fulfillment FulfillmentResponse, token string, ) *BulkIssuanceResponse`

NewBulkIssuanceResponse instantiates a new BulkIssuanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkIssuanceResponseWithDefaults

`func NewBulkIssuanceResponseWithDefaults() *BulkIssuanceResponse`

NewBulkIssuanceResponseWithDefaults instantiates a new BulkIssuanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardAllocation

`func (o *BulkIssuanceResponse) GetCardAllocation() int32`

GetCardAllocation returns the CardAllocation field if non-nil, zero value otherwise.

### GetCardAllocationOk

`func (o *BulkIssuanceResponse) GetCardAllocationOk() (*int32, bool)`

GetCardAllocationOk returns a tuple with the CardAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAllocation

`func (o *BulkIssuanceResponse) SetCardAllocation(v int32)`

SetCardAllocation sets CardAllocation field to given value.


### GetCardFulfillmentTime

`func (o *BulkIssuanceResponse) GetCardFulfillmentTime() time.Time`

GetCardFulfillmentTime returns the CardFulfillmentTime field if non-nil, zero value otherwise.

### GetCardFulfillmentTimeOk

`func (o *BulkIssuanceResponse) GetCardFulfillmentTimeOk() (*time.Time, bool)`

GetCardFulfillmentTimeOk returns a tuple with the CardFulfillmentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardFulfillmentTime

`func (o *BulkIssuanceResponse) SetCardFulfillmentTime(v time.Time)`

SetCardFulfillmentTime sets CardFulfillmentTime field to given value.

### HasCardFulfillmentTime

`func (o *BulkIssuanceResponse) HasCardFulfillmentTime() bool`

HasCardFulfillmentTime returns a boolean if a field has been set.

### GetCardProductToken

`func (o *BulkIssuanceResponse) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *BulkIssuanceResponse) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *BulkIssuanceResponse) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.


### GetCardsProcessed

`func (o *BulkIssuanceResponse) GetCardsProcessed() int32`

GetCardsProcessed returns the CardsProcessed field if non-nil, zero value otherwise.

### GetCardsProcessedOk

`func (o *BulkIssuanceResponse) GetCardsProcessedOk() (*int32, bool)`

GetCardsProcessedOk returns a tuple with the CardsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardsProcessed

`func (o *BulkIssuanceResponse) SetCardsProcessed(v int32)`

SetCardsProcessed sets CardsProcessed field to given value.

### HasCardsProcessed

`func (o *BulkIssuanceResponse) HasCardsProcessed() bool`

HasCardsProcessed returns a boolean if a field has been set.

### GetCreatedTime

`func (o *BulkIssuanceResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *BulkIssuanceResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *BulkIssuanceResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *BulkIssuanceResponse) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetExpedite

`func (o *BulkIssuanceResponse) GetExpedite() bool`

GetExpedite returns the Expedite field if non-nil, zero value otherwise.

### GetExpediteOk

`func (o *BulkIssuanceResponse) GetExpediteOk() (*bool, bool)`

GetExpediteOk returns a tuple with the Expedite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedite

`func (o *BulkIssuanceResponse) SetExpedite(v bool)`

SetExpedite sets Expedite field to given value.

### HasExpedite

`func (o *BulkIssuanceResponse) HasExpedite() bool`

HasExpedite returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *BulkIssuanceResponse) GetExpirationOffset() ExpirationOffset`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *BulkIssuanceResponse) GetExpirationOffsetOk() (*ExpirationOffset, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *BulkIssuanceResponse) SetExpirationOffset(v ExpirationOffset)`

SetExpirationOffset sets ExpirationOffset field to given value.

### HasExpirationOffset

`func (o *BulkIssuanceResponse) HasExpirationOffset() bool`

HasExpirationOffset returns a boolean if a field has been set.

### GetFulfillment

`func (o *BulkIssuanceResponse) GetFulfillment() FulfillmentResponse`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *BulkIssuanceResponse) GetFulfillmentOk() (*FulfillmentResponse, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *BulkIssuanceResponse) SetFulfillment(v FulfillmentResponse)`

SetFulfillment sets Fulfillment field to given value.


### GetNameLine1EndIndex

`func (o *BulkIssuanceResponse) GetNameLine1EndIndex() int32`

GetNameLine1EndIndex returns the NameLine1EndIndex field if non-nil, zero value otherwise.

### GetNameLine1EndIndexOk

`func (o *BulkIssuanceResponse) GetNameLine1EndIndexOk() (*int32, bool)`

GetNameLine1EndIndexOk returns a tuple with the NameLine1EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLine1EndIndex

`func (o *BulkIssuanceResponse) SetNameLine1EndIndex(v int32)`

SetNameLine1EndIndex sets NameLine1EndIndex field to given value.

### HasNameLine1EndIndex

`func (o *BulkIssuanceResponse) HasNameLine1EndIndex() bool`

HasNameLine1EndIndex returns a boolean if a field has been set.

### GetNameLine1StartIndex

`func (o *BulkIssuanceResponse) GetNameLine1StartIndex() int32`

GetNameLine1StartIndex returns the NameLine1StartIndex field if non-nil, zero value otherwise.

### GetNameLine1StartIndexOk

`func (o *BulkIssuanceResponse) GetNameLine1StartIndexOk() (*int32, bool)`

GetNameLine1StartIndexOk returns a tuple with the NameLine1StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLine1StartIndex

`func (o *BulkIssuanceResponse) SetNameLine1StartIndex(v int32)`

SetNameLine1StartIndex sets NameLine1StartIndex field to given value.

### HasNameLine1StartIndex

`func (o *BulkIssuanceResponse) HasNameLine1StartIndex() bool`

HasNameLine1StartIndex returns a boolean if a field has been set.

### GetNameLine1NumericPostfix

`func (o *BulkIssuanceResponse) GetNameLine1NumericPostfix() bool`

GetNameLine1NumericPostfix returns the NameLine1NumericPostfix field if non-nil, zero value otherwise.

### GetNameLine1NumericPostfixOk

`func (o *BulkIssuanceResponse) GetNameLine1NumericPostfixOk() (*bool, bool)`

GetNameLine1NumericPostfixOk returns a tuple with the NameLine1NumericPostfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLine1NumericPostfix

`func (o *BulkIssuanceResponse) SetNameLine1NumericPostfix(v bool)`

SetNameLine1NumericPostfix sets NameLine1NumericPostfix field to given value.

### HasNameLine1NumericPostfix

`func (o *BulkIssuanceResponse) HasNameLine1NumericPostfix() bool`

HasNameLine1NumericPostfix returns a boolean if a field has been set.

### GetNameLine1RandomPostfix

`func (o *BulkIssuanceResponse) GetNameLine1RandomPostfix() bool`

GetNameLine1RandomPostfix returns the NameLine1RandomPostfix field if non-nil, zero value otherwise.

### GetNameLine1RandomPostfixOk

`func (o *BulkIssuanceResponse) GetNameLine1RandomPostfixOk() (*bool, bool)`

GetNameLine1RandomPostfixOk returns a tuple with the NameLine1RandomPostfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLine1RandomPostfix

`func (o *BulkIssuanceResponse) SetNameLine1RandomPostfix(v bool)`

SetNameLine1RandomPostfix sets NameLine1RandomPostfix field to given value.

### HasNameLine1RandomPostfix

`func (o *BulkIssuanceResponse) HasNameLine1RandomPostfix() bool`

HasNameLine1RandomPostfix returns a boolean if a field has been set.

### GetProviderShipDate

`func (o *BulkIssuanceResponse) GetProviderShipDate() time.Time`

GetProviderShipDate returns the ProviderShipDate field if non-nil, zero value otherwise.

### GetProviderShipDateOk

`func (o *BulkIssuanceResponse) GetProviderShipDateOk() (*time.Time, bool)`

GetProviderShipDateOk returns a tuple with the ProviderShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderShipDate

`func (o *BulkIssuanceResponse) SetProviderShipDate(v time.Time)`

SetProviderShipDate sets ProviderShipDate field to given value.

### HasProviderShipDate

`func (o *BulkIssuanceResponse) HasProviderShipDate() bool`

HasProviderShipDate returns a boolean if a field has been set.

### GetProviderShippingMethod

`func (o *BulkIssuanceResponse) GetProviderShippingMethod() string`

GetProviderShippingMethod returns the ProviderShippingMethod field if non-nil, zero value otherwise.

### GetProviderShippingMethodOk

`func (o *BulkIssuanceResponse) GetProviderShippingMethodOk() (*string, bool)`

GetProviderShippingMethodOk returns a tuple with the ProviderShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderShippingMethod

`func (o *BulkIssuanceResponse) SetProviderShippingMethod(v string)`

SetProviderShippingMethod sets ProviderShippingMethod field to given value.

### HasProviderShippingMethod

`func (o *BulkIssuanceResponse) HasProviderShippingMethod() bool`

HasProviderShippingMethod returns a boolean if a field has been set.

### GetProviderTrackingNumber

`func (o *BulkIssuanceResponse) GetProviderTrackingNumber() string`

GetProviderTrackingNumber returns the ProviderTrackingNumber field if non-nil, zero value otherwise.

### GetProviderTrackingNumberOk

`func (o *BulkIssuanceResponse) GetProviderTrackingNumberOk() (*string, bool)`

GetProviderTrackingNumberOk returns a tuple with the ProviderTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderTrackingNumber

`func (o *BulkIssuanceResponse) SetProviderTrackingNumber(v string)`

SetProviderTrackingNumber sets ProviderTrackingNumber field to given value.

### HasProviderTrackingNumber

`func (o *BulkIssuanceResponse) HasProviderTrackingNumber() bool`

HasProviderTrackingNumber returns a boolean if a field has been set.

### GetToken

`func (o *BulkIssuanceResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BulkIssuanceResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BulkIssuanceResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserAssociation

`func (o *BulkIssuanceResponse) GetUserAssociation() UserAssociation`

GetUserAssociation returns the UserAssociation field if non-nil, zero value otherwise.

### GetUserAssociationOk

`func (o *BulkIssuanceResponse) GetUserAssociationOk() (*UserAssociation, bool)`

GetUserAssociationOk returns a tuple with the UserAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssociation

`func (o *BulkIssuanceResponse) SetUserAssociation(v UserAssociation)`

SetUserAssociation sets UserAssociation field to given value.

### HasUserAssociation

`func (o *BulkIssuanceResponse) HasUserAssociation() bool`

HasUserAssociation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


