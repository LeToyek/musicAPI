package entities

type Song struct {
	ID        string `json:"id" validate:"required"`
	Year      int    `json:"year"`
	Performer string `json:"performer"`
	Genre     string `json:"genre"`
	Duration  string `json:"duration"`
	Album_id  string `json:"album_id"`
}
