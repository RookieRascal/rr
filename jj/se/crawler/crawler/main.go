package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

// Fetcher is responsible for fetching web pages.
type Fetcher struct {
	client *http.Client
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher() *Fetcher {
	return &Fetcher{
		client: &http.Client{},
	}
}

// Fetch retrieves the HTML content of the given URL.
func (f *Fetcher) Fetch(url string) (string, error) {
	resp, err := f.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch url %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body from %s: %v", url, err)
	}

	return string(body), nil
}

// Crawler is responsible for managing visited URLs and crawling new ones.
type Crawler struct {
	visited map[string]bool
	mu      sync.Mutex
	wg      sync.WaitGroup
	fetcher *Fetcher
}

// NewCrawler creates a new Crawler instance.
func NewCrawler() *Crawler {
	return &Crawler{
		visited: make(map[string]bool),
		fetcher: NewFetcher(),
	}
}

// Crawl visits a URL, fetches the HTML, and extracts links to follow.
func (c *Crawler) Crawl(url string, depth int) {
	if depth <= 0 {
		return
	}

	c.mu.Lock()
	if c.visited[url] {
		c.mu.Unlock()
		return
	}
	c.visited[url] = true
	c.mu.Unlock()

	fmt.Printf("Crawling: %s\n", url)

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()

		// Fetch the URL's HTML
		htmlContent, err := c.fetcher.Fetch(url)
		if err != nil {
			log.Printf("Error fetching URL %s: %v", url, err)
			return
		}

		// Extract and crawl links
		links := ExtractLinks(htmlContent)
		for _, link := range links {
			if strings.HasPrefix(link, "http") {
				c.Crawl(link, depth-1)
			}
		}
	}()
}

// Wait waits for all goroutines to finish.
func (c *Crawler) Wait() {
	c.wg.Wait()
}

// ExtractLinks uses a regular expression to extract all href attributes in anchor tags.
func ExtractLinks(htmlContent string) []string {
	var links []string

	// Regex to match href attributes in anchor tags (simplified)
	re := regexp.MustCompile(`href="(http[s]?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(htmlContent, -1)

	// Extract the links
	for _, match := range matches {
		if len(match) > 1 {
			links = append(links, match[1])
		}
	}

	return links
}

func main() {
	// Initialize the crawler
	crawler := NewCrawler()

	// Start crawling from a given URL
	startURL := "https://example.com"
	crawler.Crawl(startURL, 2) // Depth = 2 means it will follow links twice

	// Wait for all crawling to complete
	crawler.Wait()

	fmt.Println("Crawling completed.")
}
