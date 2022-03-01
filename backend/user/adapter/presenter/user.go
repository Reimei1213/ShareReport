package presenter

type userHandler struct {
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}
