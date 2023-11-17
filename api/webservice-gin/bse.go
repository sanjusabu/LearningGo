package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func main() {
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
}
