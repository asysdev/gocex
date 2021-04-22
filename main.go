package main

import (
	"log"

	"github.com/asysdev/gocex/service"
)

func main() {
	log.Println("CEX Listing Code started")

	service.FetchBinancePairs()

	service.FetchFtxPairs()

	log.Println("CEX Listing Code ends")

}
