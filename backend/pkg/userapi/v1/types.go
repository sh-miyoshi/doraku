package userapi

//----------------------------
// Request Data
//----------------------------

// LoginRequest struct defines request data for user login
type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// UserCreateRequest struct defines request data for user create
type UserCreateRequest struct {
	EMail string `json:"email"`
}

//----------------------------
// Response Data
//----------------------------

// LoginResponse struct defines response data for user login
type LoginResponse struct {
	Token string `json:"token"`
}

// User struct defines info of user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	EMail string `json:"email"`
}
