# ImagesCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Specifies a PNG image to display on the card. | [optional] 
**ThermalColor** | Pointer to **string** | Specifies the color of the image displayed on the card. | [optional] 

## Methods

### NewImagesCard

`func NewImagesCard() *ImagesCard`

NewImagesCard instantiates a new ImagesCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImagesCardWithDefaults

`func NewImagesCardWithDefaults() *ImagesCard`

NewImagesCardWithDefaults instantiates a new ImagesCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImagesCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImagesCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImagesCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImagesCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetThermalColor

`func (o *ImagesCard) GetThermalColor() string`

GetThermalColor returns the ThermalColor field if non-nil, zero value otherwise.

### GetThermalColorOk

`func (o *ImagesCard) GetThermalColorOk() (*string, bool)`

GetThermalColorOk returns a tuple with the ThermalColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermalColor

`func (o *ImagesCard) SetThermalColor(v string)`

SetThermalColor sets ThermalColor field to given value.

### HasThermalColor

`func (o *ImagesCard) HasThermalColor() bool`

HasThermalColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


