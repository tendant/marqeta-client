# DigitalWalletTokenization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardArtId** | Pointer to **string** | Specifies the digital wallet card art identifier for the card product. Digital wallets display the card art after the initial token has been provisioned and activated. Digital wallet card art is updated for all wallets automatically whenever a tokenized card is reissued or replaced.  * If your card program is Managed by Marqeta, Marqeta populates this field on your behalf. * If your card program is Powered by Marqeta, you can obtain the correct card art identifier directly from Visa or Mastercard.  If this field is left blank, your card product inherits the card art assigned to the account BIN range. | [optional] 
**ProvisioningControls** | Pointer to [**ProvisioningControls**](ProvisioningControls.md) |  | [optional] 

## Methods

### NewDigitalWalletTokenization

`func NewDigitalWalletTokenization() *DigitalWalletTokenization`

NewDigitalWalletTokenization instantiates a new DigitalWalletTokenization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigitalWalletTokenizationWithDefaults

`func NewDigitalWalletTokenizationWithDefaults() *DigitalWalletTokenization`

NewDigitalWalletTokenizationWithDefaults instantiates a new DigitalWalletTokenization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardArtId

`func (o *DigitalWalletTokenization) GetCardArtId() string`

GetCardArtId returns the CardArtId field if non-nil, zero value otherwise.

### GetCardArtIdOk

`func (o *DigitalWalletTokenization) GetCardArtIdOk() (*string, bool)`

GetCardArtIdOk returns a tuple with the CardArtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardArtId

`func (o *DigitalWalletTokenization) SetCardArtId(v string)`

SetCardArtId sets CardArtId field to given value.

### HasCardArtId

`func (o *DigitalWalletTokenization) HasCardArtId() bool`

HasCardArtId returns a boolean if a field has been set.

### GetProvisioningControls

`func (o *DigitalWalletTokenization) GetProvisioningControls() ProvisioningControls`

GetProvisioningControls returns the ProvisioningControls field if non-nil, zero value otherwise.

### GetProvisioningControlsOk

`func (o *DigitalWalletTokenization) GetProvisioningControlsOk() (*ProvisioningControls, bool)`

GetProvisioningControlsOk returns a tuple with the ProvisioningControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningControls

`func (o *DigitalWalletTokenization) SetProvisioningControls(v ProvisioningControls)`

SetProvisioningControls sets ProvisioningControls field to given value.

### HasProvisioningControls

`func (o *DigitalWalletTokenization) HasProvisioningControls() bool`

HasProvisioningControls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


