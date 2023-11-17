package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	checkTime := time.Now()
	wg.Add(2)
	go func() {
		defer wg.Done()
		nse()
	}()
	go func() {
		defer wg.Done()
		bse()
	}()
	wg.Wait()
	fmt.Println("Time for the process ", time.Since(checkTime))
}
func nse() {

	checkTime := time.Now()
	url := "https://www.nseindia.com/api/search/autocomplete?q=SICALLOG"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Cookie", "ak_bmsc=905CCB4E4855E8EB4C60529ABBCC49FA~000000000000000000000000000000~YAAQtvY3F08Tu9SLAQAA5ieP2xWXFo0673ZSXUkMKuEmJGyy+T4NuUZHQ0jyF7sjglM2uB6x7BPpEbEaxzVVQpWTzjqSEUOrey82frVtaTKBlMk9gzxwZg5t7agfOP4CJ1+Ud28TfqCjjke1OVOMsEI9dOY1dir0DifASTOrXLV8bv27id8FBRMB6H3CmWZjCv3X68e65dhQgBZabrrI2c8XQXj+d6IYqRRg8TVYttNlcTgO6iwTx+CY07dA99WuXKbwaWEwSZKLUFgtFjDnVeY7dDoufQXZVnC3WZA2kFus7fA2E6U5QeG36Kx4bF9Npwms3oYrqSScPLcQiHaUohtp7Mso7LIvMLBihAgrMYU8kQMYqpORp6Ck7+ueRyU=")
	req.Header.Add("Accept", "*/*")
	// req.Header.Add("Content-Type", "text/html")
	req.Header.Add("User-Agent", "PostmanRuntime/7.35.0")
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
	responseJSON, err := ioutil.ReadAll(res.Body)
	// Example JSON response string

	// Define a struct to represent the JSON structure
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

	// Create an instance of the struct
	var jsonResponse Response

	// Unmarshal the JSON response into the struct
	jsonerr := json.Unmarshal([]byte(responseJSON), &jsonResponse)

	if jsonerr != nil {
		fmt.Println("Error:", jsonerr)
		return
	}

	// Access the desired fields
	if len(jsonResponse.Symbols) > 0 {
		var ans []SymbolInfo
		for _, sym := range jsonResponse.Symbols {
			ans = append(ans, sym)
		}
		fmt.Println(ans)

	} else {
		fmt.Println("No symbols found in the response.")
	}

	fmt.Println("Time for NSE :", time.Since(checkTime))
}

func bse() {
	checkTime := time.Now()
	url := "https://api.bseindia.com/Msource/1D/getQouteSearch.aspx?Type=EQ&text=INE464D01014&flag=site"
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
	// req.Header.Add("Content-Type", "text/html")
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
	fmt.Println("Content-Type: ", res.Header.Get("Content-Type"))
	symbolinfoRegex := regexp.MustCompile(`<a[^>]*>([^<]+)<br />`)
	symbolRegex := regexp.MustCompile(`>([^<]+)&nbsp;&nbsp;&nbsp;<strong>`)
	codeRegex := regexp.MustCompile(`<strong>([^<]+)</strong>`)
	bsecodeRegex := regexp.MustCompile(`/(\d+)/'`)

	// Find matches using regular expressions
	symbolinfoMatches := symbolinfoRegex.FindStringSubmatch(string(body))
	symbolMatches := symbolRegex.FindStringSubmatch(string(body))
	codeMatches := codeRegex.FindStringSubmatch(string(body))
	bsecodeMatches := bsecodeRegex.FindStringSubmatch(string(body))

	// Extract data from matches
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

	// fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	data := struct {
		Symbol_info string `json:"symbol_info"`
		Symbol      string `json:"symbol"`
		ISIN        string `json:"isin"`
		Bsecode     string `json:"bsecode"`
	}{
		Symbol_info: symbolinfo,
		Symbol:      symbol,
		ISIN:        isin,
		Bsecode:     bsecode,
	}
	// converting to json
	jsonData, jerr := json.MarshalIndent(data, "", "  ")
	if jerr != nil {
		fmt.Println("Error marshaling JSON:", jerr)
		return
	}
	fmt.Println(string(jsonData))
	fmt.Println("Time for BSE ", time.Since(checkTime))
}
