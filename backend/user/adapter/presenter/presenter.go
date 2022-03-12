package presenter

import (
	"github.com/golang/protobuf/ptypes"

	pb "share-report/proto/user"
	"share-report/user/entity"
)

type OutputPortHandler interface {
	User(user *entity.User) *pb.User
	Users(users []*entity.User) []*pb.User
	Organization(organization *entity.Organization) *pb.Organization
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

func (oph *outputPortHandler) Users(users []*entity.User) []*pb.User {
	var pbUsers []*pb.User
	for _, user := range(users) {
		createdAt, _ := ptypes.TimestampProto(user.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(user.UpdatedAt)
		pbUser := &pb.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			Valid:     user.Valid,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
		pbUsers = append(pbUsers, pbUser)
	}
	return pbUsers
}

func (oph *outputPortHandler) Organization(organization *entity.Organization) *pb.Organization {
	createdAt, _ := ptypes.TimestampProto(organization.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(organization.UpdatedAt)

	return &pb.Organization{
		Id:        organization.ID,
		UserId:    organization.UserID,
		Name:      organization.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
