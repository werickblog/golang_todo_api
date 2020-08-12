package services

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// Scrapper holds methods to scrape a site
type Scrapper struct{}

// Meta define meta data type
type Meta struct {
	Image, Description, URL, Title, Site string
}

// CallWebsite make an http request to a website
func (scrapper *Scrapper) CallWebsite(websiteURL string, c *gin.Context) Meta {
	var meta Meta = Meta{
		Image:       "",
		Description: "",
		URL:         "",
		Title:       "",
		Site:        "",
	}

	client := &http.Client{
		// Set timeout to abort if the request takes too long
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", websiteURL, nil)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err})
	}

	request.Header.Set("pragma", "no-cache")

	request.Header.Set("cache-control", "no-cache")

	request.Header.Set("dnt", "1")

	request.Header.Set("upgrade-insecure-requests", "1")

	request.Header.Set("referer", websiteURL)

	// Make website request call
	resp, err := client.Do(request)

	// If we have a successful request
	if resp.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": err})
		}

		// Map through all meta tags fetched
		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			metaProperty, _ := s.Attr("property")
			metaContent, _ := s.Attr("content")

			if metaProperty == "og:site_name" || metaProperty == "twitter:site" {
				meta.Site = metaContent
			}

			if metaProperty == "og:url" {
				meta.URL = metaContent
			}

			if metaProperty == "og:image" || metaProperty == "twitter:image" {
				meta.Image = metaContent
			}

			if metaProperty == "og:title" || metaProperty == "twitter:title" {
				meta.Title = metaContent
			}

			if metaProperty == "og:description" || metaProperty == "twitter:description" {
				meta.Description = metaContent
			}
		})
	}

	return meta
}
