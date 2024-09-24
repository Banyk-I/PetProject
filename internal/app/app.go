package app

import (
	"petProject/internal/config"
	"petProject/internal/handler"
	"petProject/internal/mongo"
	"petProject/internal/redisDB"
	"petProject/internal/repository"
	"petProject/internal/server"
	"petProject/internal/service"
)

func Run(configPath string) {
	//config
	configContainer, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	//DB
	mongoDB := mongo.ConnectMongoDB()
	redisClient := redisDB.ConnectRedisDB()

	//controller
	repo := repository.NewRepository(mongoDB, redisClient)
	svc := service.NewService(repo)
	controller := handler.NewHandler(svc)

	//server
	srv := server.NewServer(controller, configContainer.Port)
	err = srv.Start()
	if err != nil {
		panic(err)
	}
}
