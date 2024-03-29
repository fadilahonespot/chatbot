package router

import (
	"net/http"

	"github.com/fadilahonespot/chatbot/server/handler"
	"github.com/fadilahonespot/chatbot/server/middleware"
)

type Router struct {
	userHandler *handler.UserHandler
	chatHandler *handler.ChatHandler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) SetUserHandler(handler *handler.UserHandler) *Router {
	r.userHandler = handler
	return r
}

func (r *Router) SetChatHandler(handler *handler.ChatHandler) *Router {
	r.chatHandler = handler
	return r
}

func (r *Router) Validate() *Router {
	if r.userHandler == nil {
		panic("user handler is nil")
	}

	if r.chatHandler == nil {
		panic("chat handler is nil")
	}

	return r
}

// SetupRouter sets up the router
func (r *Router) SetupRouter() {
	// Register route for registering a new user
	http.HandleFunc("/register", middleware.SetLoggerMiddleware(r.userHandler.Register))
	// Register route for logging in a user
	http.HandleFunc("/login", middleware.SetLoggerMiddleware(r.userHandler.Login))

	// Register route for handling chat requests
	http.Handle("/chat", middleware.SetLoggerMiddleware(middleware.JwtMiddleware(r.chatHandler.Chat)))
}
