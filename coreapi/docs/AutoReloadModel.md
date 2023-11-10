# AutoReloadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the auto reload is active.  Only one auto reload per level, per object, can be active. | [optional] [default to true]
**Association** | Pointer to [**AutoReloadAssociation**](AutoReloadAssociation.md) |  | [optional] 
**CurrencyCode** | **string** | Three-digit link:https://www.iso.org/iso-4217-currency-codes.html[ISO 4217 currency code, window&#x3D;\&quot;_blank\&quot;]. | 
**FundingSourceAddressToken** | Pointer to **string** | Unique identifier of the funding source address to use for this auto reload.  If your funding source is an ACH account, then a &#x60;funding_source_address_token&#x60; is not required. If your funding source is a payment card, you must have at least one funding source address in order to create a GPA order.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/addresses/user/{user_token}&#x60; to retrieve address tokens for a user.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/addresses/business/{business_token}&#x60; to retrieve address tokens for a business. | [optional] 
**FundingSourceToken** | Pointer to **string** | Unique identifier of the funding source to use for this auto reload.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/user/{user_token}&#x60; to retrieve funding source tokens for a user.  Send a &#x60;GET&#x60; request to &#x60;/fundingsources/business/{business_token}&#x60; to retrieve funding source tokens for a business. | [optional] 
**OrderScope** | [**OrderScope**](OrderScope.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the auto reload.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewAutoReloadModel

`func NewAutoReloadModel(currencyCode string, orderScope OrderScope, ) *AutoReloadModel`

NewAutoReloadModel instantiates a new AutoReloadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReloadModelWithDefaults

`func NewAutoReloadModelWithDefaults() *AutoReloadModel`

NewAutoReloadModelWithDefaults instantiates a new AutoReloadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AutoReloadModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AutoReloadModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AutoReloadModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AutoReloadModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociation

`func (o *AutoReloadModel) GetAssociation() AutoReloadAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AutoReloadModel) GetAssociationOk() (*AutoReloadAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AutoReloadModel) SetAssociation(v AutoReloadAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AutoReloadModel) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AutoReloadModel) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AutoReloadModel) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AutoReloadModel) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFundingSourceAddressToken

`func (o *AutoReloadModel) GetFundingSourceAddressToken() string`

GetFundingSourceAddressToken returns the FundingSourceAddressToken field if non-nil, zero value otherwise.

### GetFundingSourceAddressTokenOk

`func (o *AutoReloadModel) GetFundingSourceAddressTokenOk() (*string, bool)`

GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceAddressToken

`func (o *AutoReloadModel) SetFundingSourceAddressToken(v string)`

SetFundingSourceAddressToken sets FundingSourceAddressToken field to given value.

### HasFundingSourceAddressToken

`func (o *AutoReloadModel) HasFundingSourceAddressToken() bool`

HasFundingSourceAddressToken returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *AutoReloadModel) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *AutoReloadModel) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *AutoReloadModel) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.

### HasFundingSourceToken

`func (o *AutoReloadModel) HasFundingSourceToken() bool`

HasFundingSourceToken returns a boolean if a field has been set.

### GetOrderScope

`func (o *AutoReloadModel) GetOrderScope() OrderScope`

GetOrderScope returns the OrderScope field if non-nil, zero value otherwise.

### GetOrderScopeOk

`func (o *AutoReloadModel) GetOrderScopeOk() (*OrderScope, bool)`

GetOrderScopeOk returns a tuple with the OrderScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderScope

`func (o *AutoReloadModel) SetOrderScope(v OrderScope)`

SetOrderScope sets OrderScope field to given value.


### GetToken

`func (o *AutoReloadModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AutoReloadModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AutoReloadModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AutoReloadModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


