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
	operationCode := analyzeMessageToken(messageToken)
	print(operationCode)
	return "getUserPhotoKeyword"

}

func analyzeMessageToken(messageToken []string) int {
	operationCode := -1
	for tokenIndex:= 0; tokenIndex < len(messageToken); tokenIndex++ {
		for opCodeIndex := 0; opCodeIndex < len(opCodeDefine); opCodeIndex++ {
			if(messageToken[tokenIndex] == opCodeDefine[opCodeIndex].keyword)
				operationCode = opCodeDefine[opCodeIndex].opCode
		}
	}
	return 1123
}
