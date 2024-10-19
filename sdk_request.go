package funpoint

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"net/url"
	"time"
)

const (
	TokenByTradeTestURL  = "https://ecpg-stage.funpoint.com.tw/Merchant/GetTokenbyTrade"
	TokenByTradeURL      = "https://ecpg.funpoint.com.tw/Merchant/GetTokenbyTrade"
	CreatePaymentTestURL = "https://ecpg-stage.funpoint.com.tw/Merchant/CreatePayment"
	CreatePaymentURL     = "https://ecpg.funpoint.com.tw/Merchant/CreatePayment"
)

type GetTokenByTradeRequestCall struct {
	Client                     *Client
	GetTokenByTradeRequest     *GetTokenByTradeRequest
	GetTokenByTradeRequestData *GetTokenByTradeRequestData
}

type CreatePaymentRequestCall struct {
	Client                   *Client
	CreatePaymentRequest     *CreatePaymentRequest
	CreatePaymentRequestData *CreatePaymentRequestData
}

type GetTokenByTradeRequest struct {
	MerchantID string `json:"MerchantID"`
	RqHeader   struct {
		Timestamp int    `json:"Timestamp"`
		Revision  string `json:"Revision"`
	} `json:"RqHeader"`
	Data string `json:"Data"`
}

type CreatePaymentRequest struct {
	MerchantID string `json:"MerchantID"`
	RqHeader   struct {
		Timestamp int    `json:"Timestamp"`
		Revision  string `json:"Revision"`
	} `json:"RqHeader"`
	Data string `json:"Data"`
}

type GetTokenByTradeRequestData struct {
	MerchantID        string `json:"MerchantID"`
	RememberCard      int    `json:"RememberCard"`
	PaymentUIType     int    `json:"PaymentUIType"`
	ChoosePaymentList string `json:"ChoosePaymentList"`
	OrderInfo         struct {
		MerchantTradeNo   string `json:"MerchantTradeNo"`
		MerchantTradeDate string `json:"MerchantTradeDate"`
		TotalAmount       int    `json:"TotalAmount"`
		ReturnURL         string `json:"ReturnURL"`
		TradeDesc         string `json:"TradeDesc"`
		ItemName          string `json:"ItemName"`
	} `json:"OrderInfo"`
	CardInfo struct {
		OrderResultURL    string `json:"OrderResultURL"`
		CreditInstallment string `json:"CreditInstallment"`
	} `json:"CardInfo"`
	UnionpayInfo struct {
		OrderResultURL string `json:"OrderResultURL"`
	} `json:"UnionpayInfo"`
	ATMInfo struct {
		ExpireDate int `json:"ExpireDate"`
	} `json:"ATMInfo"`
	ConsumerInfo struct {
		MerchantMemberID string `json:"MerchantMemberID"`
		Email            string `json:"Email"`
		Phone            string `json:"Phone"`
		Name             string `json:"Name"`
		CountryCode      string `json:"CountryCode"`
	} `json:"ConsumerInfo"`
	CustomField string `json:"CustomField"`
}

type CreatePaymentRequestData struct {
	MerchantID      string `json:"MerchantID"`
	PayToken        string `json:"PayToken"`
	MerchantTradeNo string `json:"MerchantTradeNo"`
	CVV             string `json:"CVV"`
}

type GetTokenByTradeResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type CreatePaymentResponse struct {
	MerchantID string   `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}

type GetTokenByTradeResponseData struct {
	RtnCode         int    `json:"RtnCode"`
	RtnMsg          string `json:"RtnMsg"`
	MerchantID      string `json:"MerchantID"`
	Token           string `json:"Token"`
	TokenExpireDate string `json:"TokenExpireDate"`
}

type CreatePaymentResponseData struct {
	RtnCode      int    `json:"RtnCode,omitempty"`
	RtnMsg       string `json:"RtnMsg,omitempty"`
	PlatformID   string `json:"PlatformID,omitempty"`
	MerchantID   string `json:"MerchantID,omitempty"`
	SimulatePaid int    `json:"SimulatePaid,omitempty"`
	OrderInfo    struct {
		MerchantTradeNo string  `json:"MerchantTradeNo,omitempty"`
		TradeNo         string  `json:"TradeNo,omitempty"`
		PaymentDate     string  `json:"PaymentDate,omitempty"`
		TradeAmt        int     `json:"TradeAmt,omitempty"`
		PaymentType     string  `json:"PaymentType,omitempty"`
		TradeDate       string  `json:"TradeDate,omitempty"`
		ChargeFee       float64 `json:"ChargeFee,omitempty"`
		TradeStatus     string  `json:"TradeStatus,omitempty"`
	} `json:"OrderInfo"`
	CVSInfo struct {
		PayFrom    string `json:"PayFrom,omitempty"`
		PaymentNo  string `json:"PaymentNo,omitempty"`
		PaymentURL string `json:"PaymentURL,omitempty"`
		StoreID    string `json:"StoreID,omitempty"`
		StoreName  string `json:"StoreName,omitempty"`
	} `json:"CVSInfo"`
	ATMInfo struct {
		ATMAccBank string `json:"ATMAccBank,omitempty"`
		ATMAccNo   string `json:"ATMAccNo,omitempty"`
	} `json:"ATMInfo"`
	CardInfo struct {
		AuthCode           string `json:"AuthCode,omitempty"`
		Gwsr               int    `json:"Gwsr,omitempty"`
		ProcessDate        string `json:"ProcessDate,omitempty"`
		Amount             int    `json:"Amount,omitempty"`
		Stage              int    `json:"Stage,omitempty"`
		Stast              int    `json:"Stast,omitempty"`
		Staed              int    `json:"Staed,omitempty"`
		Eci                int    `json:"Eci,omitempty"`
		Card6No            string `json:"Card6No,omitempty"`
		Card4No            string `json:"Card4No,omitempty"`
		RedDan             int    `json:"RedDan,omitempty"`
		RedDeAmt           int    `json:"RedDeAmt,omitempty"`
		RedOkAmt           int    `json:"RedOkAmt,omitempty"`
		RedYet             int    `json:"RedYet,omitempty"`
		PeriodType         string `json:"PeriodType,omitempty"`
		Frequency          int    `json:"Frequency,omitempty"`
		ExecTimes          int    `json:"ExecTimes,omitempty"`
		PeriodAmount       int    `json:"PeriodAmount,omitempty"`
		TotalSuccessTimes  int    `json:"TotalSuccessTimes,omitempty"`
		TotalSuccessAmount int    `json:"TotalSuccessAmount,omitempty"`
	} `json:"CardInfo"`
	ThreeDInfo struct {
		ThreeDURL string `json:"ThreeDURL"`
	}
	CustomField string `json:"CustomField,omitempty"`
}
type SDKResponse struct {
	MerchantID int      `json:"MerchantID"`
	RqHeader   RqHeader `json:"RqHeader"`
	TransCode  int      `json:"TransCode"`
	TransMsg   string   `json:"TransMsg"`
	Data       string   `json:"Data"`
}
type SDKResponseData struct {
	RtnCode      int    `json:"RtnCode,omitempty"`
	RtnMsg       string `json:"RtnMsg,omitempty"`
	PlatformID   string `json:"PlatformID,omitempty"`
	MerchantID   string `json:"MerchantID,omitempty"`
	SimulatePaid int    `json:"SimulatePaid,omitempty"`
	OrderInfo    struct {
		MerchantTradeNo string  `json:"MerchantTradeNo,omitempty"`
		TradeNo         string  `json:"TradeNo,omitempty"`
		PaymentDate     string  `json:"PaymentDate,omitempty"`
		TradeAmt        int     `json:"TradeAmt,omitempty"`
		PaymentType     string  `json:"PaymentType,omitempty"`
		TradeDate       string  `json:"TradeDate,omitempty"`
		ChargeFee       float64 `json:"ChargeFee,omitempty"`
		TradeStatus     string  `json:"TradeStatus,omitempty"`
	} `json:"OrderInfo"`
	CVSInfo struct {
		PayFrom    string `json:"PayFrom,omitempty"`
		PaymentNo  string `json:"PaymentNo,omitempty"`
		PaymentURL string `json:"PaymentURL,omitempty"`
		StoreID    string `json:"StoreID,omitempty"`
		StoreName  string `json:"StoreName,omitempty"`
	} `json:"CVSInfo"`
	ATMInfo struct {
		ATMAccBank string `json:"ATMAccBank,omitempty"`
		ATMAccNo   string `json:"ATMAccNo,omitempty"`
	} `json:"ATMInfo"`
	CardInfo struct {
		AuthCode           string `json:"AuthCode,omitempty"`
		Gwsr               int    `json:"Gwsr,omitempty"`
		ProcessDate        string `json:"ProcessDate,omitempty"`
		Amount             int    `json:"Amount,omitempty"`
		Stage              int    `json:"Stage,omitempty"`
		Stast              int    `json:"Stast,omitempty"`
		Staed              int    `json:"Staed,omitempty"`
		Eci                int    `json:"Eci,omitempty"`
		Card6No            string `json:"Card6No,omitempty"`
		Card4No            string `json:"Card4No,omitempty"`
		RedDan             int    `json:"RedDan,omitempty"`
		RedDeAmt           int    `json:"RedDeAmt,omitempty"`
		RedOkAmt           int    `json:"RedOkAmt,omitempty"`
		RedYet             int    `json:"RedYet,omitempty"`
		PeriodType         string `json:"PeriodType,omitempty"`
		Frequency          int    `json:"Frequency,omitempty"`
		ExecTimes          int    `json:"ExecTimes,omitempty"`
		PeriodAmount       int    `json:"PeriodAmount,omitempty"`
		TotalSuccessTimes  int    `json:"TotalSuccessTimes,omitempty"`
		TotalSuccessAmount int    `json:"TotalSuccessAmount,omitempty"`
	} `json:"CardInfo"`
	ThreeDInfo struct {
		ThreeDURL string `json:"ThreeDURL"`
	}
	CustomField string `json:"CustomField,omitempty"`
}

func (c *Client) NewGetTokenByTrade() *GetTokenByTradeRequestCall {
	r := new(GetTokenByTradeRequestCall)
	r.Client = c
	return r
}

func (c *Client) NewCreatePayment() *CreatePaymentRequestCall {
	r := new(CreatePaymentRequestCall)
	r.Client = c
	return r
}

func (c *GetTokenByTradeRequestCall) Request() *GetTokenByTradeRequestCall {
	c.GetTokenByTradeRequest = &GetTokenByTradeRequest{}
	c.GetTokenByTradeRequestData = &GetTokenByTradeRequestData{}
	return c
}

func (c *CreatePaymentRequestCall) Request() *CreatePaymentRequestCall {
	c.CreatePaymentRequest = &CreatePaymentRequest{}
	c.CreatePaymentRequestData = &CreatePaymentRequestData{}
	return c
}

func (c *GetTokenByTradeRequestCall) GetTokenByTrade() *GetTokenByTradeRequestCall {
	c.GetTokenByTradeRequest.MerchantID = c.Client.merchantID
	c.GetTokenByTradeRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.GetTokenByTradeRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.GetTokenByTradeRequestData)
	fmt.Println(string(jsonData))
	c.GetTokenByTradeRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *CreatePaymentRequestCall) CreatePayment() *CreatePaymentRequestCall {
	c.CreatePaymentRequest.MerchantID = c.Client.merchantID
	c.CreatePaymentRequest.RqHeader.Timestamp = int(time.Now().Unix())
	c.CreatePaymentRequest.RqHeader.Revision = "1.0.0"
	jsonData, _ := json.Marshal(c.CreatePaymentRequestData)
	fmt.Println(string(jsonData))
	c.CreatePaymentRequest.Data = Aes128(url.QueryEscape(string(jsonData)), c.Client.hashKey, c.Client.hashIV)
	return c
}

func (c *GetTokenByTradeRequestCall) Do(URL string) (*GetTokenByTradeResponseData, error) {
	marshal, err := json.Marshal(c.GetTokenByTradeRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(GetTokenByTradeResponse)
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
		responseData := new(GetTokenByTradeResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}

func (c *CreatePaymentRequestCall) Do(URL string) (*CreatePaymentResponseData, error) {
	marshal, err := json.Marshal(c.CreatePaymentRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(marshal))
	request, err := SendRequest(marshal, URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	response := new(CreatePaymentResponse)
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
		responseData := new(CreatePaymentResponseData)
		err = json.Unmarshal([]byte(unescape), responseData)
		if err != nil {
			return nil, err
		}
		return responseData, nil
	}
}
