package handler

import "testing"

func TestUser(t *testing.T) {
	user := new(User)
	user.UserName = "hello"
	t.Log(user)
}
