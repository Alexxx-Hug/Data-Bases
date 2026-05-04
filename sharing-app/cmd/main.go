package main

import (
	"log"
	"sharing-app/internal/database"
	"sharing-app/internal/handler"
	"sharing-app/internal/repository"
	"sharing-app/internal/service"
)

func main() {
	cfg := database.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "alex030106",
		DBName:   "sharing_db",
	}

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	log.Println("Connected to PostgreSQL successfully")

	userRepo := repository.NewUserRepository(db)
	tripRepo := repository.NewTripRepository(db)

	userService := service.NewUserService(userRepo)
	tripService := service.NewTripService(tripRepo)

	promotionRepo := repository.NewPromotionRepository(db)
	promotionService := service.NewPromotionService(promotionRepo)

	cliHandler := handler.NewCLIHandler(
		userService,
		tripService,
		promotionService,
	)
	cliHandler.Run()
}
