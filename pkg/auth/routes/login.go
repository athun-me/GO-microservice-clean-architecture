package routes

type LoginRequestBody struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func 