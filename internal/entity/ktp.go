package entity

type KTP struct {
	NoKTP         string `json:"no_ktp"`
	Name          string `json:"name"`
	BirthPlace    string `json:"birth_place"`
	BirthDate     string `json:"birth_date"`
	Gender        string `json:"gender"`
	Address       string `json:"address"`
	Religion      string `json:"religion"`
	MaritalStatus string `json:"marital_status"`
	Job           string `json:"job"`
	Nationality   string `json:"nationality"`
}
