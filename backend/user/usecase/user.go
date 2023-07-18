package usecase

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "share-report/proto/user"
	"share-report/user/entity"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := s.dh.GetUserByID(req.Id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Print("get user by id")
	return &pb.GetUserByIdResponse{
		User: s.oph.User(user),
	}, nil
}

func (s *UserService) GetUserIdByEmailAndPassword(ctx context.Context, req *pb.GetUserIdByEmailAndPasswordRequest) (*pb.GetUserIdByEmailAndPasswordResponse, error) {
	return nil, nil
}

func (s *UserService) GetUserListByOrganizationId(ctx context.Context, req *pb.GetUserListByOrganizationIdRequest) (*pb.GetUserListByOrganizationIdResponse, error) {
	users, err := s.dh.GetUserListByOrganizationId(req.OrganizationId)
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Print("get user list by organization id")
	return &pb.GetUserListByOrganizationIdResponse{
		Users: s.oph.Users(users),
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
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	message := ""
	if err == entity.ErrUserNotExist {
		err = s.dh.CreateUser(&user)
		message = "create user"
	} else {
		err = s.dh.UpdateUser(&user)
		message = "update user"
	}
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Print(message)
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteUserById(ctx context.Context, req *pb.DeleteUserByIdRequest) (*emptypb.Empty, error) {
	err := s.dh.DeleteUserByID(req.Id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Print("delete user by id")
	return &emptypb.Empty{}, nil
}
