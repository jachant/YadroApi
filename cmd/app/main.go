package main

import (
	"log"
	"os"
	"yadro/api/http/handler"
	"yadro/database"
	"yadro/pkg/server"
	"yadro/repository/ram_storage"
	"yadro/usecases/service"

	"github.com/go-chi/chi/v5"
)

func main() {

	//init database
	db := database.NewDataBase(os.Getenv("KEY"))
	//init repository
	weatherRepo := ram_storage.NewWeatherRepository(db)
	infoRepo := ram_storage.NewInfoRepository()

	//init services
	weatherSrv := service.NewWeatherService(weatherRepo)
	infoSrv := service.NewInfoService(infoRepo)

	//init Handlers
	weatherHand := handler.NewWeatherHandler(weatherSrv)
	infoHand := handler.NewInfoHndler(infoSrv)

	//init Routers
	r := chi.NewRouter()
	subRouter := chi.NewRouter()
	weatherHand.WithWeatherHandlers(subRouter)
	infoHand.WithInfoHandlers(subRouter)
	r.Mount("/info", subRouter)

	//init server
	serv := server.NewServer(":8000", r)
	if err := serv.Run(); err != nil {
		log.Fatal(err)
	}
}
