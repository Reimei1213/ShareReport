package controller

type groupHandler struct{}

func NewGroupHandler() GroupHandler {
	return &groupHandler{}
}
