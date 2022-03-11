package usecase

import (
	"context"
	pb "share-report/proto/user"
	"share-report/user/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) GetOrganizationUserByUserIdAndOrganizationId(ctx context.Context, req *pb.GetOrganizationUserByUserIdAndOrganizationIdRequest) (*pb.GetOrganizationUserByUserIdAndOrganizationIdResponse, error) {
	return &pb.GetOrganizationUserByUserIdAndOrganizationIdResponse{}, nil
}

func (s *UserService) CreateOrganizationUser(ctx context.Context, req *pb.CreateOrganizationUserRequest) (*emptypb.Empty, error) {
	organizationUser := entity.OrganizationUser{
		UserID:         req.UserId,
		OrganizationID: req.OrganizationId,
	}
	_, err := s.dh.GetOrganizationByUserIDAndOrganizationID(organizationUser.UserID, organizationUser.OrganizationID)
	if err == nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, entity.ErrOrganizationUserExist.Error())
	}
	if err != nil && err != entity.ErrOrganizationUserNotExist {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.dh.CreateOrganizationUser(&organizationUser)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteOrganizationUserByUserId(ctx context.Context, req *pb.DeleteOrganizationUserByUserIdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *UserService) DeleteOrganizationUserByOrganizationId(ctx context.Context, req *pb.DeleteOrganizationUserByOrganizationIdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
