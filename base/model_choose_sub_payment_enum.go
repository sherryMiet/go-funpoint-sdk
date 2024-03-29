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

// ChooseSubPaymentEnum **付款子項目**   若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。   請參考付款方式一覽表
type ChooseSubPaymentEnum string

// List of ChooseSubPaymentEnum
const (
	CHOOSESUBPAYMENTENUM_TAISHIN    ChooseSubPaymentEnum = "TAISHIN"
	CHOOSESUBPAYMENTENUM_ESUN       ChooseSubPaymentEnum = "ESUN"
	CHOOSESUBPAYMENTENUM_BOT        ChooseSubPaymentEnum = "BOT"
	CHOOSESUBPAYMENTENUM_FUBON      ChooseSubPaymentEnum = "FUBON"
	CHOOSESUBPAYMENTENUM_CHINATRUST ChooseSubPaymentEnum = "CHINATRUST"
	CHOOSESUBPAYMENTENUM_FIRST      ChooseSubPaymentEnum = "FIRST"
	CHOOSESUBPAYMENTENUM_CATHAY     ChooseSubPaymentEnum = "CATHAY"
	CHOOSESUBPAYMENTENUM_MEGA       ChooseSubPaymentEnum = "MEGA"
	CHOOSESUBPAYMENTENUM_LAND       ChooseSubPaymentEnum = "LAND"
	CHOOSESUBPAYMENTENUM_TACHONG    ChooseSubPaymentEnum = "TACHONG"
	CHOOSESUBPAYMENTENUM_SINOPAC    ChooseSubPaymentEnum = "SINOPAC"
	CHOOSESUBPAYMENTENUM_CVS        ChooseSubPaymentEnum = "CVS"
	CHOOSESUBPAYMENTENUM_OK         ChooseSubPaymentEnum = "OK"
	CHOOSESUBPAYMENTENUM_FAMILY     ChooseSubPaymentEnum = "FAMILY"
	CHOOSESUBPAYMENTENUM_HILIFE     ChooseSubPaymentEnum = "HILIFE"
	CHOOSESUBPAYMENTENUM_IBON       ChooseSubPaymentEnum = "IBON"
	CHOOSESUBPAYMENTENUM_BARCODE    ChooseSubPaymentEnum = "BARCODE"
	CHOOSESUBPAYMENTENUM_EMPTY      ChooseSubPaymentEnum = ""
)

// Ptr returns reference to ChooseSubPaymentEnum value
func (v ChooseSubPaymentEnum) Ptr() *ChooseSubPaymentEnum {
	return &v
}

type NullableChooseSubPaymentEnum struct {
	value *ChooseSubPaymentEnum
	isSet bool
}

func (v NullableChooseSubPaymentEnum) Get() *ChooseSubPaymentEnum {
	return v.value
}

func (v *NullableChooseSubPaymentEnum) Set(val *ChooseSubPaymentEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableChooseSubPaymentEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableChooseSubPaymentEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChooseSubPaymentEnum(val *ChooseSubPaymentEnum) *NullableChooseSubPaymentEnum {
	return &NullableChooseSubPaymentEnum{value: val, isSet: true}
}

func (v NullableChooseSubPaymentEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChooseSubPaymentEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
