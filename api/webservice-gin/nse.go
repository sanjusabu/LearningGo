package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	checkTime := time.Now()
	url := "https://www.nseindia.com/api/search/autocomplete?q=SICALLOG"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Cookie", "ak_bmsc=75FF49D9B9B88E6756ABCE5A50BF48F9~000000000000000000000000000000~YAAQtvY3Fys7s9SLAQAAZQZ52RWGpINfLtrrMJ1wQh+dc3dFq0nk1oLcvfUkpeNCjgZJhDw/WzN8mxruHUR05RbueMCcCk+QsIIA9nTUTT8/dXn9kA5ByplSi4hAEAW1dxpEl4zjVoO3mY/ftVTmIyRTUgqJSiG9hMjvP+OH9FhsUS9vG5XziNJ+qmQU/Gh2ZukPb84xDyWQez3tgrT/cqt4Gzk6DxmxNjT9v1IiVSzw30MakTCEL1z9THq7NKInN5PmRepY147oWI3i8I5jZwdsJg8mI9mmnBx5ic3r5s5TYIvimpqb+3bpp1a0e3Ft6ju+AJgqST87HT1VA+icJVLdrwAQCb03GII6Knp5vb6uMo/CTrVNsAVbbMWM0hQ=")
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

	fmt.Println("This took operation took :", time.Since(checkTime))
}
