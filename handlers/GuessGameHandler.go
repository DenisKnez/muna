package handlers

import (
	"net/http"
	"github.com/DenisKnez/muna/domains"
)


//GuessGameHandler guess game handler
type GuessGameHandler struct {
	guessGameService domains.GuessGameService
}

//NewGuessGameHandler returns a new guess game handler
func NewGuessGameHandler(ggService domains.GuessGameService) *GuessGameHandler{
	return &GuessGameHandler{ggService}
}

//Check endpoint to check the validity of the string
func (ggHandler *GuessGameHandler) Check(w http.ResponseWriter, r *http.Request) {


	gameID := checkGameSession(w, r, ggHandler.guessGameService)


	//VALUE provided

	guess := r.URL.Query().Get("guess")

	if guess == "" {
		http.Error(w, "Required parameter guess was not provided", http.StatusBadRequest)
		return
	}


	ggHandler.guessGameService.Check(gameID, guess)


}

//Stat guess game enpoint to get the current state of the game as well as the history
func (ggHandler *GuessGameHandler) Stat(w http.ResponseWriter, r *http.Request) {

	gameID := checkGameSession(w, r, ggHandler.guessGameService)

	ggHandler.guessGameService.Stat(gameID)
}

func checkGameSession(w http.ResponseWriter, r *http.Request, service domains.GuessGameService) (gameID string){

	cookieName := "player"

	//GAME SESSION
	cookie, err := r.Cookie(cookieName)

	if err != nil {
		if err != http.ErrNoCookie {
			http.Error(w, "Failed to read player information", http.StatusInternalServerError)
			return
		}
		
		gameID = service.NewGame().String()

		http.SetCookie(w, &http.Cookie{
			Name: cookieName,
			Value: gameID,
		})

		return 
	}

	gameID = cookie.Value
	return
}