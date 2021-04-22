package service

import (
	"encoding/json"
	"log"

	"github.com/asysdev/gocex/models"
	"github.com/asysdev/gocex/utils"
)

func FetchBinancePairs() {

	binance := &models.Binance{}

	err := PairsFromCex(utils.BinanceEndpoint, binance)
	if err != nil {
		return
	}

	file := utils.CreateFile(utils.BinanceFile)
	if file == nil {
		return
	}
	defer file.Close()
	for _, symbol := range binance.Symbols {
		file.WriteString(symbol.BaseAsset + "/" + symbol.QuoteAsset + "\r\n")
	}

	log.Println("Stored Binance Pair at: ", file.Name())
}

func FetchFtxPairs() {

	ftx := &models.Ftx{}

	err := PairsFromCex(utils.FtxEndpoint, ftx)
	if err != nil {
		return
	}

	file := utils.CreateFile(utils.FtxFile)
	if file == nil {
		return
	}
	defer file.Close()
	for _, symbol := range ftx.Symbols {

		if symbol.BaseAsset != "" && symbol.QuoteAsset != "" {
			file.WriteString(symbol.BaseAsset + "/" + symbol.QuoteAsset + "\r\n")
		}
	}

	log.Println("Stored Ftx Pair at: ", file.Name())
}

func PairsFromCex(endpoint string, model interface{}) error {

	rest := &RestAPI{
		URI: endpoint,
	}

	if err := rest.Get(); err != nil {
		log.Println("Error getting pairs from binance", err)
		return err
	}

	defer rest.Body.Close()
	decoder := json.NewDecoder(rest.Body)
	if err := decoder.Decode(model); err != nil {
		log.Println("Error getting pairs from binance", err)
		return err
	}

	return nil
}
