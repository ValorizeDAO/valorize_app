package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateContractsTable struct{}

func (m *CreateContractsTable) GetName() string {
	return "CreateContractsTable"
}

func (m *CreateContractsTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("contracts", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("name", 500).Nullable()
	table.String("ticker", 500).Nullable()
	table.String("abi", 1000).Nullable()
	table.Column("user_id").Type("int unsigned")
	table.ForeignKey("user_id").
		Reference("users").
		On("id").
		OnDelete("cascade").
		OnUpdate("cascade")
	table.WithTimestamps()
	table.Column("deployed").Type("boolean")
	table.MustExec()
}

func (m *CreateContractsTable) Down(con *sqlx.DB) {
	mysql.DropTable("contracts", con).MustExec()
}
