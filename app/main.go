package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	_ "net"
	_ "os"
	_ "fmt"
	_ "context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	CONNECTION_STRING = "user=logan password=leauxgan1 dbname=jsframeworks sslmode=disable"
	SOCKET_HOST = "localhost"
	SOCKET_PORT = "5000"
	SERVER_TYPE = "tcp"
) 

func getFrameworks(w http.ResponseWriter, r *http.Request) {

}



func main() {
	var err error
	db, err = sql.Open("postgres",CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	port := "3000"

	app := chi.NewRouter()
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.Timeout(60 * time.Second))

	fs := http.FileServer(http.Dir("./static/"))
	app.Handle("/game/*", http.StripPrefix("/game/",fs))

	app.Get("/hi",func(w http.ResponseWriter,r *http.Request) {
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			log.Fatal(err)
		}
	})

	app.Get("/frameworks", getFrameworks)

	log.Printf("Started server on localhost:%s",port)
	err = http.ListenAndServe(":" + port,app)
	if err != nil {
		log.Fatal(err)
	}


}

