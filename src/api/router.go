package main

import (
	controllers "example.com/controllers"

	"github.com/gin-gonic/gin"
)

type BseStruct struct {
	Symbol_info string `json:"symbol_info"`
	Symbol      string `json:"symbol"`
	ISIN        string `json:"isin"`
	Bsecode     string `json:"bsecode"`
}
type SymbolInfo struct {
	Symbol        string `json:"symbol"`
	SymbolInfo    string `json:"symbol_info"`
	ResultType    string `json:"result_type"`
	ResultSubType string `json:"result_sub_type"`
	URL           string `json:"url"`
}
type Stock struct {
	Ticker string       `json:"ticker"`
	NSE    []SymbolInfo `json:"nse"`
	BSE    BseStruct    `json:"bse"`
}

func main() {
	router := gin.Default()
	stock := &controllers.Stock{
		Ticker: "EML",
	}
	router.GET("/bsedata", func(c *gin.Context) {
		ticker := c.Query("ticker")
		querytype := c.Query("type")
		controllers.Calculatebse(ticker, stock, querytype)
		c.JSON(200, stock)
	})

	router.GET("/nsedata", func(c *gin.Context) {
		ticker := c.Query("ticker")
		controllers.Calculatense(ticker, stock)
		c.JSON(200, stock)
	})

	router.Run("localhost:8080")
}
