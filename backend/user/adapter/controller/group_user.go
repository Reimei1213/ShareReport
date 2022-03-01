package controller

type groupUserHandler struct{}

func NewGroupUserHandler() GroupUserHandler {
	return &groupUserHandler{}
}
