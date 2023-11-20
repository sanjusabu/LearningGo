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

// func main() {
// 	checkTime := time.Now()

// 	tickers := []string{
// 		"INE742O01010",
// 		"INE142M01025",
// 		"INE741K07454",
// 		"INE741K07488",
// 		"INE694X20022",
// 		"INE244L07309",
// 		"INE148I07MS6",
// 		"INE530B07138",
// 		"INE192U07293",
// 		"TIMEXWATCH",
// 		"KANELOIL",
// 		"OSWALSUG",
// 		"INE0DJ201029",
// 		"INE694C01018",
// 		"SRGINFOTEC",
// 		"JKSYNTHETC",
// 		"INE450H01022",
// 		"INE008I01026",
// 		"INE071C01019",
// 		"INE829G01011",
// 		"INE259B01020",
// 		"INE274I01016",
// 		"INE627Z01019",
// 		"INE891A07037",
// 		"INE01I507299",
// 		"INE01I507398",
// 		"INE101Q07409",
// 		"INE101Q07789",
// 		"INE101Q07714",
// 		"INE872A07UC5",
// 		"INE872A07VD1",
// 		"RUBFILINTL",
// 		"SPECIAPP",
// 		"RASOYPR",
// 		"SUNILHITEC",
// 		"INE589G01011",
// 		"INE040C01022",
// 		"INE390A01017",
// 		"INE301B01020",
// 		"INE885S01018",
// 		"INE884S01011",
// 		"INE245C01019",
// 		"STANDRDBAT",
// 		"IN3120220303",
// 		"INE787G01011",
// 		"INE852S01026",
// 		"INE549I01011",
// 		"INE438H01019",
// 		"INE785C01048",
// 		"INE466H01028",
// 		"INE804D01029",
// 		"ASOCSTONE",
// 		"USHDEVINT",
// 		"ALORA",
// 		"GITANJALI",
// 		"INE146E01015",
// 		"KLRF",
// 		"VISUINTL",
// 		"NAVKETAN",
// 		"KIRLOSELEC",
// 		"NITESHEST",
// 		"INE129D01039",
// 		"INE647D01014",
// 		"INE442A01024",
// 		"INE135A01024",
// 		"INE546B01020",
// 		"INE078A01026",
// 		"INE430G01026",
// 		"INE021B01024",
// 		"INE502K01016",
// 		"FGPIND",
// 		"IN9110D01011",
// 		"INE110D01013",
// 		"INE905A01020",
// 		"INE0PSC01024",
// 		"INE112B01013",
// 		"INE450L01024",
// 		"INE881J07FJ3",
// 		"INE881J08573",
// 		"INE872A07UR3",
// 		"INE531B01014",
// 		"INE286A01017",
// 		"INE875I01010",
// 		"INE964A01019",
// 		"INE202A01022",
// 		"INE368A01021",
// 		"INE782X01033",
// 		"IN9002A01032",
// 		"INE443L08149",
// 		"INE312K01010",
// 		"APPLEIND",
// 		"RPPINFRPP",
// 		"MAJESAUTO",
// 		"INE836H01014",
// 		"INE099J01015",
// 		"PERMAGNET",
// 		"INE528H01017",
// 		"INE538H01024",
// 		"KERALACHEM",
// 		"LAKHNNATNL",
// 		"QFSL",
// 		"MVL",
// 		"TRIVENSHET",
// 		"NETWORK",
// 		"INE0JAH07047",
// 		"INE07KI07047",
// 		"INE04VS07289",
// 		"INE998Y07121",
// 		"INE0DBJ07127",
// 		"INE08XP07043",
// 		"INE928K01013",
// 		"INE770B01026",
// 		"INE851C01014",
// 		"INE489H01020",
// 		"INE849J01021",
// 		"IN9550C01010",
// 		"INE213M01024",
// 		"INE784B01035",
// 		"INE903B01023",
// 		"MODINSULAT",
// 		"INE066F01020",
// 		"IN8359U01027",
// 		"INE968G01033",
// 		"INE0HV901016",
// 		"INE0Q3R01026",
// 		"MEDICAPS",
// 		"ATVPROJ",
// 		"PRAGBOSIMI",
// 		"SPICELEC",
// 		"NIYATILEAS",
// 		"HINDTINWRK",
// 		"VBDESAIFIN",
// 		"INE02KH01019",
// 		"INE658R01011",
// 		"INE509A01012",
// 		"INE408H01012",
// 		"INE264B01020",
// 		"INE122H01027",
// 		"INE775B01025",
// 		"KILBUNENGG",
// 		"ISIBARS",
// 		"HARIGCRANK",
// 		"GALPOWTEL",
// 		"SMSTECH",
// 		"IN8134E01010",
// 		"INE215I01027",
// 		"INE299C01024",
// 		"INE00IK01029",
// 		"INE324C01038",
// 		"INE224B01024",
// 		"INE060F01015",
// 		"AMFORGEIND",
// 		"INE559B01023",
// 		"INE904G01038",
// 		"IN8463A01029",
// 		"AMBASARABH",
// 		"BAKELHYLAM",
// 		"BALURTRANS",
// 		"LAKSELECON",
// 		"MAFATLAIND",
// 		"INE093A01033",
// 		"INE005B01027",
// 		"INE942C01045",
// 		"INE052C01027",
// 		"INE453M01018",
// 		"TANAA",
// 		"INE803A01027",
// 		"INE890I01035",
// 		"INE746C01032",
// 		"INE161B01036",
// 		"INE196B01016",
// 		"IN9002A01024",
// 		"IVRCLINFRA",
// 		"EML",
// 		"TGVSL"}
// 	var wg sync.WaitGroup

// 	var ans []Stock

// 	for _, ticker := range tickers {
// 		wg.Add(2)
// 		st := new(Stock)
// 		st.Ticker = ticker
// 		go func(ticker string, st *Stock) {
// 			defer wg.Done()
// 			Calculatebse(ticker, st, "EQ")
// 		}(ticker, st)

// 		go func(ticker string, st *Stock) {
// 			defer wg.Done()
// 			Calculatense(ticker, st)
// 		}(ticker, st)

// 		wg.Wait()

// 		fmt.Println(ticker, " Processed")

//			if st.BSE.Symbol != "" || st.NSE != nil {
//				ans = append(ans, *st)
//			} else {
//				continue
//			}
//			time.Sleep(1000)
//		}
//		// fmt.Println(ans[len(ans)-1].NSE[0].Symbol)
//		result, _ := json.MarshalIndent(ans, "", "  ")
//		fmt.Println(string(result))
//		fmt.Println("Time for the process ", time.Since(checkTime))
//	}
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

	// if symbol == "" {
	// 	return
	// }

	// jsonData, jerr := json.MarshalIndent(data, "", "  ")
	// fmt.Println(jsonData)
	// if jerr != nil {
	// 	fmt.Println("Error marshaling JSON:", jerr)
	// 	return
	// }

}
