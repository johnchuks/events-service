package service

import (
    "fmt"
    // "reflect"
    "encoding/json"
    "context"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
    "github.com/johnchuks/events-service/models"
    "gorm.io/driver/postgres"
    "gorm.io/datatypes"
)

type service struct {
    DB *gorm.DB
}


// Service interface describes a service that adds numbers
type Service interface {
    Create(ctx context.Context, email, component, environment, message string, data map[string]string) (*models.Event, error)
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
func (s *service) Create(ctx context.Context, email, component, environment, message string, data map[string]string) (*models.Event, error) {
    val, _ := json.Marshal(data)
    e := &models.Event{
        Email: email,
        Component: component,
        Environment: environment,
        Message: message,
        Data: datatypes.JSON([]byte(val)),
    }
    event, err := e.Create(s.DB)

    if err != nil {
        return nil, err
    }
    fmt.Println(event, "====EVVVETTTTHHHFF>>>>>")
    return event, nil
}
