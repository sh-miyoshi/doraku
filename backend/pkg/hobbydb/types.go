package hobbydb

// HobbyDB is structure of Hobby
type HobbyDB struct {
	ID          int
	Name        string
	NameEN      string
	GroupNo     int64
	Description string
}

// InputValue strruct defines input value for recommendation
type InputValue struct {
	Outdoor bool
	Alone   bool
	Active  bool
}
