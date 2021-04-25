package transport

import (
    "context"

    "github.com/go-kit/kit/log"
    gt "github.com/go-kit/kit/transport/grpc"
    "github.com/johnchuks/events-service/endpoints"
    "github.com/johnchuks/events-service/pb"
)

type gRPCServer struct {
	create gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.EventServiceServer {
    return &gRPCServer{
		create: gt.NewServer(
			endpoints.Create, 
			decodeCreateEventRequest,
			decodeCreateEventResponse,
		),
    }
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
    req := request.(*pb.E)
    return endpoints.MathReq{NumA: req.NumA, NumB: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
    resp := response.(endpoints.MathResp)
    return &pb.MathResponse{Result: resp.Result}, nil
}
