package configs

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	Init()
	fmt.Println(c.Host, c.Port, c.Username, c.Password, c.DBName)
	fmt.Println("Config test successfully")
}
