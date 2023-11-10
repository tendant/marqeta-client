# AuthControlExemptMidsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**EndTime** | Pointer to **time.Time** | Date and time when the exception ends, in UTC. | [optional] 
**MerchantGroupToken** | Pointer to **string** | Unique identifier of the merchant group to be exempted. This field is required if there is no entry in the &#x60;mid&#x60; field. Pass either this field or the &#x60;mid&#x60; field, not both.  For information about merchant groups, see &lt;&lt;/core-api/merchant-groups, Merchant Groups&gt;&gt;. | [optional] 
**Mid** | Pointer to **string** | Merchant to be exempted. This field is required if there is no entry in the &#x60;merchant_group_token&#x60; field. Use either this field or the &#x60;merchant_group_token&#x60; field, not both. | [optional] 
**Name** | **string** | Name of the merchant identifier authorization control exemption. | 
**StartTime** | Pointer to **time.Time** | Date and time when the exception starts, in UTC. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the merchant identifier authorization control exemption.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated. | [optional] 

## Methods

### NewAuthControlExemptMidsRequest

`func NewAuthControlExemptMidsRequest(name string, ) *AuthControlExemptMidsRequest`

NewAuthControlExemptMidsRequest instantiates a new AuthControlExemptMidsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthControlExemptMidsRequestWithDefaults

`func NewAuthControlExemptMidsRequestWithDefaults() *AuthControlExemptMidsRequest`

NewAuthControlExemptMidsRequestWithDefaults instantiates a new AuthControlExemptMidsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociation

`func (o *AuthControlExemptMidsRequest) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AuthControlExemptMidsRequest) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AuthControlExemptMidsRequest) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AuthControlExemptMidsRequest) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetEndTime

`func (o *AuthControlExemptMidsRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AuthControlExemptMidsRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AuthControlExemptMidsRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AuthControlExemptMidsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMerchantGroupToken

`func (o *AuthControlExemptMidsRequest) GetMerchantGroupToken() string`

GetMerchantGroupToken returns the MerchantGroupToken field if non-nil, zero value otherwise.

### GetMerchantGroupTokenOk

`func (o *AuthControlExemptMidsRequest) GetMerchantGroupTokenOk() (*string, bool)`

GetMerchantGroupTokenOk returns a tuple with the MerchantGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGroupToken

`func (o *AuthControlExemptMidsRequest) SetMerchantGroupToken(v string)`

SetMerchantGroupToken sets MerchantGroupToken field to given value.

### HasMerchantGroupToken

`func (o *AuthControlExemptMidsRequest) HasMerchantGroupToken() bool`

HasMerchantGroupToken returns a boolean if a field has been set.

### GetMid

`func (o *AuthControlExemptMidsRequest) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *AuthControlExemptMidsRequest) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *AuthControlExemptMidsRequest) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *AuthControlExemptMidsRequest) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetName

`func (o *AuthControlExemptMidsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthControlExemptMidsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthControlExemptMidsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *AuthControlExemptMidsRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AuthControlExemptMidsRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AuthControlExemptMidsRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AuthControlExemptMidsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetToken

`func (o *AuthControlExemptMidsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthControlExemptMidsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthControlExemptMidsRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthControlExemptMidsRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


