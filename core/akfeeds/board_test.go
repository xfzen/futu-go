package akfeeds

import (
	"testing"
)

func TestGetShBoard(t *testing.T) {
	initDB()

	params := map[string]string{}

	// 沪A
	GetSHBoard(params)

	// 深A
	GetSZBoard(params)
}
