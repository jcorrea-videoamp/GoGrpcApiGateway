package service

import (
	"context"

	proto "github.com/jcorrea-videoamp/GoGrpcApiGateway/project/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type DummyService interface {
	proto.DummyServiceServer
}

type dummyService struct {
	proto.UnimplementedDummyServiceServer
	users []*proto.UserResponse
}

func NewDummyService() (*dummyService, error) {
	return &dummyService{
		users: []*proto.UserResponse{
			{
				User: &proto.User{
					Id:   int32(1),
					Name: "Juan",
					Role: "Dev",
				},
			},
			{
				User: &proto.User{Id: int32(2),
					Name: "Crisi",
					Role: "Dev",
				},
			},
			{
				User: &proto.User{Id: int32(3),
					Name: "Carl",
					Role: "QA",
				},
			},
		},
	}, nil
}

func (ds *dummyService) ListUsers(e *emptypb.Empty, stream proto.DummyService_ListUsersServer) error {
	for _, user := range ds.users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (ds *dummyService) GetUser(ctx context.Context, req *proto.GetRequest) (*proto.UserResponse, error) {
	user := ds.users[req.Id]
	return user, nil

}

func (ds *dummyService) CreateUser(ctx context.Context, req *proto.CreateRequest) (*proto.AcknowlegeResponse, error) {
	return &proto.AcknowlegeResponse{}, nil
}
