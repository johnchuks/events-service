package service

import (
  "context"
  "github.com/go-kit/kit/log"
)

type service struct {
  logger log.Logger
}

type MetaData struct {

}

// Service interface describes a service that adds numbers
type Service interface {
  Create(ctx context.Context, email, component, environment, message string, data *MetaData) (string, error)
}

// NewService returns a Service with all of the expected dependencies
func NewService(logger log.Logger) Service {
  return &service{
    logger: logger,
  }
}

// Create func implements Service interface
func (s service) Create(ctx context.Context, email, component, environment, message string, data *MetaData) (string, error) {
  return "Event added successfully", nil
}
