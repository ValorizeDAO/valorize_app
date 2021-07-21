package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateWalletTable struct{}

func (m *CreateWalletTable) GetName() string {
	return "CreateWalletTable"
}

func (m *CreateWalletTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("wallets", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("address", 1000).NotNull()
	table.Column("user_id").Type("int unsigned")
	table.Column("is_contract").Type("bool")

	table.ForeignKey("user_id").
		Reference("users").
		On("id").
		OnDelete("cascade").
		OnUpdate("cascade")

	table.MustExec()
}

func (m *CreateWalletTable) Down(con *sqlx.DB) {
	mysql.DropTable("wallets", con).MustExec()
}
