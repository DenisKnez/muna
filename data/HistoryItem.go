package data

import (
	"time"
)

type State uint

const (
	StateSolved State = iota +1
	StateUnsolved
)

//HistoryItem item that contains played history information
type HistoryItem struct {
	TimeStamp time.Time
	Value string
}

//Info information about the rounds played
type Info struct {
	State State
	History []HistoryItem
}
