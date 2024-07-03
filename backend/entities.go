package main

type User struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Gifts    []string `json:"gifts"`
}

type Gift struct {
	ID    string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
