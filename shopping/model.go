package shopping

import "time"

type APIResponse struct {
	Timestamp time.Time  `json:"Timestamp"`
	Ack       string     `json:"Ack"`
	Build     string     `json:"Build"`
	Version   string     `json:"Version"`
	Item      Item       `json:"Item,omitempty"`
	Errors    []ErrorMsg `json:"Errors,omitempty"`
}

type ErrorMsg struct {
	ShortMessage        string `json:"ShortMessage"`
	LongMessage         string `json:"LongMessage"`
	ErrorCode           string `json:"ErrorCode"`
	SeverityCode        string `json:"SeverityCode"`
	ErrorClassification string `json:"ErrorClassification"`
}
