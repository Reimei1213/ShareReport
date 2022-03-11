package usecase

import (
	"context"
	pb "share-report/proto/user"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetOrganizationUserByUserIdAndOrganizationId(ctx context.Context, req *pb.GetOrganizationUserByUserIdAndOrganizationIdRequest) (*pb.GetOrganizationUserByUserIdAndOrganizationIdResponse, error) {
	return &pb.GetOrganizationUserByUserIdAndOrganizationIdResponse{}, nil
}

func (s *UserService) CreateOrganizationUser(ctx context.Context, req *pb.CreateOrganizationUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteOrganizationUserByUserId(ctx context.Context, req *pb.DeleteOrganizationUserByUserIdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteOrganizationUserByOrganizationId(ctx context.Context, req *pb.DeleteOrganizationUserByOrganizationIdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
