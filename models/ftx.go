package models

type FtxSymbols struct {
	BaseAsset  string `json:"baseCurrency"`
	QuoteAsset string `json:"quoteCurrency"`
}

type Ftx struct {
	Success bool         `json:"success"`
	Symbols []FtxSymbols `json:"result"`
}
