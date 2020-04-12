package mapper

import "ivar-go/src/models"

func UserToGetUserResponse(source models.User, dest models.GetUserResponse) models.GetUserResponse {

	dest.FollowerCount = len(source.Followers)
	dest.FollowingCount = len(source.Following)
	dest.LastName = source.LastName
	dest.FirstName = source.FirstName
	dest.Email = source.Email
	dest.CreatedAt = source.CreatedAt
	dest.UpdatedAt = source.UpdatedAt

	return dest
}

func UserToFollowerResponse(source models.User, dest models.GetFollowersResponse) models.GetFollowersResponse {

	dest.FirstName = source.FirstName
	dest.LastName = source.LastName

	return dest
}
