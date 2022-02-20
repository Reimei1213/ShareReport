package usecase

import(
	"context"
	pb "share-report/proto/user"
)

func (s *UserService) GetGroupById(ctx context.Context, req *pb.GetGroupByIdRequest) (*pb.GetGroupByIdResponse, error) {
	return &pb.GetGroupByIdResponse{}, nil
}

func (s *UserService) CreateGroup(ctx context.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	return &pb.CreateGroupResponse{}, nil
}

func (s *UserService) UpdateGroup(ctx context.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupResponse, error) {
	return &pb.UpdateGroupResponse{}, nil
}

func (s *UserService) DeleteGroupById(ctx context.Context, req *pb.DeleteGroupByIdRequest) (*pb.DeleteGroupByIdResponse, error) {
	return &pb.DeleteGroupByIdResponse{}, nil
}