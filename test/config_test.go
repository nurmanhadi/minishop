package test

import (
	"fmt"
	_ "minishop/internal/config"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println("test")
}
func TestXxx(t *testing.T) {
	query := "update set"
	key := ""
	v1 := "hehe"
	v2 := 8
	if v1 != "" {
		query += " nama = ?,"
		key += v1
	}
	if v2 != 0 {
		query += " age = ?"
		key += "wkwk"
	}
	fmt.Println(query)
	fmt.Println(key)
}
