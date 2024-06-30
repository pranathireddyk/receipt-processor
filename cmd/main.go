package main

import (
	"github.com/pranathireddyk/receipt-processor/internal/database"
	"github.com/pranathireddyk/receipt-processor/internal/server"
)

// main function initializes and runs the server
func main() {
	server := server.NewReceiptServer()
	db := database.NewBoltDatabase("receipts.db")
	server.DB = db
	server.Run(":8080")
	defer db.Close()
}
