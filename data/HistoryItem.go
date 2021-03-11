package data

import (
	"time"
)

//HistoryItem item that contains played history information
type HistoryItem struct {
	TimeStamp time.Time
	Value     string
}
