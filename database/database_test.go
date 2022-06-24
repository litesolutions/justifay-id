package database_test

import (
	"testing"

	"github.com/litesolutions/justifay-id/config"
	"github.com/litesolutions/justifay-id/database"
	"github.com/stretchr/testify/assert"
)

func TestNewDatabaseTypeNotSupported(t *testing.T) {
	cnf := &config.Config{
		Database: config.DatabaseConfig{
			PSN: "bogus",
		},
	}
	_, err := database.NewDatabase(cnf)

	assert.NotNil(t, err)
}
