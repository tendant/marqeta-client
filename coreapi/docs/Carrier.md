# Carrier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoFile** | Pointer to **string** | Specifies an image to display on the card carrier. | [optional] 
**LogoThumbnailFile** | Pointer to **string** | Specifies a thumbnail-sized rendering of the image specified in the &#x60;logo_file&#x60; field. | [optional] 
**MessageFile** | Pointer to **string** | Specifies a text file containing a custom message to print on the card carrier. | [optional] 
**MessageLine** | Pointer to **string** | Specifies a custom message that appears on the card carrier. | [optional] 
**TemplateId** | Pointer to **string** | Specifies the card carrier template to use. | [optional] 

## Methods

### NewCarrier

`func NewCarrier() *Carrier`

NewCarrier instantiates a new Carrier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierWithDefaults

`func NewCarrierWithDefaults() *Carrier`

NewCarrierWithDefaults instantiates a new Carrier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogoFile

`func (o *Carrier) GetLogoFile() string`

GetLogoFile returns the LogoFile field if non-nil, zero value otherwise.

### GetLogoFileOk

`func (o *Carrier) GetLogoFileOk() (*string, bool)`

GetLogoFileOk returns a tuple with the LogoFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoFile

`func (o *Carrier) SetLogoFile(v string)`

SetLogoFile sets LogoFile field to given value.

### HasLogoFile

`func (o *Carrier) HasLogoFile() bool`

HasLogoFile returns a boolean if a field has been set.

### GetLogoThumbnailFile

`func (o *Carrier) GetLogoThumbnailFile() string`

GetLogoThumbnailFile returns the LogoThumbnailFile field if non-nil, zero value otherwise.

### GetLogoThumbnailFileOk

`func (o *Carrier) GetLogoThumbnailFileOk() (*string, bool)`

GetLogoThumbnailFileOk returns a tuple with the LogoThumbnailFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoThumbnailFile

`func (o *Carrier) SetLogoThumbnailFile(v string)`

SetLogoThumbnailFile sets LogoThumbnailFile field to given value.

### HasLogoThumbnailFile

`func (o *Carrier) HasLogoThumbnailFile() bool`

HasLogoThumbnailFile returns a boolean if a field has been set.

### GetMessageFile

`func (o *Carrier) GetMessageFile() string`

GetMessageFile returns the MessageFile field if non-nil, zero value otherwise.

### GetMessageFileOk

`func (o *Carrier) GetMessageFileOk() (*string, bool)`

GetMessageFileOk returns a tuple with the MessageFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageFile

`func (o *Carrier) SetMessageFile(v string)`

SetMessageFile sets MessageFile field to given value.

### HasMessageFile

`func (o *Carrier) HasMessageFile() bool`

HasMessageFile returns a boolean if a field has been set.

### GetMessageLine

`func (o *Carrier) GetMessageLine() string`

GetMessageLine returns the MessageLine field if non-nil, zero value otherwise.

### GetMessageLineOk

`func (o *Carrier) GetMessageLineOk() (*string, bool)`

GetMessageLineOk returns a tuple with the MessageLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageLine

`func (o *Carrier) SetMessageLine(v string)`

SetMessageLine sets MessageLine field to given value.

### HasMessageLine

`func (o *Carrier) HasMessageLine() bool`

HasMessageLine returns a boolean if a field has been set.

### GetTemplateId

`func (o *Carrier) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Carrier) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Carrier) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Carrier) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


