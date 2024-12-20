package funpoint

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	GetTokenByBindingCardURL       = "https://ecpg.funpoint.com.tw/Merchant/GetTokenbyBindingCard"
	GetTokenByBindingCardTestURL   = "https://ecpg-stage.funpoint.com.tw/Merchant/GetTokenbyBindingCard"
	CreateBindCardURL              = "https://ecpg.funpoint.com.tw/Merchant/CreateBindCard"
	CreateBindCardTestURL          = "https://ecpg-stage.funpoint.com.tw/Merchant/CreateBindCard"
	CreatePaymentWithCardIDURL     = "https://ecpg.funpoint.com.tw/Merchant/CreatePaymentWithCardID"
	CreatePaymentWithCardIDTestURL = "https://ecpg-stage.funpoint.com.tw/Merchant/CreatePaymentWithCardID"
	DeleteMemberBindCardURL        = "https://ecpg.funpoint.com.tw/Merchant/DeleteMemberBindCard"
	DeleteMemberBindCardTestURL    = "https://ecpg-stage.funpoint.com.tw/Merchant/DeleteMemberBindCard"
	QueryTradeTestURL              = "https://ecpayment-stage.funpoint.com.tw/1.0.0/CreditDetail/QueryTrade"
	QueryTradeURL                  = "https://ecpayment.funpoint.com.tw/1.0.0/CreditDetail/QueryTrade"
	RefundURL                      = "https://ecpayment.funpoint.com.tw/1.0.0/Credit/DoAction"
)

type RefundCall struct {
	Client            *Client
	RefundRequest     *RefundRequest
	RefundRequestData *RefundRequestData
}

type CreateBindCardCall struct {
	Client                    *Client
	CreateBindCardRequest     *CreateBindCardRequest
	CreateBindCardRequestData *CreateBindCardRequestData
}

type GetTokenByBindingCardRequestCall struct {
	Client                           *Client
	GetTokenByBindingCardRequest     *GetTokenByBindingCardRequest
	GetTokenByBindingCardRequestData *GetTokenByBindingCardRequestData
}

type CreatePaymentWithCardIDCall struct {
	Client                             *Client
	CreatePaymentWithCardIDRequest     *CreatePaymentWithCardIDRequest
	CreatePaymentWithCardIDRequestData *CreatePaymentWithCardIDRequestData
}

type DeleteMemberBindCardCall struct {
	Client                          *Client
	DeleteMemberBindCardRequest     *DeleteMemberBindCardRequest
	DeleteMemberBindCardRequestData *DeleteMemberBindCardRequestData
}

type QueryTradeCall struct {
	Client                *Client
	QueryTradeRequest     *QueryTradeRequest
	QueryTradeRequestData *QueryTradeRequestData
}

type RefundRequest struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	Data       string   `json:"Data"`
}

type RefundRequestData struct {
	MerchantID      string `json:"MerchantID"`
	MerchantTradeNo string `json:"MerchantTradeNo"`
	TradeNo         string `json:"TradeNo"`
	Action          string `json:"Action"`
	TotalAmount     int    `json:"TotalAmount"`
}

type RefundResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type RefundResponseData struct {
	MerchantID      string `json:"MerchantID"`
	MerchantTradeNo string `json:"MerchantTradeNo"`
	TradeNo         string `json:"TradeNo"`
	RtnCode         int    `json:"RtnCode"`
	RtnMsg          string `json:"RtnMsg"`
}

type QueryTradeRequest struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	Data       string   `json:"Data"`
}

type QueryTradeRequestData struct {
	MerchantTradeNo string `json:"MerchantTradeNo"`
	MerchantID      string `json:"MerchantID"`
}

type QueryTradeResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type QueryTradeResponseData struct {
	RtnMsg   string `json:"RtnMsg"`
	RtnValue struct {
		TradeID  int    `json:"TradeID"`
		Amount   int    `json:"Amount"`
		ClsAmt   int    `json:"ClsAmt"`
		AuthTime string `json:"AuthTime"`
		Status   string `json:"Status"`
	} `json:"RtnValue"`
	CloseData struct {
		Status   string `json:"Status"`
		Amount   int    `json:"Amount"`
		DateTime string `json:"DateTime"`
	} `json:"CloseData"`
}

type DeleteMemberBindCardRequest struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	Data       string   `json:"Data"`
}

type DeleteMemberBindCardRequestData struct {
	MerchantID string `json:"MerchantID"`
	BindCardID string `json:"BindCardID"`
}

type DeleteMemberBindCardResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type DeleteMemberBindCardResponseData struct {
	RtnCode int    `json:"RtnCode"`
	RtnMsg  string `json:"RtnMsg"`
}

type CreateBindCardRequest struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	Data       string   `json:"Data"`
}

type CreateBindCardRequestData struct {
	MerchantID       string `json:"MerchantID"`
	BindCardPayToken string `json:"BindCardPayToken"`
	MerchantMemberID string `json:"MerchantMemberID"`
}

type CreateBindCardResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type CreateBindCardResponseData struct {
	RtnCode    int    `json:"RtnCode"`
	RtnMsg     string `json:"RtnMsg"`
	PlatformID string `json:"PlatformID"`
	MerchantID string `json:"MerchantID"`
	BindCardID string `json:"BindCardID"`
	IsSameCard bool   `json:"IsSameCard"`
	OrderInfo  struct {
		MerchantTradeNo string `json:"MerchantTradeNo"`
		TradeNo         string `json:"TradeNo"`
		PaymentDate     string `json:"PaymentDate"`
		TradeAmt        int    `json:"TradeAmt"`
		PaymentType     string `json:"PaymentType"`
		TradeDate       string `json:"TradeDate"`
		ChargeFee       int    `json:"ChargeFee"`
		TradeStatus     string `json:"TradeStatus"`
	} `json:"OrderInfo"`
	CardInfo struct {
		AuthCode        string `json:"AuthCode"`
		Gwsr            int    `json:"Gwsr"`
		ProcessDate     string `json:"ProcessDate"`
		Amount          int    `json:"Amount"`
		IssuingBank     string `json:"IssuingBank"`
		IssuingBankCode string `json:"IssuingBankCode "`
		Card6No         string `json:"Card6No"`
		Card4No         string `json:"Card4No"`
		Eci             int    `json:"Eci"`
	} `json:"CardInfo"`
	ThreeDInfo struct {
		ThreeDURL string `json:"ThreeDURL"`
	} `json:"ThreeDInfo"`
}

type CreatePaymentWithCardIDRequest struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	Data       string   `json:"Data"`
}

type CreatePaymentWithCardIDRequestData struct {
	MerchantID   string       `json:"MerchantID"`
	BindCardID   string       `json:"BindCardID"`
	ConsumerInfo ConsumerInfo `json:"ConsumerInfo"`
	OrderInfo    OrderInfo    `json:"OrderInfo"`
}

type CreatePaymentWithCardIDResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type CreatePaymentWithCardIDResponseData struct {
	RtnCode    int    `json:"RtnCode"`
	RtnMsg     string `json:"RtnMsg"`
	PlatformID string `json:"PlatformID"`
	MerchantID string `json:"MerchantID"`
	OrderInfo  struct {
		MerchantTradeNo string `json:"MerchantTradeNo"`
		TradeNo         string `json:"TradeNo"`
		PaymentDate     string `json:"PaymentDate"`
		TradeAmt        int    `json:"TradeAmt"`
		PaymentType     string `json:"PaymentType"`
		TradeDate       string `json:"TradeDate"`
		ChargeFee       int    `json:"ChargeFee"`
		TradeStatus     string `json:"TradeStatus"`
	} `json:"OrderInfo"`
	CardInfo struct {
		AuthCode        string `json:"AuthCode"`
		Gwsr            int    `json:"Gwsr"`
		ProcessDate     string `json:"ProcessDate"`
		Amount          int    `json:"Amount"`
		Card6No         string `json:"Card6No"`
		Card4No         string `json:"Card4No"`
		IssuingBank     string `json:"IssuingBank"`
		IssuingBankCode string `json:"IssuingBankCode"`
	} `json:"CardInfo"`
	ThreeDInfo struct {
		ThreeDURL string `json:"ThreeDURL"`
	} `json:"ThreeDInfo"`
}

type GetTokenByBindingCardRequest struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	Data       string   `json:"Data"`
}

type RqHeader struct {
	Timestamp int    `json:"Timestamp"`
	Revision  string `json:"Revision"`
}

type GetTokenByBindingCardRequestData struct {
	MerchantID     string       `json:"MerchantID"`
	ConsumerInfo   ConsumerInfo `json:"ConsumerInfo"`
	OrderInfo      OrderInfo    `json:"OrderInfo"`
	OrderResultURL string       `json:"OrderResultURL"`
}

type GetTokenByBindingCardResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type GetTokenByBindingCardResponseData struct {
	RtnCode         int    `json:"RtnCode"`
	RtnMsg          string `json:"RtnMsg"`
	MerchantID      string `json:"MerchantID"`
	Token           string `json:"Token"`
	TokenExpireDate string `json:"TokenExpireDate"`
}

type ConsumerInfo struct {
	MerchantMemberID string `json:"MerchantMemberID"`
	Email            string `json:"Email"`
	Phone            string `json:"Phone"`
	Name             string `json:"Name"`
	CountryCode      string `json:"CountryCode"`
}

type OrderInfo struct {
	MerchantTradeNo   string `json:"MerchantTradeNo"`
	MerchantTradeDate string `json:"MerchantTradeDate"`
	TotalAmount       int    `json:"TotalAmount"`
	ReturnURL         string `json:"ReturnURL"`
	TradeDesc         string `json:"TradeDesc"`
	ItemName          string `json:"ItemName"`
}

type BindCardResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type BindCardResponseData struct {
	RtnCode    int    `json:"RtnCode"`
	RtnMsg     string `json:"RtnMsg"`
	MerchantID string `json:"MerchantID"`
	OrderInfo  struct {
		MerchantTradeNo string `json:"MerchantTradeNo"`
		TradeNo         string `json:"TradeNo"`
		TradeDate       string `json:"TradeDate"`
	} `json:"OrderInfo"`
	ThreeDInfo struct {
		ThreeDURL string `json:"ThreeDURL"`
	} `json:"ThreeDInfo"`
}

func Aes128(plaintext string, key string, iv string) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS7Padding([]byte(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecodeAes128(cipherText string, key string, iv string) string {
	bIV := []byte(iv)
	bKey := []byte(key)
	decodedData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		fmt.Printf("Error decoding Base64 encoded data %v", err)
	}
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)

	mode.CryptBlocks([]byte(decodedData), []byte(decodedData))
	result := PKCS7UnPadding(decodedData, block.BlockSize())
	return string(result)
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func SendRequest(json []byte, URL string) ([]byte, error) {
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(json))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}

func (c *Client) NewGetTokenByBindingCard() *GetTokenByBindingCardRequestCall {
	r := new(GetTokenByBindingCardRequestCall)
	r.Client = c
	return r
}

func (c *Client) NewCreatePaymentWithCardID() *CreatePaymentWithCardIDCall {
	r := new(CreatePaymentWithCardIDCall)
	r.Client = c
	return r
}

func (c *Client) NewCreateBindCard() *CreateBindCardCall {
	r := new(CreateBindCardCall)
	r.Client = c
	return r
}

func (c *Client) NewDeleteMemberBindCard() *DeleteMemberBindCardCall {
	r := new(DeleteMemberBindCardCall)
	r.Client = c
	return r
}

func (c *Client) NewQueryTrade() *QueryTradeCall {
	r := new(QueryTradeCall)
	r.Client = c
	return r
}

func (c *Client) NewRefund() *RefundCall {
	r := new(RefundCall)
	r.Client = c
	return r
}

func (c *QueryTradeCall) Request() *QueryTradeCall {
	c.QueryTradeRequest = &QueryTradeRequest{}
	c.QueryTradeRequestData = &QueryTradeRequestData{}
	return c
}

func (c *CreatePaymentWithCardIDCall) Request() *CreatePaymentWithCardIDCall {
	c.CreatePaymentWithCardIDRequestData = &CreatePaymentWithCardIDRequestData{}
	c.CreatePaymentWithCardIDRequest = &CreatePaymentWithCardIDRequest{}
	return c
}

func (c *GetTokenByBindingCardRequestCall) Request() *GetTokenByBindingCardRequestCall {
	c.GetTokenByBindingCardRequestData = &GetTokenByBindingCardRequestData{}
	c.GetTokenByBindingCardRequest = &GetTokenByBindingCardRequest{}
	return c
}

func (c *CreateBindCardCall) Request() *CreateBindCardCall {
	c.CreateBindCardRequestData = &CreateBindCardRequestData{}
	c.CreateBindCardRequest = &CreateBindCardRequest{}
	return c
}

func (c *DeleteMemberBindCardCall) Request() *DeleteMemberBindCardCall {
	c.DeleteMemberBindCardRequest = &DeleteMemberBindCardRequest{}
	c.DeleteMemberBindCardRequestData = &DeleteMemberBindCardRequestData{}
	return c
}

func (c *RefundCall) Request() *RefundCall {
	c.RefundRequest = &RefundRequest{}
	c.RefundRequestData = &RefundRequestData{}
	return c
}

func (c *GetTokenByBindingCardRequestCall) GetTokenByBindingCard() *GetTokenByBindingCardRequestCall {
	c.GetTokenByBindingCardRequest.MerchantID = c.Client.merchantID
	c.GetTokenByBindingCardRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.GetTokenByBindingCardRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.GetTokenByBindingCardRequestData)
	fmt.Println(string(jsonData))
	c.GetTokenByBindingCardRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *CreatePaymentWithCardIDCall) CreatePaymentWithCardID() *CreatePaymentWithCardIDCall {
	c.CreatePaymentWithCardIDRequest.MerchantID = c.Client.merchantID
	c.CreatePaymentWithCardIDRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.CreatePaymentWithCardIDRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.CreatePaymentWithCardIDRequestData)
	c.CreatePaymentWithCardIDRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *CreateBindCardCall) CreateBindCard() *CreateBindCardCall {
	c.CreateBindCardRequest.MerchantID = c.Client.merchantID
	c.CreateBindCardRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.CreateBindCardRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.CreateBindCardRequestData)
	fmt.Println(string(jsonData))
	c.CreateBindCardRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *DeleteMemberBindCardCall) DeleteMemberBindCard() *DeleteMemberBindCardCall {
	c.DeleteMemberBindCardRequest.MerchantID = c.Client.merchantID
	c.DeleteMemberBindCardRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.DeleteMemberBindCardRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.DeleteMemberBindCardRequestData)
	fmt.Println(string(jsonData))
	c.DeleteMemberBindCardRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *RefundCall) Refund() *RefundCall {
	c.RefundRequest.MerchantID = c.Client.merchantID
	c.RefundRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.RefundRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.RefundRequestData)
	fmt.Println(string(jsonData))
	c.RefundRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *QueryTradeCall) QueryTrade() *QueryTradeCall {
	c.QueryTradeRequest.MerchantID = c.Client.merchantID
	c.QueryTradeRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.QueryTradeRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.QueryTradeRequestData)
	fmt.Println(string(jsonData))
	c.QueryTradeRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *QueryTradeCall) Do(URL string) (*QueryTradeResponseData, error) {
	marshal, err := json.Marshal(c.QueryTradeRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(QueryTradeResponse)
	err = json.Unmarshal(request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.TransCode != 1 {
		return nil, errors.New(response.TransMsg)
	} else {

		dataString := DecodeAes128(response.Data, c.Client.hashKey, c.Client.hashIV)
		unescape, err := url.QueryUnescape(dataString)
		if err != nil {
			return nil, err
		}
		responseData := new(QueryTradeResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}

func (c *GetTokenByBindingCardRequestCall) Do(URL string) (*GetTokenByBindingCardResponseData, error) {
	marshal, err := json.Marshal(c.GetTokenByBindingCardRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(GetTokenByBindingCardResponse)
	err = json.Unmarshal(request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.TransCode != 1 {
		return nil, errors.New(response.TransMsg)
	} else {

		dataString := DecodeAes128(response.Data, c.Client.hashKey, c.Client.hashIV)
		unescape, err := url.QueryUnescape(dataString)
		if err != nil {
			return nil, err
		}
		responseData := new(GetTokenByBindingCardResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}

func (c *CreatePaymentWithCardIDCall) Do(URL string) (*CreatePaymentWithCardIDResponseData, error) {
	marshal, err := json.Marshal(c.CreatePaymentWithCardIDRequest)
	if err != nil {
		return nil, err
	}
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(CreatePaymentWithCardIDResponse)
	err = json.Unmarshal(request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.TransCode != 1 {
		return nil, errors.New(response.TransMsg)
	} else {

		dataString := DecodeAes128(response.Data, c.Client.hashKey, c.Client.hashIV)
		unescape, err := url.QueryUnescape(dataString)
		if err != nil {
			return nil, err
		}
		responseData := new(CreatePaymentWithCardIDResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}

func (c *CreateBindCardCall) Do(URL string) (*CreateBindCardResponseData, error) {
	marshal, err := json.Marshal(c.CreateBindCardRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(CreateBindCardResponse)
	err = json.Unmarshal(request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.TransCode != 1 {
		return nil, errors.New(response.TransMsg)
	} else {

		dataString := DecodeAes128(response.Data, c.Client.hashKey, c.Client.hashIV)
		unescape, err := url.QueryUnescape(dataString)
		if err != nil {
			return nil, err
		}
		responseData := new(CreateBindCardResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}

func (c *DeleteMemberBindCardCall) Do(URL string) (*DeleteMemberBindCardResponseData, error) {
	marshal, err := json.Marshal(c.DeleteMemberBindCardRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(CreateBindCardResponse)
	err = json.Unmarshal(request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.TransCode != 1 {
		return nil, errors.New(response.TransMsg)
	} else {

		dataString := DecodeAes128(response.Data, c.Client.hashKey, c.Client.hashIV)
		unescape, err := url.QueryUnescape(dataString)
		if err != nil {
			return nil, err
		}
		responseData := new(DeleteMemberBindCardResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}

func (c *RefundCall) Do(URL string) (*RefundResponseData, error) {
	marshal, err := json.Marshal(c.RefundRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(RefundResponse)
	err = json.Unmarshal(request, response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.TransCode != 1 {
		return nil, errors.New(response.TransMsg)
	} else {

		dataString := DecodeAes128(response.Data, c.Client.hashKey, c.Client.hashIV)
		unescape, err := url.QueryUnescape(dataString)
		if err != nil {
			return nil, err
		}
		responseData := new(RefundResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}
