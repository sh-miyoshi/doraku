package hobbyapi

//----------------------------
// Request Data
//----------------------------

// RecommendRequest struct is request data for recommendation
type RecommendRequest struct {
	Outdoor bool `json:"outdoor"`
	Alone   bool `json:"alone"`
	Active  bool `json:"active"`
}

//----------------------------
// Response Data
//----------------------------

// HobbyInfo struct is return value of hobby API
type HobbyInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
