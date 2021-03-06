package domains

import (
	"github.com/DenisKnez/muna/data"
	"github.com/google/uuid"
)

//GuessGameRepository guess game repository
type GuessGameRepository interface {
	Stat(gameID uuid.UUID) (info data.Info, err error)
	ChangeInfoState(gameID uuid.UUID) (err error)
	NewGame(infoUUID uuid.UUID, state data.State) (err error)
	LogHistory(infoUUID string, historyItemUUID string, timestamp string, guess string) (err error)
	GameExists(gameID uuid.UUID) (ok bool)
}

//GuessGameService guess game service
type GuessGameService interface {
	Check(gameID uuid.UUID, guess string) (ok bool, err error)
	Stat(gameID uuid.UUID) (info data.Info, err error)
	NewGame() (newID uuid.UUID)
}
