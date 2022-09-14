package entities

type Playlist struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" `
	Owner string `json:"userID"`
}
