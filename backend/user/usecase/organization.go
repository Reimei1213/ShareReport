package usecase

import (
	"context"
	pb "share-report/proto/user"
)

func (s *UserService) GetOrganizationById(ctx context.Context, req *pb.GetOrganizationByIdRequest) (*pb.GetOrganizationByIdResponse, error) {
	return &pb.GetOrganizationByIdResponse{}, nil
}

func (s *UserService) CreateOrganization(ctx context.Context, req *pb.CreateOrganizationRequest) (*pb.CreateOrganizationResponse, error) {
	return &pb.CreateOrganizationResponse{}, nil
}

func (s *UserService) UpdateOrganization(ctx context.Context, req *pb.UpdateOrganizationRequest) (*pb.UpdateOrganizationResponse, error) {
	return &pb.UpdateOrganizationResponse{}, nil
}

func (s *UserService) DeleteOrganizationById(ctx context.Context, req *pb.DeleteOrganizationByIdRequest) (*pb.DeleteOrganizationByIdResponse, error) {
	return &pb.DeleteOrganizationByIdResponse{}, nil
}
