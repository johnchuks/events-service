package transport

import (
	"context"

	gkit "github.com/go-kit/kit/transport/grpc"
	"github.com/johnchuks/events-service/endpoints"
	"github.com/johnchuks/events-service/pb"
)

type gRPCServer struct {
	create   gkit.Handler
	retrieve gkit.Handler
	pb.UnimplementedEventServiceServer
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints) pb.EventServiceServer {
	return &gRPCServer{
		create: gkit.NewServer(
			endpoints.Create,
			decodeCreateEventRequest,
			encodeCreateEventResponse,
		),
		retrieve: gkit.NewServer(
			endpoints.Retrieve,
			decodeRetrieveEventRequest,
			encodeRetrieveEventResponse,
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

func (s *gRPCServer) Retrieve(ctx context.Context, req *pb.RetrieveEventRequest) (*pb.ListEventResponse, error) {
	_, resp, err := s.retrieve.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ListEventResponse), nil
}

func decodeCreateEventRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateEventRequest)
	return endpoints.CreateEventRequest{
		Email:       req.Email,
		Component:   req.Component,
		Environment: req.Environment,
		Message:     req.Message,
		Data:        req.Data,
	}, nil
}

func encodeCreateEventResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.EventResponse)
	return &pb.CreateEventResponse{
		Id:          int64(resp.ID),
		Email:       resp.Email,
		Component:   resp.Component,
		Environment: resp.Environment,
		Message:     resp.Message,
		Data:        resp.Data,
		CreatedAt:   resp.CreatedAt.String(),
	}, nil
}

func decodeRetrieveEventRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.RetrieveEventRequest)
	return endpoints.RetrieveEventRequest{
		Email:       derefString(req.Email),
		Component:   derefString(req.Component),
		Environment: derefString(req.Environment),
		Text:        derefString(req.Text),
		Date:        derefString(req.Date),
	}, nil
}

func encodeRetrieveEventResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListEventResponse)
	var events []*pb.CreateEventResponse

	for _, r := range resp.Events {
		e := &pb.CreateEventResponse{
			Id:          int64(r.ID),
			Email:       r.Email,
			Component:   r.Component,
			Environment: r.Environment,
			Message:     r.Message,
			Data:        r.Data,
			CreatedAt:   r.CreatedAt.String(),
		}
		events = append(events, e)
	}
	r := &pb.ListEventResponse{
		Events: events,
	}
	return r, nil
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
