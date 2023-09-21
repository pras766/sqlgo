package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	conndb "github.com/pras766/sqlgo/Database"
	validate "github.com/pras766/sqlgo/Middlewares"
)

var Db *sql.DB

func main() {

	godotenv.Load(".env")

	port := os.Getenv(`PORT`)

	if port == "" {
		log.Fatal("Port enviroment variable is not set")
	}

	Db = conndb.ConnectPSQL()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	viRouter := chi.NewRouter()

	viRouter.Use(validate.CheckUsername)
	viRouter.Get("/welcome", WelcomeHandler)

	tableRouter := chi.NewRouter()
	tableRouter.Get("/newTable", CreateTableHandler)
	tableRouter.Get("/getTableData", GetTableData)
	tableRouter.Post("/insertIntoTable", InsertIntoTable)

	router.Mount("/v1", viRouter)
	router.Mount("/t1", tableRouter)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
