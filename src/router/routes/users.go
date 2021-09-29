package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/followers",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowers,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/following",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowing,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/changePassword",
		Method:       http.MethodPost,
		Function:     controllers.ChangePassword,
		AuthRequired: true,
	},
}
