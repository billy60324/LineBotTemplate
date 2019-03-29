package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func httpSearchStock(stockNumber string) string {
	url := "http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_$1.tw&json=1&delay=0"
	url = strings.Replace(url, "$1", stockNumber, 1)
	resp, error := http.Get(url)
	if error != nil {
		log.Print(error)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	apiResponse := string(body)
	log.Print(apiResponse)
	return apiResponse[strings.Index(apiResponse, "[")+1 : strings.Index(apiResponse, "]")]
}
