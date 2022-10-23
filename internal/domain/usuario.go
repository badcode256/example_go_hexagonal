package domain



// Response represents the api response
type Response struct {
	Message string `json:"message"`
}

// CreateUser input create user
type IUser struct {
	User_name string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// User repository implementations

type UserRepository interface {
	CreateUser(user IUser) error
}
