package models

import (
  "database/sql"
  _ "github.com/lib/pq"
  "time"
  "github.com/dskoda1/godb/db"
)

type Player struct {
  id int
  name string
  position_id int
  created_at time.Time
  updated_at time.Time

  db *sql.DB
}

func PlayerDAO() *Player {
  db := db.GetDb("fantasy-football_development")
  p := Player{}
  p.db = db
  return &p
}

func (this *Player) Save() {
  this.db.Query("INSERT INTO players VALUES ('%s')")
}

func (this *Player) GetAll() {

}
