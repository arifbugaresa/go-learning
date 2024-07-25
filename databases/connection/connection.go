package connection

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	DBConnections    *gorm.DB
	SqlDBConnections *sql.DB
	err              error
)

func Initiator() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		viper.GetString("migration.db.postgres.db_host"),
		viper.GetInt("migration.db.postgres.db_port"),
		viper.GetString("migration.db.postgres.db_user"),
		viper.GetString("migration.db.postgres.db_password"),
		viper.GetString("migration.db.postgres.db_name"),
	)

	DBConnections, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// check connection
	SqlDBConnections, err = DBConnections.DB()
	if err != nil {
		panic(err)
	}

	err = SqlDBConnections.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
