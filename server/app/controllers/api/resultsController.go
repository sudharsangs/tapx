package api

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

type result struct {
	ProductName        string `json:"product_name,omitempty"`
	ProductPrice       string `json:"product_price,omitempty"`
	StrikeThroughPrice string `json:"strike_through_price,omitempty"`
	ProductUrl         string `json:"product_url,omitempty"`
	ProductImageUrl    string `json:"product_image_url,omitempty"`
	ProductRating      string `json:"product_rating,omitempty"`
}

func GetScrapingResults(ctx *fiber.Ctx) error {

	var resp []result
	c := colly.NewCollector()
	searchTerm := ctx.Query("query")
	searchTerm = strings.ReplaceAll(searchTerm, " ", "+")
	amazonUrl := fmt.Sprintf("https://www.amazon.in/s?k=%s", searchTerm)
	c.OnHTML("div.s-main-slot.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.s-result-item", func(_ int, e *colly.HTMLElement) {
			tempProduct := result{}
			productName := strings.TrimSpace(e.ChildText("span.a-size-medium.a-color-base.a-text-normal"))
			productPrice := e.ChildText("span.a-price > span.a-offscreen")
			productRating := e.ChildText("i.a-icon > span.a-icon-alt")
			splitProductPrice := strings.Split(productPrice, "₹")
			if len(splitProductPrice) > 1 {
				productPrice = splitProductPrice[1]
			}
			strikeThroughPrice := e.ChildText("span.a-price.a-text-price > span.a-offscreen")
			productUrl := e.ChildAttr("a.a-link-normal.s-underline-text.s-underline-link-text.s-link-style.a-text-normal", "href")
			productUrl = fmt.Sprintf("https://www.amazon.in%s", productUrl)
			productImage := e.ChildAttr("div.a-section.aok-relative.s-image-fixed-height > img.s-image", "src")
			if productName == "" {
				return
			}
			tempProduct.ProductName = productName
			tempProduct.ProductPrice = productPrice
			tempProduct.StrikeThroughPrice = strikeThroughPrice
			tempProduct.ProductUrl = productUrl
			tempProduct.ProductImageUrl = productImage
			tempProduct.ProductRating = productRating
			if productImage != "" && productPrice != "" {
				resp = append(resp, tempProduct)
			}

		})
	})
	c.Visit(amazonUrl)
	return ctx.JSON(resp)
}
