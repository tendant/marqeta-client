# TransactionControls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedCountriesToken** | Pointer to **string** | Set to &#x60;accept_us_only&#x60; to allow transactions only within the US.  Set to &#x60;decline_ofac_countries&#x60; to allow international transactions except with countries that the Financial Action Task Force (FATF) and Office of Foreign Assets Control (OFAC) have identified as high risk.  Users with the Admin role can create and update additional lists of accepted countries for transactions at the &#x60;/acceptedcountries&#x60; endpoint. See &lt;&lt;/core-api/accepted-countries, Accepted Countries&gt;&gt;. | [optional] 
**AddressVerification** | Pointer to [**AvsControls**](AvsControls.md) |  | [optional] 
**AllowChipFallback** | Pointer to **bool** | Indicates whether to allow transactions where a Europay Mastercard and Visa (EMV) chip-enabled card was processed using the magstripe as fallback. | [optional] [default to false]
**AllowFirstPinSetViaFinancialTransaction** | Pointer to **bool** | *WARNING:* This field is deprecated and will be unsupported in a future release.  Allows cardholders to define a personal identification number (PIN) as they complete their first PIN-debit transaction. | [optional] [default to false]
**AllowGpaAuth** | Pointer to **bool** | If set to &#x60;true&#x60;, transactions can be authorized using GPA funds.  *NOTE:* For most programs, this field should be set to &#x60;true&#x60;. | [optional] [default to false]
**AllowMccGroupAuthorizationControls** | Pointer to **bool** | The &lt;&lt;/core-api/mcc-groups, MCC group&gt;&gt; &#x60;authorization_controls&#x60; object allows you to automatically increase authorization holds and to specify authorization expiration times based on merchant type. By default, these settings apply to all cards in your program. You can, however, exempt cards associated with a particular card product by setting this field to &#x60;false&#x60;.  *NOTE:* Partial authorizations are disallowed if this field is set to &#x60;true&#x60;. | [optional] [default to false]
**AllowNetworkLoad** | Pointer to **bool** | Indicates whether card network loads are allowed. The associated card&#39;s state must be &#x60;ACTIVE&#x60; or the load will be rejected. | [optional] [default to false]
**AllowNetworkLoadCardActivation** | Pointer to **bool** | Indicates whether card network loads are allowed. Sets the associated card&#39;s state to &#x60;ACTIVE&#x60; if its current state is &#x60;INACTIVE&#x60;. | [optional] [default to false]
**AllowQuasiCash** | Pointer to **bool** | Indicates whether quasi-cash transactions are allowed. In a quasi-cash transaction, the cardholder purchases an item that can be directly converted to cash, such as traveler&#39;s checks, money orders, casino chips, or lottery tickets. | [optional] [default to false]
**AlwaysRequireIcc** | Pointer to **bool** | If set to &#x60;true&#x60;, cards of this card product type require an Integrated Circuit Card. | [optional] [default to false]
**AlwaysRequirePin** | Pointer to **bool** | If set to &#x60;true&#x60;, cards of this card product type require a personal identification number (PIN). | [optional] [default to false]
**EnableCreditService** | Pointer to **bool** |  | [optional] [default to false]
**EnablePartialAuthApproval** | Pointer to **bool** | Set to &#x60;true&#x60; to enable partial authorizations.  When this setting is &#x60;false&#x60; and the requested authorization amount exceeds available funds, the transaction is declined. When this setting is &#x60;true&#x60; and the requested authorization amount exceeds available funds, the transaction is authorized for the amount of available funds. | [optional] [default to false]
**IgnoreCardSuspendedState** | Pointer to **bool** | Allows transactions to be approved even if the card&#39;s &#x60;state &#x3D; SUSPENDED&#x60;. When this field is set to &#x60;true&#x60;, the card behaves as if its &#x60;state &#x3D; ACTIVE&#x60;. | [optional] [default to false]
**NotificationLanguage** | Pointer to **string** | Specifies the language for 3D Secure and digital wallet token notifications sent to cardholders under this card program.  You can send notifications to your cardholders in the following languages:  * *ces* – Czech * *deu* – German * *eng* – English * *fra* – French * *grc* – Greek * *ita* – Italian * *nld* – Dutch * *pol* – Polish * *prt* – Portuguese * *rou* – Romanian * *spa* – Spanish * *swe* – Swedish  By default, notifications are sent in English.  To specify the language for OTP notifications at the user level, see &lt;&lt;/core-api/users, Users&gt;&gt;. Languages set at the user level take precedence over the language set at the card product level. | [optional] 
**QuasiCashExemptMerchantGroupToken** | Pointer to **string** | The token of the merchant group that you want to exempt from quasi-cash transaction authorization control, allowing your cardholders to conduct quasi-cash transactions. In a quasi-cash transaction, the cardholder purchases an item that can be directly converted to cash, such as traveler&#39;s checks, money orders, casino chips, or lottery tickets.  You can specify a merchant group token in addition to whatever merchant identifiers you listed in the &#x60;quasi_cash_exempt_mids&#x60; field, if any. For more information, see &lt;&lt;/core-api/merchant-groups, Merchant Groups&gt;&gt;. | [optional] 
**QuasiCashExemptMids** | Pointer to **string** | Comma-separated list of merchant identifiers that you want to exempt from quasi-cash transaction authorization control, allowing your cardholders to conduct quasi-cash transactions. In a quasi-cash transaction, the cardholder purchases an item that can be directly converted to cash, such as traveler&#39;s checks, money orders, casino chips, or lottery tickets. | [optional] 
**RequireCardNotPresentCardSecurityCode** | Pointer to **bool** | A value of &#x60;true&#x60; indicates that if &#x60;card_presence_required&#x60; is &#x60;true&#x60;, the card&#39;s security code is required. | [optional] [default to false]
**StrongCustomerAuthenticationLimits** | Pointer to [**StrongCustomerAuthenticationLimits**](StrongCustomerAuthenticationLimits.md) |  | [optional] 

## Methods

### NewTransactionControls

`func NewTransactionControls() *TransactionControls`

NewTransactionControls instantiates a new TransactionControls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionControlsWithDefaults

`func NewTransactionControlsWithDefaults() *TransactionControls`

NewTransactionControlsWithDefaults instantiates a new TransactionControls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedCountriesToken

`func (o *TransactionControls) GetAcceptedCountriesToken() string`

GetAcceptedCountriesToken returns the AcceptedCountriesToken field if non-nil, zero value otherwise.

### GetAcceptedCountriesTokenOk

`func (o *TransactionControls) GetAcceptedCountriesTokenOk() (*string, bool)`

GetAcceptedCountriesTokenOk returns a tuple with the AcceptedCountriesToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCountriesToken

`func (o *TransactionControls) SetAcceptedCountriesToken(v string)`

SetAcceptedCountriesToken sets AcceptedCountriesToken field to given value.

### HasAcceptedCountriesToken

`func (o *TransactionControls) HasAcceptedCountriesToken() bool`

HasAcceptedCountriesToken returns a boolean if a field has been set.

### GetAddressVerification

`func (o *TransactionControls) GetAddressVerification() AvsControls`

GetAddressVerification returns the AddressVerification field if non-nil, zero value otherwise.

### GetAddressVerificationOk

`func (o *TransactionControls) GetAddressVerificationOk() (*AvsControls, bool)`

GetAddressVerificationOk returns a tuple with the AddressVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerification

`func (o *TransactionControls) SetAddressVerification(v AvsControls)`

SetAddressVerification sets AddressVerification field to given value.

### HasAddressVerification

`func (o *TransactionControls) HasAddressVerification() bool`

HasAddressVerification returns a boolean if a field has been set.

### GetAllowChipFallback

`func (o *TransactionControls) GetAllowChipFallback() bool`

GetAllowChipFallback returns the AllowChipFallback field if non-nil, zero value otherwise.

### GetAllowChipFallbackOk

`func (o *TransactionControls) GetAllowChipFallbackOk() (*bool, bool)`

GetAllowChipFallbackOk returns a tuple with the AllowChipFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChipFallback

`func (o *TransactionControls) SetAllowChipFallback(v bool)`

SetAllowChipFallback sets AllowChipFallback field to given value.

### HasAllowChipFallback

`func (o *TransactionControls) HasAllowChipFallback() bool`

HasAllowChipFallback returns a boolean if a field has been set.

### GetAllowFirstPinSetViaFinancialTransaction

`func (o *TransactionControls) GetAllowFirstPinSetViaFinancialTransaction() bool`

GetAllowFirstPinSetViaFinancialTransaction returns the AllowFirstPinSetViaFinancialTransaction field if non-nil, zero value otherwise.

### GetAllowFirstPinSetViaFinancialTransactionOk

`func (o *TransactionControls) GetAllowFirstPinSetViaFinancialTransactionOk() (*bool, bool)`

GetAllowFirstPinSetViaFinancialTransactionOk returns a tuple with the AllowFirstPinSetViaFinancialTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFirstPinSetViaFinancialTransaction

`func (o *TransactionControls) SetAllowFirstPinSetViaFinancialTransaction(v bool)`

SetAllowFirstPinSetViaFinancialTransaction sets AllowFirstPinSetViaFinancialTransaction field to given value.

### HasAllowFirstPinSetViaFinancialTransaction

`func (o *TransactionControls) HasAllowFirstPinSetViaFinancialTransaction() bool`

HasAllowFirstPinSetViaFinancialTransaction returns a boolean if a field has been set.

### GetAllowGpaAuth

`func (o *TransactionControls) GetAllowGpaAuth() bool`

GetAllowGpaAuth returns the AllowGpaAuth field if non-nil, zero value otherwise.

### GetAllowGpaAuthOk

`func (o *TransactionControls) GetAllowGpaAuthOk() (*bool, bool)`

GetAllowGpaAuthOk returns a tuple with the AllowGpaAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGpaAuth

`func (o *TransactionControls) SetAllowGpaAuth(v bool)`

SetAllowGpaAuth sets AllowGpaAuth field to given value.

### HasAllowGpaAuth

`func (o *TransactionControls) HasAllowGpaAuth() bool`

HasAllowGpaAuth returns a boolean if a field has been set.

### GetAllowMccGroupAuthorizationControls

`func (o *TransactionControls) GetAllowMccGroupAuthorizationControls() bool`

GetAllowMccGroupAuthorizationControls returns the AllowMccGroupAuthorizationControls field if non-nil, zero value otherwise.

### GetAllowMccGroupAuthorizationControlsOk

`func (o *TransactionControls) GetAllowMccGroupAuthorizationControlsOk() (*bool, bool)`

GetAllowMccGroupAuthorizationControlsOk returns a tuple with the AllowMccGroupAuthorizationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMccGroupAuthorizationControls

`func (o *TransactionControls) SetAllowMccGroupAuthorizationControls(v bool)`

SetAllowMccGroupAuthorizationControls sets AllowMccGroupAuthorizationControls field to given value.

### HasAllowMccGroupAuthorizationControls

`func (o *TransactionControls) HasAllowMccGroupAuthorizationControls() bool`

HasAllowMccGroupAuthorizationControls returns a boolean if a field has been set.

### GetAllowNetworkLoad

`func (o *TransactionControls) GetAllowNetworkLoad() bool`

GetAllowNetworkLoad returns the AllowNetworkLoad field if non-nil, zero value otherwise.

### GetAllowNetworkLoadOk

`func (o *TransactionControls) GetAllowNetworkLoadOk() (*bool, bool)`

GetAllowNetworkLoadOk returns a tuple with the AllowNetworkLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNetworkLoad

`func (o *TransactionControls) SetAllowNetworkLoad(v bool)`

SetAllowNetworkLoad sets AllowNetworkLoad field to given value.

### HasAllowNetworkLoad

`func (o *TransactionControls) HasAllowNetworkLoad() bool`

HasAllowNetworkLoad returns a boolean if a field has been set.

### GetAllowNetworkLoadCardActivation

`func (o *TransactionControls) GetAllowNetworkLoadCardActivation() bool`

GetAllowNetworkLoadCardActivation returns the AllowNetworkLoadCardActivation field if non-nil, zero value otherwise.

### GetAllowNetworkLoadCardActivationOk

`func (o *TransactionControls) GetAllowNetworkLoadCardActivationOk() (*bool, bool)`

GetAllowNetworkLoadCardActivationOk returns a tuple with the AllowNetworkLoadCardActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNetworkLoadCardActivation

`func (o *TransactionControls) SetAllowNetworkLoadCardActivation(v bool)`

SetAllowNetworkLoadCardActivation sets AllowNetworkLoadCardActivation field to given value.

### HasAllowNetworkLoadCardActivation

`func (o *TransactionControls) HasAllowNetworkLoadCardActivation() bool`

HasAllowNetworkLoadCardActivation returns a boolean if a field has been set.

### GetAllowQuasiCash

`func (o *TransactionControls) GetAllowQuasiCash() bool`

GetAllowQuasiCash returns the AllowQuasiCash field if non-nil, zero value otherwise.

### GetAllowQuasiCashOk

`func (o *TransactionControls) GetAllowQuasiCashOk() (*bool, bool)`

GetAllowQuasiCashOk returns a tuple with the AllowQuasiCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuasiCash

`func (o *TransactionControls) SetAllowQuasiCash(v bool)`

SetAllowQuasiCash sets AllowQuasiCash field to given value.

### HasAllowQuasiCash

`func (o *TransactionControls) HasAllowQuasiCash() bool`

HasAllowQuasiCash returns a boolean if a field has been set.

### GetAlwaysRequireIcc

`func (o *TransactionControls) GetAlwaysRequireIcc() bool`

GetAlwaysRequireIcc returns the AlwaysRequireIcc field if non-nil, zero value otherwise.

### GetAlwaysRequireIccOk

`func (o *TransactionControls) GetAlwaysRequireIccOk() (*bool, bool)`

GetAlwaysRequireIccOk returns a tuple with the AlwaysRequireIcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysRequireIcc

`func (o *TransactionControls) SetAlwaysRequireIcc(v bool)`

SetAlwaysRequireIcc sets AlwaysRequireIcc field to given value.

### HasAlwaysRequireIcc

`func (o *TransactionControls) HasAlwaysRequireIcc() bool`

HasAlwaysRequireIcc returns a boolean if a field has been set.

### GetAlwaysRequirePin

`func (o *TransactionControls) GetAlwaysRequirePin() bool`

GetAlwaysRequirePin returns the AlwaysRequirePin field if non-nil, zero value otherwise.

### GetAlwaysRequirePinOk

`func (o *TransactionControls) GetAlwaysRequirePinOk() (*bool, bool)`

GetAlwaysRequirePinOk returns a tuple with the AlwaysRequirePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysRequirePin

`func (o *TransactionControls) SetAlwaysRequirePin(v bool)`

SetAlwaysRequirePin sets AlwaysRequirePin field to given value.

### HasAlwaysRequirePin

`func (o *TransactionControls) HasAlwaysRequirePin() bool`

HasAlwaysRequirePin returns a boolean if a field has been set.

### GetEnableCreditService

`func (o *TransactionControls) GetEnableCreditService() bool`

GetEnableCreditService returns the EnableCreditService field if non-nil, zero value otherwise.

### GetEnableCreditServiceOk

`func (o *TransactionControls) GetEnableCreditServiceOk() (*bool, bool)`

GetEnableCreditServiceOk returns a tuple with the EnableCreditService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCreditService

`func (o *TransactionControls) SetEnableCreditService(v bool)`

SetEnableCreditService sets EnableCreditService field to given value.

### HasEnableCreditService

`func (o *TransactionControls) HasEnableCreditService() bool`

HasEnableCreditService returns a boolean if a field has been set.

### GetEnablePartialAuthApproval

`func (o *TransactionControls) GetEnablePartialAuthApproval() bool`

GetEnablePartialAuthApproval returns the EnablePartialAuthApproval field if non-nil, zero value otherwise.

### GetEnablePartialAuthApprovalOk

`func (o *TransactionControls) GetEnablePartialAuthApprovalOk() (*bool, bool)`

GetEnablePartialAuthApprovalOk returns a tuple with the EnablePartialAuthApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartialAuthApproval

`func (o *TransactionControls) SetEnablePartialAuthApproval(v bool)`

SetEnablePartialAuthApproval sets EnablePartialAuthApproval field to given value.

### HasEnablePartialAuthApproval

`func (o *TransactionControls) HasEnablePartialAuthApproval() bool`

HasEnablePartialAuthApproval returns a boolean if a field has been set.

### GetIgnoreCardSuspendedState

`func (o *TransactionControls) GetIgnoreCardSuspendedState() bool`

GetIgnoreCardSuspendedState returns the IgnoreCardSuspendedState field if non-nil, zero value otherwise.

### GetIgnoreCardSuspendedStateOk

`func (o *TransactionControls) GetIgnoreCardSuspendedStateOk() (*bool, bool)`

GetIgnoreCardSuspendedStateOk returns a tuple with the IgnoreCardSuspendedState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCardSuspendedState

`func (o *TransactionControls) SetIgnoreCardSuspendedState(v bool)`

SetIgnoreCardSuspendedState sets IgnoreCardSuspendedState field to given value.

### HasIgnoreCardSuspendedState

`func (o *TransactionControls) HasIgnoreCardSuspendedState() bool`

HasIgnoreCardSuspendedState returns a boolean if a field has been set.

### GetNotificationLanguage

`func (o *TransactionControls) GetNotificationLanguage() string`

GetNotificationLanguage returns the NotificationLanguage field if non-nil, zero value otherwise.

### GetNotificationLanguageOk

`func (o *TransactionControls) GetNotificationLanguageOk() (*string, bool)`

GetNotificationLanguageOk returns a tuple with the NotificationLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationLanguage

`func (o *TransactionControls) SetNotificationLanguage(v string)`

SetNotificationLanguage sets NotificationLanguage field to given value.

### HasNotificationLanguage

`func (o *TransactionControls) HasNotificationLanguage() bool`

HasNotificationLanguage returns a boolean if a field has been set.

### GetQuasiCashExemptMerchantGroupToken

`func (o *TransactionControls) GetQuasiCashExemptMerchantGroupToken() string`

GetQuasiCashExemptMerchantGroupToken returns the QuasiCashExemptMerchantGroupToken field if non-nil, zero value otherwise.

### GetQuasiCashExemptMerchantGroupTokenOk

`func (o *TransactionControls) GetQuasiCashExemptMerchantGroupTokenOk() (*string, bool)`

GetQuasiCashExemptMerchantGroupTokenOk returns a tuple with the QuasiCashExemptMerchantGroupToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuasiCashExemptMerchantGroupToken

`func (o *TransactionControls) SetQuasiCashExemptMerchantGroupToken(v string)`

SetQuasiCashExemptMerchantGroupToken sets QuasiCashExemptMerchantGroupToken field to given value.

### HasQuasiCashExemptMerchantGroupToken

`func (o *TransactionControls) HasQuasiCashExemptMerchantGroupToken() bool`

HasQuasiCashExemptMerchantGroupToken returns a boolean if a field has been set.

### GetQuasiCashExemptMids

`func (o *TransactionControls) GetQuasiCashExemptMids() string`

GetQuasiCashExemptMids returns the QuasiCashExemptMids field if non-nil, zero value otherwise.

### GetQuasiCashExemptMidsOk

`func (o *TransactionControls) GetQuasiCashExemptMidsOk() (*string, bool)`

GetQuasiCashExemptMidsOk returns a tuple with the QuasiCashExemptMids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuasiCashExemptMids

`func (o *TransactionControls) SetQuasiCashExemptMids(v string)`

SetQuasiCashExemptMids sets QuasiCashExemptMids field to given value.

### HasQuasiCashExemptMids

`func (o *TransactionControls) HasQuasiCashExemptMids() bool`

HasQuasiCashExemptMids returns a boolean if a field has been set.

### GetRequireCardNotPresentCardSecurityCode

`func (o *TransactionControls) GetRequireCardNotPresentCardSecurityCode() bool`

GetRequireCardNotPresentCardSecurityCode returns the RequireCardNotPresentCardSecurityCode field if non-nil, zero value otherwise.

### GetRequireCardNotPresentCardSecurityCodeOk

`func (o *TransactionControls) GetRequireCardNotPresentCardSecurityCodeOk() (*bool, bool)`

GetRequireCardNotPresentCardSecurityCodeOk returns a tuple with the RequireCardNotPresentCardSecurityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCardNotPresentCardSecurityCode

`func (o *TransactionControls) SetRequireCardNotPresentCardSecurityCode(v bool)`

SetRequireCardNotPresentCardSecurityCode sets RequireCardNotPresentCardSecurityCode field to given value.

### HasRequireCardNotPresentCardSecurityCode

`func (o *TransactionControls) HasRequireCardNotPresentCardSecurityCode() bool`

HasRequireCardNotPresentCardSecurityCode returns a boolean if a field has been set.

### GetStrongCustomerAuthenticationLimits

`func (o *TransactionControls) GetStrongCustomerAuthenticationLimits() StrongCustomerAuthenticationLimits`

GetStrongCustomerAuthenticationLimits returns the StrongCustomerAuthenticationLimits field if non-nil, zero value otherwise.

### GetStrongCustomerAuthenticationLimitsOk

`func (o *TransactionControls) GetStrongCustomerAuthenticationLimitsOk() (*StrongCustomerAuthenticationLimits, bool)`

GetStrongCustomerAuthenticationLimitsOk returns a tuple with the StrongCustomerAuthenticationLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongCustomerAuthenticationLimits

`func (o *TransactionControls) SetStrongCustomerAuthenticationLimits(v StrongCustomerAuthenticationLimits)`

SetStrongCustomerAuthenticationLimits sets StrongCustomerAuthenticationLimits field to given value.

### HasStrongCustomerAuthenticationLimits

`func (o *TransactionControls) HasStrongCustomerAuthenticationLimits() bool`

HasStrongCustomerAuthenticationLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


