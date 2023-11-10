/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.11
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marqeta_coreapi_client

import (
	"encoding/json"
	"fmt"
)

// AccountStatusEnum Status of the credit account.  *NOTE* `CHARGE_OFF` is not an allowable value for `original_status`.
type AccountStatusEnum string

// List of AccountStatusEnum
const (
	ACCOUNTSTATUSENUM_UNACTIVATED AccountStatusEnum = "UNACTIVATED"
	ACCOUNTSTATUSENUM_ACTIVE AccountStatusEnum = "ACTIVE"
	ACCOUNTSTATUSENUM_SUSPENDED AccountStatusEnum = "SUSPENDED"
	ACCOUNTSTATUSENUM_TERMINATED AccountStatusEnum = "TERMINATED"
	ACCOUNTSTATUSENUM_CHARGE_OFF AccountStatusEnum = "CHARGE_OFF"
)

// All allowed values of AccountStatusEnum enum
var AllowedAccountStatusEnumEnumValues = []AccountStatusEnum{
	"UNACTIVATED",
	"ACTIVE",
	"SUSPENDED",
	"TERMINATED",
	"CHARGE_OFF",
}

func (v *AccountStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountStatusEnum(value)
	for _, existing := range AllowedAccountStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountStatusEnum", value)
}

// NewAccountStatusEnumFromValue returns a pointer to a valid AccountStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountStatusEnumFromValue(v string) (*AccountStatusEnum, error) {
	ev := AccountStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountStatusEnum: valid values are %v", v, AllowedAccountStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountStatusEnum) IsValid() bool {
	for _, existing := range AllowedAccountStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountStatusEnum value
func (v AccountStatusEnum) Ptr() *AccountStatusEnum {
	return &v
}

type NullableAccountStatusEnum struct {
	value *AccountStatusEnum
	isSet bool
}

func (v NullableAccountStatusEnum) Get() *AccountStatusEnum {
	return v.value
}

func (v *NullableAccountStatusEnum) Set(val *AccountStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatusEnum(val *AccountStatusEnum) *NullableAccountStatusEnum {
	return &NullableAccountStatusEnum{value: val, isSet: true}
}

func (v NullableAccountStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
