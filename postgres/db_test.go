package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tonybka/go-base-persistence/config"
)

// Make sure you have run the postgresQL database from docker compose
func TestConnectDB(t *testing.T) {
	dbConfig := &config.DatabaseConfig{
		DBEndPoint: "localhost:54320",
		DBName:     "postgres",
		DBUserName: "postgres",
		DBPassword: "postgres",
	}
	assert.NotNil(t, dbConfig)

	err := ConnectDB(dbConfig)
	assert.NoError(t, err)
}
