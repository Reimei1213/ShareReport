package usecase

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"share-report/user/entity"
	pb "share-report/proto/user"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	return &pb.GetUserByIdResponse{}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	user := entity.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}
	err := s.dh.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}

func (s *UserService) DeleteUserById(ctx context.Context, req *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	return &pb.DeleteUserByIdResponse{}, nil
}