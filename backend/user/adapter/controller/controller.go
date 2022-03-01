package controller

type InputPortHandler interface {
	UserHandler
	GroupUserHandler
	GroupHandler
}

type UserHandler interface {
}

type GroupUserHandler interface {
}

type GroupHandler interface {
}

type inputPortHandler struct {
	UserHandler
	GroupUserHandler
	GroupHandler
}

func NewInputPortHandler() InputPortHandler {
	uh := NewUserHandler()
	guh := NewGroupUserHandler()
	gh := NewGroupHandler()
	return &inputPortHandler{uh, guh, gh}
}
