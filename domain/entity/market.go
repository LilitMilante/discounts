package entity

type Market struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	DirectorID int64  `json:"director_id"`
}
