package main

import (
	"database/sql"
	"fmt"
	"github.com/ehsundar/scaffold/internal/api"
	"github.com/ehsundar/scaffold/internal/storage"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	psqlconn := fmt.Sprintf("host=localhost port=5432 user=postgres password=mysecretpassword dbname=postgres sslmode=disable")
	db, err := sql.Open("postgres", psqlconn)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	s := storage.NewPostgresStorage(db)

	scaffold := api.NewScaffold(s)

	r := mux.NewRouter()
	api.RegisterRoutes(r, scaffold)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
