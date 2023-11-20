package main

import (
	b "api/controllers"

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
	st := new(Stock)
	st.Ticker = "EML"
	router.GET("/bsedata", b.Calculatebse("EML", st, "EQ"))

	router.Run("localhost:8080")
}
