package userapi

//----------------------------
// Request Data
//----------------------------

// LoginRequest struct defines request data for user login
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserCreateRequest struct defines request data for user create
type UserCreateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// RegisterUserRequest struct defines request data for validation of user create
type RegisterUserRequest struct {
	Token string `json:"token"`
}

//----------------------------
// Response Data
//----------------------------

// LoginResponse struct defines response data for user login
type LoginResponse struct {
	Token string `json:"token"`
}

// UserCreateResponse struct defines response data for user create
type UserCreateResponse struct {
	Token string `json:"token"`
}

// User struct defines info of user
type User struct {
	Name        string   `json:"name"`
	MyHobbyList []string `json:"myhobbies"`
	// TODO(place, ...)
}
