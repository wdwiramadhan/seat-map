package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"seat-map/internal/handler"
	"seat-map/internal/repository"
	"seat-map/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	dbConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalln("Unable to parse DATABASE_URL:", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalln("Unable to create connection pool:", err)
	}

	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	log.Println("Connected to the database!")

	seatMapRepository := &repository.SeatMapRepository{DB: db}
	seatMapService := &service.SeatMapService{SeatMapRepository: seatMapRepository}
	seatMapHandler := &handler.SeatMapHandler{SeatMapService: seatMapService}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /seat-map/{seatMapID}", seatMapHandler.GetSeatMapByID)

	log.Println("Server started on 8080")
	http.ListenAndServe(":8080", mux)
}
