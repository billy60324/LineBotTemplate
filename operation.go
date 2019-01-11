package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	var answer string
	doc, err := goquery.NewDocument("https://www.google.com.tw/search?q=" + keyword)
	log.Print("Google Search")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".LC201b").Each(func(i int, s *goquery.Selection) {
		answer = s.Text()
		log.Print(s.Text())
	})

	return answer
	//return "I get U"
}
