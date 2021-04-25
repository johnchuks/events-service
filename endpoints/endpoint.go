package endpoints

import (
	"strings"
	"fmt"
	"time"
    "context"
	"reflect"
	"encoding/json"
    "github.com/go-kit/kit/endpoint"
    "github.com/johnchuks/events-service/service"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
    Create endpoint.Endpoint
	Retrieve endpoint.Endpoint
}

// CreateEventRequest struct holds the endpoint request definition for a create event
type CreateEventRequest struct {
	Email string `json:"email"`
	Message string `json:"message"`
	Environment string `json:"environment"`
	Component string `json:"component"`
	Data map[string]string `json:"data"`
}

// CreareEventResponse struct holds the endpoint response definition for a create event
type EventResponse struct {
	ID uint `json:"id"`
	Email string `json:"email"`
    Message string `json:"message"`
	Environment string `json:"environment"`
	Component string `json:"component"`
	Data map[string]string `json:"data"`
	CreatedAt time.Time `json:"createdAt"`
}

type RetrieveEventRequest struct {
	Email string `json:"email"`
	Text string `json:"text"`
	Environment string `json:"environment"`
	Component string `json:"component"`
	Date string `json:"date"`
}

type ListEventResponse struct {
	Events []EventResponse `json:"events"`
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s service.Service) Endpoints {
    return Endpoints{
        Create: makeCreateEndpoint(s),
		Retrieve: makeRetrieveEndpoint(s),
    }
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
    return func(ctx context.Context, r interface{}) (response interface{}, err error)  {
        req := r.(CreateEventRequest)

        result, err := s.Create(ctx, req.Email, req.Component, req.Environment, req.Message, req.Data)
		if err != nil {
			return nil, err
		}
		var data map[string]string
		_ = json.Unmarshal([]byte(result.Data.String()), &data)

        return EventResponse{
			ID: result.ID,
			Email: result.Email,
			Message: result.Message,
			Environment: result.Environment,
			Component: result.Environment,
			Data: data,
			CreatedAt: result.CreatedAt,
		}, nil
    }
}

func makeRetrieveEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(RetrieveEventRequest)

		v := reflect.ValueOf(req)
		filter := make(map[string]interface{})

		typeOfR := v.Type()

		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Interface() != "" {
				key := strings.ToLower(typeOfR.Field(i).Name)
				filter[key] = v.Field(i).Interface()
			}
		}
		fmt.Println(filter, "====>>S>>S>")
		results, err := s.Retrieve(ctx, filter)
		var events []EventResponse
		
		for _, event := range results {
			var data map[string]string
			_ = json.Unmarshal([]byte(event.Data.String()), &data)

			e := EventResponse{
				ID: event.ID,
				Email: event.Email,
				Message: event.Message,
				Environment: event.Environment,
				Component: event.Component,
				Data: data,
				CreatedAt: event.CreatedAt,
			}
			events = append(events, e)
		}
		resp := ListEventResponse{
			Events: events,
		}
		return resp, nil
	}
}
