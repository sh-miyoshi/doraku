package hobbyapi

// Hobby struct defines info of hobby details
type Hobby struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	NameEN      string `json:"nameEN"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Group       int    `json:"group"`
}

// HobbyKey struct only includes id and name for vitualization
type HobbyKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
