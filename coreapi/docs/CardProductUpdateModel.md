# CardProductUpdateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Indicates whether the card product is active.  *NOTE:* This field has no effect on the ability to create cards from this card product. Use the &#x60;config.fulfillment.allow_card_creation&#x60; field to allow/disallow card creation. | [optional] [default to false]
**Config** | Pointer to [**CardProductConfig**](CardProductConfig.md) |  | [optional] 
**EndDate** | Pointer to **string** | End date of the range over which the card product can be active. | [optional] 
**Name** | Pointer to **string** | Name of the card product. Marqeta recommends that you use a unique string. | [optional] 
**StartDate** | Pointer to **string** | Date the card product becomes active. If the start date has passed and the card is set to &#x60;active &#x3D; false&#x60;, then the card will not be activated. | [optional] 

## Methods

### NewCardProductUpdateModel

`func NewCardProductUpdateModel() *CardProductUpdateModel`

NewCardProductUpdateModel instantiates a new CardProductUpdateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardProductUpdateModelWithDefaults

`func NewCardProductUpdateModelWithDefaults() *CardProductUpdateModel`

NewCardProductUpdateModelWithDefaults instantiates a new CardProductUpdateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CardProductUpdateModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CardProductUpdateModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CardProductUpdateModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CardProductUpdateModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *CardProductUpdateModel) GetConfig() CardProductConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CardProductUpdateModel) GetConfigOk() (*CardProductConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CardProductUpdateModel) SetConfig(v CardProductConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CardProductUpdateModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEndDate

`func (o *CardProductUpdateModel) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CardProductUpdateModel) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CardProductUpdateModel) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CardProductUpdateModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *CardProductUpdateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CardProductUpdateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CardProductUpdateModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CardProductUpdateModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *CardProductUpdateModel) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CardProductUpdateModel) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CardProductUpdateModel) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CardProductUpdateModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


