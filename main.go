package main

import (
	"fmt"
	"gin-project/config"
	"gin-project/internal/delivery/http/handler"
	"gin-project/internal/delivery/http/route"
	"gin-project/internal/infrastructure"
	"gin-project/internal/repository"
	"gin-project/internal/usecase"
)

func main() {

	config := config.New()

	fmt.Println(config)

	db, err := infrastructure.NewMySQL(config)

	if err != nil {
		panic(err)
	}

	router := infrastructure.NewGin()

	publicHandler := handler.NewPublicHandler()

	cardRepository := repository.NewSQLCardRepository(db)
	cardUsecase := usecase.NewCardUsecase(cardRepository)
	cardsHandler := handler.NewCardsHandler(cardUsecase)

	route.RegisterRoute(router, publicHandler, cardsHandler)

	router.Run(":" + config.Server.Port)

}
