package eodhistoricaldata

// Dividend models a response from /api/div/:symbol
type Dividend struct {
	Date            string `json:"date"`
	DeclarationDate string `json:"declarationDate"`
	PaymentDate     string `json:"paymentDate"`
	RecordDate      string `json:"recordDate"`
	Value           string `json:"value"`
}

// IntradayData models a response from /api/intraday/:symbol
type IntradayData struct {
	AdjustedClose float64 `json:"adjusted_close"`
	Close         float64 `json:"close"`
	Datetime      string  `json:"datetime"`
	Gmtoffset     int     `json:"gmtoffset"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Open          float64 `json:"open"`
	Timestamp     int     `json:"timestamp"`
	Volume        int     `json:"volume"`
}

// Price models a response from /api/real-time/:symbol
type Price struct {
	Change           float64 `json:"change"`
	ChangePercentage float64 `json:"change_p"`
	Close            float64 `json:"close"`
	Code             string  `json:"code"`
	GMTOffset        int     `json:"gmtoffset"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Open             float64 `json:"open"`
	PreviousClose    float64 `json:"previousClose"`
	Timestamp        int     `json:"timestamp"`
	Volume           int     `json:"volume"`
}

// PriceEOD models a response from /api/eod/:symbol
type PriceEOD struct {
	AdjustedClose float64 `json:"adjusted_close"`
	Close         float64 `json:"close"`
	Date          string  `json:"date"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Open          float64 `json:"open"`
	Volume        int     `json:"volume"`
}

// OptionContract holds data about an option contract.
type OptionContract struct {
	ContractName      string `json:"contractName"`
	ContractSize      string `json:"contractSize"`
	Currency          string `json:"currency"`
	Type              string `json:"type"`
	InTheMoney        string `json:"inTheMoney"`
	LastTradeDateTime string `json:"lastTradeDateTime"`
	ExpirationDate    string `json:"expirationDate"`
	Strike            string `json:"strike"`
	LastPrice         string `json:"lastPrice"`
	Bid               string `json:"bid"`
	Ask               string `json:"ask"`
	Change            string `json:"change"`
	ChangePercent     string `json:"changePercent"`
	Volume            int    `json:"volume"`
	OpenInterest      int    `json:"openInterest"`
	ImpliedVolatility string `json:"impliedVolatility"`
	Delta             string `json:"delta"`
	Gamma             string `json:"gamma"`
	Theta             string `json:"theta"`
	Vega              string `json:"vega"`
	Rho               string `json:"rho"`
	Theoretical       string `json:"theoretical"`
	IntrinsicValue    string `json:"intrinsicValue"`
	TimeValue         string `json:"timeValue"`
	UpdatedAt         string `json:"updatedAt"`
}

// OptionData models a response from /api/options/:symbol
type OptionData struct {
	Code     string `json:"code"`
	Exchange string `json:"exchange"`
	Data     []*struct {
		ExpirationDate string `json:"expirationDate"`
		Options        struct {
			Put  []*OptionContract `json:"PUT"`
			Call []*OptionContract `json:"CALL"`
		} `json:"options"`
	} `json:"data"`
}

// ShortInterest models a response from /api/shorts/:symbol
type ShortInterest struct {
	Date   string `json:"date"`
	Short  int    `json:"short"`
	Volume int    `json:"volume"`
}

// Split models a repsonse from /api/splits/:symbol
type Split struct {
	Date  string `json:"date"`
	Split string `json:"split"`
}

// Symbol models an item from /api/exchanges/:exchange
type Symbol struct {
	Code     string
	Name     string
	Country  string
	Exchange string
	Currency string
	Type     string
}

// TechnicalIndicator models an item from /api/technical/:symbol
type TechnicalIndicator struct {
	Date          string  `json:"string"`
	AverageVolume float64 `json:"avgvol"`
	SMA           float64 `json:"sma"`
	EMA           float64 `json:"ema"`
	WMA           float64 `json:"wma"`
	RSI           float64 `json:"rsi"`
}
