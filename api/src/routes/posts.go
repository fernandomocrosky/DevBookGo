package routes

import (
	"net/http"

	"github.com/fernandomocrosky/DevBookGo/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:           "/posts",
		Method:        http.MethodPost,
		HandleFunc:    controllers.CreatePost,
		Authenticated: true,
	},
	{
		URI:           "/posts",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetPosts,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetPostbyID,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}",
		Method:        http.MethodPut,
		HandleFunc:    controllers.UpdatePost,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}",
		Method:        http.MethodDelete,
		HandleFunc:    controllers.DeletePost,
		Authenticated: true,
	},
	{
		URI:           "/users/{userId}/posts",
		Method:        http.MethodGet,
		HandleFunc:    controllers.GetPostsByUserId,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}/like",
		Method:        http.MethodPost,
		HandleFunc:    controllers.LikePost,
		Authenticated: true,
	},
	{
		URI:           "/posts/{postId}/unlike",
		Method:        http.MethodPost,
		HandleFunc:    controllers.UnlikePost,
		Authenticated: true,
	},
}
