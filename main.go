package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	hostname = goEnvVariable("DB_HOST")
	username = goEnvVariable("DB_USER")
	password = goEnvVariable("DB_PASSWORD")
	dbname   = goEnvVariable("DB_NAME")
)

func main() {
	connStatus := connect()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if connStatus {
			_, _ = fmt.Fprintln(w, `Hello, visitor! DB Connection successful`)
		} else {
			_, _ = fmt.Fprintln(w, `Hello, visitor! DB Connection Failed`)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connect() bool {

	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Fatal("Error Connecting to DB")
		return false
	}
	defer db.Close()

	fmt.Println("Successfully Connected!")
	return true

}

func goEnvVariable(key string) string {

	result := os.Getenv(key)

	if result == "" {
		log.Fatalln("Error loading fetching env variable")
	}

	return result
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
