package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"gitlab.com/kpi-lab/microservices-demo/services/service1/handlers"
	"gitlab.com/kpi-lab/microservices-demo/services/service1/repository/postgres"
	"log"
	"net/http"
)

var (
	httpPort int
	pgHost   string
	pgUser   string
	pgPass   string
	pgDb     string
)

func init() {
	httpPort = 8080
	// pgUser = os.Getenv("POSTGRES_USER")
	pgUser = "postgres"
	// pgPass = os.Getenv("POSTGRES_PASSWORD")
	pgPass = "admin"
	// pgHost = os.Getenv("POSTGRES_HOST")
	pgHost = "localhost"
	// pgDb = os.Getenv("POSTGRES_DB")
	pgDb = "notes"
}

// type Note struct {
// 	body string `json:"body"`
// }

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConnector := fmt.Sprintf("postgres://%s:%s@%s/%s", pgUser, pgPass, pgHost, pgDb)

	conn, err := pgx.Connect(ctx, dbConnector)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	connection := postgres.New(conn)
	visits_server := handlers.NewVisitsServer(connection)

	notes_server := handlers.NewNotesServer(connection)

	r := mux.NewRouter()
	r.HandleFunc("/api/service1/ping", visits_server.Ping)
	r.HandleFunc("/api/service1/all", notes_server.GetAll).Methods("GET")
	r.HandleFunc("/api/service1/notes/", notes_server.GetNote).Methods("GET")
	r.HandleFunc("/api/service1/new", notes_server.MakeNote).Methods("POST")
	r.HandleFunc("/api/service1/notes/", notes_server.ChangeNote).Methods("PUT")
	r.HandleFunc("/api/service1/delete/", notes_server.DeleteNote).Methods("DELETE")

	fmt.Printf("Starting server at port ")
	fmt.Println(httpPort)

	err = http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r)
	if err != nil {
		log.Fatal(err)
	}
}
