package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"transactions/controllers"
)

func main() {
	db, err := sql.Open("sqlite3", "transactions.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
        id INTEGER PRIMARY KEY,
        amount REAL NOT NULL,
        category TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	transactionController := controllers.NewTransactionController(db)

	r := gin.Default()

	r.POST("/transactions", transactionController.CreateTransactionHandler)
	r.GET("/transactions/:id", transactionController.GetTransactionHandler)
	r.PUT("/transactions/:id", transactionController.UpdateTransactionHandler)
	r.DELETE("/transactions/:id", transactionController.DeleteTransactionHandler)
	r.GET("/transactions", transactionController.GetAllTransactionsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf(":%s", port)
	log.Printf("Server listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
