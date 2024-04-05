package domain

import "time"

type ReleaseDistribution struct {
	Type     string `json:"type"`
	Quantity int64  `json:"quantity"`
}

type Release struct {
	Artist       string                `json:"artist"`
	Title        string                `json:"title"`
	Genre        string                `json:"genre"`
	ReleaseDate  time.Time             `json:"releaseDate"`
	Distribution []ReleaseDistribution `json:"distribution"`
}
