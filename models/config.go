package models

import (
	"os"

	"github.com/tesh254/golang_todo_api/db"
)

// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DATABASE")

// Database name
var databaseName = "bookmarker"

// Create a connection
var dbConnect = db.NewConnection(server, databaseName)
