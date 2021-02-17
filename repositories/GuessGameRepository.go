package repositories

import (
	"database/sql"

	"github.com/DenisKnez/muna/data"
	"github.com/DenisKnez/muna/domains"
	"github.com/google/uuid"
)

//GuessGameRepository guess game repository
type GuessGameRepository struct {
	conn *sql.DB
}

//NewGuessGameRepository return a new guess game repository
func NewGuessGameRepository(conn *sql.DB) domains.GuessGameRepository {
	return &GuessGameRepository{conn}
}

//LogHistory history
func (ggRepo *GuessGameRepository) LogHistory(infoUUID uuid.UUID, historyItemUUID uuid.UUID, timestamp string, guess string) (err error) {
	stmt, err := ggRepo.conn.Prepare("INSERT INTO history_item (id, timestamp, value, info_id) VALUES ($1, $2, $3)")

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(historyItemUUID, timestamp, guess, infoUUID)

	if err != nil {
		return
	}

	return
}

//ChangeInfoState changes the state of the info
func (ggRepo *GuessGameRepository) ChangeInfoState(gameID uuid.UUID) {

	stmt, err := ggRepo.conn.Prepare("UPDATE info SET state = $1 VALUES WHERE id = $2")

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(data.StateSolved, gameID)

	if err != nil {
		return
	}

	return
}

//Stat returns games stats
func (ggRepo *GuessGameRepository) Stat(gameID uuid.UUID) (info data.Info, err error) {

	info = data.Info{}

	rows, err := ggRepo.conn.Query(`SELECT state, value, timestamp FROM info
	RIGHT JOIN history_item ON info.id = history_item.info_id WHERE info.id = $1`, gameID)

	if err != nil {
		return
	}

	for rows.Next() {

		historyItem := data.HistoryItem{}

		err = rows.Scan(&info.State, &historyItem.Value, &historyItem.TimeStamp)

		if err != nil {
			return
		}

		info.History = append(info.History, historyItem)
	}
	
	return
}

//NewGame creates a new info/game
func (ggRepo *GuessGameRepository) NewGame(infoUUID uuid.UUID, state data.State) (err error){
	stmt, err := ggRepo.conn.Prepare("INSERT INTO info (id, state) VALUES ($1, $2)")

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(infoUUID, state)

	if err != nil {
		return
	}

	return
}