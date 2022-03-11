package usecase

import (
	"context"
	pb "share-report/proto/user"
	"share-report/user/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetOrganizationById(ctx context.Context, req *pb.GetOrganizationByIdRequest) (*pb.GetOrganizationByIdResponse, error) {
	organization, err := s.dh.GetOrganizationByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.GetOrganizationByIdResponse{
		Organization: s.oph.Organization(organization),
	}, nil
}

func (s *UserService) CreateOrUpdateOrganization(ctx context.Context, req *pb.CreateOrUpdateOrganizationRequest) (*emptypb.Empty, error) {
	organization := entity.Organization{
		ID:     req.Id,
		UserID: req.UserId,
		Name:   req.Name,
	}
	_, err := s.dh.GetOrganizationByID(organization.ID)
	if err != nil && err != entity.ErrOrganizationNotExist {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err == entity.ErrOrganizationNotExist {
		err = s.dh.CreateOrganization(&organization)
	} else {
		err = s.dh.UpdateOrganization(&organization)
	}
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteOrganizationById(ctx context.Context, req *pb.DeleteOrganizationByIdRequest) (*emptypb.Empty, error) {
	err := s.dh.DeleteOrganizationByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}
