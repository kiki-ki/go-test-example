package repository_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	status := m.Run()
	os.Exit(status)
}
