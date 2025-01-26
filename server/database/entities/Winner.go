package entities

type Winner struct {
	ID          string `json:"id"`
	Season      string `json:"season"`
	Game        string `json:"game"`
	Position    string `json:"position"`
	TeamMember1 string `json:"teamMember1"`
	TeamMember2 string `json:"teamMember2"`
}
