package usecase

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "share-report/proto/user"
	"share-report/user/entity"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := s.dh.GetUserByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetUserByIdResponse{
		User: s.oph.User(user),
	}, nil
}

func (s *UserService) CreateOrUpdateUser(ctx context.Context, req *pb.CreateOrUpdateUserRequest) (*emptypb.Empty, error) {
	user := entity.User{
		ID:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	_, err := s.dh.GetUserByID(user.ID)
	if err != nil && err != entity.ErrUserNotExist {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err == entity.ErrUserNotExist {
		err = s.dh.CreateUser(&user)
	} else {
		err = s.dh.UpdateUser(&user)
	}
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteUserById(ctx context.Context, req *pb.DeleteUserByIdRequest) (*emptypb.Empty, error) {
	err := s.dh.DeleteUserByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
