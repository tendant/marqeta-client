# TransactionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountFunding** | Pointer to [**AccountFunding**](AccountFunding.md) |  | [optional] 
**Acquirer** | Pointer to [**Acquirer**](Acquirer.md) |  | [optional] 
**AcquirerFeeAmount** | Pointer to **float32** | Indicates the amount of the acquirer fee. Account holders are sometimes charged an acquirer fee for card use at ATMs, fuel dispensers, and so on. | [optional] 
**AcquirerReferenceData** | Pointer to **string** |  | [optional] 
**AcquirerReferenceId** | Pointer to **string** | Acquirer-assigned unique identifier of the transaction. Useful for settlement and reconciliation. | [optional] 
**ActingUserToken** | **string** | Unique identifier of the user who conducted the transaction. This might be a child user configured to share its parent&#39;s account balance. | 
**AddressVerification** | Pointer to [**AddressVerificationModel**](AddressVerificationModel.md) |  | [optional] 
**AdviceReasonCode** | Pointer to **string** |  | [optional] 
**AdviceReasonDetails** | Pointer to **string** |  | [optional] 
**Amount** | **float32** | Amount of the transaction. | 
**AmountToBeReleased** | Pointer to **float32** | Amount of original authorization to be released. This field appears in final clearing transactions where the clearing amount is lower than the authorization amount. | [optional] 
**ApprovalCode** | Pointer to **string** | Unique identifier assigned to an authorization, printed on the receipt at point of sale. | [optional] 
**AtcInformation** | Pointer to [**AtcInformation**](AtcInformation.md) |  | [optional] 
**AutoReload** | Pointer to [**AutoReloadModel**](AutoReloadModel.md) |  | [optional] 
**BankTransferToken** | Pointer to **string** |  | [optional] 
**BatchNumber** | Pointer to **string** | The batch number of the transaction. | [optional] 
**Business** | Pointer to [**BusinessMetadata**](BusinessMetadata.md) |  | [optional] 
**BusinessToken** | Pointer to **string** | Unique identifier of the business that owns the account that funded the transaction. | [optional] 
**Card** | Pointer to [**CardResponse**](CardResponse.md) |  | [optional] 
**CardAcceptor** | Pointer to [**TransactionCardAcceptor**](TransactionCardAcceptor.md) |  | [optional] 
**CardHolderModel** | Pointer to [**UserCardHolderResponse**](UserCardHolderResponse.md) |  | [optional] 
**CardProductToken** | Pointer to **string** |  | [optional] 
**CardSecurityCodeVerification** | Pointer to [**CardSecurityCodeVerification**](CardSecurityCodeVerification.md) |  | [optional] 
**CardToken** | Pointer to **string** | Unique identifier of the card. Useful when a single account holder has multiple cards. | [optional] 
**CardholderAuthenticationData** | Pointer to [**CardholderAuthenticationData**](CardholderAuthenticationData.md) |  | [optional] 
**CashBackAmount** | Pointer to **float32** | Amount of cash back requested by the cardholder during the transaction. Included in the total transaction amount. | [optional] 
**Chargeback** | Pointer to [**ChargebackResponse**](ChargebackResponse.md) |  | [optional] 
**ClearingRecordSequenceNumber** | Pointer to **string** | A sequence number that identifies a specific clearing message among multiple clearing messages for an authorization. | [optional] 
**CreatedTime** | Pointer to **time.Time** | Date and time when the Marqeta platform created the transaction entry, in UTC format. For example, when Marqeta processed the clearing record for a refund. | [optional] 
**CurrencyCode** | Pointer to **string** | Currency type of the transaction. | [optional] 
**CurrencyConversion** | Pointer to [**CurrencyConversion**](CurrencyConversion.md) |  | [optional] 
**DeferredSettlementDays** | Pointer to **string** |  | [optional] 
**DigitalWalletToken** | Pointer to [**DigitalWalletToken**](DigitalWalletToken.md) |  | [optional] 
**Dispute** | Pointer to [**DisputeModel**](DisputeModel.md) |  | [optional] 
**Duration** | Pointer to **int32** | Duration of the transaction on Marqeta&#39;s servers, in milliseconds. | [optional] 
**EnhancedDataToken** | Pointer to **string** | The enhanced commercial card data token for the transaction. | [optional] 
**Fee** | Pointer to [**Fee**](Fee.md) |  | [optional] 
**FeeTransfer** | Pointer to [**FeeTransferResponse**](FeeTransferResponse.md) |  | [optional] 
**Fees** | Pointer to [**[]NetworkFeeModel**](NetworkFeeModel.md) | List of fees associated with the transaction.  This array is returned if it exists in the resource. | [optional] 
**Fraud** | Pointer to [**FraudView**](FraudView.md) |  | [optional] 
**FromAccount** | Pointer to **string** | Specifies the account type for ATM transactions. | [optional] 
**Gpa** | Pointer to [**CardholderBalance**](CardholderBalance.md) |  | [optional] 
**GpaOrder** | Pointer to [**GpaResponse**](GpaResponse.md) |  | [optional] 
**GpaOrderUnload** | Pointer to [**GpaReturns**](GpaReturns.md) |  | [optional] 
**Identifier** | Pointer to **string** | Sequential identifier of the transaction. | [optional] 
**IncrementalAuthorizationTransactionTokens** | Pointer to **[]string** | An array of incremental authorization transaction tokens. | [optional] 
**InterchangeRateDescriptor** | Pointer to **string** |  | [optional] 
**IsPreauthorization** | Pointer to **bool** | Indicates if the transaction is a pre-authorization. | [optional] [default to false]
**IsaIndicator** | Pointer to **string** | The international service assessment indicator indicates if an ISA fee is applicable to the transaction. | [optional] 
**IssuerInterchangeAmount** | Pointer to **float32** | The amount of interchange charged by the card issuer. | [optional] 
**IssuerPaymentNode** | Pointer to **string** | Unique identifier of the Marqeta platform server that received the transaction from the card network. | [optional] 
**IssuerReceivedTime** | Pointer to **string** | Date and time when the Marqeta platform received the transaction from the card network, in UTC. | [optional] 
**LocalTransactionDate** | Pointer to **time.Time** |  | [optional] 
**Merchant** | Pointer to [**MerchantResponseModel**](MerchantResponseModel.md) |  | [optional] 
**MsaOrderUnload** | Pointer to [**MsaReturns**](MsaReturns.md) |  | [optional] 
**MultiClearingSequenceCount** | Pointer to **string** | If an authorization has multiple clearing transactions, this field displays their total number. For example, if an authorization has four clearing transactions, the sequence count is &#x60;04&#x60;. | [optional] 
**MultiClearingSequenceNumber** | Pointer to **string** | If an authorization has multiple clearing transactions, this field displays the sequence number for the clearing transaction. For example, if this is the second clearing transaction of four, the sequence number is &#x60;02&#x60;. | [optional] 
**NationalNetCpdOfOriginal** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** | Indicates which card network was used to complete the transactions. | [optional] 
**NetworkMetadata** | Pointer to [**NetworkMetadata**](NetworkMetadata.md) |  | [optional] 
**NetworkReferenceId** | Pointer to **string** | Network-assigned unique identifier of the transaction. Useful for settlement and reconciliation. | [optional] 
**OriginalCredit** | Pointer to [**OriginalCredit**](OriginalCredit.md) |  | [optional] 
**PeerTransfer** | Pointer to [**PeerTransferResponse**](PeerTransferResponse.md) |  | [optional] 
**Polarity** | Pointer to **string** | Indicates whether the transaction is credit or debit. | [optional] 
**Pos** | Pointer to [**Pos**](Pos.md) |  | [optional] 
**PrecedingRelatedTransactionToken** | Pointer to **string** | Returned for final transaction types.  Unique identifier of the preceding related transaction. Useful for identifying the transaction that preceded the current one.  For example, &#x60;authorization&#x60;, a temporary transaction type, precedes and is completed by &#x60;authorization.clearing&#x60;, a final transaction type. In this case, the &#x60;authorization&#x60; token is returned with this field. For which transaction types are temporary or final, see &lt;&lt;/core-api/event-types#_transaction_events, Transaction events in Event Types&gt;&gt;. | [optional] 
**PrecedingTransaction** | Pointer to [**PrecedingTransaction**](PrecedingTransaction.md) |  | [optional] 
**Program** | Pointer to [**Program**](Program.md) |  | [optional] 
**ProgramTransfer** | Pointer to [**ProgramTransferResponse**](ProgramTransferResponse.md) |  | [optional] 
**RealTimeFeeGroup** | Pointer to [**RealTimeFeeGroup**](RealTimeFeeGroup.md) |  | [optional] 
**RequestAmount** | Pointer to **float32** | Merchant-requested amount, including any fees. | [optional] 
**Response** | Pointer to [**Response**](Response.md) |  | [optional] 
**SettlementDate** | Pointer to **time.Time** | Date and time when funds were moved for a transaction, in UTC. For example, in the case of a refund, when funds were credited to the cardholder. | [optional] 
**StandinApprovedBy** | Pointer to **string** | Indicates which party approved a transaction: the card network using stand-in processing, or Marqeta using Commando Mode. Returned only when a transaction is approved. | [optional] 
**StandinBy** | Pointer to **string** | Indicates which party approved a transaction: the card network using stand-in processing, or Marqeta using Commando Mode. | [optional] 
**StandinReason** | Pointer to **string** | Indicates why the card network handled a transaction requiring stand-in processing. | [optional] 
**State** | **string** | Current state of the transaction. For more information about the &#x60;state&#x60; field, see &lt;&lt;/developer-guides/about-transactions#_the_transaction_lifecycle, The transaction lifecycle&gt;&gt;. | 
**Store** | Pointer to [**StoreResponseModel**](StoreResponseModel.md) |  | [optional] 
**Subnetwork** | Pointer to **string** | Indicates which subnetwork was used to complete the transaction. Possible values include the following:  * *VISANET*  Used for VisaNet signature-based transactions. * *VISANETDEBIT*  Used for VisaNet Debit PIN-based transaction. * *VISAINTERLINK*  Used for Visa Interlink PIN-based transactions. * *VISAPLUS*  Used for ATM withdrawals on Visa. * *MAESTRO*  Used for PIN-based transactions on Mastercard. * *CIRRUS*  Used for ATM withdrawals on Mastercard. * *MASTERCARDDEBIT*  Used for signature-based transactions on Mastercard. * *GATEWAY_JIT*  Used for Gateway JIT Funding transactions. * *MANAGED_JIT*  Used for Managed JIT Funding transactions or for transactions that occur while Commando Mode is enabled. | [optional] 
**Token** | **string** | Unique identifier of the transaction, formatted as a UUID.  *NOTE:* For subsequent related transactions, this token value appears as the &#x60;preceding_related_transaction_token&#x60;. | [readonly] 
**TransactionAttributes** | Pointer to **map[string]string** | Additional transaction attributes. | [optional] 
**TransactionMetadata** | Pointer to [**TransactionMetadata**](TransactionMetadata.md) |  | [optional] 
**Type** | **string** | Transaction event type. For more information about the &#x60;type&#x60; field, see &lt;&lt;/core-api/event-types#_transaction_events, Transaction events&gt;&gt;. | [readonly] 
**User** | Pointer to [**CardholderMetadata**](CardholderMetadata.md) |  | [optional] 
**UserToken** | Pointer to **string** | Unique identifier of the user who owns the account that funded the transaction; subsequent related transactions retain the same &#x60;user_token&#x60;, even if the card used to complete the transaction moves to another user. | [optional] 
**UserTransactionTime** | Pointer to **time.Time** | Date and time when the user initiated the transaction, in UTC. For example, when a merchant performed the original authorization for a refund. | [optional] 

## Methods

### NewTransactionModel

`func NewTransactionModel(actingUserToken string, amount float32, state string, token string, type_ string, ) *TransactionModel`

NewTransactionModel instantiates a new TransactionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionModelWithDefaults

`func NewTransactionModelWithDefaults() *TransactionModel`

NewTransactionModelWithDefaults instantiates a new TransactionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountFunding

`func (o *TransactionModel) GetAccountFunding() AccountFunding`

GetAccountFunding returns the AccountFunding field if non-nil, zero value otherwise.

### GetAccountFundingOk

`func (o *TransactionModel) GetAccountFundingOk() (*AccountFunding, bool)`

GetAccountFundingOk returns a tuple with the AccountFunding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFunding

`func (o *TransactionModel) SetAccountFunding(v AccountFunding)`

SetAccountFunding sets AccountFunding field to given value.

### HasAccountFunding

`func (o *TransactionModel) HasAccountFunding() bool`

HasAccountFunding returns a boolean if a field has been set.

### GetAcquirer

`func (o *TransactionModel) GetAcquirer() Acquirer`

GetAcquirer returns the Acquirer field if non-nil, zero value otherwise.

### GetAcquirerOk

`func (o *TransactionModel) GetAcquirerOk() (*Acquirer, bool)`

GetAcquirerOk returns a tuple with the Acquirer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirer

`func (o *TransactionModel) SetAcquirer(v Acquirer)`

SetAcquirer sets Acquirer field to given value.

### HasAcquirer

`func (o *TransactionModel) HasAcquirer() bool`

HasAcquirer returns a boolean if a field has been set.

### GetAcquirerFeeAmount

`func (o *TransactionModel) GetAcquirerFeeAmount() float32`

GetAcquirerFeeAmount returns the AcquirerFeeAmount field if non-nil, zero value otherwise.

### GetAcquirerFeeAmountOk

`func (o *TransactionModel) GetAcquirerFeeAmountOk() (*float32, bool)`

GetAcquirerFeeAmountOk returns a tuple with the AcquirerFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerFeeAmount

`func (o *TransactionModel) SetAcquirerFeeAmount(v float32)`

SetAcquirerFeeAmount sets AcquirerFeeAmount field to given value.

### HasAcquirerFeeAmount

`func (o *TransactionModel) HasAcquirerFeeAmount() bool`

HasAcquirerFeeAmount returns a boolean if a field has been set.

### GetAcquirerReferenceData

`func (o *TransactionModel) GetAcquirerReferenceData() string`

GetAcquirerReferenceData returns the AcquirerReferenceData field if non-nil, zero value otherwise.

### GetAcquirerReferenceDataOk

`func (o *TransactionModel) GetAcquirerReferenceDataOk() (*string, bool)`

GetAcquirerReferenceDataOk returns a tuple with the AcquirerReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceData

`func (o *TransactionModel) SetAcquirerReferenceData(v string)`

SetAcquirerReferenceData sets AcquirerReferenceData field to given value.

### HasAcquirerReferenceData

`func (o *TransactionModel) HasAcquirerReferenceData() bool`

HasAcquirerReferenceData returns a boolean if a field has been set.

### GetAcquirerReferenceId

`func (o *TransactionModel) GetAcquirerReferenceId() string`

GetAcquirerReferenceId returns the AcquirerReferenceId field if non-nil, zero value otherwise.

### GetAcquirerReferenceIdOk

`func (o *TransactionModel) GetAcquirerReferenceIdOk() (*string, bool)`

GetAcquirerReferenceIdOk returns a tuple with the AcquirerReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirerReferenceId

`func (o *TransactionModel) SetAcquirerReferenceId(v string)`

SetAcquirerReferenceId sets AcquirerReferenceId field to given value.

### HasAcquirerReferenceId

`func (o *TransactionModel) HasAcquirerReferenceId() bool`

HasAcquirerReferenceId returns a boolean if a field has been set.

### GetActingUserToken

`func (o *TransactionModel) GetActingUserToken() string`

GetActingUserToken returns the ActingUserToken field if non-nil, zero value otherwise.

### GetActingUserTokenOk

`func (o *TransactionModel) GetActingUserTokenOk() (*string, bool)`

GetActingUserTokenOk returns a tuple with the ActingUserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserToken

`func (o *TransactionModel) SetActingUserToken(v string)`

SetActingUserToken sets ActingUserToken field to given value.


### GetAddressVerification

`func (o *TransactionModel) GetAddressVerification() AddressVerificationModel`

GetAddressVerification returns the AddressVerification field if non-nil, zero value otherwise.

### GetAddressVerificationOk

`func (o *TransactionModel) GetAddressVerificationOk() (*AddressVerificationModel, bool)`

GetAddressVerificationOk returns a tuple with the AddressVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerification

`func (o *TransactionModel) SetAddressVerification(v AddressVerificationModel)`

SetAddressVerification sets AddressVerification field to given value.

### HasAddressVerification

`func (o *TransactionModel) HasAddressVerification() bool`

HasAddressVerification returns a boolean if a field has been set.

### GetAdviceReasonCode

`func (o *TransactionModel) GetAdviceReasonCode() string`

GetAdviceReasonCode returns the AdviceReasonCode field if non-nil, zero value otherwise.

### GetAdviceReasonCodeOk

`func (o *TransactionModel) GetAdviceReasonCodeOk() (*string, bool)`

GetAdviceReasonCodeOk returns a tuple with the AdviceReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdviceReasonCode

`func (o *TransactionModel) SetAdviceReasonCode(v string)`

SetAdviceReasonCode sets AdviceReasonCode field to given value.

### HasAdviceReasonCode

`func (o *TransactionModel) HasAdviceReasonCode() bool`

HasAdviceReasonCode returns a boolean if a field has been set.

### GetAdviceReasonDetails

`func (o *TransactionModel) GetAdviceReasonDetails() string`

GetAdviceReasonDetails returns the AdviceReasonDetails field if non-nil, zero value otherwise.

### GetAdviceReasonDetailsOk

`func (o *TransactionModel) GetAdviceReasonDetailsOk() (*string, bool)`

GetAdviceReasonDetailsOk returns a tuple with the AdviceReasonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdviceReasonDetails

`func (o *TransactionModel) SetAdviceReasonDetails(v string)`

SetAdviceReasonDetails sets AdviceReasonDetails field to given value.

### HasAdviceReasonDetails

`func (o *TransactionModel) HasAdviceReasonDetails() bool`

HasAdviceReasonDetails returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionModel) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionModel) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionModel) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAmountToBeReleased

`func (o *TransactionModel) GetAmountToBeReleased() float32`

GetAmountToBeReleased returns the AmountToBeReleased field if non-nil, zero value otherwise.

### GetAmountToBeReleasedOk

`func (o *TransactionModel) GetAmountToBeReleasedOk() (*float32, bool)`

GetAmountToBeReleasedOk returns a tuple with the AmountToBeReleased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToBeReleased

`func (o *TransactionModel) SetAmountToBeReleased(v float32)`

SetAmountToBeReleased sets AmountToBeReleased field to given value.

### HasAmountToBeReleased

`func (o *TransactionModel) HasAmountToBeReleased() bool`

HasAmountToBeReleased returns a boolean if a field has been set.

### GetApprovalCode

`func (o *TransactionModel) GetApprovalCode() string`

GetApprovalCode returns the ApprovalCode field if non-nil, zero value otherwise.

### GetApprovalCodeOk

`func (o *TransactionModel) GetApprovalCodeOk() (*string, bool)`

GetApprovalCodeOk returns a tuple with the ApprovalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalCode

`func (o *TransactionModel) SetApprovalCode(v string)`

SetApprovalCode sets ApprovalCode field to given value.

### HasApprovalCode

`func (o *TransactionModel) HasApprovalCode() bool`

HasApprovalCode returns a boolean if a field has been set.

### GetAtcInformation

`func (o *TransactionModel) GetAtcInformation() AtcInformation`

GetAtcInformation returns the AtcInformation field if non-nil, zero value otherwise.

### GetAtcInformationOk

`func (o *TransactionModel) GetAtcInformationOk() (*AtcInformation, bool)`

GetAtcInformationOk returns a tuple with the AtcInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtcInformation

`func (o *TransactionModel) SetAtcInformation(v AtcInformation)`

SetAtcInformation sets AtcInformation field to given value.

### HasAtcInformation

`func (o *TransactionModel) HasAtcInformation() bool`

HasAtcInformation returns a boolean if a field has been set.

### GetAutoReload

`func (o *TransactionModel) GetAutoReload() AutoReloadModel`

GetAutoReload returns the AutoReload field if non-nil, zero value otherwise.

### GetAutoReloadOk

`func (o *TransactionModel) GetAutoReloadOk() (*AutoReloadModel, bool)`

GetAutoReloadOk returns a tuple with the AutoReload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoReload

`func (o *TransactionModel) SetAutoReload(v AutoReloadModel)`

SetAutoReload sets AutoReload field to given value.

### HasAutoReload

`func (o *TransactionModel) HasAutoReload() bool`

HasAutoReload returns a boolean if a field has been set.

### GetBankTransferToken

`func (o *TransactionModel) GetBankTransferToken() string`

GetBankTransferToken returns the BankTransferToken field if non-nil, zero value otherwise.

### GetBankTransferTokenOk

`func (o *TransactionModel) GetBankTransferTokenOk() (*string, bool)`

GetBankTransferTokenOk returns a tuple with the BankTransferToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferToken

`func (o *TransactionModel) SetBankTransferToken(v string)`

SetBankTransferToken sets BankTransferToken field to given value.

### HasBankTransferToken

`func (o *TransactionModel) HasBankTransferToken() bool`

HasBankTransferToken returns a boolean if a field has been set.

### GetBatchNumber

`func (o *TransactionModel) GetBatchNumber() string`

GetBatchNumber returns the BatchNumber field if non-nil, zero value otherwise.

### GetBatchNumberOk

`func (o *TransactionModel) GetBatchNumberOk() (*string, bool)`

GetBatchNumberOk returns a tuple with the BatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchNumber

`func (o *TransactionModel) SetBatchNumber(v string)`

SetBatchNumber sets BatchNumber field to given value.

### HasBatchNumber

`func (o *TransactionModel) HasBatchNumber() bool`

HasBatchNumber returns a boolean if a field has been set.

### GetBusiness

`func (o *TransactionModel) GetBusiness() BusinessMetadata`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *TransactionModel) GetBusinessOk() (*BusinessMetadata, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *TransactionModel) SetBusiness(v BusinessMetadata)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *TransactionModel) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetBusinessToken

`func (o *TransactionModel) GetBusinessToken() string`

GetBusinessToken returns the BusinessToken field if non-nil, zero value otherwise.

### GetBusinessTokenOk

`func (o *TransactionModel) GetBusinessTokenOk() (*string, bool)`

GetBusinessTokenOk returns a tuple with the BusinessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessToken

`func (o *TransactionModel) SetBusinessToken(v string)`

SetBusinessToken sets BusinessToken field to given value.

### HasBusinessToken

`func (o *TransactionModel) HasBusinessToken() bool`

HasBusinessToken returns a boolean if a field has been set.

### GetCard

`func (o *TransactionModel) GetCard() CardResponse`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *TransactionModel) GetCardOk() (*CardResponse, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *TransactionModel) SetCard(v CardResponse)`

SetCard sets Card field to given value.

### HasCard

`func (o *TransactionModel) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCardAcceptor

`func (o *TransactionModel) GetCardAcceptor() TransactionCardAcceptor`

GetCardAcceptor returns the CardAcceptor field if non-nil, zero value otherwise.

### GetCardAcceptorOk

`func (o *TransactionModel) GetCardAcceptorOk() (*TransactionCardAcceptor, bool)`

GetCardAcceptorOk returns a tuple with the CardAcceptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptor

`func (o *TransactionModel) SetCardAcceptor(v TransactionCardAcceptor)`

SetCardAcceptor sets CardAcceptor field to given value.

### HasCardAcceptor

`func (o *TransactionModel) HasCardAcceptor() bool`

HasCardAcceptor returns a boolean if a field has been set.

### GetCardHolderModel

`func (o *TransactionModel) GetCardHolderModel() UserCardHolderResponse`

GetCardHolderModel returns the CardHolderModel field if non-nil, zero value otherwise.

### GetCardHolderModelOk

`func (o *TransactionModel) GetCardHolderModelOk() (*UserCardHolderResponse, bool)`

GetCardHolderModelOk returns a tuple with the CardHolderModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHolderModel

`func (o *TransactionModel) SetCardHolderModel(v UserCardHolderResponse)`

SetCardHolderModel sets CardHolderModel field to given value.

### HasCardHolderModel

`func (o *TransactionModel) HasCardHolderModel() bool`

HasCardHolderModel returns a boolean if a field has been set.

### GetCardProductToken

`func (o *TransactionModel) GetCardProductToken() string`

GetCardProductToken returns the CardProductToken field if non-nil, zero value otherwise.

### GetCardProductTokenOk

`func (o *TransactionModel) GetCardProductTokenOk() (*string, bool)`

GetCardProductTokenOk returns a tuple with the CardProductToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardProductToken

`func (o *TransactionModel) SetCardProductToken(v string)`

SetCardProductToken sets CardProductToken field to given value.

### HasCardProductToken

`func (o *TransactionModel) HasCardProductToken() bool`

HasCardProductToken returns a boolean if a field has been set.

### GetCardSecurityCodeVerification

`func (o *TransactionModel) GetCardSecurityCodeVerification() CardSecurityCodeVerification`

GetCardSecurityCodeVerification returns the CardSecurityCodeVerification field if non-nil, zero value otherwise.

### GetCardSecurityCodeVerificationOk

`func (o *TransactionModel) GetCardSecurityCodeVerificationOk() (*CardSecurityCodeVerification, bool)`

GetCardSecurityCodeVerificationOk returns a tuple with the CardSecurityCodeVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSecurityCodeVerification

`func (o *TransactionModel) SetCardSecurityCodeVerification(v CardSecurityCodeVerification)`

SetCardSecurityCodeVerification sets CardSecurityCodeVerification field to given value.

### HasCardSecurityCodeVerification

`func (o *TransactionModel) HasCardSecurityCodeVerification() bool`

HasCardSecurityCodeVerification returns a boolean if a field has been set.

### GetCardToken

`func (o *TransactionModel) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *TransactionModel) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *TransactionModel) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.

### HasCardToken

`func (o *TransactionModel) HasCardToken() bool`

HasCardToken returns a boolean if a field has been set.

### GetCardholderAuthenticationData

`func (o *TransactionModel) GetCardholderAuthenticationData() CardholderAuthenticationData`

GetCardholderAuthenticationData returns the CardholderAuthenticationData field if non-nil, zero value otherwise.

### GetCardholderAuthenticationDataOk

`func (o *TransactionModel) GetCardholderAuthenticationDataOk() (*CardholderAuthenticationData, bool)`

GetCardholderAuthenticationDataOk returns a tuple with the CardholderAuthenticationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderAuthenticationData

`func (o *TransactionModel) SetCardholderAuthenticationData(v CardholderAuthenticationData)`

SetCardholderAuthenticationData sets CardholderAuthenticationData field to given value.

### HasCardholderAuthenticationData

`func (o *TransactionModel) HasCardholderAuthenticationData() bool`

HasCardholderAuthenticationData returns a boolean if a field has been set.

### GetCashBackAmount

`func (o *TransactionModel) GetCashBackAmount() float32`

GetCashBackAmount returns the CashBackAmount field if non-nil, zero value otherwise.

### GetCashBackAmountOk

`func (o *TransactionModel) GetCashBackAmountOk() (*float32, bool)`

GetCashBackAmountOk returns a tuple with the CashBackAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBackAmount

`func (o *TransactionModel) SetCashBackAmount(v float32)`

SetCashBackAmount sets CashBackAmount field to given value.

### HasCashBackAmount

`func (o *TransactionModel) HasCashBackAmount() bool`

HasCashBackAmount returns a boolean if a field has been set.

### GetChargeback

`func (o *TransactionModel) GetChargeback() ChargebackResponse`

GetChargeback returns the Chargeback field if non-nil, zero value otherwise.

### GetChargebackOk

`func (o *TransactionModel) GetChargebackOk() (*ChargebackResponse, bool)`

GetChargebackOk returns a tuple with the Chargeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeback

`func (o *TransactionModel) SetChargeback(v ChargebackResponse)`

SetChargeback sets Chargeback field to given value.

### HasChargeback

`func (o *TransactionModel) HasChargeback() bool`

HasChargeback returns a boolean if a field has been set.

### GetClearingRecordSequenceNumber

`func (o *TransactionModel) GetClearingRecordSequenceNumber() string`

GetClearingRecordSequenceNumber returns the ClearingRecordSequenceNumber field if non-nil, zero value otherwise.

### GetClearingRecordSequenceNumberOk

`func (o *TransactionModel) GetClearingRecordSequenceNumberOk() (*string, bool)`

GetClearingRecordSequenceNumberOk returns a tuple with the ClearingRecordSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingRecordSequenceNumber

`func (o *TransactionModel) SetClearingRecordSequenceNumber(v string)`

SetClearingRecordSequenceNumber sets ClearingRecordSequenceNumber field to given value.

### HasClearingRecordSequenceNumber

`func (o *TransactionModel) HasClearingRecordSequenceNumber() bool`

HasClearingRecordSequenceNumber returns a boolean if a field has been set.

### GetCreatedTime

`func (o *TransactionModel) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *TransactionModel) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *TransactionModel) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *TransactionModel) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *TransactionModel) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *TransactionModel) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *TransactionModel) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *TransactionModel) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetCurrencyConversion

`func (o *TransactionModel) GetCurrencyConversion() CurrencyConversion`

GetCurrencyConversion returns the CurrencyConversion field if non-nil, zero value otherwise.

### GetCurrencyConversionOk

`func (o *TransactionModel) GetCurrencyConversionOk() (*CurrencyConversion, bool)`

GetCurrencyConversionOk returns a tuple with the CurrencyConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyConversion

`func (o *TransactionModel) SetCurrencyConversion(v CurrencyConversion)`

SetCurrencyConversion sets CurrencyConversion field to given value.

### HasCurrencyConversion

`func (o *TransactionModel) HasCurrencyConversion() bool`

HasCurrencyConversion returns a boolean if a field has been set.

### GetDeferredSettlementDays

`func (o *TransactionModel) GetDeferredSettlementDays() string`

GetDeferredSettlementDays returns the DeferredSettlementDays field if non-nil, zero value otherwise.

### GetDeferredSettlementDaysOk

`func (o *TransactionModel) GetDeferredSettlementDaysOk() (*string, bool)`

GetDeferredSettlementDaysOk returns a tuple with the DeferredSettlementDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferredSettlementDays

`func (o *TransactionModel) SetDeferredSettlementDays(v string)`

SetDeferredSettlementDays sets DeferredSettlementDays field to given value.

### HasDeferredSettlementDays

`func (o *TransactionModel) HasDeferredSettlementDays() bool`

HasDeferredSettlementDays returns a boolean if a field has been set.

### GetDigitalWalletToken

`func (o *TransactionModel) GetDigitalWalletToken() DigitalWalletToken`

GetDigitalWalletToken returns the DigitalWalletToken field if non-nil, zero value otherwise.

### GetDigitalWalletTokenOk

`func (o *TransactionModel) GetDigitalWalletTokenOk() (*DigitalWalletToken, bool)`

GetDigitalWalletTokenOk returns a tuple with the DigitalWalletToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletToken

`func (o *TransactionModel) SetDigitalWalletToken(v DigitalWalletToken)`

SetDigitalWalletToken sets DigitalWalletToken field to given value.

### HasDigitalWalletToken

`func (o *TransactionModel) HasDigitalWalletToken() bool`

HasDigitalWalletToken returns a boolean if a field has been set.

### GetDispute

`func (o *TransactionModel) GetDispute() DisputeModel`

GetDispute returns the Dispute field if non-nil, zero value otherwise.

### GetDisputeOk

`func (o *TransactionModel) GetDisputeOk() (*DisputeModel, bool)`

GetDisputeOk returns a tuple with the Dispute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispute

`func (o *TransactionModel) SetDispute(v DisputeModel)`

SetDispute sets Dispute field to given value.

### HasDispute

`func (o *TransactionModel) HasDispute() bool`

HasDispute returns a boolean if a field has been set.

### GetDuration

`func (o *TransactionModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TransactionModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TransactionModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TransactionModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnhancedDataToken

`func (o *TransactionModel) GetEnhancedDataToken() string`

GetEnhancedDataToken returns the EnhancedDataToken field if non-nil, zero value otherwise.

### GetEnhancedDataTokenOk

`func (o *TransactionModel) GetEnhancedDataTokenOk() (*string, bool)`

GetEnhancedDataTokenOk returns a tuple with the EnhancedDataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedDataToken

`func (o *TransactionModel) SetEnhancedDataToken(v string)`

SetEnhancedDataToken sets EnhancedDataToken field to given value.

### HasEnhancedDataToken

`func (o *TransactionModel) HasEnhancedDataToken() bool`

HasEnhancedDataToken returns a boolean if a field has been set.

### GetFee

`func (o *TransactionModel) GetFee() Fee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionModel) GetFeeOk() (*Fee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionModel) SetFee(v Fee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *TransactionModel) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeTransfer

`func (o *TransactionModel) GetFeeTransfer() FeeTransferResponse`

GetFeeTransfer returns the FeeTransfer field if non-nil, zero value otherwise.

### GetFeeTransferOk

`func (o *TransactionModel) GetFeeTransferOk() (*FeeTransferResponse, bool)`

GetFeeTransferOk returns a tuple with the FeeTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTransfer

`func (o *TransactionModel) SetFeeTransfer(v FeeTransferResponse)`

SetFeeTransfer sets FeeTransfer field to given value.

### HasFeeTransfer

`func (o *TransactionModel) HasFeeTransfer() bool`

HasFeeTransfer returns a boolean if a field has been set.

### GetFees

`func (o *TransactionModel) GetFees() []NetworkFeeModel`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *TransactionModel) GetFeesOk() (*[]NetworkFeeModel, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *TransactionModel) SetFees(v []NetworkFeeModel)`

SetFees sets Fees field to given value.

### HasFees

`func (o *TransactionModel) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetFraud

`func (o *TransactionModel) GetFraud() FraudView`

GetFraud returns the Fraud field if non-nil, zero value otherwise.

### GetFraudOk

`func (o *TransactionModel) GetFraudOk() (*FraudView, bool)`

GetFraudOk returns a tuple with the Fraud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraud

`func (o *TransactionModel) SetFraud(v FraudView)`

SetFraud sets Fraud field to given value.

### HasFraud

`func (o *TransactionModel) HasFraud() bool`

HasFraud returns a boolean if a field has been set.

### GetFromAccount

`func (o *TransactionModel) GetFromAccount() string`

GetFromAccount returns the FromAccount field if non-nil, zero value otherwise.

### GetFromAccountOk

`func (o *TransactionModel) GetFromAccountOk() (*string, bool)`

GetFromAccountOk returns a tuple with the FromAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAccount

`func (o *TransactionModel) SetFromAccount(v string)`

SetFromAccount sets FromAccount field to given value.

### HasFromAccount

`func (o *TransactionModel) HasFromAccount() bool`

HasFromAccount returns a boolean if a field has been set.

### GetGpa

`func (o *TransactionModel) GetGpa() CardholderBalance`

GetGpa returns the Gpa field if non-nil, zero value otherwise.

### GetGpaOk

`func (o *TransactionModel) GetGpaOk() (*CardholderBalance, bool)`

GetGpaOk returns a tuple with the Gpa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpa

`func (o *TransactionModel) SetGpa(v CardholderBalance)`

SetGpa sets Gpa field to given value.

### HasGpa

`func (o *TransactionModel) HasGpa() bool`

HasGpa returns a boolean if a field has been set.

### GetGpaOrder

`func (o *TransactionModel) GetGpaOrder() GpaResponse`

GetGpaOrder returns the GpaOrder field if non-nil, zero value otherwise.

### GetGpaOrderOk

`func (o *TransactionModel) GetGpaOrderOk() (*GpaResponse, bool)`

GetGpaOrderOk returns a tuple with the GpaOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpaOrder

`func (o *TransactionModel) SetGpaOrder(v GpaResponse)`

SetGpaOrder sets GpaOrder field to given value.

### HasGpaOrder

`func (o *TransactionModel) HasGpaOrder() bool`

HasGpaOrder returns a boolean if a field has been set.

### GetGpaOrderUnload

`func (o *TransactionModel) GetGpaOrderUnload() GpaReturns`

GetGpaOrderUnload returns the GpaOrderUnload field if non-nil, zero value otherwise.

### GetGpaOrderUnloadOk

`func (o *TransactionModel) GetGpaOrderUnloadOk() (*GpaReturns, bool)`

GetGpaOrderUnloadOk returns a tuple with the GpaOrderUnload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpaOrderUnload

`func (o *TransactionModel) SetGpaOrderUnload(v GpaReturns)`

SetGpaOrderUnload sets GpaOrderUnload field to given value.

### HasGpaOrderUnload

`func (o *TransactionModel) HasGpaOrderUnload() bool`

HasGpaOrderUnload returns a boolean if a field has been set.

### GetIdentifier

`func (o *TransactionModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TransactionModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TransactionModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *TransactionModel) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIncrementalAuthorizationTransactionTokens

`func (o *TransactionModel) GetIncrementalAuthorizationTransactionTokens() []string`

GetIncrementalAuthorizationTransactionTokens returns the IncrementalAuthorizationTransactionTokens field if non-nil, zero value otherwise.

### GetIncrementalAuthorizationTransactionTokensOk

`func (o *TransactionModel) GetIncrementalAuthorizationTransactionTokensOk() (*[]string, bool)`

GetIncrementalAuthorizationTransactionTokensOk returns a tuple with the IncrementalAuthorizationTransactionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalAuthorizationTransactionTokens

`func (o *TransactionModel) SetIncrementalAuthorizationTransactionTokens(v []string)`

SetIncrementalAuthorizationTransactionTokens sets IncrementalAuthorizationTransactionTokens field to given value.

### HasIncrementalAuthorizationTransactionTokens

`func (o *TransactionModel) HasIncrementalAuthorizationTransactionTokens() bool`

HasIncrementalAuthorizationTransactionTokens returns a boolean if a field has been set.

### GetInterchangeRateDescriptor

`func (o *TransactionModel) GetInterchangeRateDescriptor() string`

GetInterchangeRateDescriptor returns the InterchangeRateDescriptor field if non-nil, zero value otherwise.

### GetInterchangeRateDescriptorOk

`func (o *TransactionModel) GetInterchangeRateDescriptorOk() (*string, bool)`

GetInterchangeRateDescriptorOk returns a tuple with the InterchangeRateDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterchangeRateDescriptor

`func (o *TransactionModel) SetInterchangeRateDescriptor(v string)`

SetInterchangeRateDescriptor sets InterchangeRateDescriptor field to given value.

### HasInterchangeRateDescriptor

`func (o *TransactionModel) HasInterchangeRateDescriptor() bool`

HasInterchangeRateDescriptor returns a boolean if a field has been set.

### GetIsPreauthorization

`func (o *TransactionModel) GetIsPreauthorization() bool`

GetIsPreauthorization returns the IsPreauthorization field if non-nil, zero value otherwise.

### GetIsPreauthorizationOk

`func (o *TransactionModel) GetIsPreauthorizationOk() (*bool, bool)`

GetIsPreauthorizationOk returns a tuple with the IsPreauthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreauthorization

`func (o *TransactionModel) SetIsPreauthorization(v bool)`

SetIsPreauthorization sets IsPreauthorization field to given value.

### HasIsPreauthorization

`func (o *TransactionModel) HasIsPreauthorization() bool`

HasIsPreauthorization returns a boolean if a field has been set.

### GetIsaIndicator

`func (o *TransactionModel) GetIsaIndicator() string`

GetIsaIndicator returns the IsaIndicator field if non-nil, zero value otherwise.

### GetIsaIndicatorOk

`func (o *TransactionModel) GetIsaIndicatorOk() (*string, bool)`

GetIsaIndicatorOk returns a tuple with the IsaIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsaIndicator

`func (o *TransactionModel) SetIsaIndicator(v string)`

SetIsaIndicator sets IsaIndicator field to given value.

### HasIsaIndicator

`func (o *TransactionModel) HasIsaIndicator() bool`

HasIsaIndicator returns a boolean if a field has been set.

### GetIssuerInterchangeAmount

`func (o *TransactionModel) GetIssuerInterchangeAmount() float32`

GetIssuerInterchangeAmount returns the IssuerInterchangeAmount field if non-nil, zero value otherwise.

### GetIssuerInterchangeAmountOk

`func (o *TransactionModel) GetIssuerInterchangeAmountOk() (*float32, bool)`

GetIssuerInterchangeAmountOk returns a tuple with the IssuerInterchangeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerInterchangeAmount

`func (o *TransactionModel) SetIssuerInterchangeAmount(v float32)`

SetIssuerInterchangeAmount sets IssuerInterchangeAmount field to given value.

### HasIssuerInterchangeAmount

`func (o *TransactionModel) HasIssuerInterchangeAmount() bool`

HasIssuerInterchangeAmount returns a boolean if a field has been set.

### GetIssuerPaymentNode

`func (o *TransactionModel) GetIssuerPaymentNode() string`

GetIssuerPaymentNode returns the IssuerPaymentNode field if non-nil, zero value otherwise.

### GetIssuerPaymentNodeOk

`func (o *TransactionModel) GetIssuerPaymentNodeOk() (*string, bool)`

GetIssuerPaymentNodeOk returns a tuple with the IssuerPaymentNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerPaymentNode

`func (o *TransactionModel) SetIssuerPaymentNode(v string)`

SetIssuerPaymentNode sets IssuerPaymentNode field to given value.

### HasIssuerPaymentNode

`func (o *TransactionModel) HasIssuerPaymentNode() bool`

HasIssuerPaymentNode returns a boolean if a field has been set.

### GetIssuerReceivedTime

`func (o *TransactionModel) GetIssuerReceivedTime() string`

GetIssuerReceivedTime returns the IssuerReceivedTime field if non-nil, zero value otherwise.

### GetIssuerReceivedTimeOk

`func (o *TransactionModel) GetIssuerReceivedTimeOk() (*string, bool)`

GetIssuerReceivedTimeOk returns a tuple with the IssuerReceivedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerReceivedTime

`func (o *TransactionModel) SetIssuerReceivedTime(v string)`

SetIssuerReceivedTime sets IssuerReceivedTime field to given value.

### HasIssuerReceivedTime

`func (o *TransactionModel) HasIssuerReceivedTime() bool`

HasIssuerReceivedTime returns a boolean if a field has been set.

### GetLocalTransactionDate

`func (o *TransactionModel) GetLocalTransactionDate() time.Time`

GetLocalTransactionDate returns the LocalTransactionDate field if non-nil, zero value otherwise.

### GetLocalTransactionDateOk

`func (o *TransactionModel) GetLocalTransactionDateOk() (*time.Time, bool)`

GetLocalTransactionDateOk returns a tuple with the LocalTransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTransactionDate

`func (o *TransactionModel) SetLocalTransactionDate(v time.Time)`

SetLocalTransactionDate sets LocalTransactionDate field to given value.

### HasLocalTransactionDate

`func (o *TransactionModel) HasLocalTransactionDate() bool`

HasLocalTransactionDate returns a boolean if a field has been set.

### GetMerchant

`func (o *TransactionModel) GetMerchant() MerchantResponseModel`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *TransactionModel) GetMerchantOk() (*MerchantResponseModel, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *TransactionModel) SetMerchant(v MerchantResponseModel)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *TransactionModel) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetMsaOrderUnload

`func (o *TransactionModel) GetMsaOrderUnload() MsaReturns`

GetMsaOrderUnload returns the MsaOrderUnload field if non-nil, zero value otherwise.

### GetMsaOrderUnloadOk

`func (o *TransactionModel) GetMsaOrderUnloadOk() (*MsaReturns, bool)`

GetMsaOrderUnloadOk returns a tuple with the MsaOrderUnload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsaOrderUnload

`func (o *TransactionModel) SetMsaOrderUnload(v MsaReturns)`

SetMsaOrderUnload sets MsaOrderUnload field to given value.

### HasMsaOrderUnload

`func (o *TransactionModel) HasMsaOrderUnload() bool`

HasMsaOrderUnload returns a boolean if a field has been set.

### GetMultiClearingSequenceCount

`func (o *TransactionModel) GetMultiClearingSequenceCount() string`

GetMultiClearingSequenceCount returns the MultiClearingSequenceCount field if non-nil, zero value otherwise.

### GetMultiClearingSequenceCountOk

`func (o *TransactionModel) GetMultiClearingSequenceCountOk() (*string, bool)`

GetMultiClearingSequenceCountOk returns a tuple with the MultiClearingSequenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiClearingSequenceCount

`func (o *TransactionModel) SetMultiClearingSequenceCount(v string)`

SetMultiClearingSequenceCount sets MultiClearingSequenceCount field to given value.

### HasMultiClearingSequenceCount

`func (o *TransactionModel) HasMultiClearingSequenceCount() bool`

HasMultiClearingSequenceCount returns a boolean if a field has been set.

### GetMultiClearingSequenceNumber

`func (o *TransactionModel) GetMultiClearingSequenceNumber() string`

GetMultiClearingSequenceNumber returns the MultiClearingSequenceNumber field if non-nil, zero value otherwise.

### GetMultiClearingSequenceNumberOk

`func (o *TransactionModel) GetMultiClearingSequenceNumberOk() (*string, bool)`

GetMultiClearingSequenceNumberOk returns a tuple with the MultiClearingSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiClearingSequenceNumber

`func (o *TransactionModel) SetMultiClearingSequenceNumber(v string)`

SetMultiClearingSequenceNumber sets MultiClearingSequenceNumber field to given value.

### HasMultiClearingSequenceNumber

`func (o *TransactionModel) HasMultiClearingSequenceNumber() bool`

HasMultiClearingSequenceNumber returns a boolean if a field has been set.

### GetNationalNetCpdOfOriginal

`func (o *TransactionModel) GetNationalNetCpdOfOriginal() string`

GetNationalNetCpdOfOriginal returns the NationalNetCpdOfOriginal field if non-nil, zero value otherwise.

### GetNationalNetCpdOfOriginalOk

`func (o *TransactionModel) GetNationalNetCpdOfOriginalOk() (*string, bool)`

GetNationalNetCpdOfOriginalOk returns a tuple with the NationalNetCpdOfOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalNetCpdOfOriginal

`func (o *TransactionModel) SetNationalNetCpdOfOriginal(v string)`

SetNationalNetCpdOfOriginal sets NationalNetCpdOfOriginal field to given value.

### HasNationalNetCpdOfOriginal

`func (o *TransactionModel) HasNationalNetCpdOfOriginal() bool`

HasNationalNetCpdOfOriginal returns a boolean if a field has been set.

### GetNetwork

`func (o *TransactionModel) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionModel) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionModel) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *TransactionModel) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkMetadata

`func (o *TransactionModel) GetNetworkMetadata() NetworkMetadata`

GetNetworkMetadata returns the NetworkMetadata field if non-nil, zero value otherwise.

### GetNetworkMetadataOk

`func (o *TransactionModel) GetNetworkMetadataOk() (*NetworkMetadata, bool)`

GetNetworkMetadataOk returns a tuple with the NetworkMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMetadata

`func (o *TransactionModel) SetNetworkMetadata(v NetworkMetadata)`

SetNetworkMetadata sets NetworkMetadata field to given value.

### HasNetworkMetadata

`func (o *TransactionModel) HasNetworkMetadata() bool`

HasNetworkMetadata returns a boolean if a field has been set.

### GetNetworkReferenceId

`func (o *TransactionModel) GetNetworkReferenceId() string`

GetNetworkReferenceId returns the NetworkReferenceId field if non-nil, zero value otherwise.

### GetNetworkReferenceIdOk

`func (o *TransactionModel) GetNetworkReferenceIdOk() (*string, bool)`

GetNetworkReferenceIdOk returns a tuple with the NetworkReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkReferenceId

`func (o *TransactionModel) SetNetworkReferenceId(v string)`

SetNetworkReferenceId sets NetworkReferenceId field to given value.

### HasNetworkReferenceId

`func (o *TransactionModel) HasNetworkReferenceId() bool`

HasNetworkReferenceId returns a boolean if a field has been set.

### GetOriginalCredit

`func (o *TransactionModel) GetOriginalCredit() OriginalCredit`

GetOriginalCredit returns the OriginalCredit field if non-nil, zero value otherwise.

### GetOriginalCreditOk

`func (o *TransactionModel) GetOriginalCreditOk() (*OriginalCredit, bool)`

GetOriginalCreditOk returns a tuple with the OriginalCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCredit

`func (o *TransactionModel) SetOriginalCredit(v OriginalCredit)`

SetOriginalCredit sets OriginalCredit field to given value.

### HasOriginalCredit

`func (o *TransactionModel) HasOriginalCredit() bool`

HasOriginalCredit returns a boolean if a field has been set.

### GetPeerTransfer

`func (o *TransactionModel) GetPeerTransfer() PeerTransferResponse`

GetPeerTransfer returns the PeerTransfer field if non-nil, zero value otherwise.

### GetPeerTransferOk

`func (o *TransactionModel) GetPeerTransferOk() (*PeerTransferResponse, bool)`

GetPeerTransferOk returns a tuple with the PeerTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerTransfer

`func (o *TransactionModel) SetPeerTransfer(v PeerTransferResponse)`

SetPeerTransfer sets PeerTransfer field to given value.

### HasPeerTransfer

`func (o *TransactionModel) HasPeerTransfer() bool`

HasPeerTransfer returns a boolean if a field has been set.

### GetPolarity

`func (o *TransactionModel) GetPolarity() string`

GetPolarity returns the Polarity field if non-nil, zero value otherwise.

### GetPolarityOk

`func (o *TransactionModel) GetPolarityOk() (*string, bool)`

GetPolarityOk returns a tuple with the Polarity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolarity

`func (o *TransactionModel) SetPolarity(v string)`

SetPolarity sets Polarity field to given value.

### HasPolarity

`func (o *TransactionModel) HasPolarity() bool`

HasPolarity returns a boolean if a field has been set.

### GetPos

`func (o *TransactionModel) GetPos() Pos`

GetPos returns the Pos field if non-nil, zero value otherwise.

### GetPosOk

`func (o *TransactionModel) GetPosOk() (*Pos, bool)`

GetPosOk returns a tuple with the Pos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPos

`func (o *TransactionModel) SetPos(v Pos)`

SetPos sets Pos field to given value.

### HasPos

`func (o *TransactionModel) HasPos() bool`

HasPos returns a boolean if a field has been set.

### GetPrecedingRelatedTransactionToken

`func (o *TransactionModel) GetPrecedingRelatedTransactionToken() string`

GetPrecedingRelatedTransactionToken returns the PrecedingRelatedTransactionToken field if non-nil, zero value otherwise.

### GetPrecedingRelatedTransactionTokenOk

`func (o *TransactionModel) GetPrecedingRelatedTransactionTokenOk() (*string, bool)`

GetPrecedingRelatedTransactionTokenOk returns a tuple with the PrecedingRelatedTransactionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedingRelatedTransactionToken

`func (o *TransactionModel) SetPrecedingRelatedTransactionToken(v string)`

SetPrecedingRelatedTransactionToken sets PrecedingRelatedTransactionToken field to given value.

### HasPrecedingRelatedTransactionToken

`func (o *TransactionModel) HasPrecedingRelatedTransactionToken() bool`

HasPrecedingRelatedTransactionToken returns a boolean if a field has been set.

### GetPrecedingTransaction

`func (o *TransactionModel) GetPrecedingTransaction() PrecedingTransaction`

GetPrecedingTransaction returns the PrecedingTransaction field if non-nil, zero value otherwise.

### GetPrecedingTransactionOk

`func (o *TransactionModel) GetPrecedingTransactionOk() (*PrecedingTransaction, bool)`

GetPrecedingTransactionOk returns a tuple with the PrecedingTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedingTransaction

`func (o *TransactionModel) SetPrecedingTransaction(v PrecedingTransaction)`

SetPrecedingTransaction sets PrecedingTransaction field to given value.

### HasPrecedingTransaction

`func (o *TransactionModel) HasPrecedingTransaction() bool`

HasPrecedingTransaction returns a boolean if a field has been set.

### GetProgram

`func (o *TransactionModel) GetProgram() Program`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *TransactionModel) GetProgramOk() (*Program, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *TransactionModel) SetProgram(v Program)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *TransactionModel) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetProgramTransfer

`func (o *TransactionModel) GetProgramTransfer() ProgramTransferResponse`

GetProgramTransfer returns the ProgramTransfer field if non-nil, zero value otherwise.

### GetProgramTransferOk

`func (o *TransactionModel) GetProgramTransferOk() (*ProgramTransferResponse, bool)`

GetProgramTransferOk returns a tuple with the ProgramTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramTransfer

`func (o *TransactionModel) SetProgramTransfer(v ProgramTransferResponse)`

SetProgramTransfer sets ProgramTransfer field to given value.

### HasProgramTransfer

`func (o *TransactionModel) HasProgramTransfer() bool`

HasProgramTransfer returns a boolean if a field has been set.

### GetRealTimeFeeGroup

`func (o *TransactionModel) GetRealTimeFeeGroup() RealTimeFeeGroup`

GetRealTimeFeeGroup returns the RealTimeFeeGroup field if non-nil, zero value otherwise.

### GetRealTimeFeeGroupOk

`func (o *TransactionModel) GetRealTimeFeeGroupOk() (*RealTimeFeeGroup, bool)`

GetRealTimeFeeGroupOk returns a tuple with the RealTimeFeeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTimeFeeGroup

`func (o *TransactionModel) SetRealTimeFeeGroup(v RealTimeFeeGroup)`

SetRealTimeFeeGroup sets RealTimeFeeGroup field to given value.

### HasRealTimeFeeGroup

`func (o *TransactionModel) HasRealTimeFeeGroup() bool`

HasRealTimeFeeGroup returns a boolean if a field has been set.

### GetRequestAmount

`func (o *TransactionModel) GetRequestAmount() float32`

GetRequestAmount returns the RequestAmount field if non-nil, zero value otherwise.

### GetRequestAmountOk

`func (o *TransactionModel) GetRequestAmountOk() (*float32, bool)`

GetRequestAmountOk returns a tuple with the RequestAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAmount

`func (o *TransactionModel) SetRequestAmount(v float32)`

SetRequestAmount sets RequestAmount field to given value.

### HasRequestAmount

`func (o *TransactionModel) HasRequestAmount() bool`

HasRequestAmount returns a boolean if a field has been set.

### GetResponse

`func (o *TransactionModel) GetResponse() Response`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *TransactionModel) GetResponseOk() (*Response, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *TransactionModel) SetResponse(v Response)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *TransactionModel) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSettlementDate

`func (o *TransactionModel) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *TransactionModel) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *TransactionModel) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *TransactionModel) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetStandinApprovedBy

`func (o *TransactionModel) GetStandinApprovedBy() string`

GetStandinApprovedBy returns the StandinApprovedBy field if non-nil, zero value otherwise.

### GetStandinApprovedByOk

`func (o *TransactionModel) GetStandinApprovedByOk() (*string, bool)`

GetStandinApprovedByOk returns a tuple with the StandinApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandinApprovedBy

`func (o *TransactionModel) SetStandinApprovedBy(v string)`

SetStandinApprovedBy sets StandinApprovedBy field to given value.

### HasStandinApprovedBy

`func (o *TransactionModel) HasStandinApprovedBy() bool`

HasStandinApprovedBy returns a boolean if a field has been set.

### GetStandinBy

`func (o *TransactionModel) GetStandinBy() string`

GetStandinBy returns the StandinBy field if non-nil, zero value otherwise.

### GetStandinByOk

`func (o *TransactionModel) GetStandinByOk() (*string, bool)`

GetStandinByOk returns a tuple with the StandinBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandinBy

`func (o *TransactionModel) SetStandinBy(v string)`

SetStandinBy sets StandinBy field to given value.

### HasStandinBy

`func (o *TransactionModel) HasStandinBy() bool`

HasStandinBy returns a boolean if a field has been set.

### GetStandinReason

`func (o *TransactionModel) GetStandinReason() string`

GetStandinReason returns the StandinReason field if non-nil, zero value otherwise.

### GetStandinReasonOk

`func (o *TransactionModel) GetStandinReasonOk() (*string, bool)`

GetStandinReasonOk returns a tuple with the StandinReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandinReason

`func (o *TransactionModel) SetStandinReason(v string)`

SetStandinReason sets StandinReason field to given value.

### HasStandinReason

`func (o *TransactionModel) HasStandinReason() bool`

HasStandinReason returns a boolean if a field has been set.

### GetState

`func (o *TransactionModel) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TransactionModel) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TransactionModel) SetState(v string)`

SetState sets State field to given value.


### GetStore

`func (o *TransactionModel) GetStore() StoreResponseModel`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *TransactionModel) GetStoreOk() (*StoreResponseModel, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *TransactionModel) SetStore(v StoreResponseModel)`

SetStore sets Store field to given value.

### HasStore

`func (o *TransactionModel) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetSubnetwork

`func (o *TransactionModel) GetSubnetwork() string`

GetSubnetwork returns the Subnetwork field if non-nil, zero value otherwise.

### GetSubnetworkOk

`func (o *TransactionModel) GetSubnetworkOk() (*string, bool)`

GetSubnetworkOk returns a tuple with the Subnetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetwork

`func (o *TransactionModel) SetSubnetwork(v string)`

SetSubnetwork sets Subnetwork field to given value.

### HasSubnetwork

`func (o *TransactionModel) HasSubnetwork() bool`

HasSubnetwork returns a boolean if a field has been set.

### GetToken

`func (o *TransactionModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TransactionModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TransactionModel) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransactionAttributes

`func (o *TransactionModel) GetTransactionAttributes() map[string]string`

GetTransactionAttributes returns the TransactionAttributes field if non-nil, zero value otherwise.

### GetTransactionAttributesOk

`func (o *TransactionModel) GetTransactionAttributesOk() (*map[string]string, bool)`

GetTransactionAttributesOk returns a tuple with the TransactionAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAttributes

`func (o *TransactionModel) SetTransactionAttributes(v map[string]string)`

SetTransactionAttributes sets TransactionAttributes field to given value.

### HasTransactionAttributes

`func (o *TransactionModel) HasTransactionAttributes() bool`

HasTransactionAttributes returns a boolean if a field has been set.

### GetTransactionMetadata

`func (o *TransactionModel) GetTransactionMetadata() TransactionMetadata`

GetTransactionMetadata returns the TransactionMetadata field if non-nil, zero value otherwise.

### GetTransactionMetadataOk

`func (o *TransactionModel) GetTransactionMetadataOk() (*TransactionMetadata, bool)`

GetTransactionMetadataOk returns a tuple with the TransactionMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionMetadata

`func (o *TransactionModel) SetTransactionMetadata(v TransactionMetadata)`

SetTransactionMetadata sets TransactionMetadata field to given value.

### HasTransactionMetadata

`func (o *TransactionModel) HasTransactionMetadata() bool`

HasTransactionMetadata returns a boolean if a field has been set.

### GetType

`func (o *TransactionModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionModel) SetType(v string)`

SetType sets Type field to given value.


### GetUser

`func (o *TransactionModel) GetUser() CardholderMetadata`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransactionModel) GetUserOk() (*CardholderMetadata, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransactionModel) SetUser(v CardholderMetadata)`

SetUser sets User field to given value.

### HasUser

`func (o *TransactionModel) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserToken

`func (o *TransactionModel) GetUserToken() string`

GetUserToken returns the UserToken field if non-nil, zero value otherwise.

### GetUserTokenOk

`func (o *TransactionModel) GetUserTokenOk() (*string, bool)`

GetUserTokenOk returns a tuple with the UserToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToken

`func (o *TransactionModel) SetUserToken(v string)`

SetUserToken sets UserToken field to given value.

### HasUserToken

`func (o *TransactionModel) HasUserToken() bool`

HasUserToken returns a boolean if a field has been set.

### GetUserTransactionTime

`func (o *TransactionModel) GetUserTransactionTime() time.Time`

GetUserTransactionTime returns the UserTransactionTime field if non-nil, zero value otherwise.

### GetUserTransactionTimeOk

`func (o *TransactionModel) GetUserTransactionTimeOk() (*time.Time, bool)`

GetUserTransactionTimeOk returns a tuple with the UserTransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTransactionTime

`func (o *TransactionModel) SetUserTransactionTime(v time.Time)`

SetUserTransactionTime sets UserTransactionTime field to given value.

### HasUserTransactionTime

`func (o *TransactionModel) HasUserTransactionTime() bool`

HasUserTransactionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


