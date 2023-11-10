# Special

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantOnBoarding** | Pointer to **bool** | If set to &#x60;true&#x60;, cards of this card product type can be used for merchant onboarding. | [optional] [default to false]

## Methods

### NewSpecial

`func NewSpecial() *Special`

NewSpecial instantiates a new Special object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialWithDefaults

`func NewSpecialWithDefaults() *Special`

NewSpecialWithDefaults instantiates a new Special object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantOnBoarding

`func (o *Special) GetMerchantOnBoarding() bool`

GetMerchantOnBoarding returns the MerchantOnBoarding field if non-nil, zero value otherwise.

### GetMerchantOnBoardingOk

`func (o *Special) GetMerchantOnBoardingOk() (*bool, bool)`

GetMerchantOnBoardingOk returns a tuple with the MerchantOnBoarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantOnBoarding

`func (o *Special) SetMerchantOnBoarding(v bool)`

SetMerchantOnBoarding sets MerchantOnBoarding field to given value.

### HasMerchantOnBoarding

`func (o *Special) HasMerchantOnBoarding() bool`

HasMerchantOnBoarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


