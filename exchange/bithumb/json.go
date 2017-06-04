package bithumb

//JSON DataStruct
//Public Apis
type JsonTickerData struct {
	First       float64 `json:"opening_price,string"`
	Last        float64 `json:"closing_price,string"`
	MinPrice    float64 `json:"min_price,string"`
	MaxPrice    float64 `json:"max_price,string"`
	Average     float64 `json:"average_price,string"`
	UnitsTraded float64 `json:"units_traded,string"`
	Volume1day  float64 `json:"volume_1day,string"`
	Volume7day  float64 `json:"volume_7day,string"`
	Buy         float64 `json:"buy_price,string"`
	Sell        float64 `json:"sell_price,string"`
}
type JsonTicker struct {
	Status string `json:"status"`
	Data   map[string]JsonTickerData `json:"data"`
	Date   uint64 `json:"date,string"`
}

type JsonOrderbook struct {
	Data struct {
		Asks []struct {
			Price    float64 `json:"price,string"`
			Quantity float64 `json:"quantity,string"`
		} `json:"asks"`
		Bids []struct {
			Price    float64 `json:"price,string"`
			Quantity float64 `json:"quantity,string"`
		} `json:"bids"`
		OrderCurrency   string `json:"order_currency"`
		PaymentCurrency string `json:"payment_currency"`
		Timestamp       uint64 `json:"timestamp,string"`
	} `json:"data"`
	Status string `json:"status"`
}


type JsonRecentTransactionData struct {
	Price           float64 `json:"price"`
	Total           float64 `json:"total"`
	TransactionDate string `json:"transaction_date"`
	Type            string `json:"type"`
	UnitsTraded     string `json:"units_traded"`
}

type JsonRecentTransaction struct {
	Data []JsonRecentTransactionData `json:"data"`
	Status string `json:"status"`
}

//Private Apis
type JsonAccount struct {
	Data struct {
		AccountID string `json:"account_id"`
		Balance   string `json:"balance"`
		Created   int    `json:"created"`
		TradeFee  string `json:"trade_fee"`
	} `json:"data"`
	Status string `json:"status"`
}


type JsonBalance struct {
	Data struct {
		AvailableBtc string `json:"available_btc"`
		AvailableKrw string `json:"available_krw"`
		InUseBtc     string `json:"in_use_btc"`
		InUseKrw     string `json:"in_use_krw"`
		MisuBtc      string `json:"misu_btc"`
		MisuDepoKrw  string `json:"misu_depo_krw"`
		MisuKrw      string `json:"misu_krw"`
		TotalBtc     string `json:"total_btc"`
		TotalKrw     string `json:"total_krw"`
		XcoinLast    string `json:"xcoin_last"`
	} `json:"data"`
	Status string `json:"status"`
}
