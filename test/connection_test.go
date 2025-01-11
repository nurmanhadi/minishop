package test

import (
	"minishop/external/database/mariadb"
	_ "minishop/internal/config"
	"testing"
)

func TestConnection(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()
}
