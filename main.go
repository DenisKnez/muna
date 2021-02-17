package main

import (
	"net/http"

	"github.com/DenisKnez/muna/handlers"
	"github.com/DenisKnez/muna/repositories"
 	"github.com/DenisKnez/muna/handlers/routers"
	"github.com/DenisKnez/muna/services"
	"github.com/DenisKnez/muna/util"
)

func main() {


	sm := http.NewServeMux()

	ggRepo := repositories.NewGuessGameRepository(util.Db)
	ggService := services.NewGuessGameService(ggRepo)
	ggHandler := handlers.NewGuessGameHandler(ggService)


	routers.GuessGameRouter(sm, ggHandler)


	http.ListenAndServe(":9090", sm)

}
