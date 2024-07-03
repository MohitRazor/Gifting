package main

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Gifts    []int64 `json:"gifts"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
