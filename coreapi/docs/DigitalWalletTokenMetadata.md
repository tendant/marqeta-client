# DigitalWalletTokenMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardproductPreferredNotificationLanguage** | Pointer to **string** | Language specified in the &#x60;config.transaction_controls.notification_language&#x60; field of the card product: &#x60;ces&#x60; (Czech), &#x60;deu&#x60; (German), &#x60;eng&#x60; (English), &#x60;fra&#x60; (French), &#x60;ita&#x60; (Italian), &#x60;pol&#x60; (Polish), &#x60;spa&#x60; (Spanish), &#x60;swe&#x60; (Swedish)  The ISO maintains the link:https://www.iso.org/iso-3166-country-codes.html[full list of ISO 3166 two- and three-digit numeric country codes, window&#x3D;\&quot;_blank\&quot;]. | [optional] 
**IssuerProductConfigId** | Pointer to **string** | Unique identifier of the product configuration on the Marqeta platform. | [optional] 

## Methods

### NewDigitalWalletTokenMetadata

`func NewDigitalWalletTokenMetadata() *DigitalWalletTokenMetadata`

NewDigitalWalletTokenMetadata instantiates a new DigitalWalletTokenMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletTokenMetadataWithDefaults

`func NewDigitalWalletTokenMetadataWithDefaults() *DigitalWalletTokenMetadata`

NewDigitalWalletTokenMetadataWithDefaults instantiates a new DigitalWalletTokenMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardproductPreferredNotificationLanguage

`func (o *DigitalWalletTokenMetadata) GetCardproductPreferredNotificationLanguage() string`

GetCardproductPreferredNotificationLanguage returns the CardproductPreferredNotificationLanguage field if non-nil, zero value otherwise.

### GetCardproductPreferredNotificationLanguageOk

`func (o *DigitalWalletTokenMetadata) GetCardproductPreferredNotificationLanguageOk() (*string, bool)`

GetCardproductPreferredNotificationLanguageOk returns a tuple with the CardproductPreferredNotificationLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardproductPreferredNotificationLanguage

`func (o *DigitalWalletTokenMetadata) SetCardproductPreferredNotificationLanguage(v string)`

SetCardproductPreferredNotificationLanguage sets CardproductPreferredNotificationLanguage field to given value.

### HasCardproductPreferredNotificationLanguage

`func (o *DigitalWalletTokenMetadata) HasCardproductPreferredNotificationLanguage() bool`

HasCardproductPreferredNotificationLanguage returns a boolean if a field has been set.

### GetIssuerProductConfigId

`func (o *DigitalWalletTokenMetadata) GetIssuerProductConfigId() string`

GetIssuerProductConfigId returns the IssuerProductConfigId field if non-nil, zero value otherwise.

### GetIssuerProductConfigIdOk

`func (o *DigitalWalletTokenMetadata) GetIssuerProductConfigIdOk() (*string, bool)`

GetIssuerProductConfigIdOk returns a tuple with the IssuerProductConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerProductConfigId

`func (o *DigitalWalletTokenMetadata) SetIssuerProductConfigId(v string)`

SetIssuerProductConfigId sets IssuerProductConfigId field to given value.

### HasIssuerProductConfigId

`func (o *DigitalWalletTokenMetadata) HasIssuerProductConfigId() bool`

HasIssuerProductConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


