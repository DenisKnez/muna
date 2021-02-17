package domains

import (
	"github.com/DenisKnez/muna/data"
	"github.com/google/uuid"
)

//GuessGameRepository guess game repository
type GuessGameRepository interface {
	Stat(gameID uuid.UUID) (info data.Info, err error)
	ChangeInfoState(gameID uuid.UUID)
	NewGame(infoUUID uuid.UUID, state data.State) (err error)
	LogHistory(infoUUID uuid.UUID, historyItemUUID uuid.UUID, timestamp string, guess string) (err error)
}

//GuessGameService guess game service
type GuessGameService interface {
	Check(gameID string, guess string) (ok bool, err error)
	Stat(gameID string) (info data.Info, err error)
	NewGame() uuid.UUID
}
