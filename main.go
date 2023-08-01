package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"transactions-crud/controllers"
)

func main() {
	// Conectarse a la base de datos SQLite.
	db, err := sql.Open("sqlite3", "transactions.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Crear la tabla "transactions" si aún no existe.
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY,
		amount REAL NOT NULL,
		category TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Crear los controladores.
	transactionController := controllers.NewTransactionController(db)

	// Configurar las rutas.
	r := mux.NewRouter()
	r.HandleFunc("/transactions", transactionController.CreateTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions/{id}", transactionController.GetTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions/{id}", transactionController.UpdateTransactionHandler).Methods("PUT")

	// Ejecutar el servidor.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf(":%s", port)
	log.Printf("Server listening on %s\n", address)
	log.Fatal(http.ListenAndServe(address, r))
}
