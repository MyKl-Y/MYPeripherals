// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

// db/models.go

package db

import (
	"time"
)

type Metric struct {
	ID              int32
	Keypresses      int32
	MouseClicks     int32
	MouseDistanceIn float64
	MouseDistanceMi float64
	ScrollSteps     int32
	Timestamp       time.Time
}

type KeypressDetail struct {
	ID        int32
	Key       uint16 //string
	Count     int32
	Timestamp time.Time
}