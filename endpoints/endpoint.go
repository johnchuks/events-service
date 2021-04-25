package endpoints

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "github.com/johnchuks/events-service/service"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
    Create endpoint.Endpoint
}

// EventRequest struct holds the endpoint request definition for a create event
type CreateEventRequest struct {
	Email string `json:"email"`
	Message string `json:"message"`
	Environment string `json:"environment"`
	Component string `json:"component"`
	Data *service.MetaData `json:"data"`
}

// EventResponse struct holds the endpoint response definition for a create event
type CreateEventResponse struct {
    Message string `json:"message"`
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
    return Endpoints{
        Create: makeCreateEndpoint(s),
    }
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
    return func(ctx context.Context, r interface{}) (response interface{}, err error) {
        req := r.(CreateEventRequest)
        result, err := s.Create(ctx, req.Email, req.Component, req.Environment, req.Message, req.Data)
		if err != nil {
			return nil, err
		}
        return CreateEventResponse{Message: result}, nil
    }
}
