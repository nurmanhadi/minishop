package main

import (
	"minishop/external/database/mariadb"
	_ "minishop/internal/config"
)

func main() {
	db := mariadb.Connection()
	defer db.Close()
}
