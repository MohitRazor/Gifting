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

type Gift struct {
	ID    string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
