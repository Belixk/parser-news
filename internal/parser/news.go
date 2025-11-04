package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var cachedNews []News       // Cache for storing parsed news
var lastUpdate = time.Now() // Timestamp of last cache update

func parseRSS(link, source string) ([]News, error) {
	resp, err := http.Get(link) // Send HTTP GET request to the RSS feed URL
	if err != nil {
		return nil, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close() // Ensure response body is closed to prevent memory leaks

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: status %d - %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body) // Read the entire response body
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var rss RSS
	err = xml.Unmarshal(body, &rss) // Parse XML data into RSS struct
	if err != nil {
		return nil, fmt.Errorf("XML parsing error: %v", err)
	}

	var newList []News // Slice to store parsed news items
	for i, item := range rss.Channel.Item {
		if i == 6 { // Limit to 6 news items per source
			break
		}
		news := News{
			ID:          uuid.New().String(),
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Source:      source,
			Date:        ParsePubDate(item.PubDate),
		}
		newList = append(newList, news) // Add news item to the slice
	}

	return newList, nil // Return parsed news for further processing
}

func ParsePubDate(pubDate string) time.Time {
	// Try different date formats commonly used in RSS feeds
	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
	}

	for _, format := range formats {
		parsed, err := time.Parse(format, pubDate)
		if err == nil {
			return parsed
		}
	}
	return time.Now() // Fallback to current time if parsing fails
}

func ParseAll() ([]News, error) {
	var allNews []News
	// Return cached news if it's fresh (less than 5 minutes old) and not empty
	if time.Since(lastUpdate) <= 5*time.Minute && len(cachedNews) > 0 {
		return cachedNews, nil
	}

	// Define RSS feed sources
	sources := []struct {
		Link   string
		Source string
	}{
		{"https://tass.com/rss/v2.xml", "TASS"},
		{"https://feeds.washingtonpost.com/rss/world", "Washington"},
	}

	// Parse each news source
	for _, source := range sources {
		news, err := parseRSS(source.Link, source.Source)
		if err != nil {
			fmt.Printf("Parsing error: %v", err)
			continue // Skip failed sources but continue with others
		}
		allNews = append(allNews, news[:3]...) // Take first 3 news from each source
	}

	// Update cache with fresh data
	cachedNews = allNews
	lastUpdate = time.Now()
	return cachedNews, nil
}

func ParseBySource(source string) ([]News, error) {
	// Refresh cache if it's stale or empty
	if time.Since(lastUpdate) > 5*time.Minute || len(cachedNews) == 0 {
		ParseAll()
	}

	var result []News
	// Filter news by specified source
	for _, item := range cachedNews {
		if item.Source == source {
			result = append(result, item)
		}
	}
	return result, nil
}
