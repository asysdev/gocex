package models

type BinanceSymbols struct {
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
}

type Binance struct {
	TimeZone        string           `json:"timezone"`
	ServerTime      int              `json:"serverTime"`
	RateLimits      interface{}      `json:"rateLimits"`
	ExchangeFilters interface{}      `json:"exchangeFilters"`
	Symbols         []BinanceSymbols `json:"symbols"`
}
