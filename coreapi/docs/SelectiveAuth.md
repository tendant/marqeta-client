# SelectiveAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DmdLocationSensitivity** | Pointer to **int32** | Determines what type of merchant information is required for a match (authorization). Not relevant if &#x60;enable_regex_search_chain &#x3D; false&#x60;.  * *0*  Requires exact match on card acceptor name and postal code to existing entry in Marqeta Merchant database (most restrictive). * *1*  Partial match on card acceptor name (least restrictive). * *2*  Partial match on card acceptor name; exact match on card acceptor city. * *3*  Partial match on card acceptor name; exact match on card acceptor postal code. * *4*  Partial match on card acceptor name; exact match on street address 1 and postal code. | [optional] [default to 0]
**EnableRegexSearchChain** | Pointer to **bool** | Set to &#x60;true&#x60; to perform regular expression checking on the description received in the authorization. | [optional] [default to false]
**SaMode** | Pointer to **int32** | Specifies the selective authorization mode.  * *0* — Inactive * *1* — Active (attempts to authorize a merchant that does not have a recognized MID by matching other pieces of information) * *2* — Logging and notification (checks the transaction and logs results, but does not authorize)  Selective authorization applies to transactions that are limited to specific merchants. Matching requirements for authorization are set by the &#x60;dmd_location_sensitivity&#x60; field. | [optional] [default to 1]

## Methods

### NewSelectiveAuth

`func NewSelectiveAuth() *SelectiveAuth`

NewSelectiveAuth instantiates a new SelectiveAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectiveAuthWithDefaults

`func NewSelectiveAuthWithDefaults() *SelectiveAuth`

NewSelectiveAuthWithDefaults instantiates a new SelectiveAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDmdLocationSensitivity

`func (o *SelectiveAuth) GetDmdLocationSensitivity() int32`

GetDmdLocationSensitivity returns the DmdLocationSensitivity field if non-nil, zero value otherwise.

### GetDmdLocationSensitivityOk

`func (o *SelectiveAuth) GetDmdLocationSensitivityOk() (*int32, bool)`

GetDmdLocationSensitivityOk returns a tuple with the DmdLocationSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmdLocationSensitivity

`func (o *SelectiveAuth) SetDmdLocationSensitivity(v int32)`

SetDmdLocationSensitivity sets DmdLocationSensitivity field to given value.

### HasDmdLocationSensitivity

`func (o *SelectiveAuth) HasDmdLocationSensitivity() bool`

HasDmdLocationSensitivity returns a boolean if a field has been set.

### GetEnableRegexSearchChain

`func (o *SelectiveAuth) GetEnableRegexSearchChain() bool`

GetEnableRegexSearchChain returns the EnableRegexSearchChain field if non-nil, zero value otherwise.

### GetEnableRegexSearchChainOk

`func (o *SelectiveAuth) GetEnableRegexSearchChainOk() (*bool, bool)`

GetEnableRegexSearchChainOk returns a tuple with the EnableRegexSearchChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRegexSearchChain

`func (o *SelectiveAuth) SetEnableRegexSearchChain(v bool)`

SetEnableRegexSearchChain sets EnableRegexSearchChain field to given value.

### HasEnableRegexSearchChain

`func (o *SelectiveAuth) HasEnableRegexSearchChain() bool`

HasEnableRegexSearchChain returns a boolean if a field has been set.

### GetSaMode

`func (o *SelectiveAuth) GetSaMode() int32`

GetSaMode returns the SaMode field if non-nil, zero value otherwise.

### GetSaModeOk

`func (o *SelectiveAuth) GetSaModeOk() (*int32, bool)`

GetSaModeOk returns a tuple with the SaMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaMode

`func (o *SelectiveAuth) SetSaMode(v int32)`

SetSaMode sets SaMode field to given value.

### HasSaMode

`func (o *SelectiveAuth) HasSaMode() bool`

HasSaMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


