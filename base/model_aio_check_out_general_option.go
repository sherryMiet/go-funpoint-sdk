/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package opayBase

import (
	"encoding/json"
)

// AioCheckOutGeneralOption struct for AioCheckOutGeneralOption
type AioCheckOutGeneralOption struct {
	// **特店編號(由綠界提供)**
	MerchantID string `json:"MerchantID" form:"MerchantID"`
	// **特店交易編號(由特店提供)**   特店交易編號均為唯一值，不可重複使用。   英數字大小寫混合   如何避免訂單編號重複請參考 FAQ   如有使用 `PlatformID` ，平台商底下所有商家之訂單編號亦不可重複。
	MerchantTradeNo string `json:"MerchantTradeNo" form:"MerchantTradeNo"`
	// **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。
	StoreID *string `json:"StoreID,omitempty" form:"StoreID"`
	// **特店交易時間** 格式為 `yyyy/MM/dd HH:mm:ss`
	MerchantTradeDate ECPayDateTime           `json:"MerchantTradeDate" form:"MerchantTradeDate"`
	PaymentType       AioCheckPaymentTypeEnum `json:"PaymentType" form:"PaymentType"`
	// **交易金額**   請帶整數，不可有小數點。   僅限新台幣。   各付款金額的限制，請參考 <https://www.ecpay.com.tw/CascadeFAQ/CascadeFAQ_Qa?nID=3605>
	TotalAmount int `json:"TotalAmount" form:"TotalAmount"`
	// **交易描述** 傳送到綠界前，請將參數值先做 UrlEncode。
	TradeDesc string `json:"TradeDesc" form:"TradeDesc"`
	// **商品名稱**   1. 如果商品名稱有多筆，需在金流選擇頁一行一行顯示商品名稱的話，商品名稱請以符號#分隔。   2. 商品名稱字數限制為中英數 400 字內，超過此限制系統將自動截斷。
	ItemName string `json:"ItemName" form:"ItemName"`
	// **付款完成通知回傳網址**   當消費者付款完成後，綠界會將付款結果參數以幕後(Server POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：    1. 請勿設定與 Client 端接收付款結果網址 OrderResultURL 相同位置，避免程式判斷錯誤。   2. 請在收到 Server 端付款結果通知後，請正確回應 1|OK 給綠界。
	ReturnURL     string            `json:"ReturnURL" form:"ReturnURL"`
	ChoosePayment ChoosePaymentEnum `json:"ChoosePayment" form:"ChoosePayment"`
	// **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式
	CheckMacValue string `json:"CheckMacValue" form:"CheckMacValue"`
	// **Client端返回特店的按鈕連結**   消費者點選此按鈕後，會將頁面導回到此設定的網址   注意事項：   導回時不會帶付款結果到此網址，只是將頁面導回而已。   設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。   設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。   若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。
	ClientBackURL *string `json:"ClientBackURL,omitempty" form:"ClientBackURL"`
	// **商品銷售網址**
	ItemURL *string `json:"ItemURL,omitempty" form:"ItemURL"`
	// **備註欄位**
	Remark           *string               `json:"Remark,omitempty" form:"Remark"`
	ChooseSubPayment *ChooseSubPaymentEnum `json:"ChooseSubPayment,omitempty" form:"ChooseSubPayment"`
	// **Client端回傳付款結果網址**     當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：   1. 若與[ClientBackURL]同時設定，將會以此參數為主。   2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。
	OrderResultURL    *string                `json:"OrderResultURL,omitempty" form:"OrderResultURL"`
	NeedExtraPaidInfo *NeedExtraPaidInfoEnum `json:"NeedExtraPaidInfo,omitempty" form:"NeedExtraPaidInfo"`
	// **裝置來源** 請帶空值，由系統自動判定。
	DeviceSource *string `json:"DeviceSource,omitempty" form:"DeviceSource"`
	// **隱藏付款**   當付款方式 `ChoosePayment` 為 `ALL` 時，可隱藏不需要的付款方式，多筆請以井號分隔(#)。   可用的參數值：   - `Credit`: 信用卡   - `WebATM`: 網路 ATM   - `ATM`: 自動櫃員機   - `CVS`: 超商代碼   - `BARCODE`: 超商條碼
	IgnorePayment *string `json:"IgnorePayment,omitempty" form:"IgnorePayment"`
	// **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 `MerchantID`。
	PlatformID  *string          `json:"PlatformID,omitempty" form:"PlatformID"`
	InvoiceMark *InvoiceMarkEnum `json:"InvoiceMark,omitempty" form:"InvoiceMark"`
	// **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField1 *string `json:"CustomField1,omitempty" form:"CustomField1"`
	// **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField2 *string `json:"CustomField2,omitempty" form:"CustomField2"`
	// **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField3 *string `json:"CustomField3,omitempty" form:"CustomField3"`
	// **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField4 *string         `json:"CustomField4,omitempty" form:"CustomField4"`
	EncryptType  EncryptTypeEnum `json:"EncryptType" form:"EncryptType"`
	Language     *LanguageEnum   `json:"Language,omitempty" form:"Language"`
}

// NewAioCheckOutGeneralOption instantiates a new AioCheckOutGeneralOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutGeneralOption(merchantID string, merchantTradeNo string, merchantTradeDate ECPayDateTime, paymentType AioCheckPaymentTypeEnum, totalAmount int, tradeDesc string, itemName string, returnURL string, choosePayment ChoosePaymentEnum, checkMacValue string, encryptType EncryptTypeEnum) *AioCheckOutGeneralOption {
	this := AioCheckOutGeneralOption{}
	this.MerchantID = merchantID
	this.MerchantTradeNo = merchantTradeNo
	this.MerchantTradeDate = merchantTradeDate
	this.PaymentType = paymentType
	this.TotalAmount = totalAmount
	this.TradeDesc = tradeDesc
	this.ItemName = itemName
	this.ReturnURL = returnURL
	this.ChoosePayment = choosePayment
	this.CheckMacValue = checkMacValue
	this.EncryptType = encryptType
	return &this
}

// NewAioCheckOutGeneralOptionWithDefaults instantiates a new AioCheckOutGeneralOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutGeneralOptionWithDefaults() *AioCheckOutGeneralOption {
	this := AioCheckOutGeneralOption{}
	var paymentType AioCheckPaymentTypeEnum = "aio"
	this.PaymentType = paymentType
	var encryptType EncryptTypeEnum = ENCRYPTTYPEENUM_SHA256
	this.EncryptType = encryptType
	return &this
}

// GetMerchantID returns the MerchantID field value
func (o *AioCheckOutGeneralOption) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *AioCheckOutGeneralOption) SetMerchantID(v string) {
	o.MerchantID = v
}

// GetMerchantTradeNo returns the MerchantTradeNo field value
func (o *AioCheckOutGeneralOption) GetMerchantTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantTradeNo
}

// GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetMerchantTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTradeNo, true
}

// SetMerchantTradeNo sets field value
func (o *AioCheckOutGeneralOption) SetMerchantTradeNo(v string) {
	o.MerchantTradeNo = v
}

// GetStoreID returns the StoreID field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetStoreID() string {
	if o == nil || o.StoreID == nil {
		var ret string
		return ret
	}
	return *o.StoreID
}

// GetStoreIDOk returns a tuple with the StoreID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetStoreIDOk() (*string, bool) {
	if o == nil || o.StoreID == nil {
		return nil, false
	}
	return o.StoreID, true
}

// HasStoreID returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasStoreID() bool {
	if o != nil && o.StoreID != nil {
		return true
	}

	return false
}

// SetStoreID gets a reference to the given string and assigns it to the StoreID field.
func (o *AioCheckOutGeneralOption) SetStoreID(v string) {
	o.StoreID = &v
}

// GetMerchantTradeDate returns the MerchantTradeDate field value
func (o *AioCheckOutGeneralOption) GetMerchantTradeDate() ECPayDateTime {
	if o == nil {
		var ret ECPayDateTime
		return ret
	}

	return o.MerchantTradeDate
}

// GetMerchantTradeDateOk returns a tuple with the MerchantTradeDate field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetMerchantTradeDateOk() (*ECPayDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTradeDate, true
}

// SetMerchantTradeDate sets field value
func (o *AioCheckOutGeneralOption) SetMerchantTradeDate(v ECPayDateTime) {
	o.MerchantTradeDate = v
}

// GetPaymentType returns the PaymentType field value
func (o *AioCheckOutGeneralOption) GetPaymentType() AioCheckPaymentTypeEnum {
	if o == nil {
		var ret AioCheckPaymentTypeEnum
		return ret
	}

	return o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetPaymentTypeOk() (*AioCheckPaymentTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentType, true
}

// SetPaymentType sets field value
func (o *AioCheckOutGeneralOption) SetPaymentType(v AioCheckPaymentTypeEnum) {
	o.PaymentType = v
}

// GetTotalAmount returns the TotalAmount field value
func (o *AioCheckOutGeneralOption) GetTotalAmount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetTotalAmountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *AioCheckOutGeneralOption) SetTotalAmount(v int) {
	o.TotalAmount = v
}

// GetTradeDesc returns the TradeDesc field value
func (o *AioCheckOutGeneralOption) GetTradeDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeDesc
}

// GetTradeDescOk returns a tuple with the TradeDesc field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetTradeDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeDesc, true
}

// SetTradeDesc sets field value
func (o *AioCheckOutGeneralOption) SetTradeDesc(v string) {
	o.TradeDesc = v
}

// GetItemName returns the ItemName field value
func (o *AioCheckOutGeneralOption) GetItemName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemName
}

// GetItemNameOk returns a tuple with the ItemName field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetItemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemName, true
}

// SetItemName sets field value
func (o *AioCheckOutGeneralOption) SetItemName(v string) {
	o.ItemName = v
}

// GetReturnURL returns the ReturnURL field value
func (o *AioCheckOutGeneralOption) GetReturnURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnURL
}

// GetReturnURLOk returns a tuple with the ReturnURL field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetReturnURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnURL, true
}

// SetReturnURL sets field value
func (o *AioCheckOutGeneralOption) SetReturnURL(v string) {
	o.ReturnURL = v
}

// GetChoosePayment returns the ChoosePayment field value
func (o *AioCheckOutGeneralOption) GetChoosePayment() ChoosePaymentEnum {
	if o == nil {
		var ret ChoosePaymentEnum
		return ret
	}

	return o.ChoosePayment
}

// GetChoosePaymentOk returns a tuple with the ChoosePayment field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetChoosePaymentOk() (*ChoosePaymentEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChoosePayment, true
}

// SetChoosePayment sets field value
func (o *AioCheckOutGeneralOption) SetChoosePayment(v ChoosePaymentEnum) {
	o.ChoosePayment = v
}

// GetCheckMacValue returns the CheckMacValue field value
func (o *AioCheckOutGeneralOption) GetCheckMacValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckMacValue
}

// GetCheckMacValueOk returns a tuple with the CheckMacValue field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetCheckMacValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckMacValue, true
}

// SetCheckMacValue sets field value
func (o *AioCheckOutGeneralOption) SetCheckMacValue(v string) {
	o.CheckMacValue = v
}

// GetClientBackURL returns the ClientBackURL field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetClientBackURL() string {
	if o == nil || o.ClientBackURL == nil {
		var ret string
		return ret
	}
	return *o.ClientBackURL
}

// GetClientBackURLOk returns a tuple with the ClientBackURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetClientBackURLOk() (*string, bool) {
	if o == nil || o.ClientBackURL == nil {
		return nil, false
	}
	return o.ClientBackURL, true
}

// HasClientBackURL returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasClientBackURL() bool {
	if o != nil && o.ClientBackURL != nil {
		return true
	}

	return false
}

// SetClientBackURL gets a reference to the given string and assigns it to the ClientBackURL field.
func (o *AioCheckOutGeneralOption) SetClientBackURL(v string) {
	o.ClientBackURL = &v
}

// GetItemURL returns the ItemURL field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetItemURL() string {
	if o == nil || o.ItemURL == nil {
		var ret string
		return ret
	}
	return *o.ItemURL
}

// GetItemURLOk returns a tuple with the ItemURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetItemURLOk() (*string, bool) {
	if o == nil || o.ItemURL == nil {
		return nil, false
	}
	return o.ItemURL, true
}

// HasItemURL returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasItemURL() bool {
	if o != nil && o.ItemURL != nil {
		return true
	}

	return false
}

// SetItemURL gets a reference to the given string and assigns it to the ItemURL field.
func (o *AioCheckOutGeneralOption) SetItemURL(v string) {
	o.ItemURL = &v
}

// GetRemark returns the Remark field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetRemark() string {
	if o == nil || o.Remark == nil {
		var ret string
		return ret
	}
	return *o.Remark
}

// GetRemarkOk returns a tuple with the Remark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetRemarkOk() (*string, bool) {
	if o == nil || o.Remark == nil {
		return nil, false
	}
	return o.Remark, true
}

// HasRemark returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasRemark() bool {
	if o != nil && o.Remark != nil {
		return true
	}

	return false
}

// SetRemark gets a reference to the given string and assigns it to the Remark field.
func (o *AioCheckOutGeneralOption) SetRemark(v string) {
	o.Remark = &v
}

// GetChooseSubPayment returns the ChooseSubPayment field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetChooseSubPayment() ChooseSubPaymentEnum {
	if o == nil || o.ChooseSubPayment == nil {
		var ret ChooseSubPaymentEnum
		return ret
	}
	return *o.ChooseSubPayment
}

// GetChooseSubPaymentOk returns a tuple with the ChooseSubPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetChooseSubPaymentOk() (*ChooseSubPaymentEnum, bool) {
	if o == nil || o.ChooseSubPayment == nil {
		return nil, false
	}
	return o.ChooseSubPayment, true
}

// HasChooseSubPayment returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasChooseSubPayment() bool {
	if o != nil && o.ChooseSubPayment != nil {
		return true
	}

	return false
}

// SetChooseSubPayment gets a reference to the given ChooseSubPaymentEnum and assigns it to the ChooseSubPayment field.
func (o *AioCheckOutGeneralOption) SetChooseSubPayment(v ChooseSubPaymentEnum) {
	o.ChooseSubPayment = &v
}

// GetOrderResultURL returns the OrderResultURL field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetOrderResultURL() string {
	if o == nil || o.OrderResultURL == nil {
		var ret string
		return ret
	}
	return *o.OrderResultURL
}

// GetOrderResultURLOk returns a tuple with the OrderResultURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetOrderResultURLOk() (*string, bool) {
	if o == nil || o.OrderResultURL == nil {
		return nil, false
	}
	return o.OrderResultURL, true
}

// HasOrderResultURL returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasOrderResultURL() bool {
	if o != nil && o.OrderResultURL != nil {
		return true
	}

	return false
}

// SetOrderResultURL gets a reference to the given string and assigns it to the OrderResultURL field.
func (o *AioCheckOutGeneralOption) SetOrderResultURL(v string) {
	o.OrderResultURL = &v
}

// GetNeedExtraPaidInfo returns the NeedExtraPaidInfo field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetNeedExtraPaidInfo() NeedExtraPaidInfoEnum {
	if o == nil || o.NeedExtraPaidInfo == nil {
		var ret NeedExtraPaidInfoEnum
		return ret
	}
	return *o.NeedExtraPaidInfo
}

// GetNeedExtraPaidInfoOk returns a tuple with the NeedExtraPaidInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetNeedExtraPaidInfoOk() (*NeedExtraPaidInfoEnum, bool) {
	if o == nil || o.NeedExtraPaidInfo == nil {
		return nil, false
	}
	return o.NeedExtraPaidInfo, true
}

// HasNeedExtraPaidInfo returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasNeedExtraPaidInfo() bool {
	if o != nil && o.NeedExtraPaidInfo != nil {
		return true
	}

	return false
}

// SetNeedExtraPaidInfo gets a reference to the given NeedExtraPaidInfoEnum and assigns it to the NeedExtraPaidInfo field.
func (o *AioCheckOutGeneralOption) SetNeedExtraPaidInfo(v NeedExtraPaidInfoEnum) {
	o.NeedExtraPaidInfo = &v
}

// GetDeviceSource returns the DeviceSource field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetDeviceSource() string {
	if o == nil || o.DeviceSource == nil {
		var ret string
		return ret
	}
	return *o.DeviceSource
}

// GetDeviceSourceOk returns a tuple with the DeviceSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetDeviceSourceOk() (*string, bool) {
	if o == nil || o.DeviceSource == nil {
		return nil, false
	}
	return o.DeviceSource, true
}

// HasDeviceSource returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasDeviceSource() bool {
	if o != nil && o.DeviceSource != nil {
		return true
	}

	return false
}

// SetDeviceSource gets a reference to the given string and assigns it to the DeviceSource field.
func (o *AioCheckOutGeneralOption) SetDeviceSource(v string) {
	o.DeviceSource = &v
}

// GetIgnorePayment returns the IgnorePayment field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetIgnorePayment() string {
	if o == nil || o.IgnorePayment == nil {
		var ret string
		return ret
	}
	return *o.IgnorePayment
}

// GetIgnorePaymentOk returns a tuple with the IgnorePayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetIgnorePaymentOk() (*string, bool) {
	if o == nil || o.IgnorePayment == nil {
		return nil, false
	}
	return o.IgnorePayment, true
}

// HasIgnorePayment returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasIgnorePayment() bool {
	if o != nil && o.IgnorePayment != nil {
		return true
	}

	return false
}

// SetIgnorePayment gets a reference to the given string and assigns it to the IgnorePayment field.
func (o *AioCheckOutGeneralOption) SetIgnorePayment(v string) {
	o.IgnorePayment = &v
}

// GetPlatformID returns the PlatformID field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetPlatformID() string {
	if o == nil || o.PlatformID == nil {
		var ret string
		return ret
	}
	return *o.PlatformID
}

// GetPlatformIDOk returns a tuple with the PlatformID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetPlatformIDOk() (*string, bool) {
	if o == nil || o.PlatformID == nil {
		return nil, false
	}
	return o.PlatformID, true
}

// HasPlatformID returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasPlatformID() bool {
	if o != nil && o.PlatformID != nil {
		return true
	}

	return false
}

// SetPlatformID gets a reference to the given string and assigns it to the PlatformID field.
func (o *AioCheckOutGeneralOption) SetPlatformID(v string) {
	o.PlatformID = &v
}

// GetInvoiceMark returns the InvoiceMark field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetInvoiceMark() InvoiceMarkEnum {
	if o == nil || o.InvoiceMark == nil {
		var ret InvoiceMarkEnum
		return ret
	}
	return *o.InvoiceMark
}

// GetInvoiceMarkOk returns a tuple with the InvoiceMark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetInvoiceMarkOk() (*InvoiceMarkEnum, bool) {
	if o == nil || o.InvoiceMark == nil {
		return nil, false
	}
	return o.InvoiceMark, true
}

// HasInvoiceMark returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasInvoiceMark() bool {
	if o != nil && o.InvoiceMark != nil {
		return true
	}

	return false
}

// SetInvoiceMark gets a reference to the given InvoiceMarkEnum and assigns it to the InvoiceMark field.
func (o *AioCheckOutGeneralOption) SetInvoiceMark(v InvoiceMarkEnum) {
	o.InvoiceMark = &v
}

// GetCustomField1 returns the CustomField1 field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetCustomField1() string {
	if o == nil || o.CustomField1 == nil {
		var ret string
		return ret
	}
	return *o.CustomField1
}

// GetCustomField1Ok returns a tuple with the CustomField1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetCustomField1Ok() (*string, bool) {
	if o == nil || o.CustomField1 == nil {
		return nil, false
	}
	return o.CustomField1, true
}

// HasCustomField1 returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasCustomField1() bool {
	if o != nil && o.CustomField1 != nil {
		return true
	}

	return false
}

// SetCustomField1 gets a reference to the given string and assigns it to the CustomField1 field.
func (o *AioCheckOutGeneralOption) SetCustomField1(v string) {
	o.CustomField1 = &v
}

// GetCustomField2 returns the CustomField2 field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetCustomField2() string {
	if o == nil || o.CustomField2 == nil {
		var ret string
		return ret
	}
	return *o.CustomField2
}

// GetCustomField2Ok returns a tuple with the CustomField2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetCustomField2Ok() (*string, bool) {
	if o == nil || o.CustomField2 == nil {
		return nil, false
	}
	return o.CustomField2, true
}

// HasCustomField2 returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasCustomField2() bool {
	if o != nil && o.CustomField2 != nil {
		return true
	}

	return false
}

// SetCustomField2 gets a reference to the given string and assigns it to the CustomField2 field.
func (o *AioCheckOutGeneralOption) SetCustomField2(v string) {
	o.CustomField2 = &v
}

// GetCustomField3 returns the CustomField3 field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetCustomField3() string {
	if o == nil || o.CustomField3 == nil {
		var ret string
		return ret
	}
	return *o.CustomField3
}

// GetCustomField3Ok returns a tuple with the CustomField3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetCustomField3Ok() (*string, bool) {
	if o == nil || o.CustomField3 == nil {
		return nil, false
	}
	return o.CustomField3, true
}

// HasCustomField3 returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasCustomField3() bool {
	if o != nil && o.CustomField3 != nil {
		return true
	}

	return false
}

// SetCustomField3 gets a reference to the given string and assigns it to the CustomField3 field.
func (o *AioCheckOutGeneralOption) SetCustomField3(v string) {
	o.CustomField3 = &v
}

// GetCustomField4 returns the CustomField4 field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetCustomField4() string {
	if o == nil || o.CustomField4 == nil {
		var ret string
		return ret
	}
	return *o.CustomField4
}

// GetCustomField4Ok returns a tuple with the CustomField4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetCustomField4Ok() (*string, bool) {
	if o == nil || o.CustomField4 == nil {
		return nil, false
	}
	return o.CustomField4, true
}

// HasCustomField4 returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasCustomField4() bool {
	if o != nil && o.CustomField4 != nil {
		return true
	}

	return false
}

// SetCustomField4 gets a reference to the given string and assigns it to the CustomField4 field.
func (o *AioCheckOutGeneralOption) SetCustomField4(v string) {
	o.CustomField4 = &v
}

// GetEncryptType returns the EncryptType field value
func (o *AioCheckOutGeneralOption) GetEncryptType() EncryptTypeEnum {
	if o == nil {
		var ret EncryptTypeEnum
		return ret
	}

	return o.EncryptType
}

// GetEncryptTypeOk returns a tuple with the EncryptType field value
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetEncryptTypeOk() (*EncryptTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptType, true
}

// SetEncryptType sets field value
func (o *AioCheckOutGeneralOption) SetEncryptType(v EncryptTypeEnum) {
	o.EncryptType = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AioCheckOutGeneralOption) GetLanguage() LanguageEnum {
	if o == nil || o.Language == nil {
		var ret LanguageEnum
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutGeneralOption) GetLanguageOk() (*LanguageEnum, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AioCheckOutGeneralOption) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given LanguageEnum and assigns it to the Language field.
func (o *AioCheckOutGeneralOption) SetLanguage(v LanguageEnum) {
	o.Language = &v
}

func (o AioCheckOutGeneralOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MerchantID"] = o.MerchantID
	}
	if true {
		toSerialize["MerchantTradeNo"] = o.MerchantTradeNo
	}
	if o.StoreID != nil {
		toSerialize["StoreID"] = o.StoreID
	}
	if true {
		toSerialize["MerchantTradeDate"] = o.MerchantTradeDate
	}
	if true {
		toSerialize["PaymentType"] = o.PaymentType
	}
	if true {
		toSerialize["TotalAmount"] = o.TotalAmount
	}
	if true {
		toSerialize["TradeDesc"] = o.TradeDesc
	}
	if true {
		toSerialize["ItemName"] = o.ItemName
	}
	if true {
		toSerialize["ReturnURL"] = o.ReturnURL
	}
	if true {
		toSerialize["ChoosePayment"] = o.ChoosePayment
	}
	if true {
		toSerialize["CheckMacValue"] = o.CheckMacValue
	}
	if o.ClientBackURL != nil {
		toSerialize["ClientBackURL"] = o.ClientBackURL
	}
	if o.ItemURL != nil {
		toSerialize["ItemURL"] = o.ItemURL
	}
	if o.Remark != nil {
		toSerialize["Remark"] = o.Remark
	}
	if o.ChooseSubPayment != nil {
		toSerialize["ChooseSubPayment"] = o.ChooseSubPayment
	}
	if o.OrderResultURL != nil {
		toSerialize["OrderResultURL"] = o.OrderResultURL
	}
	if o.NeedExtraPaidInfo != nil {
		toSerialize["NeedExtraPaidInfo"] = o.NeedExtraPaidInfo
	}
	if o.DeviceSource != nil {
		toSerialize["DeviceSource"] = o.DeviceSource
	}
	if o.IgnorePayment != nil {
		toSerialize["IgnorePayment"] = o.IgnorePayment
	}
	if o.PlatformID != nil {
		toSerialize["PlatformID"] = o.PlatformID
	}
	if o.InvoiceMark != nil {
		toSerialize["InvoiceMark"] = o.InvoiceMark
	}
	if o.CustomField1 != nil {
		toSerialize["CustomField1"] = o.CustomField1
	}
	if o.CustomField2 != nil {
		toSerialize["CustomField2"] = o.CustomField2
	}
	if o.CustomField3 != nil {
		toSerialize["CustomField3"] = o.CustomField3
	}
	if o.CustomField4 != nil {
		toSerialize["CustomField4"] = o.CustomField4
	}
	if true {
		toSerialize["EncryptType"] = o.EncryptType
	}
	if o.Language != nil {
		toSerialize["Language"] = o.Language
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutGeneralOption struct {
	value *AioCheckOutGeneralOption
	isSet bool
}

func (v NullableAioCheckOutGeneralOption) Get() *AioCheckOutGeneralOption {
	return v.value
}

func (v *NullableAioCheckOutGeneralOption) Set(val *AioCheckOutGeneralOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutGeneralOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutGeneralOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutGeneralOption(val *AioCheckOutGeneralOption) *NullableAioCheckOutGeneralOption {
	return &NullableAioCheckOutGeneralOption{value: val, isSet: true}
}

func (v NullableAioCheckOutGeneralOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutGeneralOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
