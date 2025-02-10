package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestLog struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Method     string             `bson:"method" json:"method"`
	URL        string             `bson:"url" json:"url"`
	Request    interface{}        `bson:"request" json:"request"`
	Response   interface{}        `bson:"response" json:"response"`
	StatusCode int                `bson:"status_code" json:"status_code"`
	Timestamp  time.Time          `bson:"timestamp" json:"timestamp"`
}
