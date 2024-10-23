package libs

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func ConnectDB() {
  db, err := sqlx.Connect("sqlite3", "./migration/main.db")
  
  if err != nil {
    log.Fatal(err)
   }
  DB = db
}

func DisconnectDB() {
  DB.Close()
}

