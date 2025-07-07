package scraper

import (
	"etsy-agent/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func ScrapeEtsy(searchTerm string) ([]model.Product, error) {
	url := fmt.Sprintf("https://www.etsy.com/search?q=%s", searchTerm)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://google.com")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch Etsy page: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var products []model.Product
	doc.Find("li.wt-list-unstyled div.v2-listing-card__info").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3").Text()
		price := s.Find("span.currency-value").Text()
		tags := []string{}

		title = strings.TrimSpace(title)
		price = strings.TrimSpace(price)

		if title != "" && price != "" {
			products = append(products, model.Product{
				Title: title,
				Price: price,
				Tags:  tags,
			})
		}
	})

	return products, nil
}
