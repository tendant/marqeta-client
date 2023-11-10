# CardPersonalization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to [**Carrier**](Carrier.md) |  | [optional] 
**Images** | Pointer to [**Images**](Images.md) |  | [optional] 
**PersoType** | Pointer to **string** | Specifies the type of card personalization. | [optional] 
**Text** | [**Text**](Text.md) |  | 

## Methods

### NewCardPersonalization

`func NewCardPersonalization(text Text, ) *CardPersonalization`

NewCardPersonalization instantiates a new CardPersonalization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardPersonalizationWithDefaults

`func NewCardPersonalizationWithDefaults() *CardPersonalization`

NewCardPersonalizationWithDefaults instantiates a new CardPersonalization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *CardPersonalization) GetCarrier() Carrier`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *CardPersonalization) GetCarrierOk() (*Carrier, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *CardPersonalization) SetCarrier(v Carrier)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *CardPersonalization) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetImages

`func (o *CardPersonalization) GetImages() Images`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CardPersonalization) GetImagesOk() (*Images, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CardPersonalization) SetImages(v Images)`

SetImages sets Images field to given value.

### HasImages

`func (o *CardPersonalization) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetPersoType

`func (o *CardPersonalization) GetPersoType() string`

GetPersoType returns the PersoType field if non-nil, zero value otherwise.

### GetPersoTypeOk

`func (o *CardPersonalization) GetPersoTypeOk() (*string, bool)`

GetPersoTypeOk returns a tuple with the PersoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersoType

`func (o *CardPersonalization) SetPersoType(v string)`

SetPersoType sets PersoType field to given value.

### HasPersoType

`func (o *CardPersonalization) HasPersoType() bool`

HasPersoType returns a boolean if a field has been set.

### GetText

`func (o *CardPersonalization) GetText() Text`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CardPersonalization) GetTextOk() (*Text, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CardPersonalization) SetText(v Text)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


