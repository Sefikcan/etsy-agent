package main

import (
	"etsy-agent/analyzer"
	"etsy-agent/scraper"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	_ = godotenv.Load()

	searchTerm := "digital-products"

	fmt.Println("🔍 Etsy scraping is starting...")
	products, err := scraper.ScrapeEtsy(searchTerm)
	if err != nil {
		log.Fatal("Scraping err:", err)
	}

	fmt.Printf("📦 %d products found\n", len(products))

	result, err := analyzer.AnalyzeProducts(products)
	if err != nil {
		log.Fatal("GPT analyze error:", err)
	}

	fmt.Println("\\n📊 GPT Analyze Result:\\n")
	fmt.Println(result)
}
