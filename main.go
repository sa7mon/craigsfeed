package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	var searchURL string
	flag.StringVar(&searchURL, "url", "", "URL of Craigslist search")
	//flag.BoolVar(&verbose, "verbose", false, "Verbose mode")
	flag.Parse()

	if searchURL == "" {
		panic("Need a valid URL")
	}

	s := NewScraper(searchURL)
	rssItems, err := s.Scrape()
	if err != nil {
		panic(err)
	}
	now := time.Now()
	searchQuery := searchURL[strings.Index(searchURL, "?query=")+7:strings.Index(searchURL, "&")]
	feed := &feeds.Feed{
		Title:       "Craigslist Search",
		Link:        &feeds.Link{Href: ""},
		Description: fmt.Sprintf("Craigslist search for '%v'", searchQuery),
		Author:      &feeds.Author{Name: "", Email: ""},
		Created:     now,
		Items: rssItems,
	}

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rss)
}

type scraper struct {
	searchURL string
	feedItems []feeds.Item
}

func NewScraper(url string) scraper {
	return scraper{searchURL: url}
}

func (sc scraper) Scrape() ([]*feeds.Item, error) {
	// Request the HTML page.
	res, err := http.Get(sc.searchURL)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	var parsedItems []*feeds.Item

	// Find the review items
	doc.Find("ul.rows li").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a.result-title").Text()
		pageTime, timeExists := s.Find("time.result-date").Attr("title")
		if timeExists == false {
			log.Println("Couldn't find time")
			pageTime = ""
		}
		timeParsed, err := time.Parse("2006 Mon 02 Jan 04:04:05 PM", "2021 " + pageTime) // Todo add current year instead of hardcode
		if err != nil {
			log.Println(err)
		}
		resultLink, linkExists := s.Find("a.result-title").Attr("href")
		if linkExists == false {
			resultLink = ""
		}
		price := s.Find(".result-info .result-price").Text()
		location := strings.TrimSpace(s.Find(".result-info .result-hood").Text())

		var item feeds.Item
		item = feeds.Item{
			Title:       title,
			Created:     timeParsed,
			Link:        &feeds.Link{Href: resultLink},
			Description: fmt.Sprintf("%v | %v | %v", title, price, location),
			Author:      &feeds.Author{Name: "", Email: ""},
		}
		parsedItems = append(parsedItems, &item)
	})

	return parsedItems, nil
}