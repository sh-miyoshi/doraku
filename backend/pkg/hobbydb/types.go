package hobbydb

// HobbyDB is structure of Hobby in Mongo DB
type HobbyDB struct {
	ID      int    `bson:"id"`
	Name    string `bson:"name"`
	NameEN  string `bson:"nameEN"`
	GroupNo int    `bson:"groupNo"`
}
