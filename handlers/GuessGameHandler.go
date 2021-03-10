package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DenisKnez/muna/domains"
)

const (
	gameSessionCookie = "gameSession"
)

//GuessGameHandler guess game handler
type GuessGameHandler struct {
	guessGameService domains.GuessGameService
}

//NewGuessGameHandler returns a new guess game handler
func NewGuessGameHandler(ggService domains.GuessGameService) *GuessGameHandler {
	return &GuessGameHandler{ggService}
}

//Check endpoint to check the validity of the string
func (ggHandler *GuessGameHandler) Check(w http.ResponseWriter, r *http.Request) {

	gameID := ggHandler.checkGameSession(w, r)

	guess := r.URL.Query().Get("guess")

	if guess == "" {
		http.Error(w, "Required parameter guess was not provided", http.StatusBadRequest)
		return
	}

	isCorrect, err := ggHandler.guessGameService.Check(gameID, guess)

	if err != nil {
		http.Error(w, "Failed service", http.StatusInternalServerError)
		return
	}

	if isCorrect {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The guess was correct, you won!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("The guess '%s' was wrong, try to guess again", guess)))
}

//Stat guess game enpoint to get the current state of the game as well as the history
func (ggHandler *GuessGameHandler) Stat(w http.ResponseWriter, r *http.Request) {

	gameID := ggHandler.checkGameSession(w, r)

	info, err := ggHandler.guessGameService.Stat(gameID)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to get the information", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	jsonResult, err := json.Marshal(info)

	if err != nil {
		http.Error(w, "Failed json encoding", http.StatusInternalServerError)
	}

	w.Write(jsonResult)
}

//check if this is a new game (looks if the game cookie exists). If it's the first game creates a new game in the database
func (ggHandler *GuessGameHandler) checkGameSession(w http.ResponseWriter, r *http.Request) (gameID string) {

	//GAME SESSION
	cookie, err := r.Cookie(gameSessionCookie)

	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "Failed to read player information", http.StatusInternalServerError)
			return
		}

		gameID = ggHandler.guessGameService.NewGame().String()

		http.SetCookie(w, &http.Cookie{
			Name:  gameSessionCookie,
			Value: gameID,
		})

		return
	}

	gameID = cookie.Value
	return
}
