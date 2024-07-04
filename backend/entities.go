package main

type User struct {
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Gifts     []string `json:"gifts"`
	Interests []string `json:"interests"`
}

type Gift struct {
	ID    string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Link  string  `json:"link"`
	Image string  `json:"image"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
