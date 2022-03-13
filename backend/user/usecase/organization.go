package usecase

import (
	"context"
	"log"
	pb "share-report/proto/user"
	"share-report/user/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetOrganizationById(ctx context.Context, req *pb.GetOrganizationByIdRequest) (*pb.GetOrganizationByIdResponse, error) {
	organization, err := s.dh.GetOrganizationByID(req.Id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Print("get organization by id")
	return &pb.GetOrganizationByIdResponse{
		Organization: s.oph.Organization(organization),
	}, nil
}

func (s *UserService) GetOrganizationListByUserId(ctx context.Context, req *pb.GetOrganizationListByUserIdRequest) (*pb.GetOrganizationListByUserIdResponse, error) {
	organizations, err := s.dh.GetOrganizationListByUserId(req.UserId)
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Print("get organization list by user id")
	return &pb.GetOrganizationListByUserIdResponse{
		Organizations: s.oph.Organizations(organizations),
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
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	message := ""
	if err == entity.ErrOrganizationNotExist {
		err = s.dh.CreateOrganization(&organization)
		message = "create organization"
	} else {
		err = s.dh.UpdateOrganization(&organization)
		message = "update organization"
	}
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Print(message)
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteOrganizationById(ctx context.Context, req *pb.DeleteOrganizationByIdRequest) (*emptypb.Empty, error) {
	err := s.dh.DeleteOrganizationByID(req.Id)
	if err != nil {
		log.Fatal(err.Error())
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	log.Print("delete organization by id")
	return &emptypb.Empty{}, nil
}
