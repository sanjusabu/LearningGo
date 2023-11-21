package controllers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Stock struct {
	Ticker string       `json:"ticker"`
	NSE    []SymbolInfo `json:"nse"`
	BSE    BseStruct    `json:"bse"`
}
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

type Response struct {
	Symbols []SymbolInfo `json:"symbols"`
}

func Calculatense(ticker string, st *Stock) {

	url := fmt.Sprintf("https://www.nseindia.com/api/search/autocomplete?q=%s", ticker)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "PostmanRuntime/7.35.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	res, errs := client.Do(req)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}
	var responseJSON []byte
	if res.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(res.Body)
		if err != nil {
			fmt.Println("Error creating gzip reader:", err)
			return
		}
		defer reader.Close()
		responseJSON, err = ioutil.ReadAll(reader)
	} else {
		responseJSON, _ = ioutil.ReadAll(res.Body)
	}

	// Create an instance of the struct
	var jsonResponse Response

	// Unmarshal the JSON response into the struct
	jsonerr := json.Unmarshal([]byte(responseJSON), &jsonResponse)

	if jsonerr != nil {
		fmt.Println("Error:", jsonerr)
		return
	}
	var ans []SymbolInfo

	// Access the desired fields
	if len(jsonResponse.Symbols) > 0 {
		for _, sym := range jsonResponse.Symbols {
			ans = append(ans, sym)
		}
		st.NSE = ans

	} else {
		// fmt.Println("No symbols found in the NSE response for the", ticker)
		return
	}
}

func Calculatebse(ticker string, st *Stock, querytype string) {

	url := fmt.Sprintf("https://api.bseindia.com/Msource/1D/getQouteSearch.aspx?Type=%sext=%s&flag=site", querytype, ticker)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Origin", "https://www.bseindia.com")
	req.Header.Add("Referer", "https://www.bseindia.com")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "PostmanRuntime/7.34.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	res, errs := client.Do(req)

	// fmt.Print(res.Body)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	symbolinfoRegex := regexp.MustCompile(`<a[^>]*>([^<]+)<br />`)
	symbolRegex := regexp.MustCompile(`>([^<]+)&nbsp;&nbsp;&nbsp;<strong>`)
	codeRegex := regexp.MustCompile(`<strong>([^<]+)</strong>`)
	bsecodeRegex := regexp.MustCompile(`/(\d+)/'`)

	symbolinfoMatches := symbolinfoRegex.FindStringSubmatch(string(body))
	symbolMatches := symbolRegex.FindStringSubmatch(string(body))
	codeMatches := codeRegex.FindStringSubmatch(string(body))
	bsecodeMatches := bsecodeRegex.FindStringSubmatch(string(body))

	symbolinfo := ""
	symbol := ""
	isin := ""
	bsecode := ""

	if len(symbolinfoMatches) > 1 {
		symbolinfo = strings.TrimSpace(symbolinfoMatches[1])
	}

	if len(symbolMatches) > 1 {
		symbol = strings.TrimSpace(symbolMatches[1])
	}

	if len(codeMatches) > 1 {
		isin = strings.TrimSpace(codeMatches[1])
	}
	if len(bsecodeMatches) > 1 {
		bsecode = strings.TrimSpace(bsecodeMatches[1])
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	data := BseStruct{
		Symbol_info: symbolinfo,
		Symbol:      symbol,
		ISIN:        isin,
		Bsecode:     bsecode,
	}
	st.BSE = data

}
