package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationsRoutes = []Route{
	{
		URI:          "/publications",
		Method:       http.MethodPost,
		Function:     controllers.CreatePublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications",
		Method:       http.MethodGet,
		Function:     controllers.GetPublications,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationID}",
		Method:       http.MethodGet,
		Function:     controllers.GetPublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePublication,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/publications",
		Method:       http.MethodGet,
		Function:     controllers.GetPublicationByUser,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationID}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePublication,
		AuthRequired: true,
	},
	{
		URI:          "/publications/{publicationID}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnlikePublication,
		AuthRequired: true,
	},
}
