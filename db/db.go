package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

func GetDb(dbName string) *sql.DB {
  db, err := sql.Open("postgres", fmt.Sprintf("user=Home dbname=%s sslmode=disable", dbName))
  if err != nil {
    panic(err)
  }
  return db
}
