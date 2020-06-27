package users

import (
	"testing"
)

func TestQueryByID(t *testing.T) {
	user := new(User)
	info, err := user.QuerySongByID("w001")
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestQueryByName(t *testing.T) {
	user := new(User)
	info, err := user.QuerySongsByName("dddd")
	if err != nil {
		t.Error(err)
	}

	if len(info) != 0 {
		t.Log("existes")
	} else {
		t.Log("does not exist")
	}

	t.Log(info)
}
