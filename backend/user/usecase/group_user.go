package usecase

import (
	"context"
	pb "share-report/proto/user"
)

func (s *UserService) GetGroupUserByUserIdAndGroupId(ctx context.Context, req *pb.GetGroupUserByUserIdAndGroupIdRequest) (*pb.GetGroupUserByUserIdAndGroupIdResponse, error) {
	return &pb.GetGroupUserByUserIdAndGroupIdResponse{}, nil
}

func (s *UserService) CreateGroupUser(ctx context.Context, req *pb.CreateGroupUserRequest) (*pb.CreateGroupUserResponse, error) {
	return &pb.CreateGroupUserResponse{}, nil
}

func (s *UserService) DeleteGroupUserByUserId(ctx context.Context, req *pb.DeleteGroupUserByUserIdRequest) (*pb.DeleteGroupUserByUserIdResponse, error) {
	return &pb.DeleteGroupUserByUserIdResponse{}, nil
}

func (s *UserService) DeleteGroupUserByGroupId(ctx context.Context, req *pb.DeleteGroupUserByGroupIdRequest) (*pb.DeleteGroupUserByGroupIdResponse, error) {
	return &pb.DeleteGroupUserByGroupIdResponse{}, nil
}
