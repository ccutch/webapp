package models

type PostMode struct {
	OwnerId string `json:"owner_id"`
	Content string `json:"content"`
	Owner   *User  `json:"owner"`
}
