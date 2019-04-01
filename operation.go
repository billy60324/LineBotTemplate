package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	response := ""

	if operationCode == -1 {
		response = dbSearchLearnTable(messageToken)
	} else {
		if operationCode == 2 { // 學習 : 附加使用者名稱
			messageToken = []string{messageToken[0], messageToken[1], messageToken[2], profile.DisplayName}
		} else if operationCode == 7 { // 股票
			messageToken = []string{messageToken[0], profile.UserID}
		} else if operationCode == 8 { // 股票 2867
			messageToken = []string{messageToken[0], messageToken[1], profile.UserID}
		}
		response = coreOperation(operationCode, messageToken)
	}
	//operationCode := getOperationCode(messageToken)
	//operationCode := analyzeMessageToken(messageToken)
	log.Print("operation code:" + strconv.Itoa(operationCode) + "/location:" + strconv.Itoa(location))
	return response
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

//-------------------------------------------Simple Factory Pattern------------------------------------------------

func coreOperation(opCode int, messageToken []string) string {
	Operator := newOperateFactory().createOperate(opCode)
	return Operator.operate(messageToken)
}

func (*operateFactory) createOperate(operatename int) operater {
	switch operatename {
	case TeachKeyword:
		return &teachKeyword{}
	case ForgetKeyword:
		return &forgetKeyword{}
	case WhatIs:
		return &findKeywordDetail{}
	case IsWhat:
		return &findKeywordDetail{}
	case SearchStock:
		return &searchStock{}
	case GetFollowStock:
		return &getFollowStock{}
	case SetFollowStock:
		return &setFollowStock{}
	case NumberYesNo:
		return &randomYesNo{}
	case YesNo:
		return &randomYesNo{}
	case WantOrNot:
		return &randomYesNo{}
	default:
		//panic("无效运算符号")
		return nil
	}
}

type operateFactory struct {
}

func newOperateFactory() *operateFactory {
	return &operateFactory{}
}

type operater interface {
	operate([]string) string
}

type teachKeyword struct {
}

func (*teachKeyword) operate(messageToken []string) string {
	keyword := []string{messageToken[1]}
	teacher := messageToken[3]
	response := dbSearchLearnTable(keyword)
	now := time.Now()
	local, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		log.Print(err)
	}

	if response == "" {
		dbInsertLearnTable(keyword[0], messageToken[2], teacher, now.In(local).Format("2006-01-02 15:04:05"))
		response = "摁摁~原來如此~學到了呢!"
	} else {
		response = "我早就學會啦!"
	}

	return response
}

type forgetKeyword struct {
}

func (*forgetKeyword) operate(messageToken []string) string {
	keyword := []string{messageToken[1]}
	response := dbSearchLearnTable(keyword)

	if response == "" {
		response = keyword[0] + "是什麼??"
	} else {
		dbDeleteLearnTable(keyword[0])
		response = "好啦!人家忘記了!"
	}

	return response
}

type findKeywordDetail struct {
}

func (*findKeywordDetail) operate(messageToken []string) string {
	detail := ""
	var keyword string
	if messageToken[0] == "什麼是" {
		keyword = messageToken[1]
	} else {
		keyword = messageToken[0]
	}

	response, teacher, timestamp := dbSearchKeywordDetail(keyword)

	if response != "" {
		detail = keyword + "是" + response + "，" + teacher + "在" + strings.Replace(strings.Replace(timestamp, "T", " ", -1), "Z", "", -1) + "教我的"
	} else {
		detail = "我不知道" + keyword + "是什麼"
	}
	return detail
}

type searchStock struct {
}

func (*searchStock) operate(messageToken []string) string {
	response := ""
	jsonStockMessage := httpSearchStock(messageToken[1])
	log.Print(jsonStockMessage)
	if jsonStockMessage == "" {
		response = "找不到股票資訊"
		log.Print("JSON Format is wrong!")
	} else {
		var stockInformation StockInformation
		json.Unmarshal([]byte(jsonStockMessage), &stockInformation)
		log.Print(stockInformation)
		response = "***" + stockInformation.N + "***\n時間:" + stockInformation.T + "\n當盤成交價:" + stockInformation.Z + "\n今日最高:" + stockInformation.H + "\n今日最低:" + stockInformation.L + "\n開盤價:" + stockInformation.O + "\n昨收價:" + stockInformation.Y
	}
	return response
}

type getFollowStock struct {
}

func (*getFollowStock) operate(messageToken []string) string {
	response := ""
	stockNumber := messageToken[1]
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
	jsonStockMessage := apiResponse[strings.Index(apiResponse, "[")+1 : strings.Index(apiResponse, "]")]
	log.Print(jsonStockMessage)
	if jsonStockMessage == "" {
		response = "找不到股票資訊"
		log.Print("JSON Format is wrong!")
	} else {
		var stockInformation StockInformation
		json.Unmarshal([]byte(jsonStockMessage), &stockInformation)
		log.Print(stockInformation)
		response = "***" + stockInformation.N + "***\n時間:" + stockInformation.T + "\n當盤成交價:" + stockInformation.Z + "\n今日最高:" + stockInformation.H + "\n今日最低:" + stockInformation.L + "\n開盤價:" + stockInformation.O + "\n昨收價:" + stockInformation.Y
	}
	return response
}

type setFollowStock struct {
}

func (*setFollowStock) operate(messageToken []string) string {
	response := ""
	jsonStockMessage := httpSearchStock(messageToken[1])
	log.Print(jsonStockMessage)
	if jsonStockMessage == "" {
		response = "找不到該股票"
		log.Print("JSON Format is wrong!")
	} else {
		//search user
		if dbUserExist("userstock", messageToken[2]) {
			followStock := dbGetFollowStock(messageToken[2]) + messageToken[1] + ","
			dbUpdateUserStockTable(messageToken[2], followStock)
		} else {
			dbInsertUserStockTable(messageToken[2], messageToken[1])
			response = "摁哼~插入了!"
		}
		//add stock info
	}
	return response
}

type randomYesNo struct {
}

func (*randomYesNo) operate(messageToken []string) string {
	response := ""

	for tokenIndex := 0; tokenIndex < len(messageToken); tokenIndex++ {
		if strings.Contains(messageToken[tokenIndex], "484") {
			if rand.Float32() < 0.5 {
				response = "4"
			} else {
				response = "84"
			}
			goto Response
		}
		if strings.Contains(messageToken[tokenIndex], "是不是") {
			if rand.Float32() < 0.5 {
				response = "是"
			} else {
				response = "不是"
			}
			goto Response
		}
		if strings.Contains(messageToken[tokenIndex], "要不要") {
			if rand.Float32() < 0.5 {
				response = "要"
			} else {
				response = "不要"
			}
			goto Response
		}
	}
Response:
	return response
}

/* prototype
type addOperate struct {
}

func (*addOperate) operate([]string) string {
	return rhs + lhs
}

type multipleOperate struct {
}

func (*multipleOperate) operate([]string) string {
	return rhs * lhs
}
*/
