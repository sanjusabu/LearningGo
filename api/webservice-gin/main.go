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
	}
	body, err := ioutil.ReadAll(res.Body)

	assetNameRegex := regexp.MustCompile(`<a[^>]*>([^<]+)<br />`)
	tickerRegex := regexp.MustCompile(`>([^<]+)&nbsp;&nbsp;&nbsp;<strong>`)
	codeRegex := regexp.MustCompile(`<strong>([^<]+)</strong>`)
	pincodeRegex := regexp.MustCompile(`/(\d+)/'`)

	// Find matches using regular expressions
	assetNameMatches := assetNameRegex.FindStringSubmatch(string(body))
	tickerMatches := tickerRegex.FindStringSubmatch(string(body))
	codeMatches := codeRegex.FindStringSubmatch(string(body))
	pincodeMatches := pincodeRegex.FindStringSubmatch(string(body))

	// Extract data from matches
	assetName := ""
	ticker := ""
	code := ""
	pincode := ""

	if len(assetNameMatches) > 1 {
		assetName = strings.TrimSpace(assetNameMatches[1])
	}

	if len(tickerMatches) > 1 {
		ticker = strings.TrimSpace(tickerMatches[1])
	}

	if len(codeMatches) > 1 {
		code = strings.TrimSpace(codeMatches[1])
	}
	if len(pincodeMatches) > 1 {
		pincode = strings.TrimSpace(pincodeMatches[1])
	}

	// fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	data := struct {
		Company string `json:"company"`
		Ticker  string `json:"ticker"`
		Code    string `json:"code"`
		Pincode string `json:"pincode"`
	}{
		Company: assetName,
		Ticker:  ticker,
		Code:    code,
		Pincode: pincode,
	}
	// converting to json
	jsonData, jerr := json.MarshalIndent(data, "", "  ")
	if jerr != nil {
		fmt.Println("Error marshaling JSON:", jerr)
		return
	}
	fmt.Println(string(jsonData))
}
