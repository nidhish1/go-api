package main
import (
	"github.com/gorilla/mux"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/nidhish1/go-api/api"
	"github.com/nidhish1/go-api/dbCon"
)




func main() {
	dbCon.db, dbCon.err = sql.Open("mysql", "docker:docker@tcp(sqlDB)/mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/posts", api.GetPosts).Methods("GET")
	router.HandleFunc("/posts", api.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", api.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", api.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", api.DeletePost).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

