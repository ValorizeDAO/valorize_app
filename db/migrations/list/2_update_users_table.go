package list

import (
  mysql "github.com/ShkrutDenis/go-migrations/builder"
  "github.com/jmoiron/sqlx"
)

type UpdateUserTable struct{}

func (m *UpdateUserTable) GetName() string {
  return "UpdateUserTable"
}

func (m *UpdateUserTable) Up(con *sqlx.DB) {
  table := mysql.ChangeTable("users", con)
  table.String("avatar", 200).Nullable()
  table.String("about", 1000).Nullable()

  table.MustExec()
}

func (m *UpdateUserTable) Down(con *sqlx.DB) {
  table := mysql.ChangeTable("users", con)
  table.Column("avatar").Drop()
  table.Column("about").Drop()
  table.MustExec()
}

