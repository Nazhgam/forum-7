package server

import (
	router "forum/trans/custom_router"
)

func (s Server) Routers(r *router.Router) {
	r.Get("/home", s.handler.Home)

	r.Get("/logout", s.handler.LogOut)
	r.Post("/signup", s.handler.SignUp)
	r.Post("/login", s.handler.LogIn)

	r.Post("/comment", s.handler.CreateComment)
	r.Post("/comment/like", s.handler.AddEmotionToComment)
	r.Delete("/comment", s.handler.DeleteCommentByID)

	r.Post("/post", s.handler.CreatePost)
	r.Put("/post", s.handler.UpdatePost)
	r.Get("/post", s.handler.GetPostByID)
	r.Get("/post/most_liked", s.handler.MostLikedPost)
	r.Get("/post/categ", s.handler.MostLikedCategory)
	r.Post("/post/like", s.handler.AddEmotionToPost)
	r.Delete("/post", s.handler.Delete)

	// user
}
