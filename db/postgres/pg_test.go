package postgres

import "testing"

func TestPG(t *testing.T) {
	db := DBConn()
	t.Log(db)
}
