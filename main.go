package main

import (
	"fmt"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("step 1")
	connStr := "user=postgres dbname=postgres password=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step 2")

	rows, err := db.Query("SELECT 1")
	if err != nil {
		panic(err)
	}
	var (
        n   int64
    )
	rows.Next()
    if err := rows.Scan(&n); err != nil {
        log.Fatal(err)
    }
    log.Printf("id %d has role\n", n)
}

func main() {
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
