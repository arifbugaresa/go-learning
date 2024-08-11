package connection

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func Initiator() (dbConnection *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("migration.db.postgres.db_host"),
		viper.GetInt("migration.db.postgres.db_port"),
		viper.GetString("migration.db.postgres.db_user"),
		viper.GetString("migration.db.postgres.db_password"),
		viper.GetString("migration.db.postgres.db_name"),
	)

	dbConnection, err = sql.Open("postgres", dsn)

	// check connection
	err = dbConnection.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	return
}
