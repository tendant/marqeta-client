# AuthControlExemptMidsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the merchant identifier authorization control exception is active.  This field is returned if it exists in the resource. | [optional] [default to false]
**Association** | Pointer to [**SpendControlAssociation**](SpendControlAssociation.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Date and time when the resource was created, in UTC.  This field is returned if it exists in the resource. | [optional] 
**EndTime** | Pointer to **time.Time** | Date and time when the exception ends, in UTC.  This field is returned if it exists in the resource. | [optional] 
**LastUpdated** | Pointer to **time.Time** | Date and time when the resource was last updated, in UTC.  This field is returned if it exists in the resource. | [optional] 
**MerchantGroupToken** | Pointer to **string** | Unique identifier of the merchant group to be exempted.  This field is returned if it exists in the resource. | [optional] 
**Mid** | Pointer to **string** | Merchant to be exempted.  This field is returned if it exists in the resource. | [optional] 
**Name** | **string** | Name of the merchant identifier authorization control exemption. | 
**StartTime** | Pointer to **time.Time** | Date and time when the exception starts, in UTC.  This field is returned if it exists in the resource. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the merchant identifier authorization control exemption.  This field is always returned. | [optional] 

## Methods

### NewAuthControlExemptMidsResponse

`func NewAuthControlExemptMidsResponse(name string, ) *AuthControlExemptMidsResponse`

NewAuthControlExemptMidsResponse instantiates a new AuthControlExemptMidsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthControlExemptMidsResponseWithDefaults

`func NewAuthControlExemptMidsResponseWithDefaults() *AuthControlExemptMidsResponse`

NewAuthControlExemptMidsResponseWithDefaults instantiates a new AuthControlExemptMidsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AuthControlExemptMidsResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AuthControlExemptMidsResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AuthControlExemptMidsResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AuthControlExemptMidsResponse) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssociation

`func (o *AuthControlExemptMidsResponse) GetAssociation() SpendControlAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *AuthControlExemptMidsResponse) GetAssociationOk() (*SpendControlAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *AuthControlExemptMidsResponse) SetAssociation(v SpendControlAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *AuthControlExemptMidsResponse) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetCreated

`func (o *AuthControlExemptMidsResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthControlExemptMidsResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthControlExemptMidsResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthControlExemptMidsResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEndTime

`func (o *AuthControlExemptMidsResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AuthControlExemptMidsResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AuthControlExemptMidsResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AuthControlExemptMidsResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthControlExemptMidsResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthControlExemptMidsResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthControlExemptMidsResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthControlExemptMidsResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMerchantGroupToken

`func (o *AuthControlExemptMidsResponse) GetMerchantGroupToken() string`

GetMerchantGroupToken returns the MerchantGroupToken field if non-nil, zero value otherwise.

### GetMerchantGroupTokenOk

`func (o *AuthControlExemptMidsResponse) GetMerchantGroupTokenOk() (*string, bool)`

GetMerchantGroupTokenOk returns a tuple with the MerchantGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantGroupToken

`func (o *AuthControlExemptMidsResponse) SetMerchantGroupToken(v string)`

SetMerchantGroupToken sets MerchantGroupToken field to given value.

### HasMerchantGroupToken

`func (o *AuthControlExemptMidsResponse) HasMerchantGroupToken() bool`

HasMerchantGroupToken returns a boolean if a field has been set.

### GetMid

`func (o *AuthControlExemptMidsResponse) GetMid() string`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *AuthControlExemptMidsResponse) GetMidOk() (*string, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *AuthControlExemptMidsResponse) SetMid(v string)`

SetMid sets Mid field to given value.

### HasMid

`func (o *AuthControlExemptMidsResponse) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetName

`func (o *AuthControlExemptMidsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthControlExemptMidsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthControlExemptMidsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *AuthControlExemptMidsResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AuthControlExemptMidsResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AuthControlExemptMidsResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AuthControlExemptMidsResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetToken

`func (o *AuthControlExemptMidsResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthControlExemptMidsResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthControlExemptMidsResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthControlExemptMidsResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


