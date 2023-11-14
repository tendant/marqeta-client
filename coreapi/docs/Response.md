# Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInformation** | Pointer to **string** | Information about the velocity control applied to the transaction.  *NOTE:* This field is returned only in transaction response objects. It is not returned in address verification or card security verification response objects. | [optional] 
**Code** | **string** | Four-digit response code for address verification, card security code verification, or transactions.  For address verification responses, the code is an assertion by the Marqeta platform as to whether its address verification data matches that provided by the cardholder:  [cols&#x3D;\&quot;2,3,3\&quot;] !&#x3D;&#x3D;&#x3D; ! Code ! Address ! Postal Code  ! 0000 ! Match ! Match  ! 0001 ! Match ! Unmatched  ! 0100 ! Unmatched ! Match  ! 0101 ! Unmatched ! Unmatched  ! 0200 ! Data Not Present ! Match  ! 0201 ! Data Not Present ! Unmatched  ! 0002 ! Match ! Data Not Present  ! 0102 ! Unmatched ! Data Not Present  ! 0303 ! Not Validated ! Not Validated !&#x3D;&#x3D;&#x3D;  For card security verification, the code indicates whether the verification check passed and can have these possible values:  * 0000  passed * 0001  did not pass  For a transaction, the code describes the outcome of the attempted transaction. For the full list of transaction codes, see &lt;&lt;/developer-guides/about-transactions#_transaction_response_codes, Transaction response codes&gt;&gt;. | 
**Memo** | **string** | Additional text that describes the response. | 

## Methods

### NewResponse

`func NewResponse(code string, memo string, ) *Response`

NewResponse instantiates a new Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWithDefaults

`func NewResponseWithDefaults() *Response`

NewResponseWithDefaults instantiates a new Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInformation

`func (o *Response) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *Response) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *Response) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *Response) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetCode

`func (o *Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetMemo

`func (o *Response) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *Response) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *Response) SetMemo(v string)`

SetMemo sets Memo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


