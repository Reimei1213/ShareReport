package usecase

import (
	"context"
	pb "share-report/proto/user"
)

func (s *UserService) GetOrganizationUserByUserIdAndOrganizationId(ctx context.Context, req *pb.GetOrganizationUserByUserIdAndOrganizationIdRequest) (*pb.GetOrganizationUserByUserIdAndOrganizationIdResponse, error) {
	return &pb.GetOrganizationUserByUserIdAndOrganizationIdResponse{}, nil
}

func (s *UserService) CreateOrganizationUser(ctx context.Context, req *pb.CreateOrganizationUserRequest) (*pb.CreateOrganizationUserResponse, error) {
	return &pb.CreateOrganizationUserResponse{}, nil
}

func (s *UserService) DeleteOrganizationUserByUserId(ctx context.Context, req *pb.DeleteOrganizationUserByUserIdRequest) (*pb.DeleteOrganizationUserByUserIdResponse, error) {
	return &pb.DeleteOrganizationUserByUserIdResponse{}, nil
}

func (s *UserService) DeleteOrganizationUserByOrganizationId(ctx context.Context, req *pb.DeleteOrganizationUserByOrganizationIdRequest) (*pb.DeleteOrganizationUserByOrganizationIdResponse, error) {
	return &pb.DeleteOrganizationUserByOrganizationIdResponse{}, nil
}
