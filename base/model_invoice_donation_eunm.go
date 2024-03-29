/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package funpointBase

import (
	"encoding/json"
)

// InvoiceDonationEunm **捐贈註記** 若為捐贈時，參數請帶：`1`。 若為不捐贈或統一編號 `CustomerIdentifier` 有值時，參數請帶：`0`。
type InvoiceDonationEunm string

// List of InvoiceDonationEunm
const (
	INVOICEDONATIONEUNM_DONATION     InvoiceDonationEunm = "1"
	INVOICEDONATIONEUNM_NOT_DONATION InvoiceDonationEunm = "0"
)

// Ptr returns reference to InvoiceDonationEunm value
func (v InvoiceDonationEunm) Ptr() *InvoiceDonationEunm {
	return &v
}

type NullableInvoiceDonationEunm struct {
	value *InvoiceDonationEunm
	isSet bool
}

func (v NullableInvoiceDonationEunm) Get() *InvoiceDonationEunm {
	return v.value
}

func (v *NullableInvoiceDonationEunm) Set(val *InvoiceDonationEunm) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDonationEunm) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDonationEunm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDonationEunm(val *InvoiceDonationEunm) *NullableInvoiceDonationEunm {
	return &NullableInvoiceDonationEunm{value: val, isSet: true}
}

func (v NullableInvoiceDonationEunm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDonationEunm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
