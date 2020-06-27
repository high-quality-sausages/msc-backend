package models

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   int    `json:"gender"`
	Role     int    `json:"role"` // 0: Admin; 1: Normal; 2: Visitor
}
