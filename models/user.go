package models

type UserModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// GetBatch gets list of users starting at startIndex and going to lastIndex or end of list.
func (u *UserModel) GetBatch(startIndex, lastIndex int) []UserModel {
	return []models.UserModel{
		UserModel{Id: 1, Name: "Connor"},
		UserModel{Id: 2, Name: "Alex"},
	}
}
