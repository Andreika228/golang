package _struct

type Team struct {
	Person []Person `json:"person"`
}

type Person struct {
	Num        string `json:"num"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Position   string `json:"position"`
	Club       string `json:"club"`
	Height     string `json:"height"`
	KickLeg    string `json:"kick_leg"`
	Matches    string `json:"matches_for_country"`
	Goals      string `json:"goals"`
	Value      string `json:"value"`
}
