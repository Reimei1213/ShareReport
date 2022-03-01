package presenter

import (
	"github.com/golang/protobuf/ptypes"

	pb "share-report/proto/user"
	"share-report/user/entity"
)

type OutputPortHandler interface {
	User(*entity.User) *pb.User
}

type outputPortHandler struct {
}

func NewOutputPortHandler() OutputPortHandler {
	return &outputPortHandler{}
}

func (oph *outputPortHandler) User(user *entity.User) *pb.User {
	createdAt, _ := ptypes.TimestampProto(user.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(user.UpdatedAt)

	return &pb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Valid:     user.Valid,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
