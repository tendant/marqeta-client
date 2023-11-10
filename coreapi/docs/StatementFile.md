# StatementFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountToken** | Pointer to **string** | Unique identifier of the credit account on which the statement PDF file is generated. | [optional] 
**ClosingDate** | Pointer to **time.Time** | Date and time when the statement period ended. | [optional] 
**OpeningDate** | Pointer to **time.Time** | Date and time when the statement period began. | [optional] 
**SignedUrl** | Pointer to **string** | The signed URL to retrieve the statement PDF file. | [optional] 
**StatementSummaryToken** | Pointer to **string** | Unique identifier of the statement summary. | [optional] 
**Token** | Pointer to **string** | Unique identifier of the statement file. | [optional] 
**Type** | Pointer to **string** | Type of file. | [optional] 

## Methods

### NewStatementFile

`func NewStatementFile() *StatementFile`

NewStatementFile instantiates a new StatementFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementFileWithDefaults

`func NewStatementFileWithDefaults() *StatementFile`

NewStatementFileWithDefaults instantiates a new StatementFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountToken

`func (o *StatementFile) GetAccountToken() string`

GetAccountToken returns the AccountToken field if non-nil, zero value otherwise.

### GetAccountTokenOk

`func (o *StatementFile) GetAccountTokenOk() (*string, bool)`

GetAccountTokenOk returns a tuple with the AccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountToken

`func (o *StatementFile) SetAccountToken(v string)`

SetAccountToken sets AccountToken field to given value.

### HasAccountToken

`func (o *StatementFile) HasAccountToken() bool`

HasAccountToken returns a boolean if a field has been set.

### GetClosingDate

`func (o *StatementFile) GetClosingDate() time.Time`

GetClosingDate returns the ClosingDate field if non-nil, zero value otherwise.

### GetClosingDateOk

`func (o *StatementFile) GetClosingDateOk() (*time.Time, bool)`

GetClosingDateOk returns a tuple with the ClosingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDate

`func (o *StatementFile) SetClosingDate(v time.Time)`

SetClosingDate sets ClosingDate field to given value.

### HasClosingDate

`func (o *StatementFile) HasClosingDate() bool`

HasClosingDate returns a boolean if a field has been set.

### GetOpeningDate

`func (o *StatementFile) GetOpeningDate() time.Time`

GetOpeningDate returns the OpeningDate field if non-nil, zero value otherwise.

### GetOpeningDateOk

`func (o *StatementFile) GetOpeningDateOk() (*time.Time, bool)`

GetOpeningDateOk returns a tuple with the OpeningDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningDate

`func (o *StatementFile) SetOpeningDate(v time.Time)`

SetOpeningDate sets OpeningDate field to given value.

### HasOpeningDate

`func (o *StatementFile) HasOpeningDate() bool`

HasOpeningDate returns a boolean if a field has been set.

### GetSignedUrl

`func (o *StatementFile) GetSignedUrl() string`

GetSignedUrl returns the SignedUrl field if non-nil, zero value otherwise.

### GetSignedUrlOk

`func (o *StatementFile) GetSignedUrlOk() (*string, bool)`

GetSignedUrlOk returns a tuple with the SignedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedUrl

`func (o *StatementFile) SetSignedUrl(v string)`

SetSignedUrl sets SignedUrl field to given value.

### HasSignedUrl

`func (o *StatementFile) HasSignedUrl() bool`

HasSignedUrl returns a boolean if a field has been set.

### GetStatementSummaryToken

`func (o *StatementFile) GetStatementSummaryToken() string`

GetStatementSummaryToken returns the StatementSummaryToken field if non-nil, zero value otherwise.

### GetStatementSummaryTokenOk

`func (o *StatementFile) GetStatementSummaryTokenOk() (*string, bool)`

GetStatementSummaryTokenOk returns a tuple with the StatementSummaryToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementSummaryToken

`func (o *StatementFile) SetStatementSummaryToken(v string)`

SetStatementSummaryToken sets StatementSummaryToken field to given value.

### HasStatementSummaryToken

`func (o *StatementFile) HasStatementSummaryToken() bool`

HasStatementSummaryToken returns a boolean if a field has been set.

### GetToken

`func (o *StatementFile) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StatementFile) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StatementFile) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *StatementFile) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *StatementFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StatementFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StatementFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StatementFile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


