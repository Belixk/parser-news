package parser

import "time"

type News struct {
	ID          string    `json:"id" db:"id"`                   // Unique identifier for database storage
	Title       string    `json:"title" db:"title"`             // News article title
	Description string    `json:"description" db:"description"` // News article description/summary
	Link        string    `json:"link" db:"link"`               // URL to the full article
	Source      string    `json:"source" db:"source"`           // News source (TASS or Washington Post)
	Date        time.Time `json:"date" db:"date"`               // Publication date and time
	// Category    string    `json:"category" db:"category"`     // Optional category field for future use
}

type RSS struct {
	Channel Channel `xml:"channel"` // RSS channel container
}

type Channel struct {
	Item []Item `xml:"item"` // List of news items in the RSS feed
}

type Item struct {
	Title       string `xml:"title"`       // News item title from RSS
	Description string `xml:"description"` // News item description from RSS
	Link        string `xml:"link"`        // News item URL from RSS
	PubDate     string `xml:"pubDate"`     // Publication date string from RSS
}
