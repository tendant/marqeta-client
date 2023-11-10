# CardProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the card product is active.  *NOTE:* This field has no effect on the ability to create cards from this card product. Use the &#x60;config.fulfillment.allow_card_creation&#x60; field to allow/disallow card creation. | [optional] [default to false]
**Config** | Pointer to [**CardProductConfig**](CardProductConfig.md) |  | [optional] 
**EndDate** | Pointer to **string** | End date of the range over which the card product can be active. | [optional] 
**Name** | **string** | Name of the card product. Marqeta recommends that you use a unique string. | 
**StartDate** | **string** | Date when the card product becomes active. If the start date has passed and the card is set to &#x60;active &#x3D; false&#x60;, then the card will not be activated. | 
**Token** | Pointer to **string** | Unique identifier of the card product.  If you do not include a token, the system will generate one automatically. This token is required in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewCardProductRequest

`func NewCardProductRequest(name string, startDate string, ) *CardProductRequest`

NewCardProductRequest instantiates a new CardProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductRequestWithDefaults

`func NewCardProductRequestWithDefaults() *CardProductRequest`

NewCardProductRequestWithDefaults instantiates a new CardProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardProductRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProductRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProductRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardProductRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *CardProductRequest) GetConfig() CardProductConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CardProductRequest) GetConfigOk() (*CardProductConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CardProductRequest) SetConfig(v CardProductConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CardProductRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProductRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProductRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProductRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProductRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *CardProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProductRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CardProductRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProductRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProductRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetToken

`func (o *CardProductRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardProductRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardProductRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CardProductRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


