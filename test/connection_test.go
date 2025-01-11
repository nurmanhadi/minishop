package test

import (
	"context"
	"minishop/external/database/mariadb"
	_ "minishop/internal/config"
	"testing"
)

func TestConnection(t *testing.T) {
	db := mariadb.Connection()
	defer db.Close()

	_, err := db.ExecContext(context.Background(), "INSERT INTO person(name, age) VALUES(?,?)", "nurman", 23)
	if err != nil {
		panic(err)
	}
}
