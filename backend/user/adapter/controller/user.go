package controller

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}
