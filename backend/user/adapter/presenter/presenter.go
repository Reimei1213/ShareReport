package presenter

type OutputPortHandler interface {
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

type outputPortHandler struct {
	UserHandler
	GroupUserHandler
	GroupHandler
}

func NewOutputPortHandler() OutputPortHandler {
	uh := NewUserHandler()
	guh := NewGroupUserHandler()
	gh := NewGroupHandler()
	return &outputPortHandler{uh, guh, gh}
}
