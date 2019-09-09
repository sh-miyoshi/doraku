package hobbydb

// DBHandler is interface of dbHandler
type DBHandler interface {
	Initialize(connStr string) error
	GetRecommendHobby(input InputValue) (HobbyDB, error)
	GetHobbyNum() int
	GetHobbyByID(id int) (HobbyDB, error)
}

var inst = localDBHandler{}

// GetInst return instance of Database Handler
func GetInst() DBHandler {
	return &inst
}
