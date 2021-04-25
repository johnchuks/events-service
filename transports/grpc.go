package transport

import (
    "context"

    gt "github.com/go-kit/kit/transport/grpc"
    "github.com/johnchuks/events-service/endpoints"
    "github.com/johnchuks/events-service/pb"
)

type gRPCServer struct {
	create gt.Handler
	pb.UnimplementedEventServiceServer
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints) pb.EventServiceServer {
    return &gRPCServer{
		create: gt.NewServer(
			endpoints.Create, 
			decodeCreateEventRequest,
			encodeCreateEventResponse,
		),
    }
}

func (s *gRPCServer) Create(ctx context.Context, req *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
    _, resp, err := s.create.ServeGRPC(ctx, req)
    if err != nil {
        return nil, err
    }
    return resp.(*pb.CreateEventResponse), nil
}


func decodeCreateEventRequest(_ context.Context, request interface{}) (interface{}, error) {
    req := request.(*pb.CreateEventRequest)
    return endpoints.CreateEventRequest{
		Email: req.Email, 
		Component: req.Component,
		Environment: req.Environment,
		Message: req.Message,
		Data: req.Data,
	}, nil
}

func encodeCreateEventResponse(_ context.Context, response interface{}) (interface{}, error) {
    resp := response.(endpoints.CreateEventResponse)
    return &pb.CreateEventResponse{
		Id: int64(resp.ID),
		Email: resp.Email,
		Component: resp.Component,
		Environment: resp.Environment,
		Message: resp.Message,
		Data: resp.Data,
		CreatedAt: resp.CreatedAt.String(),
	}, nil
}
