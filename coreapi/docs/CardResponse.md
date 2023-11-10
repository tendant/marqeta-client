# CardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationActions** | Pointer to [**ActivationActions**](ActivationActions.md) |  | [optional] 
**Barcode** | **string** | Barcode printed on the card, expressed as numerals. | 
**BulkIssuanceToken** | Pointer to **string** | Unique identifier of the bulk card order. | [optional] 
**CardProductToken** | **string** | Unique identifier of the card product. | 
**ChipCvvNumber** | Pointer to **string** | Three-digit card verification value (ICVV) stored on the chip of the card. | [optional] 
**ContactlessExemptionCounter** | Pointer to **int32** | Running count of the contactless transactions successfully completed since the last strong customer authentication (SCA) challenge was issued. You can limit the number of contactless transactions that can be performed without issuing an SCA challenge at the card product level.  For more information about strong customer authentication, see &lt;&lt;/core-api/card-products, Card Products&gt;&gt;. | [optional] 
**ContactlessExemptionTotalAmount** | Pointer to **float32** | Running total of the money spent in contactless transactions successfully completed since the last strong customer authentication (SCA) challenge was issued. You can limit the total amount that can be spent in contactless transactions without issuing an SCA challenge at the card product level.  For more information about strong customer authentication, see &lt;&lt;/core-api/card-products, Card Products&gt;&gt;. | [optional] 
**CreatedTime** | **time.Time** | Date and time when the resource was created, in UTC. | 
**CvvNumber** | Pointer to **string** | Three-digit card verification value (CVV2 or CVC2) printed on the card. | [optional] 
**Expedite** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that you requested expedited processing of the card from your card fulfillment provider. | [optional] [default to false]
**Expiration** | **string** | Expiration date in &#x60;MMyy&#x60; format. | 
**ExpirationTime** | **time.Time** | Expiration date and time, in UTC. | 
**Fulfillment** | Pointer to [**CardFulfillmentResponse**](CardFulfillmentResponse.md) |  | [optional] 
**FulfillmentStatus** | **string** | Card fulfillment status:  * *ISSUED:* Initial state of all newly created/issued cards. * *ORDERED:* Card ordered through the card fulfillment provider. * *REORDERED:* Card reordered through the card fulfillment provider. * *REJECTED:* Card rejected by the card fulfillment provider. * *SHIPPED:* Card shipped by the card fulfillment provider. * *DELIVERED:* Card delivered by the card fulfillment provider. * *DIGITALLY_PRESENTED:* Card digitally presented using the &#x60;/cards/{token}/showpan&#x60; endpoint; does not affect the delivery of physical cards. | 
**InstrumentType** | Pointer to **string** | Instrument type of the card:  * *PHYSICAL_MSR:* A physical card with a magnetic stripe. This is the default physical card type. * *PHYSICAL_ICC:* A physical card with an integrated circuit, or \&quot;chip.\&quot; * *PHYSICAL_CONTACTLESS:* A physical card that uses radio frequency identification (RFID) or near-field communication (NFC) to enable payment over a secure radio interface. * *PHYSICAL_COMBO:* A physical card with a chip that also supports contactless payments. * *VIRTUAL_PAN:* A virtual card with a primary account number (PAN). | [optional] 
**LastFour** | **string** | Last four digits of the card primary account number (PAN). | 
**LastModifiedTime** | **time.Time** | Date and time when the resource was last modified, in UTC. | 
**Metadata** | Pointer to **map[string]string** | Associates customer-provided metadata with the card. | [optional] 
**NewPanFromCardToken** | Pointer to **string** | Reissues the specified card (known as the \&quot;source\&quot; card) with a new primary account number (PAN). | [optional] 
**Pan** | **string** | Primary account number (PAN) of the card. | 
**PinIsSet** | **bool** | Specifies if the personal identification number (PIN) has been set for the card. | [default to false]
**ReissuePanFromCardToken** | Pointer to **string** | Reissues the specified card (known as the \&quot;source\&quot; card). | [optional] 
**State** | **string** | Indicates the state of the card. | 
**StateReason** | **string** | Descriptive reason for why the card is in its current state. For example, \&quot;Card activated by cardholder\&quot;. | 
**Token** | **string** | Unique identifier of the card. | 
**TranslatePinFromCardToken** | Pointer to **string** | Copies the personal identification number (PIN) from the specified source card to the newly created card. | [optional] 
**UserToken** | **string** | Unique identifier of the cardholder. | 

## Methods

### NewCardResponse

`func NewCardResponse(barcode string, cardProductToken string, createdTime time.Time, expiration string, expirationTime time.Time, fulfillmentStatus string, lastFour string, lastModifiedTime time.Time, pan string, pinIsSet bool, state string, stateReason string, token string, userToken string, ) *CardResponse`

NewCardResponse instantiates a new CardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardResponseWithDefaults

`func NewCardResponseWithDefaults() *CardResponse`

NewCardResponseWithDefaults instantiates a new CardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationActions

`func (o *CardResponse) GetActivationActions() ActivationActions`

GetActivationActions returns the ActivationActions field if non-nil, zero value otherwise.

### GetActivationActionsOk

`func (o *CardResponse) GetActivationActionsOk() (*ActivationActions, bool)`

GetActivationActionsOk returns a tuple with the ActivationActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationActions

`func (o *CardResponse) SetActivationActions(v ActivationActions)`

SetActivationActions sets ActivationActions field to given value.

### HasActivationActions

`func (o *CardResponse) HasActivationActions() bool`

HasActivationActions returns a boolean if a field has been set.

### GetBarcode

`func (o *CardResponse) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *CardResponse) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *CardResponse) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetBulkIssuanceToken

`func (o *CardResponse) GetBulkIssuanceToken() string`

GetBulkIssuanceToken returns the BulkIssuanceToken field if non-nil, zero value otherwise.

### GetBulkIssuanceTokenOk

`func (o *CardResponse) GetBulkIssuanceTokenOk() (*string, bool)`

GetBulkIssuanceTokenOk returns a tuple with the BulkIssuanceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkIssuanceToken

`func (o *CardResponse) SetBulkIssuanceToken(v string)`

SetBulkIssuanceToken sets BulkIssuanceToken field to given value.

### HasBulkIssuanceToken

`func (o *CardResponse) HasBulkIssuanceToken() bool`

HasBulkIssuanceToken returns a boolean if a field has been set.

### GetCardProductToken

`func (o *CardResponse) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *CardResponse) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *CardResponse) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.


### GetChipCvvNumber

`func (o *CardResponse) GetChipCvvNumber() string`

GetChipCvvNumber returns the ChipCvvNumber field if non-nil, zero value otherwise.

### GetChipCvvNumberOk

`func (o *CardResponse) GetChipCvvNumberOk() (*string, bool)`

GetChipCvvNumberOk returns a tuple with the ChipCvvNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChipCvvNumber

`func (o *CardResponse) SetChipCvvNumber(v string)`

SetChipCvvNumber sets ChipCvvNumber field to given value.

### HasChipCvvNumber

`func (o *CardResponse) HasChipCvvNumber() bool`

HasChipCvvNumber returns a boolean if a field has been set.

### GetContactlessExemptionCounter

`func (o *CardResponse) GetContactlessExemptionCounter() int32`

GetContactlessExemptionCounter returns the ContactlessExemptionCounter field if non-nil, zero value otherwise.

### GetContactlessExemptionCounterOk

`func (o *CardResponse) GetContactlessExemptionCounterOk() (*int32, bool)`

GetContactlessExemptionCounterOk returns a tuple with the ContactlessExemptionCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactlessExemptionCounter

`func (o *CardResponse) SetContactlessExemptionCounter(v int32)`

SetContactlessExemptionCounter sets ContactlessExemptionCounter field to given value.

### HasContactlessExemptionCounter

`func (o *CardResponse) HasContactlessExemptionCounter() bool`

HasContactlessExemptionCounter returns a boolean if a field has been set.

### GetContactlessExemptionTotalAmount

`func (o *CardResponse) GetContactlessExemptionTotalAmount() float32`

GetContactlessExemptionTotalAmount returns the ContactlessExemptionTotalAmount field if non-nil, zero value otherwise.

### GetContactlessExemptionTotalAmountOk

`func (o *CardResponse) GetContactlessExemptionTotalAmountOk() (*float32, bool)`

GetContactlessExemptionTotalAmountOk returns a tuple with the ContactlessExemptionTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactlessExemptionTotalAmount

`func (o *CardResponse) SetContactlessExemptionTotalAmount(v float32)`

SetContactlessExemptionTotalAmount sets ContactlessExemptionTotalAmount field to given value.

### HasContactlessExemptionTotalAmount

`func (o *CardResponse) HasContactlessExemptionTotalAmount() bool`

HasContactlessExemptionTotalAmount returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CardResponse) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CardResponse) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CardResponse) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetCvvNumber

`func (o *CardResponse) GetCvvNumber() string`

GetCvvNumber returns the CvvNumber field if non-nil, zero value otherwise.

### GetCvvNumberOk

`func (o *CardResponse) GetCvvNumberOk() (*string, bool)`

GetCvvNumberOk returns a tuple with the CvvNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvNumber

`func (o *CardResponse) SetCvvNumber(v string)`

SetCvvNumber sets CvvNumber field to given value.

### HasCvvNumber

`func (o *CardResponse) HasCvvNumber() bool`

HasCvvNumber returns a boolean if a field has been set.

### GetExpedite

`func (o *CardResponse) GetExpedite() bool`

GetExpedite returns the Expedite field if non-nil, zero value otherwise.

### GetExpediteOk

`func (o *CardResponse) GetExpediteOk() (*bool, bool)`

GetExpediteOk returns a tuple with the Expedite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedite

`func (o *CardResponse) SetExpedite(v bool)`

SetExpedite sets Expedite field to given value.

### HasExpedite

`func (o *CardResponse) HasExpedite() bool`

HasExpedite returns a boolean if a field has been set.

### GetExpiration

`func (o *CardResponse) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CardResponse) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CardResponse) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### GetExpirationTime

`func (o *CardResponse) GetExpirationTime() time.Time`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *CardResponse) GetExpirationTimeOk() (*time.Time, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *CardResponse) SetExpirationTime(v time.Time)`

SetExpirationTime sets ExpirationTime field to given value.


### GetFulfillment

`func (o *CardResponse) GetFulfillment() CardFulfillmentResponse`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *CardResponse) GetFulfillmentOk() (*CardFulfillmentResponse, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *CardResponse) SetFulfillment(v CardFulfillmentResponse)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *CardResponse) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *CardResponse) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *CardResponse) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *CardResponse) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.


### GetInstrumentType

`func (o *CardResponse) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *CardResponse) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *CardResponse) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *CardResponse) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetLastFour

`func (o *CardResponse) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *CardResponse) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *CardResponse) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.


### GetLastModifiedTime

`func (o *CardResponse) GetLastModifiedTime() time.Time`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *CardResponse) GetLastModifiedTimeOk() (*time.Time, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *CardResponse) SetLastModifiedTime(v time.Time)`

SetLastModifiedTime sets LastModifiedTime field to given value.


### GetMetadata

`func (o *CardResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CardResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CardResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CardResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNewPanFromCardToken

`func (o *CardResponse) GetNewPanFromCardToken() string`

GetNewPanFromCardToken returns the NewPanFromCardToken field if non-nil, zero value otherwise.

### GetNewPanFromCardTokenOk

`func (o *CardResponse) GetNewPanFromCardTokenOk() (*string, bool)`

GetNewPanFromCardTokenOk returns a tuple with the NewPanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPanFromCardToken

`func (o *CardResponse) SetNewPanFromCardToken(v string)`

SetNewPanFromCardToken sets NewPanFromCardToken field to given value.

### HasNewPanFromCardToken

`func (o *CardResponse) HasNewPanFromCardToken() bool`

HasNewPanFromCardToken returns a boolean if a field has been set.

### GetPan

`func (o *CardResponse) GetPan() string`

GetPan returns the Pan field if non-nil, zero value otherwise.

### GetPanOk

`func (o *CardResponse) GetPanOk() (*string, bool)`

GetPanOk returns a tuple with the Pan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPan

`func (o *CardResponse) SetPan(v string)`

SetPan sets Pan field to given value.


### GetPinIsSet

`func (o *CardResponse) GetPinIsSet() bool`

GetPinIsSet returns the PinIsSet field if non-nil, zero value otherwise.

### GetPinIsSetOk

`func (o *CardResponse) GetPinIsSetOk() (*bool, bool)`

GetPinIsSetOk returns a tuple with the PinIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinIsSet

`func (o *CardResponse) SetPinIsSet(v bool)`

SetPinIsSet sets PinIsSet field to given value.


### GetReissuePanFromCardToken

`func (o *CardResponse) GetReissuePanFromCardToken() string`

GetReissuePanFromCardToken returns the ReissuePanFromCardToken field if non-nil, zero value otherwise.

### GetReissuePanFromCardTokenOk

`func (o *CardResponse) GetReissuePanFromCardTokenOk() (*string, bool)`

GetReissuePanFromCardTokenOk returns a tuple with the ReissuePanFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReissuePanFromCardToken

`func (o *CardResponse) SetReissuePanFromCardToken(v string)`

SetReissuePanFromCardToken sets ReissuePanFromCardToken field to given value.

### HasReissuePanFromCardToken

`func (o *CardResponse) HasReissuePanFromCardToken() bool`

HasReissuePanFromCardToken returns a boolean if a field has been set.

### GetState

`func (o *CardResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CardResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CardResponse) SetState(v string)`

SetState sets State field to given value.


### GetStateReason

`func (o *CardResponse) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *CardResponse) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *CardResponse) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.


### GetToken

`func (o *CardResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CardResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CardResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetTranslatePinFromCardToken

`func (o *CardResponse) GetTranslatePinFromCardToken() string`

GetTranslatePinFromCardToken returns the TranslatePinFromCardToken field if non-nil, zero value otherwise.

### GetTranslatePinFromCardTokenOk

`func (o *CardResponse) GetTranslatePinFromCardTokenOk() (*string, bool)`

GetTranslatePinFromCardTokenOk returns a tuple with the TranslatePinFromCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatePinFromCardToken

`func (o *CardResponse) SetTranslatePinFromCardToken(v string)`

SetTranslatePinFromCardToken sets TranslatePinFromCardToken field to given value.

### HasTranslatePinFromCardToken

`func (o *CardResponse) HasTranslatePinFromCardToken() bool`

HasTranslatePinFromCardToken returns a boolean if a field has been set.

### GetUserToken

`func (o *CardResponse) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *CardResponse) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *CardResponse) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


