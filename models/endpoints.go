package models

import "time"

type Endpoint struct {
	ID           string                 `bson:"_id,omitempty" json:"id"`
	URL          string                 `bson:"url" json:"url"`
	Method       string                 `bson:"method" json:"method"`
	ResponseBody map[string]interface{} `bson:"response_body" json:"response"` // Fix: Use `map[string]interface{}` to match JSON
	StatusCode   int                    `bson:"status_code" json:"status_code"`
	Headers      map[string]string      `bson:"headers" json:"headers"`
	CreatedAt    time.Time              `bson:"created_at" json:"created_at"`
}
