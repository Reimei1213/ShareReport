package usecase

import (
	"context"
	pb "share-report/proto/user"
)

func (s *UserService) GetOrganizationById(ctx context.Context, req *pb.GetOrganizationByIdRequest) (*pb.GetOrganizationByIdResponse, error) {
	return &pb.GetOrganizationByIdResponse{}, nil
}

func (s *UserService) CreateOrUpdateOrganization(ctx context.Context, req *pb.CreateOrUpdateOrganizationRequest) (*pb.CreateOrUpdateOrganizationResponse, error) {
	return &pb.CreateOrUpdateOrganizationResponse{}, nil
}

func (s *UserService) DeleteOrganizationById(ctx context.Context, req *pb.DeleteOrganizationByIdRequest) (*pb.DeleteOrganizationByIdResponse, error) {
	return &pb.DeleteOrganizationByIdResponse{}, nil
}
