package data

type State uint

const (
	StateSolved State = iota + 1
	StateUnsolved
)

//Info information about the rounds played
type Info struct {
	State   State
	History []HistoryItem
}
