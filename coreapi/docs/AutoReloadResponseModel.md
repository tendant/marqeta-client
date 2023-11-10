# AutoReloadResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Specifies whether the auto reload is active.  This field is returned if it exists in the resource. | [optional] [default to true]
**Association** | Pointer to [**AutoReloadAssociation**](AutoReloadAssociation.md) |  | [optional] 
**CreatedTime** | **time.Time** | Date and time when the auto reload object was created, in UTC. | 
**CurrencyCode** | **string** | Three-digit link:https://www.iso.org/iso-4217-currency-codes.html[ISO 4217 currency code, window&#x3D;\&quot;_blank\&quot;]. | 
**FundingSourceAddressToken** | Pointer to **string** | Unique identifier of the funding source address to use for this auto reload.  This field is returned if it exists in the resource. | [optional] 
**FundingSourceToken** | Pointer to **string** | Unique identifier of the funding source to use for this auto reload.  This field is returned if it exists in the resource. | [optional] 
**LastModifiedTime** | **time.Time** | Date and time when the auto reload object was last modified, in UTC. | 
**OrderScope** | [**OrderScope**](OrderScope.md) |  | 
**Token** | Pointer to **string** | Unique identifier of the auto reload.  This field is always returned. | [optional] 

## Methods

### NewAutoReloadResponseModel

`func NewAutoReloadResponseModel(createdTime time.Time, currencyCode string, lastModifiedTime time.Time, orderScope OrderScope, ) *AutoReloadResponseModel`

NewAutoReloadResponseModel instantiates a new AutoReloadResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoReloadResponseModelWithDefaults

`func NewAutoReloadResponseModelWithDefaults() *AutoReloadResponseModel`

NewAutoReloadResponseModelWithDefaults instantiates a new AutoReloadResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AutoReloadResponseModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AutoReloadResponseModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AutoReloadResponseModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AutoReloadResponseModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociation

`func (o *AutoReloadResponseModel) GetAssociation() AutoReloadAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AutoReloadResponseModel) GetAssociationOk() (*AutoReloadAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AutoReloadResponseModel) SetAssociation(v AutoReloadAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AutoReloadResponseModel) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AutoReloadResponseModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AutoReloadResponseModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AutoReloadResponseModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCurrencyCode

`func (o *AutoReloadResponseModel) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AutoReloadResponseModel) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AutoReloadResponseModel) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetFundingSourceAddressToken

`func (o *AutoReloadResponseModel) GetFundingSourceAddressToken() string`

GetFundingSourceAddressToken returns the FundingSourceAddressToken field if non-nil, zero value otherwise.

### GetFundingSourceAddressTokenOk

`func (o *AutoReloadResponseModel) GetFundingSourceAddressTokenOk() (*string, bool)`

GetFundingSourceAddressTokenOk returns a tuple with the FundingSourceAddressToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceAddressToken

`func (o *AutoReloadResponseModel) SetFundingSourceAddressToken(v string)`

SetFundingSourceAddressToken sets FundingSourceAddressToken field to given value.

### HasFundingSourceAddressToken

`func (o *AutoReloadResponseModel) HasFundingSourceAddressToken() bool`

HasFundingSourceAddressToken returns a boolean if a field has been set.

### GetFundingSourceToken

`func (o *AutoReloadResponseModel) GetFundingSourceToken() string`

GetFundingSourceToken returns the FundingSourceToken field if non-nil, zero value otherwise.

### GetFundingSourceTokenOk

`func (o *AutoReloadResponseModel) GetFundingSourceTokenOk() (*string, bool)`

GetFundingSourceTokenOk returns a tuple with the FundingSourceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSourceToken

`func (o *AutoReloadResponseModel) SetFundingSourceToken(v string)`

SetFundingSourceToken sets FundingSourceToken field to given value.

### HasFundingSourceToken

`func (o *AutoReloadResponseModel) HasFundingSourceToken() bool`

HasFundingSourceToken returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *AutoReloadResponseModel) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *AutoReloadResponseModel) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *AutoReloadResponseModel) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetOrderScope

`func (o *AutoReloadResponseModel) GetOrderScope() OrderScope`

GetOrderScope returns the OrderScope field if non-nil, zero value otherwise.

### GetOrderScopeOk

`func (o *AutoReloadResponseModel) GetOrderScopeOk() (*OrderScope, bool)`

GetOrderScopeOk returns a tuple with the OrderScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderScope

`func (o *AutoReloadResponseModel) SetOrderScope(v OrderScope)`

SetOrderScope sets OrderScope field to given value.


### GetToken

`func (o *AutoReloadResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AutoReloadResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AutoReloadResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AutoReloadResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


