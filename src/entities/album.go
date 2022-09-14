package entities

type Album struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required,min:2"`
	Year string `json:"year"`
}
