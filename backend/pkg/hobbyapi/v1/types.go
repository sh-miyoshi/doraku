package hobbyapi

//----------------------------
// Response Data
//----------------------------

// HobbyInfo struct is return value of hobby API
type HobbyInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
