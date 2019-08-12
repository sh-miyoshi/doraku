package hobbydb

// HobbyDB is structure of Hobby
type HobbyDB struct {
	ID      int
	Name    string
	GroupNo int64
}

// InputValue strruct defines input value for recommendation
type InputValue struct {
	Outdoor bool
	Alone   bool
	Active  bool
}

// DBHandler is interface of dbHandler
type DBHandler interface {
	Initialize(connStr string) error
	GetRecommedHobby(input InputValue) (HobbyDB, error)
	GetHobbyNum() int
}
