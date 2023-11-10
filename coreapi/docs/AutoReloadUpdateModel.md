# AutoReloadUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the auto reload is active.  Only one auto reload per level, per object, can be active. | [optional] [default to true]
**Association** | Pointer to [**AutoReloadAssociation**](AutoReloadAssociation.md) |  | [optional] 
**CurrencyCode** | Pointer to **string** | Three-digit link:https://www.iso.org/iso-4217-currency-codes.html[ISO 4217 currency code, window&#x3D;\&quot;_blank\&quot;]. | [optional] 
**FundingSourceAddressToken** | Pointer to **string** | Unique identifier of the funding source address to use for this auto reload.  If your funding source is an ACH account, then a &#x60;funding_source_address_token&#x60; is not required. If your funding source is a payment card, you must have at least one funding source address in order to create a GPA order. | [optional] 
**FundingSourceToken** | Pointer to **string** | Unique identifier of the funding source to use for this auto reload. | [optional] 
**OrderScope** | Pointer to [**OrderScope**](OrderScope.md) |  | [optional] 
**Token** | Pointer to **string** | The token in the path parameter takes precedence over the &#x60;token&#x60; body field. | [optional] 

## Methods

### NewAutoReloadUpdateModel

`func NewAutoReloadUpdateModel() *AutoReloadUpdateModel`

NewAutoReloadUpdateModel instantiates a new AutoReloadUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReloadUpdateModelWithDefaults

`func NewAutoReloadUpdateModelWithDefaults() *AutoReloadUpdateModel`

NewAutoReloadUpdateModelWithDefaults instantiates a new AutoReloadUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AutoReloadUpdateModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AutoReloadUpdateModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AutoReloadUpdateModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AutoReloadUpdateModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociation

`func (o *AutoReloadUpdateModel) GetAssociation() AutoReloadAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AutoReloadUpdateModel) GetAssociationOk() (*AutoReloadAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AutoReloadUpdateModel) SetAssociation(v AutoReloadAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AutoReloadUpdateModel) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AutoReloadUpdateModel) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AutoReloadUpdateModel) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AutoReloadUpdateModel) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AutoReloadUpdateModel) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFundingSourceAddressToken

`func (o *AutoReloadUpdateModel) GetFundingSourceAddressToken() string`

GetFundingSourceAddressToken returns the FundingSourceAddressToken field if non-nil, zero value otherwise.

### GetFundingSourceAddressTokenOk

`func (o *AutoReloadUpdateModel) GetFundingSourceAddressTokenOk() (*string, bool)`

GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceAddressToken

`func (o *AutoReloadUpdateModel) SetFundingSourceAddressToken(v string)`

SetFundingSourceAddressToken sets FundingSourceAddressToken field to given value.

### HasFundingSourceAddressToken

`func (o *AutoReloadUpdateModel) HasFundingSourceAddressToken() bool`

HasFundingSourceAddressToken returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *AutoReloadUpdateModel) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *AutoReloadUpdateModel) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *AutoReloadUpdateModel) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.

### HasFundingSourceToken

`func (o *AutoReloadUpdateModel) HasFundingSourceToken() bool`

HasFundingSourceToken returns a boolean if a field has been set.

### GetOrderScope

`func (o *AutoReloadUpdateModel) GetOrderScope() OrderScope`

GetOrderScope returns the OrderScope field if non-nil, zero value otherwise.

### GetOrderScopeOk

`func (o *AutoReloadUpdateModel) GetOrderScopeOk() (*OrderScope, bool)`

GetOrderScopeOk returns a tuple with the OrderScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderScope

`func (o *AutoReloadUpdateModel) SetOrderScope(v OrderScope)`

SetOrderScope sets OrderScope field to given value.

### HasOrderScope

`func (o *AutoReloadUpdateModel) HasOrderScope() bool`

HasOrderScope returns a boolean if a field has been set.

### GetToken

`func (o *AutoReloadUpdateModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AutoReloadUpdateModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AutoReloadUpdateModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AutoReloadUpdateModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


