package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/line/line-bot-sdk-go/linebot"
)

func parseMessage(message string) []string {
	return strings.Split(message, " ")
}

func httpGet() string {
	resp, err := http.Get("https://tw.yahoo.com/")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	log.Print(string(body))
	return string(len(body))
}

func googleSearch(keyword string) string {
	// URL https://www.google.com.tw/search?q=
	//html := "<body>	<div id=\"div1\">DIV1</div> <div class=\"name\">DIV2</div> <span>SPAN</span> </body>"
	var answer string
	doc, err := goquery.NewDocument("https://www.google.com.tw/search?q=" + keyword)
	//doc, err := goquery.NewDocument(html)
	log.Print("Google Search")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".name").Each(func(i int, s *goquery.Selection) {
		answer = s.Text()
		log.Print(answer)
	})

	return answer
	//return "I get U"
}

func googleMapSearch(latitude float64, longitude float64) string {
	// URL https://www.google.com/maps?q=24.773911999999999,121.00657699999999
	return "https://www.google.com/maps?q=" + strconv.FormatFloat(latitude, 'f', 30, 64) + "," + strconv.FormatFloat(longitude, 'f', 30, 64)
}

func botResponse(profile *linebot.UserProfileResponse, humanRequest string) string {
	messageToken := parseMessage(humanRequest)
	operationCode, location := getOperationCode(messageToken)

	if operationCode == 99 {
		log.Fatal(strings.Replace(humanRequest, "DB", "", 1))
		answer := dbtesting("777")
		log.Fatal(answer)
	}

	//operationCode := getOperationCode(messageToken)
	//operationCode := analyzeMessageToken(messageToken)
	print(operationCode + location)
	return strconv.Itoa(operationCode)
	//return strconv.Itoa(operationCode)

}

func getOperationCode(messageToken []string) (int, int) {
	operationCode := -1
	location := -1
	messageLength := len(messageToken)
	matchKeyword := false
	for tokenIndex := 0; tokenIndex < messageLength; tokenIndex++ {
		for _, opCodeDefine := range OpCodeDefine {
			if opCodeDefine.complete {
				matchKeyword = (messageToken[tokenIndex] == opCodeDefine.keyword)
			} else {
				matchKeyword = strings.Contains(messageToken[tokenIndex], opCodeDefine.keyword)
			}

			if matchKeyword && checkSyntax(opCodeDefine.opCode, tokenIndex, messageLength) {
				operationCode = opCodeDefine.opCode
				location = tokenIndex
				goto Response
			}
		}
	}

Response:
	return operationCode, location
}

func analyzeMessageToken(messageToken []string) int {
	var operationCode = -1

	return operationCode
}

func checkSyntax(operationCode int, location int, length int) bool {
	for _, opCodeSyntaxDefine := range OpCodeSyntaxDefine {
		if operationCode == opCodeSyntaxDefine.opCode {
			if opCodeSyntaxDefine.location != NoNeed && opCodeSyntaxDefine.location != location {
				return false
			}

			if opCodeSyntaxDefine.length != NoNeed && opCodeSyntaxDefine.length != length {
				return false
			}
			return true
		}
	}
	return false
}
