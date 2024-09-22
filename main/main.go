package main

import (
	"github.com/sherryMiet/go-funpoint-sdk"
	"time"
)

var (
	Loc, _ = time.LoadLocation("Asia/Taipei")
)

func main() {
	client := funpoint.NewStageClient()
	r := client.NewCreateBindCard().Request()
	r.CreateBindCardRequestData = &funpoint.CreateBindCardRequestData{
		MerchantID:       r.Client.MerchantID(),
		BindCardPayToken: "7aae762a2d514b6386e0becfa246cf13",
		MerchantMemberID: "cbe0f04c-1627-11ed-a2de-42010a32b004",
	}
	do, err := r.CreateBindCard().Do(funpoint.CreateBindCardTestURL)
	//if err != nil {
	//	return
	//}

	//r := client.NewGetTokenByBindingCard().Request()
	//r.GetTokenByBindingCardRequestData = &funpoint.GetTokenByBindingCardRequestData{
	//	MerchantID: r.Client.MerchantID(),
	//	ConsumerInfo: funpoint.ConsumerInfo{
	//		MerchantMemberID: "cbe0f04c-1627-11ed-a2de-42010a32b004",
	//		Email:            "sherry2000307@gmail.com",
	//		Phone:            "0909508777",
	//		Name:             "YU YU JU",
	//		CountryCode:      "158",
	//	},
	//	OrderInfo: funpoint.OrderInfo{
	//		MerchantTradeNo:   strconv.Itoa(int(time.Now().Unix())),
	//		MerchantTradeDate: time.Now().In(Loc).Format("2006/01/02 15:04:05"),
	//		TotalAmount:       100,
	//		ReturnURL:         "https://cashflow.ramatetech.com/funpoint/bind_card/return",
	//		TradeDesc:         "測試用",
	//		ItemName:          "幫開課",
	//	},
	//	OrderResultURL: "https://cashflow.ramatetech.com/funpoint/bind_card/return",
	//}
	//do, err := r.GetTokenByBindingCard().Do(funpoint.GetTokenByBindingCardTestURL)
	if err != nil {
		return
	}
	print(do)

}
