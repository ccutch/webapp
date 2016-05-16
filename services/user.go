package services

import "github.com/ccutch/webapp/models"

// UserService service for accessing model and/or thirdparty api
type UserService struct{}

// GetUsers gets users from page and limit from datastore
func (u *UserService) GetUsers(page, limit int) []models.UserModel {
	// [TODO] Write to analytics
	// Fix data from user input to model requirements
	start := limit * (page - 1)
	end := limit * page
	return new(models.UserModel).GetBatch(start, end)
}

/* To sick to implement right now but you get the idea.
// GetPostsForUser get posts for user from id with page and limit
func (u *UserService) GetPostsForUser(id, page, limit int) []models.PostModel {
	// [TODO] Write to analytics
	// Fix data from user input to model requirements
	start := limit * (page - 1)
	end := limit * page
	user := new(models.UserModel).GetUser(id)
	if user.Id != id {
		return []models.PostModel{}
	}
	posts := new(models.PostModel).Query(map[string]string{
		"Owner": user.Name,
	}).Range(start, end).Execute()
  return posts
}
*/
