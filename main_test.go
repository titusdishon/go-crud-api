package main_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("failed to load env files")
	}
	code := m.Run()
	os.Exit(code)
}
