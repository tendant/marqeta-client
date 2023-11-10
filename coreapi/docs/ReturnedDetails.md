# ReturnedDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnCode** | **string** | Return code for the returned payment. For more, see &lt;&lt;/developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes&gt;&gt;. | 
**ReturnReason** | **string** | Reason the payment was returned. For more, see &lt;&lt;/developer-guides/ach-origination#_nacha_ach_return_codes, NACHA ACH return codes&gt;&gt;. | 

## Methods

### NewReturnedDetails

`func NewReturnedDetails(returnCode string, returnReason string, ) *ReturnedDetails`

NewReturnedDetails instantiates a new ReturnedDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnedDetailsWithDefaults

`func NewReturnedDetailsWithDefaults() *ReturnedDetails`

NewReturnedDetailsWithDefaults instantiates a new ReturnedDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnCode

`func (o *ReturnedDetails) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *ReturnedDetails) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *ReturnedDetails) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.


### GetReturnReason

`func (o *ReturnedDetails) GetReturnReason() string`

GetReturnReason returns the ReturnReason field if non-nil, zero value otherwise.

### GetReturnReasonOk

`func (o *ReturnedDetails) GetReturnReasonOk() (*string, bool)`

GetReturnReasonOk returns a tuple with the ReturnReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReason

`func (o *ReturnedDetails) SetReturnReason(v string)`

SetReturnReason sets ReturnReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


