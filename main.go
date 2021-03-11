package main

import (
	"fmt"
	"net/http"

	"github.com/DenisKnez/muna/handlers"
	"github.com/DenisKnez/muna/handlers/routers"
	"github.com/DenisKnez/muna/repositories"
	"github.com/DenisKnez/muna/services"
	"github.com/DenisKnez/muna/util"
)

func main() {

	config := util.GetConfig()
	conn := util.GetDatabaseConnection(config)
	redis := util.GetRedisConnection(config)

	sm := http.NewServeMux()

	ggRepo := repositories.NewGuessGameRepository(conn)
	ggService := services.NewGuessGameService(ggRepo, redis)
	ggHandler := handlers.NewGuessGameHandler(ggService)

	routers.GuessGameRouter(sm, ggHandler)
	fmt.Println("server starting")
	http.ListenAndServe(":9090", sm)

}
