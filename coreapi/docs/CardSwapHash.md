# CardSwapHash

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewCardToken** | **string** | Unique identifier of the new card resource to which the digital wallet tokens are assigned. | 
**PreviousCardToken** | **string** | Unique identifier of the existing card resource that has digital wallet tokens assigned to it. | 

## Methods

### NewCardSwapHash

`func NewCardSwapHash(newCardToken string, previousCardToken string, ) *CardSwapHash`

NewCardSwapHash instantiates a new CardSwapHash object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardSwapHashWithDefaults

`func NewCardSwapHashWithDefaults() *CardSwapHash`

NewCardSwapHashWithDefaults instantiates a new CardSwapHash object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewCardToken

`func (o *CardSwapHash) GetNewCardToken() string`

GetNewCardToken returns the NewCardToken field if non-nil, zero value otherwise.

### GetNewCardTokenOk

`func (o *CardSwapHash) GetNewCardTokenOk() (*string, bool)`

GetNewCardTokenOk returns a tuple with the NewCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCardToken

`func (o *CardSwapHash) SetNewCardToken(v string)`

SetNewCardToken sets NewCardToken field to given value.


### GetPreviousCardToken

`func (o *CardSwapHash) GetPreviousCardToken() string`

GetPreviousCardToken returns the PreviousCardToken field if non-nil, zero value otherwise.

### GetPreviousCardTokenOk

`func (o *CardSwapHash) GetPreviousCardTokenOk() (*string, bool)`

GetPreviousCardTokenOk returns a tuple with the PreviousCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCardToken

`func (o *CardSwapHash) SetPreviousCardToken(v string)`

SetPreviousCardToken sets PreviousCardToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


