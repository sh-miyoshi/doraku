package hobbyapi

// Hobby struct defines info of hobby
type Hobby struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	NameEN      string `json:"nameEN"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Group       int    `json:"group"`
}
