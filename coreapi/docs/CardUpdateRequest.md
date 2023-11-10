# CardUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expedite** | Pointer to **bool** | Set to &#x60;true&#x60; to request expedited processing of the card by your card fulfillment provider.  This expedited service is available for cards fulfilled by link:http://perfectplastic.com/[Perfect Plastic Printing, window&#x3D;\&quot;_blank\&quot;], link:http://www.idemia.com[IDEMIA, window&#x3D;\&quot;_blank\&quot;], and link:https://www.arroweye.com/[Arroweye Solutions, window&#x3D;\&quot;_blank\&quot;].  *NOTE:* Contact your Marqeta representative for information regarding the cost of expedited service. | [optional] [default to false]
**Fulfillment** | Pointer to [**CardFulfillmentRequest**](CardFulfillmentRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Associates customer-provided metadata with the card. | [optional] 
**Token** | **string** | Unique identifier of the card you want to update. | 
**UserToken** | Pointer to **string** | Specifies the user you want to associate with the card. | [optional] 

## Methods

### NewCardUpdateRequest

`func NewCardUpdateRequest(token string, ) *CardUpdateRequest`

NewCardUpdateRequest instantiates a new CardUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardUpdateRequestWithDefaults

`func NewCardUpdateRequestWithDefaults() *CardUpdateRequest`

NewCardUpdateRequestWithDefaults instantiates a new CardUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpedite

`func (o *CardUpdateRequest) GetExpedite() bool`

GetExpedite returns the Expedite field if non-nil, zero value otherwise.

### GetExpediteOk

`func (o *CardUpdateRequest) GetExpediteOk() (*bool, bool)`

GetExpediteOk returns a tuple with the Expedite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedite

`func (o *CardUpdateRequest) SetExpedite(v bool)`

SetExpedite sets Expedite field to given value.

### HasExpedite

`func (o *CardUpdateRequest) HasExpedite() bool`

HasExpedite returns a boolean if a field has been set.

### GetFulfillment

`func (o *CardUpdateRequest) GetFulfillment() CardFulfillmentRequest`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *CardUpdateRequest) GetFulfillmentOk() (*CardFulfillmentRequest, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *CardUpdateRequest) SetFulfillment(v CardFulfillmentRequest)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *CardUpdateRequest) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetMetadata

`func (o *CardUpdateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CardUpdateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CardUpdateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CardUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetToken

`func (o *CardUpdateRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardUpdateRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardUpdateRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserToken

`func (o *CardUpdateRequest) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardUpdateRequest) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardUpdateRequest) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *CardUpdateRequest) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


