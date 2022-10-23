package domain

// Response represents the api response
type Response struct {
	Message string `json:"message"`
}

// InsertUser input  user
type IUser struct {
	User_name string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UpdateUser input  user
type UUser struct {
	Id        int32  `json:"id"`
	User_name string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

//List User
type Users struct {
	Id        int32  `json:"id"`
	User_name string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// User repository implementations

type UserRepository interface {
	CreateUser(user IUser) error
	UpdateUser(user UUser) error
	DeleteUser(id int64) error
	ListUsers() (*[]Users, error)
}
