# ResultCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | For any &#x60;PENDING&#x60; or &#x60;FAILURE&#x60; outcome, see the &lt;&lt;user_kyc_failure_codes, User KYC failure codes&gt;&gt; table, the &lt;&lt;outcome_reasons_for_the_business, Outcome reasons for the business&gt;&gt; table, or the &lt;&lt;outcome_reasons_for_individuals_associated_with_a_business, Outcome reasons for individuals associated with a business&gt;&gt; table. | [optional] 
**Message** | Pointer to **string** | Result code description. | [optional] 

## Methods

### NewResultCode

`func NewResultCode() *ResultCode`

NewResultCode instantiates a new ResultCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCodeWithDefaults

`func NewResultCodeWithDefaults() *ResultCode`

NewResultCodeWithDefaults instantiates a new ResultCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ResultCode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ResultCode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ResultCode) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ResultCode) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ResultCode) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResultCode) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResultCode) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResultCode) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


