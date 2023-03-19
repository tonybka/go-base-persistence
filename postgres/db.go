package postgres

import (
	"fmt"
	"strings"

	"github.com/tonybka/go-base-persistence/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

// ConnectDB connects to postgresQL database and create singleton DB connection
func ConnectDB(configs *config.DatabaseConfig) error {
	var err error

	urls := strings.Split(configs.DBEndPoint, ":")
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v sslmode=disable", urls[0], urls[1], configs.DBName)

	if configs.DBUserName != "" {
		dsn += fmt.Sprintf(" user=%v", configs.DBUserName)
	}

	if configs.DBPassword != "" {
		dsn += fmt.Sprintf(" password=%v", configs.DBPassword)
	}

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Printf("connected to database: %s\n", configs.DBEndPoint)
	return nil
}

// CloseDBConn closes given database connection
func CloseDBConn(db *gorm.DB) error {
	dbSQL, err := db.DB()
	if err != nil {
		return err
	}
	dbSQL.Close()

	return nil
}
