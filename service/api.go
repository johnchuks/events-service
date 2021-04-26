package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/johnchuks/events-service/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type service struct {
	DB *gorm.DB
}

// Service interface describes a service that adds numbers
type Service interface {
	Create(ctx context.Context, email, component, environment, message string, data map[string]string) (*models.Event, error)
	Retrieve(ctx context.Context, filter map[string]interface{}) ([]*models.Event, error)
}

// NewService returns a Service with all of the expected dependencies
func NewService(host, port, user, password, dbname string) Service {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("An error occurred:", err)
	} else {
		log.Printf("Successfully connected to database %s", dbname)
	}

	db.Debug().AutoMigrate(&models.Event{}) // database migration

	return &service{
		DB: db,
	}
}

// Create func implements Service interface
func (s service) Create(ctx context.Context, email, component, environment, message string, data map[string]string) (*models.Event, error) {
	val, _ := json.Marshal(data)
	e := &models.Event{
		Email:       email,
		Component:   component,
		Environment: environment,
		Message:     message,
		Data:        datatypes.JSON([]byte(val)),
	}
	event, err := e.Create(s.DB)

	if err != nil {
		return nil, err
	}
	return event, nil
}

// Retrieve func implements Service interface Retrieve Method
func (s service) Retrieve(ctx context.Context, filters map[string]interface{}) ([]*models.Event, error) {
	e := &models.Event{}
	events, err := e.Retrieve(s.DB, filters)

	if err != nil {
		return nil, err
	}

	return events, nil
}
