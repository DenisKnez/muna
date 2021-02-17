package routers

import (
	"net/http"

	"github.com/DenisKnez/muna/handlers"
)

func GuessGameRouter(sm *http.ServeMux, ggHandler *handlers.GuessGameHandler) {

	//POST
	sm.HandleFunc("/check", ggHandler.Check)

	//GET
	sm.HandleFunc("/stat", ggHandler.Stat)
}
